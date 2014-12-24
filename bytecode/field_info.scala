import java.io.DataInputStream
import java.io.DataOutputStream


class FieldInfo(
        c: ClassInfo,
        n: ConstUtf8Info,
        f: FieldType) extends AttributeOwner {
    def this(c: ClassInfo) = this(c, null, null)

    var _owner = c

    var _access = new FieldAccessFlags(this)
    var _name: ConstUtf8Info = n

    var _descriptor: FieldType = f
    var _descriptorString: ConstUtf8Info = null
    if (f != null) {
        _descriptorString = constants().getUtf8(_descriptor.descriptorString())
    }

    var _attributes = new FieldAttributes(this)

    def constants(): ConstantPool = _owner.constants()

    def access(): FieldAccessFlags = _access
    def name(): String = _name.value()
    def descriptorString(): String = _descriptorString.value()
    def descriptor(): FieldType = _descriptor
    def attributes(): FieldAttributes = _attributes

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
        _descriptor = parser.parseFieldDescriptor()

        _attributes.deserialize(input)
    }
}
