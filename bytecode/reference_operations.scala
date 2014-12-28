import java.io.DataInputStream
import java.io.DataOutputStream


class FieldOp(
        owner: MethodInfo,
        opCode: Int,
        mnemonic: String,
        className: String,
        fieldName: String,
        fieldType: FieldType) extends Operation(owner) {
    val _opCode = opCode
    val _mnemonic = mnemonic

    var _constFieldRef: ConstFieldRefInfo = null
    if (className != null) {
        _constFieldRef = _owner.constants().getFieldRef(
                className,
                fieldName,
                fieldType)
    }

    def serialize(output: DataOutputStream) {
        output.writeByte(_opCode)
        output.writeShort(_constFieldRef.index)
    }

    def deserialize(startAddress: Int, opCode: Int, input: DataInputStream) {
        if (opCode != _opCode) {
            throw new Exception("Unexpected op-code: " + opCode)
        }

        val index = input.readUnsignedShort()
        _constFieldRef = _owner.constants().getFieldRefByIndex(index)
    }

    def debugString(indent: String): String = {
        var name = "???"
        if (_constFieldRef != null) {
            name = _constFieldRef.debugString()
        }
        return indent + _mnemonic + " " + name
    }
}

// stack: ... -> ..., value
class Getstatic(
        owner: MethodInfo,
        className: String,
        fieldName: String,
        fieldType: FieldType) extends FieldOp(
                owner,
                OpCode.GETSTATIC,
                "getstatic",
                className,
                fieldName,
                fieldType) {
    def this(owner: MethodInfo) = this(owner, null, null, null)
}

// stack: ..., value -> ...
class Putstatic(
        owner: MethodInfo,
        className: String,
        fieldName: String,
        fieldType: FieldType) extends FieldOp(
                owner,
                OpCode.PUTSTATIC,
                "putstatic",
                className,
                fieldName,
                fieldType) {
    def this(owner: MethodInfo) = this(owner, null, null, null)
}

// stack: ..., obj ref -> ..., value
class Getfield(
        owner: MethodInfo,
        className: String,
        fieldName: String,
        fieldType: FieldType) extends FieldOp(
                owner,
                OpCode.GETFIELD,
                "getfield",
                className,
                fieldName,
                fieldType) {
    def this(owner: MethodInfo) = this(owner, null, null, null)
}

// stack: ..., obj ref, value -> ...
class Putfield(
        owner: MethodInfo,
        className: String,
        fieldName: String,
        fieldType: FieldType) extends FieldOp(
                owner,
                OpCode.PUTFIELD,
                "putfield",
                className,
                fieldName,
                fieldType) {
    def this(owner: MethodInfo) = this(owner, null, null, null)
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

    override def deserialize(
            startAddress: Int,
            opCode: Int,
            input: DataInputStream) {
        super.deserialize(startAddress: Int, opCode, input)
        _constClass = _owner.constants().getClassByIndex(operand)
    }

    override def debugString(indent: String): String = {
        var name = "???"
        if (_constClass != null) {
            name = _constClass.className()
        }
        return indent + "new " + name
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
