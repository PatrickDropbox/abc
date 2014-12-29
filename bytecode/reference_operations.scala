import java.io.DataInputStream
import java.io.DataOutputStream


class ClassOp(
        owner: MethodInfo,
        opCode: Int,
        mnemonic: String,
        className: String) extends Operation(owner) {
    def this(owner: MethodInfo,
             opCode: Int,
             mnemonic: String) = this(owner, opCode, mnemonic, null)

    val _opCode = opCode
    val _mnemonic = mnemonic

    var _constClass: ConstClassInfo = null
    if (className != null) {
        _constClass = _owner.constants().getClass(className)
    }

    def serialize(output: DataOutputStream) {
        output.writeByte(_opCode)
        output.writeShort(_constClass.index)
    }

    def deserialize(
            startAddress: Int,
            opCode: Int,
            input: DataInputStream) {
        if (opCode != _opCode) {
            throw new Exception("Unexpected op-code: " + opCode)
        }

        _constClass = _owner.constants().getClassByIndex(
                input.readUnsignedShort())
    }

    def debugString(indent: String): String = {
        var name = "???"
        if (_constClass != null) {
            name = _constClass.className()
        }
        return indent + _mnemonic + " " + name
    }
}

class FieldOp(
        owner: MethodInfo,
        opCode: Int,
        mnemonic: String,
        className: String,
        fieldName: String,
        fieldType: FieldType) extends Operation(owner) {
    val _opCode = opCode
    val _mnemonic = mnemonic

    var _constFieldRef: ConstFieldRefInfo = null
    if (className != null) {
        _constFieldRef = _owner.constants().getFieldRef(
                className,
                fieldName,
                fieldType)
    }

    def serialize(output: DataOutputStream) {
        output.writeByte(_opCode)
        output.writeShort(_constFieldRef.index)
    }

    def deserialize(startAddress: Int, opCode: Int, input: DataInputStream) {
        if (opCode != _opCode) {
            throw new Exception("Unexpected op-code: " + opCode)
        }

        val index = input.readUnsignedShort()
        _constFieldRef = _owner.constants().getFieldRefByIndex(index)
    }

    def debugString(indent: String): String = {
        var name = "???"
        if (_constFieldRef != null) {
            name = _constFieldRef.debugString()
        }
        return indent + _mnemonic + " " + name
    }
}

class MethodOp(
        owner: MethodInfo,
        opCode: Int,
        mnemonic: String,
        className: String,
        methodName: String,
        methodType: MethodType,
        useInterfaceMethodRef: Boolean) extends Operation(owner) {
    val _opCode = opCode
    val _mnemonic = mnemonic

    var _constMethodRef: ConstBaseMethodRefInfo = null
    if (className != null) {
        if (useInterfaceMethodRef) {
            _constMethodRef = _owner.constants().getInterfaceMethodRef(
                    className,
                    methodName,
                    methodType)
        } else {
            _constMethodRef = _owner.constants().getMethodRef(
                    className,
                    methodName,
                    methodType)
        }
    }

    def serialize(output: DataOutputStream) {
        output.writeByte(_opCode)
        output.writeShort(_constMethodRef.index)
    }

    def deserialize(startAddress: Int, opCode: Int, input: DataInputStream) {
        if (opCode != _opCode) {
            throw new Exception("Unexpected op-code: " + opCode)
        }

        val index = input.readUnsignedShort()
        _constMethodRef = _owner.constants().getBaseMethodRefByIndex(index)
    }

    def debugString(indent: String): String = {
        var name = "???"
        if (_constMethodRef != null) {
            name = _constMethodRef.debugString()
        }
        return indent + _mnemonic + " " + name
    }
}

// stack: ... -> ..., value
class Getstatic(
        owner: MethodInfo,
        className: String,
        fieldName: String,
        fieldType: FieldType) extends FieldOp(
                owner,
                OpCode.GETSTATIC,
                "getstatic",
                className,
                fieldName,
                fieldType) {
    def this(owner: MethodInfo) = this(owner, null, null, null)
}

// stack: ..., value -> ...
class Putstatic(
        owner: MethodInfo,
        className: String,
        fieldName: String,
        fieldType: FieldType) extends FieldOp(
                owner,
                OpCode.PUTSTATIC,
                "putstatic",
                className,
                fieldName,
                fieldType) {
    def this(owner: MethodInfo) = this(owner, null, null, null)
}

// stack: ..., obj ref -> ..., value
class Getfield(
        owner: MethodInfo,
        className: String,
        fieldName: String,
        fieldType: FieldType) extends FieldOp(
                owner,
                OpCode.GETFIELD,
                "getfield",
                className,
                fieldName,
                fieldType) {
    def this(owner: MethodInfo) = this(owner, null, null, null)
}

// stack: ..., obj ref, value -> ...
class Putfield(
        owner: MethodInfo,
        className: String,
        fieldName: String,
        fieldType: FieldType) extends FieldOp(
                owner,
                OpCode.PUTFIELD,
                "putfield",
                className,
                fieldName,
                fieldType) {
    def this(owner: MethodInfo) = this(owner, null, null, null)
}

