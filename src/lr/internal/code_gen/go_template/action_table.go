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
	EndSymbolId      string
	WildcardSymbolId string
	ActionTableType  string
	ActionType       string
	AcceptAction     string
	ShiftAction      string
	ReduceAction     string
	OrderedStates    []*lr.ItemSet
	NonTerminals     []*lr.Term
	ActionTable      string
	OrderedSymbolIds []string
	SymbolIdToConst  map[string]string
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
	EndSymbolId := _template.EndSymbolId
	WildcardSymbolId := _template.WildcardSymbolId
	ActionTableType := _template.ActionTableType
	ActionType := _template.ActionType
	AcceptAction := _template.AcceptAction
	ShiftAction := _template.ShiftAction
	ReduceAction := _template.ReduceAction
	OrderedStates := _template.OrderedStates
	NonTerminals := _template.NonTerminals
	ActionTable := _template.ActionTable
	OrderedSymbolIds := _template.OrderedSymbolIds
	SymbolIdToConst := _template.SymbolIdToConst

	// action_table.template:35:0
	{
		_n, _err := _output.Write([]byte(`type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:35:5
	{
		_n, _err := _template.writeValue(
			_output,
			(TableKeyType),
			"action_table.template:35:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:35:18
	{
		_n, _err := _output.Write([]byte(` struct {
    `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:36:4
	{
		_n, _err := _template.writeValue(
			_output,
			(StateIdType),
			"action_table.template:36:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:36:16
	{
		_n, _err := _output.Write([]byte(`
    `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:37:4
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"action_table.template:37:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:37:17
	{
		_n, _err := _output.Write([]byte(`
}

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:40:5
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionTableType),
			"action_table.template:40:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:40:21
	{
		_n, _err := _output.Write([]byte(` map[`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:40:26
	{
		_n, _err := _template.writeValue(
			_output,
			(TableKeyType),
			"action_table.template:40:26")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:40:39
	{
		_n, _err := _output.Write([]byte(`]*`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:40:41
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionType),
			"action_table.template:40:41")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:40:52
	{
		_n, _err := _output.Write([]byte(`

func (table `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:42:12
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionTableType),
			"action_table.template:42:12")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:42:28
	{
		_n, _err := _output.Write([]byte(`) Get(
    stateId `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:43:12
	{
		_n, _err := _template.writeValue(
			_output,
			(StateIdType),
			"action_table.template:43:12")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:43:24
	{
		_n, _err := _output.Write([]byte(`,
    symbolId `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:44:13
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"action_table.template:44:13")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:44:26
	{
		_n, _err := _output.Write([]byte(`) (
    *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:45:5
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionType),
			"action_table.template:45:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:45:16
	{
		_n, _err := _output.Write([]byte(`,
    bool) {

    action, ok := table[`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:48:24
	{
		_n, _err := _template.writeValue(
			_output,
			(TableKeyType),
			"action_table.template:48:24")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:48:37
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
	// action_table.template:53:23
	{
		_n, _err := _template.writeValue(
			_output,
			(TableKeyType),
			"action_table.template:53:23")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:53:36
	{
		_n, _err := _output.Write([]byte(`{stateId, `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:53:46
	{
		_n, _err := _template.writeValue(
			_output,
			(WildcardSymbolId),
			"action_table.template:53:46")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:53:63
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
	// action_table.template:58:0
	for _, state := range OrderedStates {
		// action_table.template:58:40
		{
			_n, _err := _output.Write([]byte(`
    `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// action_table.template:59:4
		{
			_n, _err := _template.writeValue(
				_output,
				(state.CodeGenAction),
				"action_table.template:59:4")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// action_table.template:59:26
		{
			_n, _err := _output.Write([]byte(` = &`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// action_table.template:59:30
		{
			_n, _err := _template.writeValue(
				_output,
				(ActionType),
				"action_table.template:59:30")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// action_table.template:59:41
		{
			_n, _err := _output.Write([]byte(`{`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// action_table.template:59:42
		{
			_n, _err := _template.writeValue(
				_output,
				(ShiftAction),
				"action_table.template:59:42")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// action_table.template:59:54
		{
			_n, _err := _output.Write([]byte(`, `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// action_table.template:59:56
		{
			_n, _err := _template.writeValue(
				_output,
				(state.CodeGenConst),
				"action_table.template:59:56")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// action_table.template:59:77
		{
			_n, _err := _output.Write([]byte(`, 0}`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// action_table.template:62:0
	for _, term := range NonTerminals {
		// action_table.template:63:4
		for _, clause := range term.Clauses {
			// action_table.template:63:44
			{
				_n, _err := _output.Write([]byte(`
    `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:64:4
			{
				_n, _err := _template.writeValue(
					_output,
					(clause.CodeGenReduceAction),
					"action_table.template:64:4")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:64:33
			{
				_n, _err := _output.Write([]byte(` = &`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:64:37
			{
				_n, _err := _template.writeValue(
					_output,
					(ActionType),
					"action_table.template:64:37")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:64:48
			{
				_n, _err := _output.Write([]byte(`{`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:64:49
			{
				_n, _err := _template.writeValue(
					_output,
					(ReduceAction),
					"action_table.template:64:49")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:64:62
			{
				_n, _err := _output.Write([]byte(`, 0, `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:64:67
			{
				_n, _err := _template.writeValue(
					_output,
					(clause.CodeGenReducerNameConst),
					"action_table.template:64:67")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:64:100
			{
				_n, _err := _output.Write([]byte(`}`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
	}
	// action_table.template:66:8
	{
		_n, _err := _output.Write([]byte(`
)

var `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:69:4
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionTable),
			"action_table.template:69:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:69:16
	{
		_n, _err := _output.Write([]byte(` = `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:69:19
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionTableType),
			"action_table.template:69:19")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:69:35
	{
		_n, _err := _output.Write([]byte(`{`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:70:0
	for _, state := range OrderedStates {
		// action_table.template:71:4
		for _, item := range state.Items {
			// action_table.template:72:8

			if !item.IsReduce {
				continue
			}

			if item.Name != lr.AcceptRule || item.LookAhead != lr.EndMarker {
				continue
			}

			// action_table.template:82:10
			{
				_n, _err := _output.Write([]byte(`
    {`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:83:5
			{
				_n, _err := _template.writeValue(
					_output,
					(state.CodeGenConst),
					"action_table.template:83:5")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:83:26
			{
				_n, _err := _output.Write([]byte(`, `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:83:28
			{
				_n, _err := _template.writeValue(
					_output,
					(EndSymbolId),
					"action_table.template:83:28")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:83:40
			{
				_n, _err := _output.Write([]byte(`}: &`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:83:44
			{
				_n, _err := _template.writeValue(
					_output,
					(ActionType),
					"action_table.template:83:44")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:83:55
			{
				_n, _err := _output.Write([]byte(`{`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:83:56
			{
				_n, _err := _template.writeValue(
					_output,
					(AcceptAction),
					"action_table.template:83:56")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:83:69
			{
				_n, _err := _output.Write([]byte(`, 0, 0},`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
	}
	// action_table.template:87:0
	for _, state := range OrderedStates {
		// action_table.template:88:4
		for _, symbol := range OrderedSymbolIds {
			// action_table.template:89:8

			child, ok := state.Goto[symbol]
			if !ok {
				continue
			}

			// action_table.template:96:10
			{
				_n, _err := _output.Write([]byte(`
    {`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:97:5
			{
				_n, _err := _template.writeValue(
					_output,
					(state.CodeGenConst),
					"action_table.template:97:5")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:97:26
			{
				_n, _err := _output.Write([]byte(`, `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:97:28
			{
				_n, _err := _template.writeValue(
					_output,
					(SymbolIdToConst[symbol]),
					"action_table.template:97:28")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:97:54
			{
				_n, _err := _output.Write([]byte(`}: `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:97:57
			{
				_n, _err := _template.writeValue(
					_output,
					(child.CodeGenAction),
					"action_table.template:97:57")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:97:79
			{
				_n, _err := _output.Write([]byte(`,`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
	}
	// action_table.template:101:0
	for _, state := range OrderedStates {
		// action_table.template:102:4
		for _, item := range state.Items {
			// action_table.template:103:8

			if !item.IsReduce {
				continue
			}

			if item.Name == lr.AcceptRule && item.LookAhead == lr.EndMarker {
				continue
			}

			idConst := SymbolIdToConst[item.LookAhead]
			reduceAction := item.Clause.CodeGenReduceAction

			// action_table.template:116:10
			{
				_n, _err := _output.Write([]byte(`
    {`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:117:5
			{
				_n, _err := _template.writeValue(
					_output,
					(state.CodeGenConst),
					"action_table.template:117:5")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:117:26
			{
				_n, _err := _output.Write([]byte(`, `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:117:28
			{
				_n, _err := _template.writeValue(
					_output,
					(idConst),
					"action_table.template:117:28")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:117:36
			{
				_n, _err := _output.Write([]byte(`}: `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:117:39
			{
				_n, _err := _template.writeValue(
					_output,
					(reduceAction),
					"action_table.template:117:39")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:117:52
			{
				_n, _err := _output.Write([]byte(`,`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
	}
	// action_table.template:119:8
	{
		_n, _err := _output.Write([]byte(`
}
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}

	return _numWritten, nil
}
