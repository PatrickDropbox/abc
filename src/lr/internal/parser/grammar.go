// Auto-generated from source: grammar.lr

package parser

import (
	fmt "fmt"
	io "io"
)

type LRSymbolId string

const (
	LRTokenSymbol          = LRSymbolId("TOKEN")
	LRTypeSymbol           = LRSymbolId("TYPE")
	LRStartSymbol          = LRSymbolId("START")
	LRRuleDefSymbol        = LRSymbolId("RULE_DEF")
	LRLabelSymbol          = LRSymbolId("LABEL")
	LRLtSymbol             = LRSymbolId("LT")
	LRGtSymbol             = LRSymbolId("GT")
	LROrSymbol             = LRSymbolId("OR")
	LRSemicolonSymbol      = LRSymbolId("SEMICOLON")
	LRIdentifierSymbol     = LRSymbolId("IDENTIFIER")
	LRSectionMarkerSymbol  = LRSymbolId("SECTION_MARKER")
	LRSectionContentSymbol = LRSymbolId("SECTION_CONTENT")
)

type LRLocation struct {
	FileName string
	Line     int
	Column   int
}

func (l LRLocation) String() string {
	return fmt.Sprintf("%v:%v:%v", l.FileName, l.Line, l.Column)
}

func (l LRLocation) ShortString() string {
	return fmt.Sprintf("%v:%v", l.Line, l.Column)
}

type LRSymbol interface {
	Id() LRSymbolId
	Location() LRLocation
}

type LRLexer interface {
	// Note: Return io.EOF to indicate end of stream
	Next() (LRSymbol, error)
}

type LRReducer interface {
	// 42:11: grammar -> ...
	ToGrammar(Defs_ []Definition, AdditionalSections_ []*AdditionalSection) (*Grammar, error)

	// 45:4: additional_sections -> add: ...
	AddToAdditionalSections(AdditionalSections_ []*AdditionalSection, AdditionalSection_ *AdditionalSection) ([]*AdditionalSection, error)

	// 46:4: additional_sections -> nil: ...
	NilToAdditionalSections() ([]*AdditionalSection, error)

	// 48:22: additional_section -> ...
	ToAdditionalSection(SectionMarker_ *Token, Identifier_ *Token, SectionContent_ *Token) (*AdditionalSection, error)

	// 51:4: defs -> add: ...
	AddToDefs(Defs_ []Definition, Def_ Definition) ([]Definition, error)

	// 52:4: defs -> add_explicit: ...
	AddExplicitToDefs(Defs_ []Definition, Def_ Definition, Semicolon_ *Token) ([]Definition, error)

	// 53:4: defs -> def: ...
	DefToDefs(Def_ Definition) ([]Definition, error)

	// 54:4: defs -> explicit_def: ...
	ExplicitDefToDefs(Def_ Definition, Semicolon_ *Token) ([]Definition, error)

	// 59:4: def -> term_decl: ...
	TermDeclToDef(Rword_ *Token, Lt_ *Token, Identifier_ *Token, Gt_ *Token, NonemptyIdentList_ []*Token) (Definition, error)

	// 61:4: def -> start_decl: ...
	StartDeclToDef(Start_ *Token, Identifier_ *Token) (Definition, error)

	// 62:4: def -> rule: ...
	RuleToDef(Rule_ *Rule) (Definition, error)

	// 65:4: rword -> token: ...
	TokenToRword(Token_ *Token) (*Token, error)

	// 66:4: rword -> type: ...
	TypeToRword(Type_ *Token) (*Token, error)

	// 69:4: nonempty_ident_list -> add: ...
	AddToNonemptyIdentList(NonemptyIdentList_ []*Token, Identifier_ *Token) ([]*Token, error)

	// 70:4: nonempty_ident_list -> ident: ...
	IdentToNonemptyIdentList(Identifier_ *Token) ([]*Token, error)

	// 73:4: ident_list -> non_empty_list: ...
	NonEmptyListToIdentList(NonemptyIdentList_ []*Token) ([]*Token, error)

	// 74:4: ident_list -> nil: ...
	NilToIdentList() ([]*Token, error)

	// 77:4: rule -> unlabeled_clause: ...
	UnlabeledClauseToRule(RuleDef_ *Token, IdentList_ []*Token) (*Rule, error)

	// 78:4: rule -> clauses: ...
	ClausesToRule(RuleDef_ *Token, LabeledClauses_ []*Clause) (*Rule, error)

	// 81:4: labeled_clauses -> add: ...
	AddToLabeledClauses(LabeledClauses_ []*Clause, Or_ *Token, LabeledClause_ *Clause) ([]*Clause, error)

	// 82:4: labeled_clauses -> clause: ...
	ClauseToLabeledClauses(LabeledClause_ *Clause) ([]*Clause, error)

	// 84:18: labeled_clause -> ...
	ToLabeledClause(Label_ *Token, IdentList_ []*Token) (*Clause, error)
}

type LRParseErrorHandler interface {
	Error(nextToken LRSymbol, parseStack _LRStack) error
}

