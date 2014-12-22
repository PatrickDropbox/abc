import java.io.ByteArrayOutputStream
import java.io.DataInputStream
import java.io.DataOutputStream
import java.util.HashMap


object ConstInfo {
    // see table 4.4-A (page 79)
    val CLASS = 7
    val FIELD_REF = 9
    val METHOD_REF = 10
    val INTERFACE_METHOD_REF = 11
    val STRING = 8
    val INTEGER = 3
    val FLOAT = 4
    val LONG = 5
    val DOUBLE = 6
    val NAME_AND_TYPE = 12
    val UTF8 = 1
    val METHOD_HANDLE = 15
    val METHOD_TYPE = 16
    val INVOKE_DYNAMIC = 18

    // tag -> order
    var tagTopoOrder = new HashMap[Int, Int]()

    // no dependencies
    tagTopoOrder.put(UTF8, 0)
    tagTopoOrder.put(INTEGER, 1)
    tagTopoOrder.put(LONG, 2)
    tagTopoOrder.put(FLOAT, 3)
    tagTopoOrder.put(DOUBLE, 4)
    // depends on utf8
    tagTopoOrder.put(STRING, 5)
    tagTopoOrder.put(CLASS, 6)
    tagTopoOrder.put(METHOD_TYPE, 7)
    tagTopoOrder.put(NAME_AND_TYPE, 8)
    // depends on class / name and type
    tagTopoOrder.put(FIELD_REF, 9)
    tagTopoOrder.put(METHOD_REF, 10)
    tagTopoOrder.put(INTERFACE_METHOD_REF, 11)
    // depends on ref infos
    tagTopoOrder.put(METHOD_HANDLE, 12)
    // depends on (external attribute) bootstrap method index
    tagTopoOrder.put(INVOKE_DYNAMIC, 13)

    def tagOrder(tag: Int): Int = {
        if (!tagTopoOrder.containsKey(tag)) {
            throw new Exception("Unknown tag type: " + tag)
        }
        return tagTopoOrder.get(tag)
    }
}

trait ConstInfo extends Comparable[ConstInfo] {
    var index = 0

    def indexSize(): Int = 1

    def tag(): Int

    def typeName(): String

    def debugString(): String = {
        val indexValue = _debugIndexValue()
        if (indexValue != "") {
            return String.format(
                    "%7s = %-18s %-15s  // %s",
                    "#" + index,
                    typeName(),
                    indexValue,
                    debugValue())
        }

        return String.format(
                "%7s = %-18s %s",
                "#" + index,
                typeName(),
                debugValue())
    }

    def _debugIndexValue(): String = ""

    def debugValue(): String

    def serialize(output: DataOutputStream)

    def deserialize(parsedTag: Int, input: DataInputStream)

    def bindConstReferences(pool: ConstantPool)

    def compareTo(other: ConstInfo): Int = {
        if (tag() < other.tag()) {
            return -1
        }
        if (tag() > other.tag()) {
            return 1
        }
        return _compareTo(other)
    }

    def _compareTo(other: ConstInfo): Int
}

// See section 4.4.7 page 85-86
//
// TODO: Fix encoding / decoding.  jvm uses a non-standard "modified-utf8"
// encoding.  Why use the standard when you can reinvent your own ...
class ConstUtf8Info extends ConstInfo {
    var value: String = ""

    def tag(): Int = ConstInfo.UTF8

    def typeName(): String = "Utf8"

    def debugValue(): String = {
        return value.replaceAll("\\p{C}", "?")
    }

    def serialize(output: DataOutputStream) {
        var buffer = new ByteArrayOutputStream()
        (new DataOutputStream(buffer)).writeUTF(value)
        val utf8Value = buffer.toByteArray()

        if (utf8Value.length > 65535) {
            throw new Exception("utf8 string too long")
        }

        output.writeByte(tag())
        output.writeShort(utf8Value.length)
        output.write(utf8Value)
    }

    def deserialize(parsedTag: Int, input: DataInputStream) {
        if (parsedTag != tag()) {
            throw new Exception("unexpected tag")
        }

        val length = input.readUnsignedShort()

        var utf8Bytes = new Array[Byte](length)
        input.readFully(utf8Bytes)

        value = new String(utf8Bytes, "UTF-8")
    }

    def bindConstReferences(pool: ConstantPool) {
        // nothing to bind
    }

    def _compareTo(o: ConstInfo): Int = {
        o match {
            case other: ConstUtf8Info => {
                return value.compareTo(other.value)
            }
            case _ => throw new Exception("unexpected other type")
        }
    }
}

// See section 4.4.4 page 82-83
class ConstIntegerInfo extends ConstInfo {
    var value: Int = 0

    def tag(): Int = ConstInfo.INTEGER

    def typeName(): String = "Integer"

    def debugValue(): String = {
        return "" + value
    }