// stack: ..., objref, [arg1, [arg2, ...]] -> ..., [result]
class Invokeinterface(
        owner: MethodInfo,
        className: String,
        methodName: String,
        methodType: MethodType) extends MethodOp(
                owner,
                OpCode.INVOKEINTERFACE,
                "invokeinterface",
                className,
                methodName,
                methodType,
                true) {  // use interface method ref
    def this(owner: MethodInfo) = this(owner, null, null, null)

    override def serialize(output: DataOutputStream) {
        super.serialize(output)
        val methodType = _constMethodRef.methodDescriptor()
        output.writeByte(1 + methodType.parameters.size())
        output.writeByte(0)
    }

    override def deserialize(pc: Int, opCode: Int, input: DataInputStream) {
        super.deserialize(pc, opCode, input)
        input.readUnsignedShort()  // ignore padding
    }
}

// stack: ..., objref, [arg1, [arg2, ...]] -> ..., [result]
class Invokespecial(
        owner: MethodInfo,
        className: String,
        methodName: String,
        methodType: MethodType) extends MethodOp(
                owner,
                OpCode.INVOKESPECIAL,
                "invokespecial",
                className,
                methodName,
                methodType,
                false) {  // use interface method ref
    def this(owner: MethodInfo) = this(owner, null, null, null)
}

// stack: ..., [arg1, [arg2, ...]] -> ..., [result]
class Invokestatic(
        owner: MethodInfo,
        className: String,
        methodName: String,
        methodType: MethodType) extends MethodOp(
                owner,
                OpCode.INVOKESTATIC,
                "invokestatic",
                className,
                methodName,
                methodType,
                false) {  // use interface method ref
    def this(owner: MethodInfo) = this(owner, null, null, null)
}

// stack: ..., objref, [arg1, [arg2, ...]] -> ..., [result]
class Invokevirtual(
        owner: MethodInfo,
        className: String,
        methodName: String,
        methodType: MethodType) extends MethodOp(
                owner,
                OpCode.INVOKEVIRTUAL,
                "invokevirtual",
                className,
                methodName,
                methodType,
                false) {  // use interface method ref
    def this(owner: MethodInfo) = this(owner, null, null, null)
}

// stack: ... -> ..., obj ref
class New(owner: MethodInfo, className: String)
        extends ClassOp(owner, OpCode.NEW, "new", className) {
    def this(owner: MethodInfo) = this(owner, null)
}

// stack: ... obj ref -> ..., obj ref
class Checkcast(owner: MethodInfo, className: String)
        extends ClassOp(owner, OpCode.CHECKCAST, "checkcast", className) {
    def this(owner: MethodInfo) = this(owner, null)
}

// stack: ... obj ref -> ..., int result
class Instanceof(owner: MethodInfo, className: String)
        extends ClassOp(owner, OpCode.INSTANCEOF, "instanceof", className) {
    def this(owner: MethodInfo) = this(owner, null)
}

// stack: ..., count -> ..., array ref
class Newarray(owner: MethodInfo, arrayType: BaseType)
        extends Operation(owner) {
    def this(owner: MethodInfo) = this(owner, null)

    var _arrayType = arrayType

    def serialize(output: DataOutputStream) {
        output.writeByte(OpCode.NEWARRAY)
        output.writeByte(_arrayType.arrayType())
    }

    def deserialize(
            startAddress: Int,
            opCode: Int,
            input: DataInputStream) {
        if (opCode != OpCode.NEWARRAY) {
            throw new Exception("Unexpected op-code: " + opCode)
        }

        val v = input.readUnsignedByte()
        _arrayType = v match {
            case 4 => new BoolType()
            case 5 => new CharType()
            case 6 => new FloatType()
            case 7 => new DoubleType()
            case 8 => new ByteType()
            case 9 => new ShortType()
            case 10 => new IntType()
            case 11 => new LongType()
            case _ => throw new Exception("Unknown base array type " + v)
        }
    }

    def debugString(indent: String): String = {
        return indent + "newarray " + _arrayType.descriptorString()
    }
}

// stack: ..., count -> ..., array ref
// NOTE: for array type, use the descriptor string as class name
class Anewarray(owner: MethodInfo, className: String)
        extends ClassOp(owner, OpCode.ANEWARRAY, "anewarray", className) {
    def this(owner: MethodInfo) = this(owner, null)
}

// stack: ..., count1, [count2, ...] -> ..., array ref
class Multianewarray(
        owner: MethodInfo,
        arrayDescriptorString: String,
        dimensions: Int) extends ClassOp(
                owner,
                OpCode.MULTIANEWARRAY,
                "multianewarray",
                arrayDescriptorString) {
    def this(owner: MethodInfo) = this (owner, null, 0)

    var _dimensions = dimensions

    override def serialize(output: DataOutputStream) {
        super.serialize(output)
        output.writeByte(_dimensions)
    }

    override def deserialize(pc: Int, opCode: Int, input: DataInputStream) {
        super.deserialize(pc, opCode, input)
        _dimensions = input.readUnsignedByte()
    }

    override def debugString(indent: String): String = {
        return super.debugString(indent) + " " + _dimensions
    }
}

// stack: ..., array ref -> ...
class Arraylength(owner: MethodInfo)
        extends NoOperandOp(owner, OpCode.ARRAYLENGTH, "arraylength") {
}

// stack: ..., obj ref -> ...
class Athrow(owner: MethodInfo)
        extends NoOperandOp(owner, OpCode.ATHROW, "athrow") {
}

// stack: ..., obj ref -> ...
class Monitorenter(owner: MethodInfo)
        extends NoOperandOp(owner, OpCode.MONITORENTER, "monitorenter") {
}

// stack: ..., obj ref -> ...
class Monitorexit(owner: MethodInfo)
        extends NoOperandOp(owner, OpCode.MONITOREXIT, "monitorexit") {
}

