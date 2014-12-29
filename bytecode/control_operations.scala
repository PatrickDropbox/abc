import java.io.DataInputStream
import java.io.DataOutputStream
import java.util.TreeMap


class Goto(owner: AttributeOwner,
           current: CodeBlock,
           target: CodeBlock)
        extends Operation(owner) {
    def this(owner: AttributeOwner, offset: Int) = {
        this(owner, null, null)
        _offset = offset
    }

    def this(owner: AttributeOwner) = this(owner, 0)

    var _currentBlock = current
    var _targetBlock = target
    var _offset = 0

    def serialize(output: DataOutputStream) {
        if (_currentBlock.segmentNumber + 1 == _targetBlock.segmentNumber) {
            // skip writing goto since the two code block are next to each other
            return
        }
        _offset = _targetBlock.pc - pc
        if (_offset <= 65535) {
            output.writeByte(OpCode.GOTO)
            output.writeShort(_offset)
        } else {
            output.writeByte(OpCode.GOTO_W)
            output.writeInt(_offset)
        }
    }

    def deserialize(startAddress: Int, opCode: Int, input: DataInputStream) {
        if (opCode != OpCode.GOTO) {
            throw new Exception("Unexpected op-code: " + opCode)
        }

        _offset = input.readShort()
    }

    override def bindBlockRefs(table: TreeMap[Int, CodeBlock]) {
        val entry = table.floorEntry(pc)
        if (entry == null) {
            throw new Exception("can't find current block")
        }
        _currentBlock = entry.getValue()

        _targetBlock = table.get(pc + _offset)
        if (_targetBlock == null) {
            throw new Exception("can't find target block")
        }
    }

    def debugString(indent: String): String = {
        return indent + _pcLine() + ": goto (offset) " + _offset
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
// => 2 * n < (h - l +2)
// otherwise use table switch
