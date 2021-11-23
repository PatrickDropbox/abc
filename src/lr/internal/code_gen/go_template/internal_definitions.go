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

	// internal_definitions.template:47:0
	{
		_n, _err := _output.Write([]byte(`func (i `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:47:8
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"internal_definitions.template:47:8")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:47:21
	{
		_n, _err := _output.Write([]byte(`) String() string {
    switch i {
    case `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:49:9
	{
		_n, _err := _template.writeValue(
			_output,
			(EndSymbolId),
			"internal_definitions.template:49:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:49:21
	{
		_n, _err := _output.Write([]byte(`:
        return "$"
    case `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:51:9
	{
		_n, _err := _template.writeValue(
			_output,
			(WildcardSymbolId),
			"internal_definitions.template:51:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:51:26
	{
		_n, _err := _output.Write([]byte(`:
        return "*"`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:53:0
	for _, term := range OrderedSymbols[3:] {
		// internal_definitions.template:54:4
		if term.SymbolId == parser.LRCharacterToken {
			// internal_definitions.template:55:8

			escaped := term.Name
			if term.Name == "'\"'" {
				escaped = "'\\\"'"
			} else if escaped[1] == '\\' {
				escaped = "'\\\\" + term.Name[2:]
			}

			// internal_definitions.template:64:10
			{
				_n, _err := _output.Write([]byte(`
    case `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// internal_definitions.template:65:9
			{
				_n, _err := _template.writeValue(
					_output,
					(term.CodeGenSymbolConst),
					"internal_definitions.template:65:9")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// internal_definitions.template:65:35
			{
				_n, _err := _output.Write([]byte(`:
        return "`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// internal_definitions.template:66:16
			{
				_n, _err := _template.writeValue(
					_output,
					(escaped),
					"internal_definitions.template:66:16")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// internal_definitions.template:66:24
			{
				_n, _err := _output.Write([]byte(`"`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		} else {
			// internal_definitions.template:67:13
			{
				_n, _err := _output.Write([]byte(`
    case `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// internal_definitions.template:68:9
			{
				_n, _err := _template.writeValue(
					_output,
					(term.CodeGenSymbolConst),
					"internal_definitions.template:68:9")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// internal_definitions.template:68:35
			{
				_n, _err := _output.Write([]byte(`:
        return "`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// internal_definitions.template:69:16
			{
				_n, _err := _template.writeValue(
					_output,
					(term.Name),
					"internal_definitions.template:69:16")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// internal_definitions.template:69:28
			{
				_n, _err := _output.Write([]byte(`"`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
	}
	// internal_definitions.template:71:8
	{
		_n, _err := _output.Write([]byte(`
    default:
        return `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:73:15
	{
		_n, _err := _template.writeValue(
			_output,
			(Sprintf),
			"internal_definitions.template:73:15")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:73:23
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
	// internal_definitions.template:78:4
	{
		_n, _err := _template.writeValue(
			_output,
			(EndSymbolId),
			"internal_definitions.template:78:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:78:16
	{
		_n, _err := _output.Write([]byte(` = `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:78:19
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"internal_definitions.template:78:19")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:78:32
	{
		_n, _err := _output.Write([]byte(`(0)
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
			(WildcardSymbolId),
			"internal_definitions.template:79:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:79:21
	{
		_n, _err := _output.Write([]byte(` = `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:79:24
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"internal_definitions.template:79:24")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:79:37
	{
		_n, _err := _output.Write([]byte(`(-1)
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:81:0
	for idx, term := range OrderedSymbols[3:] {
		// internal_definitions.template:82:4
		if term.IsTerminal {
			// internal_definitions.template:83:8
			continue
		}
		// internal_definitions.template:84:12
		{
			_n, _err := _output.Write([]byte(`
    `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// internal_definitions.template:85:4
		{
			_n, _err := _template.writeValue(
				_output,
				(term.CodeGenSymbolConst),
				"internal_definitions.template:85:4")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// internal_definitions.template:85:30
		{
			_n, _err := _output.Write([]byte(` = `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// internal_definitions.template:85:33
		{
			_n, _err := _template.writeValue(
				_output,
				(SymbolIdType),
				"internal_definitions.template:85:33")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// internal_definitions.template:85:46
		{
			_n, _err := _output.Write([]byte(`(`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// internal_definitions.template:85:47
		{
			_n, _err := _template.writeValue(
				_output,
				(256 + idx),
				"internal_definitions.template:85:47")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// internal_definitions.template:85:57
		{
			_n, _err := _output.Write([]byte(`)`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// internal_definitions.template:86:8
	{
		_n, _err := _output.Write([]byte(`
)

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:89:5
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionIdType),
			"internal_definitions.template:89:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:89:18
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
	// internal_definitions.template:93:4
	{
		_n, _err := _template.writeValue(
			_output,
			(ShiftAction),
			"internal_definitions.template:93:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:93:16
	{
		_n, _err := _output.Write([]byte(` = `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:93:19
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionIdType),
			"internal_definitions.template:93:19")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:93:32
	{
		_n, _err := _output.Write([]byte(`(0)
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
			(ReduceAction),
			"internal_definitions.template:94:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:94:17
	{
		_n, _err := _output.Write([]byte(` = `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:94:20
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionIdType),
			"internal_definitions.template:94:20")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:94:33
	{
		_n, _err := _output.Write([]byte(`(1)
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
			(AcceptAction),
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
		_n, _err := _output.Write([]byte(`(2)
)

func (i `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:98:8
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionIdType),
			"internal_definitions.template:98:8")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:98:21
	{
		_n, _err := _output.Write([]byte(`) String() string {
    switch i {
    case `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:100:9
	{
		_n, _err := _template.writeValue(
			_output,
			(ShiftAction),
			"internal_definitions.template:100:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:100:21
	{
		_n, _err := _output.Write([]byte(`:
        return "shift"
    case `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:102:9
	{
		_n, _err := _template.writeValue(
			_output,
			(ReduceAction),
			"internal_definitions.template:102:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:102:22
	{
		_n, _err := _output.Write([]byte(`:
        return "reduce"
    case `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:104:9
	{
		_n, _err := _template.writeValue(
			_output,
			(AcceptAction),
			"internal_definitions.template:104:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:104:22
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
	// internal_definitions.template:107:15
	{
		_n, _err := _template.writeValue(
			_output,
			(Sprintf),
			"internal_definitions.template:107:15")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:107:23
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
	// internal_definitions.template:111:5
	{
		_n, _err := _template.writeValue(
			_output,
			(ReduceType),
			"internal_definitions.template:111:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:111:16
	{
		_n, _err := _output.Write([]byte(` int

const (`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:114:0
	clauseIdx := 1
	// internal_definitions.template:115:0
	for _, rule := range OrderedSymbols {
		// internal_definitions.template:116:4
		for _, clause := range rule.Clauses {
			// internal_definitions.template:116:44
			{
				_n, _err := _output.Write([]byte(`
    `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// internal_definitions.template:117:4
			{
				_n, _err := _template.writeValue(
					_output,
					(clause.CodeGenReducerNameConst),
					"internal_definitions.template:117:4")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// internal_definitions.template:117:37
			{
				_n, _err := _output.Write([]byte(` = `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// internal_definitions.template:117:40
			{
				_n, _err := _template.writeValue(
					_output,
					(ReduceType),
					"internal_definitions.template:117:40")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// internal_definitions.template:117:51
			{
				_n, _err := _output.Write([]byte(`(`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// internal_definitions.template:117:52
			{
				_n, _err := _template.writeValue(
					_output,
					(clauseIdx),
					"internal_definitions.template:117:52")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// internal_definitions.template:117:62
			{
				_n, _err := _output.Write([]byte(`)`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// internal_definitions.template:118:8
			clauseIdx += 1
		}
	}
	// internal_definitions.template:120:8
	{
		_n, _err := _output.Write([]byte(`
)

func (i `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:123:8
	{
		_n, _err := _template.writeValue(
			_output,
			(ReduceType),
			"internal_definitions.template:123:8")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:123:19
	{
		_n, _err := _output.Write([]byte(`) String() string {
    switch i {`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:125:0
	for _, rule := range OrderedSymbols {
		// internal_definitions.template:126:4
		for _, clause := range rule.Clauses {
			// internal_definitions.template:126:44
			{
				_n, _err := _output.Write([]byte(`
    case `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// internal_definitions.template:127:9
			{
				_n, _err := _template.writeValue(
					_output,
					(clause.CodeGenReducerNameConst),
					"internal_definitions.template:127:9")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// internal_definitions.template:127:42
			{
				_n, _err := _output.Write([]byte(`:
        return "`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// internal_definitions.template:128:16
			{
				_n, _err := _template.writeValue(
					_output,
					(clause.CodeGenReducerName),
					"internal_definitions.template:128:16")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// internal_definitions.template:128:44
			{
				_n, _err := _output.Write([]byte(`"`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
	}
	// internal_definitions.template:130:8
	{
		_n, _err := _output.Write([]byte(`
    default:
        return `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:132:15
	{
		_n, _err := _template.writeValue(
			_output,
			(Sprintf),
			"internal_definitions.template:132:15")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:132:23
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
	// internal_definitions.template:136:5
	{
		_n, _err := _template.writeValue(
			_output,
			(StateIdType),
			"internal_definitions.template:136:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:136:17
	{
		_n, _err := _output.Write([]byte(` int

func (id `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:138:9
	{
		_n, _err := _template.writeValue(
			_output,
			(StateIdType),
			"internal_definitions.template:138:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:138:21
	{
		_n, _err := _output.Write([]byte(`) String() string {
    return `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:139:11
	{
		_n, _err := _template.writeValue(
			_output,
			(Sprintf),
			"internal_definitions.template:139:11")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:139:19
	{
		_n, _err := _output.Write([]byte(`("State %d", int(id))
}

const (`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:143:0
	for _, state := range OrderedStates {
		// internal_definitions.template:143:40
		{
			_n, _err := _output.Write([]byte(`
    `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// internal_definitions.template:144:4
		{
			_n, _err := _template.writeValue(
				_output,
				(state.CodeGenConst),
				"internal_definitions.template:144:4")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// internal_definitions.template:144:25
		{
			_n, _err := _output.Write([]byte(` = `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// internal_definitions.template:144:28
		{
			_n, _err := _template.writeValue(
				_output,
				(StateIdType),
				"internal_definitions.template:144:28")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// internal_definitions.template:144:40
		{
			_n, _err := _output.Write([]byte(`(`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// internal_definitions.template:144:41
		{
			_n, _err := _template.writeValue(
				_output,
				(state.StateNum),
				"internal_definitions.template:144:41")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// internal_definitions.template:144:58
		{
			_n, _err := _output.Write([]byte(`)`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// internal_definitions.template:145:8
	{
		_n, _err := _output.Write([]byte(`
)

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:148:5
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"internal_definitions.template:148:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:148:16
	{
		_n, _err := _output.Write([]byte(` struct {
    SymbolId_ `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:149:14
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"internal_definitions.template:149:14")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:149:27
	{
		_n, _err := _output.Write([]byte(`

    Generic_ *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:151:14
	{
		_n, _err := _template.writeValue(
			_output,
			(GenericSymbolType),
			"internal_definitions.template:151:14")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:151:32
	{
		_n, _err := _output.Write([]byte(`
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:153:0
	for _, valueType := range OrderedValueTypes {
		// internal_definitions.template:154:4
		if valueType.Name == lr.Generic {
			// internal_definitions.template:155:8
			continue
		}
		// internal_definitions.template:156:12
		{
			_n, _err := _output.Write([]byte(`
    `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// internal_definitions.template:157:4
		{
			_n, _err := _template.writeValue(
				_output,
				(valueType.Name),
				"internal_definitions.template:157:4")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// internal_definitions.template:157:21
		{
			_n, _err := _output.Write([]byte(` `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// internal_definitions.template:157:22
		{
			_n, _err := _template.writeValue(
				_output,
				(valueType.ParamType),
				"internal_definitions.template:157:22")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// internal_definitions.template:158:8
	{
		_n, _err := _output.Write([]byte(`
}
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:161:0

	valueTerms := map[string][]*lr.Term{}
	for _, term := range OrderedSymbols {
		if term.Name == lr.StartMarker ||
			term.Name == lr.Wildcard {

			continue
		}

		valueTerms[term.ValueType] = append(valueTerms[term.ValueType], term)
	}

	// internal_definitions.template:174:3
	{
		_n, _err := _output.Write([]byte(`func NewSymbol(token `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:175:21
	{
		_n, _err := _template.writeValue(
			_output,
			(TokenType),
			"internal_definitions.template:175:21")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:175:31
	{
		_n, _err := _output.Write([]byte(`) (*`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:175:35
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"internal_definitions.template:175:35")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:175:46
	{
		_n, _err := _output.Write([]byte(`, error) {
    symbol, ok := token.(*`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:176:26
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"internal_definitions.template:176:26")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:176:37
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
	// internal_definitions.template:181:14
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"internal_definitions.template:181:14")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:181:25
	{
		_n, _err := _output.Write([]byte(`{SymbolId_: token.Id()}
    switch token.Id() {`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:183:0
	for _, valueType := range OrderedValueTypes {
		// internal_definitions.template:184:4

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

		// internal_definitions.template:199:6
		{
			_n, _err := _output.Write([]byte(`
    case `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// internal_definitions.template:201:4
		for idx, kconst := range consts {
			// internal_definitions.template:202:0
			{
				_n, _err := _template.writeValue(
					_output,
					(kconst),
					"internal_definitions.template:202:0")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// internal_definitions.template:202:8
			if idx != len(consts)-1 {
				// internal_definitions.template:202:37
				{
					_n, _err := _output.Write([]byte(`, `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
			}
		}
		// internal_definitions.template:203:13
		{
			_n, _err := _output.Write([]byte(`:
        val, ok := token.(`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// internal_definitions.template:205:26
		{
			_n, _err := _template.writeValue(
				_output,
				(valueType.ParamType),
				"internal_definitions.template:205:26")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// internal_definitions.template:205:48
		{
			_n, _err := _output.Write([]byte(`)
        if !ok {
            return nil, `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// internal_definitions.template:207:24
		{
			_n, _err := _template.writeValue(
				_output,
				(Errorf),
				"internal_definitions.template:207:24")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// internal_definitions.template:207:31
		{
			_n, _err := _output.Write([]byte(`(
                "Invalid value type for token %s.  "+
                    "Expecting `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// internal_definitions.template:209:31
		{
			_n, _err := _template.writeValue(
				_output,
				(valueType.ParamType),
				"internal_definitions.template:209:31")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// internal_definitions.template:209:53
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
		// internal_definitions.template:213:15
		{
			_n, _err := _template.writeValue(
				_output,
				(valueType.Name),
				"internal_definitions.template:213:15")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// internal_definitions.template:213:32
		{
			_n, _err := _output.Write([]byte(` = val`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// internal_definitions.template:214:8
	{
		_n, _err := _output.Write([]byte(`
    default:
        return nil, `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:216:20
	{
		_n, _err := _template.writeValue(
			_output,
			(Errorf),
			"internal_definitions.template:216:20")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:216:27
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
	// internal_definitions.template:221:9
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"internal_definitions.template:221:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:221:20
	{
		_n, _err := _output.Write([]byte(`) Id() `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:221:27
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"internal_definitions.template:221:27")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:221:40
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
	// internal_definitions.template:225:9
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"internal_definitions.template:225:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:225:20
	{
		_n, _err := _output.Write([]byte(`) Loc() `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:225:28
	{
		_n, _err := _template.writeValue(
			_output,
			(LocationType),
			"internal_definitions.template:225:28")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:225:41
	{
		_n, _err := _output.Write([]byte(` {
    type locator interface { Loc() `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:226:35
	{
		_n, _err := _template.writeValue(
			_output,
			(LocationType),
			"internal_definitions.template:226:35")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:226:48
	{
		_n, _err := _output.Write([]byte(` }
    switch s.SymbolId_ {`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:228:0
	for _, field := range OrderedValueTypes {
		// internal_definitions.template:229:4
		if field.Name == lr.Generic {
			// internal_definitions.template:230:8
			continue
		}
		// internal_definitions.template:232:4
		terms := valueTerms[field.Name]
		// internal_definitions.template:232:42
		{
			_n, _err := _output.Write([]byte(`
    case `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// internal_definitions.template:234:4
		for idx, term := range terms {
			// internal_definitions.template:235:0
			{
				_n, _err := _template.writeValue(
					_output,
					(term.CodeGenSymbolConst),
					"internal_definitions.template:235:0")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// internal_definitions.template:236:8
			if idx != len(terms)-1 {
				// internal_definitions.template:236:37
				{
					_n, _err := _output.Write([]byte(`, `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
			}
		}
		// internal_definitions.template:237:13
		{
			_n, _err := _output.Write([]byte(`:
        loc, ok := interface{}(s.`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// internal_definitions.template:239:33
		{
			_n, _err := _template.writeValue(
				_output,
				(field.Name),
				"internal_definitions.template:239:33")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// internal_definitions.template:239:46
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
	// internal_definitions.template:243:8
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
	// internal_definitions.template:248:11
	{
		_n, _err := _template.writeValue(
			_output,
			(LocationType),
			"internal_definitions.template:248:11")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:248:24
	{
		_n, _err := _output.Write([]byte(`{}
}

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:251:5
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolStackType),
			"internal_definitions.template:251:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:251:21
	{
		_n, _err := _output.Write([]byte(` struct {
    lexer `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:252:10
	{
		_n, _err := _template.writeValue(
			_output,
			(LexerType),
			"internal_definitions.template:252:10")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:252:20
	{
		_n, _err := _output.Write([]byte(`
    top []*`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:253:11
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"internal_definitions.template:253:11")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:253:22
	{
		_n, _err := _output.Write([]byte(`
}

func (stack *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:256:13
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolStackType),
			"internal_definitions.template:256:13")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:256:29
	{
		_n, _err := _output.Write([]byte(`) Top() (*`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:256:39
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"internal_definitions.template:256:39")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:256:50
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
	// internal_definitions.template:260:22
	{
		_n, _err := _template.writeValue(
			_output,
			(EOF),
			"internal_definitions.template:260:22")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:260:26
	{
		_n, _err := _output.Write([]byte(` {
                return nil, `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:261:28
	{
		_n, _err := _template.writeValue(
			_output,
			(Errorf),
			"internal_definitions.template:261:28")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:261:35
	{
		_n, _err := _output.Write([]byte(`("Unexpected lex error: %s", err)
            }
            token = &`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:263:21
	{
		_n, _err := _template.writeValue(
			_output,
			(GenericSymbolType),
			"internal_definitions.template:263:21")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:263:39
	{
		_n, _err := _output.Write([]byte(`{`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:263:40
	{
		_n, _err := _template.writeValue(
			_output,
			(EndSymbolId),
			"internal_definitions.template:263:40")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:263:52
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
	// internal_definitions.template:274:13
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolStackType),
			"internal_definitions.template:274:13")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:274:29
	{
		_n, _err := _output.Write([]byte(`) Push(symbol *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:274:44
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"internal_definitions.template:274:44")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:274:55
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
	// internal_definitions.template:278:13
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolStackType),
			"internal_definitions.template:278:13")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:278:29
	{
		_n, _err := _output.Write([]byte(`) Pop() (*`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:278:39
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"internal_definitions.template:278:39")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:278:50
	{
		_n, _err := _output.Write([]byte(`, error) {
    if len(stack.top) == 0 {
        return nil, `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:280:20
	{
		_n, _err := _template.writeValue(
			_output,
			(Errorf),
			"internal_definitions.template:280:20")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:280:27
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
	// internal_definitions.template:287:5
	{
		_n, _err := _template.writeValue(
			_output,
			(StackItemType),
			"internal_definitions.template:287:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:287:19
	{
		_n, _err := _output.Write([]byte(` struct {
    StateId `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:288:12
	{
		_n, _err := _template.writeValue(
			_output,
			(StateIdType),
			"internal_definitions.template:288:12")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:288:24
	{
		_n, _err := _output.Write([]byte(`

    *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:290:5
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"internal_definitions.template:290:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:290:16
	{
		_n, _err := _output.Write([]byte(`
}

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:293:5
	{
		_n, _err := _template.writeValue(
			_output,
			(StackType),
			"internal_definitions.template:293:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:293:15
	{
		_n, _err := _output.Write([]byte(` []*`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:293:19
	{
		_n, _err := _template.writeValue(
			_output,
			(StackItemType),
			"internal_definitions.template:293:19")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:293:33
	{
		_n, _err := _output.Write([]byte(`

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:295:5
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionType),
			"internal_definitions.template:295:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:295:16
	{
		_n, _err := _output.Write([]byte(` struct {
    ActionType `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:296:15
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionIdType),
			"internal_definitions.template:296:15")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:296:28
	{
		_n, _err := _output.Write([]byte(`

    ShiftStateId `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:298:17
	{
		_n, _err := _template.writeValue(
			_output,
			(StateIdType),
			"internal_definitions.template:298:17")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:298:29
	{
		_n, _err := _output.Write([]byte(`
    ReduceType `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:299:15
	{
		_n, _err := _template.writeValue(
			_output,
			(ReduceType),
			"internal_definitions.template:299:15")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:299:26
	{
		_n, _err := _output.Write([]byte(`
}

func (act *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:302:11
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionType),
			"internal_definitions.template:302:11")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:302:22
	{
		_n, _err := _output.Write([]byte(`) ShiftItem(symbol *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:302:42
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"internal_definitions.template:302:42")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:302:53
	{
		_n, _err := _output.Write([]byte(`) *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:302:56
	{
		_n, _err := _template.writeValue(
			_output,
			(StackItemType),
			"internal_definitions.template:302:56")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:302:70
	{
		_n, _err := _output.Write([]byte(` {
    return &`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:303:12
	{
		_n, _err := _template.writeValue(
			_output,
			(StackItemType),
			"internal_definitions.template:303:12")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:303:26
	{
		_n, _err := _output.Write([]byte(`{StateId: act.ShiftStateId, `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:303:54
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"internal_definitions.template:303:54")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:303:65
	{
		_n, _err := _output.Write([]byte(`: symbol}
}

func (act *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:306:11
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionType),
			"internal_definitions.template:306:11")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:306:22
	{
		_n, _err := _output.Write([]byte(`) ReduceSymbol(
    reducer `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:307:12
	{
		_n, _err := _template.writeValue(
			_output,
			(ReducerType),
			"internal_definitions.template:307:12")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:307:24
	{
		_n, _err := _output.Write([]byte(`,
    stack `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:308:10
	{
		_n, _err := _template.writeValue(
			_output,
			(StackType),
			"internal_definitions.template:308:10")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:308:20
	{
		_n, _err := _output.Write([]byte(`) (
    `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:309:4
	{
		_n, _err := _template.writeValue(
			_output,
			(StackType),
			"internal_definitions.template:309:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:309:14
	{
		_n, _err := _output.Write([]byte(`,
    *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:310:5
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"internal_definitions.template:310:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:310:16
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
	// internal_definitions.template:314:15
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"internal_definitions.template:314:15")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:314:26
	{
		_n, _err := _output.Write([]byte(`{}
    switch act.ReduceType {`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:316:0
	for _, rule := range OrderedSymbols {
		// internal_definitions.template:317:4
		for _, clause := range rule.Clauses {
			// internal_definitions.template:317:44
			{
				_n, _err := _output.Write([]byte(`
    case `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// internal_definitions.template:318:9
			{
				_n, _err := _template.writeValue(
					_output,
					(clause.CodeGenReducerNameConst),
					"internal_definitions.template:318:9")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// internal_definitions.template:318:42
			{
				_n, _err := _output.Write([]byte(`:`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// internal_definitions.template:319:8
			if len(clause.Bindings) > 0 {
				// internal_definitions.template:319:40
				{
					_n, _err := _output.Write([]byte(`
        args := stack[len(stack)-`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// internal_definitions.template:320:33
				{
					_n, _err := _template.writeValue(
						_output,
						(len(clause.Bindings)),
						"internal_definitions.template:320:33")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// internal_definitions.template:320:56
				{
					_n, _err := _output.Write([]byte(`:]
        stack = stack[:len(stack)-`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// internal_definitions.template:321:34
				{
					_n, _err := _template.writeValue(
						_output,
						(len(clause.Bindings)),
						"internal_definitions.template:321:34")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// internal_definitions.template:321:57
				{
					_n, _err := _output.Write([]byte(`]`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
			}
			// internal_definitions.template:322:16
			{
				_n, _err := _output.Write([]byte(`
        symbol.SymbolId_ = `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// internal_definitions.template:323:27
			{
				_n, _err := _template.writeValue(
					_output,
					(rule.CodeGenSymbolConst),
					"internal_definitions.template:323:27")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// internal_definitions.template:323:53
			{
				_n, _err := _output.Write([]byte(`
        symbol.`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// internal_definitions.template:324:15
			{
				_n, _err := _template.writeValue(
					_output,
					(rule.ValueType),
					"internal_definitions.template:324:15")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// internal_definitions.template:324:32
			{
				_n, _err := _output.Write([]byte(`, err = reducer.`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// internal_definitions.template:324:48
			{
				_n, _err := _template.writeValue(
					_output,
					(clause.CodeGenReducerName),
					"internal_definitions.template:324:48")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// internal_definitions.template:324:76
			{
				_n, _err := _output.Write([]byte(`(`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// internal_definitions.template:325:8
			for idx, term := range clause.Bindings {
				// internal_definitions.template:325:52
				{
					_n, _err := _output.Write([]byte(`args[`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// internal_definitions.template:326:5
				{
					_n, _err := _template.writeValue(
						_output,
						(idx),
						"internal_definitions.template:326:5")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// internal_definitions.template:326:9
				{
					_n, _err := _output.Write([]byte(`].`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// internal_definitions.template:326:11
				{
					_n, _err := _template.writeValue(
						_output,
						(term.ValueType),
						"internal_definitions.template:326:11")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// internal_definitions.template:327:12
				if idx != len(clause.Bindings)-1 {
					// internal_definitions.template:327:49
					{
						_n, _err := _output.Write([]byte(`,`))
						_numWritten += int64(_n)
						if _err != nil {
							return _numWritten, _err
						}
					}
				}
			}
			// internal_definitions.template:328:17
			{
				_n, _err := _output.Write([]byte(`)`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
	}
	// internal_definitions.template:331:8
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
	// internal_definitions.template:337:14
	{
		_n, _err := _template.writeValue(
			_output,
			(Errorf),
			"internal_definitions.template:337:14")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// internal_definitions.template:337:21
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
