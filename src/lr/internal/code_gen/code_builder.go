package code_gen

import (
	"fmt"
	"io"
	"path/filepath"
	"regexp"
)

var (
	nameRe = regexp.MustCompile(`((?:(?:\[\])*\**)*)(?:(.+)\.)?(\w+(?:{})?)$`)
)

type line struct {
	indent      string
	indentLevel int
	format      string
	args        []interface{}
}

func (l *line) WriteTo(writer io.Writer) (int64, error) {
	total := int64(0)

	if l.format != "" {
		for i := 0; i < l.indentLevel; i++ {
			n, err := writer.Write([]byte(l.indent))
			total += int64(n)

			if err != nil {
				return total, err
			}
		}

		n, err := fmt.Fprintf(writer, l.format, l.args...)
		total += int64(n)

		if err != nil {
			return total, err
		}
	}

	n, err := writer.Write([]byte{'\n'})
	total += int64(n)
	if err != nil {
		return total, err
	}

	return total, nil
}

type CodeBuilder struct {
	chunks []io.WriterTo

	indent      string
	indentLevel int
}

func NewCodeBuilder() *CodeBuilder {
	return &CodeBuilder{
		chunks:      nil,
		indent:      "    ",
		indentLevel: 0,
	}
}

func (cb *CodeBuilder) PushIndent() {
	cb.indentLevel += 1
}

func (cb *CodeBuilder) PopIndent() {
	if cb.indentLevel > 0 {
		cb.indentLevel -= 1
	}
}

func (cb *CodeBuilder) Line(format string, args ...interface{}) {
	cb.chunks = append(
		cb.chunks,
		&line{cb.indent, cb.indentLevel, format, args})
}

func (cb *CodeBuilder) Embed(template io.WriterTo) {
	cb.chunks = append(cb.chunks, template)
}

func (cb *CodeBuilder) WriteTo(output io.Writer) (int64, error) {
	total := int64(0)
	for _, chunk := range cb.chunks {
		n, err := chunk.WriteTo(output)
		total += n

		if err != nil {
			return total, err
		}
	}

	return total, nil
}

type importEntry struct {
	path  string
	alias string
}

type importObject struct {
	*importEntry

	prefix string // * or &
	name   string
}

func (obj *importObject) String() string {
	if obj.alias == "" {
		return obj.prefix + obj.name
	}

	return obj.prefix + obj.alias + "." + obj.name
}

type goHeader struct {
	HeaderBoilerplate *CodeBuilder
	pkg               string
	imports           map[string]*importEntry
}

func newGoHeader(pkg string) *goHeader {
	return &goHeader{
		HeaderBoilerplate: NewCodeBuilder(),
		pkg:               pkg,
		imports:           map[string]*importEntry{},
	}
}

// This supports accessing objects of the form:
//     *(\[\])*(\*)*)*(<full module path>\.)?<object>({})?
// map objects are not supported
func (header *goHeader) Obj(fullName string) *importObject {
	match := nameRe.FindStringSubmatch(fullName)
	if match == nil {
		panic("Invalid fullName: " + fullName)
	}

	prefix := match[1]
	modulePath := match[2]
	name := match[3]

	if modulePath == "" {
		return &importObject{&importEntry{}, prefix, name}
	}

	entry, ok := header.imports[modulePath]
	if !ok {
		entry = &importEntry{
			path: modulePath,
		}

		header.imports[modulePath] = entry
	}

	return &importObject{entry, prefix, name}
}

func (header *goHeader) assignAlias() error {
	aliasCount := map[string]int{}
	for pkg, entry := range header.imports {
		base := filepath.Base(pkg)
		if base == "." || base == "/" {
			return fmt.Errorf("Invalid import path: %s", pkg)
		}

		alias := base
		aliasCount[alias] += 1
		if aliasCount[alias] > 1 {
			alias = fmt.Sprintf("%s%d", alias, aliasCount[alias])
		}

		entry.alias = alias
	}

	return nil
}

func (header *goHeader) WriteTo(output io.Writer) (int64, error) {
	err := header.assignAlias()
	if err != nil {
		return 0, err
	}

	numWritten, err := header.HeaderBoilerplate.WriteTo(output)
	if err != nil {
		return numWritten, err
	}

	builder := NewCodeBuilder()
	builder.Line("package %s", header.pkg)
	builder.Line("")

	if len(header.imports) > 0 {
		builder.Line("import (")
		builder.PushIndent()
		// Maybe separate built-in packages from other packages
		for _, entry := range header.imports {
			builder.Line("%s \"%s\"", entry.alias, entry.path)
		}
		builder.PopIndent()
		builder.Line(")")
		builder.Line("")
	}

	numWritten2, err := builder.WriteTo(output)
	return numWritten + numWritten2, err
}
