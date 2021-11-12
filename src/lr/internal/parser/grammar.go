// Auto-generated from source: grammar.lr

package parser

import (
	fmt "fmt"
	io "io"
)

type LRSymbolId int

const (
	LRTokenToken          = LRSymbolId(256)
	LRTypeToken           = LRSymbolId(257)
	LRStartToken          = LRSymbolId(258)
	LRRuleDefToken        = LRSymbolId(259)
	LRLabelToken          = LRSymbolId(260)
	LRLtToken             = LRSymbolId(261)
	LRGtToken             = LRSymbolId(262)
	LROrToken             = LRSymbolId(263)
	LRSemicolonToken      = LRSymbolId(264)
	LRSectionMarkerToken  = LRSymbolId(265)
	LRIdentifierToken     = LRSymbolId(266)
	LRSectionContentToken = LRSymbolId(267)
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
	StartDeclToDef(Start_ *LRGenericSymbol, NonemptyIdentList_ []*Token) (Definition, error)

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
	item, err := _LRParse(lexer, reducer, errHandler, _LRState1)
	if err != nil {
		var errRetVal *Grammar
		return errRetVal, err
	}
	return item.Grammar, nil
}

// ================================================================
// Parser internal implementation
// User should normally avoid directly accessing the following code
// ================================================================

func _LRParse(lexer LRLexer, reducer LRReducer, errHandler LRParseErrorHandler, startState _LRStateId) (*_LRStackItem, error) {
	stateStack := _LRStack{
		// Note: we don't have to populate the start symbol since its value is never accessed
		&_LRStackItem{startState, nil},
	}
	symbolStack := &_LRPseudoSymbolStack{lexer: lexer}

	for {
		nextSymbol, err := symbolStack.Top()
		if err != nil {
			return nil, err
		}

		action, ok := _LRActionTable.Get(stateStack[len(stateStack)-1].StateId, nextSymbol.Id())
		if !ok {
			return nil, errHandler.Error(nextSymbol, stateStack)
		}
		if action.ActionType == _LRShiftAction {
			stateStack = append(stateStack, action.ShiftItem(nextSymbol))

			_, err = symbolStack.Pop()
			if err != nil {
				return nil, err
			}
		} else if action.ActionType == _LRReduceAction {
			var reduceSymbol *LRSymbol
			stateStack, reduceSymbol, err = action.ReduceSymbol(reducer, stateStack)
			if err != nil {
				return nil, err
			}

			symbolStack.Push(reduceSymbol)
		} else if action.ActionType == _LRAcceptAction {
			if len(stateStack) != 2 {
				panic("This should never happen")
			}
			return stateStack[1], nil

		} else {
			panic("Unknown action type: " + action.ActionType.String())
		}
	}
}

func (i LRSymbolId) String() string {
	switch i {
	case _LREndMarker:
		return "$"
	case _LRWildcardMarker:
		return "*"
	case LRTokenToken:
		return "TOKEN"
	case LRTypeToken:
		return "TYPE"
	case LRStartToken:
		return "START"
	case LRRuleDefToken:
		return "RULE_DEF"
	case LRLabelToken:
		return "LABEL"
	case LRLtToken:
		return "LT"
	case LRGtToken:
		return "GT"
	case LROrToken:
		return "OR"
	case LRSemicolonToken:
		return "SEMICOLON"
	case LRSectionMarkerToken:
		return "SECTION_MARKER"
	case LRIdentifierToken:
		return "IDENTIFIER"
	case LRSectionContentToken:
		return "SECTION_CONTENT"
	case LRGrammarType:
		return "grammar"
	case LRAdditionalSectionsType:
		return "additional_sections"
	case LRAdditionalSectionType:
		return "additional_section"
	case LRDefsType:
		return "defs"
	case LRDefType:
		return "def"
	case LRRwordType:
		return "rword"
	case LRNonemptyIdentListType:
		return "nonempty_ident_list"
	case LRIdentListType:
		return "ident_list"
	case LRRuleType:
		return "rule"
	case LRLabeledClausesType:
		return "labeled_clauses"
	case LRLabeledClauseType:
		return "labeled_clause"
	default:
		return fmt.Sprintf("?unknown symbol %d?", int(i))
	}
}

