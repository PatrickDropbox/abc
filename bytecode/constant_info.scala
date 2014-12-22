
object ConstInfo {
    // see table 4.4-A (page 79)
    val CLASS = 7
    val FIELD_REF = 9
    val METHOD_REF = 10
    val INTERFACE_METHOD_REF = 11
    val STRING = 8
    val INTEGER = 3
    val FLOAT = 4
    val LONG = 5
    val DOUBLE = 6
    val NAME_AND_TYPE = 12
    val UTF8 = 1
    val METHOD_HANDLE = 15
    val METHOD_TYPE = 16
    val INVOKE_DYNAMIC = 18
}

trait ConstInfo {
    var index = 0

    def tag(): Int
}

class ConstClassInfo extends ConstInfo {
    def tag(): Int = ConstInfo.CLASS
}