type LRDefaultParseErrorHandler struct{}

func (LRDefaultParseErrorHandler) Error(nextToken LRSymbol, stack _LRStack) error {
	return fmt.Errorf("Syntax error: unexpected symbol %v. Expecting: %v (%v)", nextToken.Id(), _LRExpectedTerminals[stack[len(stack)-1].StateId], nextToken.Location())
}

func LRParse(lexer LRLexer, reducer LRReducer) (*Grammar, error) {
	return LRParseWithCustomErrorHandler(lexer, reducer, LRDefaultParseErrorHandler{})
}

func LRParseWithCustomErrorHandler(lexer LRLexer, reducer LRReducer, errHandler LRParseErrorHandler) (*Grammar, error) {
	var errRetVal *Grammar
	stateStack := _LRStack{
		&_LRStackItem{_LRState0, &_LRSymbol{SymbolId_: _LRStartMarker}},
	}
	symbolStack := &_LRPseudoSymbolStack{lexer: lexer}

	for {
		nextSymbol, err := symbolStack.Top()
		if err != nil {
			return errRetVal, err
		}

		action, ok := _LRActionTable.Get(stateStack[len(stateStack)-1].StateId, nextSymbol.Id())
		if !ok {
			return errRetVal, errHandler.Error(nextSymbol, stateStack)
		}
		if action.ActionType == _LRShiftAction {
			stateStack = append(stateStack, action.ShiftItem(nextSymbol))

			_, err = symbolStack.Pop()
			if err != nil {
				return errRetVal, err
			}
		} else if action.ActionType == _LRReduceAction {
			var reduceSymbol LRSymbol
			stateStack, reduceSymbol, err = action.ReduceSymbol(reducer, stateStack)
			if err != nil {
				return errRetVal, err
			}

			symbolStack.Push(reduceSymbol)
		} else if action.ActionType == _LRAcceptAction {
			if len(stateStack) != 2 {
				panic("This should never happen")
			}
			return stateStack[1].Grammar, nil

		} else {
			panic("Unknown action type: " + action.ActionType)
		}
	}
}

// =======================================================
// Parser internal implementation
// User should avoid directly accessing the following code
// =======================================================

const (
	_LRAcceptMarker   = LRSymbolId("#accept")
	_LRStartMarker    = LRSymbolId("^")
	_LREndMarker      = LRSymbolId("$")
	_LRWildcardMarker = LRSymbolId("*")

	_LRGrammarSymbol            = LRSymbolId("grammar")
	_LRAdditionalSectionsSymbol = LRSymbolId("additional_sections")
	_LRAdditionalSectionSymbol  = LRSymbolId("additional_section")
	_LRDefsSymbol               = LRSymbolId("defs")
	_LRDefSymbol                = LRSymbolId("def")
	_LRRwordSymbol              = LRSymbolId("rword")
	_LRNonemptyIdentListSymbol  = LRSymbolId("nonempty_ident_list")
	_LRIdentListSymbol          = LRSymbolId("ident_list")
	_LRRuleSymbol               = LRSymbolId("rule")
	_LRLabeledClausesSymbol     = LRSymbolId("labeled_clauses")
	_LRLabeledClauseSymbol      = LRSymbolId("labeled_clause")
)

type _LRActionType string

const (
	// NOTE: error action is implicit
	_LRShiftAction  = _LRActionType("shift")
	_LRReduceAction = _LRActionType("reduce")
	_LRAcceptAction = _LRActionType("accept")
)

type _LRReduceType string

const (
	_LRReduceToGrammar                = _LRReduceType("ToGrammar")
	_LRReduceAddToAdditionalSections  = _LRReduceType("AddToAdditionalSections")
	_LRReduceNilToAdditionalSections  = _LRReduceType("NilToAdditionalSections")
	_LRReduceToAdditionalSection      = _LRReduceType("ToAdditionalSection")
	_LRReduceAddToDefs                = _LRReduceType("AddToDefs")
	_LRReduceAddExplicitToDefs        = _LRReduceType("AddExplicitToDefs")
	_LRReduceDefToDefs                = _LRReduceType("DefToDefs")
	_LRReduceExplicitDefToDefs        = _LRReduceType("ExplicitDefToDefs")
	_LRReduceTermDeclToDef            = _LRReduceType("TermDeclToDef")
	_LRReduceStartDeclToDef           = _LRReduceType("StartDeclToDef")
	_LRReduceRuleToDef                = _LRReduceType("RuleToDef")
	_LRReduceTokenToRword             = _LRReduceType("TokenToRword")
	_LRReduceTypeToRword              = _LRReduceType("TypeToRword")
	_LRReduceAddToNonemptyIdentList   = _LRReduceType("AddToNonemptyIdentList")
	_LRReduceIdentToNonemptyIdentList = _LRReduceType("IdentToNonemptyIdentList")
	_LRReduceNonEmptyListToIdentList  = _LRReduceType("NonEmptyListToIdentList")
	_LRReduceNilToIdentList           = _LRReduceType("NilToIdentList")
	_LRReduceUnlabeledClauseToRule    = _LRReduceType("UnlabeledClauseToRule")
	_LRReduceClausesToRule            = _LRReduceType("ClausesToRule")
	_LRReduceAddToLabeledClauses      = _LRReduceType("AddToLabeledClauses")
	_LRReduceClauseToLabeledClauses   = _LRReduceType("ClauseToLabeledClauses")
	_LRReduceToLabeledClause          = _LRReduceType("ToLabeledClause")
)

