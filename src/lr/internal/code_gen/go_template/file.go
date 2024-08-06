// Auto-generated from source: file.template

package go_template

import (
	_fmt "fmt"
	_io "io"

	"fmt"
	"io"

	"github.com/pattyshack/abc/src/lr/codegenutil"
	lr "github.com/pattyshack/abc/src/lr/internal"
	"github.com/pattyshack/abc/src/lr/internal/code_gen/debug_template"
	parser "github.com/pattyshack/abc/src/lr/internal/parser"
)

type File struct {
	Package                   string
	Imports                   io.WriterTo
	ActionType                string
	ActionIdType              string
	ShiftAction               string
	ReduceAction              string
	AcceptAction              string
	StateIdType               string
	ReduceType                string
	SymbolType                string
	GenericSymbolType         string
	StackItemType             string
	StackType                 string
	SymbolStackType           string
	SymbolIdType              string
	EndSymbolId               string
	WildcardSymbolId          string
	LocationType              string
	TokenType                 string
	LexerType                 string
	ReducerType               string
	ErrHandlerType            string
	DefaultErrHandlerType     string
	ExpectedTerminalsFunc     string
	ParseFuncPrefix           string
	InternalParseFunc         string
	TableKeyType              string
	ActionTableType           string
	ActionTable               string
	SortSlice                 interface{}
	Sprintf                   interface{}
	Errorf                    interface{}
	EOF                       interface{}
	OrderedSymbolNames        []string
	Grammar                   *lr.Grammar
	States                    *lr.LRStates
	OrderedValueTypes         lr.ParamList
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

	Package := _template.Package
	Imports := _template.Imports
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
	SymbolStackType := _template.SymbolStackType
	SymbolIdType := _template.SymbolIdType
	EndSymbolId := _template.EndSymbolId
	WildcardSymbolId := _template.WildcardSymbolId
	LocationType := _template.LocationType
	TokenType := _template.TokenType
	LexerType := _template.LexerType
	ReducerType := _template.ReducerType
	ErrHandlerType := _template.ErrHandlerType
	DefaultErrHandlerType := _template.DefaultErrHandlerType
	ExpectedTerminalsFunc := _template.ExpectedTerminalsFunc
	ParseFuncPrefix := _template.ParseFuncPrefix
	InternalParseFunc := _template.InternalParseFunc
	TableKeyType := _template.TableKeyType
	ActionTableType := _template.ActionTableType
	ActionTable := _template.ActionTable
	SortSlice := _template.SortSlice
	Sprintf := _template.Sprintf
	Errorf := _template.Errorf
	EOF := _template.EOF
	OrderedSymbolNames := _template.OrderedSymbolNames
	Grammar := _template.Grammar
	States := _template.States
	OrderedValueTypes := _template.OrderedValueTypes
	OutputDebugNonKernelItems := _template.OutputDebugNonKernelItems

	// file.template:74:0
	{
		_n, _err := _output.Write([]byte(`// Auto-generated from source: `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:74:31
	{
		_n, _err := _template.writeValue(
			_output,
			(Grammar.Source),
			"file.template:74:31")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:74:48
	{
		_n, _err := _output.Write([]byte(`

package `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:76:8
	{
		_n, _err := _template.writeValue(
			_output,
			(Package),
			"file.template:76:8")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:76:16
	{
		_n, _err := _output.Write([]byte(`
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:78:0
	{
		_n, _err := (Imports).WriteTo(_output)
		_numWritten += _n
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:78:20
	{
		_n, _err := _output.Write([]byte(`
type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:80:5
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"file.template:80:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:80:18
	{
		_n, _err := _output.Write([]byte(` int

const (`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:83:0
	nextId := 256
	// file.template:84:0
	for _, term := range Grammar.Terminals {
		// file.template:85:4
		if term.SymbolId == parser.LRIdentifierToken {
			// file.template:85:53
			{
				_n, _err := _output.Write([]byte(`
    `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:86:4
			{
				_n, _err := _template.writeValue(
					_output,
					(term.CodeGenSymbolConst),
					"file.template:86:4")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:86:30
			{
				_n, _err := _output.Write([]byte(` = `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:86:33
			{
				_n, _err := _template.writeValue(
					_output,
					(SymbolIdType),
					"file.template:86:33")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:86:48
			{
				_n, _err := _output.Write([]byte(`(`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:86:49
			{
				_n, _err := _template.writeValue(
					_output,
					(nextId),
					"file.template:86:49")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:86:56
			{
				_n, _err := _output.Write([]byte(`)`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:87:8
			nextId += 1
		}
	}
	// file.template:89:8
	{
		_n, _err := _output.Write([]byte(`
)

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:92:5
	{
		_n, _err := _template.writeValue(
			_output,
			(LocationType),
			"file.template:92:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:92:18
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
	// file.template:98:8
	{
		_n, _err := _template.writeValue(
			_output,
			(LocationType),
			"file.template:98:8")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:98:21
	{
		_n, _err := _output.Write([]byte(`) String() string {
    return `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:99:11
	{
		_n, _err := _template.writeValue(
			_output,
			(Sprintf),
			"file.template:99:11")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:99:19
	{
		_n, _err := _output.Write([]byte(`("%v:%v:%v", l.FileName, l.Line, l.Column)
}

func (l `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:102:8
	{
		_n, _err := _template.writeValue(
			_output,
			(LocationType),
			"file.template:102:8")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:102:21
	{
		_n, _err := _output.Write([]byte(`) ShortString() string {
    return `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:103:11
	{
		_n, _err := _template.writeValue(
			_output,
			(Sprintf),
			"file.template:103:11")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:103:19
	{
		_n, _err := _output.Write([]byte(`("%v:%v", l.Line, l.Column)
}

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:106:5
	{
		_n, _err := _template.writeValue(
			_output,
			(TokenType),
			"file.template:106:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:106:15
	{
		_n, _err := _output.Write([]byte(` interface {
    Id() `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:107:9
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"file.template:107:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:107:22
	{
		_n, _err := _output.Write([]byte(`
    Loc() `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:108:10
	{
		_n, _err := _template.writeValue(
			_output,
			(LocationType),
			"file.template:108:10")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:108:23
	{
		_n, _err := _output.Write([]byte(`
}

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:111:5
	{
		_n, _err := _template.writeValue(
			_output,
			(GenericSymbolType),
			"file.template:111:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:111:23
	{
		_n, _err := _output.Write([]byte(` struct {
    `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:112:4
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"file.template:112:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:112:17
	{
		_n, _err := _output.Write([]byte(`
    `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:113:4
	{
		_n, _err := _template.writeValue(
			_output,
			(LocationType),
			"file.template:113:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:113:17
	{
		_n, _err := _output.Write([]byte(`
}

func (t *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:116:9
	{
		_n, _err := _template.writeValue(
			_output,
			(GenericSymbolType),
			"file.template:116:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:116:27
	{
		_n, _err := _output.Write([]byte(`) Id() `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:116:34
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"file.template:116:34")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:116:47
	{
		_n, _err := _output.Write([]byte(` { return t.`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:116:59
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"file.template:116:59")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:116:72
	{
		_n, _err := _output.Write([]byte(` }

func (t *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:118:9
	{
		_n, _err := _template.writeValue(
			_output,
			(GenericSymbolType),
			"file.template:118:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:118:27
	{
		_n, _err := _output.Write([]byte(`) Loc() `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:118:35
	{
		_n, _err := _template.writeValue(
			_output,
			(LocationType),
			"file.template:118:35")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:118:48
	{
		_n, _err := _output.Write([]byte(` { return t.`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:118:60
	{
		_n, _err := _template.writeValue(
			_output,
			(LocationType),
			"file.template:118:60")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:118:73
	{
		_n, _err := _output.Write([]byte(` }

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:120:5
	{
		_n, _err := _template.writeValue(
			_output,
			(LexerType),
			"file.template:120:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:120:15
	{
		_n, _err := _output.Write([]byte(` interface {
    // Note: Return io.EOF to indicate end of stream
    // Token with unspecified value type should return *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:122:56
	{
		_n, _err := _template.writeValue(
			_output,
			(GenericSymbolType),
			"file.template:122:56")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:122:74
	{
		_n, _err := _output.Write([]byte(`
    Next() (`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:123:12
	{
		_n, _err := _template.writeValue(
			_output,
			(TokenType),
			"file.template:123:12")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:123:22
	{
		_n, _err := _output.Write([]byte(`, error)

    CurrentLocation() `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:125:22
	{
		_n, _err := _template.writeValue(
			_output,
			(LocationType),
			"file.template:125:22")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:125:35
	{
		_n, _err := _output.Write([]byte(`
}

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:128:5
	{
		_n, _err := _template.writeValue(
			_output,
			(ReducerType),
			"file.template:128:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:128:17
	{
		_n, _err := _output.Write([]byte(` interface {`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:129:0
	firstRule := true
	// file.template:130:0
	for _, rule := range Grammar.NonTerminals {
		// file.template:131:4
		if len(rule.Clauses) == 0 {
			// file.template:132:8
			continue
		}
		// file.template:135:4
		if !firstRule {
			// file.template:135:22
			{
				_n, _err := _output.Write([]byte(`
`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
		// file.template:138:4
		firstRule = false
		// file.template:140:4
		for clauseIdx, clause := range rule.Clauses {
			// file.template:141:8
			if clauseIdx > 0 {
				// file.template:141:29
				{
					_n, _err := _output.Write([]byte(`
`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
			}
			// file.template:145:8
			if clause.Label == "" {
				// file.template:145:34
				{
					_n, _err := _output.Write([]byte(`
    // `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:146:7
				{
					_n, _err := _template.writeValue(
						_output,
						(clause.LRLocation.ShortString()),
						"file.template:146:7")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:146:41
				{
					_n, _err := _output.Write([]byte(`: `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:146:43
				{
					_n, _err := _template.writeValue(
						_output,
						(rule.Name),
						"file.template:146:43")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:146:55
				{
					_n, _err := _output.Write([]byte(` -> ...`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
			} else {
				// file.template:147:17
				{
					_n, _err := _output.Write([]byte(`
    // `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:148:7
				{
					_n, _err := _template.writeValue(
						_output,
						(clause.LRLocation.ShortString()),
						"file.template:148:7")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:148:41
				{
					_n, _err := _output.Write([]byte(`: `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:148:43
				{
					_n, _err := _template.writeValue(
						_output,
						(rule.Name),
						"file.template:148:43")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:148:55
				{
					_n, _err := _output.Write([]byte(` -> `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:148:59
				{
					_n, _err := _template.writeValue(
						_output,
						(clause.Label),
						"file.template:148:59")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:148:74
				{
					_n, _err := _output.Write([]byte(`: ...`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
			}
			// file.template:151:8
			paramNameCount := map[string]int{}
			// file.template:151:49
			{
				_n, _err := _output.Write([]byte(`
    `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:152:4
			{
				_n, _err := _template.writeValue(
					_output,
					(clause.CodeGenReducerName),
					"file.template:152:4")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:152:32
			{
				_n, _err := _output.Write([]byte(`(`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:153:8
			for termIdx, term := range clause.Bindings {
				// file.template:155:12

				paramName := ""
				if term.SymbolId == parser.LRCharacterToken {
					paramName = "char"
				} else {
					// hack: append "_" to the end of the name ensures the
					// name is never a go keyword
					paramName = codegenutil.SnakeToCamel(term.Name) + "_"
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

				// file.template:178:0
				{
					_n, _err := _template.writeValue(
						_output,
						(paramName),
						"file.template:178:0")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:178:10
				{
					_n, _err := _output.Write([]byte(` `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:178:11
				{
					_n, _err := _template.writeValue(
						_output,
						(term.CodeGenType),
						"file.template:178:11")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:178:30
				{
					_n, _err := _template.writeValue(
						_output,
						(suffix),
						"file.template:178:30")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
			}
			// file.template:179:17
			{
				_n, _err := _output.Write([]byte(`) (`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:180:3
			{
				_n, _err := _template.writeValue(
					_output,
					(rule.CodeGenType),
					"file.template:180:3")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:180:22
			{
				_n, _err := _output.Write([]byte(`, error)`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
	}
	// file.template:182:8
	{
		_n, _err := _output.Write([]byte(`
}

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:185:5
	{
		_n, _err := _template.writeValue(
			_output,
			(ErrHandlerType),
			"file.template:185:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:185:20
	{
		_n, _err := _output.Write([]byte(` interface {
    Error(nextToken `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:186:20
	{
		_n, _err := _template.writeValue(
			_output,
			(TokenType),
			"file.template:186:20")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:186:30
	{
		_n, _err := _output.Write([]byte(`, parseStack `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:186:43
	{
		_n, _err := _template.writeValue(
			_output,
			(StackType),
			"file.template:186:43")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:186:53
	{
		_n, _err := _output.Write([]byte(`) error
}

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:189:5
	{
		_n, _err := _template.writeValue(
			_output,
			(DefaultErrHandlerType),
			"file.template:189:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:189:27
	{
		_n, _err := _output.Write([]byte(` struct {}

func (`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:191:6
	{
		_n, _err := _template.writeValue(
			_output,
			(DefaultErrHandlerType),
			"file.template:191:6")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:191:28
	{
		_n, _err := _output.Write([]byte(`) Error(nextToken `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:191:46
	{
		_n, _err := _template.writeValue(
			_output,
			(TokenType),
			"file.template:191:46")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:191:56
	{
		_n, _err := _output.Write([]byte(`, stack `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:191:64
	{
		_n, _err := _template.writeValue(
			_output,
			(StackType),
			"file.template:191:64")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:191:74
	{
		_n, _err := _output.Write([]byte(`) error {
    return `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:192:11
	{
		_n, _err := _template.writeValue(
			_output,
			(Errorf),
			"file.template:192:11")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:192:18
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
	// file.template:195:8
	{
		_n, _err := _template.writeValue(
			_output,
			(ExpectedTerminalsFunc),
			"file.template:195:8")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:195:30
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
	// file.template:199:5
	{
		_n, _err := _template.writeValue(
			_output,
			(ExpectedTerminalsFunc),
			"file.template:199:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:199:27
	{
		_n, _err := _output.Write([]byte(`(id `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:199:31
	{
		_n, _err := _template.writeValue(
			_output,
			(StateIdType),
			"file.template:199:31")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:199:43
	{
		_n, _err := _output.Write([]byte(`) []`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:199:47
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"file.template:199:47")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:199:60
	{
		_n, _err := _output.Write([]byte(` {
    result := []`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:200:16
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"file.template:200:16")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:200:29
	{
		_n, _err := _output.Write([]byte(`{}
    for key, _ := range `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:201:24
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionTable),
			"file.template:201:24")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:201:36
	{
		_n, _err := _output.Write([]byte(` {
        if key.`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:202:15
	{
		_n, _err := _template.writeValue(
			_output,
			(StateIdType),
			"file.template:202:15")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:202:27
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
	// file.template:205:36
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"file.template:205:36")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:205:49
	{
		_n, _err := _output.Write([]byte(`)
    }

    `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:208:4
	{
		_n, _err := _template.writeValue(
			_output,
			(SortSlice),
			"file.template:208:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:208:14
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
	// file.template:212:0
	for idx, start := range Grammar.Starts {
		// file.template:213:4

		parseSuffix := ""
		if len(Grammar.Starts) > 1 {
			parseSuffix = codegenutil.SnakeToCamel(start.Name)
		}

		// file.template:220:6
		{
			_n, _err := _output.Write([]byte(`
func `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:221:5
		{
			_n, _err := _template.writeValue(
				_output,
				(ParseFuncPrefix),
				"file.template:221:5")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:221:23
		{
			_n, _err := _template.writeValue(
				_output,
				(parseSuffix),
				"file.template:221:23")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:221:37
		{
			_n, _err := _output.Write([]byte(`(lexer `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:221:44
		{
			_n, _err := _template.writeValue(
				_output,
				(LexerType),
				"file.template:221:44")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:221:54
		{
			_n, _err := _output.Write([]byte(`, reducer `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:221:64
		{
			_n, _err := _template.writeValue(
				_output,
				(ReducerType),
				"file.template:221:64")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:221:76
		{
			_n, _err := _output.Write([]byte(`) (`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:223:0
		{
			_n, _err := _template.writeValue(
				_output,
				(start.CodeGenType),
				"file.template:223:0")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:223:20
		{
			_n, _err := _output.Write([]byte(`, error) {

    return `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:225:11
		{
			_n, _err := _template.writeValue(
				_output,
				(ParseFuncPrefix),
				"file.template:225:11")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:225:29
		{
			_n, _err := _template.writeValue(
				_output,
				(parseSuffix),
				"file.template:225:29")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:225:43
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
		// file.template:228:8
		{
			_n, _err := _template.writeValue(
				_output,
				(DefaultErrHandlerType),
				"file.template:228:8")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:228:30
		{
			_n, _err := _output.Write([]byte(`{})
}

func `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:231:5
		{
			_n, _err := _template.writeValue(
				_output,
				(ParseFuncPrefix),
				"file.template:231:5")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:231:23
		{
			_n, _err := _template.writeValue(
				_output,
				(parseSuffix),
				"file.template:231:23")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:231:37
		{
			_n, _err := _output.Write([]byte(`WithCustomErrorHandler(
    lexer `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:232:10
		{
			_n, _err := _template.writeValue(
				_output,
				(LexerType),
				"file.template:232:10")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:232:20
		{
			_n, _err := _output.Write([]byte(`,
    reducer `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:233:12
		{
			_n, _err := _template.writeValue(
				_output,
				(ReducerType),
				"file.template:233:12")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:233:24
		{
			_n, _err := _output.Write([]byte(`,
    errHandler `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:234:15
		{
			_n, _err := _template.writeValue(
				_output,
				(ErrHandlerType),
				"file.template:234:15")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:234:30
		{
			_n, _err := _output.Write([]byte(`) (
    `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:235:4
		{
			_n, _err := _template.writeValue(
				_output,
				(start.CodeGenType),
				"file.template:235:4")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:235:24
		{
			_n, _err := _output.Write([]byte(`,
    error) {

    item, err := `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:238:17
		{
			_n, _err := _template.writeValue(
				_output,
				(InternalParseFunc),
				"file.template:238:17")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:238:35
		{
			_n, _err := _output.Write([]byte(`(lexer, reducer, errHandler, `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:238:64
		{
			_n, _err := _template.writeValue(
				_output,
				(States.OrderedStates[idx].CodeGenConst),
				"file.template:238:64")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:239:47
		{
			_n, _err := _output.Write([]byte(`)
    if err != nil {
        var errRetVal `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:241:22
		{
			_n, _err := _template.writeValue(
				_output,
				(start.CodeGenType),
				"file.template:241:22")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:241:42
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
		// file.template:244:16
		{
			_n, _err := _template.writeValue(
				_output,
				(start.ValueType),
				"file.template:244:16")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:244:34
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
	// file.template:246:7
	{
		_n, _err := _output.Write([]byte(`

// ================================================================
// Parser internal implementation
// User should normally avoid directly accessing the following code
// ================================================================

func `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:253:5
	{
		_n, _err := _template.writeValue(
			_output,
			(InternalParseFunc),
			"file.template:253:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:253:23
	{
		_n, _err := _output.Write([]byte(`(
    lexer `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:254:10
	{
		_n, _err := _template.writeValue(
			_output,
			(LexerType),
			"file.template:254:10")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:254:20
	{
		_n, _err := _output.Write([]byte(`,
    reducer `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:255:12
	{
		_n, _err := _template.writeValue(
			_output,
			(ReducerType),
			"file.template:255:12")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:255:24
	{
		_n, _err := _output.Write([]byte(`,
    errHandler `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:256:15
	{
		_n, _err := _template.writeValue(
			_output,
			(ErrHandlerType),
			"file.template:256:15")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:256:30
	{
		_n, _err := _output.Write([]byte(`,
    startState `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:257:15
	{
		_n, _err := _template.writeValue(
			_output,
			(StateIdType),
			"file.template:257:15")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:257:27
	{
		_n, _err := _output.Write([]byte(`) (
    *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:258:5
	{
		_n, _err := _template.writeValue(
			_output,
			(StackItemType),
			"file.template:258:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:258:19
	{
		_n, _err := _output.Write([]byte(`,
    error) {

    stateStack := `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:261:18
	{
		_n, _err := _template.writeValue(
			_output,
			(StackType),
			"file.template:261:18")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:261:28
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
	// file.template:264:9
	{
		_n, _err := _template.writeValue(
			_output,
			(StackItemType),
			"file.template:264:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:264:23
	{
		_n, _err := _output.Write([]byte(`{startState, nil},
    }

    symbolStack := &`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:267:20
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolStackType),
			"file.template:267:20")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:267:36
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
	// file.template:275:22
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionTable),
			"file.template:275:22")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:275:34
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
	// file.template:282:32
	{
		_n, _err := _template.writeValue(
			_output,
			(ShiftAction),
			"file.template:282:32")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:282:44
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
	// file.template:289:39
	{
		_n, _err := _template.writeValue(
			_output,
			(ReduceAction),
			"file.template:289:39")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:289:52
	{
		_n, _err := _output.Write([]byte(` {
            var reduceSymbol *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:290:30
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:290:30")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:290:41
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
	// file.template:299:39
	{
		_n, _err := _template.writeValue(
			_output,
			(AcceptAction),
			"file.template:299:39")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:299:52
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

func (i `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:310:8
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"file.template:310:8")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:310:21
	{
		_n, _err := _output.Write([]byte(`) String() string {
    switch i {
    case `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:312:9
	{
		_n, _err := _template.writeValue(
			_output,
			(EndSymbolId),
			"file.template:312:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:312:21
	{
		_n, _err := _output.Write([]byte(`:
        return "$"
    case `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:314:9
	{
		_n, _err := _template.writeValue(
			_output,
			(WildcardSymbolId),
			"file.template:314:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:314:26
	{
		_n, _err := _output.Write([]byte(`:
        return "*"`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:316:0
	for _, symbolName := range OrderedSymbolNames[3:] {
		// file.template:317:4
		term := Grammar.Terms[symbolName]
		// file.template:317:44
		{
			_n, _err := _output.Write([]byte(`
    case `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:318:9
		{
			_n, _err := _template.writeValue(
				_output,
				(term.CodeGenSymbolConst),
				"file.template:318:9")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:318:35
		{
			_n, _err := _output.Write([]byte(`:`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:319:4
		if term.SymbolId == parser.LRCharacterToken {
			// file.template:320:8

			escaped := term.Name
			if term.Name == "'\"'" {
				escaped = "'\\\"'"
			} else if escaped[1] == '\\' {
				escaped = "'\\\\" + term.Name[2:]
			}

			// file.template:329:10
			{
				_n, _err := _output.Write([]byte(`
        return "`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:330:16
			{
				_n, _err := _template.writeValue(
					_output,
					(escaped),
					"file.template:330:16")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:330:24
			{
				_n, _err := _output.Write([]byte(`"`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		} else {
			// file.template:331:13
			{
				_n, _err := _output.Write([]byte(`
        return "`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:332:16
			{
				_n, _err := _template.writeValue(
					_output,
					(term.Name),
					"file.template:332:16")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:332:28
			{
				_n, _err := _output.Write([]byte(`"`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
	}
	// file.template:334:8
	{
		_n, _err := _output.Write([]byte(`
    default:
        return `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:336:15
	{
		_n, _err := _template.writeValue(
			_output,
			(Sprintf),
			"file.template:336:15")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:336:23
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
	// file.template:341:4
	{
		_n, _err := _template.writeValue(
			_output,
			(EndSymbolId),
			"file.template:341:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:341:16
	{
		_n, _err := _output.Write([]byte(` = `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:341:19
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"file.template:341:19")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:341:32
	{
		_n, _err := _output.Write([]byte(`(0)
    `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:342:4
	{
		_n, _err := _template.writeValue(
			_output,
			(WildcardSymbolId),
			"file.template:342:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:342:21
	{
		_n, _err := _output.Write([]byte(` = `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:342:24
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"file.template:342:24")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:342:37
	{
		_n, _err := _output.Write([]byte(`(-1)
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:344:0
	for idx, term := range Grammar.NonTerminals {
		// file.template:344:48
		{
			_n, _err := _output.Write([]byte(`
    `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:345:4
		{
			_n, _err := _template.writeValue(
				_output,
				(term.CodeGenSymbolConst),
				"file.template:345:4")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:345:30
		{
			_n, _err := _output.Write([]byte(` = `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:345:33
		{
			_n, _err := _template.writeValue(
				_output,
				(SymbolIdType),
				"file.template:345:33")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:345:46
		{
			_n, _err := _output.Write([]byte(`(`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:345:47
		{
			_n, _err := _template.writeValue(
				_output,
				(256 + len(Grammar.Terminals) + idx),
				"file.template:345:47")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:345:80
		{
			_n, _err := _output.Write([]byte(`)`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// file.template:346:8
	{
		_n, _err := _output.Write([]byte(`
)

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:349:5
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionIdType),
			"file.template:349:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:349:18
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
	// file.template:353:4
	{
		_n, _err := _template.writeValue(
			_output,
			(ShiftAction),
			"file.template:353:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:353:16
	{
		_n, _err := _output.Write([]byte(` = `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:353:19
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionIdType),
			"file.template:353:19")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:353:32
	{
		_n, _err := _output.Write([]byte(`(0)
    `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:354:4
	{
		_n, _err := _template.writeValue(
			_output,
			(ReduceAction),
			"file.template:354:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:354:17
	{
		_n, _err := _output.Write([]byte(` = `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:354:20
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionIdType),
			"file.template:354:20")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:354:33
	{
		_n, _err := _output.Write([]byte(`(1)
    `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:355:4
	{
		_n, _err := _template.writeValue(
			_output,
			(AcceptAction),
			"file.template:355:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:355:17
	{
		_n, _err := _output.Write([]byte(` = `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:355:20
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionIdType),
			"file.template:355:20")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:355:33
	{
		_n, _err := _output.Write([]byte(`(2)
)

func (i `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:358:8
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionIdType),
			"file.template:358:8")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:358:21
	{
		_n, _err := _output.Write([]byte(`) String() string {
    switch i {
    case `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:360:9
	{
		_n, _err := _template.writeValue(
			_output,
			(ShiftAction),
			"file.template:360:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:360:21
	{
		_n, _err := _output.Write([]byte(`:
        return "shift"
    case `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:362:9
	{
		_n, _err := _template.writeValue(
			_output,
			(ReduceAction),
			"file.template:362:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:362:22
	{
		_n, _err := _output.Write([]byte(`:
        return "reduce"
    case `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:364:9
	{
		_n, _err := _template.writeValue(
			_output,
			(AcceptAction),
			"file.template:364:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:364:22
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
	// file.template:367:15
	{
		_n, _err := _template.writeValue(
			_output,
			(Sprintf),
			"file.template:367:15")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:367:23
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
	// file.template:371:5
	{
		_n, _err := _template.writeValue(
			_output,
			(ReduceType),
			"file.template:371:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:371:16
	{
		_n, _err := _output.Write([]byte(` int

const (`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:374:0
	clauseIdx := 1
	// file.template:375:0
	for _, rule := range Grammar.NonTerminals {
		// file.template:376:4
		for _, clause := range rule.Clauses {
			// file.template:376:44
			{
				_n, _err := _output.Write([]byte(`
    `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:377:4
			{
				_n, _err := _template.writeValue(
					_output,
					(clause.CodeGenReducerNameConst),
					"file.template:377:4")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:377:37
			{
				_n, _err := _output.Write([]byte(` = `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:377:40
			{
				_n, _err := _template.writeValue(
					_output,
					(ReduceType),
					"file.template:377:40")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:377:51
			{
				_n, _err := _output.Write([]byte(`(`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:377:52
			{
				_n, _err := _template.writeValue(
					_output,
					(clauseIdx),
					"file.template:377:52")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:377:62
			{
				_n, _err := _output.Write([]byte(`)`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:378:8
			clauseIdx += 1
		}
	}
	// file.template:380:8
	{
		_n, _err := _output.Write([]byte(`
)

func (i `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:383:8
	{
		_n, _err := _template.writeValue(
			_output,
			(ReduceType),
			"file.template:383:8")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:383:19
	{
		_n, _err := _output.Write([]byte(`) String() string {
    switch i {`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:385:0
	for _, rule := range Grammar.NonTerminals {
		// file.template:386:4
		for _, clause := range rule.Clauses {
			// file.template:386:44
			{
				_n, _err := _output.Write([]byte(`
    case `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:387:9
			{
				_n, _err := _template.writeValue(
					_output,
					(clause.CodeGenReducerNameConst),
					"file.template:387:9")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:387:42
			{
				_n, _err := _output.Write([]byte(`:
        return "`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:388:16
			{
				_n, _err := _template.writeValue(
					_output,
					(clause.CodeGenReducerName),
					"file.template:388:16")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:388:44
			{
				_n, _err := _output.Write([]byte(`"`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
	}
	// file.template:390:8
	{
		_n, _err := _output.Write([]byte(`
    default:
        return `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:392:15
	{
		_n, _err := _template.writeValue(
			_output,
			(Sprintf),
			"file.template:392:15")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:392:23
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
	// file.template:396:5
	{
		_n, _err := _template.writeValue(
			_output,
			(StateIdType),
			"file.template:396:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:396:17
	{
		_n, _err := _output.Write([]byte(` int

func (id `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:398:9
	{
		_n, _err := _template.writeValue(
			_output,
			(StateIdType),
			"file.template:398:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:398:21
	{
		_n, _err := _output.Write([]byte(`) String() string {
    return `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:399:11
	{
		_n, _err := _template.writeValue(
			_output,
			(Sprintf),
			"file.template:399:11")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:399:19
	{
		_n, _err := _output.Write([]byte(`("State %d", int(id))
}

const (`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:403:0
	for _, state := range States.OrderedStates {
		// file.template:403:47
		{
			_n, _err := _output.Write([]byte(`
    `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:404:4
		{
			_n, _err := _template.writeValue(
				_output,
				(state.CodeGenConst),
				"file.template:404:4")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:404:25
		{
			_n, _err := _output.Write([]byte(` = `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:404:28
		{
			_n, _err := _template.writeValue(
				_output,
				(StateIdType),
				"file.template:404:28")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:404:40
		{
			_n, _err := _output.Write([]byte(`(`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:404:41
		{
			_n, _err := _template.writeValue(
				_output,
				(state.StateNum),
				"file.template:404:41")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:404:58
		{
			_n, _err := _output.Write([]byte(`)`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// file.template:405:8
	{
		_n, _err := _output.Write([]byte(`
)

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:408:5
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:408:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:408:16
	{
		_n, _err := _output.Write([]byte(` struct {
    SymbolId_ `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:409:14
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"file.template:409:14")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:409:27
	{
		_n, _err := _output.Write([]byte(`

    Generic_ *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:411:14
	{
		_n, _err := _template.writeValue(
			_output,
			(GenericSymbolType),
			"file.template:411:14")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:411:32
	{
		_n, _err := _output.Write([]byte(`
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:413:0
	for _, valueType := range OrderedValueTypes {
		// file.template:414:4
		if valueType.Name == lr.Generic {
			// file.template:415:8
			continue
		}
		// file.template:416:12
		{
			_n, _err := _output.Write([]byte(`
    `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:417:4
		{
			_n, _err := _template.writeValue(
				_output,
				(valueType.Name),
				"file.template:417:4")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:417:21
		{
			_n, _err := _output.Write([]byte(` `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:417:22
		{
			_n, _err := _template.writeValue(
				_output,
				(valueType.ParamType),
				"file.template:417:22")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// file.template:418:8
	{
		_n, _err := _output.Write([]byte(`
}
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:421:0

	valueTerms := map[string][]*lr.Term{}
	for _, symbolName := range OrderedSymbolNames[2:] {
		term := Grammar.Terms[symbolName]
		valueTerms[term.ValueType] = append(valueTerms[term.ValueType], term)
	}

	// file.template:429:3
	{
		_n, _err := _output.Write([]byte(`func NewSymbol(token `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:430:21
	{
		_n, _err := _template.writeValue(
			_output,
			(TokenType),
			"file.template:430:21")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:430:31
	{
		_n, _err := _output.Write([]byte(`) (*`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:430:35
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:430:35")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:430:46
	{
		_n, _err := _output.Write([]byte(`, error) {
    symbol, ok := token.(*`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:431:26
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:431:26")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:431:37
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
	// file.template:436:14
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:436:14")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:436:25
	{
		_n, _err := _output.Write([]byte(`{SymbolId_: token.Id()}
    switch token.Id() {`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:438:0
	for _, valueType := range OrderedValueTypes {
		// file.template:439:4

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

		// file.template:454:6
		{
			_n, _err := _output.Write([]byte(`
    case `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:456:4
		for idx, kconst := range consts {
			// file.template:457:0
			{
				_n, _err := _template.writeValue(
					_output,
					(kconst),
					"file.template:457:0")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:457:8
			if idx != len(consts)-1 {
				// file.template:457:37
				{
					_n, _err := _output.Write([]byte(`, `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
			}
		}
		// file.template:458:13
		{
			_n, _err := _output.Write([]byte(`:
        val, ok := token.(`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:460:26
		{
			_n, _err := _template.writeValue(
				_output,
				(valueType.ParamType),
				"file.template:460:26")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:460:48
		{
			_n, _err := _output.Write([]byte(`)
        if !ok {
            return nil, `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:462:24
		{
			_n, _err := _template.writeValue(
				_output,
				(Errorf),
				"file.template:462:24")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:462:31
		{
			_n, _err := _output.Write([]byte(`(
                "Invalid value type for token %s.  "+
                    "Expecting `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:464:31
		{
			_n, _err := _template.writeValue(
				_output,
				(valueType.ParamType),
				"file.template:464:31")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:464:53
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
		// file.template:468:15
		{
			_n, _err := _template.writeValue(
				_output,
				(valueType.Name),
				"file.template:468:15")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:468:32
		{
			_n, _err := _output.Write([]byte(` = val`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// file.template:469:8
	{
		_n, _err := _output.Write([]byte(`
    default:
        return nil, `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:471:20
	{
		_n, _err := _template.writeValue(
			_output,
			(Errorf),
			"file.template:471:20")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:471:27
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
	// file.template:476:9
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:476:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:476:20
	{
		_n, _err := _output.Write([]byte(`) Id() `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:476:27
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"file.template:476:27")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:476:40
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
	// file.template:480:9
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:480:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:480:20
	{
		_n, _err := _output.Write([]byte(`) Loc() `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:480:28
	{
		_n, _err := _template.writeValue(
			_output,
			(LocationType),
			"file.template:480:28")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:480:41
	{
		_n, _err := _output.Write([]byte(` {
    type locator interface { Loc() `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:481:35
	{
		_n, _err := _template.writeValue(
			_output,
			(LocationType),
			"file.template:481:35")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:481:48
	{
		_n, _err := _output.Write([]byte(` }
    switch s.SymbolId_ {`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:483:0
	for _, field := range OrderedValueTypes {
		// file.template:484:4
		if field.Name == lr.Generic {
			// file.template:485:8
			continue
		}
		// file.template:487:4
		terms := valueTerms[field.Name]
		// file.template:487:42
		{
			_n, _err := _output.Write([]byte(`
    case `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:489:4
		for idx, term := range terms {
			// file.template:490:0
			{
				_n, _err := _template.writeValue(
					_output,
					(term.CodeGenSymbolConst),
					"file.template:490:0")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:491:8
			if idx != len(terms)-1 {
				// file.template:491:37
				{
					_n, _err := _output.Write([]byte(`, `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
			}
		}
		// file.template:492:13
		{
			_n, _err := _output.Write([]byte(`:
        loc, ok := interface{}(s.`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:494:33
		{
			_n, _err := _template.writeValue(
				_output,
				(field.Name),
				"file.template:494:33")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:494:46
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
	// file.template:498:8
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
	// file.template:503:11
	{
		_n, _err := _template.writeValue(
			_output,
			(LocationType),
			"file.template:503:11")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:503:24
	{
		_n, _err := _output.Write([]byte(`{}
}

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:506:5
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolStackType),
			"file.template:506:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:506:21
	{
		_n, _err := _output.Write([]byte(` struct {
    lexer `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:507:10
	{
		_n, _err := _template.writeValue(
			_output,
			(LexerType),
			"file.template:507:10")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:507:20
	{
		_n, _err := _output.Write([]byte(`
    top []*`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:508:11
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:508:11")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:508:22
	{
		_n, _err := _output.Write([]byte(`
}

func (stack *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:511:13
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolStackType),
			"file.template:511:13")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:511:29
	{
		_n, _err := _output.Write([]byte(`) Top() (*`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:511:39
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:511:39")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:511:50
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
	// file.template:515:22
	{
		_n, _err := _template.writeValue(
			_output,
			(EOF),
			"file.template:515:22")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:515:26
	{
		_n, _err := _output.Write([]byte(` {
                return nil, `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:516:28
	{
		_n, _err := _template.writeValue(
			_output,
			(Errorf),
			"file.template:516:28")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:516:35
	{
		_n, _err := _output.Write([]byte(`("Unexpected lex error: %s", err)
            }
            token = &`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:518:21
	{
		_n, _err := _template.writeValue(
			_output,
			(GenericSymbolType),
			"file.template:518:21")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:518:39
	{
		_n, _err := _output.Write([]byte(`{`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:518:40
	{
		_n, _err := _template.writeValue(
			_output,
			(EndSymbolId),
			"file.template:518:40")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:518:52
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
	// file.template:529:13
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolStackType),
			"file.template:529:13")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:529:29
	{
		_n, _err := _output.Write([]byte(`) Push(symbol *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:529:44
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:529:44")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:529:55
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
	// file.template:533:13
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolStackType),
			"file.template:533:13")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:533:29
	{
		_n, _err := _output.Write([]byte(`) Pop() (*`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:533:39
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:533:39")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:533:50
	{
		_n, _err := _output.Write([]byte(`, error) {
    if len(stack.top) == 0 {
        return nil, `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:535:20
	{
		_n, _err := _template.writeValue(
			_output,
			(Errorf),
			"file.template:535:20")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:535:27
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
	// file.template:542:5
	{
		_n, _err := _template.writeValue(
			_output,
			(StackItemType),
			"file.template:542:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:542:19
	{
		_n, _err := _output.Write([]byte(` struct {
    StateId `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:543:12
	{
		_n, _err := _template.writeValue(
			_output,
			(StateIdType),
			"file.template:543:12")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:543:24
	{
		_n, _err := _output.Write([]byte(`

    *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:545:5
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:545:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:545:16
	{
		_n, _err := _output.Write([]byte(`
}

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:548:5
	{
		_n, _err := _template.writeValue(
			_output,
			(StackType),
			"file.template:548:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:548:15
	{
		_n, _err := _output.Write([]byte(` []*`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:548:19
	{
		_n, _err := _template.writeValue(
			_output,
			(StackItemType),
			"file.template:548:19")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:548:33
	{
		_n, _err := _output.Write([]byte(`

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:550:5
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionType),
			"file.template:550:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:550:16
	{
		_n, _err := _output.Write([]byte(` struct {
    ActionType `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:551:15
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionIdType),
			"file.template:551:15")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:551:28
	{
		_n, _err := _output.Write([]byte(`

    ShiftStateId `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:553:17
	{
		_n, _err := _template.writeValue(
			_output,
			(StateIdType),
			"file.template:553:17")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:553:29
	{
		_n, _err := _output.Write([]byte(`
    ReduceType `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:554:15
	{
		_n, _err := _template.writeValue(
			_output,
			(ReduceType),
			"file.template:554:15")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:554:26
	{
		_n, _err := _output.Write([]byte(`
}

func (act *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:557:11
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionType),
			"file.template:557:11")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:557:22
	{
		_n, _err := _output.Write([]byte(`) ShiftItem(symbol *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:557:42
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:557:42")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:557:53
	{
		_n, _err := _output.Write([]byte(`) *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:557:56
	{
		_n, _err := _template.writeValue(
			_output,
			(StackItemType),
			"file.template:557:56")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:557:70
	{
		_n, _err := _output.Write([]byte(` {
    return &`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:558:12
	{
		_n, _err := _template.writeValue(
			_output,
			(StackItemType),
			"file.template:558:12")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:558:26
	{
		_n, _err := _output.Write([]byte(`{StateId: act.ShiftStateId, `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:558:54
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:558:54")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:558:65
	{
		_n, _err := _output.Write([]byte(`: symbol}
}

func (act *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:561:11
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionType),
			"file.template:561:11")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:561:22
	{
		_n, _err := _output.Write([]byte(`) ReduceSymbol(
    reducer `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:562:12
	{
		_n, _err := _template.writeValue(
			_output,
			(ReducerType),
			"file.template:562:12")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:562:24
	{
		_n, _err := _output.Write([]byte(`,
    stack `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:563:10
	{
		_n, _err := _template.writeValue(
			_output,
			(StackType),
			"file.template:563:10")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:563:20
	{
		_n, _err := _output.Write([]byte(`) (
    `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:564:4
	{
		_n, _err := _template.writeValue(
			_output,
			(StackType),
			"file.template:564:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:564:14
	{
		_n, _err := _output.Write([]byte(`,
    *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:565:5
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:565:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:565:16
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
	// file.template:569:15
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:569:15")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:569:26
	{
		_n, _err := _output.Write([]byte(`{}
    switch act.ReduceType {`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:571:0
	for _, rule := range Grammar.NonTerminals {
		// file.template:572:4
		for _, clause := range rule.Clauses {
			// file.template:572:44
			{
				_n, _err := _output.Write([]byte(`
    case `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:573:9
			{
				_n, _err := _template.writeValue(
					_output,
					(clause.CodeGenReducerNameConst),
					"file.template:573:9")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:573:42
			{
				_n, _err := _output.Write([]byte(`:`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:574:8
			if len(clause.Bindings) > 0 {
				// file.template:574:40
				{
					_n, _err := _output.Write([]byte(`
        args := stack[len(stack)-`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:575:33
				{
					_n, _err := _template.writeValue(
						_output,
						(len(clause.Bindings)),
						"file.template:575:33")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:575:56
				{
					_n, _err := _output.Write([]byte(`:]
        stack = stack[:len(stack)-`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:576:34
				{
					_n, _err := _template.writeValue(
						_output,
						(len(clause.Bindings)),
						"file.template:576:34")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:576:57
				{
					_n, _err := _output.Write([]byte(`]`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
			}
			// file.template:577:16
			{
				_n, _err := _output.Write([]byte(`
        symbol.SymbolId_ = `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:578:27
			{
				_n, _err := _template.writeValue(
					_output,
					(rule.CodeGenSymbolConst),
					"file.template:578:27")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:578:53
			{
				_n, _err := _output.Write([]byte(`
        symbol.`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:579:15
			{
				_n, _err := _template.writeValue(
					_output,
					(rule.ValueType),
					"file.template:579:15")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:579:32
			{
				_n, _err := _output.Write([]byte(`, err = reducer.`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:579:48
			{
				_n, _err := _template.writeValue(
					_output,
					(clause.CodeGenReducerName),
					"file.template:579:48")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:579:76
			{
				_n, _err := _output.Write([]byte(`(`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:580:8
			for idx, term := range clause.Bindings {
				// file.template:580:52
				{
					_n, _err := _output.Write([]byte(`args[`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:581:5
				{
					_n, _err := _template.writeValue(
						_output,
						(idx),
						"file.template:581:5")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:581:9
				{
					_n, _err := _output.Write([]byte(`].`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:581:11
				{
					_n, _err := _template.writeValue(
						_output,
						(term.ValueType),
						"file.template:581:11")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:582:12
				if idx != len(clause.Bindings)-1 {
					// file.template:582:49
					{
						_n, _err := _output.Write([]byte(`,`))
						_numWritten += int64(_n)
						if _err != nil {
							return _numWritten, _err
						}
					}
				}
			}
			// file.template:583:17
			{
				_n, _err := _output.Write([]byte(`)`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
	}
	// file.template:586:8
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
	// file.template:592:14
	{
		_n, _err := _template.writeValue(
			_output,
			(Errorf),
			"file.template:592:14")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:592:21
	{
		_n, _err := _output.Write([]byte(`("Unexpected %s reduce error: %s", act.ReduceType, err)
    }

    return stack, symbol, err
}

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:598:5
	{
		_n, _err := _template.writeValue(
			_output,
			(TableKeyType),
			"file.template:598:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:598:18
	{
		_n, _err := _output.Write([]byte(` struct {
    `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:599:4
	{
		_n, _err := _template.writeValue(
			_output,
			(StateIdType),
			"file.template:599:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:599:16
	{
		_n, _err := _output.Write([]byte(`
    `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:600:4
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"file.template:600:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:600:17
	{
		_n, _err := _output.Write([]byte(`
}

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:603:5
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionTableType),
			"file.template:603:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:603:21
	{
		_n, _err := _output.Write([]byte(` map[`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:603:26
	{
		_n, _err := _template.writeValue(
			_output,
			(TableKeyType),
			"file.template:603:26")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:603:39
	{
		_n, _err := _output.Write([]byte(`]*`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:603:41
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionType),
			"file.template:603:41")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:603:52
	{
		_n, _err := _output.Write([]byte(`

func (table `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:605:12
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionTableType),
			"file.template:605:12")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:605:28
	{
		_n, _err := _output.Write([]byte(`) Get(
    stateId `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:606:12
	{
		_n, _err := _template.writeValue(
			_output,
			(StateIdType),
			"file.template:606:12")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:606:24
	{
		_n, _err := _output.Write([]byte(`,
    symbolId `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:607:13
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"file.template:607:13")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:607:26
	{
		_n, _err := _output.Write([]byte(`) (
    *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:608:5
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionType),
			"file.template:608:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:608:16
	{
		_n, _err := _output.Write([]byte(`,
    bool) {

    action, ok := table[`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:611:24
	{
		_n, _err := _template.writeValue(
			_output,
			(TableKeyType),
			"file.template:611:24")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:611:37
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
	// file.template:616:23
	{
		_n, _err := _template.writeValue(
			_output,
			(TableKeyType),
			"file.template:616:23")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:616:36
	{
		_n, _err := _output.Write([]byte(`{stateId, `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:616:46
	{
		_n, _err := _template.writeValue(
			_output,
			(WildcardSymbolId),
			"file.template:616:46")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:616:63
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
	// file.template:621:0
	for _, state := range States.OrderedStates {
		// file.template:621:47
		{
			_n, _err := _output.Write([]byte(`
    `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:622:4
		{
			_n, _err := _template.writeValue(
				_output,
				(state.CodeGenAction),
				"file.template:622:4")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:622:26
		{
			_n, _err := _output.Write([]byte(` = &`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:622:30
		{
			_n, _err := _template.writeValue(
				_output,
				(ActionType),
				"file.template:622:30")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:622:41
		{
			_n, _err := _output.Write([]byte(`{`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:622:42
		{
			_n, _err := _template.writeValue(
				_output,
				(ShiftAction),
				"file.template:622:42")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:622:54
		{
			_n, _err := _output.Write([]byte(`, `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:622:56
		{
			_n, _err := _template.writeValue(
				_output,
				(state.CodeGenConst),
				"file.template:622:56")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:622:77
		{
			_n, _err := _output.Write([]byte(`, 0}`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// file.template:625:0
	for _, term := range Grammar.NonTerminals {
		// file.template:626:4
		for _, clause := range term.Clauses {
			// file.template:626:44
			{
				_n, _err := _output.Write([]byte(`
    `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:627:4
			{
				_n, _err := _template.writeValue(
					_output,
					(clause.CodeGenReduceAction),
					"file.template:627:4")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:627:33
			{
				_n, _err := _output.Write([]byte(` = &`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:627:37
			{
				_n, _err := _template.writeValue(
					_output,
					(ActionType),
					"file.template:627:37")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:627:48
			{
				_n, _err := _output.Write([]byte(`{`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:627:49
			{
				_n, _err := _template.writeValue(
					_output,
					(ReduceAction),
					"file.template:627:49")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:627:62
			{
				_n, _err := _output.Write([]byte(`, 0, `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:627:67
			{
				_n, _err := _template.writeValue(
					_output,
					(clause.CodeGenReducerNameConst),
					"file.template:627:67")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:627:100
			{
				_n, _err := _output.Write([]byte(`}`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
	}
	// file.template:629:8
	{
		_n, _err := _output.Write([]byte(`
)

var `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:632:4
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionTable),
			"file.template:632:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:632:16
	{
		_n, _err := _output.Write([]byte(` = `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:632:19
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionTableType),
			"file.template:632:19")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:632:35
	{
		_n, _err := _output.Write([]byte(`{`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:633:0
	for _, state := range States.OrderedStates {
		// file.template:634:4
		for _, item := range state.Items {
			// file.template:635:8

			if !item.IsReduce {
				continue
			}

			if item.Name != lr.AcceptRule || item.LookAhead != lr.EndMarker {
				continue
			}

			// file.template:645:10
			{
				_n, _err := _output.Write([]byte(`
    {`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:646:5
			{
				_n, _err := _template.writeValue(
					_output,
					(state.CodeGenConst),
					"file.template:646:5")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:646:26
			{
				_n, _err := _output.Write([]byte(`, `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:646:28
			{
				_n, _err := _template.writeValue(
					_output,
					(EndSymbolId),
					"file.template:646:28")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:646:40
			{
				_n, _err := _output.Write([]byte(`}: &`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:646:44
			{
				_n, _err := _template.writeValue(
					_output,
					(ActionType),
					"file.template:646:44")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:646:55
			{
				_n, _err := _output.Write([]byte(`{`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:646:56
			{
				_n, _err := _template.writeValue(
					_output,
					(AcceptAction),
					"file.template:646:56")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:646:69
			{
				_n, _err := _output.Write([]byte(`, 0, 0},`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
	}
	// file.template:650:0
	for _, state := range States.OrderedStates {
		// file.template:651:4
		for _, symbolName := range OrderedSymbolNames {
			// file.template:652:8

			symbol := Grammar.Terms[symbolName]
			child, ok := state.Goto[symbolName]
			if !ok {
				continue
			}

			// file.template:660:10
			{
				_n, _err := _output.Write([]byte(`
    {`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:661:5
			{
				_n, _err := _template.writeValue(
					_output,
					(state.CodeGenConst),
					"file.template:661:5")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:661:26
			{
				_n, _err := _output.Write([]byte(`, `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:661:28
			{
				_n, _err := _template.writeValue(
					_output,
					(symbol.CodeGenSymbolConst),
					"file.template:661:28")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:661:56
			{
				_n, _err := _output.Write([]byte(`}: `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:661:59
			{
				_n, _err := _template.writeValue(
					_output,
					(child.CodeGenAction),
					"file.template:661:59")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:661:81
			{
				_n, _err := _output.Write([]byte(`,`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
	}
	// file.template:665:0
	for _, state := range States.OrderedStates {
		// file.template:666:4
		for _, item := range state.Items {
			// file.template:667:8

			if !item.IsReduce {
				continue
			}

			if item.Name == lr.AcceptRule && item.LookAhead == lr.EndMarker {
				continue
			}

			idConst := Grammar.Terms[item.LookAhead].CodeGenSymbolConst
			reduceAction := item.Clause.CodeGenReduceAction

			// file.template:680:10
			{
				_n, _err := _output.Write([]byte(`
    {`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:681:5
			{
				_n, _err := _template.writeValue(
					_output,
					(state.CodeGenConst),
					"file.template:681:5")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:681:26
			{
				_n, _err := _output.Write([]byte(`, `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:681:28
			{
				_n, _err := _template.writeValue(
					_output,
					(idConst),
					"file.template:681:28")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:681:36
			{
				_n, _err := _output.Write([]byte(`}: `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:681:39
			{
				_n, _err := _template.writeValue(
					_output,
					(reduceAction),
					"file.template:681:39")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:681:52
			{
				_n, _err := _output.Write([]byte(`,`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
	}
	// file.template:683:8
	{
		_n, _err := _output.Write([]byte(`
}

/*
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:687:0
	{
		_n, _err := (&debug_template.File{
			OrderedSymbolNames:        OrderedSymbolNames,
			States:                    States,
			OutputDebugNonKernelItems: OutputDebugNonKernelItems,
		}).WriteTo(_output)
		_numWritten += _n
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:693:3
	{
		_n, _err := _output.Write([]byte(`*/
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}

	return _numWritten, nil
}
