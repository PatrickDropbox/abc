import java.io.DataOutputStream
import java.io.DataInputStream
import java.util.Collection
import java.util.TreeMap

import scala.collection.JavaConversions._


class FieldPool(c: ConstantPool) {
    var _constants: ConstantPool = c

    var _fields = new TreeMap[String, FieldInfo]()

    def fields(): Collection[FieldInfo] = _fields.values()

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
            var field = new FieldInfo()
            field.deserialize(input, _constants)
            if (_fields.containsKey(field.name())) {
                throw new Exception("duplicate field name: " + field.name())
            }
            _fields.put(field.name(), field)
        }
    }
}
