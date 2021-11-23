// Auto-generated from source: internal_definitions.template

package go_template

import (
	_fmt "fmt"
	_io "io"

	lr "github.com/pattyshack/abc/src/lr/internal"
	parser "github.com/pattyshack/abc/src/lr/internal/parser"
)

type InternalDefinitions struct {
	ActionType        string
	ActionIdType      string
	ShiftAction       string
	ReduceAction      string
	AcceptAction      string
	StateIdType       string
	ReduceType        string
	SymbolType        string
	GenericSymbolType string
	StackItemType     string
	StackType         string
	LexerType         string
	ReducerType       string
	SymbolStackType   string
	SymbolIdType      string
	EndSymbolId       string
	WildcardSymbolId  string
	LocationType      string
	TokenType         string
	Sprintf           interface{}
	Errorf            interface{}
	EOF               interface{}
	OrderedSymbols    []*lr.Term
	OrderedStates     []*lr.ItemSet
	OrderedValueTypes lr.ParamList
}

func (InternalDefinitions) Name() string { return "InternalDefinitions" }

func (template *InternalDefinitions) writeValue(
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

func (_template *InternalDefinitions) WriteTo(
	_output _io.Writer) (
	int64,
	error) {

	_numWritten := int64(0)

	ActionType := _template.ActionType
	ActionIdType := _template.ActionIdType
	ShiftAction := _template.ShiftAction
	ReduceAction := _template.ReduceAction
	AcceptAction := _template.AcceptAction
	StateIdType := _template.StateIdType
	ReduceType := _template.ReduceType
	SymbolType := _template.SymbolType
	GenericSymbolType := _template.GenericSymbolType
	StackItemType := _template.StackItemType
	StackType := _template.StackType
	LexerType := _template.LexerType
	ReducerType := _template.ReducerType
	SymbolStackType := _template.SymbolStackType
	SymbolIdType := _template.SymbolIdType
	EndSymbolId := _template.EndSymbolId
	WildcardSymbolId := _template.WildcardSymbolId
	LocationType := _template.LocationType
	TokenType := _template.TokenType
	Sprintf := _template.Sprintf
	Errorf := _template.Errorf
	EOF := _template.EOF
	OrderedSymbols := _template.OrderedSymbols
	OrderedStates := _template.OrderedStates
	OrderedValueTypes := _template.OrderedValueTypes

	// internal_definitions.template:48:0
	{
		_n, _err := _output.Write([]byte(`func (i `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:48:8
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"internal_definitions.template:48:8")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:48:21
	{
		_n, _err := _output.Write([]byte(`) String() string {
    switch i {
    case `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:50:9
	{
		_n, _err := _template.writeValue(
			_output,
			(EndSymbolId),
			"internal_definitions.template:50:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:50:21
	{
		_n, _err := _output.Write([]byte(`:
        return "$"
    case `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:52:9
	{
		_n, _err := _template.writeValue(
			_output,
			(WildcardSymbolId),
			"internal_definitions.template:52:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:52:26
	{
		_n, _err := _output.Write([]byte(`:
        return "*"`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:54:0
	for _, term := range OrderedSymbols[3:] {
		// internal_definitions.template:55:4
		if term.SymbolId == parser.LRCharacterToken {
			// internal_definitions.template:56:8

			escaped := term.Name
			if term.Name == "'\"'" {
				escaped = "'\\\"'"
			} else if escaped[1] == '\\' {
				escaped = "'\\\\" + term.Name[2:]
			}

			// internal_definitions.template:65:10
			{
				_n, _err := _output.Write([]byte(`
    case `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// internal_definitions.template:66:9
			{
				_n, _err := _template.writeValue(
					_output,
					(term.CodeGenSymbolConst),
					"internal_definitions.template:66:9")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// internal_definitions.template:66:35
			{
				_n, _err := _output.Write([]byte(`:
        return "`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// internal_definitions.template:67:16
			{
				_n, _err := _template.writeValue(
					_output,
					(escaped),
					"internal_definitions.template:67:16")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// internal_definitions.template:67:24
			{
				_n, _err := _output.Write([]byte(`"`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		} else {
			// internal_definitions.template:68:13
			{
				_n, _err := _output.Write([]byte(`
    case `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// internal_definitions.template:69:9
			{
				_n, _err := _template.writeValue(
					_output,
					(term.CodeGenSymbolConst),
					"internal_definitions.template:69:9")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// internal_definitions.template:69:35
			{
				_n, _err := _output.Write([]byte(`:
        return "`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// internal_definitions.template:70:16
			{
				_n, _err := _template.writeValue(
					_output,
					(term.Name),
					"internal_definitions.template:70:16")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// internal_definitions.template:70:28
			{
				_n, _err := _output.Write([]byte(`"`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
	}
	// internal_definitions.template:72:8
	{
		_n, _err := _output.Write([]byte(`
    default:
        return `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:74:15
	{
		_n, _err := _template.writeValue(
			_output,
			(Sprintf),
			"internal_definitions.template:74:15")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:74:23
	{
		_n, _err := _output.Write([]byte(`("?unknown symbol %d?", int(i))
    }
}

const (
    `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:79:4
	{
		_n, _err := _template.writeValue(
			_output,
			(EndSymbolId),
			"internal_definitions.template:79:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:79:16
	{
		_n, _err := _output.Write([]byte(` = `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:79:19
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"internal_definitions.template:79:19")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:79:32
	{
		_n, _err := _output.Write([]byte(`(0)
    `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:80:4
	{
		_n, _err := _template.writeValue(
			_output,
			(WildcardSymbolId),
			"internal_definitions.template:80:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:80:21
	{
		_n, _err := _output.Write([]byte(` = `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:80:24
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"internal_definitions.template:80:24")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:80:37
	{
		_n, _err := _output.Write([]byte(`(-1)
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:82:0
	for idx, term := range OrderedSymbols[3:] {
		// internal_definitions.template:83:4
		if term.IsTerminal {
			// internal_definitions.template:84:8
			continue
		}
		// internal_definitions.template:85:12
		{
			_n, _err := _output.Write([]byte(`
    `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// internal_definitions.template:86:4
		{
			_n, _err := _template.writeValue(
				_output,
				(term.CodeGenSymbolConst),
				"internal_definitions.template:86:4")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// internal_definitions.template:86:30
		{
			_n, _err := _output.Write([]byte(` = `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// internal_definitions.template:86:33
		{
			_n, _err := _template.writeValue(
				_output,
				(SymbolIdType),
				"internal_definitions.template:86:33")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// internal_definitions.template:86:46
		{
			_n, _err := _output.Write([]byte(`(`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// internal_definitions.template:86:47
		{
			_n, _err := _template.writeValue(
				_output,
				(256 + idx),
				"internal_definitions.template:86:47")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// internal_definitions.template:86:57
		{
			_n, _err := _output.Write([]byte(`)`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// internal_definitions.template:87:8
	{
		_n, _err := _output.Write([]byte(`
)

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:90:5
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionIdType),
			"internal_definitions.template:90:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:90:18
	{
		_n, _err := _output.Write([]byte(` int

const (
    // NOTE: error action is implicit
    `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:94:4
	{
		_n, _err := _template.writeValue(
			_output,
			(ShiftAction),
			"internal_definitions.template:94:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:94:16
	{
		_n, _err := _output.Write([]byte(` = `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:94:19
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionIdType),
			"internal_definitions.template:94:19")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:94:32
	{
		_n, _err := _output.Write([]byte(`(0)
    `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:95:4
	{
		_n, _err := _template.writeValue(
			_output,
			(ReduceAction),
			"internal_definitions.template:95:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:95:17
	{
		_n, _err := _output.Write([]byte(` = `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:95:20
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionIdType),
			"internal_definitions.template:95:20")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:95:33
	{
		_n, _err := _output.Write([]byte(`(1)
    `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:96:4
	{
		_n, _err := _template.writeValue(
			_output,
			(AcceptAction),
			"internal_definitions.template:96:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:96:17
	{
		_n, _err := _output.Write([]byte(` = `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:96:20
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionIdType),
			"internal_definitions.template:96:20")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:96:33
	{
		_n, _err := _output.Write([]byte(`(2)
)

func (i `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:99:8
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionIdType),
			"internal_definitions.template:99:8")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:99:21
	{
		_n, _err := _output.Write([]byte(`) String() string {
    switch i {
    case `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:101:9
	{
		_n, _err := _template.writeValue(
			_output,
			(ShiftAction),
			"internal_definitions.template:101:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:101:21
	{
		_n, _err := _output.Write([]byte(`:
        return "shift"
    case `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:103:9
	{
		_n, _err := _template.writeValue(
			_output,
			(ReduceAction),
			"internal_definitions.template:103:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:103:22
	{
		_n, _err := _output.Write([]byte(`:
        return "reduce"
    case `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:105:9
	{
		_n, _err := _template.writeValue(
			_output,
			(AcceptAction),
			"internal_definitions.template:105:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:105:22
	{
		_n, _err := _output.Write([]byte(`:
        return "accept"
    default:
        return `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:108:15
	{
		_n, _err := _template.writeValue(
			_output,
			(Sprintf),
			"internal_definitions.template:108:15")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:108:23
	{
		_n, _err := _output.Write([]byte(`("?Unknown action %d?", int(i))
    }
}

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:112:5
	{
		_n, _err := _template.writeValue(
			_output,
			(ReduceType),
			"internal_definitions.template:112:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:112:16
	{
		_n, _err := _output.Write([]byte(` int

const (`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:115:0
	clauseIdx := 1
	// internal_definitions.template:116:0
	for _, rule := range OrderedSymbols {
		// internal_definitions.template:117:4
		for _, clause := range rule.Clauses {
			// internal_definitions.template:117:44
			{
				_n, _err := _output.Write([]byte(`
    `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// internal_definitions.template:118:4
			{
				_n, _err := _template.writeValue(
					_output,
					(clause.CodeGenReducerNameConst),
					"internal_definitions.template:118:4")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// internal_definitions.template:118:37
			{
				_n, _err := _output.Write([]byte(` = `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// internal_definitions.template:118:40
			{
				_n, _err := _template.writeValue(
					_output,
					(ReduceType),
					"internal_definitions.template:118:40")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// internal_definitions.template:118:51
			{
				_n, _err := _output.Write([]byte(`(`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// internal_definitions.template:118:52
			{
				_n, _err := _template.writeValue(
					_output,
					(clauseIdx),
					"internal_definitions.template:118:52")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// internal_definitions.template:118:62
			{
				_n, _err := _output.Write([]byte(`)`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// internal_definitions.template:119:8
			clauseIdx += 1
		}
	}
	// internal_definitions.template:121:8
	{
		_n, _err := _output.Write([]byte(`
)

func (i `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:124:8
	{
		_n, _err := _template.writeValue(
			_output,
			(ReduceType),
			"internal_definitions.template:124:8")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:124:19
	{
		_n, _err := _output.Write([]byte(`) String() string {
    switch i {`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:126:0
	for _, rule := range OrderedSymbols {
		// internal_definitions.template:127:4
		for _, clause := range rule.Clauses {
			// internal_definitions.template:127:44
			{
				_n, _err := _output.Write([]byte(`
    case `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// internal_definitions.template:128:9
			{
				_n, _err := _template.writeValue(
					_output,
					(clause.CodeGenReducerNameConst),
					"internal_definitions.template:128:9")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// internal_definitions.template:128:42
			{
				_n, _err := _output.Write([]byte(`:
        return "`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// internal_definitions.template:129:16
			{
				_n, _err := _template.writeValue(
					_output,
					(clause.CodeGenReducerName),
					"internal_definitions.template:129:16")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// internal_definitions.template:129:44
			{
				_n, _err := _output.Write([]byte(`"`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
	}
	// internal_definitions.template:131:8
	{
		_n, _err := _output.Write([]byte(`
    default:
        return `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:133:15
	{
		_n, _err := _template.writeValue(
			_output,
			(Sprintf),
			"internal_definitions.template:133:15")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:133:23
	{
		_n, _err := _output.Write([]byte(`("?unknown reduce type %d?", int(i))
    }
}

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:137:5
	{
		_n, _err := _template.writeValue(
			_output,
			(StateIdType),
			"internal_definitions.template:137:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:137:17
	{
		_n, _err := _output.Write([]byte(` int

func (id `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:139:9
	{
		_n, _err := _template.writeValue(
			_output,
			(StateIdType),
			"internal_definitions.template:139:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:139:21
	{
		_n, _err := _output.Write([]byte(`) String() string {
    return `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:140:11
	{
		_n, _err := _template.writeValue(
			_output,
			(Sprintf),
			"internal_definitions.template:140:11")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:140:19
	{
		_n, _err := _output.Write([]byte(`("State %d", int(id))
}

const (`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:144:0
	for _, state := range OrderedStates {
		// internal_definitions.template:144:40
		{
			_n, _err := _output.Write([]byte(`
    `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// internal_definitions.template:145:4
		{
			_n, _err := _template.writeValue(
				_output,
				(state.CodeGenConst),
				"internal_definitions.template:145:4")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// internal_definitions.template:145:25
		{
			_n, _err := _output.Write([]byte(` = `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// internal_definitions.template:145:28
		{
			_n, _err := _template.writeValue(
				_output,
				(StateIdType),
				"internal_definitions.template:145:28")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// internal_definitions.template:145:40
		{
			_n, _err := _output.Write([]byte(`(`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// internal_definitions.template:145:41
		{
			_n, _err := _template.writeValue(
				_output,
				(state.StateNum),
				"internal_definitions.template:145:41")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// internal_definitions.template:145:58
		{
			_n, _err := _output.Write([]byte(`)`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// internal_definitions.template:146:8
	{
		_n, _err := _output.Write([]byte(`
)

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:149:5
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"internal_definitions.template:149:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:149:16
	{
		_n, _err := _output.Write([]byte(` struct {
    SymbolId_ `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:150:14
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"internal_definitions.template:150:14")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:150:27
	{
		_n, _err := _output.Write([]byte(`

    Generic_ *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:152:14
	{
		_n, _err := _template.writeValue(
			_output,
			(GenericSymbolType),
			"internal_definitions.template:152:14")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:152:32
	{
		_n, _err := _output.Write([]byte(`
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:154:0
	for _, valueType := range OrderedValueTypes {
		// internal_definitions.template:155:4
		if valueType.Name == lr.Generic {
			// internal_definitions.template:156:8
			continue
		}
		// internal_definitions.template:157:12
		{
			_n, _err := _output.Write([]byte(`
    `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// internal_definitions.template:158:4
		{
			_n, _err := _template.writeValue(
				_output,
				(valueType.Name),
				"internal_definitions.template:158:4")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// internal_definitions.template:158:21
		{
			_n, _err := _output.Write([]byte(` `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// internal_definitions.template:158:22
		{
			_n, _err := _template.writeValue(
				_output,
				(valueType.ParamType),
				"internal_definitions.template:158:22")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// internal_definitions.template:159:8
	{
		_n, _err := _output.Write([]byte(`
}
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:162:0

	valueTerms := map[string][]*lr.Term{}
	for _, term := range OrderedSymbols {
		if term.Name == lr.StartMarker ||
			term.Name == lr.Wildcard {

			continue
		}

		valueTerms[term.ValueType] = append(valueTerms[term.ValueType], term)
	}

	// internal_definitions.template:175:3
	{
		_n, _err := _output.Write([]byte(`func NewSymbol(token `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:176:21
	{
		_n, _err := _template.writeValue(
			_output,
			(TokenType),
			"internal_definitions.template:176:21")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:176:31
	{
		_n, _err := _output.Write([]byte(`) (*`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:176:35
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"internal_definitions.template:176:35")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:176:46
	{
		_n, _err := _output.Write([]byte(`, error) {
    symbol, ok := token.(*`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:177:26
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"internal_definitions.template:177:26")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:177:37
	{
		_n, _err := _output.Write([]byte(`)
    if ok {
        return symbol, nil
    }

    symbol = &`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:182:14
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"internal_definitions.template:182:14")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:182:25
	{
		_n, _err := _output.Write([]byte(`{SymbolId_: token.Id()}
    switch token.Id() {`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:184:0
	for _, valueType := range OrderedValueTypes {
		// internal_definitions.template:185:4

		consts := []string{}
		for _, term := range valueTerms[valueType.Name] {
			if !term.IsTerminal {
				break
			}

			consts = append(consts, term.CodeGenSymbolConst)
		}

		if len(consts) == 0 {
			continue
		}

		// internal_definitions.template:200:6
		{
			_n, _err := _output.Write([]byte(`
    case `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// internal_definitions.template:202:4
		for idx, kconst := range consts {
			// internal_definitions.template:203:0
			{
				_n, _err := _template.writeValue(
					_output,
					(kconst),
					"internal_definitions.template:203:0")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// internal_definitions.template:203:8
			if idx != len(consts)-1 {
				// internal_definitions.template:203:37
				{
					_n, _err := _output.Write([]byte(`, `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
			}
		}
		// internal_definitions.template:204:13
		{
			_n, _err := _output.Write([]byte(`:
        val, ok := token.(`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// internal_definitions.template:206:26
		{
			_n, _err := _template.writeValue(
				_output,
				(valueType.ParamType),
				"internal_definitions.template:206:26")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// internal_definitions.template:206:48
		{
			_n, _err := _output.Write([]byte(`)
        if !ok {
            return nil, `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// internal_definitions.template:208:24
		{
			_n, _err := _template.writeValue(
				_output,
				(Errorf),
				"internal_definitions.template:208:24")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// internal_definitions.template:208:31
		{
			_n, _err := _output.Write([]byte(`(
                "Invalid value type for token %s.  "+
                    "Expecting `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// internal_definitions.template:210:31
		{
			_n, _err := _template.writeValue(
				_output,
				(valueType.ParamType),
				"internal_definitions.template:210:31")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// internal_definitions.template:210:53
		{
			_n, _err := _output.Write([]byte(` (%v)",
                token.Id(),
                token.Loc())
        }
        symbol.`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// internal_definitions.template:214:15
		{
			_n, _err := _template.writeValue(
				_output,
				(valueType.Name),
				"internal_definitions.template:214:15")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// internal_definitions.template:214:32
		{
			_n, _err := _output.Write([]byte(` = val`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// internal_definitions.template:215:8
	{
		_n, _err := _output.Write([]byte(`
    default:
        return nil, `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:217:20
	{
		_n, _err := _template.writeValue(
			_output,
			(Errorf),
			"internal_definitions.template:217:20")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:217:27
	{
		_n, _err := _output.Write([]byte(`("Unexpected token type: %s", symbol.Id())
    }
    return symbol, nil
}

func (s *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:222:9
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"internal_definitions.template:222:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:222:20
	{
		_n, _err := _output.Write([]byte(`) Id() `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:222:27
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"internal_definitions.template:222:27")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:222:40
	{
		_n, _err := _output.Write([]byte(` {
    return s.SymbolId_
}

func (s *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:226:9
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"internal_definitions.template:226:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:226:20
	{
		_n, _err := _output.Write([]byte(`) Loc() `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:226:28
	{
		_n, _err := _template.writeValue(
			_output,
			(LocationType),
			"internal_definitions.template:226:28")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:226:41
	{
		_n, _err := _output.Write([]byte(` {
    type locator interface { Loc() `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:227:35
	{
		_n, _err := _template.writeValue(
			_output,
			(LocationType),
			"internal_definitions.template:227:35")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:227:48
	{
		_n, _err := _output.Write([]byte(` }
    switch s.SymbolId_ {`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:229:0
	for _, field := range OrderedValueTypes {
		// internal_definitions.template:230:4
		if field.Name == lr.Generic {
			// internal_definitions.template:231:8
			continue
		}
		// internal_definitions.template:233:4
		terms := valueTerms[field.Name]
		// internal_definitions.template:233:42
		{
			_n, _err := _output.Write([]byte(`
    case `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// internal_definitions.template:235:4
		for idx, term := range terms {
			// internal_definitions.template:236:0
			{
				_n, _err := _template.writeValue(
					_output,
					(term.CodeGenSymbolConst),
					"internal_definitions.template:236:0")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// internal_definitions.template:237:8
			if idx != len(terms)-1 {
				// internal_definitions.template:237:37
				{
					_n, _err := _output.Write([]byte(`, `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
			}
		}
		// internal_definitions.template:238:13
		{
			_n, _err := _output.Write([]byte(`:
        loc, ok := interface{}(s.`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// internal_definitions.template:240:33
		{
			_n, _err := _template.writeValue(
				_output,
				(field.Name),
				"internal_definitions.template:240:33")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// internal_definitions.template:240:46
		{
			_n, _err := _output.Write([]byte(`).(locator)
        if ok {
            return loc.Loc()
        }`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// internal_definitions.template:244:8
	{
		_n, _err := _output.Write([]byte(`
    }
    if s.Generic_ != nil {
        return s.Generic_.Loc()
    }
    return `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:249:11
	{
		_n, _err := _template.writeValue(
			_output,
			(LocationType),
			"internal_definitions.template:249:11")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:249:24
	{
		_n, _err := _output.Write([]byte(`{}
}

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:252:5
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolStackType),
			"internal_definitions.template:252:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:252:21
	{
		_n, _err := _output.Write([]byte(` struct {
    lexer `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:253:10
	{
		_n, _err := _template.writeValue(
			_output,
			(LexerType),
			"internal_definitions.template:253:10")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:253:20
	{
		_n, _err := _output.Write([]byte(`
    top []*`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:254:11
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"internal_definitions.template:254:11")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:254:22
	{
		_n, _err := _output.Write([]byte(`
}

func (stack *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:257:13
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolStackType),
			"internal_definitions.template:257:13")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:257:29
	{
		_n, _err := _output.Write([]byte(`) Top() (*`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:257:39
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"internal_definitions.template:257:39")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:257:50
	{
		_n, _err := _output.Write([]byte(`, error) {
    if len(stack.top) == 0 {
        token, err := stack.lexer.Next()
        if err != nil {
            if err != `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:261:22
	{
		_n, _err := _template.writeValue(
			_output,
			(EOF),
			"internal_definitions.template:261:22")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:261:26
	{
		_n, _err := _output.Write([]byte(` {
                return nil, `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:262:28
	{
		_n, _err := _template.writeValue(
			_output,
			(Errorf),
			"internal_definitions.template:262:28")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:262:35
	{
		_n, _err := _output.Write([]byte(`("Unexpected lex error: %s", err)
            }
            token = &`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:264:21
	{
		_n, _err := _template.writeValue(
			_output,
			(GenericSymbolType),
			"internal_definitions.template:264:21")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:264:39
	{
		_n, _err := _output.Write([]byte(`{`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:264:40
	{
		_n, _err := _template.writeValue(
			_output,
			(EndSymbolId),
			"internal_definitions.template:264:40")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:264:52
	{
		_n, _err := _output.Write([]byte(`, stack.lexer.CurrentLocation()}
        }
        item, err := NewSymbol(token)
        if err != nil {
            return nil, err
        }
        stack.top = append(stack.top, item)
    }
    return stack.top[len(stack.top)-1], nil
}

func (stack *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:275:13
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolStackType),
			"internal_definitions.template:275:13")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:275:29
	{
		_n, _err := _output.Write([]byte(`) Push(symbol *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:275:44
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"internal_definitions.template:275:44")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:275:55
	{
		_n, _err := _output.Write([]byte(`) {
    stack.top = append(stack.top, symbol)
}

func (stack *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:279:13
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolStackType),
			"internal_definitions.template:279:13")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:279:29
	{
		_n, _err := _output.Write([]byte(`) Pop() (*`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:279:39
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"internal_definitions.template:279:39")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:279:50
	{
		_n, _err := _output.Write([]byte(`, error) {
    if len(stack.top) == 0 {
        return nil, `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:281:20
	{
		_n, _err := _template.writeValue(
			_output,
			(Errorf),
			"internal_definitions.template:281:20")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:281:27
	{
		_n, _err := _output.Write([]byte(`("internal error: cannot pop an empty top")
    }
    ret := stack.top[len(stack.top)-1]
    stack.top = stack.top[:len(stack.top)-1]
    return ret, nil
}

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:288:5
	{
		_n, _err := _template.writeValue(
			_output,
			(StackItemType),
			"internal_definitions.template:288:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:288:19
	{
		_n, _err := _output.Write([]byte(` struct {
    StateId `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:289:12
	{
		_n, _err := _template.writeValue(
			_output,
			(StateIdType),
			"internal_definitions.template:289:12")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:289:24
	{
		_n, _err := _output.Write([]byte(`

    *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:291:5
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"internal_definitions.template:291:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:291:16
	{
		_n, _err := _output.Write([]byte(`
}

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:294:5
	{
		_n, _err := _template.writeValue(
			_output,
			(StackType),
			"internal_definitions.template:294:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:294:15
	{
		_n, _err := _output.Write([]byte(` []*`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:294:19
	{
		_n, _err := _template.writeValue(
			_output,
			(StackItemType),
			"internal_definitions.template:294:19")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:294:33
	{
		_n, _err := _output.Write([]byte(`

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:296:5
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionType),
			"internal_definitions.template:296:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:296:16
	{
		_n, _err := _output.Write([]byte(` struct {
    ActionType `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:297:15
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionIdType),
			"internal_definitions.template:297:15")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:297:28
	{
		_n, _err := _output.Write([]byte(`

    ShiftStateId `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:299:17
	{
		_n, _err := _template.writeValue(
			_output,
			(StateIdType),
			"internal_definitions.template:299:17")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:299:29
	{
		_n, _err := _output.Write([]byte(`
    ReduceType `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:300:15
	{
		_n, _err := _template.writeValue(
			_output,
			(ReduceType),
			"internal_definitions.template:300:15")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:300:26
	{
		_n, _err := _output.Write([]byte(`
}

func (act *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:303:11
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionType),
			"internal_definitions.template:303:11")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:303:22
	{
		_n, _err := _output.Write([]byte(`) ShiftItem(symbol *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:303:42
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"internal_definitions.template:303:42")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:303:53
	{
		_n, _err := _output.Write([]byte(`) *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:303:56
	{
		_n, _err := _template.writeValue(
			_output,
			(StackItemType),
			"internal_definitions.template:303:56")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:303:70
	{
		_n, _err := _output.Write([]byte(` {
    return &`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:304:12
	{
		_n, _err := _template.writeValue(
			_output,
			(StackItemType),
			"internal_definitions.template:304:12")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:304:26
	{
		_n, _err := _output.Write([]byte(`{StateId: act.ShiftStateId, `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:304:54
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"internal_definitions.template:304:54")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:304:65
	{
		_n, _err := _output.Write([]byte(`: symbol}
}

func (act *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:307:11
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionType),
			"internal_definitions.template:307:11")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:307:22
	{
		_n, _err := _output.Write([]byte(`) ReduceSymbol(
    reducer `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:308:12
	{
		_n, _err := _template.writeValue(
			_output,
			(ReducerType),
			"internal_definitions.template:308:12")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:308:24
	{
		_n, _err := _output.Write([]byte(`,
    stack `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:309:10
	{
		_n, _err := _template.writeValue(
			_output,
			(StackType),
			"internal_definitions.template:309:10")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:309:20
	{
		_n, _err := _output.Write([]byte(`) (
    `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:310:4
	{
		_n, _err := _template.writeValue(
			_output,
			(StackType),
			"internal_definitions.template:310:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:310:14
	{
		_n, _err := _output.Write([]byte(`,
    *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:311:5
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"internal_definitions.template:311:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:311:16
	{
		_n, _err := _output.Write([]byte(`,
    error) {

    var err error
    symbol := &`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:315:15
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"internal_definitions.template:315:15")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:315:26
	{
		_n, _err := _output.Write([]byte(`{}
    switch act.ReduceType {`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:317:0
	for _, rule := range OrderedSymbols {
		// internal_definitions.template:318:4
		for _, clause := range rule.Clauses {
			// internal_definitions.template:318:44
			{
				_n, _err := _output.Write([]byte(`
    case `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// internal_definitions.template:319:9
			{
				_n, _err := _template.writeValue(
					_output,
					(clause.CodeGenReducerNameConst),
					"internal_definitions.template:319:9")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// internal_definitions.template:319:42
			{
				_n, _err := _output.Write([]byte(`:`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// internal_definitions.template:320:8
			if len(clause.Bindings) > 0 {
				// internal_definitions.template:320:40
				{
					_n, _err := _output.Write([]byte(`
        args := stack[len(stack)-`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// internal_definitions.template:321:33
				{
					_n, _err := _template.writeValue(
						_output,
						(len(clause.Bindings)),
						"internal_definitions.template:321:33")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// internal_definitions.template:321:56
				{
					_n, _err := _output.Write([]byte(`:]
        stack = stack[:len(stack)-`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// internal_definitions.template:322:34
				{
					_n, _err := _template.writeValue(
						_output,
						(len(clause.Bindings)),
						"internal_definitions.template:322:34")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// internal_definitions.template:322:57
				{
					_n, _err := _output.Write([]byte(`]`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
			}
			// internal_definitions.template:323:16
			{
				_n, _err := _output.Write([]byte(`
        symbol.SymbolId_ = `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// internal_definitions.template:324:27
			{
				_n, _err := _template.writeValue(
					_output,
					(rule.CodeGenSymbolConst),
					"internal_definitions.template:324:27")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// internal_definitions.template:324:53
			{
				_n, _err := _output.Write([]byte(`
        symbol.`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// internal_definitions.template:325:15
			{
				_n, _err := _template.writeValue(
					_output,
					(rule.ValueType),
					"internal_definitions.template:325:15")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// internal_definitions.template:325:32
			{
				_n, _err := _output.Write([]byte(`, err = reducer.`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// internal_definitions.template:325:48
			{
				_n, _err := _template.writeValue(
					_output,
					(clause.CodeGenReducerName),
					"internal_definitions.template:325:48")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// internal_definitions.template:325:76
			{
				_n, _err := _output.Write([]byte(`(`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// internal_definitions.template:326:8
			for idx, term := range clause.Bindings {
				// internal_definitions.template:326:52
				{
					_n, _err := _output.Write([]byte(`args[`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// internal_definitions.template:327:5
				{
					_n, _err := _template.writeValue(
						_output,
						(idx),
						"internal_definitions.template:327:5")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// internal_definitions.template:327:9
				{
					_n, _err := _output.Write([]byte(`].`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// internal_definitions.template:327:11
				{
					_n, _err := _template.writeValue(
						_output,
						(term.ValueType),
						"internal_definitions.template:327:11")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// internal_definitions.template:328:12
				if idx != len(clause.Bindings)-1 {
					// internal_definitions.template:328:49
					{
						_n, _err := _output.Write([]byte(`,`))
						_numWritten += int64(_n)
						if _err != nil {
							return _numWritten, _err
						}
					}
				}
			}
			// internal_definitions.template:329:17
			{
				_n, _err := _output.Write([]byte(`)`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
	}
	// internal_definitions.template:332:8
	{
		_n, _err := _output.Write([]byte(`
    default:
        panic("Unknown reduce type: " + act.ReduceType.String())
    }

    if err != nil {
        err = `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:338:14
	{
		_n, _err := _template.writeValue(
			_output,
			(Errorf),
			"internal_definitions.template:338:14")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:338:21
	{
		_n, _err := _output.Write([]byte(`("Unexpected %s reduce error: %s", act.ReduceType, err)
    }

    return stack, symbol, err
}
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}

	return _numWritten, nil
}
