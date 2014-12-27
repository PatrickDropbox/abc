import java.io.DataInputStream
import java.io.DataOutputStream
import java.util.Vector


abstract class Operation(o: MethodInfo) {
    var _owner: MethodInfo = o
    var line = -1
    var pc = -1

    def canonicalForm(): Operation = this

    def serialize(output: DataOutputStream)

    def deserialize(opCode: Int, input: DataInputStream)

    def debugString(): String
}

// operations of the form: <op code>
abstract class NoOperandOp(
        owner: MethodInfo,
        opCode: Int,
        mnemonic: String) extends Operation(owner) {
    val _opCode = opCode
    val _mnemonic = mnemonic

    def serialize(output: DataOutputStream) {
        output.writeByte(_opCode)
    }

    def deserialize(opCode: Int, input: DataInputStream) {
        if (opCode != _opCode) {
            throw new Exception("Unexpected op-code: " + opCode)
        }
    }

    def debugString(): String = _mnemonic
}

// operations of the form: <op code> <byte operand>
abstract class ByteOperandOp(
        owner: MethodInfo,
        opCode: Int,
        mnemonic: String,
        signed: Boolean,
        v: Int) extends Operation(owner) {
    val _opCode = opCode
    val _mnemonic = mnemonic
    var isSigned = signed
    var operand = v

    def serialize(output: DataOutputStream) {
        output.writeByte(_opCode)
        output.writeByte(operand)
    }

    def deserialize(opCode: Int, input: DataInputStream) {
        if (opCode != _opCode) {
            throw new Exception("Unexpected op-code: " + opCode)
        }

        if (isSigned) {
            operand = input.readByte()
        } else {
            operand = input.readUnsignedByte()
        }
    }

    def debugString(): String = _mnemonic + " " + operand
}

// operations of the form: <op code> <short operand>
class ShortOperandOp(
        owner: MethodInfo,
        opCode: Int,
        mnemonic: String,
        signed: Boolean,
        v: Int) extends Operation(owner) {
    val _opCode = opCode
    val _mnemonic = mnemonic
    var isSigned = signed
    var operand = v

    def serialize(output: DataOutputStream) {
        output.writeByte(_opCode)
        output.writeShort(operand)
    }

    def deserialize(opCode: Int, input: DataInputStream) {
        if (opCode != _opCode) {
            throw new Exception("Unexpected op-code: " + opCode)
        }

        if (isSigned) {
            operand = input.readShort()
        } else {
            operand = input.readUnsignedShort()
        }
    }

    def debugString(): String = _mnemonic + " " + operand
}

