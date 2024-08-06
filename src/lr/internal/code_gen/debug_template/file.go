// Auto-generated from source: file.template

package debug_template

import (
	_fmt "fmt"
	_io "io"

	lr "github.com/pattyshack/abc/src/lr/internal"
)

type File struct {
	OrderedSymbolNames        []string
	States                    *lr.LRStates
	OutputDebugNonKernelItems bool
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

	OrderedSymbolNames := _template.OrderedSymbolNames
	States := _template.States
	OutputDebugNonKernelItems := _template.OutputDebugNonKernelItems

	// file.template:14:0

	gotoCount := 0
	reduceCount := 0
	shiftReduceCount := 0
	reduceReduceCount := 0

	// file.template:21:3
	{
		_n, _err := _output.Write([]byte(`Parser Debug States:`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:23:0
	for _, state := range States.OrderedStates {
		// file.template:23:47
		{
			_n, _err := _output.Write([]byte(`
  State `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:24:8
		{
			_n, _err := _template.writeValue(
				_output,
				(state.StateNum),
				"file.template:24:8")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:24:25
		{
			_n, _err := _output.Write([]byte(`:
    Kernel Items:`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:26:4
		firstNonKernel := true
		// file.template:27:4
		for _, item := range state.Items {
			// file.template:28:8
			if !item.IsKernel && firstNonKernel {
				// file.template:29:12

				if !OutputDebugNonKernelItems &&
					len(state.ShiftReduceConflictSymbols) == 0 &&
					len(state.ReduceReduceConflictSymbols) == 0 {

					break
				}

				firstNonKernel = false

				// file.template:40:14
				{
					_n, _err := _output.Write([]byte(`
    Non-kernel Items:`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
			}
			// file.template:42:17
			{
				_n, _err := _output.Write([]byte(`
      `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:44:6
			{
				_n, _err := _template.writeValue(
					_output,
					(item),
					"file.template:44:6")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
		// file.template:45:13
		{
			_n, _err := _output.Write([]byte(`
    Reduce:`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:48:4
		if len(state.Reduce) == 0 {
			// file.template:48:34
			{
				_n, _err := _output.Write([]byte(`
      (nil)`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
		// file.template:52:4
		for _, symbolName := range OrderedSymbolNames {
			// file.template:53:8

			items := state.Reduce[symbolName]
			reduceCount += len(items)

			if len(items) == 0 {
				continue
			}

			// file.template:62:11
			{
				_n, _err := _output.Write([]byte(`
      `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:64:6
			{
				_n, _err := _template.writeValue(
					_output,
					(symbolName),
					"file.template:64:6")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:64:17
			{
				_n, _err := _output.Write([]byte(` -> [`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:65:8
			for idx, item := range items {
				// file.template:66:0
				{
					_n, _err := _template.writeValue(
						_output,
						(item.Name),
						"file.template:66:0")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:67:12
				if idx != len(items)-1 {
					// file.template:67:41
					{
						_n, _err := _output.Write([]byte(` `))
						_numWritten += int64(_n)
						if _err != nil {
							return _numWritten, _err
						}
					}
				}
			}
			// file.template:68:17
			{
				_n, _err := _output.Write([]byte(`]`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
		// file.template:70:13
		{
			_n, _err := _output.Write([]byte(`
    Goto:`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:73:4
		gotoCount += len(state.Goto)
		// file.template:74:4
		if len(state.Goto) == 0 {
			// file.template:74:33
			{
				_n, _err := _output.Write([]byte(`
      (nil)`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
		// file.template:78:4
		for _, symbolName := range OrderedSymbolNames {
			// file.template:79:8
			child, ok := state.Goto[symbolName]
			// file.template:80:8
			if ok {
				// file.template:80:18
				{
					_n, _err := _output.Write([]byte(`
      `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:81:6
				{
					_n, _err := _template.writeValue(
						_output,
						(symbolName),
						"file.template:81:6")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:81:17
				{
					_n, _err := _output.Write([]byte(` -> State `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:81:27
				{
					_n, _err := _template.writeValue(
						_output,
						(child.StateNum),
						"file.template:81:27")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
			}
		}
		// file.template:85:4
		if len(state.ShiftReduceConflictSymbols) > 0 {
			// file.template:86:8
			shiftReduceCount += len(state.ShiftReduceConflictSymbols)
			// file.template:86:73
			{
				_n, _err := _output.Write([]byte(`
    Shift/reduce conflict symbols:
      [`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:89:8
			for idx, symbol := range state.ShiftReduceConflictSymbols {
				// file.template:90:0
				{
					_n, _err := _template.writeValue(
						_output,
						(symbol),
						"file.template:90:0")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:91:12
				if idx != len(state.ShiftReduceConflictSymbols)-1 {
					// file.template:91:68
					{
						_n, _err := _output.Write([]byte(` `))
						_numWritten += int64(_n)
						if _err != nil {
							return _numWritten, _err
						}
					}
				}
			}
			// file.template:92:17
			{
				_n, _err := _output.Write([]byte(`]`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
		// file.template:96:4
		if len(state.ReduceReduceConflictSymbols) > 0 {
			// file.template:97:8
			reduceReduceCount += len(state.ReduceReduceConflictSymbols)
			// file.template:97:75
			{
				_n, _err := _output.Write([]byte(`
    Reduce/reduce conflict symbols:
      [`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:100:8
			for idx, symbol := range state.ReduceReduceConflictSymbols {
				// file.template:101:0
				{
					_n, _err := _template.writeValue(
						_output,
						(symbol),
						"file.template:101:0")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:102:12
				if idx != len(state.ShiftReduceConflictSymbols)-1 {
					// file.template:102:68
					{
						_n, _err := _output.Write([]byte(` `))
						_numWritten += int64(_n)
						if _err != nil {
							return _numWritten, _err
						}
					}
				}
			}
			// file.template:103:17
			{
				_n, _err := _output.Write([]byte(`]`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
		// file.template:105:13
		{
			_n, _err := _output.Write([]byte(`
`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// file.template:107:7
	{
		_n, _err := _output.Write([]byte(`
Number of states: `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:108:18
	{
		_n, _err := _template.writeValue(
			_output,
			(len(States.OrderedStates)),
			"file.template:108:18")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:108:46
	{
		_n, _err := _output.Write([]byte(`
Number of shift actions: `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:109:25
	{
		_n, _err := _template.writeValue(
			_output,
			(gotoCount),
			"file.template:109:25")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:109:35
	{
		_n, _err := _output.Write([]byte(`
Number of reduce actions: `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:110:26
	{
		_n, _err := _template.writeValue(
			_output,
			(reduceCount),
			"file.template:110:26")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:110:38
	{
		_n, _err := _output.Write([]byte(`
Number of shift/reduce conflicts: `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:111:34
	{
		_n, _err := _template.writeValue(
			_output,
			(shiftReduceCount),
			"file.template:111:34")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:111:51
	{
		_n, _err := _output.Write([]byte(`
Number of reduce/reduce conflicts: `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:112:35
	{
		_n, _err := _template.writeValue(
			_output,
			(reduceReduceCount),
			"file.template:112:35")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:112:53
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
