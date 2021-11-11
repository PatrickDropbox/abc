// Auto-generated from source: grammar.lr

package parser

import (
	fmt "fmt"
	io "io"
)

type LRSymbolId string

const (
	LRTokenToken          = LRSymbolId("TOKEN")
	LRTypeToken           = LRSymbolId("TYPE")
	LRStartToken          = LRSymbolId("START")
	LRRuleDefToken        = LRSymbolId("RULE_DEF")
	LRLabelToken          = LRSymbolId("LABEL")
	LRLtToken             = LRSymbolId("LT")
	LRGtToken             = LRSymbolId("GT")
	LROrToken             = LRSymbolId("OR")
	LRSemicolonToken      = LRSymbolId("SEMICOLON")
	LRSectionMarkerToken  = LRSymbolId("SECTION_MARKER")
	LRIdentifierToken     = LRSymbolId("IDENTIFIER")
	LRSectionContentToken = LRSymbolId("SECTION_CONTENT")
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

type LRToken interface {
	Id() LRSymbolId
	Loc() LRLocation
}

type LRGenericSymbol struct {
	LRSymbolId
	LRLocation
}

func (t *LRGenericSymbol) Id() LRSymbolId { return t.LRSymbolId }

func (t *LRGenericSymbol) Loc() LRLocation { return t.LRLocation }

type LRLexer interface {
	// Note: Return io.EOF to indicate end of stream
	// Token with unspecified value type should return *LRGenericSymbol
	Next() (LRToken, error)

	CurrentLocation() LRLocation
}

type LRReducer interface {
	// 42:11: grammar -> ...
	ToGrammar(Defs_ []Definition, AdditionalSections_ []*AdditionalSection) (*Grammar, error)

	// 45:4: additional_sections -> add: ...
	AddToAdditionalSections(AdditionalSections_ []*AdditionalSection, AdditionalSection_ *AdditionalSection) ([]*AdditionalSection, error)

	// 46:4: additional_sections -> nil: ...
	NilToAdditionalSections() ([]*AdditionalSection, error)

	// 48:22: additional_section -> ...
	ToAdditionalSection(SectionMarker_ *LRGenericSymbol, Identifier_ *Token, SectionContent_ *Token) (*AdditionalSection, error)

	// 51:4: defs -> add: ...
	AddToDefs(Defs_ []Definition, Def_ Definition) ([]Definition, error)

	// 52:4: defs -> add_explicit: ...
	AddExplicitToDefs(Defs_ []Definition, Def_ Definition, Semicolon_ *LRGenericSymbol) ([]Definition, error)

	// 53:4: defs -> def: ...
	DefToDefs(Def_ Definition) ([]Definition, error)

	// 54:4: defs -> explicit_def: ...
	ExplicitDefToDefs(Def_ Definition, Semicolon_ *LRGenericSymbol) ([]Definition, error)

	// 59:4: def -> term_decl: ...
	TermDeclToDef(Rword_ *LRGenericSymbol, Lt_ *LRGenericSymbol, Identifier_ *Token, Gt_ *LRGenericSymbol, NonemptyIdentList_ []*Token) (Definition, error)

	// 60:4: def -> untyped_term_decl: ...
	UntypedTermDeclToDef(Rword_ *LRGenericSymbol, NonemptyIdentList_ []*Token) (Definition, error)

	// 62:4: def -> start_decl: ...
	StartDeclToDef(Start_ *LRGenericSymbol, Identifier_ *Token) (Definition, error)

	// 63:4: def -> rule: ...
	RuleToDef(Rule_ *Rule) (Definition, error)

	// 66:4: rword -> token: ...
	TokenToRword(Token_ *LRGenericSymbol) (*LRGenericSymbol, error)

	// 67:4: rword -> type: ...
	TypeToRword(Type_ *LRGenericSymbol) (*LRGenericSymbol, error)

	// 70:4: nonempty_ident_list -> add: ...
	AddToNonemptyIdentList(NonemptyIdentList_ []*Token, Identifier_ *Token) ([]*Token, error)

	// 71:4: nonempty_ident_list -> ident: ...
	IdentToNonemptyIdentList(Identifier_ *Token) ([]*Token, error)

	// 74:4: ident_list -> non_empty_list: ...
	NonEmptyListToIdentList(NonemptyIdentList_ []*Token) ([]*Token, error)

	// 75:4: ident_list -> nil: ...
	NilToIdentList() ([]*Token, error)

	// 78:4: rule -> unlabeled_clause: ...
	UnlabeledClauseToRule(RuleDef_ *Token, IdentList_ []*Token) (*Rule, error)

	// 79:4: rule -> clauses: ...
	ClausesToRule(RuleDef_ *Token, LabeledClauses_ []*Clause) (*Rule, error)

	// 82:4: labeled_clauses -> add: ...
	AddToLabeledClauses(LabeledClauses_ []*Clause, Or_ *LRGenericSymbol, LabeledClause_ *Clause) ([]*Clause, error)

	// 83:4: labeled_clauses -> clause: ...
	ClauseToLabeledClauses(LabeledClause_ *Clause) ([]*Clause, error)

	// 85:18: labeled_clause -> ...
	ToLabeledClause(Label_ *Token, IdentList_ []*Token) (*Clause, error)
}

type LRParseErrorHandler interface {
	Error(nextToken LRToken, parseStack _LRStack) error
}

type LRDefaultParseErrorHandler struct{}

func (LRDefaultParseErrorHandler) Error(nextToken LRToken, stack _LRStack) error {
	return fmt.Errorf("Syntax error: unexpected symbol %v. Expecting: %v (%v)", nextToken.Id(), _LRExpectedTerminals[stack[len(stack)-1].StateId], nextToken.Loc())
}

func LRParse(lexer LRLexer, reducer LRReducer) (*Grammar, error) {
	return LRParseWithCustomErrorHandler(lexer, reducer, LRDefaultParseErrorHandler{})
}

func LRParseWithCustomErrorHandler(lexer LRLexer, reducer LRReducer, errHandler LRParseErrorHandler) (*Grammar, error) {
	var errRetVal *Grammar
	stateStack := _LRStack{
		// Note: we don't have to populate the start symbol since its value is never accessed
		&_LRStackItem{_LRState0, nil},
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
			var reduceSymbol *LRSymbol
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
	_LREndMarker      = LRSymbolId("$")
	_LRWildcardMarker = LRSymbolId("*")

	LRGrammarType            = LRSymbolId("grammar")
	LRAdditionalSectionsType = LRSymbolId("additional_sections")
	LRAdditionalSectionType  = LRSymbolId("additional_section")
	LRDefsType               = LRSymbolId("defs")
	LRDefType                = LRSymbolId("def")
	LRRwordType              = LRSymbolId("rword")
	LRNonemptyIdentListType  = LRSymbolId("nonempty_ident_list")
	LRIdentListType          = LRSymbolId("ident_list")
	LRRuleType               = LRSymbolId("rule")
	LRLabeledClausesType     = LRSymbolId("labeled_clauses")
	LRLabeledClauseType      = LRSymbolId("labeled_clause")
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
	_LRReduceUntypedTermDeclToDef     = _LRReduceType("UntypedTermDeclToDef")
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

type _LRStateId int

func (id _LRStateId) String() string {
	return fmt.Sprintf("State %d", int(id))
}

const (
	_LRState0  = _LRStateId(0)
	_LRState1  = _LRStateId(1)
	_LRState2  = _LRStateId(2)
	_LRState3  = _LRStateId(3)
	_LRState4  = _LRStateId(4)
	_LRState5  = _LRStateId(5)
	_LRState6  = _LRStateId(6)
	_LRState7  = _LRStateId(7)
	_LRState8  = _LRStateId(8)
	_LRState9  = _LRStateId(9)
	_LRState10 = _LRStateId(10)
	_LRState11 = _LRStateId(11)
	_LRState12 = _LRStateId(12)
	_LRState13 = _LRStateId(13)
	_LRState14 = _LRStateId(14)
	_LRState15 = _LRStateId(15)
	_LRState16 = _LRStateId(16)
	_LRState17 = _LRStateId(17)
	_LRState18 = _LRStateId(18)
	_LRState19 = _LRStateId(19)
	_LRState20 = _LRStateId(20)
	_LRState21 = _LRStateId(21)
	_LRState22 = _LRStateId(22)
	_LRState23 = _LRStateId(23)
	_LRState24 = _LRStateId(24)
	_LRState25 = _LRStateId(25)
	_LRState26 = _LRStateId(26)
	_LRState27 = _LRStateId(27)
	_LRState28 = _LRStateId(28)
	_LRState29 = _LRStateId(29)
	_LRState30 = _LRStateId(30)
	_LRState31 = _LRStateId(31)
	_LRState32 = _LRStateId(32)
	_LRState33 = _LRStateId(33)
)

type LRSymbol struct {
	SymbolId_ LRSymbolId

	Generic_ *LRGenericSymbol

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

func NewSymbol(token LRToken) (*LRSymbol, error) {
	symbol, ok := token.(*LRSymbol)
	if ok {
		return symbol, nil
	}

	symbol = &LRSymbol{SymbolId_: token.Id()}
	switch token.Id() {
	case _LREndMarker, LRTokenToken, LRTypeToken, LRStartToken, LRLtToken, LRGtToken, LROrToken, LRSemicolonToken, LRSectionMarkerToken:
		val, ok := token.(*LRGenericSymbol)
		if !ok {
			return nil, fmt.Errorf("Invalid value type for token %s.  Expecting *LRGenericSymbol (%v)", token.Id(), token.Loc())
		}
		symbol.Generic_ = val
	case LRRuleDefToken, LRLabelToken, LRIdentifierToken, LRSectionContentToken:
		val, ok := token.(*Token)
		if !ok {
			return nil, fmt.Errorf("Invalid value type for token %s.  Expecting *Token (%v)", token.Id(), token.Loc())
		}
		symbol.Token = val
	default:
		return nil, fmt.Errorf("Unexpected token type: %s", symbol.Id())
	}
	return symbol, nil
}

func (s *LRSymbol) Id() LRSymbolId {
	return s.SymbolId_
}

func (s *LRSymbol) Loc() LRLocation {
	type locator interface{ Loc() LRLocation }
	switch s.SymbolId_ {
	case LRAdditionalSectionType:
		loc, ok := interface{}(s.AdditionalSection).(locator)
		if ok {
			return loc.Loc()
		}
	case LRAdditionalSectionsType:
		loc, ok := interface{}(s.AdditionalSections).(locator)
		if ok {
			return loc.Loc()
		}
	case LRLabeledClauseType:
		loc, ok := interface{}(s.Clause).(locator)
		if ok {
			return loc.Loc()
		}
	case LRLabeledClausesType:
		loc, ok := interface{}(s.Clauses).(locator)
		if ok {
			return loc.Loc()
		}
	case LRDefType:
		loc, ok := interface{}(s.Definition).(locator)
		if ok {
			return loc.Loc()
		}
	case LRDefsType:
		loc, ok := interface{}(s.Definitions).(locator)
		if ok {
			return loc.Loc()
		}
	case LRGrammarType:
		loc, ok := interface{}(s.Grammar).(locator)
		if ok {
			return loc.Loc()
		}
	case LRRuleType:
		loc, ok := interface{}(s.Rule).(locator)
		if ok {
			return loc.Loc()
		}
	case LRRuleDefToken, LRLabelToken, LRIdentifierToken, LRSectionContentToken:
		loc, ok := interface{}(s.Token).(locator)
		if ok {
			return loc.Loc()
		}
	case LRNonemptyIdentListType, LRIdentListType:
		loc, ok := interface{}(s.Tokens).(locator)
		if ok {
			return loc.Loc()
		}
	}
	if s.Generic_ != nil {
		return s.Generic_.Loc()
	}
	return LRLocation{}
}

type _LRPseudoSymbolStack struct {
	lexer LRLexer
	top   []*LRSymbol
}

func (stack *_LRPseudoSymbolStack) Top() (*LRSymbol, error) {
	if len(stack.top) == 0 {
		token, err := stack.lexer.Next()
		if err != nil {
			if err != io.EOF {
				return nil, fmt.Errorf("Unexpected lex error: %s", err)
			}
			token = &LRGenericSymbol{_LREndMarker, stack.lexer.CurrentLocation()}
		}
		item, err := NewSymbol(token)
		if err != nil {
			return nil, err
		}
		stack.top = append(stack.top, item)
	}
	return stack.top[len(stack.top)-1], nil
}

func (stack *_LRPseudoSymbolStack) Push(symbol *LRSymbol) {
	stack.top = append(stack.top, symbol)
}

func (stack *_LRPseudoSymbolStack) Pop() (LRToken, error) {
	if len(stack.top) == 0 {
		return nil, fmt.Errorf("internal error: cannot pop an empty top")
	}
	ret := stack.top[len(stack.top)-1]
	stack.top = stack.top[:len(stack.top)-1]
	return ret, nil
}

type _LRStackItem struct {
	StateId _LRStateId

	*LRSymbol
}

type _LRStack []*_LRStackItem

type _LRAction struct {
	ActionType _LRActionType

	ShiftStateId _LRStateId
	ReduceType   _LRReduceType
}

func (act *_LRAction) ShiftItem(symbol *LRSymbol) *_LRStackItem {
	return &_LRStackItem{StateId: act.ShiftStateId, LRSymbol: symbol}
}

func (act *_LRAction) ReduceSymbol(reducer LRReducer, stack _LRStack) (_LRStack, *LRSymbol, error) {
	var err error
	symbol := &LRSymbol{}
	switch act.ReduceType {
	case _LRReduceToGrammar:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = LRGrammarType
		symbol.Grammar, err = reducer.ToGrammar(args[0].Definitions, args[1].AdditionalSections)
	case _LRReduceAddToAdditionalSections:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = LRAdditionalSectionsType
		symbol.AdditionalSections, err = reducer.AddToAdditionalSections(args[0].AdditionalSections, args[1].AdditionalSection)
	case _LRReduceNilToAdditionalSections:
		symbol.SymbolId_ = LRAdditionalSectionsType
		symbol.AdditionalSections, err = reducer.NilToAdditionalSections()
	case _LRReduceToAdditionalSection:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = LRAdditionalSectionType
		symbol.AdditionalSection, err = reducer.ToAdditionalSection(args[0].Generic_, args[1].Token, args[2].Token)
	case _LRReduceAddToDefs:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = LRDefsType
		symbol.Definitions, err = reducer.AddToDefs(args[0].Definitions, args[1].Definition)
	case _LRReduceAddExplicitToDefs:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = LRDefsType
		symbol.Definitions, err = reducer.AddExplicitToDefs(args[0].Definitions, args[1].Definition, args[2].Generic_)
	case _LRReduceDefToDefs:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = LRDefsType
		symbol.Definitions, err = reducer.DefToDefs(args[0].Definition)
	case _LRReduceExplicitDefToDefs:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = LRDefsType
		symbol.Definitions, err = reducer.ExplicitDefToDefs(args[0].Definition, args[1].Generic_)
	case _LRReduceTermDeclToDef:
		args := stack[len(stack)-5:]
		stack = stack[:len(stack)-5]
		symbol.SymbolId_ = LRDefType
		symbol.Definition, err = reducer.TermDeclToDef(args[0].Generic_, args[1].Generic_, args[2].Token, args[3].Generic_, args[4].Tokens)
	case _LRReduceUntypedTermDeclToDef:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = LRDefType
		symbol.Definition, err = reducer.UntypedTermDeclToDef(args[0].Generic_, args[1].Tokens)
	case _LRReduceStartDeclToDef:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = LRDefType
		symbol.Definition, err = reducer.StartDeclToDef(args[0].Generic_, args[1].Token)
	case _LRReduceRuleToDef:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = LRDefType
		symbol.Definition, err = reducer.RuleToDef(args[0].Rule)
	case _LRReduceTokenToRword:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = LRRwordType
		symbol.Generic_, err = reducer.TokenToRword(args[0].Generic_)
	case _LRReduceTypeToRword:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = LRRwordType
		symbol.Generic_, err = reducer.TypeToRword(args[0].Generic_)
	case _LRReduceAddToNonemptyIdentList:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = LRNonemptyIdentListType
		symbol.Tokens, err = reducer.AddToNonemptyIdentList(args[0].Tokens, args[1].Token)
	case _LRReduceIdentToNonemptyIdentList:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = LRNonemptyIdentListType
		symbol.Tokens, err = reducer.IdentToNonemptyIdentList(args[0].Token)
	case _LRReduceNonEmptyListToIdentList:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = LRIdentListType
		symbol.Tokens, err = reducer.NonEmptyListToIdentList(args[0].Tokens)
	case _LRReduceNilToIdentList:
		symbol.SymbolId_ = LRIdentListType
		symbol.Tokens, err = reducer.NilToIdentList()
	case _LRReduceUnlabeledClauseToRule:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = LRRuleType
		symbol.Rule, err = reducer.UnlabeledClauseToRule(args[0].Token, args[1].Tokens)
	case _LRReduceClausesToRule:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = LRRuleType
		symbol.Rule, err = reducer.ClausesToRule(args[0].Token, args[1].Clauses)
	case _LRReduceAddToLabeledClauses:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = LRLabeledClausesType
		symbol.Clauses, err = reducer.AddToLabeledClauses(args[0].Clauses, args[1].Generic_, args[2].Clause)
	case _LRReduceClauseToLabeledClauses:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = LRLabeledClausesType
		symbol.Clauses, err = reducer.ClauseToLabeledClauses(args[0].Clause)
	case _LRReduceToLabeledClause:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = LRLabeledClauseType
		symbol.Clause, err = reducer.ToLabeledClause(args[0].Token, args[1].Tokens)
	default:
		panic("Unknown reduce type: " + act.ReduceType)
	}

	if err != nil {
		err = fmt.Errorf("Unexpected %s reduce error: %s", act.ReduceType, err)
	}

	return stack, symbol, err
}

var (
	_LRGotoState0Action                     = &_LRAction{_LRShiftAction, _LRState0, ""}
	_LRGotoState1Action                     = &_LRAction{_LRShiftAction, _LRState1, ""}
	_LRGotoState2Action                     = &_LRAction{_LRShiftAction, _LRState2, ""}
	_LRGotoState3Action                     = &_LRAction{_LRShiftAction, _LRState3, ""}
	_LRGotoState4Action                     = &_LRAction{_LRShiftAction, _LRState4, ""}
	_LRGotoState5Action                     = &_LRAction{_LRShiftAction, _LRState5, ""}
	_LRGotoState6Action                     = &_LRAction{_LRShiftAction, _LRState6, ""}
	_LRGotoState7Action                     = &_LRAction{_LRShiftAction, _LRState7, ""}
	_LRGotoState8Action                     = &_LRAction{_LRShiftAction, _LRState8, ""}
	_LRGotoState9Action                     = &_LRAction{_LRShiftAction, _LRState9, ""}
	_LRGotoState10Action                    = &_LRAction{_LRShiftAction, _LRState10, ""}
	_LRGotoState11Action                    = &_LRAction{_LRShiftAction, _LRState11, ""}
	_LRGotoState12Action                    = &_LRAction{_LRShiftAction, _LRState12, ""}
	_LRGotoState13Action                    = &_LRAction{_LRShiftAction, _LRState13, ""}
	_LRGotoState14Action                    = &_LRAction{_LRShiftAction, _LRState14, ""}
	_LRGotoState15Action                    = &_LRAction{_LRShiftAction, _LRState15, ""}
	_LRGotoState16Action                    = &_LRAction{_LRShiftAction, _LRState16, ""}
	_LRGotoState17Action                    = &_LRAction{_LRShiftAction, _LRState17, ""}
	_LRGotoState18Action                    = &_LRAction{_LRShiftAction, _LRState18, ""}
	_LRGotoState19Action                    = &_LRAction{_LRShiftAction, _LRState19, ""}
	_LRGotoState20Action                    = &_LRAction{_LRShiftAction, _LRState20, ""}
	_LRGotoState21Action                    = &_LRAction{_LRShiftAction, _LRState21, ""}
	_LRGotoState22Action                    = &_LRAction{_LRShiftAction, _LRState22, ""}
	_LRGotoState23Action                    = &_LRAction{_LRShiftAction, _LRState23, ""}
	_LRGotoState24Action                    = &_LRAction{_LRShiftAction, _LRState24, ""}
	_LRGotoState25Action                    = &_LRAction{_LRShiftAction, _LRState25, ""}
	_LRGotoState26Action                    = &_LRAction{_LRShiftAction, _LRState26, ""}
	_LRGotoState27Action                    = &_LRAction{_LRShiftAction, _LRState27, ""}
	_LRGotoState28Action                    = &_LRAction{_LRShiftAction, _LRState28, ""}
	_LRGotoState29Action                    = &_LRAction{_LRShiftAction, _LRState29, ""}
	_LRGotoState30Action                    = &_LRAction{_LRShiftAction, _LRState30, ""}
	_LRGotoState31Action                    = &_LRAction{_LRShiftAction, _LRState31, ""}
	_LRGotoState32Action                    = &_LRAction{_LRShiftAction, _LRState32, ""}
	_LRGotoState33Action                    = &_LRAction{_LRShiftAction, _LRState33, ""}
	_LRReduceToGrammarAction                = &_LRAction{_LRReduceAction, 0, _LRReduceToGrammar}
	_LRReduceAddToAdditionalSectionsAction  = &_LRAction{_LRReduceAction, 0, _LRReduceAddToAdditionalSections}
	_LRReduceNilToAdditionalSectionsAction  = &_LRAction{_LRReduceAction, 0, _LRReduceNilToAdditionalSections}
	_LRReduceToAdditionalSectionAction      = &_LRAction{_LRReduceAction, 0, _LRReduceToAdditionalSection}
	_LRReduceAddToDefsAction                = &_LRAction{_LRReduceAction, 0, _LRReduceAddToDefs}
	_LRReduceAddExplicitToDefsAction        = &_LRAction{_LRReduceAction, 0, _LRReduceAddExplicitToDefs}
	_LRReduceDefToDefsAction                = &_LRAction{_LRReduceAction, 0, _LRReduceDefToDefs}
	_LRReduceExplicitDefToDefsAction        = &_LRAction{_LRReduceAction, 0, _LRReduceExplicitDefToDefs}
	_LRReduceTermDeclToDefAction            = &_LRAction{_LRReduceAction, 0, _LRReduceTermDeclToDef}
	_LRReduceUntypedTermDeclToDefAction     = &_LRAction{_LRReduceAction, 0, _LRReduceUntypedTermDeclToDef}
	_LRReduceStartDeclToDefAction           = &_LRAction{_LRReduceAction, 0, _LRReduceStartDeclToDef}
	_LRReduceRuleToDefAction                = &_LRAction{_LRReduceAction, 0, _LRReduceRuleToDef}
	_LRReduceTokenToRwordAction             = &_LRAction{_LRReduceAction, 0, _LRReduceTokenToRword}
	_LRReduceTypeToRwordAction              = &_LRAction{_LRReduceAction, 0, _LRReduceTypeToRword}
	_LRReduceAddToNonemptyIdentListAction   = &_LRAction{_LRReduceAction, 0, _LRReduceAddToNonemptyIdentList}
	_LRReduceIdentToNonemptyIdentListAction = &_LRAction{_LRReduceAction, 0, _LRReduceIdentToNonemptyIdentList}
	_LRReduceNonEmptyListToIdentListAction  = &_LRAction{_LRReduceAction, 0, _LRReduceNonEmptyListToIdentList}
	_LRReduceNilToIdentListAction           = &_LRAction{_LRReduceAction, 0, _LRReduceNilToIdentList}
	_LRReduceUnlabeledClauseToRuleAction    = &_LRAction{_LRReduceAction, 0, _LRReduceUnlabeledClauseToRule}
	_LRReduceClausesToRuleAction            = &_LRAction{_LRReduceAction, 0, _LRReduceClausesToRule}
	_LRReduceAddToLabeledClausesAction      = &_LRAction{_LRReduceAction, 0, _LRReduceAddToLabeledClauses}
	_LRReduceClauseToLabeledClausesAction   = &_LRAction{_LRReduceAction, 0, _LRReduceClauseToLabeledClauses}
	_LRReduceToLabeledClauseAction          = &_LRAction{_LRReduceAction, 0, _LRReduceToLabeledClause}
)

type _LRActionTableKey struct {
	_LRStateId
	LRSymbolId
}

type _LRActionTableType map[_LRActionTableKey]*_LRAction

func (table _LRActionTableType) Get(stateId _LRStateId, symbol LRSymbolId) (*_LRAction, bool) {
	action, ok := table[_LRActionTableKey{stateId, symbol}]
	if ok {
		return action, ok
	}
	action, ok = table[_LRActionTableKey{stateId, _LRWildcardMarker}]
	return action, ok
}

var _LRActionTable = _LRActionTableType{
	{_LRState1, _LREndMarker}:             &_LRAction{_LRAcceptAction, 0, ""},
	{_LRState0, LRTokenToken}:             _LRGotoState4Action,
	{_LRState0, LRTypeToken}:              _LRGotoState5Action,
	{_LRState0, LRStartToken}:             _LRGotoState3Action,
	{_LRState0, LRRuleDefToken}:           _LRGotoState2Action,
	{_LRState0, LRGrammarType}:            _LRGotoState1Action,
	{_LRState0, LRDefsType}:               _LRGotoState7Action,
	{_LRState0, LRDefType}:                _LRGotoState6Action,
	{_LRState0, LRRwordType}:              _LRGotoState9Action,
	{_LRState0, LRRuleType}:               _LRGotoState8Action,
	{_LRState2, LRLabelToken}:             _LRGotoState11Action,
	{_LRState2, LRIdentifierToken}:        _LRGotoState10Action,
	{_LRState2, LRNonemptyIdentListType}:  _LRGotoState15Action,
	{_LRState2, LRIdentListType}:          _LRGotoState12Action,
	{_LRState2, LRLabeledClausesType}:     _LRGotoState14Action,
	{_LRState2, LRLabeledClauseType}:      _LRGotoState13Action,
	{_LRState3, LRIdentifierToken}:        _LRGotoState16Action,
	{_LRState6, LRSemicolonToken}:         _LRGotoState17Action,
	{_LRState7, LRTokenToken}:             _LRGotoState4Action,
	{_LRState7, LRTypeToken}:              _LRGotoState5Action,
	{_LRState7, LRStartToken}:             _LRGotoState3Action,
	{_LRState7, LRRuleDefToken}:           _LRGotoState2Action,
	{_LRState7, LRAdditionalSectionsType}: _LRGotoState18Action,
	{_LRState7, LRDefType}:                _LRGotoState19Action,
	{_LRState7, LRRwordType}:              _LRGotoState9Action,
	{_LRState7, LRRuleType}:               _LRGotoState8Action,
	{_LRState9, LRLtToken}:                _LRGotoState20Action,
	{_LRState9, LRIdentifierToken}:        _LRGotoState10Action,
	{_LRState9, LRNonemptyIdentListType}:  _LRGotoState21Action,
	{_LRState11, LRIdentifierToken}:       _LRGotoState10Action,
	{_LRState11, LRNonemptyIdentListType}: _LRGotoState15Action,
	{_LRState11, LRIdentListType}:         _LRGotoState22Action,
	{_LRState14, LROrToken}:               _LRGotoState23Action,
	{_LRState15, LRIdentifierToken}:       _LRGotoState24Action,
	{_LRState18, LRSectionMarkerToken}:    _LRGotoState25Action,
	{_LRState18, LRAdditionalSectionType}: _LRGotoState26Action,
	{_LRState19, LRSemicolonToken}:        _LRGotoState27Action,
	{_LRState20, LRIdentifierToken}:       _LRGotoState28Action,
	{_LRState21, LRIdentifierToken}:       _LRGotoState24Action,
	{_LRState23, LRLabelToken}:            _LRGotoState11Action,
	{_LRState23, LRLabeledClauseType}:     _LRGotoState29Action,
	{_LRState25, LRIdentifierToken}:       _LRGotoState30Action,
	{_LRState28, LRGtToken}:               _LRGotoState31Action,
	{_LRState30, LRSectionContentToken}:   _LRGotoState32Action,
	{_LRState31, LRIdentifierToken}:       _LRGotoState10Action,
	{_LRState31, LRNonemptyIdentListType}: _LRGotoState33Action,
	{_LRState33, LRIdentifierToken}:       _LRGotoState24Action,
	{_LRState2, _LRWildcardMarker}:        _LRReduceNilToIdentListAction,
	{_LRState4, _LRWildcardMarker}:        _LRReduceTokenToRwordAction,
	{_LRState5, _LRWildcardMarker}:        _LRReduceTypeToRwordAction,
	{_LRState6, _LRWildcardMarker}:        _LRReduceDefToDefsAction,
	{_LRState7, _LRWildcardMarker}:        _LRReduceNilToAdditionalSectionsAction,
	{_LRState8, _LRWildcardMarker}:        _LRReduceRuleToDefAction,
	{_LRState10, _LRWildcardMarker}:       _LRReduceIdentToNonemptyIdentListAction,
	{_LRState11, _LRWildcardMarker}:       _LRReduceNilToIdentListAction,
	{_LRState12, _LRWildcardMarker}:       _LRReduceUnlabeledClauseToRuleAction,
	{_LRState13, _LRWildcardMarker}:       _LRReduceClauseToLabeledClausesAction,
	{_LRState14, _LRWildcardMarker}:       _LRReduceClausesToRuleAction,
	{_LRState15, _LRWildcardMarker}:       _LRReduceNonEmptyListToIdentListAction,
	{_LRState16, _LRWildcardMarker}:       _LRReduceStartDeclToDefAction,
	{_LRState17, _LRWildcardMarker}:       _LRReduceExplicitDefToDefsAction,
	{_LRState18, _LREndMarker}:            _LRReduceToGrammarAction,
	{_LRState19, _LRWildcardMarker}:       _LRReduceAddToDefsAction,
	{_LRState21, _LRWildcardMarker}:       _LRReduceUntypedTermDeclToDefAction,
	{_LRState22, _LRWildcardMarker}:       _LRReduceToLabeledClauseAction,
	{_LRState24, _LRWildcardMarker}:       _LRReduceAddToNonemptyIdentListAction,
	{_LRState26, _LRWildcardMarker}:       _LRReduceAddToAdditionalSectionsAction,
	{_LRState27, _LRWildcardMarker}:       _LRReduceAddExplicitToDefsAction,
	{_LRState29, _LRWildcardMarker}:       _LRReduceAddToLabeledClausesAction,
	{_LRState32, _LRWildcardMarker}:       _LRReduceToAdditionalSectionAction,
	{_LRState33, _LRWildcardMarker}:       _LRReduceTermDeclToDefAction,
}

var _LRExpectedTerminals = map[_LRStateId][]LRSymbolId{
	_LRState0:  []LRSymbolId{LRTokenToken, LRTypeToken, LRStartToken, LRRuleDefToken},
	_LRState1:  []LRSymbolId{_LREndMarker},
	_LRState2:  []LRSymbolId{LRLabelToken, LRIdentifierToken},
	_LRState3:  []LRSymbolId{LRIdentifierToken},
	_LRState6:  []LRSymbolId{LRSemicolonToken},
	_LRState7:  []LRSymbolId{LRTokenToken, LRTypeToken, LRStartToken, LRRuleDefToken},
	_LRState9:  []LRSymbolId{LRLtToken, LRIdentifierToken},
	_LRState11: []LRSymbolId{LRIdentifierToken},
	_LRState14: []LRSymbolId{LROrToken},
	_LRState15: []LRSymbolId{LRIdentifierToken},
	_LRState18: []LRSymbolId{LRSectionMarkerToken, _LREndMarker},
	_LRState19: []LRSymbolId{LRSemicolonToken},
	_LRState20: []LRSymbolId{LRIdentifierToken},
	_LRState21: []LRSymbolId{LRIdentifierToken},
	_LRState23: []LRSymbolId{LRLabelToken},
	_LRState25: []LRSymbolId{LRIdentifierToken},
	_LRState28: []LRSymbolId{LRGtToken},
	_LRState30: []LRSymbolId{LRSectionContentToken},
	_LRState31: []LRSymbolId{LRIdentifierToken},
	_LRState33: []LRSymbolId{LRIdentifierToken},
}

/*
Parser Debug States:
  State 0:
    Kernel Items:
      #accept: ^.grammar
    Non-kernel Items:
      def:.START IDENTIFIER
      def:.rule
      def:.rword LT IDENTIFIER GT nonempty_ident_list
      def:.rword nonempty_ident_list
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
      #accept: ^ grammar., $
    Reduce:
      $ -> [#accept]
    Goto:
      (nil)

  State 2:
    Kernel Items:
      rule: RULE_DEF.ident_list
      rule: RULE_DEF.labeled_clauses
    Non-kernel Items:
      ident_list:., *
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
      def: START.IDENTIFIER
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 16

  State 4:
    Kernel Items:
      rword: TOKEN., *
    Reduce:
      * -> [rword]
    Goto:
      (nil)

  State 5:
    Kernel Items:
      rword: TYPE., *
    Reduce:
      * -> [rword]
    Goto:
      (nil)

  State 6:
    Kernel Items:
      defs: def., *
      defs: def.SEMICOLON
    Reduce:
      * -> [defs]
    Goto:
      SEMICOLON -> State 17

  State 7:
    Kernel Items:
      defs: defs.def
      defs: defs.def SEMICOLON
      grammar: defs.additional_sections
    Non-kernel Items:
      additional_sections:., *
      additional_sections:.additional_sections additional_section
      def:.START IDENTIFIER
      def:.rule
      def:.rword LT IDENTIFIER GT nonempty_ident_list
      def:.rword nonempty_ident_list
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
      def: rule., *
    Reduce:
      * -> [def]
    Goto:
      (nil)

  State 9:
    Kernel Items:
      def: rword.LT IDENTIFIER GT nonempty_ident_list
      def: rword.nonempty_ident_list
    Non-kernel Items:
      nonempty_ident_list:.IDENTIFIER
      nonempty_ident_list:.nonempty_ident_list IDENTIFIER
    Reduce:
      (nil)
    Goto:
      LT -> State 20
      IDENTIFIER -> State 10
      nonempty_ident_list -> State 21

  State 10:
    Kernel Items:
      nonempty_ident_list: IDENTIFIER., *
    Reduce:
      * -> [nonempty_ident_list]
    Goto:
      (nil)

  State 11:
    Kernel Items:
      labeled_clause: LABEL.ident_list
    Non-kernel Items:
      ident_list:., *
      ident_list:.nonempty_ident_list
      nonempty_ident_list:.IDENTIFIER
      nonempty_ident_list:.nonempty_ident_list IDENTIFIER
    Reduce:
      * -> [ident_list]
    Goto:
      IDENTIFIER -> State 10
      nonempty_ident_list -> State 15
      ident_list -> State 22

  State 12:
    Kernel Items:
      rule: RULE_DEF ident_list., *
    Reduce:
      * -> [rule]
    Goto:
      (nil)

  State 13:
    Kernel Items:
      labeled_clauses: labeled_clause., *
    Reduce:
      * -> [labeled_clauses]
    Goto:
      (nil)

  State 14:
    Kernel Items:
      labeled_clauses: labeled_clauses.OR labeled_clause
      rule: RULE_DEF labeled_clauses., *
    Reduce:
      * -> [rule]
    Goto:
      OR -> State 23

  State 15:
    Kernel Items:
      ident_list: nonempty_ident_list., *
      nonempty_ident_list: nonempty_ident_list.IDENTIFIER
    Reduce:
      * -> [ident_list]
    Goto:
      IDENTIFIER -> State 24

  State 16:
    Kernel Items:
      def: START IDENTIFIER., *
    Reduce:
      * -> [def]
    Goto:
      (nil)

  State 17:
    Kernel Items:
      defs: def SEMICOLON., *
    Reduce:
      * -> [defs]
    Goto:
      (nil)

  State 18:
    Kernel Items:
      additional_sections: additional_sections.additional_section
      grammar: defs additional_sections., $
    Non-kernel Items:
      additional_section:.SECTION_MARKER IDENTIFIER SECTION_CONTENT
    Reduce:
      $ -> [grammar]
    Goto:
      SECTION_MARKER -> State 25
      additional_section -> State 26

  State 19:
    Kernel Items:
      defs: defs def., *
      defs: defs def.SEMICOLON
    Reduce:
      * -> [defs]
    Goto:
      SEMICOLON -> State 27

  State 20:
    Kernel Items:
      def: rword LT.IDENTIFIER GT nonempty_ident_list
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 28

  State 21:
    Kernel Items:
      def: rword nonempty_ident_list., *
      nonempty_ident_list: nonempty_ident_list.IDENTIFIER
    Reduce:
      * -> [def]
    Goto:
      IDENTIFIER -> State 24

  State 22:
    Kernel Items:
      labeled_clause: LABEL ident_list., *
    Reduce:
      * -> [labeled_clause]
    Goto:
      (nil)

  State 23:
    Kernel Items:
      labeled_clauses: labeled_clauses OR.labeled_clause
    Non-kernel Items:
      labeled_clause:.LABEL ident_list
    Reduce:
      (nil)
    Goto:
      LABEL -> State 11
      labeled_clause -> State 29

  State 24:
    Kernel Items:
      nonempty_ident_list: nonempty_ident_list IDENTIFIER., *
    Reduce:
      * -> [nonempty_ident_list]
    Goto:
      (nil)

  State 25:
    Kernel Items:
      additional_section: SECTION_MARKER.IDENTIFIER SECTION_CONTENT
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 30

  State 26:
    Kernel Items:
      additional_sections: additional_sections additional_section., *
    Reduce:
      * -> [additional_sections]
    Goto:
      (nil)

  State 27:
    Kernel Items:
      defs: defs def SEMICOLON., *
    Reduce:
      * -> [defs]
    Goto:
      (nil)

  State 28:
    Kernel Items:
      def: rword LT IDENTIFIER.GT nonempty_ident_list
    Reduce:
      (nil)
    Goto:
      GT -> State 31

  State 29:
    Kernel Items:
      labeled_clauses: labeled_clauses OR labeled_clause., *
    Reduce:
      * -> [labeled_clauses]
    Goto:
      (nil)

  State 30:
    Kernel Items:
      additional_section: SECTION_MARKER IDENTIFIER.SECTION_CONTENT
    Reduce:
      (nil)
    Goto:
      SECTION_CONTENT -> State 32

  State 31:
    Kernel Items:
      def: rword LT IDENTIFIER GT.nonempty_ident_list
    Non-kernel Items:
      nonempty_ident_list:.IDENTIFIER
      nonempty_ident_list:.nonempty_ident_list IDENTIFIER
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 10
      nonempty_ident_list -> State 33

  State 32:
    Kernel Items:
      additional_section: SECTION_MARKER IDENTIFIER SECTION_CONTENT., *
    Reduce:
      * -> [additional_section]
    Goto:
      (nil)

  State 33:
    Kernel Items:
      def: rword LT IDENTIFIER GT nonempty_ident_list., *
      nonempty_ident_list: nonempty_ident_list.IDENTIFIER
    Reduce:
      * -> [def]
    Goto:
      IDENTIFIER -> State 24

Number of states: 34
Number of shift actions: 46
Number of reduce actions: 25
Number of shift/reduce conflicts: 0
Number of reduce/reduce conflicts: 0
*/
