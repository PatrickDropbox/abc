import java.io.DataInputStream
import java.io.DataOutputStream


class MethodInfo(
        c: ClassInfo,
        n: ConstUtf8Info,
        s: ConstUtf8Info,
        f: MethodType) extends AttributeOwner {
    def this(c: ClassInfo) = this(c, null, null, null)

    var _owner = c

    var _access = new MethodAccessFlags(this)
    var _name: ConstUtf8Info = n

    var _descriptorString: ConstUtf8Info = s
    var _descriptor: MethodType = f

    var _attributes = new MethodAttributes(this)

    def constants(): ConstantPool = _owner.constants()

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

    def deserialize(input: DataInputStream, constants: ConstantPool) {
        _access.deserialize(input)
        _name = constants.getUtf8ByIndex(input.readUnsignedShort())

        _descriptorString = constants.getUtf8ByIndex(input.readUnsignedShort())
        var parser = new DescriptorParser(_descriptorString.value())
        _descriptor = parser.parseMethodDescriptor()

        _attributes.deserialize(input)
    }
}
