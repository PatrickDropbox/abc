import java.util.Stack
import java.util.Vector

import scala.collection.JavaConversions._


class StackFrame {
    var maxStack = 0
    var maxLocals = 0

    var stack = new Stack[FieldType]()
    var locals = new Vector[FieldType]()

    def mergeFrom(other: StackFrame): Boolean = {
        return _mergeStackFrom(other) || mergeLocalsFrom(other)
    }

    def _mergeStackFrom(other: StackFrame): Boolean = {
        if (stack.size() != other.stack.size()) {
            throw new Exception("different stack height ...")
        }

        var modified = false
        var updatedStack = new Stack[FieldType]()
        for (i <- 0.to(stack.size() - 1)) {
            val origType = stack.elementAt(i)
            val newType = _mergeTypes(origType, other.stack.elementAt(i), false)
            if (origType.compareTo(newType) != 0) {
                modified = true
            }
            updatedStack.add(newType)
        }

        stack = updatedStack
        return modified
    }

    def mergeLocalsFrom(other: StackFrame): Boolean = {
        var modified = locals.size() != other.locals.size()
        var updatedLocals = new Vector[FieldType]()

        var minSize = locals.size()
        if (other.locals.size() < minSize) {
            minSize = other.locals.size()
        }

        for (i <- 0.to(minSize - 1)) {
            val origType = stack.elementAt(i)
            val newType = _mergeTypes(origType, other.stack.elementAt(i), true)
            if (origType.compareTo(newType) != 0) {
                modified = true
            }
            updatedLocals.add(newType)
        }

        locals = updatedLocals
        return modified
    }

    def cloneFrame(): StackFrame = {
        var frame = cloneLocals()
        for (field <- stack) {
            frame.stack.add(field)
        }
        return frame
    }

    def cloneLocals(): StackFrame = {
        var frame = new StackFrame()
        for (field <- locals) {
            frame.locals.add(field)
        }
        return frame
    }

    def _mergeTypes(
            v1: FieldType,
            v2: FieldType,
            mergeIncompatibleTypes: Boolean): FieldType = {
        if (v1.compareTo(v2) == 0) {
            return v1
        }

        var result: FieldType = v1 match {
            case a1: ArrayType => {
                v2 match {
                    case a2: ArrayType => _mergeArrayTypes(a1, a2)
                    case _: RefType => return new ObjectType(Const.JAVA_OBJECT)
                    case _ => null
                }
            }
            case o1: ObjectType => {
                v2 match {
                    case o2: ObjectType => _mergeObjectTypes(o1, o2)
                    case _: RefType => return new ObjectType(Const.JAVA_OBJECT)
                    case _ => null
                }
            }
            case _ => null
        }

        if (result != null) {
            return result
        }

        if (mergeIncompatibleTypes) {
            return new UnusableType()  // not "top"
        }

        throw new Exception(
                "merging incompatible types: " + v1.descriptorString() +
                        "vs. " + v2.descriptorString())
    }

    // This assumes v1.compareTo(v2) != 0
    def _mergeObjectTypes(v1: ObjectType, v2: ObjectType): ObjectType = {
        // TODO: load class catalog and find most specific common supertype.
        return new ObjectType(Const.JAVA_OBJECT)
    }

    // This assumes v1.compareTo(v2) != 0
    def _mergeArrayTypes(v1: ArrayType, v2: ArrayType): RefType = {
        // TODO: fix array merge semantic (see page 350 / section 4.10.2.2)
        val commonItemType: FieldType = v1.itemType match {
            case _: BaseType => new ObjectType(Const.JAVA_OBJECT)
            case a1: ArrayType => {
                v2.itemType match {
                    case a2: ArrayType => _mergeArrayTypes(a1, a2)
                    case _ => new ObjectType(Const.JAVA_OBJECT)
                }
            }
            case o1: ObjectType => {
                v2.itemType match {
                    case o2: ObjectType => _mergeObjectTypes(o1, o2)
                    case _ => new ObjectType(Const.JAVA_OBJECT)
                }
            }
        }

        return new ArrayType(commonItemType)
    }

}
