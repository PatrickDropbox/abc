IMPORTS = [
        'java.io.DataInputStream',
        'java.io.DataOutputStream',
        'java.util.Vector',
        'scala.collection.JavaConversions._'
        ]

# class name -> (owner type, comment, {name -> bit})
ACCESS_FLAGS = {
        'ClassAccessFlags': (
                'ClassInfo',
                'see table 4.1-A / page 71-72',
                {
                    'public':     0x0001,
                    'final':      0x0010,
                    'super':      0x0020,
                    'interface':  0x0200,
                    'abstract':   0x0400,
                    'synthetic':  0x1000,
                    'annotation': 0x2000,
                    'enum':       0x4000,
                }),
        'FieldAccessFlags': (
                'FieldInfo',
                'see table 4.5-A / page 90-91',
                {
                    'public':    0x0001,
                    'private':   0x0002,
                    'protected': 0x0004,
                    'static':    0x0008,
                    'final':     0x0010,
                    'volatile':  0x0040,
                    'transient': 0x0080,
                    'synthetic': 0x1000,
                    'enum':      0x4000,
                }),
        'MethodAccessFlags': (
                'MethodInfo',
                'see table 4.6-A / page 93-94',
                {
                    'public':       0x0001,
                    'private':      0x0002,
                    'protected':    0x0004,
                    'static':       0x0008,
                    'final':        0x0010,
                    'synchronized': 0x0020,
                    'bridge':       0x0040,
                    'varargs':      0x0080,
                    'native':       0x0100,
                    'abstract':     0x0400,
                    'strict':       0x0800,
                    'synthetic':    0x1000,
                }),
        'InnerClassAccessFlags': (
                'InnerClassEntry',
                'see table 4.7.6-A / page 114-118',
                {
                    'public':     0x0001,
                    'private':    0x0002,
                    'protected':  0x0004,
                    'static':     0x0008,
                    'final':      0x0010,
                    'interface':  0x0200,
                    'abstract':   0x0400,
                    'synthetic':  0x1000,
                    'annotation': 0x2000,
                    'enum':       0x4000,
                }),
        }


class Writer(object):
    def __init__(self):
        self.indent = 0
        self.content = []

    def push(self):
        self.indent += 1

    def pop(self):
        self.indent -= 1
        assert self.indent >= 0

    def write(self, s):
        line = '%s%s' % ('    ' * self.indent, s) if s else ''
        self.content.append(line)

    def body(self):
        return '\n'.join(self.content)

def main():
    w = Writer()
    push = w.push
    pop = w.pop
    write = w.write

    write('// AUTO-GENERATED BY gen_access_flags.py')
    IMPORTS.sort()
    for i in IMPORTS:
        write('import %s' % i)
    write('')
    write('')

    for class_name, (owner_type, comment, name_bits) in ACCESS_FLAGS.items():
        name_bits = name_bits.items()
        name_bits.sort(key=lambda x: x[1])

        write('// %s' % comment)
        write('class %s(o: %s) {' % (class_name, owner_type))
        push()

        # track owner
        write('')
        write('var _owner = o')
        write('')

        # fields
        for name, _ in name_bits:
            write('var is%s = false' % name.title())
        write('')

        # debug string
        write('def debugString(): String = {')
        push()

        write('var flagStrings = new Vector[String]()')
        write('')

        for name, bit in name_bits:
            write('if (is%s) {' % name.title())
            push()
            write('flagStrings.add("%s")' % name.upper())
            pop()
            write('}')

        write('')
        write('var result = ""')
        write('for (s <- flagStrings) {')
        push()
        write('if (result.equals("")) {')
        push()
        write('result = s')
        pop()
        write('} else {')
        push()
        write('result += ", " + s')
        pop()
        write('}')
        pop()
        write('}')
        write('')
        write('return result')

        pop()
        write('}')
        write('')

        # serialize
        write('def serialize(output: DataOutputStream) {')
        push()

        write('var result = 0')
        write('')

        for name, bit in name_bits:
            write('if (is%s) {' % name.title())
            push()
            write('result |= 0x%04x' % bit)
            pop()
            write('}')

        write('')
        write('output.writeShort(result)')

        pop()
        write('}')
        write('')

        # deserialize
        write('def deserialize(input: DataInputStream) {')
        push()
        write('val flags = input.readUnsignedShort()')
        write('')
        for name, bit in name_bits:
            write('is%s = ((flags & 0x%04x) != 0)' % (name.title(), bit))
        pop()
        write('}')

        pop()
        write('}')
        write('')

    f = open('access_flags.scala', 'w')
    f.write(w.body())
    f.close()


if __name__ == '__main__':
    main()
