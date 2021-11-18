package templated_codegen

import (
	fmt "fmt"
	io "io"
)

import (
	"github.com/pattyshack/abc/src/template/internal"
)

type For struct {
	ind  string
	stmt *template.For
}

func (For) Name() string { return "For" }

func (template *For) writeValue(
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

func (_template *For) WriteTo(_output io.Writer) (int64, error) {
	_numWritten := int64(0)

	ind := _template.ind
	stmt := _template.stmt

	// for.template:13:0
	{
		_n, _err := _template.writeValue(
			_output,
			(ind),
			"for.template:13:0")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}

	// for.template:13:4
	{
		_n, _err := _output.Write([]byte(`// `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}

	// for.template:13:7
	{
		_n, _err := _template.writeValue(
			_output,
			(stmt.Loc()),
			"for.template:13:7")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}

	// for.template:13:20
	{
		_n, _err := _output.Write([]byte(`
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}

	// for.template:14:0
	{
		_n, _err := _template.writeValue(
			_output,
			(ind),
			"for.template:14:0")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}

	// for.template:14:6
	{
		_n, _err := _output.Write([]byte(`for `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}

	// for.template:14:10
	{
		_n, _err := _template.writeValue(
			_output,
			(stmt.Predicate.Value),
			"for.template:14:10")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}

	// for.template:14:33
	{
		_n, _err := _output.Write([]byte(` {
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}

	// for.template:16:4
	{
		_n, _err := (&Body{ind + "\t", stmt.Body}).WriteTo(_output)
		_numWritten += _n
		if _err != nil {
			return _numWritten, _err
		}
	}

	// for.template:17:0
	{
		_n, _err := _template.writeValue(
			_output,
			(ind),
			"for.template:17:0")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}

	// for.template:17:4
	{
		_n, _err := _output.Write([]byte(`}
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}

	return _numWritten, nil
}
