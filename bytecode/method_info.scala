import java.io.DataInputStream
import java.io.DataOutputStream


class MethodSignature(
        n: String,
        p: ParameterTypes) extends Comparable[MethodSignature] {
    val name = n
    val parameters = p

    def compareTo(other: MethodSignature): Int = {
        val c = name.compareTo(other.name)
        if (c != 0) {
            return c
        }
        return parameters.compareTo(other.parameters)
    }

    override def toString(): String = name + parameters.descriptorString()
}

class MethodInfo(
        c: ClassInfo,
        n: String,
        f: MethodType) extends AttributeOwner with Comparable[MethodInfo] {
    def this(c: ClassInfo) = this(c, null, null)

    var _owner = c

    var _access = new MethodAccessFlags(this)
    var _name: ConstUtf8Info = null
    if (n != null) {
        _name = _owner.constants().getUtf8(n)
    }

    var _descriptor: MethodType = f
    var _descriptorString: ConstUtf8Info = null
    if (f != null) {
        _descriptorString = constants.getUtf8(_descriptor.descriptorString())
    }

    var _attributes = new MethodAttributes(this)

    def constants(): ConstantPool = _owner.constants()

    def compareTo(other: MethodInfo): Int = {
        if (_owner != other._owner) {
            throw new Exception(
                    "comparing methods from different class info objects")
        }

        val c = _name.compareTo(other._name)
        if (c != 0) {
            return c
        }
        return _descriptorString.compareTo(other._descriptorString)
    }

    def signature(): MethodSignature = {
        return new MethodSignature(name(), descriptor().parameters)
    }

    def access(): MethodAccessFlags = _access
    def name(): String = _name.value()
    def descriptorString(): String = _descriptorString.value()
    def descriptor(): MethodType = _descriptor
    def attributes(): MethodAttributes = _attributes

    def serialize(output: DataOutputStream) {
        _access.serialize(output)
        output.writeShort(_name.index)
        output.writeShort(_descriptorString.index)
        _attributes.serialize(output)
    }

    def deserialize(input: DataInputStream) {
        _access.deserialize(input)
        _name = constants().getUtf8ByIndex(input.readUnsignedShort())

        _descriptorString = constants().getUtf8ByIndex(
                input.readUnsignedShort())
        var parser = new DescriptorParser(_descriptorString.value())
        _descriptor = parser.parseMethodDescriptor()

        _attributes.deserialize(input)
    }
}
