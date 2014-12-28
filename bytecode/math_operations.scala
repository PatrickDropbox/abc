import java.io.DataInputStream
import java.io.DataOutputStream


// stack: ..., value1, value2 -> ..., result
abstract class BinaryIOp(owner: MethodInfo, opCode: Int, mnemonic: String)
        extends NoOperandOp(owner, opCode, mnemonic) {
}

// stack: ..., value1 -> ..., result
abstract class UnaryIOp(owner: MethodInfo, opCode: Int, mnemonic: String)
        extends NoOperandOp(owner, opCode, mnemonic) {
}

// stack: ..., value1, value2 -> ..., result
abstract class BinaryLOp(owner: MethodInfo, opCode: Int, mnemonic: String)
        extends NoOperandOp(owner, opCode, mnemonic) {
}

// stack: ..., value1 -> ..., result
abstract class UnaryLOp(owner: MethodInfo, opCode: Int, mnemonic: String)
        extends NoOperandOp(owner, opCode, mnemonic) {
}

// stack: ..., (long) value1, (int) value2 -> (long) result
abstract class ShiftLOp(owner: MethodInfo, opCode: Int, mnemonic: String)
        extends NoOperandOp(owner, opCode, mnemonic) {
}

// stack: ..., value1, value2 -> ..., result
abstract class BinaryFOp(owner: MethodInfo, opCode: Int, mnemonic: String)
        extends NoOperandOp(owner, opCode, mnemonic) {
}

// stack: ..., value1 -> ..., result
abstract class UnaryFOp(owner: MethodInfo, opCode: Int, mnemonic: String)
        extends NoOperandOp(owner, opCode, mnemonic) {
}

// stack: ..., value1, value2 -> ..., result
abstract class BinaryDOp(owner: MethodInfo, opCode: Int, mnemonic: String)
        extends NoOperandOp(owner, opCode, mnemonic) {
}

// stack: ..., value1 -> ..., result
abstract class UnaryDOp(owner: MethodInfo, opCode: Int, mnemonic: String)
        extends NoOperandOp(owner, opCode, mnemonic) {
}

//
// int operations
//

class Iadd(owner: MethodInfo) extends BinaryIOp(owner, OpCode.IADD, "iadd") {
}

class Isub(owner: MethodInfo) extends BinaryIOp(owner, OpCode.ISUB, "isub") {
}

class Imul(owner: MethodInfo) extends BinaryIOp(owner, OpCode.IMUL, "imul") {
}

class Idiv(owner: MethodInfo) extends BinaryIOp(owner, OpCode.IDIV, "idiv") {
}

class Irem(owner: MethodInfo) extends BinaryIOp(owner, OpCode.IREM, "irem") {
}

class Ineg(owner: MethodInfo) extends UnaryIOp(owner, OpCode.INEG, "ineg") {
}

class Ishl(owner: MethodInfo) extends BinaryIOp(owner, OpCode.ISHL, "ishl") {
}

class Ishr(owner: MethodInfo) extends BinaryIOp(owner, OpCode.ISHR, "ishr") {
}

class Iushr(owner: MethodInfo) extends BinaryIOp(owner, OpCode.IUSHR, "iushr") {
}

class Iand(owner: MethodInfo) extends BinaryIOp(owner, OpCode.IAND, "iand") {
}

class Ior(owner: MethodInfo) extends BinaryIOp(owner, OpCode.IOR, "ior") {
}

class Ixor(owner: MethodInfo) extends BinaryIOp(owner, OpCode.IXOR, "ixor") {
}

// iinc <local var index> <const int>
// increment local variable by const without modifying stack
class Iinc(owner: MethodInfo, index: Int, v: Int)
        extends TwoByteOperandsOp(
                owner,
                OpCode.IINC,
                "iinc",
                false,  // signed
                index,
                true,  // signed
                v) {
    def this(owner: MethodInfo) = this(owner, -1, 0)

    override def serialize(output: DataOutputStream) {
        if (operand1 < 256 && (-128 <= operand2 && operand2 <= 127)) {
            super.serialize(output)
        } else {
            output.writeByte(OpCode.WIDE)
            output.writeByte(OpCode.IINC)
            output.writeShort(operand1)
            output.writeShort(operand2)
        }
    }
}

//
// long operations
//

class Ladd(owner: MethodInfo) extends BinaryLOp(owner, OpCode.LADD, "ladd") {
}

class Lsub(owner: MethodInfo) extends BinaryLOp(owner, OpCode.LSUB, "lsub") {
}

class Lmul(owner: MethodInfo) extends BinaryLOp(owner, OpCode.LMUL, "lmul") {
}

class Ldiv(owner: MethodInfo) extends BinaryLOp(owner, OpCode.LDIV, "ldiv") {
}

class Lrem(owner: MethodInfo) extends BinaryLOp(owner, OpCode.LREM, "lrem") {
}

class Lneg(owner: MethodInfo) extends UnaryLOp(owner, OpCode.LNEG, "lneg") {
}

class Lshl(owner: MethodInfo) extends ShiftLOp(owner, OpCode.LSHL, "lshl") {
}

class Lshr(owner: MethodInfo) extends ShiftLOp(owner, OpCode.LSHR, "lshr") {
}

class Lushr(owner: MethodInfo) extends ShiftLOp(owner, OpCode.LUSHR, "lushr") {
}

class Land(owner: MethodInfo) extends BinaryLOp(owner, OpCode.LAND, "land") {
}

class Lor(owner: MethodInfo) extends BinaryLOp(owner, OpCode.LOR, "lor") {
}

class Lxor(owner: MethodInfo) extends BinaryLOp(owner, OpCode.LXOR, "lxor") {
}

//
// float operations
//

class Fadd(owner: MethodInfo) extends BinaryFOp(owner, OpCode.FADD, "fadd") {
}

class Fsub(owner: MethodInfo) extends BinaryFOp(owner, OpCode.FSUB, "fsub") {
}

class Fmul(owner: MethodInfo) extends BinaryFOp(owner, OpCode.FMUL, "fmul") {
}

class Fdiv(owner: MethodInfo) extends BinaryFOp(owner, OpCode.FDIV, "fdiv") {
}

class Frem(owner: MethodInfo) extends BinaryFOp(owner, OpCode.FREM, "frem") {
}

class Fneg(owner: MethodInfo) extends UnaryFOp(owner, OpCode.FNEG, "fneg") {
}

//
// double operations
//

class Dadd(owner: MethodInfo) extends BinaryDOp(owner, OpCode.DADD, "dadd") {
}

class Dsub(owner: MethodInfo) extends BinaryDOp(owner, OpCode.DSUB, "dsub") {
}

class Dmul(owner: MethodInfo) extends BinaryDOp(owner, OpCode.DMUL, "dmul") {
}

class Ddiv(owner: MethodInfo) extends BinaryDOp(owner, OpCode.DDIV, "ddiv") {
}

class Drem(owner: MethodInfo) extends BinaryDOp(owner, OpCode.DREM, "drem") {
}

class Dneg(owner: MethodInfo) extends UnaryDOp(owner, OpCode.DNEG, "dneg") {
}