type _LRStateId string

const (
	_LRState0  = _LRStateId("State 0")
	_LRState1  = _LRStateId("State 1")
	_LRState2  = _LRStateId("State 2")
	_LRState3  = _LRStateId("State 3")
	_LRState4  = _LRStateId("State 4")
	_LRState5  = _LRStateId("State 5")
	_LRState6  = _LRStateId("State 6")
	_LRState7  = _LRStateId("State 7")
	_LRState8  = _LRStateId("State 8")
	_LRState9  = _LRStateId("State 9")
	_LRState10 = _LRStateId("State 10")
	_LRState11 = _LRStateId("State 11")
	_LRState12 = _LRStateId("State 12")
	_LRState13 = _LRStateId("State 13")
	_LRState14 = _LRStateId("State 14")
	_LRState15 = _LRStateId("State 15")
	_LRState16 = _LRStateId("State 16")
	_LRState17 = _LRStateId("State 17")
	_LRState18 = _LRStateId("State 18")
	_LRState19 = _LRStateId("State 19")
	_LRState20 = _LRStateId("State 20")
	_LRState21 = _LRStateId("State 21")
	_LRState22 = _LRStateId("State 22")
	_LRState23 = _LRStateId("State 23")
	_LRState24 = _LRStateId("State 24")
	_LRState25 = _LRStateId("State 25")
	_LRState26 = _LRStateId("State 26")
	_LRState27 = _LRStateId("State 27")
	_LRState28 = _LRStateId("State 28")
	_LRState29 = _LRStateId("State 29")
	_LRState30 = _LRStateId("State 30")
	_LRState31 = _LRStateId("State 31")
	_LRState32 = _LRStateId("State 32")
)

type _LRSymbol struct {
	SymbolId_ LRSymbolId

	AdditionalSection  *AdditionalSection
	AdditionalSections []*AdditionalSection
	Clause             *Clause
	Clauses            []*Clause
	Definition         Definition
	Definitions        []Definition
	Grammar            *Grammar
	Rule               *Rule
	Token              *Token
	Tokens             []*Token
}

func (s *_LRSymbol) Id() LRSymbolId {
	return s.SymbolId_
}

func (s *_LRSymbol) Location() LRLocation {
	type locator interface{ Location() LRLocation }
	switch s.SymbolId_ {
	case _LRAdditionalSectionSymbol:
		loc, ok := interface{}(s.AdditionalSection).(locator)
		if ok {
			return loc.Location()
		}
	case _LRAdditionalSectionsSymbol:
		loc, ok := interface{}(s.AdditionalSections).(locator)
		if ok {
			return loc.Location()
		}
	case _LRLabeledClauseSymbol:
		loc, ok := interface{}(s.Clause).(locator)
		if ok {
			return loc.Location()
		}
	case _LRLabeledClausesSymbol:
		loc, ok := interface{}(s.Clauses).(locator)
		if ok {
			return loc.Location()
		}
	case _LRDefSymbol:
		loc, ok := interface{}(s.Definition).(locator)
		if ok {
			return loc.Location()
		}
	case _LRDefsSymbol:
		loc, ok := interface{}(s.Definitions).(locator)
		if ok {
			return loc.Location()
		}
	case _LRGrammarSymbol:
		loc, ok := interface{}(s.Grammar).(locator)
		if ok {
			return loc.Location()
		}
	case _LRRuleSymbol:
		loc, ok := interface{}(s.Rule).(locator)
		if ok {
			return loc.Location()
		}
	case LRTokenSymbol, LRTypeSymbol, LRStartSymbol, LRRuleDefSymbol, LRLabelSymbol, LRLtSymbol, LRGtSymbol, LROrSymbol, LRSemicolonSymbol, LRIdentifierSymbol, LRSectionMarkerSymbol, LRSectionContentSymbol, _LRRwordSymbol:
		loc, ok := interface{}(s.Token).(locator)
		if ok {
			return loc.Location()
		}
	case _LRNonemptyIdentListSymbol, _LRIdentListSymbol:
		loc, ok := interface{}(s.Tokens).(locator)
		if ok {
			return loc.Location()
		}
	}
	return LRLocation{}
}

type _LRPseudoToken LRSymbolId

func (t _LRPseudoToken) Id() LRSymbolId { return LRSymbolId(t) }

func (_LRPseudoToken) Location() LRLocation { return LRLocation{} }

type _LRPseudoSymbolStack struct {
	lexer LRLexer
	top   []LRSymbol
}

