import java.io.DataInputStream
import java.io.DataOutputStream
import java.util.HashMap
import java.util.TreeMap
import java.util.TreeSet

import scala.collection.JavaConversions._


class MethodPool(c: ClassInfo) {
    var _owner = c

    var _methods = new TreeMap[MethodSignature, MethodInfo]()

    // Secondary index.  Unlike c++, TreeMap/TreeSet lower bound does not
    // return an iterator, and generating the upper bound is hard. =(
    var _methodsByName = new HashMap[String, TreeSet[MethodInfo]]()

    def _add(method: MethodInfo) {
        val signature = method.signature()
        if (_methods.containsKey(signature)) {
            throw new Exception("Adding duplicate method: " + signature)
        }
        _methods.put(signature, method)

        var methodSet = _methodsByName.get(method.name())
        if (methodSet == null) {
            methodSet = new TreeSet[MethodInfo]()
            _methodsByName.put(method.name(), methodSet)
        }
        methodSet.add(method)
    }

    def add(name: String, methodType: MethodType): MethodInfo = {
        val method = new MethodInfo(
                _owner,
                _owner.constants().getUtf8(name),
                methodType)
        _add(method)
        return method
    }

    def get(signature: MethodSignature): MethodInfo = {
        val method = _methods.get(signature)
        if (method == null) {
            throw new Exception("method not found: " + signature)
        }
        return method
    }

    def serialize(output: DataOutputStream) {
        // TODO
    }

    def deserialize(input: DataInputStream, constants: ConstantPool) {
        // TODO
    }
}
