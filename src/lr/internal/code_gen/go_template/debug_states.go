// Auto-generated from source: debug_states.template

package go_template

import (
	_fmt "fmt"
	_io "io"

	lr "github.com/pattyshack/abc/src/lr/internal"
)

type DebugStates struct {
	OutputDebugNonKernelItems bool
	OrderedSymbols            []string
	OrderedStates             []*lr.ItemSet
}

func (DebugStates) Name() string { return "DebugStates" }

func (template *DebugStates) writeValue(
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

func (_template *DebugStates) WriteTo(
	_output _io.Writer) (
	int64,
	error) {

	_numWritten := int64(0)

	OutputDebugNonKernelItems := _template.OutputDebugNonKernelItems
	OrderedSymbols := _template.OrderedSymbols
	OrderedStates := _template.OrderedStates

	// debug_states.template:15:0

	gotoCount := 0
	reduceCount := 0
	shiftReduceCount := 0
	reduceReduceCount := 0

	// debug_states.template:22:3
	{
		_n, _err := _output.Write([]byte(`/*
Parser Debug States:`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// debug_states.template:25:0
	for _, state := range OrderedStates {
		// debug_states.template:25:40
		{
			_n, _err := _output.Write([]byte(`
  State `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// debug_states.template:26:8
		{
			_n, _err := _template.writeValue(
				_output,
				(state.StateNum),
				"debug_states.template:26:8")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// debug_states.template:26:25
		{
			_n, _err := _output.Write([]byte(`:
    Kernel Items:`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// debug_states.template:28:4
		firstNonKernel := true
		// debug_states.template:29:4
		for _, item := range state.Items {
			// debug_states.template:30:8
			if !item.IsKernel && firstNonKernel {
				// debug_states.template:31:12

				if !OutputDebugNonKernelItems &&
					len(state.ShiftReduceConflictSymbols) == 0 &&
					len(state.ReduceReduceConflictSymbols) == 0 {

					break
				}

				firstNonKernel = false

				// debug_states.template:42:14
				{
					_n, _err := _output.Write([]byte(`
    Non-kernel Items:`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
			}
			// debug_states.template:44:17
			{
				_n, _err := _output.Write([]byte(`
      `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// debug_states.template:46:6
			{
				_n, _err := _template.writeValue(
					_output,
					(item),
					"debug_states.template:46:6")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
		// debug_states.template:47:13
		{
			_n, _err := _output.Write([]byte(`
    Reduce:`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// debug_states.template:50:4
		if len(state.Reduce) == 0 {
			// debug_states.template:50:34
			{
				_n, _err := _output.Write([]byte(`
      (nil)`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
		// debug_states.template:54:4
		for _, symbol := range OrderedSymbols {
			// debug_states.template:55:8

			items := state.Reduce[symbol]
			reduceCount += len(items)

			if len(items) == 0 {
				continue
			}

			// debug_states.template:64:11
			{
				_n, _err := _output.Write([]byte(`
      `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// debug_states.template:66:6
			{
				_n, _err := _template.writeValue(
					_output,
					(symbol),
					"debug_states.template:66:6")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// debug_states.template:66:13
			{
				_n, _err := _output.Write([]byte(` -> [`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// debug_states.template:67:8
			for idx, item := range items {
				// debug_states.template:68:0
				{
					_n, _err := _template.writeValue(
						_output,
						(item.Name),
						"debug_states.template:68:0")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// debug_states.template:69:12
				if idx != len(items)-1 {
					// debug_states.template:69:41
					{
						_n, _err := _output.Write([]byte(` `))
						_numWritten += int64(_n)
						if _err != nil {
							return _numWritten, _err
						}
					}
				}
			}
			// debug_states.template:70:17
			{
				_n, _err := _output.Write([]byte(`]`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
		// debug_states.template:72:13
		{
			_n, _err := _output.Write([]byte(`
    Goto:`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// debug_states.template:75:4
		gotoCount += len(state.Goto)
		// debug_states.template:76:4
		if len(state.Goto) == 0 {
			// debug_states.template:76:33
			{
				_n, _err := _output.Write([]byte(`
      (nil)`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
		// debug_states.template:80:4
		for _, symbol := range OrderedSymbols {
			// debug_states.template:81:8
			child, ok := state.Goto[symbol]
			// debug_states.template:82:8
			if ok {
				// debug_states.template:82:18
				{
					_n, _err := _output.Write([]byte(`
      `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// debug_states.template:83:6
				{
					_n, _err := _template.writeValue(
						_output,
						(symbol),
						"debug_states.template:83:6")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// debug_states.template:83:13
				{
					_n, _err := _output.Write([]byte(` -> State `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// debug_states.template:83:23
				{
					_n, _err := _template.writeValue(
						_output,
						(child.StateNum),
						"debug_states.template:83:23")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
			}
		}
		// debug_states.template:87:4
		if len(state.ShiftReduceConflictSymbols) > 0 {
			// debug_states.template:88:8
			shiftReduceCount += len(state.ShiftReduceConflictSymbols)
			// debug_states.template:88:73
			{
				_n, _err := _output.Write([]byte(`
    Shift/reduce conflict symbols:
      [`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// debug_states.template:91:8
			for idx, symbol := range state.ShiftReduceConflictSymbols {
				// debug_states.template:92:0
				{
					_n, _err := _template.writeValue(
						_output,
						(symbol),
						"debug_states.template:92:0")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// debug_states.template:93:12
				if idx != len(state.ShiftReduceConflictSymbols)-1 {
					// debug_states.template:93:68
					{
						_n, _err := _output.Write([]byte(` `))
						_numWritten += int64(_n)
						if _err != nil {
							return _numWritten, _err
						}
					}
				}
			}
			// debug_states.template:94:17
			{
				_n, _err := _output.Write([]byte(`]`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
		// debug_states.template:98:4
		if len(state.ReduceReduceConflictSymbols) > 0 {
			// debug_states.template:99:8
			reduceReduceCount += len(state.ReduceReduceConflictSymbols)
			// debug_states.template:99:75
			{
				_n, _err := _output.Write([]byte(`
    Reduce/reduce conflict symbols:
      [`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// debug_states.template:102:8
			for idx, symbol := range state.ReduceReduceConflictSymbols {
				// debug_states.template:103:0
				{
					_n, _err := _template.writeValue(
						_output,
						(symbol),
						"debug_states.template:103:0")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// debug_states.template:104:12
				if idx != len(state.ShiftReduceConflictSymbols)-1 {
					// debug_states.template:104:68
					{
						_n, _err := _output.Write([]byte(` `))
						_numWritten += int64(_n)
						if _err != nil {
							return _numWritten, _err
						}
					}
				}
			}
			// debug_states.template:105:17
			{
				_n, _err := _output.Write([]byte(`]`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
		// debug_states.template:107:13
		{
			_n, _err := _output.Write([]byte(`
`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// debug_states.template:109:7
	{
		_n, _err := _output.Write([]byte(`
Number of states: `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// debug_states.template:110:18
	{
		_n, _err := _template.writeValue(
			_output,
			(len(OrderedStates)),
			"debug_states.template:110:18")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// debug_states.template:110:39
	{
		_n, _err := _output.Write([]byte(`
Number of shift actions: `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// debug_states.template:111:25
	{
		_n, _err := _template.writeValue(
			_output,
			(gotoCount),
			"debug_states.template:111:25")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// debug_states.template:111:35
	{
		_n, _err := _output.Write([]byte(`
Number of reduce actions: `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// debug_states.template:112:26
	{
		_n, _err := _template.writeValue(
			_output,
			(reduceCount),
			"debug_states.template:112:26")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// debug_states.template:112:38
	{
		_n, _err := _output.Write([]byte(`
Number of shift/reduce conflicts: `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// debug_states.template:113:34
	{
		_n, _err := _template.writeValue(
			_output,
			(shiftReduceCount),
			"debug_states.template:113:34")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// debug_states.template:113:51
	{
		_n, _err := _output.Write([]byte(`
Number of reduce/reduce conflicts: `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// debug_states.template:114:35
	{
		_n, _err := _template.writeValue(
			_output,
			(reduceReduceCount),
			"debug_states.template:114:35")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// debug_states.template:114:53
	{
		_n, _err := _output.Write([]byte(`
*/
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}

	return _numWritten, nil
}
