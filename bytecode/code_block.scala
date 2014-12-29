import java.io.DataInputStream
import java.io.DataOutputStream
import java.util.Vector

class LineContext {
    var lineNumber = -1
}

abstract class CodeSegment(
        owner: AttributeOwner,
        parent: CodeSection,
        context: LineContext) extends Operation(owner) {
    var _parentScope: CodeSection = parent
    var _lineContext: LineContext = context

    var segmentNumber = -1

    def _assignAddress(startAddress: Int): Int
}

// a linear section in the control flow graph.
//
// Only the last op maybe a jump or a return.  No other ops in the block
// can be a jump (or a return).
//
// if implicitGoto is set and the last ops is not a jump/return, then
// and goto is inserted to the end of the block during verification.
class CodeBlock(parent: CodeSection)
        extends CodeSegment(parent._owner, parent, parent._lineContext) {
    var _ops = new Vector[Operation]()

    var _hasControlOp = false

    var implicitGoto: CodeBlock = null

    def _add(op: Operation) {
        if (_hasControlOp) {
            throw new Exception("Cannot add more ops after control op")
        }
        op.line = _lineContext.lineNumber
        _ops.add(op)
    }

    def _addControl(op: Operation) {
        if (_hasControlOp) {
            throw new Exception("Cannot add more ops after control op")
        }
        _hasControlOp = true
        op.line = _lineContext.lineNumber
        _ops.add(op)
    }

    //
    // Non-control operations
    //

    def pushNull() { _add(new AconstNull(_owner)) }
    def pushI(v: Int) { _add(new PushI(_owner, v)) }
    def pushL(v: Long) { _add(new PushL(_owner, v)) }
    def pushF(v: Float) { _add(new PushF(_owner, v)) }
    def pushD(v: Double) { _add(new PushD(_owner, v)) }
    def pushString(v: String) { _add(new PushString(_owner, v)) }
    // TODO: support other ldc constant types

    def loadI(index: Int) { _add(new LoadI(_owner, index)) }
    def loadL(index: Int) { _add(new LoadL(_owner, index)) }
    def loadF(index: Int) { _add(new LoadF(_owner, index)) }
    def loadD(index: Int) { _add(new LoadD(_owner, index)) }
    def loadA(index: Int) { _add(new LoadA(_owner, index)) }

    def storeI(index: Int) { _add(new StoreI(_owner, index)) }
    def storeL(index: Int) { _add(new StoreL(_owner, index)) }
    def storeF(index: Int) { _add(new StoreF(_owner, index)) }
    def storeD(index: Int) { _add(new StoreD(_owner, index)) }
    def storeA(index: Int) { _add(new StoreA(_owner, index)) }

    def loadFromIArray() { _add(new LoadFromIArray(_owner)) }
    def loadFromLArray() { _add(new LoadFromLArray(_owner)) }
    def loadFromFArray() { _add(new LoadFromFArray(_owner)) }
    def loadFromDArray() { _add(new LoadFromDArray(_owner)) }
    def loadFromAArray() { _add(new LoadFromAArray(_owner)) }
    def loadFromBArray() { _add(new LoadFromBArray(_owner)) }
    def loadFromCArray() { _add(new LoadFromCArray(_owner)) }
    def loadFromSArray() { _add(new LoadFromSArray(_owner)) }

    def storeIntoIArray() { _add(new StoreIntoIArray(_owner)) }
    def storeIntoLArray() { _add(new StoreIntoLArray(_owner)) }
    def storeIntoFArray() { _add(new StoreIntoFArray(_owner)) }
    def storeIntoDArray() { _add(new StoreIntoDArray(_owner)) }
    def storeIntoAArray() { _add(new StoreIntoAArray(_owner)) }
    def storeIntoBArray() { _add(new StoreIntoBArray(_owner)) }
    def storeIntoCArray() { _add(new StoreIntoCArray(_owner)) }
    def storeIntoSArray() { _add(new StoreIntoSArray(_owner)) }

    def pop() { _add(new Pop(_owner)) }
    def pop2() { _add(new Pop2(_owner)) }
    def dup() { _add(new Dup(_owner)) }
    def dupX1() { _add(new DupX1(_owner)) }
    def dupX2() { _add(new DupX2(_owner)) }
    def dup2() { _add(new Dup2(_owner)) }
    def dup2X1() { _add(new Dup2X1(_owner)) }
    def dup2X2() { _add(new Dup2X2(_owner)) }
    def swap() { _add(new Swap(_owner)) }

    def addI() { _add(new Iadd(_owner)) }
    def subI() { _add(new Isub(_owner)) }
    def mulI() { _add(new Imul(_owner)) }
    def divI() { _add(new Idiv(_owner)) }
    def remI() { _add(new Irem(_owner)) }
    def negI() { _add(new Ineg(_owner)) }
    def shlI() { _add(new Ishl(_owner)) }
    def shrI() { _add(new Ishr(_owner)) }
    def ushrI() { _add(new Iushr(_owner)) }
    def andI() { _add(new Iand(_owner)) }
    def orI() { _add(new Ior(_owner)) }
    def xorI() { _add(new Ixor(_owner)) }
    def incI(index: Int, v: Int) { _add(new Iinc(_owner, index, v)) }

    def addL() { _add(new Ladd(_owner)) }
    def subL() { _add(new Lsub(_owner)) }
    def mulL() { _add(new Lmul(_owner)) }
    def divL() { _add(new Ldiv(_owner)) }
    def remL() { _add(new Lrem(_owner)) }
    def negL() { _add(new Lneg(_owner)) }
    def shlL() { _add(new Lshl(_owner)) }
    def shrL() { _add(new Lshr(_owner)) }
    def ushrL() { _add(new Lushr(_owner)) }
    def andL() { _add(new Land(_owner)) }
    def orL() { _add(new Lor(_owner)) }
    def xorL() { _add(new Lxor(_owner)) }

    def addF() { _add(new Fadd(_owner)) }
    def subF() { _add(new Fsub(_owner)) }
    def mulF() { _add(new Fmul(_owner)) }
    def divF() { _add(new Fdiv(_owner)) }
    def remF() { _add(new Frem(_owner)) }
    def negF() { _add(new Fneg(_owner)) }

    def addD() { _add(new Dadd(_owner)) }
    def subD() { _add(new Dsub(_owner)) }
    def mulD() { _add(new Dmul(_owner)) }
    def divD() { _add(new Ddiv(_owner)) }
    def remD() { _add(new Drem(_owner)) }
    def negD() { _add(new Dneg(_owner)) }

    def i2l() { _add(new I2l(_owner)) }
    def i2f() { _add(new I2f(_owner)) }
    def i2d() { _add(new I2d(_owner)) }
    def l2i() { _add(new L2i(_owner)) }
    def l2f() { _add(new L2f(_owner)) }
    def l2d() { _add(new L2d(_owner)) }
    def f2i() { _add(new F2i(_owner)) }
    def f2l() { _add(new F2l(_owner)) }
    def f2d() { _add(new F2d(_owner)) }
    def d2i() { _add(new D2i(_owner)) }
    def d2l() { _add(new D2l(_owner)) }
    def d2f() { _add(new D2f(_owner)) }
    def i2b() { _add(new I2b(_owner)) }
    def i2c() { _add(new I2c(_owner)) }
    def i2s() { _add(new I2s(_owner)) }

    def cmpL() { _add(new Lcmp(_owner)) }
    def cmplF() { _add(new Fcmpl(_owner)) }
    def cmpgF() { _add(new Fcmpg(_owner)) }
    def cmplD() { _add(new Dcmpl(_owner)) }
    def cmpgD() { _add(new Dcmpg(_owner)) }

    def getStatic(
            className: String,
            fieldName: String,
            fieldType: FieldType) {
        _add(new Getstatic(_owner, className, fieldName, fieldType))
    }
    def putStatic(
            className: String,
            fieldName: String,
            fieldType: FieldType) {
        _add(new Putstatic(_owner, className, fieldName, fieldType))
    }
    def getField(
            className: String,
            fieldName: String,
            fieldType: FieldType) {
        _add(new Getfield(_owner, className, fieldName, fieldType))
    }
    def putField(
            className: String,
            fieldName: String,
            fieldType: FieldType) {
        _add(new Putfield(_owner, className, fieldName, fieldType))
    }

    def invokeVirtual(
            className: String,
            methodName: String,
            methodType: MethodType) {
        _add(new Invokevirtual(_owner, className, methodName, methodType))
    }
    def invokeSpecial(
            className: String,
            methodName: String,
            methodType: MethodType) {
        _add(new Invokespecial(_owner, className, methodName, methodType))
    }
    def invokeStatic(
            className: String,
            methodName: String,
            methodType: MethodType) {
        _add(new Invokestatic(_owner, className, methodName, methodType))
    }
    def invokeInterface(
            className: String,
            methodName: String,
            methodType: MethodType) {
        _add(new Invokeinterface(_owner, className, methodName, methodType))
    }

    def newA(t: ObjectType) { _add(new New(_owner, t)) }
    def newArray(t: BaseType) { _add(new Newarray(_owner, t)) }
    def aNewArray(t: RefType) { _add(new Anewarray(_owner, t)) }
    def multiANewArray(t: ArrayType, d: Int) {
        _add(new Multianewarray(_owner, t, d))
    }

    def arrayLength() { _add(new Arraylength(_owner)) }

    def checkCast(t: ObjectType) { _add(new Checkcast(_owner, t)) }
    def instanceOf(t: ObjectType) { _add(new Instanceof(_owner, t)) }

    def enterMonitor() { _add(new Monitorenter(_owner)) }
    def exitMonitor() { _add(new Monitorexit(_owner)) }

    // TODO table / lookup switch

    //
    // Control operations
    //

    def ifEq(ifBlock: CodeBlock, elseBlock: CodeBlock) {
        _addControl(new Ifeq(_owner, ifBlock, elseBlock))
    }
    def ifNe(ifBlock: CodeBlock, elseBlock: CodeBlock) {
        _addControl(new Ifne(_owner, ifBlock, elseBlock))
    }
    def ifLt(ifBlock: CodeBlock, elseBlock: CodeBlock) {
        _addControl(new Iflt(_owner, ifBlock, elseBlock))
    }
    def ifGe(ifBlock: CodeBlock, elseBlock: CodeBlock) {
        _addControl(new Ifge(_owner, ifBlock, elseBlock))
    }
    def ifGt(ifBlock: CodeBlock, elseBlock: CodeBlock) {
        _addControl(new Ifgt(_owner, ifBlock, elseBlock))
    }
    def ifLe(ifBlock: CodeBlock, elseBlock: CodeBlock) {
        _addControl(new Ifle(_owner, ifBlock, elseBlock))
    }
    def ifICmpEq(ifBlock: CodeBlock, elseBlock: CodeBlock) {
        _addControl(new IfIcmpeq(_owner, ifBlock, elseBlock))
    }
    def ifICmpNe(ifBlock: CodeBlock, elseBlock: CodeBlock) {
        _addControl(new IfIcmpne(_owner, ifBlock, elseBlock))
    }
    def ifICmpLt(ifBlock: CodeBlock, elseBlock: CodeBlock) {
        _addControl(new IfIcmplt(_owner, ifBlock, elseBlock))
    }
    def ifICmpGe(ifBlock: CodeBlock, elseBlock: CodeBlock) {
        _addControl(new IfIcmpge(_owner, ifBlock, elseBlock))
    }
    def ifICmpGt(ifBlock: CodeBlock, elseBlock: CodeBlock) {
        _addControl(new IfIcmpgt(_owner, ifBlock, elseBlock))
    }
    def ifICmpLe(ifBlock: CodeBlock, elseBlock: CodeBlock) {
        _addControl(new IfIcmple(_owner, ifBlock, elseBlock))
    }
    def ifACmpEq(ifBlock: CodeBlock, elseBlock: CodeBlock) {
        _addControl(new IfAcmpeq(_owner, ifBlock, elseBlock))
    }
    def ifACmpNe(ifBlock: CodeBlock, elseBlock: CodeBlock) {
        _addControl(new IfAcmpne(_owner, ifBlock, elseBlock))
    }
    def ifNull(ifBlock: CodeBlock, elseBlock: CodeBlock) {
        _addControl(new Ifnull(_owner, ifBlock, elseBlock))
    }
    def ifNonNull(ifBlock: CodeBlock, elseBlock: CodeBlock) {
        _addControl(new Ifnonnull(_owner, ifBlock, elseBlock))
    }

    def goto(target: CodeBlock) { _addControl(new Goto(_owner, this, target)) }

    // XXX: maybe infer return type from method signature?
    def returnI() { _addControl(new Ireturn(_owner)) }
    def returnL() { _addControl(new Lreturn(_owner)) }
    def returnF() { _addControl(new Freturn(_owner)) }
    def returnD() { _addControl(new Dreturn(_owner)) }
    def returnA() { _addControl(new Areturn(_owner)) }
    def returnVoid() { _addControl(new Return(_owner)) }

    def throwA() { _addControl(new Athrow(_owner)) }

    def _assignAddress(startAddress: Int): Int = {
        throw new Exception("TODO")
    }

    def serialize(output: DataOutputStream) {
        throw new Exception("TODO")
    }

    def deserialize(startAddress: Int, opCode: Int, input: DataInputStream) {
        throw new Exception("TODO")
    }

    def debugString(indent: String): String = {
        throw new Exception("TODO")
    }
}

