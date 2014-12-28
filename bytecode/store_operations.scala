import java.io.DataInputStream
import java.io.DataOutputStream

abstract class Store(
        owner: MethodInfo,
        opCode: Int,
        shortOpCodeStart: Int,
        mnemonic: String,
        v: Int) extends ByteOperandOp(owner, opCode, mnemonic, false, v) {
    val _shortOpCodeStart = shortOpCodeStart

    override def serialize(output: DataOutputStream) {
        operand match {
            case 0 => output.writeByte(_shortOpCodeStart)
            case 1 => output.writeByte(_shortOpCodeStart + 1)
            case 2 => output.writeByte(_shortOpCodeStart + 2)
            case 3 => output.writeByte(_shortOpCodeStart + 3)
            case _ => {
                if (operand < 256) {
                    super.serialize(output)
                } else {
                    output.writeByte(OpCode.WIDE)
                    output.writeByte(_opCode)
                    output.writeShort(operand)
                }
            }
        }
    }
}

//
// istore <local var index>
// stack: ..., value -> ...
//
class StoreI(owner: MethodInfo, index: Int)
        extends Store(owner, OpCode.ISTORE, OpCode.ISTORE_0, "istore", index) {
    def this(owner: MethodInfo) = this(owner, 0)
}

//
// lstore <local var index>
// stack: ..., value -> ...
//
class StoreL(owner: MethodInfo, index: Int)
        extends Store(owner, OpCode.LSTORE, OpCode.LSTORE_0, "lstore", index) {
    def this(owner: MethodInfo) = this(owner, 0)
}

//
// fstore <local var index>
// stack: ..., value -> ...
//
class StoreF(owner: MethodInfo, index: Int)
        extends Store(owner, OpCode.FSTORE, OpCode.FSTORE_0, "fstore", index) {
    def this(owner: MethodInfo) = this(owner, 0)
}

//
// dstore <local var index>
// stack: ..., value -> ...
//
class StoreD(owner: MethodInfo, index: Int)
        extends Store(owner, OpCode.DSTORE, OpCode.DSTORE_0, "dstore", index) {
    def this(owner: MethodInfo) = this(owner, 0)
}

//
// astore <local var index>
// stack: ..., value -> ...
//
class StoreA(owner: MethodInfo, index: Int)
        extends Store(owner, OpCode.ASTORE, OpCode.ASTORE_0, "astore", index) {
    def this(owner: MethodInfo) = this(owner, 0)
}

class Istore0(owner: MethodInfo)
        extends NoOperandOp(owner, OpCode.ISTORE_0, "istore_0") {
    override def canonicalForm(): Operation = {
        return new StoreI(owner, 0)
    }
}

class Istore1(owner: MethodInfo)
        extends NoOperandOp(owner, OpCode.ISTORE_1, "istore_1") {
    override def canonicalForm(): Operation = {
        return new StoreI(owner, 1)
    }
}

class Istore2(owner: MethodInfo)
        extends NoOperandOp(owner, OpCode.ISTORE_2, "istore_2") {
    override def canonicalForm(): Operation = {
        return new StoreI(owner, 2)
    }
}

class Istore3(owner: MethodInfo)
        extends NoOperandOp(owner, OpCode.ISTORE_3, "istore_3") {
    override def canonicalForm(): Operation = {
        return new StoreI(owner, 3)
    }
}

class Lstore0(owner: MethodInfo)
        extends NoOperandOp(owner, OpCode.LSTORE_0, "lstore_0") {
    override def canonicalForm(): Operation = {
        return new StoreL(owner, 0)
    }
}

class Lstore1(owner: MethodInfo)
        extends NoOperandOp(owner, OpCode.LSTORE_1, "lstore_1") {
    override def canonicalForm(): Operation = {
        return new StoreL(owner, 1)
    }
}

class Lstore2(owner: MethodInfo)
        extends NoOperandOp(owner, OpCode.LSTORE_2, "lstore_2") {
    override def canonicalForm(): Operation = {
        return new StoreL(owner, 2)
    }
}

class Lstore3(owner: MethodInfo)
        extends NoOperandOp(owner, OpCode.LSTORE_3, "lstore_3") {
    override def canonicalForm(): Operation = {
        return new StoreL(owner, 3)
    }
}

class Fstore0(owner: MethodInfo)
        extends NoOperandOp(owner, OpCode.FSTORE_0, "fstore_0") {
    override def canonicalForm(): Operation = {
        return new StoreF(owner, 0)
    }
}

