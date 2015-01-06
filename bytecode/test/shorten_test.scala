object ShortenTest {
    def main(arg: Array[String]) {
        var c = new ClassInfo("ClassName")
        var m = c._methods.add("func", new MethodType())
        var a = m._attributes
        a._code = new CodeAttribute(m)

        var code = new CodeScope(m, null, 0)
        a._code.code = code

        var b1 = code.newBlock()
        var s2 = code.newSubSection()
        var s3 = code.newSubSection()

        var b21 = s2.newBlock()
        var s22 = s2.newSubSection()
        var b23 = s2.newBlock()

        var b221 = s22.newBlock()

        var b31 = s3.newBlock()
        var b32 = s3.newBlock()
        var s33 = s3.newSubSection()

        var b331 = s33.newBlock()
        var s332 = s33.newSubSection()

        var b3321 = s332.newBlock()
        b3321.returnVoid()

        c.analyze()

        PcAssigner.assignSegmentIdsAndPcs(code)
        println(code.debugString(""))
    }
}