    def serialize(output: DataOutputStream) {
        output.writeByte(tag())
        output.writeInt(value)
    }

    def deserialize(parsedTag: Int, input: DataInputStream) {
        if (parsedTag != tag()) {
            throw new Exception("unexpected tag")
        }
        value = input.readInt()
    }

    def bindConstReferences(pool: ConstantPool) {
        // nothing to bind
    }

    def _compareTo(o: ConstInfo): Int = {
        o match {
            case other: ConstIntegerInfo => {
                if (value < other.value) {
                    return -1
                }
                if (value > other.value) {
                    return 1
                }
                return 0
            }
            case _ => throw new Exception("unexpected other type")
        }
    }
}

// See section 4.4.5 page 83-85
class ConstLongInfo extends ConstInfo {
    var value: Long = 0

    def tag(): Int = ConstInfo.LONG

    override def indexSize(): Int = 2

    def typeName(): String = "Long"

    def debugValue(): String = {
        return "" + value + "l"
    }

    def serialize(output: DataOutputStream) {
        output.writeByte(tag())
        output.writeLong(value)
    }

    def deserialize(parsedTag: Int, input: DataInputStream) {
        if (parsedTag != tag()) {
            throw new Exception("unexpected tag")
        }
        value = input.readLong()
    }

    def bindConstReferences(pool: ConstantPool) {
        // nothing to bind
    }

    def _compareTo(o: ConstInfo): Int = {
        o match {
            case other: ConstLongInfo => {
                if (value < other.value) {
                    return -1
                }
                if (value > other.value) {
                    return 1
                }
                return 0
            }
            case _ => throw new Exception("unexpected other type")
        }
    }
}

// See section 4.4.4 page 82-83
class ConstFloatInfo extends ConstInfo {
    var value: Float = 0

    def tag(): Int = ConstInfo.FLOAT

    def typeName(): String = "Float"

    def debugValue(): String = {
        return "" + value + "f"
    }

    def serialize(output: DataOutputStream) {
        output.writeByte(tag())
        output.writeFloat(value)
    }

    def deserialize(parsedTag: Int, input: DataInputStream) {
        if (parsedTag != tag()) {
            throw new Exception("unexpected tag")
        }
        value = input.readFloat()
    }

    def bindConstReferences(pool: ConstantPool) {
        // nothing to bind
    }

    def _compareTo(o: ConstInfo): Int = {
        o match {
            case other: ConstFloatInfo => {
                if (value < other.value) {
                    return -1
                }
                if (value > other.value) {
                    return 1
                }
                return 0
            }
            case _ => throw new Exception("unexpected other type")
        }
    }
}

// See section 4.4.5 page 83-85
class ConstDoubleInfo extends ConstInfo {
    var value: Double = 0

    def tag(): Int = ConstInfo.DOUBLE

    override def indexSize(): Int = 2

    def typeName(): String = "Double"

    def debugValue(): String = {
        return "" + value + "d"
    }

    def serialize(output: DataOutputStream) {
        output.writeByte(tag())
        output.writeDouble(value)
    }

    def deserialize(parsedTag: Int, input: DataInputStream) {
        if (parsedTag != tag()) {
            throw new Exception("unexpected tag")
        }
        value = input.readDouble()
    }

    def bindConstReferences(pool: ConstantPool) {
        // nothing to bind
    }

    def _compareTo(o: ConstInfo): Int = {
        o match {
            case other: ConstDoubleInfo => {
                if (value < other.value) {
                    return -1
                }
                if (value > other.value) {
                    return 1
                }
                return 0
            }
            case _ => throw new Exception("unexpected other type")
        }
    }
}

// see section 4.4.3 page 81-82
class ConstStringInfo extends ConstInfo {
    var utf8String: ConstUtf8Info = null

    // only used during deserialization
    var _tmpUtf8StringIndex = 0

    def tag(): Int = ConstInfo.STRING

    def typeName(): String = "String"

    override def _debugIndexValue(): String = {
        return "#" + utf8String.index
    }

    def debugValue(): String = {
        return utf8String.debugValue()
    }

    def serialize(output: DataOutputStream) {
        output.writeByte(tag())
        output.writeShort(utf8String.index)
    }

    def deserialize(parsedTag: Int, input: DataInputStream) {
        if (parsedTag != tag()) {
            throw new Exception("unexpected tag")
        }
        _tmpUtf8StringIndex = input.readUnsignedShort()
    }

    def bindConstReferences(pool: ConstantPool) {
        utf8String = pool.getUtf8ByIndex(_tmpUtf8StringIndex)
    }

    def _compareTo(o: ConstInfo): Int = {
        o match {
            case other: ConstStringInfo => {
                return utf8String.compareTo(other.utf8String)
            }
            case _ => throw new Exception("unexpected other type")
        }
    }
}

