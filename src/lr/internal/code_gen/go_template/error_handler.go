// Auto-generated from source: error_handler.template

package go_template

import (
	_fmt "fmt"
	_io "io"
)

type ErrorHandler struct {
	ErrHandler        string
	TokenType         string
	StackType         string
	DefaultErrHandler string
	Errorf            interface{}
	ExpectedTerminals string
}

func (ErrorHandler) Name() string { return "ErrorHandler" }

func (template *ErrorHandler) writeValue(
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

func (_template *ErrorHandler) WriteTo(
	_output _io.Writer) (
	int64,
	error) {

	_numWritten := int64(0)

	ErrHandler := _template.ErrHandler
	TokenType := _template.TokenType
	StackType := _template.StackType
	DefaultErrHandler := _template.DefaultErrHandler
	Errorf := _template.Errorf
	ExpectedTerminals := _template.ExpectedTerminals

	// error_handler.template:16:0
	{
		_n, _err := _output.Write([]byte(`type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// error_handler.template:16:5
	{
		_n, _err := _template.writeValue(
			_output,
			(ErrHandler),
			"error_handler.template:16:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// error_handler.template:16:16
	{
		_n, _err := _output.Write([]byte(` interface {
    Error(nextToken `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// error_handler.template:17:20
	{
		_n, _err := _template.writeValue(
			_output,
			(TokenType),
			"error_handler.template:17:20")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// error_handler.template:17:30
	{
		_n, _err := _output.Write([]byte(`, parseStack `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// error_handler.template:17:43
	{
		_n, _err := _template.writeValue(
			_output,
			(StackType),
			"error_handler.template:17:43")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// error_handler.template:17:53
	{
		_n, _err := _output.Write([]byte(`) error
}

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// error_handler.template:20:5
	{
		_n, _err := _template.writeValue(
			_output,
			(DefaultErrHandler),
			"error_handler.template:20:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// error_handler.template:20:23
	{
		_n, _err := _output.Write([]byte(` struct {}

func (`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// error_handler.template:22:6
	{
		_n, _err := _template.writeValue(
			_output,
			(DefaultErrHandler),
			"error_handler.template:22:6")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// error_handler.template:22:24
	{
		_n, _err := _output.Write([]byte(`) Error(nextToken `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// error_handler.template:22:42
	{
		_n, _err := _template.writeValue(
			_output,
			(TokenType),
			"error_handler.template:22:42")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// error_handler.template:22:52
	{
		_n, _err := _output.Write([]byte(`, stack `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// error_handler.template:22:60
	{
		_n, _err := _template.writeValue(
			_output,
			(StackType),
			"error_handler.template:22:60")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// error_handler.template:22:70
	{
		_n, _err := _output.Write([]byte(`) error {
    return `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// error_handler.template:23:11
	{
		_n, _err := _template.writeValue(
			_output,
			(Errorf),
			"error_handler.template:23:11")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// error_handler.template:23:18
	{
		_n, _err := _output.Write([]byte(`(
        "Syntax error: unexpected symbol %v. Expecting %v (%v)",
        nextToken.Id(),
        `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// error_handler.template:26:8
	{
		_n, _err := _template.writeValue(
			_output,
			(ExpectedTerminals),
			"error_handler.template:26:8")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// error_handler.template:26:26
	{
		_n, _err := _output.Write([]byte(`[stack[len(stack)-1].StateId],
        nextToken.Loc())
}
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}

	return _numWritten, nil
}
