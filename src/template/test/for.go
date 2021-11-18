package main

import (
	fmt "fmt"
	io "io"
)

type ForTemplate struct {
	Chan  <-chan int
	Count int
}

func (ForTemplate) Name() string { return "ForTemplate" }

func (template *ForTemplate) writeValue(
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

func (_template *ForTemplate) WriteTo(_output io.Writer) (int64, error) {
	_numWritten := int64(0)

	Chan := _template.Chan
	Count := _template.Count

	// for.template:12:3
	{
		_n, _err := _output.Write([]byte(`
Infinite Loop:`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}

	// for.template:15:0
	for {

		// for.template:16:4
		Count += 1

		// for.template:17:4
		if Count%3 == 0 {

			// for.template:17:27
			{
				_n, _err := _output.Write([]byte(`
  `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}

			// for.template:19:2
			{
				_n, _err := _template.writeValue(
					_output,
					(Count),
					"for.template:19:2")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}

			// for.template:19:8
			{
				_n, _err := _output.Write([]byte(` % 3 == 0`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}

			// for.template:20:8
			continue

		} else if Count%3 == 1 {

			// for.template:21:32
			{
				_n, _err := _output.Write([]byte(`
  `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}

			// for.template:23:2
			{
				_n, _err := _template.writeValue(
					_output,
					(Count),
					"for.template:23:2")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}

			// for.template:23:8
			{
				_n, _err := _output.Write([]byte(` % 3 == 1`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		} else {

			// for.template:24:14
			{
				_n, _err := _output.Write([]byte(`
  `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}

			// for.template:26:2
			{
				_n, _err := _template.writeValue(
					_output,
					(Count),
					"for.template:26:2")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}

			// for.template:26:8
			{
				_n, _err := _output.Write([]byte(` % 3 == 2`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}

			// for.template:27:8
			if Count > 10 {

				// for.template:28:12
				break

			}

		}

	}

	// for.template:31:7
	{
		_n, _err := _output.Write([]byte(`
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}

	// for.template:35:3
	{
		_n, _err := _output.Write([]byte(`
Predicate Loop`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}

	// for.template:38:0
	for Count < 30 {

		// for.template:38:19
		{
			_n, _err := _output.Write([]byte(`
  `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}

		// for.template:39:2
		{
			_n, _err := _template.writeValue(
				_output,
				(Count),
				"for.template:39:2")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}

		// for.template:40:4

		Count = (Count +
			1) * 2

	}

	// for.template:44:8
	{
		_n, _err := _output.Write([]byte(`

`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}

	// for.template:48:2
	{
		_n, _err := _output.Write([]byte(`

Counter Loop`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}

	// for.template:51:0
	for i := 0; i < 5; i++ {

		// for.template:51:27
		{
			_n, _err := _output.Write([]byte(`
  `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}

		// for.template:52:2
		{
			_n, _err := _template.writeValue(
				_output,
				(i),
				"for.template:52:2")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}

		// for.template:52:4
		{
			_n, _err := _output.Write([]byte(`.0`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}

	// for.template:53:8
	{
		_n, _err := _output.Write([]byte(`

`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}

	// for.template:57:2
	{
		_n, _err := _output.Write([]byte(`

Slice Range Loop`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}

	// for.template:60:0
	for _, item := range []string{"foo", "bar"} {

		// for.template:60:48
		{
			_n, _err := _output.Write([]byte(`
  `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}

		// for.template:61:2
		{
			_n, _err := _template.writeValue(
				_output,
				(fmt.Sprintf("item: %s", item)),
				"for.template:61:2")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}

	}

	// for.template:62:8
	{
		_n, _err := _output.Write([]byte(`

Map Range Loop`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}

	// for.template:65:0
	for key, val := range map[string]int{"key1": 1, "key2": 2} {

		// for.template:65:61
		{
			_n, _err := _output.Write([]byte(`
  `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}

		// for.template:66:2
		{
			_n, _err := _template.writeValue(
				_output,
				(key),
				"for.template:66:2")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}

		// for.template:66:6
		{
			_n, _err := _output.Write([]byte(` -> `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}

		// for.template:66:10
		{
			_n, _err := _template.writeValue(
				_output,
				(val),
				"for.template:66:10")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}

	}

	// for.template:67:8
	{
		_n, _err := _output.Write([]byte(`

Channel Range Loop
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}

	// for.template:70:0
	for val := range Chan {

		// for.template:70:27
		{
			_n, _err := _template.writeValue(
				_output,
				(val),
				"for.template:70:27")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}

		// for.template:70:31
		{
			_n, _err := _output.Write([]byte(`,`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}

	// for.template:70:39
	{
		_n, _err := _output.Write([]byte(`

Infinite Loop 2:`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}

	// for.template:73:0
	cnt := 0

	// for.template:74:0
	for {

		// for.template:75:4
		if cnt > 3 {

			// for.template:75:22
			{
				_n, _err := _output.Write([]byte(`
`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}

			// for.template:78:8
			return _numWritten, nil

		}

		// for.template:79:13
		{
			_n, _err := _output.Write([]byte(`
Iteration: `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}

		// for.template:80:11
		{
			_n, _err := _template.writeValue(
				_output,
				(cnt),
				"for.template:80:11")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}

		// for.template:81:4
		cnt += 1

	}

	// for.template:82:7
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
