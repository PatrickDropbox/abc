import java.io.DataInputStream
import java.io.DataOutputStream
import java.util.Vector

import scala.collection.JavaConversions._


object Attribute {
    def deserialize(
            inputStream: DataInputStream,
            constants: ConstantPool): Vector[Attribute] = {
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

class ClassAttributes {
    def serialize(output: DataOutputStream) {
        // TODO
    }

    def deserialize(input: DataInputStream, constants: ConstantPool) {
        // TODO
        Attribute.deserialize(input, constants)
    }
}

class FieldAttributes {
    def serialize(output: DataOutputStream) {
        // TODO
    }

    def deserialize(input: DataInputStream, constants: ConstantPool) {
        // TODO
        Attribute.deserialize(input, constants)
    }
}
