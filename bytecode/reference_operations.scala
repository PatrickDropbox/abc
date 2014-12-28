import java.io.DataInputStream
import java.io.DataOutputStream


// stack: ... -> ..., obj ref
class New(owner: MethodInfo, className: String)
        extends ShortOperandOp(owner, OpCode.NEW, "new", false, 0) {
    def this(owner: MethodInfo) = this(owner, null)

    var _constClass: ConstClassInfo = null
    if (className != null) {
        _constClass = _owner.constants().getClass(className)
    }

    override def serialize(output: DataOutputStream) {
        operand = _constClass.index
        super.serialize(output)
    }

    override def deserialize(opCode: Int, input: DataInputStream) {
        super.deserialize(opCode, input)
        _constClass = _owner.constants().getClassByIndex(operand)
    }

    override def debugString(): String = {
        var name = "???"
        if (_constClass != null) {
            name = _constClass.className()
        }
        return "new " + name
    }
}