const (
	_LREndMarker      = LRSymbolId(0)
	_LRWildcardMarker = LRSymbolId(-1)

	LRGrammarType            = LRSymbolId(268)
	LRAdditionalSectionsType = LRSymbolId(269)
	LRAdditionalSectionType  = LRSymbolId(270)
	LRDefsType               = LRSymbolId(271)
	LRDefType                = LRSymbolId(272)
	LRRwordType              = LRSymbolId(273)
	LRNonemptyIdentListType  = LRSymbolId(274)
	LRIdentListType          = LRSymbolId(275)
	LRRuleType               = LRSymbolId(276)
	LRLabeledClausesType     = LRSymbolId(277)
	LRLabeledClauseType      = LRSymbolId(278)
)

type _LRActionType int

const (
	// NOTE: error action is implicit
	_LRShiftAction  = _LRActionType(0)
	_LRReduceAction = _LRActionType(1)
	_LRAcceptAction = _LRActionType(2)
)

func (i _LRActionType) String() string {
	switch i {
	case _LRShiftAction:
		return "shift"
	case _LRReduceAction:
		return "reduce"
	case _LRAcceptAction:
		return "accept"
	default:
		return fmt.Sprintf("?unknown action %d", int(i))
	}
}

type _LRReduceType int

const (
	_LRReduceToGrammar                = _LRReduceType(1)
	_LRReduceAddToAdditionalSections  = _LRReduceType(2)
	_LRReduceNilToAdditionalSections  = _LRReduceType(3)
	_LRReduceToAdditionalSection      = _LRReduceType(4)
	_LRReduceAddToDefs                = _LRReduceType(5)
	_LRReduceAddExplicitToDefs        = _LRReduceType(6)
	_LRReduceDefToDefs                = _LRReduceType(7)
	_LRReduceExplicitDefToDefs        = _LRReduceType(8)
	_LRReduceTermDeclToDef            = _LRReduceType(9)
	_LRReduceUntypedTermDeclToDef     = _LRReduceType(10)
	_LRReduceStartDeclToDef           = _LRReduceType(11)
	_LRReduceRuleToDef                = _LRReduceType(12)
	_LRReduceTokenToRword             = _LRReduceType(13)
	_LRReduceTypeToRword              = _LRReduceType(14)
	_LRReduceAddToNonemptyIdentList   = _LRReduceType(15)
	_LRReduceIdentToNonemptyIdentList = _LRReduceType(16)
	_LRReduceNonEmptyListToIdentList  = _LRReduceType(17)
	_LRReduceNilToIdentList           = _LRReduceType(18)
	_LRReduceUnlabeledClauseToRule    = _LRReduceType(19)
	_LRReduceClausesToRule            = _LRReduceType(20)
	_LRReduceAddToLabeledClauses      = _LRReduceType(21)
	_LRReduceClauseToLabeledClauses   = _LRReduceType(22)
	_LRReduceToLabeledClause          = _LRReduceType(23)
)

func (i _LRReduceType) String() string {
	switch i {
	case _LRReduceToGrammar:
		return "ToGrammar"
	case _LRReduceAddToAdditionalSections:
		return "AddToAdditionalSections"
	case _LRReduceNilToAdditionalSections:
		return "NilToAdditionalSections"
	case _LRReduceToAdditionalSection:
		return "ToAdditionalSection"
	case _LRReduceAddToDefs:
		return "AddToDefs"
	case _LRReduceAddExplicitToDefs:
		return "AddExplicitToDefs"
	case _LRReduceDefToDefs:
		return "DefToDefs"
	case _LRReduceExplicitDefToDefs:
		return "ExplicitDefToDefs"
	case _LRReduceTermDeclToDef:
		return "TermDeclToDef"
	case _LRReduceUntypedTermDeclToDef:
		return "UntypedTermDeclToDef"
	case _LRReduceStartDeclToDef:
		return "StartDeclToDef"
	case _LRReduceRuleToDef:
		return "RuleToDef"
	case _LRReduceTokenToRword:
		return "TokenToRword"
	case _LRReduceTypeToRword:
		return "TypeToRword"
	case _LRReduceAddToNonemptyIdentList:
		return "AddToNonemptyIdentList"
	case _LRReduceIdentToNonemptyIdentList:
		return "IdentToNonemptyIdentList"
	case _LRReduceNonEmptyListToIdentList:
		return "NonEmptyListToIdentList"
	case _LRReduceNilToIdentList:
		return "NilToIdentList"
	case _LRReduceUnlabeledClauseToRule:
		return "UnlabeledClauseToRule"
	case _LRReduceClausesToRule:
		return "ClausesToRule"
	case _LRReduceAddToLabeledClauses:
		return "AddToLabeledClauses"
	case _LRReduceClauseToLabeledClauses:
		return "ClauseToLabeledClauses"
	case _LRReduceToLabeledClause:
		return "ToLabeledClause"
	default:
		return fmt.Sprintf("?unknown reduce type %d?", int(i))
	}
}