// see section 4.4.1 page 79-80
class ConstClassInfo extends ConstInfo {
    var className: ConstUtf8Info = null

    // only used during deserialization
    var _tmpClassNameIndex = 0

    def tag(): Int = ConstInfo.CLASS

    def typeName(): String = "Class"

    override def _debugIndexValue(): String = {
        return "#" + className.index
    }

    def debugValue(): String = {
        return className.debugValue()
    }

    def serialize(output: DataOutputStream) {
        output.writeByte(tag())
        output.writeShort(className.index)
    }

    def deserialize(parsedTag: Int, input: DataInputStream) {
        if (parsedTag != tag()) {
            throw new Exception("unexpected tag")
        }
        _tmpClassNameIndex = input.readUnsignedShort()
    }

    def bindConstReferences(pool: ConstantPool) {
        className = pool.getUtf8ByIndex(_tmpClassNameIndex)
    }

    def _compareTo(o: ConstInfo): Int = {
        o match {
            case other: ConstClassInfo => {
                return className.compareTo(other.className)
            }
            case _ => throw new Exception("unexpected other type")
        }
    }
}

// see section 4.4.9 page 89
class ConstMethodTypeInfo extends ConstInfo {
    var descriptorString: ConstUtf8Info = null
    var descriptor: MethodType = null

    // only used during deserialization
    var _tmpDescriptorIndex = 0

    def tag(): Int = ConstInfo.METHOD_TYPE

    def typeName(): String = "MethodType"

    override def _debugIndexValue(): String = {
        return "#" + descriptorString.index
    }

    def debugValue(): String = {
        return descriptorString.debugValue()
    }

    def serialize(output: DataOutputStream) {
        output.writeByte(tag())
        output.writeShort(descriptorString.index)
    }

    def deserialize(parsedTag: Int, input: DataInputStream) {
        if (parsedTag != tag()) {
            throw new Exception("unexpected tag")
        }
        _tmpDescriptorIndex = input.readUnsignedShort()
    }

    def bindConstReferences(pool: ConstantPool) {
        descriptorString = pool.getUtf8ByIndex(_tmpDescriptorIndex)
        var parser = new DescriptorParser(descriptorString.value)
        descriptor = parser.parseMethodDescriptor()
    }

    def _compareTo(o: ConstInfo): Int = {
        o match {
            case other: ConstMethodTypeInfo => {
                return descriptorString.compareTo(other.descriptorString)
            }
            case _ => throw new Exception("unexpected other type")
        }
    }
}

// see section 4.4.6 page 85
class ConstNameAndTypeInfo extends ConstInfo {
    var name: ConstUtf8Info = null
    var descriptor: ConstUtf8Info = null

    // only used during deserialization
    var _tmpNameIndex = 0
    var _tmpDescriptorIndex = 0

    def tag(): Int = ConstInfo.NAME_AND_TYPE

    def typeName(): String = "NameAndType"

    override def _debugIndexValue(): String = {
        return "#" + name.index + ":#" + descriptor.index
    }

    def debugValue(): String = {
        return name.debugValue() + ":" + descriptor.debugValue()
    }

    def serialize(output: DataOutputStream) {
        output.writeByte(tag())
        output.writeShort(name.index)
        output.writeShort(descriptor.index)
    }

    def deserialize(parsedTag: Int, input: DataInputStream) {
        if (parsedTag != tag()) {
            throw new Exception("unexpected tag")
        }
        _tmpNameIndex = input.readUnsignedShort()
        _tmpDescriptorIndex = input.readUnsignedShort()
    }

    def bindConstReferences(pool: ConstantPool) {
        name = pool.getUtf8ByIndex(_tmpNameIndex)
        descriptor = pool.getUtf8ByIndex(_tmpDescriptorIndex)
    }

    def _compareTo(o: ConstInfo): Int = {
        o match {
            case other: ConstNameAndTypeInfo => {
                val c = name.compareTo(other.name)
                if (c != 0) {
                    return c
                }
                return descriptor.compareTo(other.descriptor)
            }
            case _ => throw new Exception("unexpected other type")
        }
    }
}

// see section 4.4.2 page 80
abstract class ConstRefInfo extends ConstInfo {
    var classInfo: ConstClassInfo = null
    var nameAndType: ConstNameAndTypeInfo = null

    // only used during deserialization
    var _tmpClassIndex = 0
    var _tmpNameAndTypeIndex = 0

    override def _debugIndexValue(): String = {
        return "#" + classInfo.index + ".#" + nameAndType.index
    }

    def debugValue(): String = {
        return classInfo.debugValue() + "." + nameAndType.debugValue()
    }

    def serialize(output: DataOutputStream) {
        output.writeByte(tag())
        output.writeShort(classInfo.index)
        output.writeShort(nameAndType.index)
    }

