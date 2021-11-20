// Auto-generated from source: action_table.template

package go_template

import (
	_fmt "fmt"
	_io "io"

	lr "github.com/pattyshack/abc/src/lr/internal"
)

type ActionTable struct {
	ActionType    string
	ShiftAction   string
	ReduceAction  string
	OrderedStates []*lr.ItemSet
	NonTerminals  []*lr.Term
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

	ActionType := _template.ActionType
	ShiftAction := _template.ShiftAction
	ReduceAction := _template.ReduceAction
	OrderedStates := _template.OrderedStates
	NonTerminals := _template.NonTerminals

	// action_table.template:18:0
	{
		_n, _err := _output.Write([]byte(`var (`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// action_table.template:19:0
	for _, state := range OrderedStates {
		// action_table.template:19:40
		{
			_n, _err := _output.Write([]byte(`
    `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// action_table.template:20:4
		{
			_n, _err := _template.writeValue(
				_output,
				(state.CodeGenAction),
				"action_table.template:20:4")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// action_table.template:20:26
		{
			_n, _err := _output.Write([]byte(` = &`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// action_table.template:20:30
		{
			_n, _err := _template.writeValue(
				_output,
				(ActionType),
				"action_table.template:20:30")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// action_table.template:20:41
		{
			_n, _err := _output.Write([]byte(`{`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// action_table.template:20:42
		{
			_n, _err := _template.writeValue(
				_output,
				(ShiftAction),
				"action_table.template:20:42")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// action_table.template:20:54
		{
			_n, _err := _output.Write([]byte(`, `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// action_table.template:20:56
		{
			_n, _err := _template.writeValue(
				_output,
				(state.CodeGenConst),
				"action_table.template:20:56")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// action_table.template:20:77
		{
			_n, _err := _output.Write([]byte(`, 0}`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// action_table.template:23:0
	for _, term := range NonTerminals {
		// action_table.template:24:4
		for _, clause := range term.Clauses {
			// action_table.template:24:44
			{
				_n, _err := _output.Write([]byte(`
    `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:25:4
			{
				_n, _err := _template.writeValue(
					_output,
					(clause.CodeGenReduceAction),
					"action_table.template:25:4")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:25:33
			{
				_n, _err := _output.Write([]byte(` = &`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:25:37
			{
				_n, _err := _template.writeValue(
					_output,
					(ActionType),
					"action_table.template:25:37")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:25:48
			{
				_n, _err := _output.Write([]byte(`{`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:25:49
			{
				_n, _err := _template.writeValue(
					_output,
					(ReduceAction),
					"action_table.template:25:49")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:25:62
			{
				_n, _err := _output.Write([]byte(`, 0, `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:25:67
			{
				_n, _err := _template.writeValue(
					_output,
					(clause.CodeGenReducerNameConst),
					"action_table.template:25:67")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// action_table.template:25:100
			{
				_n, _err := _output.Write([]byte(`}`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
	}
	// action_table.template:27:8
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
