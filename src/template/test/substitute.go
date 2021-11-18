package main

import (
	fmt "fmt"
	io "io"
)

type SubstituteTemplate struct {
}

func (SubstituteTemplate) Name() string { return "SubstituteTemplate" }

func (template *SubstituteTemplate) writeValue(
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

func (_template *SubstituteTemplate) WriteTo(_output io.Writer) (int64, error) {
	_numWritten := int64(0)

	// substitute.template:7:0

	prefix := "I own "

	// substitute.template:9:2
	{
		_n, _err := _output.Write([]byte(`
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}

	// substitute.template:10:0
	{
		_n, _err := _template.writeValue(
			_output,
			(prefix),
			"substitute.template:10:0")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}

	// substitute.template:10:7
	{
		_n, _err := _output.Write([]byte(`$`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}

	// substitute.template:10:9
	{
		_n, _err := _template.writeValue(
			_output,
			(2 * (1 - 1)),
			"substitute.template:10:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}

	// substitute.template:10:19
	{
		_n, _err := _output.Write([]byte(`
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}

	return _numWritten, nil
}
