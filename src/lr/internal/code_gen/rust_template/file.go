// Auto-generated from source: file.template

package rust_template

import (
	_fmt "fmt"
	_io "io"

	"fmt"
	"sort"
	"strings"

	lr "github.com/pattyshack/abc/src/lr/internal"
	"github.com/pattyshack/abc/src/lr/internal/code_gen/debug_template"
	parser "github.com/pattyshack/abc/src/lr/internal/parser"
)

type File struct {
	Grammar            *lr.Grammar
	States             *lr.LRStates
	Cfg                *lr.RustSpec
	OrderedSymbolNames []string
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

	Grammar := _template.Grammar
	States := _template.States
	Cfg := _template.Cfg
	OrderedSymbolNames := _template.OrderedSymbolNames

	// file.template:22:0

	sortedValueDecls := []string{}
	for name, actualType := range Cfg.ValueTypes {
		sortedValueDecls = append(sortedValueDecls, name+"("+actualType+"),")
	}
	sort.Strings(sortedValueDecls)

	rustType := func(name string) string {
		actualType := Cfg.ValueTypes[name]
		if actualType != "" {
			return actualType
		}

		return "Symbol"
	}

	rustSymbolKind := func(termName string) string {
		if termName == lr.StartMarker {
			return "SymbolKind::_StartParseMarker"
		} else if termName == lr.Wildcard {
			return "SymbolKind::_WildcardMarker"
		} else if termName == lr.EndMarker {
			return "SymbolKind::EofToken"
		}

		term := Grammar.Terms[termName]
		if term == nil {
			panic("Invalid term name: " + termName)
		}

		if !term.IsTerminal {
			return "SymbolKind::" + term.CodeGenSymbolConst + "Type"
		}

		if term.SymbolId != parser.LRIdentifierToken {
			return "SymbolKind::AsciiCharToken(" + term.Name + ")"
		}

		return "SymbolKind::" + term.CodeGenSymbolConst + "Token"
	}

	// file.template:62:3
	{
		_n, _err := _output.Write([]byte(`// Auto-generated from source: `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:63:31
	{
		_n, _err := _template.writeValue(
			_output,
			(Grammar.Source),
			"file.template:63:31")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:63:48
	{
		_n, _err := _output.Write([]byte(`

use std::error;
use std::fmt;

`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:68:0
	for _, entry := range Cfg.UseList {
		// file.template:68:38
		{
			_n, _err := _output.Write([]byte(`use `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:69:4
		{
			_n, _err := _template.writeValue(
				_output,
				(entry),
				"file.template:69:4")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:69:10
		{
			_n, _err := _output.Write([]byte(`;`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// file.template:70:8
	{
		_n, _err := _output.Write([]byte(`

#[derive(Debug, Clone)]
pub enum SymbolKind {
  //
  // Token symbols.
  //
  EofToken,
  AsciiCharToken(char),
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:80:0
	for _, term := range Grammar.Terminals {
		// file.template:81:2
		if term.SymbolId == parser.LRIdentifierToken {
			// file.template:81:51
			{
				_n, _err := _output.Write([]byte(`
  `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:82:2
			{
				_n, _err := _template.writeValue(
					_output,
					(term.CodeGenSymbolConst),
					"file.template:82:2")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:82:28
			{
				_n, _err := _output.Write([]byte(`Token,`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
	}
	// file.template:84:8
	{
		_n, _err := _output.Write([]byte(`

  //
  // Type symbols.
  //
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:89:0
	for _, term := range Grammar.NonTerminals {
		// file.template:89:46
		{
			_n, _err := _output.Write([]byte(`  `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:90:2
		{
			_n, _err := _template.writeValue(
				_output,
				(term.CodeGenSymbolConst),
				"file.template:90:2")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:90:28
		{
			_n, _err := _output.Write([]byte(`Type,
`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// file.template:91:8
	{
		_n, _err := _output.Write([]byte(`
  //
  // For internal use only.
  //
  _StartParseMarker,
  _WildcardMarker,
}

impl SymbolKind {
  fn to_string(&self) -> String {
    match self {
      SymbolKind::_StartParseMarker => "^",
      SymbolKind::_WildcardMarker => "*",
      SymbolKind::EofToken => "$",
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:106:0
	for _, term := range Grammar.Terminals {
		// file.template:106:42
		{
			_n, _err := _output.Write([]byte(`
      `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:107:6
		{
			_n, _err := _template.writeValue(
				_output,
				(rustSymbolKind(term.Name)),
				"file.template:107:6")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:107:34
		{
			_n, _err := _output.Write([]byte(` => "`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:107:39
		{
			_n, _err := _template.writeValue(
				_output,
				(term.Name),
				"file.template:107:39")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:107:51
		{
			_n, _err := _output.Write([]byte(`",`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// file.template:108:8
	{
		_n, _err := _output.Write([]byte(`
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:109:0
	for _, term := range Grammar.NonTerminals {
		// file.template:109:45
		{
			_n, _err := _output.Write([]byte(`
      `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:110:6
		{
			_n, _err := _template.writeValue(
				_output,
				(rustSymbolKind(term.Name)),
				"file.template:110:6")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:110:34
		{
			_n, _err := _output.Write([]byte(` => "`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:110:39
		{
			_n, _err := _template.writeValue(
				_output,
				(term.Name),
				"file.template:110:39")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:110:51
		{
			_n, _err := _output.Write([]byte(`",`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// file.template:111:8
	{
		_n, _err := _output.Write([]byte(`
      SymbolKind::AsciiCharToken(c) => panic!("Unexpected token '{}'", c),
    }.to_string()
  }
}

#[derive(Debug)]
pub enum SymbolData {
  // Note: %token without value declaration must have Nil data.
  Nil,
  // Note: %type without value declaration must have Any data.
  Any(Box<Symbol>),
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:124:0
	for _, decl := range sortedValueDecls {
		// file.template:124:42
		{
			_n, _err := _output.Write([]byte(`
  `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:125:2
		{
			_n, _err := _template.writeValue(
				_output,
				(decl),
				"file.template:125:2")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// file.template:126:8
	{
		_n, _err := _output.Write([]byte(`
}

#[derive(Debug)]
pub struct Symbol {
  pub kind: SymbolKind,
  pub data: SymbolData,
}

impl Symbol {
  pub fn validate(&self) -> Result<(), Box<dyn error::Error>> {
    match self {
      Symbol{
        kind: SymbolKind::EofToken,
        data: _,
      } => (),`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:142:0
	for _, term := range Grammar.Terminals {
		// file.template:143:2

		enumData := "SymbolData::Nil"
		if term.ValueType != lr.Generic {
			enumData = "SymbolData::" + term.ValueType + "(_)"
		}

		// file.template:148:4
		{
			_n, _err := _output.Write([]byte(`
      Symbol{
        kind: `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:150:14
		{
			_n, _err := _template.writeValue(
				_output,
				(rustSymbolKind(term.Name)),
				"file.template:150:14")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:150:42
		{
			_n, _err := _output.Write([]byte(`,
        data: `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:151:14
		{
			_n, _err := _template.writeValue(
				_output,
				(enumData),
				"file.template:151:14")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:151:23
		{
			_n, _err := _output.Write([]byte(`,
      } => (),`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// file.template:154:0
	for _, term := range Grammar.NonTerminals {
		// file.template:155:2

		enumData := "SymbolData::Any(_)"
		if term.ValueType != lr.Generic {
			enumData = "SymbolData::" + term.ValueType + "(_)"
		}

		// file.template:160:4
		{
			_n, _err := _output.Write([]byte(`
      Symbol{
        kind: `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:162:14
		{
			_n, _err := _template.writeValue(
				_output,
				(rustSymbolKind(term.Name)),
				"file.template:162:14")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:162:42
		{
			_n, _err := _output.Write([]byte(`,
        data: `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:163:14
		{
			_n, _err := _template.writeValue(
				_output,
				(enumData),
				"file.template:163:14")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:163:23
		{
			_n, _err := _output.Write([]byte(`,
      } => (),`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// file.template:165:8
	{
		_n, _err := _output.Write([]byte(`
      _ => return Err(Box::new(Error::new(format!(
        "Unexpected symbol {:?}",
        self)))),
    };

    Ok(())
  }
}

pub trait Lexer {
  fn next(&mut self) -> Result<Symbol, Box<dyn error::Error>>;
}

pub trait Reducer {`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:180:0
	for _, rule := range Grammar.NonTerminals {
		// file.template:181:2
		if len(rule.Clauses) == 0 {
			// file.template:182:4
			continue
		}
		// file.template:185:2
		for _, clause := range rule.Clauses {
			// file.template:185:42
			{
				_n, _err := _output.Write([]byte(`
  // `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:186:5
			{
				_n, _err := _template.writeValue(
					_output,
					(clause.LRLocation.ShortString()),
					"file.template:186:5")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:186:39
			{
				_n, _err := _output.Write([]byte(`: `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:186:41
			{
				_n, _err := _template.writeValue(
					_output,
					(rule.Name),
					"file.template:186:41")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:186:53
			{
				_n, _err := _output.Write([]byte(` ->`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:187:4
			if clause.Label == "" {
				// file.template:187:31
				{
					_n, _err := _output.Write([]byte(` ...`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
			} else {
				// file.template:189:14
				{
					_n, _err := _output.Write([]byte(` `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:190:1
				{
					_n, _err := _template.writeValue(
						_output,
						(clause.Label),
						"file.template:190:1")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:190:16
				{
					_n, _err := _output.Write([]byte(`: ...`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
			}
			// file.template:191:12
			{
				_n, _err := _output.Write([]byte(`
  fn `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:192:5
			{
				_n, _err := _template.writeValue(
					_output,
					(clause.CodeGenReducerName),
					"file.template:192:5")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:192:33
			{
				_n, _err := _output.Write([]byte(`(&self, `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:193:4
			paramNameCount := map[string]int{}
			// file.template:194:4
			for termIdx, term := range clause.Bindings {
				// file.template:195:6

				// hack: append "_" to name to ensure the name isn't a keyword.
				paramName := "char_"
				if term.SymbolId != parser.LRCharacterToken {
					paramName = strings.ToLower(term.Name) + "_"
				}

				paramNameCount[paramName] += 1
				cnt := paramNameCount[paramName]
				if cnt > 1 {
					paramName = fmt.Sprintf("%s%d", paramName, cnt)
				}

				// file.template:208:0
				{
					_n, _err := _template.writeValue(
						_output,
						(paramName),
						"file.template:208:0")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:208:10
				{
					_n, _err := _output.Write([]byte(`: `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:208:12
				{
					_n, _err := _template.writeValue(
						_output,
						(rustType(term.ValueType)),
						"file.template:208:12")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:209:6
				if termIdx != len(clause.Bindings)-1 {
					// file.template:209:50
					{
						_n, _err := _output.Write([]byte(`, `))
						_numWritten += int64(_n)
						if _err != nil {
							return _numWritten, _err
						}
					}
				}
			}
			// file.template:212:13
			{
				_n, _err := _output.Write([]byte(`) -> Result<`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:213:12
			{
				_n, _err := _template.writeValue(
					_output,
					(rustType(rule.ValueType)),
					"file.template:213:12")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:213:39
			{
				_n, _err := _output.Write([]byte(`, Box<dyn error::Error>>;`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
		// file.template:214:10
		{
			_n, _err := _output.Write([]byte(`
`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// file.template:215:8
	{
		_n, _err := _output.Write([]byte(`}

#[derive(Debug)]
pub struct Error {
  msg: String,
}

impl Error {
  pub fn new(msg: String) -> Self {
    Self{
      msg: msg,
    }
  }
}

impl fmt::Display for Error {
  fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
    write!(f, "{}", self.msg)
  }
}

impl error::Error for Error {
}
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:240:0
	for idx, start := range Grammar.Starts {
		// file.template:240:44
		{
			_n, _err := _output.Write([]byte(`
pub fn parse_`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:241:13
		{
			_n, _err := _template.writeValue(
				_output,
				(start.Name),
				"file.template:241:13")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:241:26
		{
			_n, _err := _output.Write([]byte(`<L: Lexer, R: Reducer>(
  lexer: &mut L,
  reducer: &R,
) -> Result<`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:244:12
		{
			_n, _err := _template.writeValue(
				_output,
				(rustType(start.ValueType)),
				"file.template:244:12")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:244:40
		{
			_n, _err := _output.Write([]byte(`, Box<dyn error::Error>> {
  let result = parse(
    lexer,
    reducer,
    StateId::`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:248:13
		{
			_n, _err := _template.writeValue(
				_output,
				(States.OrderedStates[idx].CodeGenConst),
				"file.template:248:13")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:248:54
		{
			_n, _err := _output.Write([]byte(`)?;

  match result {`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:251:2
		if start.ValueType == lr.Generic {
			// file.template:251:39
			{
				_n, _err := _output.Write([]byte(`
    SymbolData::Any(val) => Ok(*val),`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		} else {
			// file.template:253:11
			{
				_n, _err := _output.Write([]byte(`
    SymbolData::`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:254:16
			{
				_n, _err := _template.writeValue(
					_output,
					(start.ValueType),
					"file.template:254:16")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:254:34
			{
				_n, _err := _output.Write([]byte(`(val) => return Ok(val),`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
		// file.template:255:10
		{
			_n, _err := _output.Write([]byte(`
    _ => (),
  }

  panic!("Invalid symbol data type. This should never happen");
}`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// file.template:261:8
	{
		_n, _err := _output.Write([]byte(`

// ==============================
// Parser internal implementation
// ==============================

fn parse<L: Lexer, R: Reducer>(
  lexer: &mut L,
  reducer: &R,
  start_state: StateId,
) -> Result<SymbolData, Box<dyn error::Error>> {

  let mut state_stack = vec![
    ParseStackFrame{
      state_id: start_state,
      symbol: Symbol{
        kind: SymbolKind::_StartParseMarker,
        data: SymbolData::Nil,
      }
    },
  ];

  let mut symbol_stack = SymbolStack::new(lexer);

  loop {
    let next_symbol_kind = symbol_stack.peek()?;

    let current_state_id = state_stack[state_stack.len() - 1].state_id.clone();

    match lookup_action(current_state_id, next_symbol_kind) {
      Action::Goto(next_state_id) => {
        state_stack.push(ParseStackFrame{
          state_id: next_state_id,
          symbol: symbol_stack.pop(),
        })
      },
      Action::Reduce(reduce_kind) => {
        symbol_stack.push(reduce_symbol(
            reduce_kind,
            reducer,
            &mut state_stack)?)?;
      },
      Action::Accept => {
        assert_eq!(state_stack.len(), 2, "This should never happen");
        return Ok(state_stack.pop().unwrap().symbol.data)
      },
      Action::Error => return Err(new_syntax_error(
        symbol_stack.pop(),
        state_stack.pop().unwrap().state_id)),
    }
  }
}

fn reduce_symbol<R: Reducer>(
  reduce_kind: ReduceKind,
  reducer: &R,
  state_stack: &mut Vec<ParseStackFrame>,
) -> Result<Symbol, Box<dyn error::Error>> {
  let reduced = match reduce_kind {`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:320:0
	for _, rule := range Grammar.NonTerminals {
		// file.template:321:2
		for _, clause := range rule.Clauses {
			// file.template:321:42
			{
				_n, _err := _output.Write([]byte(`
    ReduceKind::`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:322:16
			{
				_n, _err := _template.writeValue(
					_output,
					(clause.CodeGenReducerNameConst),
					"file.template:322:16")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:322:49
			{
				_n, _err := _output.Write([]byte(` => {`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:323:6
			for idx := len(clause.Bindings) - 1; idx >= 0; idx -= 1 {
				// file.template:323:67
				{
					_n, _err := _output.Write([]byte(`
      let symbol`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:324:16
				{
					_n, _err := _template.writeValue(
						_output,
						(idx),
						"file.template:324:16")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:324:20
				{
					_n, _err := _output.Write([]byte(` = state_stack.pop().unwrap().symbol;`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:325:8
				term := clause.Bindings[idx]
				// file.template:326:8
				if term.ValueType == lr.Generic {
					// file.template:326:44
					{
						_n, _err := _output.Write([]byte(`
      let arg`))
						_numWritten += int64(_n)
						if _err != nil {
							return _numWritten, _err
						}
					}
					// file.template:327:13
					{
						_n, _err := _template.writeValue(
							_output,
							(idx),
							"file.template:327:13")
						_numWritten += int64(_n)
						if _err != nil {
							return _numWritten, _err
						}
					}
					// file.template:327:17
					{
						_n, _err := _output.Write([]byte(` = symbol`))
						_numWritten += int64(_n)
						if _err != nil {
							return _numWritten, _err
						}
					}
					// file.template:327:26
					{
						_n, _err := _template.writeValue(
							_output,
							(idx),
							"file.template:327:26")
						_numWritten += int64(_n)
						if _err != nil {
							return _numWritten, _err
						}
					}
					// file.template:327:30
					{
						_n, _err := _output.Write([]byte(`;`))
						_numWritten += int64(_n)
						if _err != nil {
							return _numWritten, _err
						}
					}
				} else {
					// file.template:328:17
					{
						_n, _err := _output.Write([]byte(`
      let arg`))
						_numWritten += int64(_n)
						if _err != nil {
							return _numWritten, _err
						}
					}
					// file.template:329:13
					{
						_n, _err := _template.writeValue(
							_output,
							(idx),
							"file.template:329:13")
						_numWritten += int64(_n)
						if _err != nil {
							return _numWritten, _err
						}
					}
					// file.template:329:17
					{
						_n, _err := _output.Write([]byte(` = match symbol`))
						_numWritten += int64(_n)
						if _err != nil {
							return _numWritten, _err
						}
					}
					// file.template:329:32
					{
						_n, _err := _template.writeValue(
							_output,
							(idx),
							"file.template:329:32")
						_numWritten += int64(_n)
						if _err != nil {
							return _numWritten, _err
						}
					}
					// file.template:329:36
					{
						_n, _err := _output.Write([]byte(` {
        Symbol{data: SymbolData::`))
						_numWritten += int64(_n)
						if _err != nil {
							return _numWritten, _err
						}
					}
					// file.template:330:33
					{
						_n, _err := _template.writeValue(
							_output,
							(term.ValueType),
							"file.template:330:33")
						_numWritten += int64(_n)
						if _err != nil {
							return _numWritten, _err
						}
					}
					// file.template:330:50
					{
						_n, _err := _output.Write([]byte(`(val), ..} => val,
        _ => panic!("Failed to extract argument.  This should never happen"),
      };`))
						_numWritten += int64(_n)
						if _err != nil {
							return _numWritten, _err
						}
					}
				}
				// file.template:333:16
				{
					_n, _err := _output.Write([]byte(`
`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
			}
			// file.template:335:14
			{
				_n, _err := _output.Write([]byte(`
      let result = reducer.`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:336:27
			{
				_n, _err := _template.writeValue(
					_output,
					(clause.CodeGenReducerName),
					"file.template:336:27")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:336:55
			{
				_n, _err := _output.Write([]byte(`(`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:337:6
			for idx := 0; idx < len(clause.Bindings); idx += 1 {
				// file.template:337:63
				{
					_n, _err := _output.Write([]byte(`arg`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:338:3
				{
					_n, _err := _template.writeValue(
						_output,
						(idx),
						"file.template:338:3")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:339:8
				if idx != len(clause.Bindings)-1 {
					// file.template:339:49
					{
						_n, _err := _output.Write([]byte(`, `))
						_numWritten += int64(_n)
						if _err != nil {
							return _numWritten, _err
						}
					}
				}
			}
			// file.template:342:15
			{
				_n, _err := _output.Write([]byte(`)?;
      Symbol{
        kind: `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:345:14
			{
				_n, _err := _template.writeValue(
					_output,
					(rustSymbolKind(rule.Name)),
					"file.template:345:14")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:345:42
			{
				_n, _err := _output.Write([]byte(`,`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:346:6
			if rule.ValueType == lr.Generic {
				// file.template:346:42
				{
					_n, _err := _output.Write([]byte(`
        data: SymbolData::Any(Box::new(result)),`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
			} else {
				// file.template:348:15
				{
					_n, _err := _output.Write([]byte(`
        data: SymbolData::`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:349:26
				{
					_n, _err := _template.writeValue(
						_output,
						(rule.ValueType),
						"file.template:349:26")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:349:43
				{
					_n, _err := _output.Write([]byte(`(result),`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
			}
			// file.template:350:14
			{
				_n, _err := _output.Write([]byte(`
      }
    },`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
	}
	// file.template:354:8
	{
		_n, _err := _output.Write([]byte(`
  };

  Ok(reduced)
}

struct SymbolStack<'a, L: Lexer> {
  lexer: &'a mut L,
  top: Vec<Symbol>,
}

impl<'a, L: Lexer> SymbolStack<'a, L> {
  fn new(lexer: &'a mut L) -> Self {
    Self{
      lexer: lexer,
      top: Vec::new(),
    }
  }

  fn peek(&mut self) -> Result<SymbolKind, Box<dyn error::Error>> {
    if self.top.is_empty() {
      let symbol = self.lexer.next()?;
      let _ = symbol.validate()?;
      self.top.push(symbol);
    }

    Ok(self.top[self.top.len()-1].kind.clone())
  }

  fn push(&mut self, symbol: Symbol) -> Result<(), Box<dyn error::Error>> {
    symbol.validate()?;
    self.top.push(symbol);
    Ok(())
  }

  fn pop(&mut self) -> Symbol {
    if self.top.is_empty() {
      panic!("Cannot pop an empty symbol stack. This should never happen");
    }

    self.top.pop().unwrap()
  }
}

#[derive(Clone, Debug)]
enum StateId {`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:400:0
	for _, state := range States.OrderedStates {
		// file.template:400:47
		{
			_n, _err := _output.Write([]byte(`
  `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:401:2
		{
			_n, _err := _template.writeValue(
				_output,
				(state.CodeGenConst),
				"file.template:401:2")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:401:23
		{
			_n, _err := _output.Write([]byte(`,`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// file.template:402:8
	{
		_n, _err := _output.Write([]byte(`
}

struct ParseStackFrame {
  state_id: StateId,
  symbol: Symbol,
}

#[derive(Clone, Debug)]
enum ReduceKind {`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:412:0
	for _, term := range Grammar.NonTerminals {
		// file.template:413:2
		for _, clause := range term.Clauses {
			// file.template:413:42
			{
				_n, _err := _output.Write([]byte(`
  `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:414:2
			{
				_n, _err := _template.writeValue(
					_output,
					(clause.CodeGenReducerNameConst),
					"file.template:414:2")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:414:35
			{
				_n, _err := _output.Write([]byte(`,`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
	}
	// file.template:416:8
	{
		_n, _err := _output.Write([]byte(`
}

#[derive(Clone, Debug)]
enum Action {
  Goto(StateId),
  Reduce(ReduceKind),
  Accept,
  Error,
}

fn lookup_action(current_state: StateId, next_symbol: SymbolKind) -> Action {
  match current_state {`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:429:0
	for _, state := range States.OrderedStates {
		// file.template:429:47
		{
			_n, _err := _output.Write([]byte(`
    StateId::`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:430:13
		{
			_n, _err := _template.writeValue(
				_output,
				(state.CodeGenConst),
				"file.template:430:13")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:430:34
		{
			_n, _err := _output.Write([]byte(` => match next_symbol {`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:431:2

		wildcard := "_ => Action::Error,"

		nextState := state.Goto[lr.Wildcard]
		if nextState != nil {
			wildcard = "_ => Action::Goto(StateId::" + nextState.CodeGenConst + "),"
		}

		// file.template:441:2
		for _, item := range state.Items {
			// file.template:442:4

			if !item.IsReduce {
				continue
			}

			if item.Name != lr.AcceptRule {
				continue
			}

			if item.LookAhead != lr.EndMarker {
				panic("Expecting end marker. This should never happen")
			}

			// file.template:456:6
			{
				_n, _err := _output.Write([]byte(`
      `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:457:6
			{
				_n, _err := _template.writeValue(
					_output,
					(rustSymbolKind(item.LookAhead)),
					"file.template:457:6")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:457:39
			{
				_n, _err := _output.Write([]byte(` => Action::Accept,`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
		// file.template:460:2
		for _, name := range OrderedSymbolNames[2:] {
			// file.template:461:4

			nextState := state.Goto[name]
			if nextState == nil {
				continue
			}

			// file.template:468:6
			{
				_n, _err := _output.Write([]byte(`
      `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:469:6
			{
				_n, _err := _template.writeValue(
					_output,
					(rustSymbolKind(name)),
					"file.template:469:6")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:469:29
			{
				_n, _err := _output.Write([]byte(` => Action::Goto(StateId::`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:469:55
			{
				_n, _err := _template.writeValue(
					_output,
					(nextState.CodeGenConst),
					"file.template:469:55")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:469:80
			{
				_n, _err := _output.Write([]byte(`),`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
		// file.template:472:2
		for _, item := range state.Items {
			// file.template:473:4

			if !item.IsReduce {
				continue
			}

			if item.Name == lr.AcceptRule {
				continue
			}

			reduce := "Action::Reduce(ReduceKind::" + item.Clause.CodeGenReducerNameConst + "),"

			if item.LookAhead == lr.Wildcard {
				wildcard = "_ => " + reduce
				continue
			}

			// file.template:490:6
			{
				_n, _err := _output.Write([]byte(`
      `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:491:6
			{
				_n, _err := _template.writeValue(
					_output,
					(rustSymbolKind(item.LookAhead)),
					"file.template:491:6")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:491:39
			{
				_n, _err := _output.Write([]byte(` => `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:491:43
			{
				_n, _err := _template.writeValue(
					_output,
					(reduce),
					"file.template:491:43")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
		// file.template:492:10
		{
			_n, _err := _output.Write([]byte(`
      `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:493:6
		{
			_n, _err := _template.writeValue(
				_output,
				(wildcard),
				"file.template:493:6")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:493:17
		{
			_n, _err := _output.Write([]byte(`
    },`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// file.template:495:8
	{
		_n, _err := _output.Write([]byte(`
  }
}
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:499:0

	/*
	   symboDebugName := func(name string) string {
	     if name == lr.StartMarker {
	       return "^"
	     }
	     if name == lr.Wildcard {
	       return "*"
	     }
	     if name == lr.EndMarker {
	       return "$"
	     }

	     return name
	   }
	*/

	// file.template:515:3
	{
		_n, _err := _output.Write([]byte(`fn new_syntax_error(
  next_symbol: Symbol,
  current_state_id: StateId,
) -> Box<dyn error::Error> {
  let expected_terminals = match current_state_id {`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:521:0
	for _, state := range States.OrderedStates {
		// file.template:521:47
		{
			_n, _err := _output.Write([]byte(`
    StateId::`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:522:13
		{
			_n, _err := _template.writeValue(
				_output,
				(state.CodeGenConst),
				"file.template:522:13")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:522:34
		{
			_n, _err := _output.Write([]byte(` => "`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:523:2
		first := true
		// file.template:524:2
		for _, item := range state.Items {
			// file.template:525:4
			if !item.IsReduce {
				// file.template:526:6
				continue
			}
			// file.template:529:4
			if !first {
				// file.template:529:19
				{
					_n, _err := _output.Write([]byte(`, `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
			}
			// file.template:532:4
			first = false
			// file.template:533:0
			{
				_n, _err := _template.writeValue(
					_output,
					(item.LookAhead),
					"file.template:533:0")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
		// file.template:536:2
		for _, name := range OrderedSymbolNames[2:] {
			// file.template:537:4
			if _, ok := state.Goto[name]; ok {
				// file.template:538:6
				if !first {
					// file.template:538:21
					{
						_n, _err := _output.Write([]byte(`, `))
						_numWritten += int64(_n)
						if _err != nil {
							return _numWritten, _err
						}
					}
				}
				// file.template:541:6
				first = false
				// file.template:542:0
				{
					_n, _err := _template.writeValue(
						_output,
						(name),
						"file.template:542:0")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
			}
		}
		// file.template:544:11
		{
			_n, _err := _output.Write([]byte(`",`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// file.template:546:9
	{
		_n, _err := _output.Write([]byte(`  };

  Box::new(Error::new(format!(
    "Syntax error: unexpected symbol {}. Expecting [{}]",
    next_symbol.kind.to_string(),
    expected_terminals)))
}

/*
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:556:0
	{
		_n, _err := (&debug_template.File{
			OrderedSymbolNames:        OrderedSymbolNames,
			States:                    States,
			OutputDebugNonKernelItems: Cfg.OutputDebugNonKernelItems,
		}).WriteTo(_output)
		_numWritten += _n
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:562:3
	{
		_n, _err := _output.Write([]byte(`]
*/
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}

	return _numWritten, nil
}
