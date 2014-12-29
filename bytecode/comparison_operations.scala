import java.io.DataInputStream
import java.io.DataOutputStream


// stack: ..., long value1, long value2 -> int result
class Lcmp(owner: AttributeOwner)
        extends NoOperandOp(owner, OpCode.LCMP, "lcmp") {
}

// stack: ..., float value1, float value2 -> int result
abstract class Fcmp(owner: AttributeOwner, opCode: Int, mnemonic: String)
        extends NoOperandOp(owner, opCode, mnemonic) {
}

// stack: ..., double value1, double value2 -> int result
abstract class Dcmp(owner: AttributeOwner, opCode: Int, mnemonic: String)
        extends NoOperandOp(owner, opCode, mnemonic) {
}

class Fcmpg(owner: AttributeOwner) extends Fcmp(owner, OpCode.FCMPG, "fcmpg") {
}

class Fcmpl(owner: AttributeOwner) extends Fcmp(owner, OpCode.FCMPL, "fcmpl") {
}

class Dcmpg(owner: AttributeOwner) extends Dcmp(owner, OpCode.DCMPG, "dcmpg") {
}

class Dcmpl(owner: AttributeOwner) extends Dcmp(owner, OpCode.DCMPL, "dcmpl") {
}

abstract class IfBaseOp(
        owner: AttributeOwner,
        opCode: Int,
        mnemonic: String,
        truePcOffset: Int) extends ShortOperandOp(
                owner,
                opCode,
                mnemonic,
                false,  // signed
                truePcOffset) {
}

// stack: ..., int value -> ... (compare against 0)
abstract class IfIOp(
        owner: AttributeOwner,
        opCode: Int,
        mnemonic: String,
        truePcOffset: Int) extends IfBaseOp(
                owner,
                opCode,
                mnemonic,
                truePcOffset) {
}

// stack: ..., int value1, int value2 -> ...
abstract class IfCmpIOp(
        owner: AttributeOwner,
        opCode: Int,
        mnemonic: String,
        truePcOffset: Int) extends IfBaseOp(
                owner,
                opCode,
                mnemonic,
                truePcOffset) {
}

// stack: ..., ref value -> ...
abstract class IfAOp(
        owner: AttributeOwner,
        opCode: Int,
        mnemonic: String,
        truePcOffset: Int) extends IfBaseOp(
                owner,
                opCode,
                mnemonic,
                truePcOffset) {
}

// stack: ..., ref value1, ref value2 -> ...
abstract class IfCmpAOp(
        owner: AttributeOwner,
        opCode: Int,
        mnemonic: String,
        truePcOffset: Int) extends IfBaseOp(owner, opCode, mnemonic, truePcOffset) {
}

//
// int conditional branching
//

class Ifeq(owner: AttributeOwner, truePcOffset: Int)
        extends IfIOp(owner, OpCode.IFEQ, "ifeq", truePcOffset) {
    def this(owner: AttributeOwner) = this(owner, -1)
}

class Ifne(owner: AttributeOwner, truePcOffset: Int)
        extends IfIOp(owner, OpCode.IFNE, "ifne", truePcOffset) {
    def this(owner: AttributeOwner) = this(owner, -1)
}

class Iflt(owner: AttributeOwner, truePcOffset: Int)
        extends IfIOp(owner, OpCode.IFLT, "iflt", truePcOffset) {
    def this(owner: AttributeOwner) = this(owner, -1)
}

class Ifle(owner: AttributeOwner, truePcOffset: Int)
        extends IfIOp(owner, OpCode.IFLE, "ifle", truePcOffset) {
    def this(owner: AttributeOwner) = this(owner, -1)
}

class Ifgt(owner: AttributeOwner, truePcOffset: Int)
        extends IfIOp(owner, OpCode.IFGT, "ifgt", truePcOffset) {
    def this(owner: AttributeOwner) = this(owner, -1)
}

class Ifge(owner: AttributeOwner, truePcOffset: Int)
        extends IfIOp(owner, OpCode.IFGE, "ifge", truePcOffset) {
    def this(owner: AttributeOwner) = this(owner, -1)
}

class IfIcmpeq(owner: AttributeOwner, truePcOffset: Int)
        extends IfCmpIOp(owner, OpCode.IF_ICMPEQ, "if_icmpeq", truePcOffset) {
    def this(owner: AttributeOwner) = this(owner, -1)
}

class IfIcmpne(owner: AttributeOwner, truePcOffset: Int)
        extends IfCmpIOp(owner, OpCode.IF_ICMPNE, "if_icmpne", truePcOffset) {
    def this(owner: AttributeOwner) = this(owner, -1)
}

class IfIcmplt(owner: AttributeOwner, truePcOffset: Int)
        extends IfCmpIOp(owner, OpCode.IF_ICMPLT, "if_icmplt", truePcOffset) {
    def this(owner: AttributeOwner) = this(owner, -1)
}

class IfIcmpge(owner: AttributeOwner, truePcOffset: Int)
        extends IfCmpIOp(owner, OpCode.IF_ICMPGE, "if_icmpge", truePcOffset) {
    def this(owner: AttributeOwner) = this(owner, -1)
}

class IfIcmpgt(owner: AttributeOwner, truePcOffset: Int)
        extends IfCmpIOp(owner, OpCode.IF_ICMPGT, "if_icmpgt", truePcOffset) {
    def this(owner: AttributeOwner) = this(owner, -1)
}

class IfIcmple(owner: AttributeOwner, truePcOffset: Int)
        extends IfCmpIOp(owner, OpCode.IF_ICMPLE, "if_icmple", truePcOffset) {
    def this(owner: AttributeOwner) = this(owner, -1)
}

//
// ref conditional branching
//

class IfAcmpeq(owner: AttributeOwner, truePcOffset: Int)
        extends IfCmpAOp(owner, OpCode.IF_ACMPEQ, "if_acmpeq", truePcOffset) {
    def this(owner: AttributeOwner) = this(owner, -1)
}

class IfAcmpne(owner: AttributeOwner, truePcOffset: Int)
        extends IfCmpAOp(owner, OpCode.IF_ACMPNE, "if_acmpne", truePcOffset) {
    def this(owner: AttributeOwner) = this(owner, -1)
}

class Ifnonnull(owner: AttributeOwner, truePcOffset: Int)
        extends IfAOp(owner, OpCode.IFNONNULL, "ifnonnull", truePcOffset) {
    def this(owner: AttributeOwner) = this(owner, -1)
}

class Ifnull(owner: AttributeOwner, truePcOffset: Int)
        extends IfAOp(owner, OpCode.IFNULL, "ifnull", truePcOffset) {
    def this(owner: AttributeOwner) = this(owner, -1)
}
