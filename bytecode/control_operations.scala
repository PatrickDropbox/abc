import java.io.DataInputStream
import java.io.DataOutputStream
import java.util.TreeMap

import scala.collection.JavaConversions._


class Goto(owner: AttributeOwner,
           current: CodeBlock,
           target: CodeBlock)
        extends Operation(owner) {
    def this(owner: AttributeOwner, offset: Int) = {
        this(owner, null, null)
        _tmpOffset = offset
    }

    def this(owner: AttributeOwner) = this(owner, 0)

    var _currentBlock = current
    var _targetBlock = target

    // only used during deserialization
    var _tmpOffset = 0

    def serialize(output: DataOutputStream) {
        if (_currentBlock.segmentId + 1 == _targetBlock.segmentId) {
            // skip writing goto since the two code block are next to each other
            return
        }
        val offset = _targetBlock.pc - pc
        if (offset > Const.UINT16_MAX) {
            throw new Exception("code too large (goto_w not supported)")
        }
        output.writeByte(OpCode.GOTO)
        output.writeShort(offset)
    }

    def deserialize(startAddress: Int, opCode: Int, input: DataInputStream) {
        if (opCode != OpCode.GOTO) {
            throw new Exception("Unexpected op-code: " + opCode)
        }

        _tmpOffset = input.readShort()
    }

    override def bindBlockRefs(table: TreeMap[Int, CodeBlock]) {
        val entry = table.floorEntry(pc)
        if (entry == null) {
            throw new Exception("can't find current block")
        }
        _currentBlock = entry.getValue()

        _targetBlock = table.get(pc + _tmpOffset)
        if (_targetBlock == null) {
            throw new Exception("can't find target block")
        }
    }

    def debugString(indent: String): String = {
        return indent + _pcLine() + ": goto " + _targetBlock.pc + "\n"
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

class Switch(owner: AttributeOwner, defaultBranch: CodeBlock)
        extends Operation(owner) {
    def this(owner: AttributeOwner) = this(owner, null)

    var _defaultBranch = defaultBranch
    var _table = new TreeMap[Int, CodeBlock]()

    // only used during deserialization
    var _tmpDefaultOffset = 0
    var _tmpOffset: TreeMap[Int, Int] = null

    def add(i: Int, branch: CodeBlock) {
        if (branch != _defaultBranch) {
            _table.put(i, branch)
        }
    }

    // Use lookup switch instead of table switch when lookup is more compact:
    //  Table switch size = 4 * (high - low + 1) + 12
    //  Lookup switch size = 8 * num entries + 8
    //      4 * (high - low + 1) + 12 <= 8 * num entries + 8
    //  ->  high - low + 2 <= 2 * num entries
    def _useTableSwitch(): Boolean = {
        if (_table.isEmpty()) {
            throw new Exception("No switch cases ...")
        }

        val low = _table.firstEntry().getKey()
        val high = _table.lastEntry().getKey()

        return (high - low + 2) <= (2 * _table.size())
    }

    def _paddingSize(startAddress: Int): Int = {
        return (4 - ((startAddress + 1) % 4)) % 4
    }

    def serialize(output: DataOutputStream) {
        if (_useTableSwitch()) {
            _serializeTableSwitch(output)
        } else {
            _serializeLookupSwitch(output)
        }
    }

    def _serializeLookupSwitch(output: DataOutputStream) {
        output.writeByte(OpCode.TABLESWITCH)
        for (_ <- 1 to _paddingSize(pc)) {
            output.writeByte(0)
        }

        output.writeInt(_defaultBranch.pc - pc)  // default offset

        output.writeInt(_table.size())
        for (entry <- _table.entrySet()) {
            output.writeInt(entry.getKey())
            output.writeInt(entry.getValue().pc - pc)
        }
    }

    def _serializeTableSwitch(output: DataOutputStream) {
        output.writeByte(OpCode.TABLESWITCH)
        for (_ <- 1 to _paddingSize(pc)) {
            output.writeByte(0)
        }

        val low = _table.firstEntry().getKey()
        val high = _table.lastEntry().getKey()

        output.writeInt(_defaultBranch.pc - pc)  // default offset
        output.writeInt(low)
        output.writeInt(high)

        for (i <- low to high) {
            val branch = _table.get(i)
            if (branch == null) {
                output.writeInt(_defaultBranch.pc - pc)
            } else {
                output.writeInt(branch.pc - pc)
            }
        }
    }

    def deserialize(startAddress: Int, opCode: Int, input: DataInputStream) {
        if (!_table.isEmpty()) {
            throw new Exception("deserializing into non-empty switch")
        }

        opCode match {
            case OpCode.TABLESWITCH =>
                    _deserializeTableSwitch(startAddress, input)
            case OpCode.LOOKUPSWITCH =>
                    _deserializeLookupSwitch(startAddress, input)
            case _ => throw new Exception("Unexpected op code: " + opCode)
        }
    }

    def _deserializeTableSwitch(startAddress: Int, input: DataInputStream) {
        input.skipBytes(_paddingSize(startAddress))

        _tmpDefaultOffset = input.readInt()
        val low = input.readInt()
        val high = input.readInt()

        _tmpOffset = new TreeMap[Int, Int]()
        for (i <- low to high) {
            val offset = input.readInt()
            if (offset != _tmpDefaultOffset) {
                _tmpOffset.put(i, offset)
            }
        }
    }

    def _deserializeLookupSwitch(startAddress: Int, input: DataInputStream) {
        input.skipBytes(_paddingSize(startAddress))

        _tmpDefaultOffset = input.readInt()

        val numEntries = input.readInt()

        _tmpOffset = new TreeMap[Int, Int]()
        for (_ <- 1 to numEntries) {
            val value = input.readInt()
            val offset = input.readInt()
            if (offset != _tmpDefaultOffset) {
                _tmpOffset.put(value, offset)
            }
        }
    }

    override def bindBlockRefs(table: TreeMap[Int, CodeBlock]) {
        _defaultBranch = table.get(pc + _tmpDefaultOffset)
        if (_defaultBranch == null) {
            throw new Exception("can't find switch default block")
        }

        for (entry <- _tmpOffset.entrySet()) {
            val block = table.get(pc + entry.getValue())
            if (block == null) {
                throw new Exception("can't find switch case block")
            }
            _table.put(entry.getKey(), block)
        }

        // free up tmp map
        _tmpOffset = null
    }

    def debugString(indent: String): String = {
        var result = indent + _pcLine() + ": switch\n"
        result += indent + "  default: " + _defaultBranch.pc + "\n"
        for (entry <- _table.entrySet()) {
            result += indent + "  case " + entry.getKey() + ": " +
                    entry.getValue().pc + "\n"
        }
        return result
    }
}

