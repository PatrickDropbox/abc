import java.io.DataInputStream
import java.io.DataOutputStream


class FieldInfo {
    var _access = new FieldAccessFlags()
    var _name: ConstUtf8Info = null

    var _descriptorString: ConstUtf8Info = null
    var _descriptor: FieldType = null

    var _attributes = new FieldAttributes()

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

        _attributes.deserialize(input, constants)
    }
}
