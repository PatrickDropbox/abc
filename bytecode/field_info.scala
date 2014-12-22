import java.io.DataInputStream
import java.io.DataOutputStream


class FieldInfo {
    var access = new FieldAccessFlags()
    var name: ConstUtf8Info = null
    var descriptor: ConstUtf8Info = null
    var attributes = new FieldAttributes()
}
