import java.io.DataInputStream
import java.io.DataOutputStream


class Goto(owner: MethodInfo, pc: Int)
        extends ShortOperandOp(owner, OpCode.GOTO, "goto", false, pc) {
    def this(owner: MethodInfo) = this(owner, -1)

    override def serialize(output: DataOutputStream) {
        if (pc <= 65535) {
            super.serialize(output)
        } else {
            output.writeByte(OpCode.GOTO_W)
            output.writeInt(operand)
        }
    }
}

class GotoW(owner: MethodInfo, pc: Int)
        extends IntOperandOp(owner, OpCode.GOTO_W, "goto_w", pc) {
    def this(owner: MethodInfo) = this(owner, -1)

    override def canonicalForm(): Operation = {
        return new Goto(_owner, operand)
    }
}

// DEPRECATED: this is only kept around for deserializing older classes
// stack: ... -> ..., address
class Jsr(owner: MethodInfo, pc: Int)
        extends ShortOperandOp(owner, OpCode.JSR, "jsr", false, pc) {
    def this(owner: MethodInfo) = this(owner, -1)

    override def serialize(output: DataOutputStream) {
        throw new Exception("jsr deprecated")
    }
}

// DEPRECATED: this is only kept around for deserializing older classes
// stack: ... -> ..., address
class JsrW(owner: MethodInfo, pc: Int)
        extends IntOperandOp(owner, OpCode.JSR, "jsr_w", pc) {
    def this(owner: MethodInfo) = this(owner, -1)

    override def serialize(output: DataOutputStream) {
        throw new Exception("jsr_w deprecated")
    }

    override def canonicalForm(): Operation = {
        return new Jsr(_owner, operand)
    }
}

// DEPRECATED: this is only kept around for deserializing older classes
class Ret(owner: MethodInfo, index: Int)
        extends ByteOperandOp(owner, OpCode.RET, "ret", false, index) {
    def this(owner: MethodInfo) = this(owner, -1)

    override def serialize(output: DataOutputStream) {
        throw new Exception("ret deprecated")
    }
}


// return void
class Return(owner: MethodInfo)
        extends NoOperandOp(owner, OpCode.RETURN, "return") {
}

// stack: ..., value -> ...
abstract class ReturnValue(
        owner: MethodInfo,
        opCode: Int,
        mnemonic: String) extends NoOperandOp(owner, opCode, mnemonic) {
}

class Ireturn(owner: MethodInfo)
        extends ReturnValue(owner, OpCode.IRETURN, "ireturn") {
}

class Lreturn(owner: MethodInfo)
        extends ReturnValue(owner, OpCode.LRETURN, "lreturn") {
}

class Freturn(owner: MethodInfo)
        extends ReturnValue(owner, OpCode.FRETURN, "freturn") {
}

class Dreturn(owner: MethodInfo)
        extends ReturnValue(owner, OpCode.DRETURN, "dreturn") {
}

class Areturn(owner: MethodInfo)
        extends ReturnValue(owner, OpCode.ARETURN, "areturn") {
}


// TODO: lookup / table switches
// for serialization, use lookup switch when
//      8* (# entries) + 8 < 4 * (high - low + 1) + 12
// otherwise use table switch
