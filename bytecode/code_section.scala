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

        pc = -1
        _endPc = -1
        segmentId = -1
        _mapId = -1
    }

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
        _insertImplicitGoto()

        var blocks = PcAssigner.assignSegmentIdsAndPcs(this)
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
