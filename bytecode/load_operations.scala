import java.io.DataInputStream
import java.io.DataOutputStream

abstract class _Load(
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
// iload <local var index>
// stack: ..., value -> ...
//
class LoadI(owner: MethodInfo, index: Int)
        extends _Load(owner, OpCode.ILOAD, OpCode.ILOAD_0, "iload", index) {
    def this(owner: MethodInfo) = this(owner, 0)
}

//
// lload <local var index>
// stack: ..., value -> ...
//
class LoadL(owner: MethodInfo, index: Int)
        extends _Load(owner, OpCode.LLOAD, OpCode.LLOAD_0, "lload", index) {
    def this(owner: MethodInfo) = this(owner, 0)
}

//
// fload <local var index>
// stack: ..., value -> ...
//
class LoadF(owner: MethodInfo, index: Int)
        extends _Load(owner, OpCode.FLOAD, OpCode.FLOAD_0, "fload", index) {
    def this(owner: MethodInfo) = this(owner, 0)
}

//
// dload <local var index>
// stack: ..., value -> ...
//
class LoadD(owner: MethodInfo, index: Int)
        extends _Load(owner, OpCode.DLOAD, OpCode.DLOAD_0, "dload", index) {
    def this(owner: MethodInfo) = this(owner, 0)
}

//
// aload <local var index>
// stack: ..., value -> ...
//
class LoadA(owner: MethodInfo, index: Int)
        extends _Load(owner, OpCode.ALOAD, OpCode.ALOAD_0, "aload", index) {
    def this(owner: MethodInfo) = this(owner, 0)
}

class Iload0(owner: MethodInfo)
        extends NoOperandOp(owner, OpCode.ILOAD_0, "iload_0") {
    override def canonicalForm(): Operation = {
        return new LoadI(owner, 0)
    }
}

class Iload1(owner: MethodInfo)
        extends NoOperandOp(owner, OpCode.ILOAD_1, "iload_1") {
    override def canonicalForm(): Operation = {
        return new LoadI(owner, 1)
    }
}

class Iload2(owner: MethodInfo)
        extends NoOperandOp(owner, OpCode.ILOAD_2, "iload_2") {
    override def canonicalForm(): Operation = {
        return new LoadI(owner, 2)
    }
}

class Iload3(owner: MethodInfo)
        extends NoOperandOp(owner, OpCode.ILOAD_3, "iload_3") {
    override def canonicalForm(): Operation = {
        return new LoadI(owner, 3)
    }
}

class Lload0(owner: MethodInfo)
        extends NoOperandOp(owner, OpCode.LLOAD_0, "lload_0") {
    override def canonicalForm(): Operation = {
        return new LoadL(owner, 0)
    }
}

class Lload1(owner: MethodInfo)
        extends NoOperandOp(owner, OpCode.LLOAD_1, "lload_1") {
    override def canonicalForm(): Operation = {
        return new LoadL(owner, 1)
    }
}

class Lload2(owner: MethodInfo)
        extends NoOperandOp(owner, OpCode.LLOAD_2, "lload_2") {
    override def canonicalForm(): Operation = {
        return new LoadL(owner, 2)
    }
}

class Lload3(owner: MethodInfo)
        extends NoOperandOp(owner, OpCode.LLOAD_3, "lload_3") {
    override def canonicalForm(): Operation = {
        return new LoadL(owner, 3)
    }
}

class Fload0(owner: MethodInfo)
        extends NoOperandOp(owner, OpCode.FLOAD_0, "fload_0") {
    override def canonicalForm(): Operation = {
        return new LoadF(owner, 0)
    }
}

class Fload1(owner: MethodInfo)
        extends NoOperandOp(owner, OpCode.FLOAD_1, "fload_1") {
    override def canonicalForm(): Operation = {
        return new LoadF(owner, 1)
    }
}

class Fload2(owner: MethodInfo)
        extends NoOperandOp(owner, OpCode.FLOAD_2, "fload_2") {
    override def canonicalForm(): Operation = {
        return new LoadF(owner, 2)
    }
}

class Fload3(owner: MethodInfo)
        extends NoOperandOp(owner, OpCode.FLOAD_3, "fload_3") {
    override def canonicalForm(): Operation = {
        return new LoadF(owner, 3)
    }
}

class Dload0(owner: MethodInfo)
        extends NoOperandOp(owner, OpCode.DLOAD_0, "dload_0") {
    override def canonicalForm(): Operation = {
        return new LoadD(owner, 0)
    }
}

class Dload1(owner: MethodInfo)
        extends NoOperandOp(owner, OpCode.DLOAD_1, "dload_1") {
    override def canonicalForm(): Operation = {
        return new LoadD(owner, 1)
    }
}

class Dload2(owner: MethodInfo)
        extends NoOperandOp(owner, OpCode.DLOAD_2, "dload_2") {
    override def canonicalForm(): Operation = {
        return new LoadD(owner, 2)
    }
}

class Dload3(owner: MethodInfo)
        extends NoOperandOp(owner, OpCode.DLOAD_3, "dload_3") {
    override def canonicalForm(): Operation = {
        return new LoadD(owner, 3)
    }
}

class Aload0(owner: MethodInfo)
        extends NoOperandOp(owner, OpCode.ALOAD_0, "aload_0") {
    override def canonicalForm(): Operation = {
        return new LoadA(owner, 0)
    }
}

class Aload1(owner: MethodInfo)
        extends NoOperandOp(owner, OpCode.ALOAD_1, "aload_1") {
    override def canonicalForm(): Operation = {
        return new LoadA(owner, 1)
    }
}

class Aload2(owner: MethodInfo)
        extends NoOperandOp(owner, OpCode.ALOAD_2, "aload_2") {
    override def canonicalForm(): Operation = {
        return new LoadA(owner, 2)
    }
}

class Aload3(owner: MethodInfo)
        extends NoOperandOp(owner, OpCode.ALOAD_3, "aload_3") {
    override def canonicalForm(): Operation = {
        return new LoadA(owner, 3)
    }
}

abstract class _LoadFromArray(
        owner: MethodInfo,
        opCode: Int,
        mnemonic: String) extends NoOperandOp(owner, opCode, mnemonic) {
}

//
// iaload
// stack: ..., arrayref, index -> ..., value
//
class LoadFromIArray(owner: MethodInfo)
        extends _LoadFromArray(owner, OpCode.IALOAD, "iaload") {
}

//
// laload
// stack: ..., arrayref, index -> ..., value
//
class LoadFromLArray(owner: MethodInfo)
        extends _LoadFromArray(owner, OpCode.LALOAD, "laload") {
}

//
// faload
// stack: ..., arrayref, index -> ..., value
//
class LoadFromFArray(owner: MethodInfo)
        extends _LoadFromArray(owner, OpCode.FALOAD, "faload") {
}

//
// daload
// stack: ..., arrayref, index -> ..., value
//
class LoadFromDArray(owner: MethodInfo)
        extends _LoadFromArray(owner, OpCode.DALOAD, "daload") {
}

//
// aaload
// stack: ..., arrayref, index -> ..., value
//
class LoadFromAArray(owner: MethodInfo)
        extends _LoadFromArray(owner, OpCode.AALOAD, "aaload") {
}

//
// baload
// stack: ..., arrayref, index -> ..., value
//
class LoadFromBArray(owner: MethodInfo)
        extends _LoadFromArray(owner, OpCode.BALOAD, "baload") {
}

//
// caload
// stack: ..., arrayref, index -> ..., value
//
class LoadFromCArray(owner: MethodInfo)
        extends _LoadFromArray(owner, OpCode.CALOAD, "caload") {
}

//
// saload
// stack: ..., arrayref, index -> ..., value
//
class LoadFromSArray(owner: MethodInfo)
        extends _LoadFromArray(owner, OpCode.SALOAD, "saload") {
}
