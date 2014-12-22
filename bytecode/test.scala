class Test {
    var intValue: Int = 1
    var longValue: Long = 2
    var floatValue: Float = 3
    var doubleValue: Double = 4
    var stringValue: String = "Hello world"

    def f(): Int = intValue

    def g(): Long = f().toLong

    def h(): Double = g().toDouble

    def h2(): Double = g().toDouble + h()
}
