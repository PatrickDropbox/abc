class ShortenSimpleGotoChains extends CodeAnalysisPass {
    def analyze(root: CodeScope) {
        new GotoChainShortener(root).apply()
    }
}

class SimpleGotoChainShortener(root: CodeScope) extends CodeVisitor(root) {
    override def visitBlock(block: CodeBlock) {
        block._ops.lastElement() match {
            case g: Goto => {
                var target = followSimpleGoto(g._targetBlock)
                while (target != null) {
                    g._targetBlock = target
                    target = followSimpleGoto(target)
                }
            }
            case _ => {}
        }
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

