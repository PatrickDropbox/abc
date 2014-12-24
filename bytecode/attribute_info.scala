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

    def _readAttributes(inputStream: DataInputStream): Vector[Attribute] = {
        var attributes = new Vector[Attribute]()

        val attrCount = inputStream.readUnsignedShort()
        for (_ <- 1 to attrCount) {
            val name = _owner.constants().getUtf8ByIndex(
                    inputStream.readUnsignedShort())
            var attr = name.value() match {
                // TODO
                case "Deprecated" => new DeprecatedAttribute(_owner)
                case "SourceFile" => new SourceFileAttribute(_owner)
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
    var _sourceFile: SourceFileAttribute = null
    var _deprecated: DeprecatedAttribute = null

    def sourceFile(): String = {
        if (_sourceFile == null) {
            return null
        }
        return _sourceFile.sourceFile()
    }

    def setSourceFile(s: String) {
        _sourceFile = new SourceFileAttribute(
                _owner,
                _owner.constants().getUtf8(s))
    }

    def isDeprecated(): Boolean = _deprecated != null
    def setIsDeprecated(b: Boolean) {
        if (b) {
            _deprecated = new DeprecatedAttribute(_owner)
        } else {
            _deprecated = null
        }
    }

    def allAttributes(): Vector[Attribute] = {
        var allAttributes = new Vector[Attribute]()

        // TODO

        if (_sourceFile != null) {
            allAttributes.add(_sourceFile)
        }

        if (_deprecated != null) {
            allAttributes.add(_deprecated)
        }

        for (attr <- _unsupported) {
            allAttributes.add(attr)
        }

        return allAttributes
    }

    def deserialize(input: DataInputStream) {
        for (a <- _readAttributes(input)) {
            a match {
                // TODO
                case attr: DeprecatedAttribute => _deprecated = attr
                case attr: SourceFileAttribute => _sourceFile = attr
                case attr: UnsupportedAttribute => _unsupported.add(attr)
                case _ => throw new Exception(
                        "Unexpected class attribute: " + a.name())
            }
        }
    }
}

class FieldAttributes(f: FieldInfo) extends AttributeGroup(f) {
    var _deprecated: DeprecatedAttribute = null

    def isDeprecated(): Boolean = _deprecated != null
    def setIsDeprecated(b: Boolean) {
        if (b) {
            _deprecated = new DeprecatedAttribute(_owner)
        } else {
            _deprecated = null
        }
    }

    def allAttributes(): Vector[Attribute] = {
        var allAttributes = new Vector[Attribute]()

        // TODO

        if (_deprecated != null) {
            allAttributes.add(_deprecated)
        }

        for (attr <- _unsupported) {
            allAttributes.add(attr)
        }

        return allAttributes
    }

    def deserialize(input: DataInputStream) {
        for (a <- _readAttributes(input)) {
            a match {
                // TODO
                case attr: DeprecatedAttribute => _deprecated = attr
                case attr: UnsupportedAttribute => _unsupported.add(attr)
                case _ => throw new Exception(
                        "Unexpected class attribute: " + a.name())
            }
        }
    }
}

class MethodAttributes(m: MethodInfo) extends AttributeGroup(m) {
    var _deprecated: DeprecatedAttribute = null

    def isDeprecated(): Boolean = _deprecated != null
    def setIsDeprecated(b: Boolean) {
        if (b) {
            _deprecated = new DeprecatedAttribute(_owner)
        } else {
            _deprecated = null
        }
    }

    def allAttributes(): Vector[Attribute] = {
        var allAttributes = new Vector[Attribute]()

        // TODO

        if (_deprecated != null) {
            allAttributes.add(_deprecated)
        }

        for (attr <- _unsupported) {
            allAttributes.add(attr)
        }

        return allAttributes
    }

    def deserialize(input: DataInputStream) {
        for (a <- _readAttributes(input)) {
            a match {
                // TODO
                case attr: DeprecatedAttribute => _deprecated = attr
                case attr: UnsupportedAttribute => _unsupported.add(attr)
                case _ => throw new Exception(
                        "Unexpected class attribute: " + a.name())
            }
        }
    }
}

// TODO code attributes
