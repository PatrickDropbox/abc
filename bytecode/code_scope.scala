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


class ExceptionTarget(e: ConstClassInfo, t: CodeScope) {
    var exception: ConstClassInfo = e
    var target: CodeScope = t
}

// a group of code segments which will be written out as a single continuous
// unit.
//
// Also acts as lexical / exception try scope
//
// NOTE: scope's implicit goto is only used by the last segment. All other
// segments will implicitly goto the next segment in the scope (useless
// implicitGoto is explicitly set)
class CodeScope(
        owner: AttributeOwner,
        parent: CodeScope) extends CodeSegment(owner, parent) {
    var _mapId = -1  // unique unordered id

    def this(parent: CodeScope) = this(parent._owner, parent)

    var _segments = new Vector[CodeSegment]()
    var _blocks = new Vector[CodeBlock]()
    var _subsections = new Vector[CodeScope]()

    var _exceptionTargets = new Vector[ExceptionTarget]()

    var _entryPoint: CodeSegment = null

    def getEntryBlock(): CodeBlock = {
        if (_entryPoint == null) {
            return newBlock()
        }

        _entryPoint match {
            case b: CodeBlock => return b
            case s: CodeScope => return s.getEntryBlock()
        }
    }

    def newBlock(): CodeBlock = {
        var block = new CodeBlock(_owner)
        return _addBlock(block)
    }

    def _addSegment(seg: CodeSegment) {
        if (_entryPoint == null) {
            _entryPoint = seg
        } else {
            var last = _segments.lastElement()
            if (last.implicitGoto == null) {
                last.implicitGoto = seg
            }
        }

        _segments.add(seg)
    }

    def _addBlock(block: CodeBlock): CodeBlock = {
        block._parentScope = this
        _addSegment(block)
        _blocks.add(block)
        return block
    }

    def newSubSection(): CodeScope = {
        val section = new CodeScope(this)
        _addSegment(section)
        _subsections.add(section)
        return section
    }

    // only used for reconstruction
    def _getMostSpecificSection(pc: Int): CodeScope = {
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
    def newExceptionHandle(exceptionClassName: String): CodeScope = {
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
    def shareExceptionHandle(exceptionClassName: String, target: CodeScope) {
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

    def generateLineNumberTable(): LineNumberTableAttribute = {
        var table = new TreeMap[Int, Int]()

        var blocks = new Vector[CodeBlock]()
        _collectBlocks(blocks)
        Collections.sort(blocks)

        var currLine = -1
        for (block <- blocks) {
            for (op <- block._ops) {
                if (op.line >= 0 && op.line != currLine) {
                    table.put(op.pc, op.line)
                    currLine = op.line
                }
            }
        }

        if (table.isEmpty()) {
            return null
        }

        var attr = new LineNumberTableAttribute(_owner)
        attr.table = table
        return attr
    }

    // this assumes pc are assigned and segments are sorted
    def _collectExceptionEntries(result: Vector[ExceptionEntry]) {
        for (section <- _subsections) {
            section._collectExceptionEntries(result)
        }

        for (entry <- _exceptionTargets) {
            result.add(new ExceptionEntry(
                    _owner,
                    pc,
                    _endPc,
                    entry.target.getEntryBlock().pc,
                    entry.exception))
        }
    }

    // returns true if this equals, or contains, the other section
    def _contains(other: CodeScope): Boolean = {
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

    // only for deserialization (assume scope is sorted)
    def _fixEntryPoints() {
        _entryPoint = _segments.elementAt(0)

        for (section <- _subsections) {
            section._fixEntryPoints()
        }
    }

    def _collectBlocks(result: Vector[CodeBlock]) {
        for (seg <- _segments) {
            seg match {
                case b: CodeBlock => result.add(b)
                case s: CodeScope => s._collectBlocks(result)
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
            mapping: HashMap[Int, CodeScope]): Int = {
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

        if (_endPc > Const.UINT16_MAX) {
            // code attributes can't handle code length > 64K ...
            throw new Exception("code too large")
        }

        output.writeInt(_endPc)
        for (block <- blocks) {
            block.serialize(output)
        }

        var exceptions = new Vector[ExceptionEntry]()
        _collectExceptionEntries(exceptions)

        output.writeShort(exceptions.size())
        for (entry <- exceptions) {
            entry.serialize(output)
        }
    }

    def deserialize(startAddress: Int, opCode: Int, input: DataInputStream) {
        throw new Exception("cannot directly deserialize code section")
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

