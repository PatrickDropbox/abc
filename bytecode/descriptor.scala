import java.util.Vector

import scala.collection.JavaConversions._


trait DescriptorType extends Comparable[DescriptorType] {
    def descriptorString(): String

    def compareTo(other: DescriptorType): Int = {
        return descriptorString().compareTo(other.descriptorString())
    }
}

trait FieldType extends DescriptorType {
    // see table 2.3 / page 29
    def categorySize(): Int = 1

    def cloneType(): FieldType
}

abstract class BaseType(c: Boolean) extends FieldType {
    var isConstant = c

    def arrayType(): Int
}

trait RefType extends FieldType {
}

abstract class BaseIntType(c: Boolean, v: Int) extends BaseType(c) {
    var constValue = v
}

class BoolType extends BaseIntType(false, 0) {
    def descriptorString(): String = "Z"

    def arrayType(): Int = 4

    def cloneType(): FieldType = { return new BoolType() }
}

class ByteType extends BaseIntType(false, 0) {
    def descriptorString(): String = "B"

    def arrayType(): Int = 8

    def cloneType(): FieldType = { return new ByteType() }
}

class CharType extends BaseIntType(false, 0) {
    def descriptorString(): String = "C"

    def arrayType(): Int = 5

    def cloneType(): FieldType = { return new CharType() }
}

class ShortType extends BaseIntType(false, 0) {
    def descriptorString(): String = "S"

    def arrayType(): Int = 9

    def cloneType(): FieldType = { return new ShortType() }
}

class IntType(c: Boolean, v: Int) extends BaseIntType(c, v) {
    def this() = this(false, 0)
    def this(v: Int) = this(true, v)

    def descriptorString(): String = "I"

    def arrayType(): Int = 10

    def cloneType(): FieldType = { return new IntType(isConstant, constValue) }
}

class DoubleType(c: Boolean, v: Double) extends BaseType(c) {
    def this() = this(false, 0)
    def this(v: Double) = this(true, v)

    var constValue: Double = v

    def descriptorString(): String = "D"

    def arrayType(): Int = 7

    override def categorySize(): Int = 2

    def cloneType(): FieldType = {
        return new DoubleType(isConstant, constValue)
    }
}

class FloatType(c: Boolean, v: Float) extends BaseType(c) {
    def this() = this(false, 0)
    def this(v: Float) = this(true, v)

    var constValue: Float = v

    def descriptorString(): String = "F"

    def arrayType(): Int = 6

    def cloneType(): FieldType = {
        return new FloatType(isConstant, constValue)
    }
}

class LongType(c: Boolean, v: Long) extends BaseType(c) {
    def this() = this(false, 0)
    def this(v: Long) = this(true, v)

    var constValue: Long = v

    def descriptorString(): String = "J"

    def arrayType(): Int = 11

    override def categorySize(): Int = 2

    def cloneType(): FieldType = { return new LongType(isConstant, constValue) }
}

class ArrayType(t: FieldType) extends RefType {
    val itemType = t

    def descriptorString(): String = "[" + itemType.descriptorString()

    def cloneType(): FieldType = { return new ArrayType(t.cloneType()) }
}

class ObjectType(i: Boolean, s: String) extends RefType {
    def this(s: String) = this(false, s)

    var isUninitialized = i

    val name = s

    def descriptorString(): String = "L" + name + ";"

    def isJavaString(): Boolean = name == "java/lang/String"

    def isJavaObject(): Boolean = name == "java/lang/Object"

    def cloneType(): FieldType = {
        return new ObjectType(isUninitialized, name)
    }
}

class ParameterTypes extends Comparable[ParameterTypes] {
    var _parameters = new Vector[FieldType]()

    def add(field: FieldType) {
        _parameters.add(field)
    }

    def size(): Int = _parameters.size()

    def descriptorString(): String = {
        var result = "("
        for (p <- _parameters) {
            result += p.descriptorString()
        }
        result += ")"

        return result
    }

    def compareTo(other: ParameterTypes): Int = {
        return descriptorString().compareTo(other.descriptorString())
    }
}

class MethodType extends DescriptorType {
    var parameters = new ParameterTypes()
    var returnType: FieldType = null  // null for void

    def descriptorString(): String = {
        var result = parameters.descriptorString()

        if (returnType != null) {
            result += returnType.descriptorString()
        } else {
            result += "V"
        }
        return result
    }
}

//
// Only usable for stack frame verification
//

class UnusableType(t: Boolean) extends FieldType {
    var isTop = t  // only true for double / long

    def descriptorString(): String = "__top__"  // fake type

    def cloneType(): FieldType = { return new UnusableType(isTop) }
}

