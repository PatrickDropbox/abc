import java.io.DataInputStream
import java.io.DataOutputStream


//
// nop
// stack: ... -> ...
//
class Nop(owner: MethodInfo)
        extends NoOperandOp(owner, OpCode.NOP, "nop") {
}

//
// aconst_null
// stack: ... -> ..., null
//
class AconstNull(owner: MethodInfo)
        extends NoOperandOp(owner, OpCode.ACONST_NULL, "aconst_null") {
}

//
// "PushI" <value>
// stack: ... -> ..., value
//
class PushI(owner: MethodInfo, v: Int) extends Operation(owner) {
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

    def deserialize(startAddress: Int, opCode: Int, input: DataInputStream) {
        throw new Exception("cannot directly deserialize \"iconst\"")
    }

    def debugString(): String = "\"iconst\" " + value
}

//
// "PushL" <value>
// stack: ... -> ..., value
//
class PushL(owner: MethodInfo, v: Long) extends Operation(owner) {
    var value = v
    var _constLong: ConstLongInfo = null
    if (value != 0 && value != 1) {
        _constLong = owner.constants().getLong(value)
    }

    def serialize(output: DataOutputStream) {
        if (value == 0) {
            output.write(OpCode.LCONST_0)
        } else if (value == 1) {
            output.write(OpCode.LCONST_1)
        } else {
            output.write(OpCode.LDC2_W)
            output.writeShort(_constLong.index)
        }
    }

    def deserialize(startAddress: Int, opCode: Int, input: DataInputStream) {
        throw new Exception("cannot directly deserialize \"lconst\"")
    }

    def debugString(): String = "\"lconst\" " + value
}

//
// "PushF" <value>
// stack: ... -> ..., value
//
class PushF(owner: MethodInfo, v: Float) extends Operation(owner) {
    var value = v
    var _constFloat: ConstFloatInfo = null
    if (value != 0 && value != 1 && value != 2) {
        _constFloat = owner.constants().getFloat(value)
    }

    def serialize(output: DataOutputStream) {
        if (value == 0) {
            output.write(OpCode.FCONST_0)
        } else if (value == 1) {
            output.write(OpCode.FCONST_1)
        } else if (value == 2) {
            output.write(OpCode.FCONST_2)
        } else {
            if (_constFloat.index < 256) {
                output.write(OpCode.LDC)
                output.writeByte(_constFloat.index)
            } else {
                output.write(OpCode.LDC_W)
                output.writeShort(_constFloat.index)
            }
        }
    }

    def deserialize(startAddress: Int, opCode: Int, input: DataInputStream) {
        throw new Exception("cannot directly deserialize \"fconst\"")
    }

    def debugString(): String = "\"fconst\" " + value
}

//
// "PushD" <value>
// stack: ... -> ..., value
//
class PushD(owner: MethodInfo, v: Double) extends Operation(owner) {
    var value = v
    var _constDouble: ConstDoubleInfo = null
    if (value != 0 && value != 1) {
        _constDouble = owner.constants().getDouble(value)
    }

    def serialize(output: DataOutputStream) {
        if (value == 0) {
            output.write(OpCode.DCONST_0)
        } else if (value == 1) {
            output.write(OpCode.DCONST_1)
        } else {
            output.write(OpCode.LDC2_W)
            output.writeShort(_constDouble.index)
        }
    }

    def deserialize(startAddress: Int, opCode: Int, input: DataInputStream) {
        throw new Exception("cannot directly deserialize \"dconst\"")
    }

    def debugString(): String = "\"dconst\" " + value
}

//
// "PushString" <value>
// stack: ... -> ..., value
//
class PushString(owner: MethodInfo, v: String) extends Operation(owner) {
    var _constString: ConstStringInfo = null
    if (value != 0 && value != 1) {
        _constString = owner.constants().getString(value)
    }

    def value(): String = _constString.value()

    def serialize(output: DataOutputStream) {
        if (_constString.index < 256) {
            output.write(OpCode.LDC)
            output.writeByte(_constString.index)
        } else {
            output.write(OpCode.LDC_W)
            output.writeShort(_constString.index)
        }
    }

    def deserialize(startAddress: Int, opCode: Int, input: DataInputStream) {
        throw new Exception("cannot directly deserialize \"sconst\"")
    }

    def debugString(): String = "\"sconst\" " + value
}

class IconstM1(owner: MethodInfo)
        extends NoOperandOp(owner, OpCode.ICONST_M1, "iconst_m1") {
    override def canonicalForm(): Operation = {
        return new PushI(_owner, -1)
    }
}

class Iconst0(owner: MethodInfo)
        extends NoOperandOp(owner, OpCode.ICONST_0, "iconst_0") {
    override def canonicalForm(): Operation = {
        return new PushI(_owner, 0)
    }
}

class Iconst1(owner: MethodInfo)
        extends NoOperandOp(owner, OpCode.ICONST_1, "iconst_1") {
    override def canonicalForm(): Operation = {
        return new PushI(_owner, 1)
    }
}

class Iconst2(owner: MethodInfo)
        extends NoOperandOp(owner, OpCode.ICONST_2, "iconst_2") {
    override def canonicalForm(): Operation = {
        return new PushI(_owner, 2)
    }
}

class Iconst3(owner: MethodInfo)
        extends NoOperandOp(owner, OpCode.ICONST_3, "iconst_3") {
    override def canonicalForm(): Operation = {
        return new PushI(_owner, 3)
    }
}

