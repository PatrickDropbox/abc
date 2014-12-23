import java.io.DataInputStream
import java.io.DataOutputStream
import java.util.Vector

import scala.collection.JavaConversions._


trait AttributeOwner {
    def constants(): ConstantPool
}

object Attribute {
    def deserialize(
            inputStream: DataInputStream,
            owner: AttributeOwner): Vector[Attribute] = {
        val attrCount = inputStream.readUnsignedShort()
        for (_ <- 1 to attrCount) {
            // TODO
            val nameIndex = inputStream.readUnsignedShort()
            val attrLength = inputStream.readInt()
            var buffer = new Array[Byte](attrLength)
            inputStream.readFully(buffer)
        }
        return null
    }
}

trait Attribute {
}

class ClassAttributes(c: ClassInfo) {
    var _owner = c

    def serialize(output: DataOutputStream) {
        // TODO
    }

    def deserialize(input: DataInputStream) {
        // TODO
        Attribute.deserialize(input, _owner)
    }
}

class FieldAttributes(f: FieldInfo) {
    var _owner = f

    def serialize(output: DataOutputStream) {
        // TODO
    }

    def deserialize(input: DataInputStream) {
        // TODO
        Attribute.deserialize(input, _owner._owner)
    }
}

class MethodAttributes(f: MethodInfo) {
    var _owner = f

    def serialize(output: DataOutputStream) {
        // TODO
    }

    def deserialize(input: DataInputStream) {
        // TODO
        Attribute.deserialize(input, _owner._owner)
    }
}
