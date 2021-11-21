// Auto-generated from source: action_table.template

package go_template

import (
	_fmt "fmt"
	_io "io"

	lr "github.com/pattyshack/abc/src/lr/internal"
)

type ActionTable struct {
	TableKeyType     string
	StateIdType      string
	SymbolIdType     string
	WildcardSymbolId string
	ActionTableType  string
	ActionType       string
	ShiftAction      string
	ReduceAction     string
	OrderedStates    []*lr.ItemSet
	NonTerminals     []*lr.Term
}

func (ActionTable) Name() string { return "ActionTable" }

func (template *ActionTable) writeValue(
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

func (_template *ActionTable) WriteTo(
	_output _io.Writer) (
	int64,
	error) {

	_numWritten := int64(0)

	TableKeyType := _template.TableKeyType
	StateIdType := _template.StateIdType
	SymbolIdType := _template.SymbolIdType
	WildcardSymbolId := _template.WildcardSymbolId
	ActionTableType := _template.ActionTableType
	ActionType := _template.ActionType
	ShiftAction := _template.ShiftAction
	ReduceAction := _template.ReduceAction
	OrderedStates := _template.OrderedStates
	NonTerminals := _template.NonTerminals

	// action_table.template:26:0
	{
		_n, _err := _output.Write([]byte(`type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:26:5
	{
		_n, _err := _template.writeValue(
			_output,
			(TableKeyType),
			"action_table.template:26:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:26:18
	{
		_n, _err := _output.Write([]byte(` struct {
    `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:27:4
	{
		_n, _err := _template.writeValue(
			_output,
			(StateIdType),
			"action_table.template:27:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:27:16
	{
		_n, _err := _output.Write([]byte(`
    `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:28:4
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"action_table.template:28:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:28:17
	{
		_n, _err := _output.Write([]byte(`
}

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:31:5
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionTableType),
			"action_table.template:31:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:31:21
	{
		_n, _err := _output.Write([]byte(` map[`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:31:26
	{
		_n, _err := _template.writeValue(
			_output,
			(TableKeyType),
			"action_table.template:31:26")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:31:39
	{
		_n, _err := _output.Write([]byte(`]*`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:31:41
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionType),
			"action_table.template:31:41")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:31:52
	{
		_n, _err := _output.Write([]byte(`

func (table `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:33:12
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionTableType),
			"action_table.template:33:12")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:33:28
	{
		_n, _err := _output.Write([]byte(`) Get(
    stateId `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:34:12
	{
		_n, _err := _template.writeValue(
			_output,
			(StateIdType),
			"action_table.template:34:12")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:34:24
	{
		_n, _err := _output.Write([]byte(`,
    symbolId `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:35:13
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"action_table.template:35:13")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:35:26
	{
		_n, _err := _output.Write([]byte(`) (
    *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:36:5
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionType),
			"action_table.template:36:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:36:16
	{
		_n, _err := _output.Write([]byte(`,
    bool) {

    action, ok := table[`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:39:24
	{
		_n, _err := _template.writeValue(
			_output,
			(TableKeyType),
			"action_table.template:39:24")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:39:37
	{
		_n, _err := _output.Write([]byte(`{stateId, symbolId}]
    if ok {
        return action, ok
    }

    action, ok = table[`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:44:23
	{
		_n, _err := _template.writeValue(
			_output,
			(TableKeyType),
			"action_table.template:44:23")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:44:36
	{
		_n, _err := _output.Write([]byte(`{stateId, `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:44:46
	{
		_n, _err := _template.writeValue(
			_output,
			(WildcardSymbolId),
			"action_table.template:44:46")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:44:63
	{
		_n, _err := _output.Write([]byte(`}]
    return action, ok
}

var (`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:49:0
	for _, state := range OrderedStates {
		// action_table.template:49:40
		{
			_n, _err := _output.Write([]byte(`
    `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// action_table.template:50:4
		{
			_n, _err := _template.writeValue(
				_output,
				(state.CodeGenAction),
				"action_table.template:50:4")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// action_table.template:50:26
		{
			_n, _err := _output.Write([]byte(` = &`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// action_table.template:50:30
		{
			_n, _err := _template.writeValue(
				_output,
				(ActionType),
				"action_table.template:50:30")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// action_table.template:50:41
		{
			_n, _err := _output.Write([]byte(`{`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// action_table.template:50:42
		{
			_n, _err := _template.writeValue(
				_output,
				(ShiftAction),
				"action_table.template:50:42")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// action_table.template:50:54
		{
			_n, _err := _output.Write([]byte(`, `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// action_table.template:50:56
		{
			_n, _err := _template.writeValue(
				_output,
				(state.CodeGenConst),
				"action_table.template:50:56")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// action_table.template:50:77
		{
			_n, _err := _output.Write([]byte(`, 0}`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// action_table.template:53:0
	for _, term := range NonTerminals {
		// action_table.template:54:4
		for _, clause := range term.Clauses {
			// action_table.template:54:44
			{
				_n, _err := _output.Write([]byte(`
    `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:55:4
			{
				_n, _err := _template.writeValue(
					_output,
					(clause.CodeGenReduceAction),
					"action_table.template:55:4")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:55:33
			{
				_n, _err := _output.Write([]byte(` = &`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:55:37
			{
				_n, _err := _template.writeValue(
					_output,
					(ActionType),
					"action_table.template:55:37")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:55:48
			{
				_n, _err := _output.Write([]byte(`{`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:55:49
			{
				_n, _err := _template.writeValue(
					_output,
					(ReduceAction),
					"action_table.template:55:49")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:55:62
			{
				_n, _err := _output.Write([]byte(`, 0, `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:55:67
			{
				_n, _err := _template.writeValue(
					_output,
					(clause.CodeGenReducerNameConst),
					"action_table.template:55:67")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:55:100
			{
				_n, _err := _output.Write([]byte(`}`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
	}
	// action_table.template:57:8
	{
		_n, _err := _output.Write([]byte(`
)
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}

	return _numWritten, nil
}
