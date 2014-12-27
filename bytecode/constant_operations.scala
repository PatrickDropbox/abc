import java.io.DataInputStream
import java.io.DataOutputStream


class Nop(owner: MethodInfo) extends NoOperandOp(owner, OpCode.NOP, "nop") {
}

class AconstNull(owner: MethodInfo) extends NoOperandOp(
        owner,
        OpCode.ACONST_NULL,
        "aconst_null") {
}

class Iconst(owner: MethodInfo, v: Int) extends Operation(owner) {
    var value = v
    var _constInt: ConstIntegerInfo = null
    if (value < -32768 || value > 32767) {
        _constInt = owner.constants().getInteger(value)
    }

    def serialize(output: DataOutputStream) {
        value match {
            case -1 => output.write(OpCode.ICONST_M1)
            case 0 => output.write(OpCode.ICONST_0)
            case 1 => output.write(OpCode.ICONST_1)
            case 2 => output.write(OpCode.ICONST_2)
            case 3 => output.write(OpCode.ICONST_3)
            case 4 => output.write(OpCode.ICONST_4)
            case 5 => output.write(OpCode.ICONST_5)
            case _ => {
                if (-128 <= value && value <= 127) {
                    output.write(OpCode.BIPUSH)
                    output.writeByte(value)
                } else if (-32768 <= value && value <= 32767) {
                    output.write(OpCode.SIPUSH)
                    output.writeShort(value)
                } else {
                    if (_constInt.index < 256) {
                        output.write(OpCode.LDC)
                        output.writeShort(_constInt.index)
                    } else {
                        output.write(OpCode.LDC_W)
                        output.writeShort(_constInt.index)
                    }
                }
            }
        }
    }

    def deserialize(opCode: Int, input: DataInputStream) {
        throw new Exception("cannot directly deserialize iconst")
    }

    def debugString(): String = "iconst " + value
}

class IconstM1(owner: MethodInfo) extends NoOperandOp(
        owner,
        OpCode.ICONST_M1,
        "iconst_m1") {

    override def canonicalForm(): Operation = {
        return new Iconst(_owner, -1)
    }
}

class Iconst0(owner: MethodInfo) extends NoOperandOp(
        owner,
        OpCode.ICONST_0,
        "iconst_0") {

    override def canonicalForm(): Operation = {
        return new Iconst(_owner, 0)
    }
}

class Iconst1(owner: MethodInfo) extends NoOperandOp(
        owner,
        OpCode.ICONST_1,
        "iconst_1") {

    override def canonicalForm(): Operation = {
        return new Iconst(_owner, 1)
    }
}

class Iconst2(owner: MethodInfo) extends NoOperandOp(
        owner,
        OpCode.ICONST_2,
        "iconst_2") {

    override def canonicalForm(): Operation = {
        return new Iconst(_owner, 2)
    }
}

class Iconst3(owner: MethodInfo) extends NoOperandOp(
        owner,
        OpCode.ICONST_3,
        "iconst_3") {

    override def canonicalForm(): Operation = {
        return new Iconst(_owner, 3)
    }
}

class Iconst4(owner: MethodInfo) extends NoOperandOp(
        owner,
        OpCode.ICONST_4,
        "iconst_4") {

    override def canonicalForm(): Operation = {
        return new Iconst(_owner, 4)
    }
}

class Iconst5(owner: MethodInfo) extends NoOperandOp(
        owner,
        OpCode.ICONST_5,
        "iconst_5") {

    override def canonicalForm(): Operation = {
        return new Iconst(_owner, 5)
    }
}
