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
	ParsePrefix       string
	InternalParse     string
	Sprintf           interface{}
	Errorf            interface{}
	Terminals         []*lr.Term
	NonTerminals      []*lr.Term
	Starts            []*lr.Term
	OrderedStates     []*lr.ItemSet
	ExpectedTerminals string
	StateIdType       string
	ActionTable       string
	SortSlice         interface{}
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
	ParsePrefix := _template.ParsePrefix
	InternalParse := _template.InternalParse
	Sprintf := _template.Sprintf
	Errorf := _template.Errorf
	Terminals := _template.Terminals
	NonTerminals := _template.NonTerminals
	Starts := _template.Starts
	OrderedStates := _template.OrderedStates
	ExpectedTerminals := _template.ExpectedTerminals
	StateIdType := _template.StateIdType
	ActionTable := _template.ActionTable
	SortSlice := _template.SortSlice

	// public_definitions.template:45:0
	{
		_n, _err := _output.Write([]byte(`type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:45:5
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"public_definitions.template:45:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:45:18
	{
		_n, _err := _output.Write([]byte(` int

const (`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:48:0
	nextId := 256
	// public_definitions.template:49:0
	for _, term := range Terminals {
		// public_definitions.template:50:4
		if term.SymbolId == parser.LRIdentifierToken {
			// public_definitions.template:50:53
			{
				_n, _err := _output.Write([]byte(`
    `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// public_definitions.template:51:4
			{
				_n, _err := _template.writeValue(
					_output,
					(term.CodeGenSymbolConst),
					"public_definitions.template:51:4")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// public_definitions.template:51:30
			{
				_n, _err := _output.Write([]byte(` = `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// public_definitions.template:51:33
			{
				_n, _err := _template.writeValue(
					_output,
					(SymbolIdType),
					"public_definitions.template:51:33")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// public_definitions.template:51:48
			{
				_n, _err := _output.Write([]byte(`(`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// public_definitions.template:51:49
			{
				_n, _err := _template.writeValue(
					_output,
					(nextId),
					"public_definitions.template:51:49")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// public_definitions.template:51:56
			{
				_n, _err := _output.Write([]byte(`)`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// public_definitions.template:52:8
			nextId += 1
		}
	}
	// public_definitions.template:54:8
	{
		_n, _err := _output.Write([]byte(`
)

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:57:5
	{
		_n, _err := _template.writeValue(
			_output,
			(LocationType),
			"public_definitions.template:57:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:57:18
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
		_n, _err := _output.Write([]byte(`) String() string {
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
		_n, _err := _output.Write([]byte(`("%v:%v:%v", l.FileName, l.Line, l.Column)
}

func (l `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:67:8
	{
		_n, _err := _template.writeValue(
			_output,
			(LocationType),
			"public_definitions.template:67:8")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:67:21
	{
		_n, _err := _output.Write([]byte(`) ShortString() string {
    return `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:68:11
	{
		_n, _err := _template.writeValue(
			_output,
			(Sprintf),
			"public_definitions.template:68:11")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:68:19
	{
		_n, _err := _output.Write([]byte(`("%v:%v", l.Line, l.Column)
}

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:71:5
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"public_definitions.template:71:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:71:16
	{
		_n, _err := _output.Write([]byte(` interface {
    Id() `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:72:9
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"public_definitions.template:72:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:72:22
	{
		_n, _err := _output.Write([]byte(`
    Loc() `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:73:10
	{
		_n, _err := _template.writeValue(
			_output,
			(LocationType),
			"public_definitions.template:73:10")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:73:23
	{
		_n, _err := _output.Write([]byte(`
}

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:76:5
	{
		_n, _err := _template.writeValue(
			_output,
			(GenericSymbolType),
			"public_definitions.template:76:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:76:23
	{
		_n, _err := _output.Write([]byte(` struct {
    `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:77:4
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"public_definitions.template:77:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:77:17
	{
		_n, _err := _output.Write([]byte(`
    `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:78:4
	{
		_n, _err := _template.writeValue(
			_output,
			(LocationType),
			"public_definitions.template:78:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:78:17
	{
		_n, _err := _output.Write([]byte(`
}

func (t *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:81:9
	{
		_n, _err := _template.writeValue(
			_output,
			(GenericSymbolType),
			"public_definitions.template:81:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:81:27
	{
		_n, _err := _output.Write([]byte(`) Id() `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:81:34
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"public_definitions.template:81:34")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:81:47
	{
		_n, _err := _output.Write([]byte(` { return t.`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:81:59
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"public_definitions.template:81:59")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:81:72
	{
		_n, _err := _output.Write([]byte(` }

func (t *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:83:9
	{
		_n, _err := _template.writeValue(
			_output,
			(GenericSymbolType),
			"public_definitions.template:83:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:83:27
	{
		_n, _err := _output.Write([]byte(`) Loc() `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:83:35
	{
		_n, _err := _template.writeValue(
			_output,
			(LocationType),
			"public_definitions.template:83:35")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:83:48
	{
		_n, _err := _output.Write([]byte(` { return t.`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:83:60
	{
		_n, _err := _template.writeValue(
			_output,
			(LocationType),
			"public_definitions.template:83:60")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:83:73
	{
		_n, _err := _output.Write([]byte(` }

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:85:5
	{
		_n, _err := _template.writeValue(
			_output,
			(LexerType),
			"public_definitions.template:85:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:85:15
	{
		_n, _err := _output.Write([]byte(` interface {
    // Note: Return io.EOF to indicate end of stream
    // Token with unspecified value type should return *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:87:56
	{
		_n, _err := _template.writeValue(
			_output,
			(GenericSymbolType),
			"public_definitions.template:87:56")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:87:74
	{
		_n, _err := _output.Write([]byte(`
    Next() (`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:88:12
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"public_definitions.template:88:12")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:88:23
	{
		_n, _err := _output.Write([]byte(`, error)

    CurrentLocation() `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:90:22
	{
		_n, _err := _template.writeValue(
			_output,
			(LocationType),
			"public_definitions.template:90:22")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:90:35
	{
		_n, _err := _output.Write([]byte(`
}


type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:94:5
	{
		_n, _err := _template.writeValue(
			_output,
			(ReducerType),
			"public_definitions.template:94:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:94:17
	{
		_n, _err := _output.Write([]byte(` interface {`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:95:0
	for ruleIdx, rule := range NonTerminals {
		// public_definitions.template:96:4
		if ruleIdx > 0 {
			// public_definitions.template:96:23
			{
				_n, _err := _output.Write([]byte(`
`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
		// public_definitions.template:100:4
		for clauseIdx, clause := range rule.Clauses {
			// public_definitions.template:101:8
			if clauseIdx > 0 {
				// public_definitions.template:101:29
				{
					_n, _err := _output.Write([]byte(`
`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
			}
			// public_definitions.template:105:8
			if clause.Label == "" {
				// public_definitions.template:105:34
				{
					_n, _err := _output.Write([]byte(`
    // `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// public_definitions.template:106:7
				{
					_n, _err := _template.writeValue(
						_output,
						(clause.LRLocation.ShortString()),
						"public_definitions.template:106:7")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// public_definitions.template:106:41
				{
					_n, _err := _output.Write([]byte(`: `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// public_definitions.template:106:43
				{
					_n, _err := _template.writeValue(
						_output,
						(rule.Name),
						"public_definitions.template:106:43")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// public_definitions.template:106:55
				{
					_n, _err := _output.Write([]byte(` -> ...`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
			} else {
				// public_definitions.template:107:17
				{
					_n, _err := _output.Write([]byte(`
    // `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// public_definitions.template:108:7
				{
					_n, _err := _template.writeValue(
						_output,
						(clause.LRLocation.ShortString()),
						"public_definitions.template:108:7")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// public_definitions.template:108:41
				{
					_n, _err := _output.Write([]byte(`: `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// public_definitions.template:108:43
				{
					_n, _err := _template.writeValue(
						_output,
						(rule.Name),
						"public_definitions.template:108:43")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// public_definitions.template:108:55
				{
					_n, _err := _output.Write([]byte(` -> `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// public_definitions.template:108:59
				{
					_n, _err := _template.writeValue(
						_output,
						(clause.Label),
						"public_definitions.template:108:59")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// public_definitions.template:108:74
				{
					_n, _err := _output.Write([]byte(`: ...`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
			}
			// public_definitions.template:111:8
			paramNameCount := map[string]int{}
			// public_definitions.template:111:49
			{
				_n, _err := _output.Write([]byte(`
    `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// public_definitions.template:112:4
			{
				_n, _err := _template.writeValue(
					_output,
					(clause.CodeGenReducerName),
					"public_definitions.template:112:4")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// public_definitions.template:112:32
			{
				_n, _err := _output.Write([]byte(`(`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// public_definitions.template:113:8
			for termIdx, term := range clause.Bindings {
				// public_definitions.template:115:12

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

				// public_definitions.template:137:15
				{
					_n, _err := _output.Write([]byte(`        `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// public_definitions.template:138:8
				{
					_n, _err := _template.writeValue(
						_output,
						(paramName),
						"public_definitions.template:138:8")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// public_definitions.template:138:18
				{
					_n, _err := _output.Write([]byte(` `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// public_definitions.template:138:19
				{
					_n, _err := _template.writeValue(
						_output,
						(term.CodeGenType),
						"public_definitions.template:138:19")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// public_definitions.template:138:38
				{
					_n, _err := _template.writeValue(
						_output,
						(suffix),
						"public_definitions.template:138:38")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
			}
			// public_definitions.template:139:17
			{
				_n, _err := _output.Write([]byte(`) (`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// public_definitions.template:140:3
			{
				_n, _err := _template.writeValue(
					_output,
					(rule.CodeGenType),
					"public_definitions.template:140:3")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// public_definitions.template:140:22
			{
				_n, _err := _output.Write([]byte(`, error)`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
	}
	// public_definitions.template:142:8
	{
		_n, _err := _output.Write([]byte(`
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
			(ErrHandler),
			"public_definitions.template:145:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:145:16
	{
		_n, _err := _output.Write([]byte(` interface {
    Error(nextToken `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:146:20
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"public_definitions.template:146:20")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:146:31
	{
		_n, _err := _output.Write([]byte(`, parseStack `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:146:44
	{
		_n, _err := _template.writeValue(
			_output,
			(StackType),
			"public_definitions.template:146:44")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:146:54
	{
		_n, _err := _output.Write([]byte(`) error
}

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:149:5
	{
		_n, _err := _template.writeValue(
			_output,
			(DefaultErrHandler),
			"public_definitions.template:149:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:149:23
	{
		_n, _err := _output.Write([]byte(` struct {}

func (`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:151:6
	{
		_n, _err := _template.writeValue(
			_output,
			(DefaultErrHandler),
			"public_definitions.template:151:6")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:151:24
	{
		_n, _err := _output.Write([]byte(`) Error(nextToken `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:151:42
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"public_definitions.template:151:42")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:151:53
	{
		_n, _err := _output.Write([]byte(`, stack `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:151:61
	{
		_n, _err := _template.writeValue(
			_output,
			(StackType),
			"public_definitions.template:151:61")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:151:71
	{
		_n, _err := _output.Write([]byte(`) error {
    return `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:152:11
	{
		_n, _err := _template.writeValue(
			_output,
			(Errorf),
			"public_definitions.template:152:11")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:152:18
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
	// public_definitions.template:155:8
	{
		_n, _err := _template.writeValue(
			_output,
			(ExpectedTerminals),
			"public_definitions.template:155:8")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:155:26
	{
		_n, _err := _output.Write([]byte(`(stack[len(stack)-1].StateId),
        nextToken.Loc())
}

func `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:159:5
	{
		_n, _err := _template.writeValue(
			_output,
			(ExpectedTerminals),
			"public_definitions.template:159:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:159:23
	{
		_n, _err := _output.Write([]byte(`(id `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:159:27
	{
		_n, _err := _template.writeValue(
			_output,
			(StateIdType),
			"public_definitions.template:159:27")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:159:39
	{
		_n, _err := _output.Write([]byte(`) []`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:159:43
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"public_definitions.template:159:43")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:159:56
	{
		_n, _err := _output.Write([]byte(` {
    result := []`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:160:16
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"public_definitions.template:160:16")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:160:29
	{
		_n, _err := _output.Write([]byte(`{}
    for key, _ := range `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:161:24
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionTable),
			"public_definitions.template:161:24")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:161:36
	{
		_n, _err := _output.Write([]byte(` {
        if key.`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:162:15
	{
		_n, _err := _template.writeValue(
			_output,
			(StateIdType),
			"public_definitions.template:162:15")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:162:27
	{
		_n, _err := _output.Write([]byte(` != id {
            continue
        }
        result = append(result, key.`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:165:36
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"public_definitions.template:165:36")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:165:49
	{
		_n, _err := _output.Write([]byte(`)
    }

    `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:168:4
	{
		_n, _err := _template.writeValue(
			_output,
			(SortSlice),
			"public_definitions.template:168:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:168:14
	{
		_n, _err := _output.Write([]byte(`(result, func(i int, j int) bool {return result[i] < result[j]})
    return result
}
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// public_definitions.template:172:0
	for idx, start := range Starts {
		// public_definitions.template:173:4

		parseSuffix := ""
		if len(Starts) > 1 {
			parseSuffix = snakeToCamel(start.Name)
		}

		// public_definitions.template:180:6
		{
			_n, _err := _output.Write([]byte(`
func `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// public_definitions.template:181:5
		{
			_n, _err := _template.writeValue(
				_output,
				(ParsePrefix),
				"public_definitions.template:181:5")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// public_definitions.template:181:19
		{
			_n, _err := _template.writeValue(
				_output,
				(parseSuffix),
				"public_definitions.template:181:19")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// public_definitions.template:181:33
		{
			_n, _err := _output.Write([]byte(`(lexer `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// public_definitions.template:181:40
		{
			_n, _err := _template.writeValue(
				_output,
				(LexerType),
				"public_definitions.template:181:40")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// public_definitions.template:181:50
		{
			_n, _err := _output.Write([]byte(`, reducer `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// public_definitions.template:181:60
		{
			_n, _err := _template.writeValue(
				_output,
				(ReducerType),
				"public_definitions.template:181:60")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// public_definitions.template:181:72
		{
			_n, _err := _output.Write([]byte(`) (`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// public_definitions.template:183:0
		{
			_n, _err := _template.writeValue(
				_output,
				(start.CodeGenType),
				"public_definitions.template:183:0")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// public_definitions.template:183:20
		{
			_n, _err := _output.Write([]byte(`, error) {

    return `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// public_definitions.template:185:11
		{
			_n, _err := _template.writeValue(
				_output,
				(ParsePrefix),
				"public_definitions.template:185:11")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// public_definitions.template:185:25
		{
			_n, _err := _template.writeValue(
				_output,
				(parseSuffix),
				"public_definitions.template:185:25")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// public_definitions.template:185:39
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
		// public_definitions.template:188:8
		{
			_n, _err := _template.writeValue(
				_output,
				(DefaultErrHandler),
				"public_definitions.template:188:8")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// public_definitions.template:188:26
		{
			_n, _err := _output.Write([]byte(`{})
}

func `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// public_definitions.template:191:5
		{
			_n, _err := _template.writeValue(
				_output,
				(ParsePrefix),
				"public_definitions.template:191:5")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// public_definitions.template:191:19
		{
			_n, _err := _template.writeValue(
				_output,
				(parseSuffix),
				"public_definitions.template:191:19")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// public_definitions.template:191:33
		{
			_n, _err := _output.Write([]byte(`WithCustomErrorHandler(
    lexer `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// public_definitions.template:192:10
		{
			_n, _err := _template.writeValue(
				_output,
				(LexerType),
				"public_definitions.template:192:10")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// public_definitions.template:192:20
		{
			_n, _err := _output.Write([]byte(`,
    reducer `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// public_definitions.template:193:12
		{
			_n, _err := _template.writeValue(
				_output,
				(ReducerType),
				"public_definitions.template:193:12")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// public_definitions.template:193:24
		{
			_n, _err := _output.Write([]byte(`,
    errHandler `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// public_definitions.template:194:15
		{
			_n, _err := _template.writeValue(
				_output,
				(ErrHandler),
				"public_definitions.template:194:15")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// public_definitions.template:194:26
		{
			_n, _err := _output.Write([]byte(`) (
    `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// public_definitions.template:195:4
		{
			_n, _err := _template.writeValue(
				_output,
				(start.CodeGenType),
				"public_definitions.template:195:4")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// public_definitions.template:195:24
		{
			_n, _err := _output.Write([]byte(`,
    error) {

    item, err := `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// public_definitions.template:198:17
		{
			_n, _err := _template.writeValue(
				_output,
				(InternalParse),
				"public_definitions.template:198:17")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// public_definitions.template:198:31
		{
			_n, _err := _output.Write([]byte(`(lexer, reducer, errHandler, `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// public_definitions.template:198:60
		{
			_n, _err := _template.writeValue(
				_output,
				(OrderedStates[idx].CodeGenConst),
				"public_definitions.template:198:60")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// public_definitions.template:199:40
		{
			_n, _err := _output.Write([]byte(`)
    if err != nil {
        var errRetVal `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// public_definitions.template:201:22
		{
			_n, _err := _template.writeValue(
				_output,
				(start.CodeGenType),
				"public_definitions.template:201:22")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// public_definitions.template:201:42
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
		// public_definitions.template:204:16
		{
			_n, _err := _template.writeValue(
				_output,
				(start.ValueType),
				"public_definitions.template:204:16")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// public_definitions.template:204:34
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
	// public_definitions.template:206:7
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
