package code_gen

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	lr "github.com/pattyshack/abc/src/lr/internal"
	"github.com/pattyshack/abc/src/lr/internal/code_gen/go_template"
	"github.com/pattyshack/abc/src/lr/internal/parser"
)

var (
	escapedChar = map[string]byte{
		"'\\t'":  '\t',
		"'\\n'":  '\n',
		"'\\''":  '\'',
		"'\\\\'": '\\',
	}

	nameRe = regexp.MustCompile(`((?:(?:\[\])*\**)*)(?:(.+)\.)?(\w+(?:{})?)$`)
)

type importEntry struct {
	path  string
	alias string
}

type importObject struct {
	*importEntry

	prefix string // * or &
	name   string
}

func (obj *importObject) String() string {
	if obj.alias == "" {
		return obj.prefix + obj.name
	}

	return obj.prefix + obj.alias + "." + obj.name
}

type goImports struct {
	imports map[string]*importEntry
}

func newGoImports() *goImports {
	return &goImports{
		imports: map[string]*importEntry{},
	}
}

// This supports accessing objects of the form:
//     *(\[\])*(\*)*)*(<full module path>\.)?<object>({})?
// map objects are not supported
func (imports *goImports) Obj(fullName string) *importObject {
	match := nameRe.FindStringSubmatch(fullName)
	if match == nil {
		panic("Invalid fullName: " + fullName)
	}

	prefix := match[1]
	modulePath := match[2]
	name := match[3]

	if modulePath == "" {
		return &importObject{&importEntry{}, prefix, name}
	}

	entry, ok := imports.imports[modulePath]
	if !ok {
		entry = &importEntry{
			path: modulePath,
		}

		imports.imports[modulePath] = entry
	}

	return &importObject{entry, prefix, name}
}

func (imports *goImports) assignAlias() error {
	aliasCount := map[string]int{}
	for pkg, entry := range imports.imports {
		base := filepath.Base(pkg)
		if base == "." || base == "/" {
			return fmt.Errorf("Invalid import path: %s", pkg)
		}

		alias := base
		aliasCount[alias] += 1
		if aliasCount[alias] > 1 {
			alias = fmt.Sprintf("%s%d", alias, aliasCount[alias])
		}

		entry.alias = alias
	}

	return nil
}

func (imports *goImports) WriteTo(output io.Writer) (int64, error) {
	err := imports.assignAlias()
	if err != nil {
		return 0, err
	}

	builder := NewCodeBuilder()
	if len(imports.imports) > 0 {
		builder.Line("import (")
		builder.PushIndent()
		// Maybe separate built-in packages from other packages
		for _, entry := range imports.imports {
			builder.Line("%s \"%s\"", entry.alias, entry.path)
		}
		builder.PopIndent()
		builder.Line(")")
		builder.Line("")
	}

	return builder.WriteTo(output)
}

type GoCodeBuilder struct {
	*go_template.File
}

func (cb *GoCodeBuilder) WriteTo(output io.Writer) (int64, error) {
	buffer := bytes.NewBuffer(nil)

	_, err := cb.File.WriteTo(buffer)
	if err != nil {
		return 0, err
	}

	formatted, err := format.Source(buffer.Bytes())
	if err != nil {
		return 0, fmt.Errorf(
			"Failed to format (%s) generated code:\n%s",
			err,
			buffer.Bytes())
	}

	n, err := output.Write(formatted)
	return int64(n), err
}

// TODO handle this more gracefully
func snakeToCamel(str string) string {
	chunks := strings.Split(str, "_")

	result := ""
	for _, chunk := range chunks {
		result += strings.Title(strings.ToLower(chunk))
	}

	return result
}

type goCodeGen struct {
	*lr.Grammar
	*lr.GoSpec

	*lr.LRStates

	*goImports

	nameLocs map[string]parser.LRLocation

	endSymbol      string
	wildcardSymbol string

	genericSymbol string
}