type _LRStateId int

func (id _LRStateId) String() string {
	return fmt.Sprintf("State %d", int(id))
}

const (
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
	_LRState34 = _LRStateId(34)
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
		symbol.Definition, err = reducer.StartDeclToDef(args[0].Generic_, args[1].Tokens)
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
		panic("Unknown reduce type: " + act.ReduceType.String())
	}

	if err != nil {
		err = fmt.Errorf("Unexpected %s reduce error: %s", act.ReduceType, err)
	}

	return stack, symbol, err
}

var (
	_LRGotoState1Action                     = &_LRAction{_LRShiftAction, _LRState1, 0}
	_LRGotoState2Action                     = &_LRAction{_LRShiftAction, _LRState2, 0}
	_LRGotoState3Action                     = &_LRAction{_LRShiftAction, _LRState3, 0}
	_LRGotoState4Action                     = &_LRAction{_LRShiftAction, _LRState4, 0}
	_LRGotoState5Action                     = &_LRAction{_LRShiftAction, _LRState5, 0}
	_LRGotoState6Action                     = &_LRAction{_LRShiftAction, _LRState6, 0}
	_LRGotoState7Action                     = &_LRAction{_LRShiftAction, _LRState7, 0}
	_LRGotoState8Action                     = &_LRAction{_LRShiftAction, _LRState8, 0}
	_LRGotoState9Action                     = &_LRAction{_LRShiftAction, _LRState9, 0}
	_LRGotoState10Action                    = &_LRAction{_LRShiftAction, _LRState10, 0}
	_LRGotoState11Action                    = &_LRAction{_LRShiftAction, _LRState11, 0}
	_LRGotoState12Action                    = &_LRAction{_LRShiftAction, _LRState12, 0}
	_LRGotoState13Action                    = &_LRAction{_LRShiftAction, _LRState13, 0}
	_LRGotoState14Action                    = &_LRAction{_LRShiftAction, _LRState14, 0}
	_LRGotoState15Action                    = &_LRAction{_LRShiftAction, _LRState15, 0}
	_LRGotoState16Action                    = &_LRAction{_LRShiftAction, _LRState16, 0}
	_LRGotoState17Action                    = &_LRAction{_LRShiftAction, _LRState17, 0}
	_LRGotoState18Action                    = &_LRAction{_LRShiftAction, _LRState18, 0}
	_LRGotoState19Action                    = &_LRAction{_LRShiftAction, _LRState19, 0}
	_LRGotoState20Action                    = &_LRAction{_LRShiftAction, _LRState20, 0}
	_LRGotoState21Action                    = &_LRAction{_LRShiftAction, _LRState21, 0}
	_LRGotoState22Action                    = &_LRAction{_LRShiftAction, _LRState22, 0}
	_LRGotoState23Action                    = &_LRAction{_LRShiftAction, _LRState23, 0}
	_LRGotoState24Action                    = &_LRAction{_LRShiftAction, _LRState24, 0}
	_LRGotoState25Action                    = &_LRAction{_LRShiftAction, _LRState25, 0}
	_LRGotoState26Action                    = &_LRAction{_LRShiftAction, _LRState26, 0}
	_LRGotoState27Action                    = &_LRAction{_LRShiftAction, _LRState27, 0}
	_LRGotoState28Action                    = &_LRAction{_LRShiftAction, _LRState28, 0}
	_LRGotoState29Action                    = &_LRAction{_LRShiftAction, _LRState29, 0}
	_LRGotoState30Action                    = &_LRAction{_LRShiftAction, _LRState30, 0}
	_LRGotoState31Action                    = &_LRAction{_LRShiftAction, _LRState31, 0}
	_LRGotoState32Action                    = &_LRAction{_LRShiftAction, _LRState32, 0}
	_LRGotoState33Action                    = &_LRAction{_LRShiftAction, _LRState33, 0}
	_LRGotoState34Action                    = &_LRAction{_LRShiftAction, _LRState34, 0}
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
	{_LRState2, _LREndMarker}:             &_LRAction{_LRAcceptAction, 0, 0},
	{_LRState1, LRTokenToken}:             _LRGotoState5Action,
	{_LRState1, LRTypeToken}:              _LRGotoState6Action,
	{_LRState1, LRStartToken}:             _LRGotoState4Action,
	{_LRState1, LRRuleDefToken}:           _LRGotoState3Action,
	{_LRState1, LRGrammarType}:            _LRGotoState2Action,
	{_LRState1, LRDefsType}:               _LRGotoState8Action,
	{_LRState1, LRDefType}:                _LRGotoState7Action,
	{_LRState1, LRRwordType}:              _LRGotoState10Action,
	{_LRState1, LRRuleType}:               _LRGotoState9Action,
	{_LRState3, LRLabelToken}:             _LRGotoState12Action,
	{_LRState3, LRIdentifierToken}:        _LRGotoState11Action,
	{_LRState3, LRNonemptyIdentListType}:  _LRGotoState16Action,
	{_LRState3, LRIdentListType}:          _LRGotoState13Action,
	{_LRState3, LRLabeledClausesType}:     _LRGotoState15Action,
	{_LRState3, LRLabeledClauseType}:      _LRGotoState14Action,
	{_LRState4, LRIdentifierToken}:        _LRGotoState11Action,
	{_LRState4, LRNonemptyIdentListType}:  _LRGotoState17Action,
	{_LRState7, LRSemicolonToken}:         _LRGotoState18Action,
	{_LRState8, LRTokenToken}:             _LRGotoState5Action,
	{_LRState8, LRTypeToken}:              _LRGotoState6Action,
	{_LRState8, LRStartToken}:             _LRGotoState4Action,
	{_LRState8, LRRuleDefToken}:           _LRGotoState3Action,
	{_LRState8, LRAdditionalSectionsType}: _LRGotoState19Action,
	{_LRState8, LRDefType}:                _LRGotoState20Action,
	{_LRState8, LRRwordType}:              _LRGotoState10Action,
	{_LRState8, LRRuleType}:               _LRGotoState9Action,
	{_LRState10, LRLtToken}:               _LRGotoState21Action,
	{_LRState10, LRIdentifierToken}:       _LRGotoState11Action,
	{_LRState10, LRNonemptyIdentListType}: _LRGotoState22Action,
	{_LRState12, LRIdentifierToken}:       _LRGotoState11Action,
	{_LRState12, LRNonemptyIdentListType}: _LRGotoState16Action,
	{_LRState12, LRIdentListType}:         _LRGotoState23Action,
	{_LRState15, LROrToken}:               _LRGotoState24Action,
	{_LRState16, LRIdentifierToken}:       _LRGotoState25Action,
	{_LRState17, LRIdentifierToken}:       _LRGotoState25Action,
	{_LRState19, LRSectionMarkerToken}:    _LRGotoState26Action,
	{_LRState19, LRAdditionalSectionType}: _LRGotoState27Action,
	{_LRState20, LRSemicolonToken}:        _LRGotoState28Action,
	{_LRState21, LRIdentifierToken}:       _LRGotoState29Action,
	{_LRState22, LRIdentifierToken}:       _LRGotoState25Action,
	{_LRState24, LRLabelToken}:            _LRGotoState12Action,
	{_LRState24, LRLabeledClauseType}:     _LRGotoState30Action,
	{_LRState26, LRIdentifierToken}:       _LRGotoState31Action,
	{_LRState29, LRGtToken}:               _LRGotoState32Action,
	{_LRState31, LRSectionContentToken}:   _LRGotoState33Action,
	{_LRState32, LRIdentifierToken}:       _LRGotoState11Action,
	{_LRState32, LRNonemptyIdentListType}: _LRGotoState34Action,
	{_LRState34, LRIdentifierToken}:       _LRGotoState25Action,
	{_LRState3, _LRWildcardMarker}:        _LRReduceNilToIdentListAction,
	{_LRState5, _LRWildcardMarker}:        _LRReduceTokenToRwordAction,
	{_LRState6, _LRWildcardMarker}:        _LRReduceTypeToRwordAction,
	{_LRState7, _LRWildcardMarker}:        _LRReduceDefToDefsAction,
	{_LRState8, _LRWildcardMarker}:        _LRReduceNilToAdditionalSectionsAction,
	{_LRState9, _LRWildcardMarker}:        _LRReduceRuleToDefAction,
	{_LRState11, _LRWildcardMarker}:       _LRReduceIdentToNonemptyIdentListAction,
	{_LRState12, _LRWildcardMarker}:       _LRReduceNilToIdentListAction,
	{_LRState13, _LRWildcardMarker}:       _LRReduceUnlabeledClauseToRuleAction,
	{_LRState14, _LRWildcardMarker}:       _LRReduceClauseToLabeledClausesAction,
	{_LRState15, _LRWildcardMarker}:       _LRReduceClausesToRuleAction,
	{_LRState16, _LRWildcardMarker}:       _LRReduceNonEmptyListToIdentListAction,
	{_LRState17, _LRWildcardMarker}:       _LRReduceStartDeclToDefAction,
	{_LRState18, _LRWildcardMarker}:       _LRReduceExplicitDefToDefsAction,
	{_LRState19, _LREndMarker}:            _LRReduceToGrammarAction,
	{_LRState20, _LRWildcardMarker}:       _LRReduceAddToDefsAction,
	{_LRState22, _LRWildcardMarker}:       _LRReduceUntypedTermDeclToDefAction,
	{_LRState23, _LRWildcardMarker}:       _LRReduceToLabeledClauseAction,
	{_LRState25, _LRWildcardMarker}:       _LRReduceAddToNonemptyIdentListAction,
	{_LRState27, _LRWildcardMarker}:       _LRReduceAddToAdditionalSectionsAction,
	{_LRState28, _LRWildcardMarker}:       _LRReduceAddExplicitToDefsAction,
	{_LRState30, _LRWildcardMarker}:       _LRReduceAddToLabeledClausesAction,
	{_LRState33, _LRWildcardMarker}:       _LRReduceToAdditionalSectionAction,
	{_LRState34, _LRWildcardMarker}:       _LRReduceTermDeclToDefAction,
}

