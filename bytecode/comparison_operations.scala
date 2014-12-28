import java.io.DataInputStream
import java.io.DataOutputStream


// stack: ..., long value1, long value2 -> int result
class Lcmp(owner: MethodInfo) extends NoOperandOp(owner, OpCode.LCMP, "lcmp") {
}

// stack: ..., float value1, float value2 -> int result
abstract class Fcmp(owner: MethodInfo, opCode: Int, mnemonic: String)
        extends NoOperandOp(owner, opCode, mnemonic) {
}

// stack: ..., double value1, double value2 -> int result
abstract class Dcmp(owner: MethodInfo, opCode: Int, mnemonic: String)
        extends NoOperandOp(owner, opCode, mnemonic) {
}

class Fcmpg(owner: MethodInfo) extends Fcmp(owner, OpCode.FCMPG, "fcmpg") {
}

class Fcmpl(owner: MethodInfo) extends Fcmp(owner, OpCode.FCMPL, "fcmpl") {
}

class Dcmpg(owner: MethodInfo) extends Dcmp(owner, OpCode.DCMPG, "dcmpg") {
}

class Dcmpl(owner: MethodInfo) extends Dcmp(owner, OpCode.DCMPL, "dcmpl") {
}

abstract class IfBaseOp(
        owner: MethodInfo,
        opCode: Int,
        mnemonic: String,
        trueCondPc: Int) extends ShortOperandOp(
                owner,
                opCode,
                mnemonic,
                false,  // signed
                trueCondPc) {
}

// stack: ..., int value -> ... (compare against 0)
abstract class IfIOp(
        owner: MethodInfo,
        opCode: Int,
        mnemonic: String,
        trueCondPc: Int) extends IfBaseOp(owner, opCode, mnemonic, trueCondPc) {
}

// stack: ..., int value1, int value2 -> ...
abstract class IfCmpIOp(
        owner: MethodInfo,
        opCode: Int,
        mnemonic: String,
        trueCondPc: Int) extends IfBaseOp(owner, opCode, mnemonic, trueCondPc) {
}

// stack: ..., ref value -> ...
abstract class IfAOp(
        owner: MethodInfo,
        opCode: Int,
        mnemonic: String,
        trueCondPc: Int) extends IfBaseOp(owner, opCode, mnemonic, trueCondPc) {
}

// stack: ..., ref value1, ref value2 -> ...
abstract class IfCmpAOp(
        owner: MethodInfo,
        opCode: Int,
        mnemonic: String,
        trueCondPc: Int) extends IfBaseOp(owner, opCode, mnemonic, trueCondPc) {
}

//
// int conditional branching
//

class Ifeq(owner: MethodInfo, trueCondPc: Int)
        extends IfIOp(owner, OpCode.IFEQ, "ifeq", trueCondPc) {
    def this(owner: MethodInfo) = this(owner, -1)
}

class Ifne(owner: MethodInfo, trueCondPc: Int)
        extends IfIOp(owner, OpCode.IFNE, "ifne", trueCondPc) {
    def this(owner: MethodInfo) = this(owner, -1)
}

class Iflt(owner: MethodInfo, trueCondPc: Int)
        extends IfIOp(owner, OpCode.IFLT, "iflt", trueCondPc) {
    def this(owner: MethodInfo) = this(owner, -1)
}

class Ifle(owner: MethodInfo, trueCondPc: Int)
        extends IfIOp(owner, OpCode.IFLE, "ifle", trueCondPc) {
    def this(owner: MethodInfo) = this(owner, -1)
}

class Ifgt(owner: MethodInfo, trueCondPc: Int)
        extends IfIOp(owner, OpCode.IFGT, "ifgt", trueCondPc) {
    def this(owner: MethodInfo) = this(owner, -1)
}

class Ifge(owner: MethodInfo, trueCondPc: Int)
        extends IfIOp(owner, OpCode.IFGE, "ifge", trueCondPc) {
    def this(owner: MethodInfo) = this(owner, -1)
}

class IfIcmpeq(owner: MethodInfo, trueCondPc: Int)
        extends IfCmpIOp(owner, OpCode.IF_ICMPEQ, "if_icmpeq", trueCondPc) {
    def this(owner: MethodInfo) = this(owner, -1)
}

class IfIcmpne(owner: MethodInfo, trueCondPc: Int)
        extends IfCmpIOp(owner, OpCode.IF_ICMPNE, "if_icmpne", trueCondPc) {
    def this(owner: MethodInfo) = this(owner, -1)
}

class IfIcmplt(owner: MethodInfo, trueCondPc: Int)
        extends IfCmpIOp(owner, OpCode.IF_ICMPLT, "if_icmplt", trueCondPc) {
    def this(owner: MethodInfo) = this(owner, -1)
}

class IfIcmpge(owner: MethodInfo, trueCondPc: Int)
        extends IfCmpIOp(owner, OpCode.IF_ICMPGE, "if_icmpge", trueCondPc) {
    def this(owner: MethodInfo) = this(owner, -1)
}

class IfIcmpgt(owner: MethodInfo, trueCondPc: Int)
        extends IfCmpIOp(owner, OpCode.IF_ICMPGT, "if_icmpgt", trueCondPc) {
    def this(owner: MethodInfo) = this(owner, -1)
}

class IfIcmple(owner: MethodInfo, trueCondPc: Int)
        extends IfCmpIOp(owner, OpCode.IF_ICMPLE, "if_icmple", trueCondPc) {
    def this(owner: MethodInfo) = this(owner, -1)
}

//
// ref conditional branching
//

class IfAcmpeq(owner: MethodInfo, trueCondPc: Int)
        extends IfCmpAOp(owner, OpCode.IF_ACMPEQ, "if_acmpeq", trueCondPc) {
    def this(owner: MethodInfo) = this(owner, -1)
}

class IfAcmpne(owner: MethodInfo, trueCondPc: Int)
        extends IfCmpAOp(owner, OpCode.IF_ACMPNE, "if_acmpne", trueCondPc) {
    def this(owner: MethodInfo) = this(owner, -1)
}

class Ifnonnull(owner: MethodInfo, trueCondPc: Int)
        extends IfAOp(owner, OpCode.IFNONNULL, "ifnonnull", trueCondPc) {
    def this(owner: MethodInfo) = this(owner, -1)
}

class Ifnull(owner: MethodInfo, trueCondPc: Int)
        extends IfAOp(owner, OpCode.IFNULL, "ifnull", trueCondPc) {
    def this(owner: MethodInfo) = this(owner, -1)
}
