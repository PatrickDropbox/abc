package lr

import (
	"fmt"
	"sort"
	"strings"

	"github.com/pattyshack/abc/src/lr/internal/parser"

	"gopkg.in/yaml.v2"
)

const (
	Generic = "Generic_"
)

type Clause struct {
	*parser.Clause

	Bindings []*Term

	// Temp variable populated by the code generator.
	CodeGenReducerName      string
	CodeGenReducerNameConst string
	CodeGenReduceAction     string
}

type Term struct {
	Name string
	parser.LRLocation

	*parser.TermDeclaration

	// Rule and ClauseBindings are nil iff the term is terminal
	*parser.Rule

	Clauses []*Clause

	Reachable bool

	CodeGenSymbolConst string

	// Temp variable populated by code generator.  The generated type is
	// language specific.
	CodeGenType interface{}
}

type Grammar struct {
	Source string

	Terms map[string]*Term

	Terminals    []*Term // sorted by declaration location
	NonTerminals []*Term // sorted by rule location

	Start *Term

	*LangSpecs
}

func classifyDefinitions(
	parsed []parser.Definition) (
	map[string]*Term,
	map[string]*parser.Rule,
	string, // fist rule's name
	*parser.StartDeclaration,
	[]string) { // error strings

	terms := map[string]*Term{}

	rules := map[string]*parser.Rule{}

	startRuleName := ""
	var start *parser.StartDeclaration

	errStrs := []string{}

	for _, d := range parsed {
		switch def := d.(type) {
		case *parser.StartDeclaration:
			if start != nil {
				errStrs = append(
					errStrs,
					fmt.Sprintf(
						"Duplicate start declaration: %s %s",
						start.Location().ShortString(),
						def.Location().ShortString()))
			}

			start = def

		case *parser.TermDeclaration:
			for _, term := range def.Terms {
				if def.ValueType == nil {
					def.ValueType = &parser.Token{Value: Generic}
				}

				prev, ok := terms[term.Value]
				if ok {
					errStrs = append(
						errStrs,
						fmt.Sprintf(
							"Duplicate token/type declaration: %s %s %s",
							term.Value,
							prev.LRLocation.ShortString(),
							term.LRLocation.ShortString()))
				}

				terms[term.Value] = &Term{
					Name:            term.Value,
					LRLocation:      term.LRLocation,
					TermDeclaration: def,
					Rule:            nil,
					Reachable:       false,
				}
			}

		case *parser.Rule:
			if startRuleName == "" {
				startRuleName = def.Name.Value
			}

			prev, ok := rules[def.Name.Value]
			if ok {
				errStrs = append(
					errStrs,
					fmt.Sprintf(
						"Duplicate rule: %s %s %s",
						def.Name.Value,
						prev.Location().ShortString(),
						def.Location().ShortString()))
			}
			rules[def.Name.Value] = def
		}
	}

	return terms, rules, startRuleName, start, errStrs
}

func bindTerms(
	terms map[string]*Term,
	rules map[string]*parser.Rule,
	startRuleName string,
	start *parser.StartDeclaration) (*Term, []string) {

	errStrs := []string{}

	for name, rule := range rules {
		term, ok := terms[name]
		if !ok {
			errStrs = append(
				errStrs,
				fmt.Sprintf("Undefined type: %s %v", name, rule.Location()))
			continue
		}

		term.Rule = rule

		clauses := []*Clause{}
		for _, clause := range rule.Clauses {
			clause := &Clause{Clause: clause, Bindings: []*Term{}}

			for _, id := range clause.Body {
				t, ok := terms[id.Value]
				if ok {
					clause.Bindings = append(clause.Bindings, t)
				} else {
					errStrs = append(
						errStrs,
						fmt.Sprintf(
							"Undefined token/type: %s %v",
							id.Value,
							id.Location))
				}
			}

			clauses = append(clauses, clause)
		}

		term.Clauses = clauses
	}

	for name, term := range terms {
		if !term.IsTerminal && term.Rule == nil {
			errStrs = append(
				errStrs,
				fmt.Sprintf(
					"No rule specified for type: %s %v",
					name,
					term.LRLocation))
		} else if term.IsTerminal && term.Rule != nil {
			errStrs = append(
				errStrs,
				fmt.Sprintf(
					"token cannot have associated rule: %s %v",
					name,
					term.Rule.Location()))
		}
	}

	if start != nil {
		startRuleName = start.Id.Value
	}

	startTerm, ok := terms[startRuleName]
	if !ok || startTerm.IsTerminal {
		errStrs = append(
			errStrs,
			fmt.Sprintf("Invalid start rule: %s", startRuleName))
	}

	return startTerm, errStrs
}

