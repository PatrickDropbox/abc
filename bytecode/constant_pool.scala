import java.io.DataInputStream
import java.io.DataOutputStream


class ConstantPool {

    def getClassByIndex(index: Int): ConstClassInfo = {
        // TODO
        return null
    }

    def serialize(output: DataOutputStream) {
        // TODO
    }

    def deserialize(input: DataInputStream) {
        _parse(input)
        _bindConstReferences()
    }

    def _parse(input: DataInputStream) {
        // TODO
    }

    def _bindConstReferences() {
        // TODO
    }

    // TODO invoke dynamic reference index
}