object Operation {
    def deserialize(owner: MethodInfo, input: DataInputStream): Operation = {
        val opCode = input.readUnsignedByte()
        var operation: Operation = opCode match {
            case 0 => new Nop(owner)
            case 1 => new AconstNull(owner)
            case 2 => new IconstM1(owner)
            case 3 => new Iconst0(owner)
            case 4 => new Iconst1(owner)
            case 5 => new Iconst2(owner)
            case 6 => new Iconst3(owner)
            case 7 => new Iconst4(owner)
            case 8 => new Iconst5(owner)
            case 9 => new Lconst0(owner)
            case 10 => new Lconst1(owner)
            case 11 => new Fconst0(owner)
            case 12 => new Fconst1(owner)
            case 13 => new Fconst2(owner)
            case 14 => new Dconst0(owner)
            case 15 => new Dconst1(owner)
            case 16 => new Bipush(owner)
            case 17 => new Sipush(owner)
            case 18 => new Ldc(owner)
            case 19 => new LdcW(owner)
            case 20 => new Ldc2W(owner)
            case 21 => new LoadI(owner)
            case 22 => new LoadL(owner)
            case 23 => new LoadF(owner)
            case 24 => new LoadD(owner)
            case 25 => new LoadA(owner)
            case 26 => new Iload0(owner)
            case 27 => new Iload1(owner)
            case 28 => new Iload2(owner)
            case 29 => new Iload3(owner)
            case 30 => new Lload0(owner)
            case 31 => new Lload1(owner)
            case 32 => new Lload2(owner)
            case 33 => new Lload3(owner)
            case 34 => new Fload0(owner)
            case 35 => new Fload1(owner)
            case 36 => new Fload2(owner)
            case 37 => new Fload3(owner)
            case 38 => new Dload0(owner)
            case 39 => new Dload1(owner)
            case 40 => new Dload2(owner)
            case 41 => new Dload3(owner)
            case 42 => new Aload0(owner)
            case 43 => new Aload1(owner)
            case 44 => new Aload2(owner)
            case 45 => new Aload3(owner)
            case 46 => new LoadFromIArray(owner)
            case 47 => new LoadFromLArray(owner)
            case 48 => new LoadFromFArray(owner)
            case 49 => new LoadFromDArray(owner)
            case 50 => new LoadFromAArray(owner)
            case 51 => new LoadFromBArray(owner)
            case 52 => new LoadFromCArray(owner)
            case 53 => new LoadFromSArray(owner)
            case 54 => throw new Exception("TODO")
            case 55 => throw new Exception("TODO")
            case 56 => throw new Exception("TODO")
            case 57 => throw new Exception("TODO")
            case 58 => throw new Exception("TODO")
            case 59 => throw new Exception("TODO")
            case 60 => throw new Exception("TODO")
            case 61 => throw new Exception("TODO")
            case 62 => throw new Exception("TODO")
            case 63 => throw new Exception("TODO")
            case 64 => throw new Exception("TODO")
            case 65 => throw new Exception("TODO")
            case 66 => throw new Exception("TODO")
            case 67 => throw new Exception("TODO")
            case 68 => throw new Exception("TODO")
            case 69 => throw new Exception("TODO")
            case 70 => throw new Exception("TODO")
            case 71 => throw new Exception("TODO")
            case 72 => throw new Exception("TODO")
            case 73 => throw new Exception("TODO")
            case 74 => throw new Exception("TODO")
            case 75 => throw new Exception("TODO")
            case 76 => throw new Exception("TODO")
            case 77 => throw new Exception("TODO")
            case 78 => throw new Exception("TODO")
            case 79 => throw new Exception("TODO")
            case 80 => throw new Exception("TODO")
            case 81 => throw new Exception("TODO")
            case 82 => throw new Exception("TODO")
            case 83 => throw new Exception("TODO")
            case 84 => throw new Exception("TODO")
            case 85 => throw new Exception("TODO")
            case 86 => throw new Exception("TODO")
            case 87 => throw new Exception("TODO")
            case 88 => throw new Exception("TODO")
            case 89 => throw new Exception("TODO")
            case 90 => throw new Exception("TODO")
            case 91 => throw new Exception("TODO")
            case 92 => throw new Exception("TODO")
            case 93 => throw new Exception("TODO")
            case 94 => throw new Exception("TODO")
            case 95 => throw new Exception("TODO")
            case 96 => throw new Exception("TODO")
            case 97 => throw new Exception("TODO")
            case 98 => throw new Exception("TODO")
            case 99 => throw new Exception("TODO")
            case 100 => throw new Exception("TODO")
            case 101 => throw new Exception("TODO")
            case 102 => throw new Exception("TODO")
            case 103 => throw new Exception("TODO")
            case 104 => throw new Exception("TODO")
            case 105 => throw new Exception("TODO")
            case 106 => throw new Exception("TODO")
            case 107 => throw new Exception("TODO")
            case 108 => throw new Exception("TODO")
            case 109 => throw new Exception("TODO")
            case 110 => throw new Exception("TODO")
            case 111 => throw new Exception("TODO")
            case 112 => throw new Exception("TODO")
            case 113 => throw new Exception("TODO")
            case 114 => throw new Exception("TODO")
            case 115 => throw new Exception("TODO")
            case 116 => throw new Exception("TODO")
            case 117 => throw new Exception("TODO")
            case 118 => throw new Exception("TODO")
            case 119 => throw new Exception("TODO")
            case 120 => throw new Exception("TODO")
            case 121 => throw new Exception("TODO")
            case 122 => throw new Exception("TODO")
            case 123 => throw new Exception("TODO")
            case 124 => throw new Exception("TODO")
            case 125 => throw new Exception("TODO")
            case 126 => throw new Exception("TODO")
            case 127 => throw new Exception("TODO")
            case 128 => throw new Exception("TODO")
            case 129 => throw new Exception("TODO")
            case 130 => throw new Exception("TODO")
            case 131 => throw new Exception("TODO")
            case 132 => throw new Exception("TODO")
            case 133 => throw new Exception("TODO")
            case 134 => throw new Exception("TODO")
            case 135 => throw new Exception("TODO")
            case 136 => throw new Exception("TODO")
            case 137 => throw new Exception("TODO")
            case 138 => throw new Exception("TODO")
            case 139 => throw new Exception("TODO")
            case 140 => throw new Exception("TODO")
            case 141 => throw new Exception("TODO")
            case 142 => throw new Exception("TODO")
            case 143 => throw new Exception("TODO")
            case 144 => throw new Exception("TODO")
            case 145 => throw new Exception("TODO")
            case 146 => throw new Exception("TODO")
            case 147 => throw new Exception("TODO")
            case 148 => throw new Exception("TODO")
            case 149 => throw new Exception("TODO")
            case 150 => throw new Exception("TODO")
            case 151 => throw new Exception("TODO")
            case 152 => throw new Exception("TODO")
            case 153 => throw new Exception("TODO")
            case 154 => throw new Exception("TODO")
            case 155 => throw new Exception("TODO")
            case 156 => throw new Exception("TODO")
            case 157 => throw new Exception("TODO")
            case 158 => throw new Exception("TODO")
            case 159 => throw new Exception("TODO")
            case 160 => throw new Exception("TODO")
            case 161 => throw new Exception("TODO")
            case 162 => throw new Exception("TODO")
            case 163 => throw new Exception("TODO")
            case 164 => throw new Exception("TODO")
            case 165 => throw new Exception("TODO")
            case 166 => throw new Exception("TODO")
            case 167 => throw new Exception("TODO")
            case 168 => throw new Exception("TODO")
            case 169 => throw new Exception("TODO")
            case 170 => throw new Exception("TODO")
            case 171 => throw new Exception("TODO")
            case 172 => throw new Exception("TODO")
            case 173 => throw new Exception("TODO")
            case 174 => throw new Exception("TODO")
            case 175 => throw new Exception("TODO")
            case 176 => throw new Exception("TODO")
            case 177 => throw new Exception("TODO")
            case 178 => throw new Exception("TODO")
            case 179 => throw new Exception("TODO")
            case 180 => throw new Exception("TODO")
            case 181 => throw new Exception("TODO")
            case 182 => throw new Exception("TODO")
            case 183 => throw new Exception("TODO")
            case 184 => throw new Exception("TODO")
            case 185 => throw new Exception("TODO")
            case 186 => throw new Exception("TODO")
            case 187 => throw new Exception("TODO")
            case 188 => throw new Exception("TODO")
            case 189 => throw new Exception("TODO")
            case 190 => throw new Exception("TODO")
            case 191 => throw new Exception("TODO")
            case 192 => throw new Exception("TODO")
            case 193 => throw new Exception("TODO")
            case 194 => throw new Exception("TODO")
            case 195 => throw new Exception("TODO")
            case 196 => throw new Exception("TODO")
            case 197 => throw new Exception("TODO")
            case 198 => throw new Exception("TODO")
            case 199 => throw new Exception("TODO")
            case 200 => throw new Exception("TODO")
            case 201 => throw new Exception("TODO")
            case _ => throw new Exception("Unknown op code" + opCode)
        }

        operation.deserialize(opCode, input)
        return operation.canonicalForm()
    }
}
