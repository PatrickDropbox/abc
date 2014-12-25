import java.io.ByteArrayOutputStream
import java.io.DataInputStream
import java.io.DataOutputStream
import java.util.Vector

import scala.collection.JavaConversions._


trait AttributeOwner {
    def constants(): ConstantPool

    // only applicable if the owner is FieldInfo
    def fieldType(): FieldType = null

    // only applicable if the owner is MethodInfo
    def methodType(): MethodType = null
}

abstract class Attribute(o: AttributeOwner, attributeName: String) {
    var _owner: AttributeOwner = o
    var _name: ConstUtf8Info = null
    if (attributeName != null) {
        _name = _owner.constants().getUtf8(attributeName)
    }

    def name(): String = _name.value()

    def serialize(output: DataOutputStream)

    def deserialize(name: ConstUtf8Info,
                    attrLength: Int,
                    input: DataInputStream)

    def debugString(indent: String): String
}

abstract class RawBytesAttribute(
        o: AttributeOwner,
        attributeName: String,
        b: Array[Byte]) extends Attribute(o, attributeName) {
    var _bytes: Array[Byte] = b

    def bytes(): Array[Byte] = _bytes

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
}

class NoValueAttribute(
        o: AttributeOwner,
        attributeName: String) extends Attribute(o, attributeName) {

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

    def debugString(indent: String): String = indent + name()
}

class StringValueAttribute(
        o: AttributeOwner,
        attributeName: String,
        v: String) extends Attribute(o, attributeName) {

    var _value: ConstUtf8Info = null
    if (v != null) {
        _value = _owner.constants().getUtf8(v)
    }

    def value(): String = _value.value()

    def serialize(output: DataOutputStream) {
        output.writeShort(_name.index)
        output.writeInt(2)
        output.writeShort(_value.index)
    }

    def deserialize(n: ConstUtf8Info, attrLength: Int, input: DataInputStream) {
        if (n.compareTo(_name) != 0) {
            throw new Exception("Unexpected attribute name: " + n.value())
        }
        if (attrLength != 2) {
            throw new Exception("Unexpected attribute length")
        }
        _value = _owner.constants().getUtf8ByIndex(
                input.readUnsignedShort())
    }

    def debugString(indent: String): String = {
        return indent + name() + ": " + value()
    }
}

class UnsupportedAttribute(
        o: AttributeOwner,
        attributeName: String,
        b: Array[Byte]) extends RawBytesAttribute(o, attributeName, b) {
    def this(o: AttributeOwner) = this(o, null, null)

    def debugString(indent: String): String = indent + name() + " (Unsupported)"
}

object SourceDebugExtensionAttribute {
    def modifiedUtf8(s: String): Array[Byte] = {
        if (s == null) {
            return null
        }
        var buffer = new ByteArrayOutputStream()
        (new DataOutputStream(buffer)).writeUTF(s)
        return buffer.toByteArray()
    }
}

class SourceDebugExtensionAttribute(
        o: AttributeOwner,
        s: String) extends RawBytesAttribute(
                o,
                "SourceDebugExtension",
                SourceDebugExtensionAttribute.modifiedUtf8(s)) {
    def this(o: AttributeOwner) = this(o, null)

    def debugString(indent: String): String = indent + name()
}

class SourceFileAttribute(
        o: AttributeOwner,
        n: String) extends StringValueAttribute(o, "SourceFile", n) {
    def this(o: AttributeOwner) = this(o, null)
}

// see page 118-123 for signature syntax
class SignatureAttribute(
        o: AttributeOwner,
        n: String) extends StringValueAttribute(o, "Signature", n) {
    def this(o: AttributeOwner) = this(o, null)
}

class DeprecatedAttribute(
        o: AttributeOwner) extends NoValueAttribute(o, "Deprecated") {
}

class SyntheticAttribute(
        o: AttributeOwner) extends NoValueAttribute(o, "Synthetic") {
}

