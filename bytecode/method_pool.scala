import java.io.DataInputStream
import java.io.DataOutputStream


class MethodPool(c: ClassInfo) {
    var _owner = c

    def serialize(output: DataOutputStream) {
        // TODO
    }

    def deserialize(input: DataInputStream, constants: ConstantPool) {
        // TODO
    }
}
