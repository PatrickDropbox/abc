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
