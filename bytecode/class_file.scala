import java.io.DataInputStream
import java.io.DataOutputStream
import java.util.Vector

import scala.collection.JavaConversions._


object ClassFile {
    val MAGIC = 0xcafebabe
}

class ClassFile {
    var minorVersion = 0
    var majorVersion = 51  // jvm7

    var constants = new ConstantPool()

    var access = new ClassAccessFlags()

    var thisClass: ConstClassInfo = null  // the current class
    var superClass: ConstClassInfo = null  // null if Object; non-null otherwise
    var interfaces = new Vector[ConstClassInfo]()

    var fields = new FieldPool()
    var methods = new MethodPool()
    var attributes = new ClassAttributes()

    def serialize(output: DataOutputStream) {
        output.writeInt(ClassFile.MAGIC)

        output.writeShort(minorVersion)
        output.writeShort(majorVersion)

        constants.serialize(output)

        access.serialize(output)

        output.writeShort(thisClass.index)
        var superClassIndex = 0
        if (superClass != null) {  // i.e., not Object
            superClassIndex = superClass.index
        }
        output.writeShort(superClassIndex)
        output.writeShort(interfaces.length)
        for (iface <- interfaces) {
            output.writeShort(iface.index)
        }

        fields.serialize(output)
        methods.serialize(output)
        attributes.serialize(output)
    }

    def deserialize(input: DataInputStream) {
        if (input.readInt() != ClassFile.MAGIC) {
            throw new Exception("Invalid magic")
        }

        minorVersion = input.readUnsignedShort()
        majorVersion = input.readUnsignedShort()

        constants.deserialize(input)

        access.deserialize(input)

        val thisClassIndex = input.readUnsignedShort()
        thisClass = constants.getClassByIndex(thisClassIndex)

        val superClassIndex = input.readUnsignedShort()
        if (superClassIndex == 0) {
            superClass = null
        } else {
            superClass = constants.getClassByIndex(superClassIndex)
        }

        val interfacesCount = input.readUnsignedShort()
        for (_ <- 1 to interfacesCount) {
            interfaces.add(constants.getClassByIndex(input.readUnsignedShort()))
        }

        fields.deserialize(input, constants)
        methods.deserialize(input, constants)
        attributes.deserialize(input, constants)
    }
}
