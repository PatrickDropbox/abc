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
	ActionTable      string
	OrderedSymbols   []*lr.Term
	Symbols          map[string]*lr.Term
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
	ActionTable := _template.ActionTable
	OrderedSymbols := _template.OrderedSymbols
	Symbols := _template.Symbols

	// action_table.template:34:0
	{
		_n, _err := _output.Write([]byte(`type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:34:5
	{
		_n, _err := _template.writeValue(
			_output,
			(TableKeyType),
			"action_table.template:34:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:34:18
	{
		_n, _err := _output.Write([]byte(` struct {
    `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:35:4
	{
		_n, _err := _template.writeValue(
			_output,
			(StateIdType),
			"action_table.template:35:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:35:16
	{
		_n, _err := _output.Write([]byte(`
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
			(SymbolIdType),
			"action_table.template:36:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:36:17
	{
		_n, _err := _output.Write([]byte(`
}

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:39:5
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionTableType),
			"action_table.template:39:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:39:21
	{
		_n, _err := _output.Write([]byte(` map[`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:39:26
	{
		_n, _err := _template.writeValue(
			_output,
			(TableKeyType),
			"action_table.template:39:26")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:39:39
	{
		_n, _err := _output.Write([]byte(`]*`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:39:41
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionType),
			"action_table.template:39:41")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:39:52
	{
		_n, _err := _output.Write([]byte(`

func (table `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:41:12
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionTableType),
			"action_table.template:41:12")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:41:28
	{
		_n, _err := _output.Write([]byte(`) Get(
    stateId `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:42:12
	{
		_n, _err := _template.writeValue(
			_output,
			(StateIdType),
			"action_table.template:42:12")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:42:24
	{
		_n, _err := _output.Write([]byte(`,
    symbolId `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:43:13
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"action_table.template:43:13")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:43:26
	{
		_n, _err := _output.Write([]byte(`) (
    *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:44:5
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionType),
			"action_table.template:44:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:44:16
	{
		_n, _err := _output.Write([]byte(`,
    bool) {

    action, ok := table[`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:47:24
	{
		_n, _err := _template.writeValue(
			_output,
			(TableKeyType),
			"action_table.template:47:24")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:47:37
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
	// action_table.template:52:23
	{
		_n, _err := _template.writeValue(
			_output,
			(TableKeyType),
			"action_table.template:52:23")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:52:36
	{
		_n, _err := _output.Write([]byte(`{stateId, `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:52:46
	{
		_n, _err := _template.writeValue(
			_output,
			(WildcardSymbolId),
			"action_table.template:52:46")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:52:63
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
	// action_table.template:57:0
	for _, state := range OrderedStates {
		// action_table.template:57:40
		{
			_n, _err := _output.Write([]byte(`
    `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// action_table.template:58:4
		{
			_n, _err := _template.writeValue(
				_output,
				(state.CodeGenAction),
				"action_table.template:58:4")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// action_table.template:58:26
		{
			_n, _err := _output.Write([]byte(` = &`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// action_table.template:58:30
		{
			_n, _err := _template.writeValue(
				_output,
				(ActionType),
				"action_table.template:58:30")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// action_table.template:58:41
		{
			_n, _err := _output.Write([]byte(`{`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// action_table.template:58:42
		{
			_n, _err := _template.writeValue(
				_output,
				(ShiftAction),
				"action_table.template:58:42")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// action_table.template:58:54
		{
			_n, _err := _output.Write([]byte(`, `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// action_table.template:58:56
		{
			_n, _err := _template.writeValue(
				_output,
				(state.CodeGenConst),
				"action_table.template:58:56")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// action_table.template:58:77
		{
			_n, _err := _output.Write([]byte(`, 0}`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// action_table.template:61:0
	for _, term := range OrderedSymbols {
		// action_table.template:62:4
		for _, clause := range term.Clauses {
			// action_table.template:62:44
			{
				_n, _err := _output.Write([]byte(`
    `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:63:4
			{
				_n, _err := _template.writeValue(
					_output,
					(clause.CodeGenReduceAction),
					"action_table.template:63:4")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:63:33
			{
				_n, _err := _output.Write([]byte(` = &`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:63:37
			{
				_n, _err := _template.writeValue(
					_output,
					(ActionType),
					"action_table.template:63:37")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:63:48
			{
				_n, _err := _output.Write([]byte(`{`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:63:49
			{
				_n, _err := _template.writeValue(
					_output,
					(ReduceAction),
					"action_table.template:63:49")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:63:62
			{
				_n, _err := _output.Write([]byte(`, 0, `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:63:67
			{
				_n, _err := _template.writeValue(
					_output,
					(clause.CodeGenReducerNameConst),
					"action_table.template:63:67")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:63:100
			{
				_n, _err := _output.Write([]byte(`}`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
	}
	// action_table.template:65:8
	{
		_n, _err := _output.Write([]byte(`
)

var `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:68:4
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionTable),
			"action_table.template:68:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:68:16
	{
		_n, _err := _output.Write([]byte(` = `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:68:19
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionTableType),
			"action_table.template:68:19")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:68:35
	{
		_n, _err := _output.Write([]byte(`{`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:69:0
	for _, state := range OrderedStates {
		// action_table.template:70:4
		for _, item := range state.Items {
			// action_table.template:71:8

			if !item.IsReduce {
				continue
			}

			if item.Name != lr.AcceptRule || item.LookAhead != lr.EndMarker {
				continue
			}

			// action_table.template:81:10
			{
				_n, _err := _output.Write([]byte(`
    {`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:82:5
			{
				_n, _err := _template.writeValue(
					_output,
					(state.CodeGenConst),
					"action_table.template:82:5")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:82:26
			{
				_n, _err := _output.Write([]byte(`, `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:82:28
			{
				_n, _err := _template.writeValue(
					_output,
					(EndSymbolId),
					"action_table.template:82:28")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:82:40
			{
				_n, _err := _output.Write([]byte(`}: &`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:82:44
			{
				_n, _err := _template.writeValue(
					_output,
					(ActionType),
					"action_table.template:82:44")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:82:55
			{
				_n, _err := _output.Write([]byte(`{`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:82:56
			{
				_n, _err := _template.writeValue(
					_output,
					(AcceptAction),
					"action_table.template:82:56")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:82:69
			{
				_n, _err := _output.Write([]byte(`, 0, 0},`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
	}
	// action_table.template:86:0
	for _, state := range OrderedStates {
		// action_table.template:87:4
		for _, symbol := range OrderedSymbols {
			// action_table.template:88:8

			child, ok := state.Goto[symbol.Name]
			if !ok {
				continue
			}

			// action_table.template:95:10
			{
				_n, _err := _output.Write([]byte(`
    {`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:96:5
			{
				_n, _err := _template.writeValue(
					_output,
					(state.CodeGenConst),
					"action_table.template:96:5")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:96:26
			{
				_n, _err := _output.Write([]byte(`, `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:96:28
			{
				_n, _err := _template.writeValue(
					_output,
					(symbol.CodeGenSymbolConst),
					"action_table.template:96:28")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:96:56
			{
				_n, _err := _output.Write([]byte(`}: `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:96:59
			{
				_n, _err := _template.writeValue(
					_output,
					(child.CodeGenAction),
					"action_table.template:96:59")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:96:81
			{
				_n, _err := _output.Write([]byte(`,`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
	}
	// action_table.template:100:0
	for _, state := range OrderedStates {
		// action_table.template:101:4
		for _, item := range state.Items {
			// action_table.template:102:8

			if !item.IsReduce {
				continue
			}

			if item.Name == lr.AcceptRule && item.LookAhead == lr.EndMarker {
				continue
			}

			idConst := Symbols[item.LookAhead].CodeGenSymbolConst
			reduceAction := item.Clause.CodeGenReduceAction

			// action_table.template:115:10
			{
				_n, _err := _output.Write([]byte(`
    {`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:116:5
			{
				_n, _err := _template.writeValue(
					_output,
					(state.CodeGenConst),
					"action_table.template:116:5")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:116:26
			{
				_n, _err := _output.Write([]byte(`, `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:116:28
			{
				_n, _err := _template.writeValue(
					_output,
					(idConst),
					"action_table.template:116:28")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:116:36
			{
				_n, _err := _output.Write([]byte(`}: `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:116:39
			{
				_n, _err := _template.writeValue(
					_output,
					(reduceAction),
					"action_table.template:116:39")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:116:52
			{
				_n, _err := _output.Write([]byte(`,`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
	}
	// action_table.template:118:8
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