class NullType extends RefType {
    def descriptorString(): String = "__null__"  // fake type

    def cloneType(): FieldType = { return new NullType() }
}

//
// maybe use a real lexer/parser ...
//

object DescriptorToken {
    val BYTE = "B"
    val CHAR = "C"
    val DOUBLE = "D"
    val FLOAT = "F"
    val INT = "I"
    val LONG = "J"
    val OBJECT = "L"
    val SHORT = "S"
    val BOOL = "Z"
    val ARRAY = "["
    val VOID = "V"
    val LPARAM = "("
    val RPARAM = ")"
    val EOF = ""
}

class DescriptorTokenizer(d: String) {
    var descriptorString = d

    var nextPos = 0
    var value = ""

    def lookAhead(): Char = descriptorString.charAt(nextPos)

    def nextToken(): String = {
        value = ""
        if (nextPos >= descriptorString.length) {
            return DescriptorToken.EOF
        }

        value += descriptorString.charAt(nextPos)
        nextPos += 1
        value match {
            case DescriptorToken.BYTE => return DescriptorToken.BYTE
            case DescriptorToken.CHAR => return DescriptorToken.CHAR
            case DescriptorToken.DOUBLE => return DescriptorToken.DOUBLE
            case DescriptorToken.FLOAT => return DescriptorToken.FLOAT
            case DescriptorToken.INT => return DescriptorToken.INT
            case DescriptorToken.LONG => return DescriptorToken.LONG
            case DescriptorToken.SHORT => return DescriptorToken.SHORT
            case DescriptorToken.BOOL => return DescriptorToken.BOOL
            case DescriptorToken.ARRAY => return DescriptorToken.ARRAY
            case DescriptorToken.OBJECT => {
                val end = descriptorString.indexOf(';', nextPos - 1)
                if (end == -1) {
                    throw new Exception(
                            "malformed descriptor: " + descriptorString)
                }
                value = descriptorString.substring(nextPos, end)
                nextPos = end + 1
                return DescriptorToken.OBJECT
            }
            case DescriptorToken.VOID => return DescriptorToken.VOID
            case DescriptorToken.LPARAM => return DescriptorToken.LPARAM
            case DescriptorToken.RPARAM => return DescriptorToken.RPARAM
            case _ => throw new Exception(
                    "malformed descriptor: " + descriptorString)
        }
    }
}

class DescriptorParser(s: String) {
    val descriptorString = s
    var tokenizer = new DescriptorTokenizer(s)

    def _parseField(): FieldType = {
        tokenizer.nextToken() match {
            case DescriptorToken.BYTE => return new ByteType()
            case DescriptorToken.CHAR => return new CharType()
            case DescriptorToken.DOUBLE => return new DoubleType()
            case DescriptorToken.FLOAT => return new FloatType()
            case DescriptorToken.INT => return new IntType()
            case DescriptorToken.LONG => return new LongType()
            case DescriptorToken.SHORT => return new ShortType()
            case DescriptorToken.BOOL => return new BoolType()
            case DescriptorToken.ARRAY => return new ArrayType(_parseField())
            case DescriptorToken.OBJECT =>
                    return new ObjectType(tokenizer.value)
            case _ => throw new Exception(
                    "malformed descriptor: " + descriptorString)
        }
    }

    def parseFieldDescriptor(): FieldType = {
        val f = _parseField()
        tokenizer.nextToken() match {
            case DescriptorToken.EOF => return f
            case _ => throw new Exception(
                    "malformed descriptor: " + descriptorString)
        }
    }

    def parseMethodDescriptor(): MethodType = {
        if (tokenizer.nextToken() != DescriptorToken.LPARAM) {
            throw new Exception("malformed descriptor: " + descriptorString)
        }

        var method = new MethodType()

        var shouldContinue = true
        while (shouldContinue) {
            try {
                if (tokenizer.lookAhead() == ')') {
                    tokenizer.nextToken()
                    shouldContinue = false
                } else {
                    method.parameters.add(_parseField())
                }
            } catch {
                case ex: IndexOutOfBoundsException => throw new Exception(
                        "malformed descriptor: " + descriptorString)
            }
        }

        try {
            if (tokenizer.lookAhead() == 'V') {
                tokenizer.nextToken()
            } else {
                method.returnType = _parseField()
            }
        } catch {
            case ex: IndexOutOfBoundsException => throw new Exception(
                    "malformed descriptor: " + descriptorString)
        }

        tokenizer.nextToken() match {
            case DescriptorToken.EOF => return method
            case _ => throw new Exception(
                    "malformed descriptor: " + descriptorString)
        }
    }
}