class ExceptionTarget(e: ConstClassInfo, t: CodeSection) {
    var exception: ConstClassInfo = e
    var target: CodeSection = t
}

// a group of code segments which will be written out as a single continuous
// unit.
class CodeSection(
        owner: AttributeOwner,
        parent: CodeSection,
        context: LineContext) extends CodeSegment(owner, parent, context) {
    def this(parent: CodeSection) = this(
            parent._owner,
            parent,
            parent._lineContext)

    var _segments = new Vector[CodeSegment]()
    _segments.add(new CodeBlock(this))  // entry block

    var _exceptionTargets = new Vector[ExceptionTarget]()

    def entryBlock(): CodeBlock = {
        _segments.elementAt(0) match {
            case block: CodeBlock => return block
            case _ => throw new Exception("entry point should be a block ...")
        }
    }

    def newBlock(): CodeBlock = {
        val block = new CodeBlock(this)
        _segments.add(block)
        return block
    }

    def newSubSection(): CodeSection = {
        val section = new CodeSection(this)
        _segments.add(section)
        return section
    }

    // add new exception handle for the current section.
    // NOTE:
    //   - call order matters.
    //   - pass in null to catch all
    //   - cannot use this on top scope
    def newExceptionHandle(exceptionClassName: String): CodeSection = {
        if (_parentScope == null) {
            throw new Exception(
                    "cannot create exception handle for top level scope")
        }

        var exception: ConstClassInfo = null
        if (exceptionClassName != null) {
            exception =_owner.constants().getClass(exceptionClassName)
        }

        val target = _parentScope.newSubSection()

        _exceptionTargets.add(new ExceptionTarget(exception, target))
        return target
    }

    def _assignSegmentNumber(startNumber: Int): Int = {
        throw new Exception("TODO")
    }

    def _assignAddress(startAddress: Int): Int = {
        throw new Exception("TODO")
    }

    def serialize(output: DataOutputStream) {
        throw new Exception("TODO")
    }

    def deserialize(startAddress: Int, opCode: Int, input: DataInputStream) {
        throw new Exception("TODO")
    }

    def debugString(indent: String): String = {
        throw new Exception("TODO")
    }
}