class Iconst4(owner: MethodInfo)
        extends NoOperandOp(owner, OpCode.ICONST_4, "iconst_4") {
    override def canonicalForm(): Operation = {
        return new PushI(_owner, 4)
    }
}

class Iconst5(owner: MethodInfo)
        extends NoOperandOp(owner, OpCode.ICONST_5, "iconst_5") {
    override def canonicalForm(): Operation = {
        return new PushI(_owner, 5)
    }
}

class Lconst0(owner: MethodInfo)
        extends NoOperandOp(owner, OpCode.LCONST_0, "lconst_0") {
    override def canonicalForm(): Operation = {
        return new PushL(_owner, 0)
    }
}

class Lconst1(owner: MethodInfo)
        extends NoOperandOp(owner, OpCode.LCONST_1, "lconst_1") {
    override def canonicalForm(): Operation = {
        return new PushL(_owner, 1)
    }
}

class Fconst0(owner: MethodInfo)
        extends NoOperandOp(owner, OpCode.FCONST_0, "fconst_0") {
    override def canonicalForm(): Operation = {
        return new PushF(_owner, 0)
    }
}

class Fconst1(owner: MethodInfo)
        extends NoOperandOp(owner, OpCode.FCONST_1, "fconst_1") {
    override def canonicalForm(): Operation = {
        return new PushF(_owner, 1)
    }
}

class Fconst2(owner: MethodInfo)
        extends NoOperandOp(owner, OpCode.FCONST_2, "fconst_2") {

    override def canonicalForm(): Operation = {
        return new PushF(_owner, 2)
    }
}

class Dconst0(owner: MethodInfo)
        extends NoOperandOp(owner, OpCode.DCONST_0, "dconst_0") {
    override def canonicalForm(): Operation = {
        return new PushD(_owner, 0)
    }
}

class Dconst1(owner: MethodInfo)
        extends NoOperandOp(owner, OpCode.DCONST_1, "dconst_1") {

    override def canonicalForm(): Operation = {
        return new PushD(_owner, 1)
    }
}

class Bipush(owner: MethodInfo, v: Int) extends ByteOperandOp(
        owner,
        OpCode.BIPUSH,
        "bipush",
        true,  // is signed
        v) {
    def this(owner: MethodInfo) = this(owner, 0)

    override def canonicalForm(): Operation = {
        return new PushI(_owner, operand)
    }
}

class Sipush(owner: MethodInfo, v: Int) extends ShortOperandOp(
        owner,
        OpCode.SIPUSH,
        "sipush",
        true,  // is signed
        v) {
    def this(owner: MethodInfo) = this(owner, 0)

    override def canonicalForm(): Operation = {
        return new PushI(_owner, operand)
    }
}

class Ldc(owner: MethodInfo, v: ConstInfo) extends ByteOperandOp(
        owner,
        OpCode.LDC,
        "ldc",
        false,  // signed
        0) {
    def this(owner: MethodInfo) = this(owner, null)

    var _const = v

    override def serialize(output: DataOutputStream) {
        operand = _const.index
        super.serialize(output)
    }

    override def deserialize(
            startAddress: Int,
            opCode: Int,
            input: DataInputStream) {
        super.deserialize(startAddress, opCode, input)
        _const = _owner.constants().getByIndex(operand)
    }

    override def canonicalForm(): Operation = {
        return _const match {
            case c: ConstIntegerInfo => new PushI(_owner, c.value())
            case c: ConstFloatInfo => new PushF(_owner, c.value())
            case c: ConstStringInfo => new PushString(_owner, c.value())
            // TODO: support class / method type / method handle
            case _ => this
        }
    }
}

class LdcW(owner: MethodInfo, v: ConstInfo) extends ShortOperandOp(
        owner,
        OpCode.LDC_W,
        "ldc_w",
        false,  // signed
        0) {
    def this(owner: MethodInfo) = this(owner, null)

    var _const = v

    override def serialize(output: DataOutputStream) {
        operand = _const.index
        super.serialize(output)
    }

    override def deserialize(
            startAddress: Int,
            opCode: Int,
            input: DataInputStream) {
        super.deserialize(startAddress, opCode, input)
        _const = _owner.constants().getByIndex(operand)
    }

    override def canonicalForm(): Operation = {
        return _const match {
            case c: ConstIntegerInfo => new PushI(_owner, c.value())
            case c: ConstFloatInfo => new PushF(_owner, c.value())
            case c: ConstStringInfo => new PushString(_owner, c.value())
            // TODO: support class / method type / method handle
            case _ => this
        }
    }
}

class Ldc2W(owner: MethodInfo, v: ConstInfo) extends ShortOperandOp(
        owner,
        OpCode.LDC2_W,
        "ldc2_w",
        false,  // signed
        0) {
    def this(owner: MethodInfo) = this(owner, null)

    var _const = v

    override def serialize(output: DataOutputStream) {
        operand = _const.index
        super.serialize(output)
    }

    override def deserialize(
            startAddress: Int,
            opCode: Int,
            input: DataInputStream) {
        super.deserialize(startAddress, opCode, input)
        _const = _owner.constants().getByIndex(operand)
    }

    override def canonicalForm(): Operation = {
        return _const match {
            case c: ConstDoubleInfo => new PushD(_owner, c.value())
            case c: ConstLongInfo => new PushL(_owner, c.value())
            case _ => throw new Exception(
                    "Unexpected constant type: " + _const.typeName())
        }
    }
}
