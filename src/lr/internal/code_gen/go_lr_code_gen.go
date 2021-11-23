package code_gen

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
	"sort"
	"strings"

	lr "github.com/pattyshack/abc/src/lr/internal"
	"github.com/pattyshack/abc/src/lr/internal/code_gen/go_template"
	"github.com/pattyshack/abc/src/lr/internal/parser"
)

var escapedChar = map[string]byte{
	"'\\t'":  '\t',
	"'\\n'":  '\n',
	"'\\''":  '\'',
	"'\\\\'": '\\',
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

	*goHeader

	nameLocs map[string]parser.LRLocation

	location string

	symbolId string
	stateId  string

	symbol string

	endSymbol      string
	wildcardSymbol string

	token   string
	lexer   string
	reducer string

	errHandler        string
	defaultErrHandler string
	expectedTerminals string

	genericSymbol string

	symbolStack string

	stackItem string
	stack     string

	reduceType string
	actionType string

	shiftAction  string
	reduceAction string
	acceptAction string

	action string

	tableKey        string
	actionTableType string
	actionTable     string

	parse string
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
		Grammar:  grammar,
		GoSpec:   cfg,
		LRStates: states,
		goHeader: newGoHeader(),
		nameLocs: map[string]parser.LRLocation{},
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
	gen.location = gen.Prefix + "Location"
	gen.symbolId = gen.Prefix + "SymbolId"
	gen.symbol = gen.Prefix + "Symbol"
	gen.stateId = "_" + gen.Prefix + "StateId"
	gen.endSymbol = "_" + gen.Prefix + "EndMarker"
	gen.wildcardSymbol = "_" + gen.Prefix + "WildcardMarker"
	gen.token = gen.Prefix + "Token"
	gen.lexer = gen.Prefix + "Lexer"
	gen.reducer = gen.Prefix + "Reducer"
	gen.errHandler = gen.Prefix + "ParseErrorHandler"
	gen.defaultErrHandler = gen.Prefix + "DefaultParseErrorHandler"
	gen.expectedTerminals = gen.Prefix + "ExpectedTerminals"
	gen.genericSymbol = gen.Prefix + "GenericSymbol"
	gen.symbolStack = "_" + gen.Prefix + "PseudoSymbolStack"
	gen.stackItem = "_" + gen.Prefix + "StackItem"
	gen.stack = "_" + gen.Prefix + "Stack"
	gen.reduceType = "_" + gen.Prefix + "ReduceType"
	gen.actionType = "_" + gen.Prefix + "ActionType"
	gen.shiftAction = "_" + gen.Prefix + "ShiftAction"
	gen.reduceAction = "_" + gen.Prefix + "ReduceAction"
	gen.acceptAction = "_" + gen.Prefix + "AcceptAction"
	gen.action = "_" + gen.Prefix + "Action"
	gen.tableKey = "_" + gen.Prefix + "ActionTableKey"
	gen.actionTableType = "_" + gen.Prefix + "ActionTableType"
	gen.actionTable = "_" + gen.Prefix + "ActionTable"
	gen.parse = "_" + gen.Prefix + "Parse"

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
		Package:                   gen.Package,
		Imports:                   gen.goHeader,
		ActionType:                gen.action,
		ActionIdType:              gen.actionType,
		ShiftAction:               gen.shiftAction,
		ReduceAction:              gen.reduceAction,
		AcceptAction:              gen.acceptAction,
		StateIdType:               gen.stateId,
		ReduceType:                gen.reduceType,
		SymbolType:                gen.symbol,
		GenericSymbolType:         gen.genericSymbol,
		StackItemType:             gen.stackItem,
		StackType:                 gen.stack,
		SymbolStackType:           gen.symbolStack,
		SymbolIdType:              gen.symbolId,
		EndSymbolId:               gen.endSymbol,
		WildcardSymbolId:          gen.wildcardSymbol,
		LocationType:              gen.location,
		TokenType:                 gen.token,
		LexerType:                 gen.lexer,
		ReducerType:               gen.reducer,
		ErrHandlerType:            gen.errHandler,
		DefaultErrHandlerType:     gen.defaultErrHandler,
		ExpectedTerminalsFunc:     gen.expectedTerminals,
		ParseFuncPrefix:           gen.Prefix + "Parse",
		InternalParseFunc:         gen.parse,
		TableKeyType:              gen.tableKey,
		ActionTableType:           gen.actionTableType,
		ActionTable:               gen.actionTable,
		SortSlice:                 gen.Obj("sort.Slice"),
		Sprintf:                   gen.Obj("fmt.Sprintf"),
		Errorf:                    gen.Obj("fmt.Errorf"),
		EOF:                       gen.Obj("io.EOF"),
		OrderedSymbols:            orderedSymbols,
		Symbols:                   symbols,
		StartSymbols:              gen.Starts,
		OrderedStates:             gen.OrderedStates,
		OrderedValueTypes:         orderedValueTypes,
		OutputDebugNonKernelItems: gen.OutputDebugNonKernelItems,
	}

	return &GoCodeBuilder{file}, nil
}
