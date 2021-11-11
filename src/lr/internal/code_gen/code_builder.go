package code_gen

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
	"path/filepath"
	"regexp"
)

var (
	nameRe = regexp.MustCompile(`((?:(?:\[\])*\**)*)(?:(.+)\.)?(\w+(?:{})?)$`)
)

type Code interface {
	WriteTo(io.Writer) (int64, error)
}

type line struct {
	indentLevel int
	format      string
	args        []interface{}
}

func (l *line) writeTo(writer io.Writer, indent string) (int64, error) {
	total := int64(0)

	if l.format != "" {
		for i := 0; i < l.indentLevel; i++ {
			n, err := writer.Write([]byte(indent))
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
	lines []*line

	indent      string
	indentLevel int
}

func NewCodeBuilder() *CodeBuilder {
	return &CodeBuilder{
		lines:       nil,
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
	cb.lines = append(
		cb.lines,
		&line{cb.indentLevel, format, args})
}

func (cb *CodeBuilder) WriteTo(output io.Writer) (int64, error) {
	total := int64(0)
	for _, line := range cb.lines {
		n, err := line.writeTo(output, cb.indent)
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

type GoCodeBuilder struct {
	*CodeBuilder

	*goHeader
}

func NewGoCodeBuilder(pkg string) *GoCodeBuilder {
	return &GoCodeBuilder{
		NewCodeBuilder(),
		newGoHeader(pkg),
	}
}

func (cb *GoCodeBuilder) WriteTo(output io.Writer) (int64, error) {
	buffer := bytes.NewBuffer(nil)

	_, err := cb.goHeader.WriteTo(buffer)
	if err != nil {
		return 0, err
	}

	_, err = cb.CodeBuilder.WriteTo(buffer)
	if err != nil {
		return 0, err
	}

	formatted, err := format.Source(buffer.Bytes())
	if err != nil {
		return 0, fmt.Errorf(
			"Failed to format (%s) generated code:\n%s",
			err,
			buffer.Bytes())
	}

	n, err := output.Write(formatted)
	return int64(n), err
}