func (stack *_LRPseudoSymbolStack) Top() (LRSymbol, error) {
	if len(stack.top) == 0 {
		token, err := stack.lexer.Next()
		if err != nil {
			if err != io.EOF {
				return nil, fmt.Errorf("Unexpected lex error: %s", err)
			}
			token = _LRPseudoToken(_LREndMarker)
		}
		stack.top = append(stack.top, token)
	}
	return stack.top[len(stack.top)-1], nil
}

func (stack *_LRPseudoSymbolStack) Push(symbol LRSymbol) {
	stack.top = append(stack.top, symbol)
}

func (stack *_LRPseudoSymbolStack) Pop() (LRSymbol, error) {
	if len(stack.top) == 0 {
		return nil, fmt.Errorf("internal error: cannot pop an empty top")
	}
	ret := stack.top[len(stack.top)-1]
	stack.top = stack.top[:len(stack.top)-1]
	return ret, nil
}

type _LRStackItem struct {
	StateId _LRStateId

	*_LRSymbol
}

type _LRStack []*_LRStackItem

type _LRAction struct {
	ActionType _LRActionType

	ShiftStateId _LRStateId
	ReduceType   _LRReduceType
}

func (act *_LRAction) ShiftItem(symbol LRSymbol) *_LRStackItem {
	item := &_LRStackItem{StateId: act.ShiftStateId}

	reducedSymbol, ok := symbol.(*_LRSymbol)
	if ok {
		item._LRSymbol = reducedSymbol
		return item
	}

	item._LRSymbol = &_LRSymbol{SymbolId_: symbol.Id()}

	switch symbol.Id() {
	case LRTokenSymbol, LRTypeSymbol, LRStartSymbol, LRRuleDefSymbol, LRLabelSymbol, LRLtSymbol, LRGtSymbol, LROrSymbol, LRSemicolonSymbol, LRIdentifierSymbol, LRSectionMarkerSymbol, LRSectionContentSymbol:
		item.Token = symbol.(*Token)
	case _LREndMarker: // EOF has no value
	default:
		panic("Unexpected symbol type: " + symbol.Id())
	}
	return item
}

func (act *_LRAction) ReduceSymbol(reducer LRReducer, stack _LRStack) (_LRStack, *_LRSymbol, error) {
	var err error
	symbol := &_LRSymbol{}
	switch act.ReduceType {
	case _LRReduceToGrammar:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = _LRGrammarSymbol
		symbol.Grammar, err = reducer.ToGrammar(args[0].Definitions, args[1].AdditionalSections)
	case _LRReduceAddToAdditionalSections:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = _LRAdditionalSectionsSymbol
		symbol.AdditionalSections, err = reducer.AddToAdditionalSections(args[0].AdditionalSections, args[1].AdditionalSection)
	case _LRReduceNilToAdditionalSections:
		symbol.SymbolId_ = _LRAdditionalSectionsSymbol
		symbol.AdditionalSections, err = reducer.NilToAdditionalSections()
	case _LRReduceToAdditionalSection:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = _LRAdditionalSectionSymbol
		symbol.AdditionalSection, err = reducer.ToAdditionalSection(args[0].Token, args[1].Token, args[2].Token)
	case _LRReduceAddToDefs:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = _LRDefsSymbol
		symbol.Definitions, err = reducer.AddToDefs(args[0].Definitions, args[1].Definition)
	case _LRReduceAddExplicitToDefs:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = _LRDefsSymbol
		symbol.Definitions, err = reducer.AddExplicitToDefs(args[0].Definitions, args[1].Definition, args[2].Token)
	case _LRReduceDefToDefs:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _LRDefsSymbol
		symbol.Definitions, err = reducer.DefToDefs(args[0].Definition)
	case _LRReduceExplicitDefToDefs:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = _LRDefsSymbol
		symbol.Definitions, err = reducer.ExplicitDefToDefs(args[0].Definition, args[1].Token)
	case _LRReduceTermDeclToDef:
		args := stack[len(stack)-5:]
		stack = stack[:len(stack)-5]
		symbol.SymbolId_ = _LRDefSymbol
		symbol.Definition, err = reducer.TermDeclToDef(args[0].Token, args[1].Token, args[2].Token, args[3].Token, args[4].Tokens)
	case _LRReduceStartDeclToDef:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = _LRDefSymbol
		symbol.Definition, err = reducer.StartDeclToDef(args[0].Token, args[1].Token)
	case _LRReduceRuleToDef:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _LRDefSymbol
		symbol.Definition, err = reducer.RuleToDef(args[0].Rule)
	case _LRReduceTokenToRword:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _LRRwordSymbol
		symbol.Token, err = reducer.TokenToRword(args[0].Token)
	case _LRReduceTypeToRword:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _LRRwordSymbol
		symbol.Token, err = reducer.TypeToRword(args[0].Token)
	case _LRReduceAddToNonemptyIdentList:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = _LRNonemptyIdentListSymbol
		symbol.Tokens, err = reducer.AddToNonemptyIdentList(args[0].Tokens, args[1].Token)
	case _LRReduceIdentToNonemptyIdentList:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _LRNonemptyIdentListSymbol
		symbol.Tokens, err = reducer.IdentToNonemptyIdentList(args[0].Token)
	case _LRReduceNonEmptyListToIdentList:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _LRIdentListSymbol
		symbol.Tokens, err = reducer.NonEmptyListToIdentList(args[0].Tokens)
	case _LRReduceNilToIdentList:
		symbol.SymbolId_ = _LRIdentListSymbol
		symbol.Tokens, err = reducer.NilToIdentList()
	case _LRReduceUnlabeledClauseToRule:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = _LRRuleSymbol
		symbol.Rule, err = reducer.UnlabeledClauseToRule(args[0].Token, args[1].Tokens)
	case _LRReduceClausesToRule:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = _LRRuleSymbol
		symbol.Rule, err = reducer.ClausesToRule(args[0].Token, args[1].Clauses)
	case _LRReduceAddToLabeledClauses:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = _LRLabeledClausesSymbol
		symbol.Clauses, err = reducer.AddToLabeledClauses(args[0].Clauses, args[1].Token, args[2].Clause)
	case _LRReduceClauseToLabeledClauses:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _LRLabeledClausesSymbol
		symbol.Clauses, err = reducer.ClauseToLabeledClauses(args[0].Clause)
	case _LRReduceToLabeledClause:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = _LRLabeledClauseSymbol
		symbol.Clause, err = reducer.ToLabeledClause(args[0].Token, args[1].Tokens)
	default:
		panic("Unknown reduce type: " + act.ReduceType)
	}

	if err != nil {
		err = fmt.Errorf("Unexpected %s reduce error: %s", act.ReduceType, err)
	}

	return stack, symbol, err
}

