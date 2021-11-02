package lr

import (
	"fmt"
	"sort"
	"strings"
)

const (
	AcceptRule  = "#accept"
	StartMarker = "^"
	EndMarker   = "$"
	Wildcard    = "*"
)

// LR(1) production item of the form: A -> [a B] . [c] , x
type Item struct {
	Name     string
	Parsed   []string
	Expected []string

	LookAhead string

	*Term
	*Clause
}

func (item *Item) String() string {
	return item.Name + ":" +
		strings.Join(item.Parsed, " ") + "." +
		strings.Join(item.Expected, " ") + "," +
		item.LookAhead
}

type Items []*Item

func (items Items) KernelItems() Items {
	kernel := Items{}
	for _, item := range items {
		if len(item.Parsed) == 0 {
			continue
		}
		kernel = append(kernel, item)
	}

	sort.Sort(kernel)

	return kernel
}

func (items Items) Len() int {
	return len(items)
}

// sort order:
//  kernel items before non-kernel items
//  rule name
//  rule "path component"
//  longer parsed
//  parsed len (longer before shorter)
//  look ahead
func (items Items) Less(iIdx int, jIdx int) bool {
	i := items[iIdx]
	j := items[jIdx]

	if len(i.Parsed) > 0 {
		if len(j.Parsed) == 0 {
			return true
		}
	} else if len(j.Parsed) > 0 {
		return false
	}

	if i.Name != j.Name {
		return i.Name < j.Name
	}

	iRuleLen := len(i.Parsed) + len(i.Expected)
	jRuleLen := len(j.Parsed) + len(j.Expected)

	idx := 0
	for idx < iRuleLen && idx < jRuleLen {
		iPath := ""
		if len(i.Parsed) > idx {
			iPath = i.Parsed[idx]
		} else {
			iPath = i.Expected[idx-len(i.Parsed)]
		}

		jPath := ""
		if len(j.Parsed) > idx {
			jPath = j.Parsed[idx]
		} else {
			jPath = j.Expected[idx-len(j.Parsed)]
		}

		if iPath != jPath {
			return iPath < jPath
		}

		idx += 1
	}

	if idx < iRuleLen {
		return false
	}

	if idx < jRuleLen {
		return true
	}

	if len(i.Parsed) != len(j.Parsed) {
		return len(i.Parsed) > len(j.Parsed)
	}

	return i.LookAhead < j.LookAhead
}

func (items Items) Swap(i int, j int) {
	items[i], items[j] = items[j], items[i]
}

type ItemSet struct {
	Kernel string
	Items  // sorted

	CanonicalItems Items // LR1 canonical item set.  Sorted

	StateNum int

	Goto map[string]*ItemSet

	CodeGenConst string
}

type itemSetBuilder map[string]*Item

func (builder itemSetBuilder) maybeAdd(item *Item) bool {
	key := item.String()
	_, ok := builder[key]
	if ok {
		return false
	}

	builder[key] = item
	return true
}

func (builder itemSetBuilder) finalize() *ItemSet {
	canonicalItems := Items{}
	for _, item := range builder {
		canonicalItems = append(canonicalItems, item)
	}

	sort.Sort(canonicalItems)

	numReduce := 0
	wildcardReduce := itemSetBuilder{}
	for _, item := range builder {
		if len(item.Expected) == 0 {
			lookAhead := Wildcard
			if item.Name == AcceptRule {
				lookAhead = item.LookAhead
			}
			if wildcardReduce.maybeAdd(
				&Item{
					Name:      item.Name,
					Parsed:    item.Parsed,
					Expected:  item.Expected,
					LookAhead: lookAhead,

					Term:   item.Term,
					Clause: item.Clause,
				}) {

				numReduce += 1
			}
		} else {
			wildcardReduce.maybeAdd(item)
		}
	}

	items := Items{}
	if numReduce == 1 {
		for _, item := range wildcardReduce {
			items = append(items, item)
		}
	} else {
		for _, item := range builder {
			items = append(items, item)
		}
	}

	sort.Sort(items)

	chunks := []string{}
	for _, item := range items.KernelItems() {
		chunks = append(chunks, item.String())
	}

	return &ItemSet{
		Kernel:         strings.Join(chunks, ";"),
		Items:          items,
		CanonicalItems: canonicalItems,
		Goto:           map[string]*ItemSet{},
	}
}

type LRStates struct {
	*Grammar

	FirstTerms map[string]map[string]struct{}

	States        map[string]*ItemSet
	OrderedStates []*ItemSet

	Errors []string
}

func (states *LRStates) maybeAdd(builder itemSetBuilder) (*ItemSet, bool) {
	if len(builder) == 0 {
		return nil, false
	}

	states.populateClosure(builder)
	state := builder.finalize()

	origState, ok := states.States[state.Kernel]
	if ok {
		return origState, false
	}

	state.StateNum = len(states.OrderedStates)
	states.States[state.Kernel] = state
	states.OrderedStates = append(states.OrderedStates, state)

	return state, true
}