var _LRExpectedTerminals = map[_LRStateId][]LRSymbolId{
	_LRState1:  []LRSymbolId{LRTokenToken, LRTypeToken, LRStartToken, LRRuleDefToken},
	_LRState2:  []LRSymbolId{_LREndMarker},
	_LRState3:  []LRSymbolId{LRLabelToken, LRIdentifierToken},
	_LRState4:  []LRSymbolId{LRIdentifierToken},
	_LRState7:  []LRSymbolId{LRSemicolonToken},
	_LRState8:  []LRSymbolId{LRTokenToken, LRTypeToken, LRStartToken, LRRuleDefToken},
	_LRState10: []LRSymbolId{LRLtToken, LRIdentifierToken},
	_LRState12: []LRSymbolId{LRIdentifierToken},
	_LRState15: []LRSymbolId{LROrToken},
	_LRState16: []LRSymbolId{LRIdentifierToken},
	_LRState17: []LRSymbolId{LRIdentifierToken},
	_LRState19: []LRSymbolId{LRSectionMarkerToken, _LREndMarker},
	_LRState20: []LRSymbolId{LRSemicolonToken},
	_LRState21: []LRSymbolId{LRIdentifierToken},
	_LRState22: []LRSymbolId{LRIdentifierToken},
	_LRState24: []LRSymbolId{LRLabelToken},
	_LRState26: []LRSymbolId{LRIdentifierToken},
	_LRState29: []LRSymbolId{LRGtToken},
	_LRState31: []LRSymbolId{LRSectionContentToken},
	_LRState32: []LRSymbolId{LRIdentifierToken},
	_LRState34: []LRSymbolId{LRIdentifierToken},
}