func newGoCodeGen(
	grammar *lr.Grammar,
	states *lr.LRStates) (
	*goCodeGen,
	error) {

	cfg := grammar.LangSpecs.Go
	if cfg == nil {
		return nil, fmt.Errorf("go configuration not specified in lang_specs")
	}

	if cfg.Package == "" {
		return nil, fmt.Errorf("package name not specified")
	}

	return &goCodeGen{
		Grammar:   grammar,
		GoSpec:    cfg,
		LRStates:  states,
		goImports: newGoImports(),
		nameLocs:  map[string]parser.LRLocation{},
	}, nil
}

func (gen *goCodeGen) check(name string, loc parser.LRLocation) error {
	prev, ok := gen.nameLocs[name]
	if ok {
		return fmt.Errorf(
			"Generated conflicting name: %s (from %s and %s)",
			name,
			prev,
			loc)
	}
	gen.nameLocs[name] = loc

	return nil
}

func (gen *goCodeGen) populateCodeGenVariables() error {
	gen.endSymbol = "_" + gen.Prefix + "EndMarker"
	gen.wildcardSymbol = "_" + gen.Prefix + "WildcardMarker"
	gen.genericSymbol = gen.Prefix + "GenericSymbol"

	for _, term := range gen.Terms {
		valueType := gen.ValueTypes[term.ValueType]
		if valueType == "" {
			if term.ValueType != lr.Generic {
				return fmt.Errorf(
					"Undefined value type for <%s> %s",
					term.ValueType,
					term.LRLocation)
			}
			valueType = "*" + gen.genericSymbol
		}

		term.CodeGenType = gen.Obj(valueType)

		symbolConst := ""
		if term.SymbolId == parser.LRCharacterToken {
			symbolConst = term.Name
		} else {
			symbolConst = gen.Prefix + snakeToCamel(term.Name)
			if term.IsTerminal {
				symbolConst += "Token"
			} else {
				symbolConst += "Type"
			}
		}

		err := gen.check(symbolConst, term.LRLocation)
		if err != nil {
			return err
		}

		term.CodeGenSymbolConst = symbolConst

		for _, clause := range term.Clauses {
			reducerName := snakeToCamel(clause.Label) + "To" +
				snakeToCamel(term.Name)

			err := gen.check(reducerName, clause.LRLocation)
			if err != nil {
				return err
			}

			clause.CodeGenReducerName = reducerName

			reducerConst := "_" + gen.Prefix + "Reduce" + reducerName
			err = gen.check(reducerConst, clause.LRLocation)
			if err != nil {
				return err
			}

			clause.CodeGenReducerNameConst = reducerConst

			actionConst := reducerConst + "Action"
			err = gen.check(actionConst, clause.LRLocation)
			if err != nil {
				return err
			}

			clause.CodeGenReduceAction = actionConst
		}
	}

	for _, state := range gen.OrderedStates {
		state.CodeGenConst = fmt.Sprintf(
			"_%sState%d",
			gen.Prefix,
			state.StateNum)

		actionConst := fmt.Sprintf(
			"_%sGotoState%dAction",
			gen.Prefix,
			state.StateNum)
		err := gen.check(actionConst, parser.LRLocation{})
		if err != nil {
			return err
		}

		state.CodeGenAction = actionConst
	}

	return nil
}