func (states *LRStates) populateStartAcceptStates() {
	startState := itemSetBuilder{}
	startState.maybeAdd(
		&Item{
			Name:      AcceptRule,
			Parsed:    []string{StartMarker},
			Expected:  []string{states.Start.Name},
			LookAhead: EndMarker,
		})
	states.maybeAdd(startState)

	// Pre-populate accept state for my own sanity.  This way, the first state
	// is always the start state, while the second state is always the accept
	// state.
	acceptState := itemSetBuilder{}
	acceptState.maybeAdd(
		&Item{
			Name:      AcceptRule,
			Parsed:    []string{StartMarker, states.Start.Name},
			Expected:  []string{},
			LookAhead: EndMarker,
		})
	states.maybeAdd(acceptState)
}

func (states *LRStates) generateStates() {

	symbols := []string{}
	for symbol, _ := range states.Terms {
		symbols = append(symbols, symbol)
	}
	sort.Strings(symbols)

	modified := true
	for modified {
		modified = false

		for _, state := range states.OrderedStates {
			for _, symbol := range symbols {
				gotoState, added := states.generateGotoState(state, symbol)
				if gotoState == nil {
					continue
				}

				if added {
					modified = true
				}

				nextState := state.Goto[symbol]
				if nextState == nil {
					state.Goto[symbol] = gotoState
				} else if nextState != gotoState {
					panic("This should never happen")
				}
			}
		}
	}
}

func (states *LRStates) generateGotoState(
	parentState *ItemSet,
	symbol string) (
	*ItemSet,
	bool) {

	state := itemSetBuilder{}

	for _, item := range parentState.Items {
		if len(item.Expected) > 0 && item.Expected[0] == symbol {
			state.maybeAdd(
				&Item{
					Name:      item.Name,
					Parsed:    append(item.Parsed, symbol),
					Expected:  item.Expected[1:],
					LookAhead: item.LookAhead,

					Term:   item.Term,
					Clause: item.Clause,
				})

		}
	}

	if len(state) == 0 {
		return nil, false
	}

	return states.maybeAdd(state)
}

func (states *LRStates) populateClosure(state itemSetBuilder) {
	modified := true
	for modified {
		modified = false
		for _, item := range state {
			if len(item.Expected) == 0 {
				continue
			}

			rule := states.Terms[item.Expected[0]]
			if rule.IsTerminal {
				continue
			}

			terminals := states.firstTerminals(
				append(item.Expected[1:], item.LookAhead))

			for _, clause := range rule.Clauses {
				clauseSymbols := []string{}
				for _, term := range clause.Bindings {
					clauseSymbols = append(clauseSymbols, term.Name)
				}

				for terminal, _ := range terminals {
					added := state.maybeAdd(
						&Item{
							Name:      rule.Name,
							Expected:  clauseSymbols,
							LookAhead: terminal,

							Term:   rule,
							Clause: clause,
						})
					if added {
						modified = true
					}
				}
			}
		}
	}
}

func (states *LRStates) firstTerminals(symbols []string) map[string]struct{} {
	result := map[string]struct{}{}
	for _, symbol := range symbols {
		hasNil := false
		for term, _ := range states.FirstTerms[symbol] {
			if term == "" {
				hasNil = true
			} else {
				result[term] = struct{}{}
			}
		}
		if !hasNil {
			return result
		}
	}

	panic(fmt.Sprintf("Shouldn't reach here: %v", symbols))
}

func (states *LRStates) computeFirstTerminals() {
	firstTerms := map[string]map[string]struct{}{
		"$": map[string]struct{}{
			"$": struct{}{},
		},
	}
	states.FirstTerms = firstTerms

	for _, term := range states.Terms {
		set := map[string]struct{}{}

		firstTerms[term.Name] = set
		if term.IsTerminal {
			set[term.Name] = struct{}{}
		} else {
			for _, clause := range term.Clauses {
				if len(clause.Body) == 0 {
					set[""] = struct{}{}
				}
			}
		}
	}

	modified := true
	for modified {
		modified = false

		for _, rule := range states.NonTerminals {
			set := firstTerms[rule.Name]
			add := func(symbol string) {
				_, ok := set[symbol]
				if !ok {
					modified = true
					set[symbol] = struct{}{}
				}
			}

			for _, clause := range rule.Clauses {
				nilTermCount := 0
				for _, term := range clause.Bindings {
					hasNil := false
					for symbol, _ := range firstTerms[term.Name] {
						if symbol != "" {
							add(symbol)
						} else {
							hasNil = true
							nilTermCount += 1
						}
					}

					if !hasNil {
						break
					}
				}

				if nilTermCount == len(clause.Bindings) {
					add("")
				}
			}
		}
	}
}

func NewLRStates(grammar *Grammar) *LRStates {
	states := &LRStates{
		Grammar: grammar,
		States:  map[string]*ItemSet{},
	}

	states.computeFirstTerminals()
	states.populateStartAcceptStates()
	states.generateStates()

    // TODO check for shift/reduce & reduce/reduce errors
	return states
}
