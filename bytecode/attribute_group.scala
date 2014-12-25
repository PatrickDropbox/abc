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
/*
EnclosingMethod
BootstrapMethods
Code
Exceptions
RuntimeVisibleParameterAnnotations
RuntimeInvisibleParameterAnnotations
AnnotationDefault
MethodParameters
RuntimeVisibleAnnotations
RuntimeInvisibleAnnotations
LineNumberTable
LocalVariableTable
LocalVariableTypeTable
StackMapTable
RuntimeVisibleTypeAnnotations
RuntimeInvisibleTypeAnnotations

*/
                case "ConstantValue" => new ConstantValueAttribute(_owner)
                case "Deprecated" => new DeprecatedAttribute(_owner)
                case "InnerClasses" => new InnerClassesAttribute(_owner)
                case "Signature" => new SignatureAttribute(_owner)
                case "SourceDebugExtension" =>
                        new SourceDebugExtensionAttribute(_owner)
                case "SourceFile" => new SourceFileAttribute(_owner)
                case "Synthetic" => new SyntheticAttribute(_owner)
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
    var _signature: SignatureAttribute = null
    var _innerClasses: InnerClassesAttribute = null
    var _deprecated: DeprecatedAttribute = null
    var _synthetic: SyntheticAttribute = null
    var _sourceDebugExtension: SourceDebugExtensionAttribute = null

    def sourceFile(): String = {
        if (_sourceFile == null) {
            return null
        }
        return _sourceFile.value()
    }
    def setSourceFile(s: String) {
        if (s == null) {
            _sourceFile = null
        } else {
            _sourceFile = new SourceFileAttribute(_owner, s)
        }
    }

    def signature(): String = {
        if (_signature == null) {
            return null
        }
        return _signature.value()
    }
    def setSignature(s: String) {
        if (s == null) {
            _signature = null
        } else {
            _signature = new SignatureAttribute(_owner, s)
        }
    }

    def innerClasses(): InnerClassesAttribute = innerClasses
    def addInnerClass(
            innerClass: String,
            outerClass: String,
            innerName: String): InnerClassInfo = {
        val inner = new InnerClassInfo(
                _owner,
                innerClass,
                outerClass,
                innerName)

        if (_innerClasses == null) {
            _innerClasses = new InnerClassesAttribute(_owner)
        }

        _innerClasses.add(inner)
        return inner
    }

    def isDeprecated(): Boolean = _deprecated != null
    def setIsDeprecated(b: Boolean) {
        if (b) {
            _deprecated = new DeprecatedAttribute(_owner)
        } else {
            _deprecated = null
        }
    }

    def isSynthetic(): Boolean = _synthetic != null
    def setIsSynthetic(b: Boolean) {
        if (b) {
            _synthetic = new SyntheticAttribute(_owner)
        } else {
            _synthetic = null
        }
    }

    def sourceDebugExtension(): String = {
        return new String(_sourceDebugExtension.bytes(), "UTF-8")
    }
    def setSourceDebugExtension(s: String) {
        if (s == null) {
            _sourceDebugExtension = null
        } else {
            _sourceDebugExtension = new SourceDebugExtensionAttribute(_owner, s)
        }
    }

    def allAttributes(): Vector[Attribute] = {
        var allAttributes = new Vector[Attribute]()

        // TODO

        if (_sourceFile != null) {
            allAttributes.add(_sourceFile)
        }
        if (_signature != null) {
            allAttributes.add(_signature)
        }
        if (_innerClasses != null) {
            allAttributes.add(_innerClasses)
        }
        if (_deprecated != null) {
            allAttributes.add(_deprecated)
        }
        if (_synthetic != null) {
            allAttributes.add(_synthetic)
        }
        if (_sourceDebugExtension != null) {
            allAttributes.add(_sourceDebugExtension)
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
                case attr: DeprecatedAttribute => {
                    if (_deprecated != null) {
                        throw new Exception("multiple deprecated attribute")
                    }
                    _deprecated = attr
                }
                case attr: SignatureAttribute => {
                    if (_signature != null) {
                        throw new Exception("multiple signature attribute")
                    }
                    _signature = attr
                }
                case attr: InnerClassesAttribute => {
                    if (_innerClasses != null) {
                        throw new Exception("multiple inner classes attribute")
                    }
                    _innerClasses = attr
                }
                case attr: SourceDebugExtensionAttribute => {
                    if (_sourceDebugExtension != null) {
                        throw new Exception(
                                "multiple source debug extension attribute")
                    }
                    _sourceDebugExtension = attr
                }
                case attr: SourceFileAttribute => {
                    if (_sourceFile != null) {
                        throw new Exception("multiple sourceFile attribute")
                    }
                    _sourceFile = attr
                }
                case attr: SyntheticAttribute => {
                    if (_synthetic != null) {
                        throw new Exception("multiple synthetic attribute")
                    }
                    _synthetic = attr
                }
                case attr: UnsupportedAttribute => _unsupported.add(attr)
                case _ => throw new Exception(
                        "Unexpected class attribute: " + a.name())
            }
        }
    }
}