class Fstore1(owner: MethodInfo)
        extends NoOperandOp(owner, OpCode.FSTORE_1, "fstore_1") {
    override def canonicalForm(): Operation = {
        return new StoreF(owner, 1)
    }
}

class Fstore2(owner: MethodInfo)
        extends NoOperandOp(owner, OpCode.FSTORE_2, "fstore_2") {
    override def canonicalForm(): Operation = {
        return new StoreF(owner, 2)
    }
}

class Fstore3(owner: MethodInfo)
        extends NoOperandOp(owner, OpCode.FSTORE_3, "fstore_3") {
    override def canonicalForm(): Operation = {
        return new StoreF(owner, 3)
    }
}

class Dstore0(owner: MethodInfo)
        extends NoOperandOp(owner, OpCode.DSTORE_0, "dstore_0") {
    override def canonicalForm(): Operation = {
        return new StoreD(owner, 0)
    }
}

class Dstore1(owner: MethodInfo)
        extends NoOperandOp(owner, OpCode.DSTORE_1, "dstore_1") {
    override def canonicalForm(): Operation = {
        return new StoreD(owner, 1)
    }
}

class Dstore2(owner: MethodInfo)
        extends NoOperandOp(owner, OpCode.DSTORE_2, "dstore_2") {
    override def canonicalForm(): Operation = {
        return new StoreD(owner, 2)
    }
}

class Dstore3(owner: MethodInfo)
        extends NoOperandOp(owner, OpCode.DSTORE_3, "dstore_3") {
    override def canonicalForm(): Operation = {
        return new StoreD(owner, 3)
    }
}

class Astore0(owner: MethodInfo)
        extends NoOperandOp(owner, OpCode.ASTORE_0, "astore_0") {
    override def canonicalForm(): Operation = {
        return new StoreA(owner, 0)
    }
}

class Astore1(owner: MethodInfo)
        extends NoOperandOp(owner, OpCode.ASTORE_1, "astore_1") {
    override def canonicalForm(): Operation = {
        return new StoreA(owner, 1)
    }
}

class Astore2(owner: MethodInfo)
        extends NoOperandOp(owner, OpCode.ASTORE_2, "astore_2") {
    override def canonicalForm(): Operation = {
        return new StoreA(owner, 2)
    }
}

class Astore3(owner: MethodInfo)
        extends NoOperandOp(owner, OpCode.ASTORE_3, "astore_3") {
    override def canonicalForm(): Operation = {
        return new StoreA(owner, 3)
    }
}

abstract class StoreIntoArray(
        owner: MethodInfo,
        opCode: Int,
        mnemonic: String) extends NoOperandOp(owner, opCode, mnemonic) {
}

//
// iastore
// stack: ..., arrayref, index, value -> ...
//
class StoreIntoIArray(owner: MethodInfo)
        extends StoreIntoArray(owner, OpCode.IASTORE, "iastore") {
}

//
// lastore
// stack: ..., arrayref, index, value -> ...
//
class StoreIntoLArray(owner: MethodInfo)
        extends StoreIntoArray(owner, OpCode.LASTORE, "lastore") {
}

//
// fastore
// stack: ..., arrayref, index, value -> ...
//
class StoreIntoFArray(owner: MethodInfo)
        extends StoreIntoArray(owner, OpCode.FASTORE, "fastore") {
}

//
// dastore
// stack: ..., arrayref, index, value -> ...
//
class StoreIntoDArray(owner: MethodInfo)
        extends StoreIntoArray(owner, OpCode.DASTORE, "dastore") {
}

//
// aastore
// stack: ..., arrayref, index, value -> ...
//
class StoreIntoAArray(owner: MethodInfo)
        extends StoreIntoArray(owner, OpCode.AASTORE, "aastore") {
}

//
// bastore
// stack: ..., arrayref, index, value -> ...
//
class StoreIntoBArray(owner: MethodInfo)
        extends StoreIntoArray(owner, OpCode.BASTORE, "bastore") {
}

//
// castore
// stack: ..., arrayref, index, value -> ...
//
class StoreIntoCArray(owner: MethodInfo)
        extends StoreIntoArray(owner, OpCode.CASTORE, "castore") {
}

//
// sastore
// stack: ..., arrayref, index, value -> ...
//
class StoreIntoSArray(owner: MethodInfo)
        extends StoreIntoArray(owner, OpCode.SASTORE, "sastore") {
}
