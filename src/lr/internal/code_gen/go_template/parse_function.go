// Auto-generated from source: parse_function.template

package go_template

import (
	_fmt "fmt"
	_io "io"
)

type ParseFunc struct {
	ParseFuncName   string
	LexerType       string
	ReducerType     string
	ErrHandlerType  string
	StateIdType     string
	SymbolType      string
	StackItemType   string
	StackType       string
	SymbolStackType string
	ActionTable     string
	AcceptAction    string
	ShiftAction     string
	ReduceAction    string
}

func (ParseFunc) Name() string { return "ParseFunc" }

func (template *ParseFunc) writeValue(
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

func (_template *ParseFunc) WriteTo(
	_output _io.Writer) (
	int64,
	error) {

	_numWritten := int64(0)

	ParseFuncName := _template.ParseFuncName
	LexerType := _template.LexerType
	ReducerType := _template.ReducerType
	ErrHandlerType := _template.ErrHandlerType
	StateIdType := _template.StateIdType
	SymbolType := _template.SymbolType
	StackItemType := _template.StackItemType
	StackType := _template.StackType
	SymbolStackType := _template.SymbolStackType
	ActionTable := _template.ActionTable
	AcceptAction := _template.AcceptAction
	ShiftAction := _template.ShiftAction
	ReduceAction := _template.ReduceAction

	// parse_function.template:24:0
	{
		_n, _err := _output.Write([]byte(`func `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// parse_function.template:24:5
	{
		_n, _err := _template.writeValue(
			_output,
			(ParseFuncName),
			"parse_function.template:24:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// parse_function.template:24:19
	{
		_n, _err := _output.Write([]byte(`(
    lexer `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// parse_function.template:25:10
	{
		_n, _err := _template.writeValue(
			_output,
			(LexerType),
			"parse_function.template:25:10")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// parse_function.template:25:20
	{
		_n, _err := _output.Write([]byte(`,
    reducer `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// parse_function.template:26:12
	{
		_n, _err := _template.writeValue(
			_output,
			(ReducerType),
			"parse_function.template:26:12")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// parse_function.template:26:24
	{
		_n, _err := _output.Write([]byte(`,
    errHandler `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// parse_function.template:27:15
	{
		_n, _err := _template.writeValue(
			_output,
			(ErrHandlerType),
			"parse_function.template:27:15")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// parse_function.template:27:30
	{
		_n, _err := _output.Write([]byte(`,
    startState `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// parse_function.template:28:15
	{
		_n, _err := _template.writeValue(
			_output,
			(StateIdType),
			"parse_function.template:28:15")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// parse_function.template:28:27
	{
		_n, _err := _output.Write([]byte(`) (
    *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// parse_function.template:29:5
	{
		_n, _err := _template.writeValue(
			_output,
			(StackItemType),
			"parse_function.template:29:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// parse_function.template:29:19
	{
		_n, _err := _output.Write([]byte(`,
    error) {

    stateStack := `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// parse_function.template:32:18
	{
		_n, _err := _template.writeValue(
			_output,
			(StackType),
			"parse_function.template:32:18")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// parse_function.template:32:28
	{
		_n, _err := _output.Write([]byte(`{
        // Note: we don't have to populate the start symbol since its value
        // is never accessed.
        &`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// parse_function.template:35:9
	{
		_n, _err := _template.writeValue(
			_output,
			(StackItemType),
			"parse_function.template:35:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// parse_function.template:35:23
	{
		_n, _err := _output.Write([]byte(`{startState, nil},
    }

    symbolStack := &`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// parse_function.template:38:20
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolStackType),
			"parse_function.template:38:20")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// parse_function.template:38:36
	{
		_n, _err := _output.Write([]byte(`{lexer: lexer}

    for {
        nextSymbol, err := symbolStack.Top()
        if err != nil {
            return nil, err
        }

        action, ok := `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// parse_function.template:46:22
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionTable),
			"parse_function.template:46:22")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// parse_function.template:46:34
	{
		_n, _err := _output.Write([]byte(`.Get(
            stateStack[len(stateStack)-1].StateId,
            nextSymbol.Id())
        if !ok {
            return nil, errHandler.Error(nextSymbol, stateStack)
        }

        if action.ActionType == `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// parse_function.template:53:32
	{
		_n, _err := _template.writeValue(
			_output,
			(ShiftAction),
			"parse_function.template:53:32")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// parse_function.template:53:44
	{
		_n, _err := _output.Write([]byte(` {
            stateStack = append(stateStack, action.ShiftItem(nextSymbol))

            _, err = symbolStack.Pop()
            if err != nil {
                return nil, err
            }
        } else if action.ActionType == `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// parse_function.template:60:39
	{
		_n, _err := _template.writeValue(
			_output,
			(ReduceAction),
			"parse_function.template:60:39")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// parse_function.template:60:52
	{
		_n, _err := _output.Write([]byte(` {
            var reduceSymbol *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// parse_function.template:61:30
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"parse_function.template:61:30")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// parse_function.template:61:41
	{
		_n, _err := _output.Write([]byte(`
            stateStack, reduceSymbol, err = action.ReduceSymbol(
                reducer,
                stateStack)
            if err != nil {
                return nil, err
            }

            symbolStack.Push(reduceSymbol)
        } else if action.ActionType == `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// parse_function.template:70:39
	{
		_n, _err := _template.writeValue(
			_output,
			(AcceptAction),
			"parse_function.template:70:39")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// parse_function.template:70:52
	{
		_n, _err := _output.Write([]byte(` {
            if len(stateStack) != 2 {
                panic("This should never happen")
            }
            return stateStack[1], nil
        } else {
            panic("Unknown action type: " + action.ActionType.String())
        }
    }
}
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}

	return _numWritten, nil
}
