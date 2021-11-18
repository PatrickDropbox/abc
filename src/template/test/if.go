package main

import (
	fmt "fmt"
	io "io"
)

type IfTemplate struct {
}

func (IfTemplate) Name() string { return "IfTemplate" }

func (template *IfTemplate) writeValue(
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

func (_template *IfTemplate) WriteTo(_output io.Writer) (int64, error) {
	_numWritten := int64(0)

	// if.template:6:0
	{
		_n, _err := _output.Write([]byte(`
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}

	// if.template:7:0

	val := 3

	pred := func() bool { return false }

	// if.template:11:2
	{
		_n, _err := _output.Write([]byte(`

If Without Else If / Else 1
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}

	// if.template:14:0
	if val*7 == 21 {

		// if.template:14:20
		{
			_n, _err := _output.Write([]byte(`val * 7 = 21`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}

	// if.template:14:39
	{
		_n, _err := _output.Write([]byte(`

If Without Else If / Else 2
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}

	// if.template:17:0
	if val == 3 && pred() {

		// if.template:17:27
		{
			_n, _err := _output.Write([]byte(`sadness`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}

	// if.template:17:41
	{
		_n, _err := _output.Write([]byte(`

With Else Branch
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}

	// if.template:20:0
	if val != 3 {

		// if.template:20:15
		{
			_n, _err := _output.Write([]byte(`val not equal 3`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	} else {

		// if.template:20:38
		{
			_n, _err := _output.Write([]byte(`val equals 3`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}

	// if.template:20:57
	{
		_n, _err := _output.Write([]byte(`

With Else If Branches 1
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}

	// if.template:23:0
	if val == 0 {

		// if.template:23:15
		{
			_n, _err := _output.Write([]byte(`
0
`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	} else if val == 1 {

		// if.template:25:20
		{
			_n, _err := _output.Write([]byte(`
1
`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	} else if val == 2 {

		// if.template:27:20
		{
			_n, _err := _output.Write([]byte(`
2
`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}

	// if.template:29:7
	{
		_n, _err := _output.Write([]byte(`

With Else If Branches 2
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}

	// if.template:32:0
	if val == 0 {

		// if.template:32:15
		{
			_n, _err := _output.Write([]byte(`
0
`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	} else if val == 1 {

		// if.template:34:20
		{
			_n, _err := _output.Write([]byte(`
1
`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	} else if val == 2 {

		// if.template:36:20
		{
			_n, _err := _output.Write([]byte(`
2
`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	} else if val == 3 {

		// if.template:38:20
		{
			_n, _err := _output.Write([]byte(`
3
`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}

	// if.template:40:7
	{
		_n, _err := _output.Write([]byte(`

With Else If And Else Branches
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}

	// if.template:43:0
	if val == 0 {

		// if.template:43:15
		{
			_n, _err := _output.Write([]byte(`
0
`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	} else if val == 1 {

		// if.template:45:20
		{
			_n, _err := _output.Write([]byte(`
1
`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	} else if val == 2 {

		// if.template:47:20
		{
			_n, _err := _output.Write([]byte(`
2
`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	} else {

		// if.template:49:8
		{
			_n, _err := _output.Write([]byte(`
other
`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}

	// if.template:51:7
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
