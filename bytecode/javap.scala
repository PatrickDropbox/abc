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

        println("Classfile " + file.getAbsolutePath())
        println("  minor version: " + classInfo.minorVersion())
        println("  major version: " + classInfo.majorVersion())
        println("  flags: " + classInfo.access.debugString())

        println("Constants:")
        for (info <- classInfo.constants()._tmpConstInfosByIndex.values()) {
            println(info.debugString())
        }
        println("Fields:")
        for (field <- classInfo.fields.fields()) {
            println("  " + field.name())
            println("    type: " + field.descriptor().descriptorString())
            println("    flags: " + field.access().debugString())
        }
    }
}
