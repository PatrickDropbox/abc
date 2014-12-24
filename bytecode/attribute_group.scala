import java.io.DataInputStream
import java.io.DataOutputStream
import java.util.Vector

import scala.collection.JavaConversions._


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
