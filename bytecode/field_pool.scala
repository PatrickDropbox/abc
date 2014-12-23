import java.io.DataOutputStream
import java.io.DataInputStream
import java.util.Collection
import java.util.TreeMap

import scala.collection.JavaConversions._


class FieldPool(c: ClassFile) {
    var _ownerClass = c

    var _fields = new TreeMap[String, FieldInfo]()

    var _constants: ConstantPool = c.constants()

    def fields(): Collection[FieldInfo] = _fields.values()

    def addField(name: String, fieldType: FieldType): FieldInfo = {
        if (_fields.containsKey(name)) {
            throw new Exception("adding duplicate field: " + name)
        }
        _fields.put(
                name,
                new FieldInfo(
                        _ownerClass,
                        _constants.getUtf8(name),
                        _constants.getUtf8(fieldType.descriptorString()),
                        fieldType))
    }

    def getField(name: String): FieldInfo = {
        val f = _fields.get(name)
        if (f == null) {
            throw new Exception("missing field: " + name)
        }
        return f
    }

    def serialize(output: DataOutputStream) {
        output.writeShort(_fields.size())
        for (field <- fields()) {
            field.serialize(output)
        }
    }

    def deserialize(input: DataInputStream) {
        if (!_fields.isEmpty()) {
            throw new Exception("deserializing into non-empty field pool")
        }

        val fieldCount = input.readUnsignedShort()
        for (i <- 1 to fieldCount) {
            var field = new FieldInfo(_ownerClass)
            field.deserialize(input, _constants)
            if (_fields.containsKey(field.name())) {
                throw new Exception("duplicate field name: " + field.name())
            }
            _fields.put(field.name(), field)
        }
    }
}
