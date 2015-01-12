import java.util.Stack
import java.util.Vector

import scala.collection.JavaConversions._
import scala.reflect.ClassTag


class StackFrame {
    var maxStack = 0
    var maxLocals = 0

    var stack = new Stack[FieldType]()
    var locals = new Vector[FieldType]()

    def mergeFrom(other: StackFrame, mergeStack: Boolean): Boolean = {
        if (other.maxStack > maxStack) {
            maxStack = other.maxStack
        }
        if (other.maxLocals > maxLocals) {
            maxLocals = other.maxLocals
        }

        var modified = false
        if (mergeStack) {
            modified = _mergeStackFrom(other)
        }

        modified |= _mergeLocalsFrom(other)

        return modified
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

    def _mergeLocalsFrom(other: StackFrame): Boolean = {
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
        frame.maxStack = maxStack
        frame.maxLocals = maxLocals
        for (field <- locals) {
            frame.locals.add(field)
        }
        return frame
    }

    def push(f: FieldType) {
        f match {
            case _: UnusableType => {
                throw new Exception("cannot push unusable type")
            }
            case _: TopType => {
                throw new Exception("cannot directly push top type")
            }
            case _: LongType => {
                stack.push(f)
                stack.push(new TopType())
            }
            case _: DoubleType => {
                stack.push(f)
                stack.push(new TopType())
            }
            case _ => {
                stack.push(f)
            }
        }

        if (stack.size() > maxStack) {
            maxStack = stack.size()
        }
    }

    def pop(expected: FieldType): FieldType = {
        expected match {
            case _: UnusableType => {
                throw new Exception("cannot push unusable type")
            }
            case _: TopType => {
                throw new Exception("cannot directly push top type")
            }
            case _: IntType => return _pop[IntType]()
            case _: FloatType => return _pop[FloatType]()
            case _: LongType => {
                _pop[TopType]()
                return _pop[LongType]()
            }
            case _: DoubleType => {
                _pop[TopType]()
                return _pop[DoubleType]()
            }
            case _: ArrayType => return _pop[ArrayType]()
            case _: RefType => return _pop[RefType]()
            case _ => throw new Exception(
                    "Pop-ing unexpected type: " + expected.descriptorString())
        }
    }

    def _pop[T <: FieldType: ClassTag](): FieldType = {
        val f = stack.pop()

        val cls = implicitly[ClassTag[T]].runtimeClass
        f match {
            case t: T if cls.isInstance(t) => return t
            case _ => throw new Exception(
                    "Unexpected type: " + f.descriptorString())
        }
    }

    def store(expected: FieldType, i: Int) {
        val f = pop(expected)

        val needed = i + f.categorySize()
        while (needed > locals.size()) {
            locals.add(new UnusableType())
        }

        locals.elementAt(i) match {
            case _: TopType => {
                // invalidate previous long / double
                locals.setElementAt(new UnusableType(), i - 1)
            }
            case _: DoubleType => {
                // invalidate top
                locals.setElementAt(new UnusableType(), i + 1)
            }
            case _: LongType => {
                // invalidate top
                locals.setElementAt(new UnusableType(), i + 1)
            }
        }

        f match {
            case _: UnusableType => {
                throw new Exception("cannot store unusable type")
            }
            case _: TopType => {
                throw new Exception("cannot directly store top type")
            }
            case _: LongType => {
                locals.setElementAt(f, i)
                locals.setElementAt(new TopType(), i + 1)
            }
            case _: DoubleType => {
                locals.setElementAt(f, i)
                locals.setElementAt(new TopType(), i + 1)
            }
            case _ => {
                locals.setElementAt(f, i)
            }
        }

        if (locals.size() > maxLocals) {
            maxLocals = locals.size()
        }
    }

    def load(expected: FieldType, i: Int) {
        val f = locals.elementAt(i)
        f match {
            case _: UnusableType => {
                throw new Exception("cannot load unusable type")
            }
            case _: TopType => {
                throw new Exception("cannot load top type")
            }
            case _ => {}  // continue
        }

        expected match {
            case _: IntType => _check[IntType](f)
            case _: FloatType => _check[FloatType](f)
            case _: LongType => _check[LongType](f)
            case _: DoubleType => _check[DoubleType](f)
            case _: ArrayType => _check[ArrayType](f)
            case _: RefType => _check[RefType](f)
            case _ => throw new Exception(
                    "unexpected load type: " + expected.descriptorString())
        }

        push(f)
    }

    def _check[T <: FieldType: ClassTag](f: FieldType) {
        val cls = implicitly[ClassTag[T]].runtimeClass
        f match {
            case t: T if cls.isInstance(t) => {}
            case _ => throw new Exception(
                    "Unexpected type: " + f.descriptorString())
        }
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

    def apply(op: Operation) {
        if (_applyConstantOp(op)) {
            return
        }
        if (_applyLoadOp(op)) {
            return
        }
        if (_applyStoreOp(op)) {
            return
        }
        throw new Exception(
                "op not implemented in stack frame: " + op.debugString(""))
    }

    def _applyConstantOp(op: Operation): Boolean = {
        op match {
            case o: PushI => push(new IntType())
            case o: PushL => push(new LongType())
            case o: PushF => push(new FloatType())
            case o: PushD => push(new DoubleType())
            case o: PushString => push(new ObjectType(Const.JAVA_STRING))
            case o: Ldc => {
                // TODO: support class / method type / method handle ...
                throw new Exception("raw ldc unsupported")
            }
            case o: LdcW => {
                // TODO: support class / method type / method handle ...
                throw new Exception("raw ldc_w unsupported")
            }
            case _ => return false
        }

        return true
    }

    def _applyLoadOp(op: Operation): Boolean = {
        op match {
            case o: LoadI => load(new IntType(), o.index)
            case o: LoadL => load(new LongType(), o.index)
            case o: LoadF => load(new FloatType(), o.index)
            case o: LoadD => load(new DoubleType(), o.index)
            case o: LoadA => load(new CheckRefType(), o.index)
            case o: LoadFromIArray => {
                throw new Exception("TODO")
            }
            case o: LoadFromBArray => {  // byte (or bool) array
                throw new Exception("TODO")
            }
            case o: LoadFromCArray => {  // char array
                throw new Exception("TODO")
            }
            case o: LoadFromSArray => {  // short array
                throw new Exception("TODO")
            }
            case o: LoadFromLArray => {
                throw new Exception("TODO")
            }
            case o: LoadFromFArray => {
                throw new Exception("TODO")
            }
            case o: LoadFromDArray => {
                throw new Exception("TODO")
            }
            case o: LoadFromAArray => {
                throw new Exception("TODO")
            }
            case _ => return false
        }

        return true
    }

    def _applyStoreOp(op: Operation): Boolean = {
        op match {
            case o: StoreI => store(new IntType(), o.index)
            case o: StoreL => store(new LongType(), o.index)
            case o: StoreF => store(new FloatType(), o.index)
            case o: StoreD => store(new DoubleType(), o.index)
            case o: StoreA => store(new CheckRefType(), o.index)
            case o: StoreIntoIArray => {
                throw new Exception("TODO")
            }
            case o: StoreIntoBArray => {  // byte (or bool) array
                throw new Exception("TODO")
            }
            case o: StoreIntoCArray => {  // char array
                throw new Exception("TODO")
            }
            case o: StoreIntoSArray => {  // short array
                throw new Exception("TODO")
            }
            case o: StoreIntoLArray => {
                throw new Exception("TODO")
            }
            case o: StoreIntoFArray => {
                throw new Exception("TODO")
            }
            case o: StoreIntoDArray => {
                throw new Exception("TODO")
            }
            case o: StoreIntoAArray => {
                throw new Exception("TODO")
            }
            case _ => return false
        }

        return true
    }
}
