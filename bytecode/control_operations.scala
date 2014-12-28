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

