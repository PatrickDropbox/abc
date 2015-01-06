import scala.collection.JavaConversions._


class ShortenSimpleGotoChains extends CodeAnalysisPass {
    def analyze(root: CodeScope) {
        new SimpleGotoChainShortener(root).apply()
    }
}

class SimpleGotoChainShortener(root: CodeScope) extends CodeVisitor(root) {
    override def visitBlock(block: CodeBlock) {
        block._ops.lastElement() match {
            case g: Goto => {
                g._targetBlock = shorten(g._targetBlock)
            }
            case i: IfBaseOp => {
                i._ifBranch = shorten(i._ifBranch)
            }
            case s: Switch => {
                s._defaultBranch = shorten(s._defaultBranch)
                for (c <- s._table.entrySet()) {
                    c.setValue(shorten(c.getValue()))
                }
            }
            case _ => {}
        }
    }

    def shorten(block: CodeBlock): CodeBlock = {
        var target = block
        while (true) {
            var newTarget = followSimpleGoto(target)
            if (newTarget == null) {
                return target
            } else {
                target = newTarget
            }
        }
        return null
    }

    def followSimpleGoto(block: CodeBlock): CodeBlock = {
        if (block._ops.size() != 1) {
            return null
        }

        block._ops.lastElement() match {
            case g: Goto => return g._targetBlock
            case _ => return null
        }
    }
}

