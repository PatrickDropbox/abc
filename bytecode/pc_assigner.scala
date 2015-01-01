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

