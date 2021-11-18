package templated_codegen

import (
	fmt "fmt"
	io "io"
)

import (
	"github.com/pattyshack/abc/src/template/internal"
)

type File struct {
	spec *template.File
}

func (File) Name() string { return "File" }

func (template *File) writeValue(
	output io.Writer, value interface{}, loc string) (int, error) {
	var valueBytes []byte
	switch val := value.(type) {
	case fmt.Stringer:
		valueBytes = []byte(val.String())
	case string:
		valueBytes = []byte(val)
	case []byte:
		valueBytes = val
	case bool:
		valueBytes = []byte(fmt.Sprintf("%v", val))
	case uint:
		valueBytes = []byte(fmt.Sprintf("%v", val))
	case uint8:
		valueBytes = []byte(fmt.Sprintf("%v", val))
	case uint16:
		valueBytes = []byte(fmt.Sprintf("%v", val))
	case uint32:
		valueBytes = []byte(fmt.Sprintf("%v", val))
	case uint64:
		valueBytes = []byte(fmt.Sprintf("%v", val))
	case int:
		valueBytes = []byte(fmt.Sprintf("%v", val))
	case int8:
		valueBytes = []byte(fmt.Sprintf("%v", val))
	case int16:
		valueBytes = []byte(fmt.Sprintf("%v", val))
	case int32:
		valueBytes = []byte(fmt.Sprintf("%v", val))
	case int64:
		valueBytes = []byte(fmt.Sprintf("%v", val))
	case float32:
		valueBytes = []byte(fmt.Sprintf("%v", val))
	case float64:
		valueBytes = []byte(fmt.Sprintf("%v", val))
	case complex64:
		valueBytes = []byte(fmt.Sprintf("%v", val))
	case complex128:
		valueBytes = []byte(fmt.Sprintf("%v", val))
	default:
		return 0, fmt.Errorf("Unsupported output value type (%s): %v", loc, value)
	}

	return output.Write(valueBytes)
}

func (_template *File) WriteTo(_output io.Writer) (int64, error) {
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
	_io "io"
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}

	// file.template:18:0
	if spec.Imports != "" {

		// file.template:18:26
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

	// file.template:20:9
	{
		_n, _err := _output.Write([]byte(`
)

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}

	// file.template:24:5
	{
		_n, _err := _template.writeValue(
			_output,
			(spec.TemplateName),
			"file.template:24:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}

	// file.template:24:25
	{
		_n, _err := _output.Write([]byte(` struct {
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}

	// file.template:26:0
	for _, arg := range spec.Arguments {

		// file.template:26:39
		{
			_n, _err := _output.Write([]byte(`
	`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}

		// file.template:27:1
		{
			_n, _err := _template.writeValue(
				_output,
				(arg.Name),
				"file.template:27:1")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}

		// file.template:27:12
		{
			_n, _err := _output.Write([]byte(` `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}

		// file.template:27:13
		{
			_n, _err := _template.writeValue(
				_output,
				(arg.Type),
				"file.template:27:13")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}

	}

	// file.template:28:9
	{
		_n, _err := _output.Write([]byte(`
}

func (`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}

	// file.template:32:6
	{
		_n, _err := _template.writeValue(
			_output,
			(spec),
			"file.template:32:6")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}

	// file.template:32:11
	{
		_n, _err := _output.Write([]byte(`.TemplateName) Name() string { return "`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}

	// file.template:32:50
	{
		_n, _err := _template.writeValue(
			_output,
			(spec.TemplateName),
			"file.template:32:50")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}

	// file.template:32:70
	{
		_n, _err := _output.Write([]byte(`" }

func (template *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}

	// file.template:34:16
	{
		_n, _err := _template.writeValue(
			_output,
			(spec.TemplateName),
			"file.template:34:16")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}

	// file.template:34:36
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
		valueBytes = val
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}

	// file.template:50:0
	for _, primitive := range template.OutputablePrimitiveTypes {

		// file.template:50:65
		{
			_n, _err := _output.Write([]byte(`
	case `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}

		// file.template:52:6
		{
			_n, _err := _template.writeValue(
				_output,
				(primitive),
				"file.template:52:6")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}

		// file.template:52:16
		{
			_n, _err := _output.Write([]byte(`:
		valueBytes = []byte(_fmt.Sprintf("%v", val))
`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}

	// file.template:55:9
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

	// file.template:67:17
	{
		_n, _err := _template.writeValue(
			_output,
			(spec.TemplateName),
			"file.template:67:17")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}

	// file.template:67:37
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

	// file.template:74:0
	for _, arg := range spec.Arguments {

		// file.template:74:40
		{
			_n, _err := _output.Write([]byte(`
	`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}

		// file.template:76:1
		{
			_n, _err := _template.writeValue(
				_output,
				(arg.Name),
				"file.template:76:1")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}

		// file.template:76:12
		{
			_n, _err := _output.Write([]byte(` := _template.`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}

		// file.template:76:26
		{
			_n, _err := _template.writeValue(
				_output,
				(arg.Name),
				"file.template:76:26")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}

		// file.template:76:37
		{
			_n, _err := _output.Write([]byte(`
`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}

	// file.template:78:8
	{
		_n, _err := _output.Write([]byte(`
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}

	// file.template:80:0
	{
		_n, _err := (&Body{"\t", spec.Body}).WriteTo(_output)
		_numWritten += _n
		if _err != nil {
			return _numWritten, _err
		}
	}

	// file.template:80:33
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
