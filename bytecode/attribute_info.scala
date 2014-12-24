import java.io.DataInputStream
import java.io.DataOutputStream
import java.util.Vector

import scala.collection.JavaConversions._


trait AttributeOwner {
    def constants(): ConstantPool
}

abstract class Attribute(o: AttributeOwner, n: ConstUtf8Info) {
    var _owner: AttributeOwner = o
    var _name: ConstUtf8Info = n

    def name(): String = _name.value()

    def serialize(output: DataOutputStream)

    def deserialize(name: ConstUtf8Info,
                    attrLength: Int,
                    input: DataInputStream)

    def debugString(indent: String): String
}

class UnsupportedAttribute(
        o: AttributeOwner,
        n: ConstUtf8Info,
        b: Array[Byte]) extends Attribute(o, n) {
    def this(o: AttributeOwner) = this(o, null, null)

    var _bytes: Array[Byte] = b

    def serialize(output: DataOutputStream) {
        output.writeShort(_name.index)
        output.writeInt(_bytes.length)
        output.write(_bytes)
    }

    def deserialize(n: ConstUtf8Info, attrLength: Int, input: DataInputStream) {
        _name = n
        _bytes = new Array[Byte](attrLength)
        input.readFully(_bytes)
    }

    def debugString(indent: String): String = indent + name() + " (Unsupported)"
}

class SourceFileAttribute(
        o: AttributeOwner,
        n: ConstUtf8Info) extends Attribute(
                o,
                o.constants().getUtf8("SourceFile")) {
    def this(o: AttributeOwner) = this(o, null)

    var _sourceFile: ConstUtf8Info = n

    def sourceFile(): String = _sourceFile.value()

    def serialize(output: DataOutputStream) {
        output.writeShort(_name.index)
        output.writeInt(2)
        output.writeShort(_sourceFile.index)
    }

    def deserialize(n: ConstUtf8Info, attrLength: Int, input: DataInputStream) {
        if (n.compareTo(_name) != 0) {
            throw new Exception("Unexpected attribute name: " + n.value())
        }
        if (attrLength != 2) {
            throw new Exception("Unexpected attribute length")
        }
        _sourceFile = _owner.constants().getUtf8ByIndex(
                input.readUnsignedShort())
    }

    def debugString(indent: String): String = {
        return indent + name() + ": " + sourceFile()
    }
}

class DeprecatedAttribute(
        o: AttributeOwner) extends Attribute(
                o,
                o.constants().getUtf8("Deprecated")) {

    def serialize(output: DataOutputStream) {
        output.writeShort(_name.index)
        output.writeInt(0)
    }

    def deserialize(n: ConstUtf8Info, attrLength: Int, input: DataInputStream) {
        if (n.compareTo(_name) != 0) {
            throw new Exception("Unexpected attribute name: " + n.value())
        }
        if (attrLength != 0) {
            throw new Exception("Unexpected attribute length")
        }
    }

    def debugString(indent: String): String = indent + "Deprecated"
}

