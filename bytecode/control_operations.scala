import java.io.DataInputStream
import java.io.DataOutputStream


class Goto(owner: AttributeOwner, pcOffset: Int)
        extends ShortOperandOp(owner, OpCode.GOTO, "goto", false, pcOffset) {
    def this(owner: AttributeOwner) = this(owner, -1)

    override def serialize(output: DataOutputStream) {
        if (pcOffset <= 65535) {
            super.serialize(output)
        } else {
            output.writeByte(OpCode.GOTO_W)
            output.writeInt(operand)
        }
    }
}

class GotoW(owner: AttributeOwner, pcOffset: Int)
        extends IntOperandOp(owner, OpCode.GOTO_W, "goto_w", pcOffset) {
    def this(owner: AttributeOwner) = this(owner, -1)

    override def canonicalForm(): Operation = {
        return new Goto(_owner, operand)
    }
}

// DEPRECATED: this is only kept around for deserializing older classes
// stack: ... -> ..., address
class Jsr(owner: AttributeOwner, pc: Int)
        extends ShortOperandOp(owner, OpCode.JSR, "jsr", false, pc) {
    def this(owner: AttributeOwner) = this(owner, -1)

    override def serialize(output: DataOutputStream) {
        throw new Exception("jsr deprecated")
    }
}

// DEPRECATED: this is only kept around for deserializing older classes
// stack: ... -> ..., address
class JsrW(owner: AttributeOwner, pc: Int)
        extends IntOperandOp(owner, OpCode.JSR, "jsr_w", pc) {
    def this(owner: AttributeOwner) = this(owner, -1)

    override def serialize(output: DataOutputStream) {
        throw new Exception("jsr_w deprecated")
    }

    override def canonicalForm(): Operation = {
        return new Jsr(_owner, operand)
    }
}

// DEPRECATED: this is only kept around for deserializing older classes
class Ret(owner: AttributeOwner, index: Int)
        extends ByteOperandOp(owner, OpCode.RET, "ret", false, index) {
    def this(owner: AttributeOwner) = this(owner, -1)

    override def serialize(output: DataOutputStream) {
        throw new Exception("ret deprecated")
    }
}


// return void
class Return(owner: AttributeOwner)
        extends NoOperandOp(owner, OpCode.RETURN, "return") {
}

// stack: ..., value -> ...
abstract class ReturnValue(
        owner: AttributeOwner,
        opCode: Int,
        mnemonic: String) extends NoOperandOp(owner, opCode, mnemonic) {
}

class Ireturn(owner: AttributeOwner)
        extends ReturnValue(owner, OpCode.IRETURN, "ireturn") {
}

class Lreturn(owner: AttributeOwner)
        extends ReturnValue(owner, OpCode.LRETURN, "lreturn") {
}

class Freturn(owner: AttributeOwner)
        extends ReturnValue(owner, OpCode.FRETURN, "freturn") {
}

class Dreturn(owner: AttributeOwner)
        extends ReturnValue(owner, OpCode.DRETURN, "dreturn") {
}

class Areturn(owner: AttributeOwner)
        extends ReturnValue(owner, OpCode.ARETURN, "areturn") {
}


// TODO: lookup / table switches
// for serialization, use lookup switch when
//      8* (# entries) + 8 < 4 * (high - low + 1) + 12
// otherwise use table switch
