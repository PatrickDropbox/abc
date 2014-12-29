import java.io.DataInputStream
import java.io.DataOutputStream
import java.util.Vector

import scala.collection.JavaConversions._


class ExceptionEntry(
        owner: AttributeOwner,
        spc: Int,
        epc: Int,
        hpc: Int,
        className: String) {
    def this(owner: AttributeOwner) = this(owner, 0, 0, 0, null)

    var _owner = owner

    var startPc = spc  // inclusive
    var endPc = epc  // exclusive
    var handlerPc = hpc
    var classType: ConstClassInfo = null  // null means catch all
    if (className != null) {
        classType = _owner.constants().getClass(className)
    }

    def serialize(output: DataOutputStream) {
        output.writeShort(startPc)
        output.writeShort(endPc)
        output.writeShort(handlerPc)
        if (classType == null) {
            output.writeShort(0)
        } else {
            output.writeShort(classType.index)
        }
    }

    def deserialize(input: DataInputStream) {
        startPc = input.readUnsignedShort()
        endPc = input.readUnsignedShort()
        handlerPc = input.readUnsignedShort()

        val index = input.readUnsignedShort()
        if (index == 0) {
            classType = null
        } else {
            classType = _owner.constants().getClassByIndex(index)
        }
    }

    def debugString(indent: String): String = {
        var name = "(any)"
        if (classType != null) {
            name = classType.className()
        }

        var result = indent + "[" + startPc + ", " + endPc + ") -> " + handlerPc
        result += " for " + name + "\n"
        return result
    }
}

class CodeAttribute(o: AttributeOwner)
        extends Attribute(o, "Code")
        with AttributeOwner {

    var maxStack = 0
    var maxLocals = 0

    var operations = new Vector[Operation]()

    var exceptionEntries = new Vector[ExceptionEntry]()

    var attributes = new CodeAttributes(this)

    def constants(): ConstantPool = _owner.constants()

    def serialize(output: DataOutputStream) {
        // TODO
    }

    def deserialize(name: ConstUtf8Info,
                    attrLength: Int,
                    input: DataInputStream) {
        if (!operations.isEmpty()) {
            throw new Exception("deserializing into non-empty code attribute")
        }

        maxStack = input.readUnsignedShort()
        maxLocals = input.readUnsignedShort()

        val codeLength = input.readInt()
        var codeBytes = new Array[Byte](codeLength)
        input.readFully(codeBytes)

        operations = Operation.deserialize(this, codeBytes)

        exceptionEntries = new Vector[ExceptionEntry]()
        val numExceptionEntries = input.readUnsignedShort()
        for (_ <- 1 to numExceptionEntries) {
            var entry = new ExceptionEntry(_owner)
            entry.deserialize(input)
            exceptionEntries.add(entry)
        }

        attributes = new CodeAttributes(this)
        attributes.deserialize(input)
    }

    def debugString(indent: String): String = {
        var result = indent + "Code:\n"
        result += indent + "  Max stack: " + maxStack + "\n"
        result += indent + "  Max locals: " + maxLocals + "\n"
        val subIndent = indent + "    "
        for (op <- operations) {
            result += op.debugString(subIndent)
        }
        result += indent + "  Exceptions:\n"
        for (entry <- exceptionEntries) {
            result += entry.debugString(subIndent)
        }
        result += indent + "  Attributes:\n"
        result += attributes.debugString(subIndent)

        return result
    }
}