type _LRActionTableKey struct {
	_LRStateId
	LRSymbolId
}

type _LRActionTableType map[_LRActionTableKey]_LRAction

func (table _LRActionTableType) Get(stateId _LRStateId, symbol LRSymbolId) (_LRAction, bool) {
	action, ok := table[_LRActionTableKey{stateId, symbol}]
	if ok {
		return action, ok
	}
	action, ok = table[_LRActionTableKey{stateId, _LRWildcardMarker}]
	return action, ok
}

var _LRActionTable = _LRActionTableType{
	{_LRState1, _LREndMarker}:                {_LRAcceptAction, "", ""},
	{_LRState0, LRTokenSymbol}:               {_LRShiftAction, _LRState4, ""},
	{_LRState0, LRTypeSymbol}:                {_LRShiftAction, _LRState5, ""},
	{_LRState0, LRStartSymbol}:               {_LRShiftAction, _LRState3, ""},
	{_LRState0, LRRuleDefSymbol}:             {_LRShiftAction, _LRState2, ""},
	{_LRState0, _LRGrammarSymbol}:            {_LRShiftAction, _LRState1, ""},
	{_LRState0, _LRDefsSymbol}:               {_LRShiftAction, _LRState7, ""},
	{_LRState0, _LRDefSymbol}:                {_LRShiftAction, _LRState6, ""},
	{_LRState0, _LRRwordSymbol}:              {_LRShiftAction, _LRState9, ""},
	{_LRState0, _LRRuleSymbol}:               {_LRShiftAction, _LRState8, ""},
	{_LRState2, LRLabelSymbol}:               {_LRShiftAction, _LRState11, ""},
	{_LRState2, LRIdentifierSymbol}:          {_LRShiftAction, _LRState10, ""},
	{_LRState2, _LRNonemptyIdentListSymbol}:  {_LRShiftAction, _LRState15, ""},
	{_LRState2, _LRIdentListSymbol}:          {_LRShiftAction, _LRState12, ""},
	{_LRState2, _LRLabeledClausesSymbol}:     {_LRShiftAction, _LRState14, ""},
	{_LRState2, _LRLabeledClauseSymbol}:      {_LRShiftAction, _LRState13, ""},
	{_LRState3, LRIdentifierSymbol}:          {_LRShiftAction, _LRState16, ""},
	{_LRState6, LRSemicolonSymbol}:           {_LRShiftAction, _LRState17, ""},
	{_LRState7, LRTokenSymbol}:               {_LRShiftAction, _LRState4, ""},
	{_LRState7, LRTypeSymbol}:                {_LRShiftAction, _LRState5, ""},
	{_LRState7, LRStartSymbol}:               {_LRShiftAction, _LRState3, ""},
	{_LRState7, LRRuleDefSymbol}:             {_LRShiftAction, _LRState2, ""},
	{_LRState7, _LRAdditionalSectionsSymbol}: {_LRShiftAction, _LRState18, ""},
	{_LRState7, _LRDefSymbol}:                {_LRShiftAction, _LRState19, ""},
	{_LRState7, _LRRwordSymbol}:              {_LRShiftAction, _LRState9, ""},
	{_LRState7, _LRRuleSymbol}:               {_LRShiftAction, _LRState8, ""},
	{_LRState9, LRLtSymbol}:                  {_LRShiftAction, _LRState20, ""},
	{_LRState11, LRIdentifierSymbol}:         {_LRShiftAction, _LRState10, ""},
	{_LRState11, _LRNonemptyIdentListSymbol}: {_LRShiftAction, _LRState15, ""},
	{_LRState11, _LRIdentListSymbol}:         {_LRShiftAction, _LRState21, ""},
	{_LRState14, LROrSymbol}:                 {_LRShiftAction, _LRState22, ""},
	{_LRState15, LRIdentifierSymbol}:         {_LRShiftAction, _LRState23, ""},
	{_LRState18, LRSectionMarkerSymbol}:      {_LRShiftAction, _LRState24, ""},
	{_LRState18, _LRAdditionalSectionSymbol}: {_LRShiftAction, _LRState25, ""},
	{_LRState19, LRSemicolonSymbol}:          {_LRShiftAction, _LRState26, ""},
	{_LRState20, LRIdentifierSymbol}:         {_LRShiftAction, _LRState27, ""},
	{_LRState22, LRLabelSymbol}:              {_LRShiftAction, _LRState11, ""},
	{_LRState22, _LRLabeledClauseSymbol}:     {_LRShiftAction, _LRState28, ""},
	{_LRState24, LRIdentifierSymbol}:         {_LRShiftAction, _LRState29, ""},
	{_LRState27, LRGtSymbol}:                 {_LRShiftAction, _LRState30, ""},
	{_LRState29, LRSectionContentSymbol}:     {_LRShiftAction, _LRState31, ""},
	{_LRState30, LRIdentifierSymbol}:         {_LRShiftAction, _LRState10, ""},
	{_LRState30, _LRNonemptyIdentListSymbol}: {_LRShiftAction, _LRState32, ""},
	{_LRState32, LRIdentifierSymbol}:         {_LRShiftAction, _LRState23, ""},
	{_LRState2, _LRWildcardMarker}:           {_LRReduceAction, "", _LRReduceNilToIdentList},
	{_LRState4, LRLtSymbol}:                  {_LRReduceAction, "", _LRReduceTokenToRword},
	{_LRState5, LRLtSymbol}:                  {_LRReduceAction, "", _LRReduceTypeToRword},
	{_LRState6, _LRWildcardMarker}:           {_LRReduceAction, "", _LRReduceDefToDefs},
	{_LRState7, _LRWildcardMarker}:           {_LRReduceAction, "", _LRReduceNilToAdditionalSections},
	{_LRState8, _LRWildcardMarker}:           {_LRReduceAction, "", _LRReduceRuleToDef},
	{_LRState10, _LRWildcardMarker}:          {_LRReduceAction, "", _LRReduceIdentToNonemptyIdentList},
	{_LRState11, _LRWildcardMarker}:          {_LRReduceAction, "", _LRReduceNilToIdentList},
	{_LRState12, _LRWildcardMarker}:          {_LRReduceAction, "", _LRReduceUnlabeledClauseToRule},
	{_LRState13, _LRWildcardMarker}:          {_LRReduceAction, "", _LRReduceClauseToLabeledClauses},
	{_LRState14, _LRWildcardMarker}:          {_LRReduceAction, "", _LRReduceClausesToRule},
	{_LRState15, _LRWildcardMarker}:          {_LRReduceAction, "", _LRReduceNonEmptyListToIdentList},
	{_LRState16, _LRWildcardMarker}:          {_LRReduceAction, "", _LRReduceStartDeclToDef},
	{_LRState17, _LRWildcardMarker}:          {_LRReduceAction, "", _LRReduceExplicitDefToDefs},
	{_LRState18, _LREndMarker}:               {_LRReduceAction, "", _LRReduceToGrammar},
	{_LRState19, _LRWildcardMarker}:          {_LRReduceAction, "", _LRReduceAddToDefs},
	{_LRState21, _LRWildcardMarker}:          {_LRReduceAction, "", _LRReduceToLabeledClause},
	{_LRState23, _LRWildcardMarker}:          {_LRReduceAction, "", _LRReduceAddToNonemptyIdentList},
	{_LRState25, _LRWildcardMarker}:          {_LRReduceAction, "", _LRReduceAddToAdditionalSections},
	{_LRState26, _LRWildcardMarker}:          {_LRReduceAction, "", _LRReduceAddExplicitToDefs},
	{_LRState28, _LRWildcardMarker}:          {_LRReduceAction, "", _LRReduceAddToLabeledClauses},
	{_LRState31, _LRWildcardMarker}:          {_LRReduceAction, "", _LRReduceToAdditionalSection},
	{_LRState32, _LRWildcardMarker}:          {_LRReduceAction, "", _LRReduceTermDeclToDef},
}

