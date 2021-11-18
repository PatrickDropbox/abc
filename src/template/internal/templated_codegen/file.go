package templated_codegen

import (
	_fmt "fmt"
	_io "io"

	"github.com/pattyshack/abc/src/template/internal"
)

type File struct {
	spec *template.File
}

func (File) Name() string { return "File" }

func (template *File) writeValue(
	output _io.Writer,
	value interface{},
	loc string) (
	int,
	error) {

	var valueBytes []byte
	switch val := value.(type) {
	case _fmt.Stringer:
		valueBytes = []byte(val.String())
	case string:
		valueBytes = []byte(val)
	case []byte:
		valueBytes = val
	case bool:
		valueBytes = []byte(_fmt.Sprintf("%v", val))
	case uint:
		valueBytes = []byte(_fmt.Sprintf("%v", val))
	case uint8:
		valueBytes = []byte(_fmt.Sprintf("%v", val))
	case uint16:
		valueBytes = []byte(_fmt.Sprintf("%v", val))
	case uint32:
		valueBytes = []byte(_fmt.Sprintf("%v", val))
	case uint64:
		valueBytes = []byte(_fmt.Sprintf("%v", val))
	case int:
		valueBytes = []byte(_fmt.Sprintf("%v", val))
	case int8:
		valueBytes = []byte(_fmt.Sprintf("%v", val))
	case int16:
		valueBytes = []byte(_fmt.Sprintf("%v", val))
	case int32:
		valueBytes = []byte(_fmt.Sprintf("%v", val))
	case int64:
		valueBytes = []byte(_fmt.Sprintf("%v", val))
	case float32:
		valueBytes = []byte(_fmt.Sprintf("%v", val))
	case float64:
		valueBytes = []byte(_fmt.Sprintf("%v", val))
	case complex64:
		valueBytes = []byte(_fmt.Sprintf("%v", val))
	case complex128:
		valueBytes = []byte(_fmt.Sprintf("%v", val))
	default:
		return 0, _fmt.Errorf(
			"Unsupported output value type (%s): %v",
			loc,
			value)
	}

	return output.Write(valueBytes)
}

func (_template *File) WriteTo(
	_output _io.Writer) (
	int64,
	error) {

	_numWritten := int64(0)

	spec := _template.spec

	// file.template:12:0
	{
		_n, _err := _output.Write([]byte(`package `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:12:8
	{
		_n, _err := _template.writeValue(
			_output,
			(spec.PackageName),
			"file.template:12:8")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:12:27
	{
		_n, _err := _output.Write([]byte(`

import (
	_fmt "fmt"
	_io "io"`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:17:0
	if spec.Imports != "" {
		// file.template:17:26
		{
			_n, _err := _output.Write([]byte(`

`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:19:0
		{
			_n, _err := _template.writeValue(
				_output,
				(spec.Imports),
				"file.template:19:0")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// file.template:20:8
	{
		_n, _err := _output.Write([]byte(`
)

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:23:5
	{
		_n, _err := _template.writeValue(
			_output,
			(spec.TemplateName),
			"file.template:23:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:23:25
	{
		_n, _err := _output.Write([]byte(` struct {`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:24:0
	for _, arg := range spec.Arguments {
		// file.template:24:39
		{
			_n, _err := _output.Write([]byte(`
	`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:25:1
		{
			_n, _err := _template.writeValue(
				_output,
				(arg.Name),
				"file.template:25:1")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:25:12
		{
			_n, _err := _output.Write([]byte(` `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:25:13
		{
			_n, _err := _template.writeValue(
				_output,
				(arg.Type),
				"file.template:25:13")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// file.template:26:8
	{
		_n, _err := _output.Write([]byte(`
}

func (`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:29:6
	{
		_n, _err := _template.writeValue(
			_output,
			(spec.TemplateName),
			"file.template:29:6")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:29:26
	{
		_n, _err := _output.Write([]byte(`) Name() string { return "`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:29:52
	{
		_n, _err := _template.writeValue(
			_output,
			(spec.TemplateName),
			"file.template:29:52")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:29:72
	{
		_n, _err := _output.Write([]byte(`" }

func (template *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:31:16
	{
		_n, _err := _template.writeValue(
			_output,
			(spec.TemplateName),
			"file.template:31:16")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:31:36
	{
		_n, _err := _output.Write([]byte(`) writeValue(
	output _io.Writer,
	value interface{},
	loc string) (
	int,
	error) {

	var valueBytes []byte
	switch val := value.(type) {
	case _fmt.Stringer:
		valueBytes = []byte(val.String())
	case string:
		valueBytes = []byte(val)
	case []byte:
		valueBytes = val`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:46:0
	for _, primitive := range template.OutputablePrimitiveTypes {
		// file.template:46:64
		{
			_n, _err := _output.Write([]byte(`
	case `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:47:6
		{
			_n, _err := _template.writeValue(
				_output,
				(primitive),
				"file.template:47:6")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:47:16
		{
			_n, _err := _output.Write([]byte(`:
		valueBytes = []byte(_fmt.Sprintf("%v", val))`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// file.template:49:9
	{
		_n, _err := _output.Write([]byte(`
	default:
		return 0, _fmt.Errorf(
			"Unsupported output value type (%s): %v",
			loc,
			value)
	}

	return output.Write(valueBytes)
}

func (_template *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:61:17
	{
		_n, _err := _template.writeValue(
			_output,
			(spec.TemplateName),
			"file.template:61:17")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:61:37
	{
		_n, _err := _output.Write([]byte(`) WriteTo(
	_output _io.Writer) (
	int64,
	error) {

	_numWritten := int64(0)
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:68:0
	for idx, arg := range spec.Arguments {
		// file.template:69:4
		if idx == 0 {
			// file.template:69:20
			{
				_n, _err := _output.Write([]byte(`
`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
		// file.template:71:13
		{
			_n, _err := _output.Write([]byte(`	`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:72:1
		{
			_n, _err := _template.writeValue(
				_output,
				(arg.Name),
				"file.template:72:1")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:72:12
		{
			_n, _err := _output.Write([]byte(` := _template.`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:72:26
		{
			_n, _err := _template.writeValue(
				_output,
				(arg.Name),
				"file.template:72:26")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:72:37
		{
			_n, _err := _output.Write([]byte(`
`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// file.template:73:7
	{
		_n, _err := _output.Write([]byte(`
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:75:0
	{
		_n, _err := (&Body{"\t", spec.Body}).WriteTo(_output)
		_numWritten += _n
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:75:34
	{
		_n, _err := _output.Write([]byte(`
	return _numWritten, nil
}
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}

	return _numWritten, nil
}