func GenerateGoLRCode(
	grammar *lr.Grammar,
	states *lr.LRStates) (
	io.WriterTo,
	error) {

	cfg := grammar.LangSpecs.Go
	if cfg == nil {
		return nil, fmt.Errorf("go configuration not specified in lang_specs")
	}

	if cfg.Package == "" {
		return nil, fmt.Errorf("package name not specified")
	}

	gen, err := newGoCodeGen(grammar, states)
	if err != nil {
		return nil, err
	}

	err = gen.populateCodeGenVariables()
	if err != nil {
		return nil, err
	}

	orderedValueTypes := lr.ParamList{
		&lr.Param{lr.Generic, gen.Obj("*" + gen.genericSymbol)},
	}
	for name, valueType := range gen.ValueTypes {
		orderedValueTypes = append(
			orderedValueTypes,
			&lr.Param{name, gen.Obj(valueType)})
	}
	sort.Sort(orderedValueTypes)

	genericPtr := "*" + gen.genericSymbol
	orderedSymbols := []*lr.Term{
		&lr.Term{
			Name:               lr.StartMarker,
			IsTerminal:         true,
			ValueType:          lr.Generic,
			CodeGenSymbolConst: lr.StartMarker,
			CodeGenType:        genericPtr,
		},
		&lr.Term{
			Name:               lr.Wildcard,
			IsTerminal:         true,
			ValueType:          lr.Generic,
			CodeGenSymbolConst: gen.wildcardSymbol,
			CodeGenType:        genericPtr,
		},
		&lr.Term{
			Name:               lr.EndMarker,
			IsTerminal:         true,
			ValueType:          lr.Generic,
			CodeGenSymbolConst: gen.endSymbol,
			CodeGenType:        genericPtr,
		},
	}
	orderedSymbols = append(orderedSymbols, gen.Terminals...)
	orderedSymbols = append(orderedSymbols, gen.NonTerminals...)

	symbols := make(map[string]*lr.Term, len(orderedSymbols))
	for _, symbol := range orderedSymbols {
		symbols[symbol.Name] = symbol
	}

	file := &go_template.File{
		Source:                    grammar.Source,
		Package:                   cfg.Package,
		Imports:                   gen.goImports,
		ActionType:                "_" + cfg.Prefix + "Action",
		ActionIdType:              "_" + cfg.Prefix + "ActionType",
		ShiftAction:               "_" + cfg.Prefix + "ShiftAction",
		ReduceAction:              "_" + cfg.Prefix + "ReduceAction",
		AcceptAction:              "_" + cfg.Prefix + "AcceptAction",
		StateIdType:               "_" + cfg.Prefix + "StateId",
		ReduceType:                "_" + cfg.Prefix + "ReduceType",
		SymbolType:                cfg.Prefix + "Symbol",
		GenericSymbolType:         gen.genericSymbol,
		StackItemType:             "_" + cfg.Prefix + "StackItem",
		StackType:                 "_" + cfg.Prefix + "Stack",
		SymbolStackType:           "_" + cfg.Prefix + "PseudoSymbolStack",
		SymbolIdType:              cfg.Prefix + "SymbolId",
		EndSymbolId:               gen.endSymbol,
		WildcardSymbolId:          gen.wildcardSymbol,
		LocationType:              cfg.Prefix + "Location",
		TokenType:                 cfg.Prefix + "Token",
		LexerType:                 cfg.Prefix + "Lexer",
		ReducerType:               cfg.Prefix + "Reducer",
		ErrHandlerType:            cfg.Prefix + "ParseErrorHandler",
		DefaultErrHandlerType:     cfg.Prefix + "DefaultParseErrorHandler",
		ExpectedTerminalsFunc:     cfg.Prefix + "ExpectedTerminals",
		ParseFuncPrefix:           cfg.Prefix + "Parse",
		InternalParseFunc:         "_" + cfg.Prefix + "Parse",
		TableKeyType:              "_" + cfg.Prefix + "ActionTableKey",
		ActionTableType:           "_" + cfg.Prefix + "ActionTableType",
		ActionTable:               "_" + cfg.Prefix + "ActionTable",
		SortSlice:                 gen.Obj("sort.Slice"),
		Sprintf:                   gen.Obj("fmt.Sprintf"),
		Errorf:                    gen.Obj("fmt.Errorf"),
		EOF:                       gen.Obj("io.EOF"),
		OrderedSymbols:            orderedSymbols,
		Symbols:                   symbols,
		StartSymbols:              grammar.Starts,
		OrderedStates:             states.OrderedStates,
		OrderedValueTypes:         orderedValueTypes,
		OutputDebugNonKernelItems: cfg.OutputDebugNonKernelItems,
	}

	return &GoCodeBuilder{file}, nil
}