var _LRExpectedTerminals = map[_LRStateId][]LRSymbolId{
	_LRState0:  []LRSymbolId{LRTokenSymbol, LRTypeSymbol, LRStartSymbol, LRRuleDefSymbol},
	_LRState1:  []LRSymbolId{_LREndMarker},
	_LRState2:  []LRSymbolId{LRLabelSymbol, LRIdentifierSymbol},
	_LRState3:  []LRSymbolId{LRIdentifierSymbol},
	_LRState4:  []LRSymbolId{LRLtSymbol},
	_LRState5:  []LRSymbolId{LRLtSymbol},
	_LRState6:  []LRSymbolId{LRSemicolonSymbol},
	_LRState7:  []LRSymbolId{LRTokenSymbol, LRTypeSymbol, LRStartSymbol, LRRuleDefSymbol},
	_LRState9:  []LRSymbolId{LRLtSymbol},
	_LRState11: []LRSymbolId{LRIdentifierSymbol},
	_LRState14: []LRSymbolId{LROrSymbol},
	_LRState15: []LRSymbolId{LRIdentifierSymbol},
	_LRState18: []LRSymbolId{LRSectionMarkerSymbol, _LREndMarker},
	_LRState19: []LRSymbolId{LRSemicolonSymbol},
	_LRState20: []LRSymbolId{LRIdentifierSymbol},
	_LRState22: []LRSymbolId{LRLabelSymbol},
	_LRState24: []LRSymbolId{LRIdentifierSymbol},
	_LRState27: []LRSymbolId{LRGtSymbol},
	_LRState29: []LRSymbolId{LRSectionContentSymbol},
	_LRState30: []LRSymbolId{LRIdentifierSymbol},
	_LRState32: []LRSymbolId{LRIdentifierSymbol},
}