    def deserialize(parsedTag: Int, input: DataInputStream) {
        if (parsedTag != tag()) {
            throw new Exception("unexpected tag")
        }
        _tmpClassIndex = input.readUnsignedShort()
        _tmpNameAndTypeIndex = input.readUnsignedShort()
    }

    def bindConstReferences(pool: ConstantPool) {
        classInfo = pool.getClassByIndex(_tmpClassIndex)
        nameAndType = pool.getNameAndTypeByIndex(_tmpNameAndTypeIndex)
    }

    def _compareTo(o: ConstInfo): Int = {
        o match {
            case other: ConstRefInfo => {
                val c = classInfo.compareTo(other.classInfo)
                if (c != 0) {
                    return c
                }
                return nameAndType.compareTo(other.nameAndType)
            }
            case _ => throw new Exception("unexpected other type")
        }
    }
}

// see section 4.4.2 page 80
class ConstFieldRefInfo extends ConstRefInfo {
    def tag(): Int = ConstInfo.FIELD_REF

    def typeName(): String = "FieldRef"
}

// see section 4.4.2 page 80
class ConstMethodRefInfo extends ConstRefInfo {
    def tag(): Int = ConstInfo.METHOD_REF

    def typeName(): String = "MethodRef"
}

// see section 4.4.2 page 80
class ConstInterfaceMethodRefInfo extends ConstRefInfo {
    def tag(): Int = ConstInfo.INTERFACE_METHOD_REF

    def typeName(): String = "InterfaceMethodRef"
}

// see section 4.4.8 page 87-89
class ConstMethodHandleInfo extends ConstInfo {
    var referenceKind: Byte = 0
    var reference: ConstRefInfo = null

    // only used during deserialization
    var _tmpReferenceIndex = 0

    def tag(): Int = ConstInfo.METHOD_HANDLE

    def typeName(): String = "MethodHandle"

    override def _debugIndexValue(): String = {
        return "" + referenceKind + " #" + reference.index
    }

    def debugValue(): String = {
        return "" + referenceKind + " " + reference.debugValue()
    }

    def serialize(output: DataOutputStream) {
        output.writeByte(tag())
        output.writeByte(referenceKind)
        output.writeShort(reference.index)
    }

    def deserialize(parsedTag: Int, input: DataInputStream) {
        if (parsedTag != tag()) {
            throw new Exception("unexpected tag")
        }
        referenceKind = input.readByte()
        _tmpReferenceIndex = input.readUnsignedShort()
    }

    def bindConstReferences(pool: ConstantPool) {
        reference = pool.getRefByIndex(_tmpReferenceIndex)
    }

    def _compareTo(o: ConstInfo): Int = {
        o match {
            case other: ConstMethodHandleInfo => {
                if (referenceKind < other.referenceKind) {
                    return -1
                }
                if (referenceKind > other.referenceKind) {
                    return 1
                }
                return reference.compareTo(other.reference)
            }
            case _ => throw new Exception("unexpected other type")
        }
    }
}

// TODO: bind bootstrap method attr index to real entry
class ConstInvokeDynamicInfo extends ConstInfo {
    var bootstrapMethodAttrIndex = 0
    var nameAndType: ConstNameAndTypeInfo = null

    // only used during deserialization
    var _tmpNameAndTypeIndex = 0

    def tag(): Int = ConstInfo.INVOKE_DYNAMIC

    def typeName(): String = "InvokeDynamic"

    override def _debugIndexValue(): String = {
        return "" + bootstrapMethodAttrIndex + " #" + nameAndType.index
    }

    def debugValue(): String = {
        return "" + bootstrapMethodAttrIndex + " " + nameAndType.debugValue()
    }

    def serialize(output: DataOutputStream) {
        output.writeByte(tag())
        output.writeShort(bootstrapMethodAttrIndex)
        output.writeShort(nameAndType.index)
    }

    def deserialize(parsedTag: Int, input: DataInputStream) {
        if (parsedTag != tag()) {
            throw new Exception("unexpected tag")
        }
        bootstrapMethodAttrIndex = input.readUnsignedShort()
        _tmpNameAndTypeIndex = input.readUnsignedShort()
    }

    def bindConstReferences(pool: ConstantPool) {
        nameAndType = pool.getNameAndTypeByIndex(_tmpNameAndTypeIndex)
    }

    def _compareTo(o: ConstInfo): Int = {
        o match {
            case other: ConstInvokeDynamicInfo => {
                if (bootstrapMethodAttrIndex < other.bootstrapMethodAttrIndex) {
                    return -1
                }
                if (bootstrapMethodAttrIndex > other.bootstrapMethodAttrIndex) {
                    return 1
                }
                return nameAndType.compareTo(other.nameAndType)
            }
            case _ => throw new Exception("unexpected other type")
        }
    }
}
