package code_gen

import (
	"fmt"
	"sort"
	"strings"

	lr "github.com/pattyshack/abc/src/lr/internal"
	"github.com/pattyshack/abc/src/lr/internal/parser"
)

type param struct {
	name      string
	paramType interface{}
}

type paramList []*param

func (list paramList) Len() int {
	return len(list)
}

func (list paramList) Less(i int, j int) bool {
	return list[i].name < list[j].name
}

func (list paramList) Swap(i int, j int) {
	list[i], list[j] = list[j], list[i]
}

func (list paramList) String() string {
	params := []string{}
	for _, param := range list {
		params = append(
			params,
			fmt.Sprintf("%s %v", param.name, param.paramType))
	}

	return strings.Join(params, ", ")
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

	*GoCodeBuilder

	nameLocs map[string]parser.LRLocation

	location string

	symbolId string
	stateId  string

	symbol string

	acceptSymbol   string
	startSymbol    string
	endSymbol      string
	wildcardSymbol string

	token   string
	lexer   string
	reducer string

	errHandler        string
	defaultErrHandler string
	expectedTerminals string

	pseudoToken string

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

	builder := NewGoCodeBuilder(cfg.Package)

	builder.HeaderBoilerplate.Line(
		"// Auto-generated from source: %s",
		grammar.Source)
	builder.HeaderBoilerplate.Line("")

	return &goCodeGen{
		Grammar:       grammar,
		GoSpec:        cfg,
		LRStates:      states,
		GoCodeBuilder: builder,
		nameLocs:      map[string]parser.LRLocation{},
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
	gen.symbol = "_" + gen.Prefix + "Symbol"
	gen.stateId = "_" + gen.Prefix + "StateId"
	gen.acceptSymbol = "_" + gen.Prefix + "AcceptMarker"
	gen.startSymbol = "_" + gen.Prefix + "StartMarker"
	gen.endSymbol = "_" + gen.Prefix + "EndMarker"
	gen.wildcardSymbol = "_" + gen.Prefix + "WildcardMarker"
	gen.token = gen.Prefix + "Symbol"
	gen.lexer = gen.Prefix + "Lexer"
	gen.reducer = gen.Prefix + "Reducer"
	gen.errHandler = gen.Prefix + "ParseErrorHandler"
	gen.defaultErrHandler = gen.Prefix + "DefaultParseErrorHandler"
	gen.expectedTerminals = "_" + gen.Prefix + "ExpectedTerminals"
	gen.pseudoToken = "_" + gen.Prefix + "PseudoToken"
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

	for _, term := range gen.Terms {
		valueType := gen.ValueTypes[term.ValueType.Value]
		if valueType == "" {
			return fmt.Errorf(
				"Undefined value type for <%s> %s",
				term.ValueType.Value,
				term.LRLocation)
		}

		term.CodeGenType = gen.Obj(valueType)

		symbolConst := gen.Prefix + snakeToCamel(term.Name) + "Symbol"

		if !term.IsTerminal {
			symbolConst = "_" + symbolConst
		}

		err := gen.check(symbolConst, term.LRLocation)
		if err != nil {
			return err
		}

		term.CodeGenSymbolConst = symbolConst

		for _, clause := range term.Clauses {
			reducerName := ""
			if clause.Label != nil {
				reducerName = snakeToCamel(clause.Label.Value)
			}
			reducerName += "To" + snakeToCamel(term.Name)

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
		}
	}

	for _, state := range gen.OrderedStates {
		state.CodeGenConst = fmt.Sprintf(
			"_%sState%d",
			gen.Prefix,
			state.StateNum)
	}

	return nil
}

func (gen *goCodeGen) generateTerminalSymbolIds() {
	l := gen.Line

	l("type %s string", gen.symbolId)
	l("")
	l("const (")
	gen.PushIndent()
	for _, term := range gen.Terminals {
		l("%s = %s(\"%s\")", term.CodeGenSymbolConst, gen.symbolId, term.Name)
	}
	gen.PopIndent()
	l(")")
	l("")
}

func (gen *goCodeGen) generateNonTerminalSymbolIds() {
	l := gen.Line

	l("const (")
	gen.PushIndent()

	l("%s = %s(\"%s\")", gen.acceptSymbol, gen.symbolId, lr.AcceptRule)
	l("%s = %s(\"%s\")", gen.startSymbol, gen.symbolId, lr.StartMarker)
	l("%s = %s(\"%s\")", gen.endSymbol, gen.symbolId, lr.EndMarker)
	l("%s = %s(\"%s\")", gen.wildcardSymbol, gen.symbolId, lr.Wildcard)

	l("")

	for _, term := range gen.NonTerminals {
		l("%s = %s(\"%s\")", term.CodeGenSymbolConst, gen.symbolId, term.Name)
	}
	gen.PopIndent()
	l(")")
	l("")
}

func (gen *goCodeGen) generateTokenInterface() {
	l := gen.Line
	push := gen.PushIndent
	pop := gen.PopIndent

	l("type %s struct {", gen.location)
	push()
	l("FileName string")
	l("Line int")
	l("Column int")
	pop()
	l("}")
	l("")

	l("func (l %s) String() string {", gen.location)
	push()
	l("return %v(\"%%v:%%v:%%v\", l.FileName, l.Line, l.Column)",
		gen.Obj("fmt.Sprintf"))
	pop()
	l("}")
	l("")

	l("func (l %s) ShortString() string {", gen.location)
	push()
	l("return %v(\"%%v:%%v\", l.Line, l.Column)", gen.Obj("fmt.Sprintf"))
	pop()
	l("}")
	l("")

	l("type %s interface {", gen.token)
	push()
	l("Id() %s", gen.symbolId)
	l("Location() %v", gen.location)
	pop()
	l("}")
	l("")
}

func (gen *goCodeGen) generateLexerInterface() {
	l := gen.Line

	l("type %s interface {", gen.lexer)
	gen.PushIndent()
	l("// Note: Return io.EOF to indicate end of stream")
	l("Next() (%s, error)", gen.token)
	gen.PopIndent()
	l("}")
	l("")
}

func (gen *goCodeGen) generateReducerInterface() {
	l := gen.Line

	l("type %s interface {", gen.reducer)
	gen.PushIndent()

	for ruleIdx, rule := range gen.NonTerminals {
		if ruleIdx > 0 {
			l("")
		}

		for clauseIdx, clause := range rule.Clauses {
			if clauseIdx > 0 {
				l("")
			}

			if clause.Label == nil {
				l("// %s: %s -> ...",
					clause.LRLocation.ShortString(),
					rule.Name)
			} else {
				l("// %s: %s -> %s: ...",
					clause.LRLocation.ShortString(),
					rule.Name,
					clause.Label.Value)
			}

			paramNameCount := map[string]int{}

			params := paramList{}
			for _, term := range clause.Bindings {
				paramName := term.Name

				// hack: append "_" to the end of the name ensures the
				// name is never a go keyword
				paramName = snakeToCamel(paramName) + "_"

				paramNameCount[paramName] += 1
				cnt := paramNameCount[paramName]
				if cnt > 1 {
					paramName = fmt.Sprintf("%s%d", paramName, cnt)
				}

				params = append(params, &param{paramName, term.CodeGenType})
			}

			l("%s(%v) (%v, error)",
				clause.CodeGenReducerName,
				params,
				rule.CodeGenType)
		}
	}

	gen.PopIndent()
	l("}")
	l("")
}

func (gen *goCodeGen) generateReduceTypes() {
	l := gen.Line

	l("type %s string", gen.reduceType)
	l("")
	l("const (")
	gen.PushIndent()
	for _, rule := range gen.NonTerminals {
		for _, clause := range rule.Clauses {
			l("%s = %s(\"%s\")",
				clause.CodeGenReducerNameConst,
				gen.reduceType,
				clause.CodeGenReducerName)
		}
	}
	gen.PopIndent()
	l(")")
}

func (gen *goCodeGen) generateActionTypes() {
	l := gen.Line

	l("type %s string", gen.actionType)
	l("")
	l("const (")
	gen.PushIndent()
	l("// NOTE: error action is implicit")
	l("%s = %s(\"shift\")", gen.shiftAction, gen.actionType)
	l("%s = %s(\"reduce\")", gen.reduceAction, gen.actionType)
	l("%s = %s(\"accept\")", gen.acceptAction, gen.actionType)
	gen.PopIndent()
	l(")")
	l("")
}

func (gen *goCodeGen) generateAction() {
	l := gen.Line
	push := gen.PushIndent
	pop := gen.PopIndent

	l("type %s struct {", gen.action)
	push()
	l("ActionType %s", gen.actionType)
	l("")
	l("ShiftStateId %s", gen.stateId)
	l("ReduceType %s", gen.reduceType)
	pop()
	l("}")
	l("")

	valueTerms := map[string][]*lr.Term{}
	for _, term := range gen.Terminals {
		valueTerms[term.ValueType.Value] = append(
			valueTerms[term.ValueType.Value],
			term)
	}

	values := []string{}
	for value, _ := range valueTerms {
		values = append(values, value)
	}

	sort.Strings(values)

	l("func (act *%s) ShiftItem(symbol %s) *%s {",
		gen.action,
		gen.token,
		gen.stackItem)
	push()
	l("item := &%s{StateId: act.ShiftStateId}", gen.stackItem)
	l("")
	l("reducedSymbol, ok := symbol.(*%s)", gen.symbol)
	l("if ok {")
	push()
	l("item.%s = reducedSymbol", gen.symbol)
	l("return item")
	pop()
	l("}")
	l("")

	l("item.%s = &%s{SymbolId_: symbol.Id()}", gen.symbol, gen.symbol)
	l("")
	l("switch symbol.Id() {")
	for _, value := range values {
		terms := valueTerms[value]
		consts := []string{}
		for _, term := range terms {
			consts = append(consts, term.CodeGenSymbolConst)
		}

		l("case %s:", strings.Join(consts, ", "))
		push()
		l("item.%s = symbol.(%s)", value, terms[0].CodeGenType)
		pop()
	}

	l("case %s: // EOF has no value", gen.endSymbol)
	l("default:")
	l("panic(\"Unexpected symbol type: \" + symbol.Id())")
	l("}")

	l("return item")
	pop()
	l("}")
	l("")

	l("func (act *%s) ReduceSymbol(reducer %s, stack %s) (%s, *%s, error) {",
		gen.action,
		gen.reducer,
		gen.stack,
		gen.stack,
		gen.symbol)
	push()
	l("var err error")
	l("symbol := &%s{}", gen.symbol)
	l("switch act.ReduceType {")
	for _, rule := range gen.NonTerminals {
		for _, clause := range rule.Clauses {
			l("case %s:", clause.CodeGenReducerNameConst)
			push()

			if len(clause.Bindings) > 0 {
				l("args := stack[len(stack)-%d:]", len(clause.Bindings))
				l("stack = stack[:len(stack)-%d]", len(clause.Bindings))
			}

			args := []string{}
			for i, term := range clause.Bindings {
				args = append(
					args,
					fmt.Sprintf("args[%d].%s", i, term.ValueType.Value))
			}

			l("symbol.SymbolId_ = %s", rule.CodeGenSymbolConst)
			l("symbol.%s, err = reducer.%s(%s)",
				rule.ValueType.Value,
				clause.CodeGenReducerName,
				strings.Join(args, ", "))
			pop()
		}
	}
	l("default:")
	push()
	l("panic(\"Unknown reduce type: \" + act.ReduceType)")
	pop()
	l("}")
	l("")
	l("if err != nil {")
	push()
	l("err = %s(\"Unexpected %%s reduce error: %%s\", act.ReduceType, err)",
		gen.Obj("fmt.Errorf"))
	pop()
	l("}")
	l("")
	l("return stack, symbol, err")
	pop()
	l("}")
	l("")
}

func (gen *goCodeGen) generateSymbolType() {
	l := gen.Line
	push := gen.PushIndent
	pop := gen.PopIndent

	l("type %s struct {", gen.symbol)
	push()
	l("SymbolId_ %s", gen.symbolId)
	l("")
	fields := paramList{}
	for name, valueType := range gen.ValueTypes {
		fields = append(fields, &param{name, gen.Obj(valueType)})
	}
	sort.Sort(fields)

	for _, field := range fields {
		l("%s %v", field.name, field.paramType)
	}

	pop()
	l("}")
	l("")

	valueTermConsts := map[string][]string{}
	for _, terms := range [][]*lr.Term{gen.Terminals, gen.NonTerminals} {
		for _, term := range terms {
			valueTermConsts[term.ValueType.Value] = append(
				valueTermConsts[term.ValueType.Value],
				term.CodeGenSymbolConst)
		}
	}

	l("func (s *%s) Id() %s {", gen.symbol, gen.symbolId)
	push()
	l("return s.SymbolId_")
	pop()
	l("}")
	l("")

	l("func (s *%s) Location() %s {", gen.symbol, gen.location)
	push()
	l("type locator interface { Location() %s }", gen.location)
	l("switch s.SymbolId_ {")
	for _, field := range fields {
		list := valueTermConsts[field.name]
		l("case %s:", strings.Join(list, ", "))
		push()
		l("loc, ok := interface{}(s.%s).(locator)", field.name)
		l("if ok {")
		push()
		l("return loc.Location()")
		pop()
		l("}")
		pop()
	}
	l("}")
	l("return %s{}", gen.location)
	pop()
	l("}")
}

func (gen *goCodeGen) generateStack() {
	l := gen.Line

	l("type %s struct {", gen.stackItem)
	gen.PushIndent()
	l("StateId %s", gen.stateId)
	l("")
	l("*%s", gen.symbol)
	gen.PopIndent()
	l("}")
	l("")

	l("type %s []*%s", gen.stack, gen.stackItem)
	l("")
}

func (gen *goCodeGen) generateDebugStates() {
	l := gen.Line

	symbols := []string{"$", "*"}
	for _, terms := range [][]*lr.Term{gen.Terminals, gen.NonTerminals} {
		for _, term := range terms {
			symbols = append(symbols, term.Name)
		}
	}

	gotoCount := 0
	reduceCount := 0
	shiftReduceCount := 0
	reduceReduceCount := 0

	l("/*")
	l("Parser Debug States:")

	for _, state := range gen.OrderedStates {
		reduce := map[string][]string{}
		l("  State %d:", state.StateNum)
		l("    Kernel Items:")
		firstNonKernel := true
		for _, item := range state.Items {
			if !item.IsKernel() && firstNonKernel {
				if !gen.OutputDebugNonKernelItems &&
					len(state.ShiftReduceConflictSymbols) == 0 &&
					len(state.ReduceReduceConflictSymbols) == 0 {
					break
				}

				firstNonKernel = false
				l("    Non-kernel Items:")
			}

			if item.IsReduce() {
				reduceCount += 1
				reduce[item.LookAhead] = append(
					reduce[item.LookAhead],
					item.Name)
			}
			l("      %s", item)
		}
		l("    Reduce:")
		if len(reduce) == 0 {
			l("      (nil)")
		}
		for _, symbol := range symbols {
			list := reduce[symbol]
			if len(list) > 0 {
				l("      %s -> %v", symbol, list)
			}
		}
		l("    Goto:")
		gotoCount += len(state.Goto)
		if len(state.Goto) == 0 {
			l("      (nil)")
		}
		for _, symbol := range symbols {
			child, ok := state.Goto[symbol]
			if ok {
				l("      %s -> State %d", symbol, child.StateNum)
			}
		}
		if len(state.ShiftReduceConflictSymbols) > 0 {
			l("    Shift/reduce conflict symbols:")
			l("      %v", state.ShiftReduceConflictSymbols)
			shiftReduceCount += len(state.ShiftReduceConflictSymbols)
		}
		if len(state.ReduceReduceConflictSymbols) > 0 {
			l("    Reduce/reduce conflict symbols:")
			l("      %v", state.ReduceReduceConflictSymbols)
			for _, symbol := range state.ReduceReduceConflictSymbols {
				reduceReduceCount += len(state.Reduce[symbol])
			}
		}

		l("")
	}

	l("Number of states: %d", len(gen.States))
	l("Number of shift actions: %d", gotoCount)
	l("Number of reduce actions: %d", reduceCount)
	l("Number of shift/reduce conflicts: %d", shiftReduceCount)
	l("Number of reduce/reduce conflicts: %d", reduceReduceCount)
	l("*/")
	l("")
}

func (gen *goCodeGen) generateStateIds() {
	l := gen.Line

	l("type %s string", gen.stateId)
	l("")
	l("const (")
	gen.PushIndent()
	for _, state := range gen.OrderedStates {
		l("%s = %s(\"State %d\")",
			state.CodeGenConst,
			gen.stateId,
			state.StateNum)
	}
	gen.PopIndent()
	l(")")
}

func (gen *goCodeGen) generateActionTable() {
	l := gen.Line
	push := gen.PushIndent
	pop := gen.PopIndent

	l("type %s struct {", gen.tableKey)
	push()
	l("%s", gen.stateId)
	l("%s", gen.symbolId)
	pop()
	l("}")
	l("")

	l("type %s map[%s]%s", gen.actionTableType, gen.tableKey, gen.action)
	l("")
	l("func (table %s) Get(stateId %s, symbol %s) (%s, bool) {",
		gen.actionTableType,
		gen.stateId,
		gen.symbolId,
		gen.action)
	push()
	l("action, ok := table[%s{stateId, symbol}]", gen.tableKey)
	l("if ok {")
	push()
	l("return action, ok")
	pop()
	l("}")
	l("action, ok = table[%s{stateId, %s}]", gen.tableKey, gen.wildcardSymbol)
	l("return action, ok")
	pop()
	l("}")
	l("")

	symbols := []string{"$", "*"}
	for _, terms := range [][]*lr.Term{gen.Terminals, gen.NonTerminals} {
		for _, term := range terms {
			symbols = append(symbols, term.Name)
		}
	}

	idToConst := map[string]string{
		lr.AcceptRule:  gen.acceptSymbol,
		lr.StartMarker: gen.startSymbol,
		lr.EndMarker:   gen.endSymbol,
		lr.Wildcard:    gen.wildcardSymbol,
	}

	for _, term := range gen.Terms {
		idToConst[term.Name] = term.CodeGenSymbolConst
	}

	l("var %s = %s{", gen.actionTable, gen.actionTableType)
	push()

	l("{%s, %s}: {%s, \"\", \"\"},",
		gen.OrderedStates[1].CodeGenConst,
		gen.endSymbol,
		gen.acceptAction)

	for _, state := range gen.OrderedStates {
		for _, symbol := range symbols {
			child, ok := state.Goto[symbol]
			if ok {
				l("{%s, %s}: {%s, %s, \"\"},",
					state.CodeGenConst,
					idToConst[symbol],
					gen.shiftAction,
					child.CodeGenConst)
			}
		}
	}

	for _, state := range gen.OrderedStates {
		for _, item := range state.Items {
			if !item.IsReduce() {
				continue
			}

			if state.StateNum == 1 && item.LookAhead == lr.EndMarker {
				continue
			}

			l("{%s, %s}: {%s, \"\", %s},",
				state.CodeGenConst,
				idToConst[item.LookAhead],
				gen.reduceAction,
				item.Clause.CodeGenReducerNameConst)
		}
	}
	pop()
	l("}")
	l("")
}

func (gen *goCodeGen) generatePseudoToken() {
	l := gen.Line

	l("type %s %s", gen.pseudoToken, gen.symbolId)
	l("")
	l("func (t %s) Id() %s { return %s(t) }",
		gen.pseudoToken,
		gen.symbolId,
		gen.symbolId)
	l("")
	l("func (%s) Location() %v { return %v{} }",
		gen.pseudoToken,
		gen.location,
		gen.location)
	l("")
}

func (gen *goCodeGen) generateSymbolStack() {
	l := gen.Line
	push := gen.PushIndent
	pop := gen.PopIndent

	l("type %s struct {", gen.symbolStack)
	push()
	l("lexer %s", gen.lexer)
	l("top []%s", gen.token)
	pop()
	l("}")
	l("")

	l("func (stack *%s) Top() (%s, error) {", gen.symbolStack, gen.token)
	push()
	l("if len(stack.top) == 0 {")
	push()
	l("token, err := stack.lexer.Next()")
	l("if err != nil {")
	push()
	l("if err != %v {", gen.Obj("io.EOF"))
	push()
	l("return nil, %v(\"Unexpected lex error: %%s\", err)",
		gen.Obj("fmt.Errorf"))
	pop()
	l("}")
	l("token = %s(%s)", gen.pseudoToken, gen.endSymbol)
	pop()
	l("}")
	l("stack.top = append(stack.top, token)")
	pop()
	l("}")
	l("return stack.top[len(stack.top)-1], nil")
	pop()
	l("}")
	l("")

	l("func (stack *%s) Push(symbol %s) {", gen.symbolStack, gen.token)
	push()
	l("stack.top = append(stack.top, symbol)")
	pop()
	l("}")
	l("")

	l("func (stack *%s) Pop() (%s, error) {", gen.symbolStack, gen.token)
	push()

	l("if len(stack.top) == 0 {")
	push()
	l("return nil, %v(\"internal error: cannot pop an empty top\")",
		gen.Obj("fmt.Errorf"))
	pop()
	l("}")

	l("ret := stack.top[len(stack.top)-1]")
	l("stack.top = stack.top[:len(stack.top)-1]")
	l("return ret, nil")

	pop()
	l("}")
	l("")
}

func (gen *goCodeGen) generateExpectedTerminals() {
	l := gen.Line

	idToConst := map[string]string{
		lr.EndMarker: gen.endSymbol,
	}
	symbols := []string{}

	for _, term := range gen.Terminals {
		idToConst[term.Name] = term.CodeGenSymbolConst
		symbols = append(symbols, term.Name)
	}

	l("var %s = map[%s][]%s{", gen.expectedTerminals, gen.stateId, gen.symbolId)
	gen.PushIndent()

	for _, state := range gen.OrderedStates {
		consts := []string{}

		for _, symbol := range symbols {
			_, ok := state.Goto[symbol]
			if !ok {
				continue
			}
			consts = append(consts, idToConst[symbol])
		}

		for _, item := range state.Items {
			if item.IsReduce() && item.LookAhead != lr.Wildcard {
				consts = append(consts, idToConst[item.LookAhead])
			}
		}

		if len(consts) > 0 {
			l("%s: []%s{%s},",
				state.CodeGenConst,
				gen.symbolId,
				strings.Join(consts, ", "))
		}
	}

	gen.PopIndent()
	l("}")
	l("")
}

func (gen *goCodeGen) generateParseErrorHandler() {
	l := gen.Line
	push := gen.PushIndent
	pop := gen.PopIndent

	l("type %s interface {", gen.errHandler)
	push()
	l("Error(nextToken %s, parseStack %s) error", gen.token, gen.stack)
	pop()
	l("}")
	l("")

	l("type %s struct {}", gen.defaultErrHandler)
	l("")
	l("func (%s) Error(nextToken %s, stack %s) error {",
		gen.defaultErrHandler,
		gen.token,
		gen.stack)
	push()
	l("return %v(\"Syntax error: unexpected symbol %%v. Expecting: %%v (%%v)\", nextToken.Id(), %s[stack[len(stack)-1].StateId], nextToken.Location())",
		gen.Obj("fmt.Errorf"),
		gen.expectedTerminals)
	pop()
	l("}")
	l("")
}

func (gen *goCodeGen) generateParse() {
	l := gen.Line
	push := gen.PushIndent
	pop := gen.PopIndent

	l("func %sParse(lexer %s, reducer %s) (%v, error) {",
		gen.Prefix,
		gen.lexer,
		gen.reducer,
		gen.Start.CodeGenType)
	push()
	l("return %sParseWithCustomErrorHandler(lexer, reducer, %s{})",
		gen.Prefix,
		gen.defaultErrHandler)
	pop()
	l("}")
	l("")

	l("func %sParseWithCustomErrorHandler(lexer %s, reducer %s, errHandler %s) (%v, error) {",
		gen.Prefix,
		gen.lexer,
		gen.reducer,
		gen.errHandler,
		gen.Start.CodeGenType)
	push()

	l("var errRetVal %v", gen.Start.CodeGenType)
	l("stateStack := %s{", gen.stack)
	push()
	l("&%s{%s, &%s{SymbolId_: %s}},",
		gen.stackItem,
		gen.OrderedStates[0].CodeGenConst,
		gen.symbol,
		gen.startSymbol)
	pop()
	l("}")

	l("symbolStack := &%s{lexer: lexer}", gen.symbolStack)
	l("")

	l("for {")
	push()

	l("nextSymbol, err := symbolStack.Top()")
	l("if err != nil {")
	push()

	l("return errRetVal, err")

	pop()
	l("}")

	l("")
	l("action, ok := %s.Get(stateStack[len(stateStack)-1].StateId, nextSymbol.Id())",
		gen.actionTable)
	l("if !ok {")
	push()
	l("return errRetVal, errHandler.Error(nextSymbol, stateStack)")
	pop()
	l("}")

	l("if action.ActionType == %s {", gen.shiftAction)
	push()
	l("stateStack = append(stateStack, action.ShiftItem(nextSymbol))")
	l("")
	l("_, err = symbolStack.Pop()")
	l("if err != nil {")
	push()
	l("return errRetVal, err")
	pop()
	l("}")

	pop()

	l("} else if action.ActionType == %s {", gen.reduceAction)

	push()
	l("var reduceSymbol %s", gen.token)
	l("stateStack, reduceSymbol, err = action.ReduceSymbol(reducer, stateStack)")
	l("if err != nil {")
	push()
	l("return errRetVal, err")
	pop()
	l("}")
	l("")

	l("symbolStack.Push(reduceSymbol)")

	pop()
	l("} else if action.ActionType == %s {", gen.acceptAction)
	push()

	l("if len(stateStack) != 2 {")
	push()
	l("panic(\"This should never happen\")")
	pop()
	l("}")

	l("return stateStack[1].%s, nil", gen.Start.ValueType.Value)
	l("")

	pop()
	l("} else {")
	push()
	l("panic(\"Unknown action type: \" + action.ActionType)")
	pop()
	l("}")
	pop()

	l("}")

	pop()
	l("}")
}

func GenerateGoLRCode(
	grammar *lr.Grammar,
	states *lr.LRStates) (
	*GoCodeBuilder,
	error) {

	gen, err := newGoCodeGen(grammar, states)
	if err != nil {
		return nil, err
	}

	err = gen.populateCodeGenVariables()
	if err != nil {
		return nil, err
	}

	gen.generateTerminalSymbolIds()
	gen.generateTokenInterface()
	gen.generateLexerInterface()
	gen.generateReducerInterface()

	gen.generateParseErrorHandler()

	gen.generateParse()

	l := gen.Line
	l("// =======================================================")
	l("// Parser internal implementation")
	l("// User should avoid directly accessing the following code")
	l("// =======================================================")
	l("")

	gen.generateNonTerminalSymbolIds()
	gen.generateActionTypes()
	gen.generateReduceTypes()

	gen.generateStateIds()

	gen.generateSymbolType()

	gen.generatePseudoToken()
	gen.generateSymbolStack()

	gen.generateStack()

	gen.generateAction()

	gen.generateActionTable()

	// Maybe make the action table map[StateId]map[Symbol]Action and
	// extract the expected terminal symbols from the action table
	gen.generateExpectedTerminals()

	gen.generateDebugStates()

	return gen.GoCodeBuilder, nil
}