/*
Parser Debug States:
  State 0:
    Kernel Items:
      #accept:^.grammar
    Non-kernel Items:
      def:.START IDENTIFIER
      def:.rule
      def:.rword LT IDENTIFIER GT nonempty_ident_list
      defs:.def
      defs:.def SEMICOLON
      defs:.defs def
      defs:.defs def SEMICOLON
      grammar:.defs additional_sections
      rule:.RULE_DEF ident_list
      rule:.RULE_DEF labeled_clauses
      rword:.TOKEN
      rword:.TYPE
    Reduce:
      (nil)
    Goto:
      TOKEN -> State 4
      TYPE -> State 5
      START -> State 3
      RULE_DEF -> State 2
      grammar -> State 1
      defs -> State 7
      def -> State 6
      rword -> State 9
      rule -> State 8

  State 1:
    Kernel Items:
      #accept:^ grammar.,$
    Reduce:
      $ -> [#accept]
    Goto:
      (nil)

  State 2:
    Kernel Items:
      rule:RULE_DEF.ident_list
      rule:RULE_DEF.labeled_clauses
    Non-kernel Items:
      ident_list:.,*
      ident_list:.nonempty_ident_list
      labeled_clause:.LABEL ident_list
      labeled_clauses:.labeled_clause
      labeled_clauses:.labeled_clauses OR labeled_clause
      nonempty_ident_list:.IDENTIFIER
      nonempty_ident_list:.nonempty_ident_list IDENTIFIER
    Reduce:
      * -> [ident_list]
    Goto:
      LABEL -> State 11
      IDENTIFIER -> State 10
      nonempty_ident_list -> State 15
      ident_list -> State 12
      labeled_clauses -> State 14
      labeled_clause -> State 13

  State 3:
    Kernel Items:
      def:START.IDENTIFIER
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 16

  State 4:
    Kernel Items:
      rword:TOKEN.,LT
    Reduce:
      LT -> [rword]
    Goto:
      (nil)

  State 5:
    Kernel Items:
      rword:TYPE.,LT
    Reduce:
      LT -> [rword]
    Goto:
      (nil)

  State 6:
    Kernel Items:
      defs:def.,*
      defs:def.SEMICOLON
    Reduce:
      * -> [defs]
    Goto:
      SEMICOLON -> State 17

  State 7:
    Kernel Items:
      defs:defs.def
      defs:defs.def SEMICOLON
      grammar:defs.additional_sections
    Non-kernel Items:
      additional_sections:.,*
      additional_sections:.additional_sections additional_section
      def:.START IDENTIFIER
      def:.rule
      def:.rword LT IDENTIFIER GT nonempty_ident_list
      rule:.RULE_DEF ident_list
      rule:.RULE_DEF labeled_clauses
      rword:.TOKEN
      rword:.TYPE
    Reduce:
      * -> [additional_sections]
    Goto:
      TOKEN -> State 4
      TYPE -> State 5
      START -> State 3
      RULE_DEF -> State 2
      additional_sections -> State 18
      def -> State 19
      rword -> State 9
      rule -> State 8

  State 8:
    Kernel Items:
      def:rule.,*
    Reduce:
      * -> [def]
    Goto:
      (nil)

  State 9:
    Kernel Items:
      def:rword.LT IDENTIFIER GT nonempty_ident_list
    Reduce:
      (nil)
    Goto:
      LT -> State 20

  State 10:
    Kernel Items:
      nonempty_ident_list:IDENTIFIER.,*
    Reduce:
      * -> [nonempty_ident_list]
    Goto:
      (nil)

  State 11:
    Kernel Items:
      labeled_clause:LABEL.ident_list
    Non-kernel Items:
      ident_list:.,*
      ident_list:.nonempty_ident_list
      nonempty_ident_list:.IDENTIFIER
      nonempty_ident_list:.nonempty_ident_list IDENTIFIER
    Reduce:
      * -> [ident_list]
    Goto:
      IDENTIFIER -> State 10
      nonempty_ident_list -> State 15
      ident_list -> State 21

  State 12:
    Kernel Items:
      rule:RULE_DEF ident_list.,*
    Reduce:
      * -> [rule]
    Goto:
      (nil)

  State 13:
    Kernel Items:
      labeled_clauses:labeled_clause.,*
    Reduce:
      * -> [labeled_clauses]
    Goto:
      (nil)

  State 14:
    Kernel Items:
      labeled_clauses:labeled_clauses.OR labeled_clause
      rule:RULE_DEF labeled_clauses.,*
    Reduce:
      * -> [rule]
    Goto:
      OR -> State 22

  State 15:
    Kernel Items:
      ident_list:nonempty_ident_list.,*
      nonempty_ident_list:nonempty_ident_list.IDENTIFIER
    Reduce:
      * -> [ident_list]
    Goto:
      IDENTIFIER -> State 23

  State 16:
    Kernel Items:
      def:START IDENTIFIER.,*
    Reduce:
      * -> [def]
    Goto:
      (nil)

  State 17:
    Kernel Items:
      defs:def SEMICOLON.,*
    Reduce:
      * -> [defs]
    Goto:
      (nil)

  State 18:
    Kernel Items:
      additional_sections:additional_sections.additional_section
      grammar:defs additional_sections.,$
    Non-kernel Items:
      additional_section:.SECTION_MARKER IDENTIFIER SECTION_CONTENT
    Reduce:
      $ -> [grammar]
    Goto:
      SECTION_MARKER -> State 24
      additional_section -> State 25

  State 19:
    Kernel Items:
      defs:defs def.,*
      defs:defs def.SEMICOLON
    Reduce:
      * -> [defs]
    Goto:
      SEMICOLON -> State 26

  State 20:
    Kernel Items:
      def:rword LT.IDENTIFIER GT nonempty_ident_list
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 27

  State 21:
    Kernel Items:
      labeled_clause:LABEL ident_list.,*
    Reduce:
      * -> [labeled_clause]
    Goto:
      (nil)

  State 22:
    Kernel Items:
      labeled_clauses:labeled_clauses OR.labeled_clause
    Non-kernel Items:
      labeled_clause:.LABEL ident_list
    Reduce:
      (nil)
    Goto:
      LABEL -> State 11
      labeled_clause -> State 28

  State 23:
    Kernel Items:
      nonempty_ident_list:nonempty_ident_list IDENTIFIER.,*
    Reduce:
      * -> [nonempty_ident_list]
    Goto:
      (nil)

  State 24:
    Kernel Items:
      additional_section:SECTION_MARKER.IDENTIFIER SECTION_CONTENT
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 29

  State 25:
    Kernel Items:
      additional_sections:additional_sections additional_section.,*
    Reduce:
      * -> [additional_sections]
    Goto:
      (nil)

  State 26:
    Kernel Items:
      defs:defs def SEMICOLON.,*
    Reduce:
      * -> [defs]
    Goto:
      (nil)

  State 27:
    Kernel Items:
      def:rword LT IDENTIFIER.GT nonempty_ident_list
    Reduce:
      (nil)
    Goto:
      GT -> State 30

  State 28:
    Kernel Items:
      labeled_clauses:labeled_clauses OR labeled_clause.,*
    Reduce:
      * -> [labeled_clauses]
    Goto:
      (nil)

  State 29:
    Kernel Items:
      additional_section:SECTION_MARKER IDENTIFIER.SECTION_CONTENT
    Reduce:
      (nil)
    Goto:
      SECTION_CONTENT -> State 31

  State 30:
    Kernel Items:
      def:rword LT IDENTIFIER GT.nonempty_ident_list
    Non-kernel Items:
      nonempty_ident_list:.IDENTIFIER
      nonempty_ident_list:.nonempty_ident_list IDENTIFIER
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 10
      nonempty_ident_list -> State 32

  State 31:
    Kernel Items:
      additional_section:SECTION_MARKER IDENTIFIER SECTION_CONTENT.,*
    Reduce:
      * -> [additional_section]
    Goto:
      (nil)

  State 32:
    Kernel Items:
      def:rword LT IDENTIFIER GT nonempty_ident_list.,*
      nonempty_ident_list:nonempty_ident_list.IDENTIFIER
    Reduce:
      * -> [def]
    Goto:
      IDENTIFIER -> State 23

Number of states: 33
Number of shift actions: 43
Number of reduce actions: 24
Number of shift/reduce conflicts: 0
Number of reduce/reduce conflicts: 0
*/
