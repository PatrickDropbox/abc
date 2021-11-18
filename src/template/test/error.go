package main

import (
	fmt "fmt"
	io "io"
)

type ErrorTemplate struct {
}

func (ErrorTemplate) Name() string { return "ErrorTemplate" }

func (template *ErrorTemplate) writeValue(
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

func (_template *ErrorTemplate) WriteTo(_output io.Writer) (int64, error) {
	_numWritten := int64(0)

	// error.template:5:0
	if true {

		// error.template:6:4
		{
			_err := (fmt.Errorf("bad"))
			if _err == nil {
				_err = fmt.Errorf("Unexpected error (error.template:6:4)")
			}
			return _numWritten, _err
		}

	}

	return _numWritten, nil
}
