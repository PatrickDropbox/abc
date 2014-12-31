import java.io.DataInputStream
import java.io.File
import java.io.FileInputStream

import scala.collection.JavaConversions._


object Javap {
    def main(args: Array[String]) {
        if (args.length != 1) {
            println("USAGE: Javap <class file>")
            System.exit(1)
        }

        var file = new File(args(0))
        var input = new DataInputStream(new FileInputStream(file))

        var classInfo = new ClassInfo()
        try {
            classInfo.deserialize(input)
        } finally {
            input.close()
        }

// TODO: REMOVE HACK
for (m <- classInfo.methods.methods()) {
    var code = m.attributes().code.code
    code.serialize()
}



        println("Classfile " + file.getAbsolutePath())
        println("  Minor version: " + classInfo.minorVersion())
        println("  Major version: " + classInfo.majorVersion())
        println("  Flags: " + classInfo.access.debugString())

        println("\nAttributes:")
        println(classInfo.attributes().debugString("  "))

        println("Constants:")
        for (info <- classInfo.constants()._tmpConstInfosByIndex.values()) {
            println(info.debugString())
        }

        println("\nFields:")
        for (field <- classInfo.fields.fields()) {
            println("  " + field.name())
            println("    Type: " + field.fieldTypeString())
            println("    Flags: " + field.access().debugString())
            println("    Attributes:")
            println(field.attributes().debugString("      "))
        }

        println("Methods:")
        for (method <- classInfo.methods.methods()) {
            println("  " + method.name())
            println("    Type: " + method.methodTypeString())
            println("    Flags: " + method.access().debugString())
            println("    Attributes:")
            println(method.attributes().debugString("      "))
        }
    }
}