func checkReachability(start *Term, terms map[string]*Term) []string {
	if start == nil {
		return nil
	}

	exploreSet := map[string]*Term{
		start.Name: start,
	}

	for len(exploreSet) > 0 {
		nextExploreSet := map[string]*Term{}

		for _, term := range exploreSet {
			if term.Reachable {
				continue
			}
			term.Reachable = true

			for _, clause := range term.Clauses {
				for _, item := range clause.Bindings {
					if !item.Reachable {
						nextExploreSet[item.Name] = item
					}
				}
			}
		}

		exploreSet = nextExploreSet
	}

	errStrs := []string{}
	for _, term := range terms {
		if !term.Reachable {
			errStrs = append(
				errStrs,
				fmt.Sprintf(
					"Unused token/type. Not reachable from start rule: %s %v",
					term.Name,
					term.LRLocation))
		}
	}

	return errStrs
}

func extractLangSpecs(
	sections []*parser.AdditionalSection) (
	*LangSpecs,
	[]string) {

	errStrs := []string{}

	var langSpecsSection *parser.AdditionalSection
	for _, section := range sections {
		if section.Name.Value != "lang_specs" {
			errStrs = append(
				errStrs,
				fmt.Sprintf(
					"Unexpected additional section: %s %v",
					section.Name.Value,
					section.Name.Location))
			continue
		}

		if langSpecsSection != nil {
			errStrs = append(
				errStrs,
				fmt.Sprintf(
					"Duplicated lang_specs section specified: %v %v",
					langSpecsSection.Name.Location,
					section.Name.Location))
		}

		langSpecsSection = section
	}

	langSpecs := &LangSpecs{}
	if langSpecsSection != nil {
		err := yaml.Unmarshal([]byte(langSpecsSection.Content.Value), langSpecs)
		if err != nil {
			errStrs = append(
				errStrs,
				fmt.Sprintf("Failed to unmarshal lang_specs: %s", err))
		}
	}

	return langSpecs, errStrs
}

func NewGrammar(
	sourceFile string,
	parsed *parser.Grammar) (
	*Grammar,
	error) {

	terms, rules, firstRuleName, start, errStrs := classifyDefinitions(
		parsed.Definitions)

	if len(rules) == 0 {
		errStrs = append(errStrs, "No rules specified in grammar.")
	}

	startTerm, bindErrStrs := bindTerms(terms, rules, firstRuleName, start)
	errStrs = append(errStrs, bindErrStrs...)

	errStrs = append(errStrs, checkReachability(startTerm, terms)...)

	langSpecs, asErrStrs := extractLangSpecs(parsed.AdditionalSections)
	errStrs = append(errStrs, asErrStrs...)

	if len(errStrs) > 0 {
		return nil, fmt.Errorf(strings.Join(errStrs, "\n"))
	}

	terminals := []*Term{}
	nonTerminals := []*Term{}

	for _, term := range terms {
		if term.IsTerminal {
			terminals = append(terminals, term)
		} else {
			nonTerminals = append(nonTerminals, term)
		}
	}

	sort.Sort(ByDeclLoc(terminals))
	sort.Sort(ByRuleLoc(nonTerminals))

	return &Grammar{
		Source:       sourceFile,
		Terms:        terms,
		Terminals:    terminals,
		NonTerminals: nonTerminals,
		Start:        startTerm,
		LangSpecs:    langSpecs,
	}, nil
}
