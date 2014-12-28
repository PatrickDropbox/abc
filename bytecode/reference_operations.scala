import java.io.DataInputStream
import java.io.DataOutputStream


// stack: ... -> ..., value
class Getstatic(
        owner: MethodInfo,
        className: String,
        fieldName: String,
        fieldType: FieldType) extends ShortOperandOp(
                owner,
                OpCode.GETSTATIC,
                "getstatic",
                false,
                0) {
    def this(owner: MethodInfo) = this(owner, null, null, null)

    var _constFieldRef: ConstFieldRefInfo = null
    if (className != null) {
        _constFieldRef = _owner.constants().getFieldRef(
                className,
                fieldName,
                fieldType)
    }

    override def serialize(output: DataOutputStream) {
        operand = _constFieldRef.index
        super.serialize(output)
    }

    override def deserialize(opCode: Int, input: DataInputStream) {
        super.deserialize(opCode, input)
        _constFieldRef = _owner.constants().getFieldRefByIndex(operand)
    }

    override def debugString(): String = {
        var name = "???"
        if (_constFieldRef != null) {
            name = _constFieldRef.debugString()
        }
        return "getstatic " + name
    }
}

// stack: ... -> ..., obj ref
class New(owner: MethodInfo, className: String)
        extends ShortOperandOp(owner, OpCode.NEW, "new", false, 0) {
    def this(owner: MethodInfo) = this(owner, null)

    var _constClass: ConstClassInfo = null
    if (className != null) {
        _constClass = _owner.constants().getClass(className)
    }

    override def serialize(output: DataOutputStream) {
        operand = _constClass.index
        super.serialize(output)
    }

    override def deserialize(opCode: Int, input: DataInputStream) {
        super.deserialize(opCode, input)
        _constClass = _owner.constants().getClassByIndex(operand)
    }

    override def debugString(): String = {
        var name = "???"
        if (_constClass != null) {
            name = _constClass.className()
        }
        return "new " + name
    }
}

// stack: ..., array ref -> ...
class Arraylength(owner: MethodInfo)
        extends NoOperandOp(owner, OpCode.ARRAYLENGTH, "arraylength") {
}

// stack: ..., obj ref -> ...
class Monitorenter(owner: MethodInfo)
        extends NoOperandOp(owner, OpCode.MONITORENTER, "monitorenter") {
}

// stack: ..., obj ref -> ...
class Monitorexit(owner: MethodInfo)
        extends NoOperandOp(owner, OpCode.MONITOREXIT, "monitorexit") {
}
