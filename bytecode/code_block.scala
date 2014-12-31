import java.io.DataInputStream
import java.io.DataOutput
import java.util.Collections
import java.util.HashMap
import java.util.HashSet
import java.util.Stack
import java.util.TreeMap
import java.util.TreeSet
import java.util.Vector

import scala.collection.JavaConversions._


abstract class CodeSegment(
        owner: AttributeOwner,
        parent: CodeSection)
        extends Operation(owner) with Comparable[CodeSegment] {
    var _parentScope: CodeSection = parent

    // not inclusive.
    var _endPc = -1

    var segmentId = -1

    def _assignAddress(startAddress: Int): Int

    def _insertImplicitGoto(): CodeBlock

    def _resetPcIds()

    def compareTo(other: CodeSegment): Int = {
        if (segmentId < other.segmentId) {
            return -1
        } else if (segmentId > other.segmentId) {
            return 1
        }

        if (pc < other.pc) {
            return -1
        } else if (pc > other.pc) {
            return 1
        }

        return 0
    }
}

// a linear section in the control flow graph.
//
// Only the last op maybe a jump or a return.  No other ops in the block
// can be a jump (or a return).
//
// if implicitGoto is set and the last ops is not a jump/return, then
// and goto is inserted to the end of the block during verification.
class CodeBlock(owner: AttributeOwner)
        extends CodeSegment(owner, null) {

    var isEntryPoint = false  // for the method.  there can only be one.

    var lineContext = -1

    var _ops = new Vector[Operation]()

    var _hasControlOp = false

    var implicitGoto: CodeBlock = null

    def _add(op: Operation) {
        if (_hasControlOp) {
            throw new Exception(
                    "Cannot add more ops after control op: " +
                            _ops.lastElement().pc)
        }

        if (lineContext >= 0) {
            op.line = lineContext
        }
        _ops.add(op)

        _hasControlOp = op match {
            case _: Return => true
            case _: ReturnValue => true
            case _: Athrow => true
            case g: Goto => true
            case i: IfBaseOp => true
            case s: Switch => true
            case _ => false
        }
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
        _add(new Ifeq(_owner, ifBlock, elseBlock))
    }
    def ifNe(ifBlock: CodeBlock, elseBlock: CodeBlock) {
        _add(new Ifne(_owner, ifBlock, elseBlock))
    }
    def ifLt(ifBlock: CodeBlock, elseBlock: CodeBlock) {
        _add(new Iflt(_owner, ifBlock, elseBlock))
    }
    def ifGe(ifBlock: CodeBlock, elseBlock: CodeBlock) {
        _add(new Ifge(_owner, ifBlock, elseBlock))
    }
    def ifGt(ifBlock: CodeBlock, elseBlock: CodeBlock) {
        _add(new Ifgt(_owner, ifBlock, elseBlock))
    }
    def ifLe(ifBlock: CodeBlock, elseBlock: CodeBlock) {
        _add(new Ifle(_owner, ifBlock, elseBlock))
    }
    def ifICmpEq(ifBlock: CodeBlock, elseBlock: CodeBlock) {
        _add(new IfIcmpeq(_owner, ifBlock, elseBlock))
    }
    def ifICmpNe(ifBlock: CodeBlock, elseBlock: CodeBlock) {
        _add(new IfIcmpne(_owner, ifBlock, elseBlock))
    }
    def ifICmpLt(ifBlock: CodeBlock, elseBlock: CodeBlock) {
        _add(new IfIcmplt(_owner, ifBlock, elseBlock))
    }
    def ifICmpGe(ifBlock: CodeBlock, elseBlock: CodeBlock) {
        _add(new IfIcmpge(_owner, ifBlock, elseBlock))
    }
    def ifICmpGt(ifBlock: CodeBlock, elseBlock: CodeBlock) {
        _add(new IfIcmpgt(_owner, ifBlock, elseBlock))
    }
    def ifICmpLe(ifBlock: CodeBlock, elseBlock: CodeBlock) {
        _add(new IfIcmple(_owner, ifBlock, elseBlock))
    }
    def ifACmpEq(ifBlock: CodeBlock, elseBlock: CodeBlock) {
        _add(new IfAcmpeq(_owner, ifBlock, elseBlock))
    }
    def ifACmpNe(ifBlock: CodeBlock, elseBlock: CodeBlock) {
        _add(new IfAcmpne(_owner, ifBlock, elseBlock))
    }
    def ifNull(ifBlock: CodeBlock, elseBlock: CodeBlock) {
        _add(new Ifnull(_owner, ifBlock, elseBlock))
    }
    def ifNonNull(ifBlock: CodeBlock, elseBlock: CodeBlock) {
        _add(new Ifnonnull(_owner, ifBlock, elseBlock))
    }

    def goto(target: CodeBlock) { _add(new Goto(_owner, this, target)) }

    // XXX: maybe infer return type from method signature?
    def returnI() { _add(new Ireturn(_owner)) }
    def returnL() { _add(new Lreturn(_owner)) }
    def returnF() { _add(new Freturn(_owner)) }
    def returnD() { _add(new Dreturn(_owner)) }
    def returnA() { _add(new Areturn(_owner)) }
    def returnVoid() { _add(new Return(_owner)) }

    def throwA() { _add(new Athrow(_owner)) }

    def _insertImplicitGoto(): CodeBlock = {
        if (!_hasControlOp) {
            if (implicitGoto == null) {
                throw new Exception("no implicit goto - pc: " + pc)
            }
            goto(implicitGoto)
            return null
        }

        var lastOp = _ops.lastElement()
        lastOp match {
            // this simplifies control flow flattening since we no
            // longer need to pair the else branch immediately after the
            // condition operation.
            case i: IfBaseOp => {
                var indirection = new CodeBlock(_owner)
                indirection.goto(i._elseBranch)
                i._elseBranch = indirection
                return indirection
            }
            case _ => {}
        }

        return null
    }

    def _assignAddress(startAddress: Int): Int = {
        throw new Exception("TODO")
    }

    def _resetPcIds() {
        /* TODO
        pc = -1
        _endPc = -1
        */
        segmentId = -1
    }

    def serialize(output: DataOutput) {
        throw new Exception("TODO")
    }

    def deserialize(startAddress: Int, opCode: Int, input: DataInputStream) {
        throw new Exception("TODO")
    }

    def debugString(indent: String): String = {
        var result = indent + "Block (pc: [" + pc + ", " + _endPc +
                ") segment: " + segmentId + ")\n"
        for (op <- _ops) {
            result += op.debugString(indent + "  ")
        }
        return result
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
        parent: CodeSection) extends CodeSegment(owner, parent) {
    var _mapId = -1  // unique unordered id

    def this(parent: CodeSection) = this(parent._owner, parent)

    var _segments = new Vector[CodeSegment]()
    var _blocks = new Vector[CodeBlock]()
    var _subsections = new Vector[CodeSection]()

    var _exceptionTargets = new Vector[ExceptionTarget]()

    def newBlock(): CodeBlock = {
        return _addBlock(new CodeBlock(_owner))
    }

    def _addBlock(block: CodeBlock): CodeBlock = {
        block._parentScope = this
        _segments.add(block)
        _blocks.add(block)
        return block
    }

    def newSubSection(): CodeSection = {
        val section = new CodeSection(this)
        _segments.add(section)
        _subsections.add(section)
        return section
    }

    // only used for reconstruction
    def _getMostSpecificSection(pc: Int): CodeSection = {
        for (s <- _subsections) {
            if (s.pc <= pc && pc < s._endPc) {
                return s._getMostSpecificSection(pc)
            }
        }

        return this
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

    // For sharing exception section (needed for javac)
    def shareExceptionHandle(exceptionClassName: String, target: CodeSection) {
        if (_parentScope == null) {
            throw new Exception(
                    "cannot create exception handle for top level scope")
        }

        var exception: ConstClassInfo = null
        if (exceptionClassName != null) {
            exception =_owner.constants().getClass(exceptionClassName)
        }

        _exceptionTargets.add(new ExceptionTarget(exception, target))
    }

    // this assumes pc are assigned and segments are sorted
    def collectExceptionEntries(result: Vector[ExceptionEntry]) {
        for (section <- _subsections) {
            section.collectExceptionEntries(result)
        }

        for (entry <- _exceptionTargets) {
            result.add(new ExceptionEntry(
                    _owner,
                    pc,
                    _endPc,
                    entry.target.pc,
                    entry.exception))
        }
    }

    // returns true if this equals, or contains, the other section
    def _contains(other: CodeSection): Boolean = {
        if (this == other) {
            return true
        }

        var tmp = other
        while (tmp._parentScope != null) {
            if (this == tmp._parentScope) {
                return true
            }
            tmp = tmp._parentScope
        }

        return false
    }

    def getEntryPoint(): CodeBlock = {
        var results = new Vector[CodeBlock]
        _collectEntryPoints(results)

        if (results.isEmpty()) {
            throw new Exception("no entry point")
        }

        if (results.size() > 1) {
            throw new Exception("multiple entry points")
        }

        return results.firstElement()
    }

    def _collectEntryPoints(entries: Vector[CodeBlock]) {
        for (seg <- _segments) {
            seg match {
                case b: CodeBlock => {
                    if (b.isEntryPoint) {
                        entries.add(b)
                    }
                }
                case s: CodeSection => s._collectEntryPoints(entries)
            }
        }
    }

    def _collectBlocks(result: Vector[CodeBlock]) {
        for (seg <- _segments) {
            seg match {
                case b: CodeBlock => result.add(b)
                case s: CodeSection => s._collectEntryPoints(result)
            }
        }
    }

    def _insertImplicitGoto(): CodeBlock = {
        var indirections = new Vector[CodeBlock]()
        for (seg <- _segments) {
            val block = seg._insertImplicitGoto()
            if (block != null) {
                indirections.add(block)
            }
        }
        for (block <- indirections) {
            _addBlock(block)
        }
        return null
    }

    def _assignMapId(
            startNumber: Int,
            mapping: HashMap[Int, CodeSection]): Int = {
        var number = startNumber
        for (seg <- _subsections) {
            number = seg._assignMapId(number, mapping)
        }

        _mapId = number
        mapping.put(number, this)

        number += 1
        return number
    }

    def _resetPcIds() {
        for (seg <- _subsections) {
            seg._resetPcIds()
        }

        /* TODO
        pc = -1
        _endPc = -1
        */
        segmentId = -1
        _mapId = -1
    }

    def _assignAddress(startAddress: Int): Int = {
        throw new Exception("TODO")
    }

    // only used for reconstruction
    def _fixPcs() {
        for (s <- _subsections) {
            s._fixPcs()
        }

        pc = -1
        _endPc = -1
        for (seg <- _segments) {
            if (pc == -1) {
                pc = seg.pc
                _endPc = seg._endPc
            } else {
                if (seg.pc < pc) {
                    pc = seg.pc
                }

                if (seg._endPc > _endPc) {
                    _endPc = seg._endPc
                }
            }
        }
    }

    def serialize(output: DataOutput) {
        _resetPcIds()

        _insertImplicitGoto()

        var segmentIdAssigner = new SegmentIdAssigner(this)
        segmentIdAssigner.assignIds()

        // TODO
    }

    def deserialize(startAddress: Int, opCode: Int, input: DataInputStream) {
        throw new Exception("TODO")
    }

    def sort() {
        Collections.sort(_segments)
        Collections.sort(_blocks)
        Collections.sort(_subsections)

        for (section <- _subsections) {
            section.sort()
        }
    }

    def debugString(indent: String): String = {
        sort()

        var result = indent + "Section (pc: [" + pc + ", " + _endPc +
                ") segment: " + segmentId + ")\n"
        for (segment <- _segments) {
            result += segment.debugString(indent + "  ")
        }

        return result
    }
}

// NOTE/WARNING: the algorithm extremely inefficient/naive, but should be ok
// since # segment << # of ops.  Fix/optimize later as needed.
class SegmentIdAssigner(root: CodeSection) {
    if (root._parentScope != null) {
        throw new Exception("not root ...")
    }

    var rootSection = root

    // section's map id -> section
    var sectionMap: HashMap[Int, CodeSection] = null

    // section's map id -> stack
    var stacksMap: HashMap[Int, Stack[CodeBlock]] = null

    var scopeStack: Stack[CodeSection] = null

    var currentScope: CodeSection = null
    var currentStack: Stack[CodeBlock] = null

    var nextSegmentId = 1

    def _init() {
        sectionMap = new HashMap[Int, CodeSection]()
        rootSection._assignMapId(0, sectionMap)

        scopeStack = new Stack[CodeSection]()

        var entryBlock = rootSection.getEntryPoint()
        var tmp = entryBlock._parentScope
        while (tmp != null) {
            scopeStack.add(0, tmp)
            tmp = tmp._parentScope
        }

        stacksMap = new HashMap[Int, Stack[CodeBlock]]()
        for (i <- sectionMap.keySet()) {
            stacksMap.put(i, new Stack[CodeBlock]())
        }

        currentScope = entryBlock._parentScope
        currentStack = stacksMap.get(entryBlock._parentScope._mapId)
        currentStack.push(entryBlock)
    }

    def assignIds() {
        _init()
        while (!scopeStack.isEmpty()) {
            if (currentStack.isEmpty()) {
                _updateStacks()
            } else {
                _assignId()
            }
        }
    }

    def _updateStacks() {

        var minId = nextSegmentId

        // try assigning blocks in current scope first, since subsections tend
        // to be exceptions related.
        for (block <- currentScope._blocks) {
            if (block.segmentId < 0) {
                currentStack.push(block)
                return
            }
            if (minId > block.segmentId) {
                minId = block.segmentId
            }
        }

        for (section <- currentScope._subsections) {
            if (section.segmentId < 0) {
                currentScope = section
                currentStack = stacksMap.get(section._mapId)
                scopeStack.push(section)
                return
            }
            if (minId > section.segmentId) {
                minId = section.segmentId
            }
        }

        currentScope.segmentId = minId
        scopeStack.pop()

        if (scopeStack.isEmpty()) {
            currentScope = null
            currentStack = null
        } else {
            currentScope = scopeStack.peek()
            currentStack = stacksMap.get(currentScope._mapId)
        }
    }

    def _assignId() {
        var block = currentStack.pop()
        if (block.segmentId > 0) {  // already visited
            return
        }

        block.segmentId = nextSegmentId
        nextSegmentId += 1

        var candidates = new Vector[CodeBlock]()
        block._ops.lastElement() match {
            case g: Goto => candidates.add(g._targetBlock)
            case i: IfBaseOp => {
                candidates.add(i._ifBranch)

                // else branch should be an implicit goto and must be next
                // to the current block
                if (i._elseBranch._parentScope != currentScope) {
                    throw new Exception("unexpected")
                }
                if (i._elseBranch._ops.size() != 1) {
                    throw new Exception("unexpected")
                }
                i._elseBranch._ops.lastElement() match {
                    case g: Goto => {
                        i._elseBranch.segmentId = nextSegmentId
                        nextSegmentId += 1
                        candidates.add(g._targetBlock)
                    }
                    case _ => throw new Exception("unexpected")
                }
            }
            case s: Switch => {
                for (block <- s._table.values()) {
                    candidates.add(block)
                }
                candidates.add(s._defaultBranch)
            }
            case _ => {}
        }

        var candidateScope: CodeSection = null
        for (block <- candidates) {
            if (block.segmentId < 0) {
                val blockScope = block._parentScope

                stacksMap.get(blockScope._mapId).push(block)
                if (currentScope._contains(blockScope)) {
                    candidateScope = blockScope
                }
            }
        }

        if (candidateScope != null && currentScope != candidateScope) {
            currentScope = candidateScope
            currentStack = stacksMap.get(candidateScope._mapId)
            scopeStack.push(currentScope)
        }
    }
}

object CodeSection {
    def reconstructFlowGraph(
            owner: AttributeOwner,
            exceptions: Vector[ExceptionEntry],
            ops: Vector[Operation]): CodeSection = {

        _checkExceptionOverlaps(exceptions)

        val jumpTargets = _collectJumpTargets(exceptions, ops)

        var result = new CodeSection(owner, null)
        result.pc = 0
        result._endPc = ops.lastElement().pc + 5 // fake it

        _createExceptionSubsections(exceptions, result)

        var pcBlockMap = new TreeMap[Int, CodeBlock]()

        var prevBlock: CodeBlock = null
        var currBlock: CodeBlock = null
        for (op <- ops) {
            if (jumpTargets.contains(op.pc)) {
                if (prevBlock != null) {
                    prevBlock.implicitGoto = currBlock
                    prevBlock._endPc = currBlock.pc
                }
                prevBlock = currBlock

                var section = result._getMostSpecificSection(op.pc)
                currBlock = section.newBlock()
                currBlock.pc = op.pc
                pcBlockMap.put(op.pc, currBlock)

                if (op.pc == 0) {
                    currBlock.isEntryPoint = true
                }
            }
            currBlock._add(op)
        }
        if (currBlock != null) {
            if (prevBlock != null) {
                prevBlock.implicitGoto = currBlock
                prevBlock._endPc = currBlock.pc
            }
            currBlock._endPc = currBlock.pc + 5  // fake it
        }

        result._fixPcs()

        for (op <- ops) {
            op.bindBlockRefs(pcBlockMap)
        }

        return result
    }

    // 1. ensure there are no partial overlaps
    // 2. ensure more specific scope is before less specific scope
    def _checkExceptionOverlaps(exceptions: Vector[ExceptionEntry]) {
        // TODO
    }

    def _createExceptionSubsections(
            exceptions: Vector[ExceptionEntry],
            global: CodeSection) {
        // create try sections
        for (i <- (exceptions.size() - 1).to(0, -1)) {
            var entry = exceptions.elementAt(i)

            var section = global._getMostSpecificSection(entry.startPc)
            if (section.pc < entry.startPc || entry.endPc < section._endPc) {
                section = section.newSubSection()
                section.pc = entry.startPc
                section._endPc = entry.endPc
            }
            entry._tmpSection = section
        }

        // create catch sections.
        for (entry <- exceptions) {
            var section = global._getMostSpecificSection(entry.handlerPc)
            if (section.pc < entry.handlerPc) {
                section = section.newSubSection()
                section.pc = entry.handlerPc
                section._endPc = entry.handlerPc + 1  // fake it
            }
            entry._tmpSection.shareExceptionHandle(entry.className(), section)
        }
    }

    def _collectJumpTargets(
            exceptions: Vector[ExceptionEntry],
            ops: Vector[Operation]): TreeSet[Int] = {
        var jumpTargets = new TreeSet[Int]()

        for (ex <- exceptions) {
            jumpTargets.add(ex.startPc)
            jumpTargets.add(ex.endPc)
            jumpTargets.add(ex.handlerPc)
        }

        for (op <- ops) {
            op match {
                case _: CodeSegment => throw new Exception(
                        "cannot group non-basic ops")
                case _: Return => jumpTargets.add(op.pc + 1)
                case _: ReturnValue => jumpTargets.add(op.pc + 1)
                case _: Athrow => jumpTargets.add(op.pc + 1)
                case g: Goto => jumpTargets.add(g.pc + g._tmpOffset)
                case i: IfBaseOp => {
                    jumpTargets.add(i.pc + i._tmpOffset)  // if branch
                    jumpTargets.add(i.pc + 3)  // else branch
                }
                case s: Switch => {
                    jumpTargets.add(s.pc + s._tmpDefaultOffset)
                    for (offset <- s._tmpOffset.values()) {
                        jumpTargets.add(s.pc + offset)
                    }
                }
                case _ => {}
            }
        }

        jumpTargets.add(0)
        return jumpTargets
    }
}