/*
Parser Debug States:
  State 1:
    Kernel Items:
      #accept: ^.grammar
    Non-kernel Items:
      def:.START nonempty_ident_list
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
      TOKEN -> State 5
      TYPE -> State 6
      START -> State 4
      RULE_DEF -> State 3
      grammar -> State 2
      defs -> State 8
      def -> State 7
      rword -> State 10
      rule -> State 9

  State 2:
    Kernel Items:
      #accept: ^ grammar., $
    Reduce:
      $ -> [#accept]
    Goto:
      (nil)

  State 3:
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
      LABEL -> State 12
      IDENTIFIER -> State 11
      nonempty_ident_list -> State 16
      ident_list -> State 13
      labeled_clauses -> State 15
      labeled_clause -> State 14

  State 4:
    Kernel Items:
      def: START.nonempty_ident_list
    Non-kernel Items:
      nonempty_ident_list:.IDENTIFIER
      nonempty_ident_list:.nonempty_ident_list IDENTIFIER
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 11
      nonempty_ident_list -> State 17

  State 5:
    Kernel Items:
      rword: TOKEN., *
    Reduce:
      * -> [rword]
    Goto:
      (nil)

  State 6:
    Kernel Items:
      rword: TYPE., *
    Reduce:
      * -> [rword]
    Goto:
      (nil)

  State 7:
    Kernel Items:
      defs: def., *
      defs: def.SEMICOLON
    Reduce:
      * -> [defs]
    Goto:
      SEMICOLON -> State 18

  State 8:
    Kernel Items:
      defs: defs.def
      defs: defs.def SEMICOLON
      grammar: defs.additional_sections
    Non-kernel Items:
      additional_sections:., *
      additional_sections:.additional_sections additional_section
      def:.START nonempty_ident_list
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
      TOKEN -> State 5
      TYPE -> State 6
      START -> State 4
      RULE_DEF -> State 3
      additional_sections -> State 19
      def -> State 20
      rword -> State 10
      rule -> State 9

  State 9:
    Kernel Items:
      def: rule., *
    Reduce:
      * -> [def]
    Goto:
      (nil)

  State 10:
    Kernel Items:
      def: rword.LT IDENTIFIER GT nonempty_ident_list
      def: rword.nonempty_ident_list
    Non-kernel Items:
      nonempty_ident_list:.IDENTIFIER
      nonempty_ident_list:.nonempty_ident_list IDENTIFIER
    Reduce:
      (nil)
    Goto:
      LT -> State 21
      IDENTIFIER -> State 11
      nonempty_ident_list -> State 22

  State 11:
    Kernel Items:
      nonempty_ident_list: IDENTIFIER., *
    Reduce:
      * -> [nonempty_ident_list]
    Goto:
      (nil)

  State 12:
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
      IDENTIFIER -> State 11
      nonempty_ident_list -> State 16
      ident_list -> State 23

  State 13:
    Kernel Items:
      rule: RULE_DEF ident_list., *
    Reduce:
      * -> [rule]
    Goto:
      (nil)

  State 14:
    Kernel Items:
      labeled_clauses: labeled_clause., *
    Reduce:
      * -> [labeled_clauses]
    Goto:
      (nil)

  State 15:
    Kernel Items:
      labeled_clauses: labeled_clauses.OR labeled_clause
      rule: RULE_DEF labeled_clauses., *
    Reduce:
      * -> [rule]
    Goto:
      OR -> State 24

  State 16:
    Kernel Items:
      ident_list: nonempty_ident_list., *
      nonempty_ident_list: nonempty_ident_list.IDENTIFIER
    Reduce:
      * -> [ident_list]
    Goto:
      IDENTIFIER -> State 25

  State 17:
    Kernel Items:
      def: START nonempty_ident_list., *
      nonempty_ident_list: nonempty_ident_list.IDENTIFIER
    Reduce:
      * -> [def]
    Goto:
      IDENTIFIER -> State 25

  State 18:
    Kernel Items:
      defs: def SEMICOLON., *
    Reduce:
      * -> [defs]
    Goto:
      (nil)

  State 19:
    Kernel Items:
      additional_sections: additional_sections.additional_section
      grammar: defs additional_sections., $
    Non-kernel Items:
      additional_section:.SECTION_MARKER IDENTIFIER SECTION_CONTENT
    Reduce:
      $ -> [grammar]
    Goto:
      SECTION_MARKER -> State 26
      additional_section -> State 27

  State 20:
    Kernel Items:
      defs: defs def., *
      defs: defs def.SEMICOLON
    Reduce:
      * -> [defs]
    Goto:
      SEMICOLON -> State 28

  State 21:
    Kernel Items:
      def: rword LT.IDENTIFIER GT nonempty_ident_list
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 29

  State 22:
    Kernel Items:
      def: rword nonempty_ident_list., *
      nonempty_ident_list: nonempty_ident_list.IDENTIFIER
    Reduce:
      * -> [def]
    Goto:
      IDENTIFIER -> State 25

  State 23:
    Kernel Items:
      labeled_clause: LABEL ident_list., *
    Reduce:
      * -> [labeled_clause]
    Goto:
      (nil)

  State 24:
    Kernel Items:
      labeled_clauses: labeled_clauses OR.labeled_clause
    Non-kernel Items:
      labeled_clause:.LABEL ident_list
    Reduce:
      (nil)
    Goto:
      LABEL -> State 12
      labeled_clause -> State 30

  State 25:
    Kernel Items:
      nonempty_ident_list: nonempty_ident_list IDENTIFIER., *
    Reduce:
      * -> [nonempty_ident_list]
    Goto:
      (nil)

  State 26:
    Kernel Items:
      additional_section: SECTION_MARKER.IDENTIFIER SECTION_CONTENT
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 31

  State 27:
    Kernel Items:
      additional_sections: additional_sections additional_section., *
    Reduce:
      * -> [additional_sections]
    Goto:
      (nil)

  State 28:
    Kernel Items:
      defs: defs def SEMICOLON., *
    Reduce:
      * -> [defs]
    Goto:
      (nil)

  State 29:
    Kernel Items:
      def: rword LT IDENTIFIER.GT nonempty_ident_list
    Reduce:
      (nil)
    Goto:
      GT -> State 32

  State 30:
    Kernel Items:
      labeled_clauses: labeled_clauses OR labeled_clause., *
    Reduce:
      * -> [labeled_clauses]
    Goto:
      (nil)

  State 31:
    Kernel Items:
      additional_section: SECTION_MARKER IDENTIFIER.SECTION_CONTENT
    Reduce:
      (nil)
    Goto:
      SECTION_CONTENT -> State 33

  State 32:
    Kernel Items:
      def: rword LT IDENTIFIER GT.nonempty_ident_list
    Non-kernel Items:
      nonempty_ident_list:.IDENTIFIER
      nonempty_ident_list:.nonempty_ident_list IDENTIFIER
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 11
      nonempty_ident_list -> State 34

  State 33:
    Kernel Items:
      additional_section: SECTION_MARKER IDENTIFIER SECTION_CONTENT., *
    Reduce:
      * -> [additional_section]
    Goto:
      (nil)

  State 34:
    Kernel Items:
      def: rword LT IDENTIFIER GT nonempty_ident_list., *
      nonempty_ident_list: nonempty_ident_list.IDENTIFIER
    Reduce:
      * -> [def]
    Goto:
      IDENTIFIER -> State 25

Number of states: 34
Number of shift actions: 48
Number of reduce actions: 25
Number of shift/reduce conflicts: 0
Number of reduce/reduce conflicts: 0
*/