class FieldAttributes(f: FieldInfo) extends AttributeGroup(f) {
    var _signature: SignatureAttribute = null
    var _constantValue: ConstantValueAttribute = null
    var _deprecated: DeprecatedAttribute = null
    var _synthetic: SyntheticAttribute = null

    def signature(): String = {
        if (_signature == null) {
            return null
        }
        return _signature.value()
    }
    def setSignature(s: String) {
        if (s == null) {
            _signature = null
        } else {
            _signature = new SignatureAttribute(_owner, s)
        }
    }

    def constantValue(): ConstantValueAttribute = _constantValue

    def isDeprecated(): Boolean = _deprecated != null
    def setIsDeprecated(b: Boolean) {
        if (b) {
            _deprecated = new DeprecatedAttribute(_owner)
        } else {
            _deprecated = null
        }
    }

    def isSynthetic(): Boolean = _synthetic != null
    def setIsSynthetic(b: Boolean) {
        if (b) {
            _synthetic = new SyntheticAttribute(_owner)
        } else {
            _synthetic = null
        }
    }

    def allAttributes(): Vector[Attribute] = {
        var allAttributes = new Vector[Attribute]()

        // TODO

        if (_signature != null) {
            allAttributes.add(_signature)
        }
        if (_deprecated != null) {
            allAttributes.add(_deprecated)
        }
        if (_synthetic != null) {
            allAttributes.add(_synthetic)
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
                case attr: ConstantValueAttribute => {
                    if (_constantValue != null) {
                        throw new Exception("multiple constant value attribute")
                    }
                    _constantValue = attr
                }
                case attr: DeprecatedAttribute => {
                    if (_deprecated != null) {
                        throw new Exception("multiple deprecated attribute")
                    }
                    _deprecated = attr
                }
                case attr: SignatureAttribute => {
                    if (_signature != null) {
                        throw new Exception("multiple signature attribute")
                    }
                    _signature = attr
                }
                case attr: SyntheticAttribute => {
                    if (_synthetic != null) {
                        throw new Exception("multiple synthetic attribute")
                    }
                    _synthetic = attr
                }
                case attr: UnsupportedAttribute => _unsupported.add(attr)
                case _ => throw new Exception(
                        "Unexpected field attribute: " + a.name())
            }
        }
    }
}

class MethodAttributes(m: MethodInfo) extends AttributeGroup(m) {
    var _signature: SignatureAttribute = null
    var _deprecated: DeprecatedAttribute = null
    var _synthetic: SyntheticAttribute = null

    def signature(): String = {
        if (_signature == null) {
            return null
        }
        return _signature.value()
    }
    def setSignature(s: String) {
        if (s == null) {
            _signature = null
        } else {
            _signature = new SignatureAttribute(_owner, s)
        }
    }

    def isDeprecated(): Boolean = _deprecated != null
    def setIsDeprecated(b: Boolean) {
        if (b) {
            _deprecated = new DeprecatedAttribute(_owner)
        } else {
            _deprecated = null
        }
    }

    def isSynthetic(): Boolean = _synthetic != null
    def setIsSynthetic(b: Boolean) {
        if (b) {
            _synthetic = new SyntheticAttribute(_owner)
        } else {
            _synthetic = null
        }
    }

    def allAttributes(): Vector[Attribute] = {
        var allAttributes = new Vector[Attribute]()

        // TODO

        if (_signature != null) {
            allAttributes.add(_signature)
        }
        if (_deprecated != null) {
            allAttributes.add(_deprecated)
        }
        if (_synthetic != null) {
            allAttributes.add(_synthetic)
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
                case attr: DeprecatedAttribute => {
                    if (_deprecated != null) {
                        throw new Exception("multiple deprecated attribute")
                    }
                    _deprecated = attr
                }
                case attr: SignatureAttribute => {
                    if (_signature != null) {
                        throw new Exception("multiple signature attribute")
                    }
                    _signature = attr
                }
                case attr: SyntheticAttribute => {
                    if (_synthetic != null) {
                        throw new Exception("multiple synthetic attribute")
                    }
                    _synthetic = attr
                }
                case attr: UnsupportedAttribute => _unsupported.add(attr)
                case _ => throw new Exception(
                        "Unexpected method attribute: " + a.name())
            }
        }
    }
}

// TODO code attributes
