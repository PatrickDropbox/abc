import java.util.Vector

import scala.collection.JavaConversions._


trait DescriptorType extends Comparable[DescriptorType] {
    def descriptorString(): String

    def compareTo(other: DescriptorType): Int = {
        return descriptorString().compareTo(other.descriptorString())
    }
}

trait FieldType extends DescriptorType {
}

trait BaseType extends FieldType {
}

class ByteType extends BaseType {
    def descriptorString(): String = "B"
}

class CharType extends BaseType {
    def descriptorString(): String = "C"
}

class DoubleType extends BaseType {
    def descriptorString(): String = "D"
}

class FloatType extends BaseType {
    def descriptorString(): String = "F"
}

class IntType extends BaseType {
    def descriptorString(): String = "I"
}

class LongType extends BaseType {
    def descriptorString(): String = "J"
}

class ShortType extends BaseType {
    def descriptorString(): String = "S"
}

class BoolType extends BaseType {
    def descriptorString(): String = "Z"
}

class ArrayType(t: FieldType) extends FieldType {
    val itemType = t

    def descriptorString(): String = "[" + itemType.descriptorString()
}

class ObjectType(s: String) extends FieldType {
    val name = s

    def descriptorString(): String = name
}

class MethodType extends DescriptorType {
    var parameters = new Vector[FieldType]()
    var returnType: FieldType = null  // null for void

    def descriptorString(): String = {
        var result = "("
        for (p <- parameters) {
            result += p.descriptorString()
        }
        result += ")"

        if (returnType != null) {
            result += returnType.descriptorString()
        } else {
            result += "V"
        }
        return result
    }
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
                value = descriptorString.substring(nextPos - 1, end + 1)
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

