package parser

import (
	"fmt"
)

type Definition interface {
	Location() LRLocation
	String() string
}

var _ LRSymbol = &Token{}

type Token struct {
	LRLocation

	LRSymbolId
	Value string
}

func (t *Token) Id() LRSymbolId {
	return t.LRSymbolId
}

func (t *Token) Location() LRLocation {
	return t.LRLocation
}

func (t *Token) String() string {
	return fmt.Sprintf("%v: %s (%v)", t.LRSymbolId, t.Value, t.Location)
}

type StartDeclaration struct {
	LRLocation

	Id *Token
}

func NewStartDeclaration(start *Token, id *Token) *StartDeclaration {
	return &StartDeclaration{
		LRLocation: start.LRLocation,
		Id:         id,
	}
}

func (sd *StartDeclaration) Location() LRLocation {
	return sd.LRLocation
}

func (sd *StartDeclaration) String() string {
	return "%start " + sd.Id.Value
}

type TermDeclaration struct {
	TermType *Token

	IsTerminal bool

	ValueType *Token

	Terms []*Token
}

func NewTermDeclaration(
	termType *Token,
	valueType *Token,
	terms []*Token) *TermDeclaration {

	return &TermDeclaration{
		TermType:   termType,
		IsTerminal: termType.Value == TokenKeyword,
		ValueType:  valueType,
		Terms:      terms,
	}
}

func (td *TermDeclaration) Location() LRLocation {
	return td.TermType.LRLocation
}

func (td *TermDeclaration) String() string {
	terms := ""
	for _, term := range td.Terms {
		terms += " " + term.Value
	}

	return td.TermType.Value + " <" + td.ValueType.Value + ">" + terms
}

type Clause struct {
	Label *Token // optional
	Body  []*Token

	// set by NewRule
	LRLocation
	Parent *Rule
}

func NewClause(label *Token, body []*Token) *Clause {
	return &Clause{
		Label: label,
		Body:  body,
	}
}

func (clause *Clause) String() string {
	result := "(none):"
	if clause.Label != nil {
		result = clause.Label.Value + ":"
	}

	for _, item := range clause.Body {
		result += " " + item.Value
	}

	return result
}

type Rule struct {
	Name    *Token
	Clauses []*Clause
}

func NewRule(name *Token, clauses []*Clause) *Rule {
	rule := &Rule{
		Name:    name,
		Clauses: clauses,
	}

	for _, clause := range clauses {
		loc := name.LRLocation
		if clause.Label != nil {
			loc = clause.Label.LRLocation
		} else if len(clause.Body) > 0 {
			loc = clause.Body[0].LRLocation
		}

		clause.LRLocation = loc
		clause.Parent = rule
	}

	return rule
}

func (r *Rule) Location() LRLocation {
	return r.Name.LRLocation
}

func (r *Rule) String() string {
	result := r.Name.Value + " ->"
	for i, clause := range r.Clauses {
		result += "\n    " + clause.String()
		if i < len(r.Clauses)-1 {
			result += " |"
		}
	}

	return result
}

type AdditionalSection struct {
	Name    *Token
	Content *Token
}

func NewAdditionalSection(name *Token, content *Token) *AdditionalSection {
	return &AdditionalSection{
		Name:    name,
		Content: content,
	}
}

type Grammar struct {
	Definitions []Definition

	AdditionalSections []*AdditionalSection
}

func NewGrammar(
	defs []Definition,
	additionalSections []*AdditionalSection) *Grammar {

	return &Grammar{defs, additionalSections}
}
