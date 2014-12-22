import java.io.DataInputStream
import java.io.DataOutputStream
import java.util.Collection
import java.util.TreeMap
import java.util.Vector

import scala.collection.JavaConversions._
import scala.reflect.ClassTag


class ConstantPool {
    // using TreeMap instead of TreeSet to simplify deduplication.
    var _constInfos = new TreeMap[ConstInfo, ConstInfo]()

    // only used for deserialization.  invalidate when _constInfos is modified.
    var _tmpConstInfosByIndex: TreeMap[Int, ConstInfo] = null

    def _get[T <: ConstInfo: ClassTag](info: T): T = {
        _tmpConstInfosByIndex = null
        if (_constInfos.containsKey(info)) {
            val cls = implicitly[ClassTag[T]].runtimeClass
            _constInfos.get(info) match {
                case t: T if cls.isInstance(t) => return t
                case _ => throw new Exception("unexpected const info type")
            }
        }
        _constInfos.put(info, info)
        return info
    }

    def getUtf8(value: String): ConstUtf8Info = {
        return _get[ConstUtf8Info](new ConstUtf8Info(value))
    }

    def getInteger(value: Int): ConstIntegerInfo = {
        return _get[ConstIntegerInfo](new ConstIntegerInfo(value))
    }

    def getLong(value: Long): ConstLongInfo = {
        return _get[ConstLongInfo](new ConstLongInfo(value))
    }

    def getFloat(value: Float): ConstFloatInfo = {
        return _get[ConstFloatInfo](new ConstFloatInfo(value))
    }

    def getDouble(value: Double): ConstDoubleInfo = {
        return _get[ConstDoubleInfo](new ConstDoubleInfo(value))
    }

    def getString(value: String): ConstStringInfo = {
        return _get[ConstStringInfo](new ConstStringInfo(getUtf8(value)))
    }

    def getClass(name: String): ConstClassInfo = {
        return _get[ConstClassInfo](new ConstClassInfo(getUtf8(name)))
    }

    def getMethodType(methodType: MethodType): ConstMethodTypeInfo = {
        return _get[ConstMethodTypeInfo](
                new ConstMethodTypeInfo(
                        methodType,
                        getUtf8(methodType.descriptorString())))
    }

    def _getByIndex[T <: ConstInfo : ClassTag](index: Int): T = {
        if (_tmpConstInfosByIndex == null) {
            throw new Exception(
                    "can only lookup by index during deserialization")
        }

        val info = _tmpConstInfosByIndex.get(index)
        if (info == null) {
            throw new Exception("entry not found")
        }

        val cls = implicitly[ClassTag[T]].runtimeClass
        info match {
            case t: T if cls.isInstance(t) => return t
            case _ => throw new Exception(
                    "unexpected const info type at " + index)
        }
    }

    def getUtf8ByIndex(index: Int): ConstUtf8Info = {
        return _getByIndex[ConstUtf8Info](index)
    }

    def getIntegerByIndex(index: Int): ConstIntegerInfo = {
        return _getByIndex[ConstIntegerInfo](index)
    }

    def getLongByIndex(index: Int): ConstLongInfo = {
        return _getByIndex[ConstLongInfo](index)
    }

    def getFloatByIndex(index: Int): ConstFloatInfo = {
        return _getByIndex[ConstFloatInfo](index)
    }

    def getStringByIndex(index: Int): ConstStringInfo = {
        return _getByIndex[ConstStringInfo](index)
    }

    def getNameAndTypeByIndex(index: Int): ConstNameAndTypeInfo = {
        return _getByIndex[ConstNameAndTypeInfo](index)
    }

    def getRefByIndex(index: Int): ConstRefInfo = {
        return _getByIndex[ConstRefInfo](index)
    }

    def getClassByIndex(index: Int): ConstClassInfo = {
        return _getByIndex[ConstClassInfo](index)
    }

    def _assignIndex(): Int = {
        var nextIndex = 1
        for (info <- _constInfos.keySet()) {
            info.index = nextIndex
            nextIndex += info.indexSize()
        }

        if (nextIndex > 65535) {
            throw new Exception("const pool too large")
        }
        return nextIndex
    }

    def serialize(output: DataOutputStream) {
        val count = _assignIndex()
        output.writeShort(count)

        for (info <- _constInfos.keySet()) {
            info.serialize(output)
        }
    }

    // TODO bind invoke dynamic bootstrap method attr index to reference
    def deserialize(input: DataInputStream) {
        if (!_constInfos.isEmpty()) {
            throw new Exception("deserializing into non-empty constant pool")
        }

        var constants = _parse(input)
        _generateIndexMap(constants)
        _bindConstReferences(constants)
        _populateAndDedup(constants)
    }

    def _parse(input: DataInputStream): Vector[ConstInfo] = {
        var result = new Vector[ConstInfo]()

        val constPoolCount = input.readUnsignedShort()

        var nextIndex = 1
        while (nextIndex < constPoolCount) {
            val tag = input.readByte()
            var info = tag match {
                case ConstInfo.UTF8 => new ConstUtf8Info()
                case ConstInfo.INTEGER => new ConstIntegerInfo()
                case ConstInfo.LONG => new ConstLongInfo()
                case ConstInfo.FLOAT => new ConstFloatInfo()
                case ConstInfo.DOUBLE => new ConstDoubleInfo()
                case ConstInfo.STRING => new ConstStringInfo()
                case ConstInfo.CLASS => new ConstClassInfo()
                case ConstInfo.NAME_AND_TYPE => new ConstNameAndTypeInfo()
                case ConstInfo.METHOD_TYPE => new ConstMethodTypeInfo()
                case ConstInfo.FIELD_REF => new ConstFieldRefInfo()
                case ConstInfo.METHOD_REF => new ConstMethodRefInfo()
                case ConstInfo.INTERFACE_METHOD_REF =>
                        new ConstInterfaceMethodRefInfo()
                case ConstInfo.METHOD_HANDLE => new ConstMethodHandleInfo()
                case ConstInfo.INVOKE_DYNAMIC => new ConstInvokeDynamicInfo()
                case _ => throw new Exception("Unknown const info type: " + tag)
            }
            info.index = nextIndex
            info.deserialize(tag, input)

            nextIndex += info.indexSize()

            result.add(info)
        }

        return result
    }

    def _generateIndexMap(constants: Collection[ConstInfo]) {
        _tmpConstInfosByIndex = new TreeMap[Int, ConstInfo]()

        for (info <- constants) {
            if (info.index < 1) {
                throw new Exception("invalid index")
            }
            if (_tmpConstInfosByIndex.containsKey(info.index)) {
                throw new Exception("duplicate index")
            }
            _tmpConstInfosByIndex.put(info.index, info)
        }
    }

    def _bindConstReferences(constants: Vector[ConstInfo]) {
        for (info <- constants) {
            info.bindConstReferences(this)
        }
    }

    def _populateAndDedup(constants: Vector[ConstInfo]) {
        for (info <- constants) {
            val first = _constInfos.get(info)
            if (first == null) {
                _constInfos.put(info, info)
            } else {
                println("Dedup " + info.index + " -> " + first.index)
                println("  old: " + info.debugString())
                println("  new: " + info.debugString())
                _tmpConstInfosByIndex.put(info.index, first)
            }
        }
    }
}
