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

    def debugString(indent: String): String = indent + name + " (Unsupported)"
}

abstract class AttributeGroup(o: AttributeOwner) {
    var _owner: AttributeOwner = o

    var _unsupported = new Vector[UnsupportedAttribute]()

    def unsupported(): Vector[UnsupportedAttribute] = _unsupported

    def allAttributes(): Vector[Attribute]

    def debugString(indent: String): String = {
        val attributes = allAttributes()

        if (attributes.isEmpty()) {
            return indent + "(no attributes)\n"
        }

        var result = ""
        for (attr <- attributes) {
            result += attr.debugString(indent) + "\n"
        }

        return result
    }

    def serialize(output: DataOutputStream) {
        val attributes = allAttributes()

        output.writeShort(attributes.size())
        for (attr <- attributes) {
            attr.serialize(output)
        }
    }

    def _parse(inputStream: DataInputStream): Vector[Attribute] = {
        var attributes = new Vector[Attribute]()

        val attrCount = inputStream.readUnsignedShort()
        for (_ <- 1 to attrCount) {
            val name = _owner.constants().getUtf8ByIndex(
                    inputStream.readUnsignedShort())
            var attr = name.value() match {
                case _ => new UnsupportedAttribute(_owner)
            }

            attr.deserialize(name, inputStream.readInt(), inputStream)
            attributes.add(attr)
        }

        return attributes
    }

    def deserialize(input: DataInputStream)
}

class ClassAttributes(c: ClassInfo) extends AttributeGroup(c) {
    def allAttributes(): Vector[Attribute] = {
        var allAttributes = new Vector[Attribute]()

        // TODO

        for (attr <- _unsupported) {
            allAttributes.add(attr)
        }

        return allAttributes
    }

    def deserialize(input: DataInputStream) {
        for (attr <- _parse(input)) {
            attr match {
                // TODO
                case u: UnsupportedAttribute => _unsupported.add(u)
                case _ => throw new Exception(
                        "Unexpected class attribute: " + attr.name())
            }
        }
    }
}

class FieldAttributes(f: FieldInfo) extends AttributeGroup(f) {
    def allAttributes(): Vector[Attribute] = {
        var allAttributes = new Vector[Attribute]()

        // TODO

        for (attr <- _unsupported) {
            allAttributes.add(attr)
        }

        return allAttributes
    }

    def deserialize(input: DataInputStream) {
        for (attr <- _parse(input)) {
            attr match {
                // TODO
                case u: UnsupportedAttribute => _unsupported.add(u)
                case _ => throw new Exception(
                        "Unexpected class attribute: " + attr.name())
            }
        }
    }
}

class MethodAttributes(m: MethodInfo) extends AttributeGroup(m) {
    def allAttributes(): Vector[Attribute] = {
        var allAttributes = new Vector[Attribute]()

        // TODO

        for (attr <- _unsupported) {
            allAttributes.add(attr)
        }

        return allAttributes
    }

    def deserialize(input: DataInputStream) {
        for (attr <- _parse(input)) {
            attr match {
                // TODO
                case u: UnsupportedAttribute => _unsupported.add(u)
                case _ => throw new Exception(
                        "Unexpected class attribute: " + attr.name())
            }
        }
    }
}
