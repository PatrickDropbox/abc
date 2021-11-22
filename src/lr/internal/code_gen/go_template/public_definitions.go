// Auto-generated from source: public_definitions.template

package go_template

import (
	_fmt "fmt"
	_io "io"

	"fmt"

	lr "github.com/pattyshack/abc/src/lr/internal"
	"github.com/pattyshack/abc/src/lr/internal/parser"
)

type PublicDefinitions struct {
	LocationType      string
	SymbolIdType      string
	SymbolType        string
	GenericSymbolType string
	LexerType         string
	ReducerType       string
	ErrHandler        string
	DefaultErrHandler string
	StackType         string
	ExpectedTerminals string
	ParsePrefix       string
	InternalParse     string
	Sprintf           interface{}
	Errorf            interface{}
	Terminals         []*lr.Term
	NonTerminals      []*lr.Term
	Starts            []*lr.Term
	OrderedStates     []*lr.ItemSet
}

func (PublicDefinitions) Name() string { return "PublicDefinitions" }

func (template *PublicDefinitions) writeValue(
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

func (_template *PublicDefinitions) WriteTo(
	_output _io.Writer) (
	int64,
	error) {

	_numWritten := int64(0)

	LocationType := _template.LocationType
	SymbolIdType := _template.SymbolIdType
	SymbolType := _template.SymbolType
	GenericSymbolType := _template.GenericSymbolType
	LexerType := _template.LexerType
	ReducerType := _template.ReducerType
	ErrHandler := _template.ErrHandler
	DefaultErrHandler := _template.DefaultErrHandler
	StackType := _template.StackType
	ExpectedTerminals := _template.ExpectedTerminals
	ParsePrefix := _template.ParsePrefix
	InternalParse := _template.InternalParse
	Sprintf := _template.Sprintf
	Errorf := _template.Errorf
	Terminals := _template.Terminals
	NonTerminals := _template.NonTerminals
	Starts := _template.Starts
	OrderedStates := _template.OrderedStates

	// public_definitions.template:41:0
	{
		_n, _err := _output.Write([]byte(`type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:41:5
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"public_definitions.template:41:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:41:18
	{
		_n, _err := _output.Write([]byte(` int

const (`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:44:0
	nextId := 256
	// public_definitions.template:45:0
	for _, term := range Terminals {
		// public_definitions.template:46:4
		if term.SymbolId == parser.LRIdentifierToken {
			// public_definitions.template:46:53
			{
				_n, _err := _output.Write([]byte(`
    `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// public_definitions.template:47:4
			{
				_n, _err := _template.writeValue(
					_output,
					(term.CodeGenSymbolConst),
					"public_definitions.template:47:4")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// public_definitions.template:47:30
			{
				_n, _err := _output.Write([]byte(` = `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// public_definitions.template:47:33
			{
				_n, _err := _template.writeValue(
					_output,
					(SymbolIdType),
					"public_definitions.template:47:33")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// public_definitions.template:47:48
			{
				_n, _err := _output.Write([]byte(`(`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// public_definitions.template:47:49
			{
				_n, _err := _template.writeValue(
					_output,
					(nextId),
					"public_definitions.template:47:49")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// public_definitions.template:47:56
			{
				_n, _err := _output.Write([]byte(`)`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// public_definitions.template:48:8
			nextId += 1
		}
	}
	// public_definitions.template:50:8
	{
		_n, _err := _output.Write([]byte(`
)

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:53:5
	{
		_n, _err := _template.writeValue(
			_output,
			(LocationType),
			"public_definitions.template:53:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:53:18
	{
		_n, _err := _output.Write([]byte(` struct {
    FileName string
    Line int
    Column int
}

func (l `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:59:8
	{
		_n, _err := _template.writeValue(
			_output,
			(LocationType),
			"public_definitions.template:59:8")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:59:21
	{
		_n, _err := _output.Write([]byte(`) String() string {
    return `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:60:11
	{
		_n, _err := _template.writeValue(
			_output,
			(Sprintf),
			"public_definitions.template:60:11")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:60:19
	{
		_n, _err := _output.Write([]byte(`("%v:%v:%v", l.FileName, l.Line, l.Column)
}

func (l `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:63:8
	{
		_n, _err := _template.writeValue(
			_output,
			(LocationType),
			"public_definitions.template:63:8")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:63:21
	{
		_n, _err := _output.Write([]byte(`) ShortString() string {
    return `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:64:11
	{
		_n, _err := _template.writeValue(
			_output,
			(Sprintf),
			"public_definitions.template:64:11")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:64:19
	{
		_n, _err := _output.Write([]byte(`("%v:%v", l.Line, l.Column)
}

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:67:5
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"public_definitions.template:67:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:67:16
	{
		_n, _err := _output.Write([]byte(` interface {
    Id() `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:68:9
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"public_definitions.template:68:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:68:22
	{
		_n, _err := _output.Write([]byte(`
    Loc() `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:69:10
	{
		_n, _err := _template.writeValue(
			_output,
			(LocationType),
			"public_definitions.template:69:10")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:69:23
	{
		_n, _err := _output.Write([]byte(`
}

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:72:5
	{
		_n, _err := _template.writeValue(
			_output,
			(GenericSymbolType),
			"public_definitions.template:72:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:72:23
	{
		_n, _err := _output.Write([]byte(` struct {
    `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:73:4
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"public_definitions.template:73:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:73:17
	{
		_n, _err := _output.Write([]byte(`
    `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:74:4
	{
		_n, _err := _template.writeValue(
			_output,
			(LocationType),
			"public_definitions.template:74:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:74:17
	{
		_n, _err := _output.Write([]byte(`
}

func (t *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:77:9
	{
		_n, _err := _template.writeValue(
			_output,
			(GenericSymbolType),
			"public_definitions.template:77:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:77:27
	{
		_n, _err := _output.Write([]byte(`) Id() `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:77:34
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"public_definitions.template:77:34")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:77:47
	{
		_n, _err := _output.Write([]byte(` { return t.`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:77:59
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"public_definitions.template:77:59")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:77:72
	{
		_n, _err := _output.Write([]byte(` }

func (t *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:79:9
	{
		_n, _err := _template.writeValue(
			_output,
			(GenericSymbolType),
			"public_definitions.template:79:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:79:27
	{
		_n, _err := _output.Write([]byte(`) Loc() `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:79:35
	{
		_n, _err := _template.writeValue(
			_output,
			(LocationType),
			"public_definitions.template:79:35")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:79:48
	{
		_n, _err := _output.Write([]byte(` { return t.`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:79:60
	{
		_n, _err := _template.writeValue(
			_output,
			(LocationType),
			"public_definitions.template:79:60")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:79:73
	{
		_n, _err := _output.Write([]byte(` }

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:81:5
	{
		_n, _err := _template.writeValue(
			_output,
			(LexerType),
			"public_definitions.template:81:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:81:15
	{
		_n, _err := _output.Write([]byte(` interface {
    // Note: Return io.EOF to indicate end of stream
    // Token with unspecified value type should return *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:83:56
	{
		_n, _err := _template.writeValue(
			_output,
			(GenericSymbolType),
			"public_definitions.template:83:56")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:83:74
	{
		_n, _err := _output.Write([]byte(`
    Next() (`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:84:12
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"public_definitions.template:84:12")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:84:23
	{
		_n, _err := _output.Write([]byte(`, error)

    CurrentLocation() `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:86:22
	{
		_n, _err := _template.writeValue(
			_output,
			(LocationType),
			"public_definitions.template:86:22")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:86:35
	{
		_n, _err := _output.Write([]byte(`
}


type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:90:5
	{
		_n, _err := _template.writeValue(
			_output,
			(ReducerType),
			"public_definitions.template:90:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:90:17
	{
		_n, _err := _output.Write([]byte(` interface {`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:91:0
	for ruleIdx, rule := range NonTerminals {
		// public_definitions.template:92:4
		if ruleIdx > 0 {
			// public_definitions.template:92:23
			{
				_n, _err := _output.Write([]byte(`
`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
		// public_definitions.template:96:4
		for clauseIdx, clause := range rule.Clauses {
			// public_definitions.template:97:8
			if clauseIdx > 0 {
				// public_definitions.template:97:29
				{
					_n, _err := _output.Write([]byte(`
`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
			}
			// public_definitions.template:101:8
			if clause.Label == "" {
				// public_definitions.template:101:34
				{
					_n, _err := _output.Write([]byte(`
    // `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// public_definitions.template:102:7
				{
					_n, _err := _template.writeValue(
						_output,
						(clause.LRLocation.ShortString()),
						"public_definitions.template:102:7")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// public_definitions.template:102:41
				{
					_n, _err := _output.Write([]byte(`: `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// public_definitions.template:102:43
				{
					_n, _err := _template.writeValue(
						_output,
						(rule.Name),
						"public_definitions.template:102:43")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// public_definitions.template:102:55
				{
					_n, _err := _output.Write([]byte(` -> ...`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
			} else {
				// public_definitions.template:103:17
				{
					_n, _err := _output.Write([]byte(`
    // `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// public_definitions.template:104:7
				{
					_n, _err := _template.writeValue(
						_output,
						(clause.LRLocation.ShortString()),
						"public_definitions.template:104:7")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// public_definitions.template:104:41
				{
					_n, _err := _output.Write([]byte(`: `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// public_definitions.template:104:43
				{
					_n, _err := _template.writeValue(
						_output,
						(rule.Name),
						"public_definitions.template:104:43")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// public_definitions.template:104:55
				{
					_n, _err := _output.Write([]byte(` -> `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// public_definitions.template:104:59
				{
					_n, _err := _template.writeValue(
						_output,
						(clause.Label),
						"public_definitions.template:104:59")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// public_definitions.template:104:74
				{
					_n, _err := _output.Write([]byte(`: ...`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
			}
			// public_definitions.template:107:8
			paramNameCount := map[string]int{}
			// public_definitions.template:107:49
			{
				_n, _err := _output.Write([]byte(`
    `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// public_definitions.template:108:4
			{
				_n, _err := _template.writeValue(
					_output,
					(clause.CodeGenReducerName),
					"public_definitions.template:108:4")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// public_definitions.template:108:32
			{
				_n, _err := _output.Write([]byte(`(`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// public_definitions.template:109:8
			for termIdx, term := range clause.Bindings {
				// public_definitions.template:111:12

				paramName := ""
				if term.SymbolId == parser.LRCharacterToken {
					paramName = "char"
				} else {
					// hack: append "_" to the end of the name ensures the
					// name is never a go keyword
					paramName = snakeToCamel(term.Name) + "_"
				}

				paramNameCount[paramName] += 1
				cnt := paramNameCount[paramName]
				if cnt > 1 {
					paramName = fmt.Sprintf("%s%d", paramName, cnt)
				}

				suffix := ""
				if termIdx != len(clause.Bindings) {
					suffix = ", "
				}

				// public_definitions.template:133:15
				{
					_n, _err := _output.Write([]byte(`        `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// public_definitions.template:134:8
				{
					_n, _err := _template.writeValue(
						_output,
						(paramName),
						"public_definitions.template:134:8")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// public_definitions.template:134:18
				{
					_n, _err := _output.Write([]byte(` `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// public_definitions.template:134:19
				{
					_n, _err := _template.writeValue(
						_output,
						(term.CodeGenType),
						"public_definitions.template:134:19")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// public_definitions.template:134:38
				{
					_n, _err := _template.writeValue(
						_output,
						(suffix),
						"public_definitions.template:134:38")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
			}
			// public_definitions.template:135:17
			{
				_n, _err := _output.Write([]byte(`) (`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// public_definitions.template:136:3
			{
				_n, _err := _template.writeValue(
					_output,
					(rule.CodeGenType),
					"public_definitions.template:136:3")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// public_definitions.template:136:22
			{
				_n, _err := _output.Write([]byte(`, error)`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
	}
	// public_definitions.template:138:8
	{
		_n, _err := _output.Write([]byte(`
}

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:141:5
	{
		_n, _err := _template.writeValue(
			_output,
			(ErrHandler),
			"public_definitions.template:141:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:141:16
	{
		_n, _err := _output.Write([]byte(` interface {
    Error(nextToken `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:142:20
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"public_definitions.template:142:20")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:142:31
	{
		_n, _err := _output.Write([]byte(`, parseStack `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:142:44
	{
		_n, _err := _template.writeValue(
			_output,
			(StackType),
			"public_definitions.template:142:44")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:142:54
	{
		_n, _err := _output.Write([]byte(`) error
}

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:145:5
	{
		_n, _err := _template.writeValue(
			_output,
			(DefaultErrHandler),
			"public_definitions.template:145:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:145:23
	{
		_n, _err := _output.Write([]byte(` struct {}

func (`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:147:6
	{
		_n, _err := _template.writeValue(
			_output,
			(DefaultErrHandler),
			"public_definitions.template:147:6")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:147:24
	{
		_n, _err := _output.Write([]byte(`) Error(nextToken `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:147:42
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"public_definitions.template:147:42")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:147:53
	{
		_n, _err := _output.Write([]byte(`, stack `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:147:61
	{
		_n, _err := _template.writeValue(
			_output,
			(StackType),
			"public_definitions.template:147:61")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:147:71
	{
		_n, _err := _output.Write([]byte(`) error {
    return `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:148:11
	{
		_n, _err := _template.writeValue(
			_output,
			(Errorf),
			"public_definitions.template:148:11")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:148:18
	{
		_n, _err := _output.Write([]byte(`(
        "Syntax error: unexpected symbol %v. Expecting %v (%v)",
        nextToken.Id(),
        `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:151:8
	{
		_n, _err := _template.writeValue(
			_output,
			(ExpectedTerminals),
			"public_definitions.template:151:8")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:151:26
	{
		_n, _err := _output.Write([]byte(`[stack[len(stack)-1].StateId],
        nextToken.Loc())
}
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:155:0
	for idx, start := range Starts {
		// public_definitions.template:156:4

		parseSuffix := ""
		if len(Starts) > 1 {
			parseSuffix = snakeToCamel(start.Name)
		}

		// public_definitions.template:163:6
		{
			_n, _err := _output.Write([]byte(`
func `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// public_definitions.template:164:5
		{
			_n, _err := _template.writeValue(
				_output,
				(ParsePrefix),
				"public_definitions.template:164:5")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// public_definitions.template:164:19
		{
			_n, _err := _template.writeValue(
				_output,
				(parseSuffix),
				"public_definitions.template:164:19")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// public_definitions.template:164:33
		{
			_n, _err := _output.Write([]byte(`(lexer `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// public_definitions.template:164:40
		{
			_n, _err := _template.writeValue(
				_output,
				(LexerType),
				"public_definitions.template:164:40")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// public_definitions.template:164:50
		{
			_n, _err := _output.Write([]byte(`, reducer `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// public_definitions.template:164:60
		{
			_n, _err := _template.writeValue(
				_output,
				(ReducerType),
				"public_definitions.template:164:60")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// public_definitions.template:164:72
		{
			_n, _err := _output.Write([]byte(`) (`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// public_definitions.template:166:0
		{
			_n, _err := _template.writeValue(
				_output,
				(start.CodeGenType),
				"public_definitions.template:166:0")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// public_definitions.template:166:20
		{
			_n, _err := _output.Write([]byte(`, error) {

    return `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// public_definitions.template:168:11
		{
			_n, _err := _template.writeValue(
				_output,
				(ParsePrefix),
				"public_definitions.template:168:11")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// public_definitions.template:168:25
		{
			_n, _err := _template.writeValue(
				_output,
				(parseSuffix),
				"public_definitions.template:168:25")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// public_definitions.template:168:39
		{
			_n, _err := _output.Write([]byte(`WithCustomErrorHandler(
        lexer,
        reducer,
        `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// public_definitions.template:171:8
		{
			_n, _err := _template.writeValue(
				_output,
				(DefaultErrHandler),
				"public_definitions.template:171:8")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// public_definitions.template:171:26
		{
			_n, _err := _output.Write([]byte(`{})
}

func `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// public_definitions.template:174:5
		{
			_n, _err := _template.writeValue(
				_output,
				(ParsePrefix),
				"public_definitions.template:174:5")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// public_definitions.template:174:19
		{
			_n, _err := _template.writeValue(
				_output,
				(parseSuffix),
				"public_definitions.template:174:19")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// public_definitions.template:174:33
		{
			_n, _err := _output.Write([]byte(`WithCustomErrorHandler(
    lexer `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// public_definitions.template:175:10
		{
			_n, _err := _template.writeValue(
				_output,
				(LexerType),
				"public_definitions.template:175:10")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// public_definitions.template:175:20
		{
			_n, _err := _output.Write([]byte(`,
    reducer `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// public_definitions.template:176:12
		{
			_n, _err := _template.writeValue(
				_output,
				(ReducerType),
				"public_definitions.template:176:12")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// public_definitions.template:176:24
		{
			_n, _err := _output.Write([]byte(`,
    errHandler `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// public_definitions.template:177:15
		{
			_n, _err := _template.writeValue(
				_output,
				(ErrHandler),
				"public_definitions.template:177:15")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// public_definitions.template:177:26
		{
			_n, _err := _output.Write([]byte(`) (
    `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// public_definitions.template:178:4
		{
			_n, _err := _template.writeValue(
				_output,
				(start.CodeGenType),
				"public_definitions.template:178:4")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// public_definitions.template:178:24
		{
			_n, _err := _output.Write([]byte(`,
    error) {

    item, err := `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// public_definitions.template:181:17
		{
			_n, _err := _template.writeValue(
				_output,
				(InternalParse),
				"public_definitions.template:181:17")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// public_definitions.template:181:31
		{
			_n, _err := _output.Write([]byte(`(lexer, reducer, errHandler, `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// public_definitions.template:181:60
		{
			_n, _err := _template.writeValue(
				_output,
				(OrderedStates[idx].CodeGenConst),
				"public_definitions.template:181:60")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// public_definitions.template:182:40
		{
			_n, _err := _output.Write([]byte(`)
    if err != nil {
        var errRetVal `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// public_definitions.template:184:22
		{
			_n, _err := _template.writeValue(
				_output,
				(start.CodeGenType),
				"public_definitions.template:184:22")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// public_definitions.template:184:42
		{
			_n, _err := _output.Write([]byte(`
        return errRetVal, err
    }
    return item.`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// public_definitions.template:187:16
		{
			_n, _err := _template.writeValue(
				_output,
				(start.ValueType),
				"public_definitions.template:187:16")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// public_definitions.template:187:34
		{
			_n, _err := _output.Write([]byte(`, nil
}
`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// public_definitions.template:189:7
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
