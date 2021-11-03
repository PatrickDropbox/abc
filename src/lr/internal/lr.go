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
	Name string

	Rule []string
	Dot  int

	LookAhead string

	*Term
	*Clause

	Key  string
	Next *Item
	Core *Item
}

func NewCoreItem(name string, rule []string, term *Term, clause *Clause) *Item {
	return &Item{
		Name:   name,
		Rule:   rule,
		Term:   term,
		Clause: clause,
	}
}

func (item *Item) Shift() *Item {
	if item.Next == nil {
		next := *item
		next.Dot += 1
		next.Key = ""
		if item.Core != nil {
			next.Core = item.Core.Shift()
		}

		item.Next = &next
	}

	return item.Next
}

func (item *Item) ReplaceLookAhead(symbol string) *Item {
	if symbol == item.LookAhead {
		return item
	}

	newItem := *item

	if item.LookAhead == "" {
		newItem.Core = item
	}

	newItem.LookAhead = symbol
	newItem.Key = ""
	newItem.Next = nil

	return &newItem
}

func (item *Item) IsKernel() bool {
	return item.Dot != 0
}

func (item *Item) IsReduce() bool {
	return item.Dot == len(item.Rule)
}

func (item *Item) String() string {
	if item.Key == "" {
		result := ""
		if item.Core != nil {
			result = item.Core.String()
		} else {
			result += item.Name + ":"

			separator := ""
			for idx, symbol := range item.Rule {
				if idx == item.Dot {
					separator = "."
				}

				result += separator + symbol

				separator = " "
			}

			if item.Dot == len(item.Rule) {
				result += "."
			}
		}

		if item.LookAhead != "" {
			result += "," + item.LookAhead
		}

		item.Key = result
	}

	return item.Key
}

type itemPoolKey struct {
	Rule      string
	Label     string
	LookAhead string
}

type itemPool struct {
	// (rule, clause label, "") -> item
	coreItems map[itemPoolKey]*Item

	// (rule, clause label, look ahead) -> item
	firstItems map[itemPoolKey]*Item
}

func newItemPool() *itemPool {
	return &itemPool{map[itemPoolKey]*Item{}, map[itemPoolKey]*Item{}}
}

func (pool *itemPool) Get(term *Term, clause *Clause, lookAhead string) *Item {
	label := ""
	if clause.Label != nil {
		label = clause.Label.Value
	}
	key := itemPoolKey{term.Name, label, lookAhead}

	first, ok := pool.firstItems[key]
	if ok {
		return first
	}

	coreKey := itemPoolKey{term.Name, label, ""}
	core, ok := pool.coreItems[coreKey]
	if !ok {
		symbols := []string{}
		for _, term := range clause.Bindings {
			symbols = append(symbols, term.Name)
		}

		core = NewCoreItem(term.Name, symbols, term, clause)
		pool.coreItems[coreKey] = core
	}

	first = core.ReplaceLookAhead(lookAhead)
	pool.firstItems[key] = first

	return first
}

type Items []*Item

func (items Items) KernelItems() Items {
	kernel := Items{}
	for _, item := range items {
		if !item.IsKernel() {
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

	if i.IsKernel() {
		if !j.IsKernel() {
			return true
		}
	} else if j.IsKernel() {
		return false
	}

	if i.Name != j.Name {
		return i.Name < j.Name
	}

	idx := 0
	for idx < len(i.Rule) && idx < len(j.Rule) {
		iPath := i.Rule[idx]
		jPath := j.Rule[idx]

		if iPath != jPath {
			return iPath < jPath
		}

		idx += 1
	}

	if idx < len(i.Rule) {
		return false
	}

	if idx < len(j.Rule) {
		return true
	}

	if i.Dot != j.Dot {
		return i.Dot > j.Dot
	}

	return i.LookAhead < j.LookAhead
}

func (items Items) Swap(i int, j int) {
	items[i], items[j] = items[j], items[i]
}

type stateAction struct {
	shift  *ItemSet
	reduce string
}

type ItemSet struct {
	Kernel string
	Items  // sorted

	StateNum int

	Goto map[string]*ItemSet

	Reduce map[string]Items

	ReduceReduceConflictSymbols []string
	ShiftReduceConflictSymbols  []string

	CodeGenConst string
}

func (set *ItemSet) canMergeFrom(other *ItemSet) bool {
	actions := map[string]*stateAction{}
	for symbol, items := range set.Reduce {
		if len(items) > 1 {
			return false // don't merge state with reduce/reduce error
		}
		actions[symbol] = &stateAction{reduce: items[0].String()}
	}

	for symbol, next := range set.Goto {
		actions[symbol] = &stateAction{shift: next}
	}

	for symbol, items := range other.Reduce {
		if len(items) > 1 {
			return false // don't merge state with reduce/reduce error
		}

		action, ok := actions[symbol]
		if !ok {
			continue
		}

		if action.shift != nil {
			// This introduces new shift/reduce error.  In theory, this should
			// never happen
			return false
		}

		if action.reduce != items[0].String() {
			return false // This introduces new reduce/reduce error
		}
	}

	for symbol, next := range set.Goto {
		action, ok := actions[symbol]
		if !ok {
			continue
		}

		if action.reduce != "" {
			// This introduces new shift/reduce error.  In theory, this should
			// never happen
			return false
		}

		if action.shift.StateNum != next.StateNum {
			// Don't merge different state machines
			return false
		}
	}

	return true
}

func (set *ItemSet) mergeFrom(other *ItemSet) {
	builder := itemSetBuilder{}
	for _, item := range set.Items {
		builder.maybeAdd(item)
	}

	for _, item := range other.Items {
		builder.maybeAdd(item)
	}

	temp := builder.finalize()

	set.Kernel = temp.Kernel
	set.Items = temp.Items

	for symbol, next := range other.Goto {
		set.Goto[symbol] = next
	}

	for symbol, items := range other.Reduce {
		set.Reduce[symbol] = items
	}
}

func (set *ItemSet) clone() *ItemSet {
	gotoMap := map[string]*ItemSet{}
	for symbol, state := range set.Goto {
		gotoMap[symbol] = state
	}

	reduce := map[string]Items{}
	for symbol, items := range set.Reduce {
		reduce[symbol] = items
	}

	return &ItemSet{
		Kernel:   set.Kernel,
		StateNum: set.StateNum,
		Items:    set.Items,
		Goto:     gotoMap,
		Reduce:   reduce,
	}
}

func (set *ItemSet) computeConflictSymbols() {
	for symbol, items := range set.Reduce {
		_, ok := set.Goto[symbol]
		if ok {
			set.ShiftReduceConflictSymbols = append(
				set.ShiftReduceConflictSymbols,
				symbol)
		}

		if len(items) > 1 {
			// NOTE: In theory, we can handle reducing multiple productions
			// as long as neither production rule is a prefix of the other
			// production rule.  However, this require inspecting the stack
			// which is kind of painful to code gen.

			set.ReduceReduceConflictSymbols = append(
				set.ReduceReduceConflictSymbols,
				symbol)
		}
	}
}

func (set *ItemSet) compress() {
	if len(set.ShiftReduceConflictSymbols) > 0 ||
		len(set.ReduceReduceConflictSymbols) > 0 {

		return // Don't compress error state to output more debug info
	}

	type itemCount struct {
		*Item
		Count int
	}

	builder := itemSetBuilder{}
	counts := map[string]*itemCount{}
	for _, item := range set.Items {
		if item.IsReduce() {
			key := item.Core.String()
			entry, ok := counts[key]
			if !ok {
				entry = &itemCount{item, 0}
				counts[key] = entry
			}
			entry.Count += 1
		} else {
			// Shift production does not depend on look ahead symbols
			builder.maybeAdd(item.ReplaceLookAhead(""))
		}
	}

	max := 0
	maxKey := ""
	for key, entry := range counts {
		if entry.Count > max {
			max = entry.Count
			maxKey = key
		}
	}

	if max == 1 {
		maxKey = ""
	}

	reduce := map[string]Items{}
	for key, entry := range counts {
		if key == maxKey {
			entry.Item = entry.ReplaceLookAhead(Wildcard)
		}

		builder.maybeAdd(entry.Item)
		reduce[entry.LookAhead] = append(reduce[entry.LookAhead], entry.Item)
	}

	temp := builder.finalize()

	set.Kernel = temp.Kernel
	set.Items = temp.Items

	set.Reduce = reduce
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
	items := Items{}
	for _, item := range builder {
		items = append(items, item)
	}

	sort.Sort(items)

	chunks := []string{}
	for _, item := range items.KernelItems() {
		chunks = append(chunks, item.String())
	}

	reduce := map[string]Items{}

	for _, item := range items {
		if item.IsReduce() {
			reduce[item.LookAhead] = append(reduce[item.LookAhead], item)
		}
	}

	return &ItemSet{
		Kernel: strings.Join(chunks, ";"),
		Items:  items,
		Goto:   map[string]*ItemSet{},
		Reduce: reduce,
	}
}

type LRStates struct {
	*Grammar

	FirstTerms map[string]map[string]struct{}

	ItemPool *itemPool

	States        map[string]*ItemSet
	OrderedStates []*ItemSet
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

func (states *LRStates) populateStartStates() {
	core := NewCoreItem(
		AcceptRule,
		[]string{StartMarker, states.Start.Name},
		nil,
		nil)

	startState := itemSetBuilder{}
	startState.maybeAdd(core.ReplaceLookAhead(EndMarker).Shift())
	states.maybeAdd(startState)

}

func (states *LRStates) generateStates() {

	symbols := []string{}
	for symbol, _ := range states.Terms {
		if symbol == states.Start.Name {
			continue
		}
		symbols = append(symbols, symbol)
	}
	sort.Strings(symbols)

	// NOTE: by placing the start symbol at the front of the list,
	// we ensure that the second state is always the accept state.
	symbols = append([]string{states.Start.Name}, symbols...)

	exploredIdx := 0
	for exploredIdx < len(states.OrderedStates) {
		for _, state := range states.OrderedStates[exploredIdx:] {
			for _, symbol := range symbols {
				gotoState, _ := states.generateGotoState(state, symbol)
				if gotoState == nil {
					continue
				}

				nextState := state.Goto[symbol]
				if nextState == nil {
					state.Goto[symbol] = gotoState
				} else if nextState != gotoState {
					panic("This should never happen")
				}
			}

			exploredIdx += 1
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
		if !item.IsReduce() && item.Rule[item.Dot] == symbol {
			state.maybeAdd(item.Shift())
		}
	}

	if len(state) == 0 {
		return nil, false
	}

	return states.maybeAdd(state)
}

func (states *LRStates) populateClosure(state itemSetBuilder) {
	toExplore := state
	for len(toExplore) > 0 {
		nextToExplore := itemSetBuilder{}

		for _, item := range toExplore {
			if item.IsReduce() {
				continue
			}

			rule := states.Terms[item.Rule[item.Dot]]
			if rule.IsTerminal {
				continue
			}

			terminals := states.firstTerminals(
				append(item.Rule[item.Dot+1:], item.LookAhead))

			for _, clause := range rule.Clauses {
				for terminal, _ := range terminals {
					item := states.ItemPool.Get(rule, clause, terminal)

					added := state.maybeAdd(item)
					if added {
						nextToExplore.maybeAdd(item)
					}
				}
			}
		}

		toExplore = nextToExplore
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

func (states *LRStates) mergeStates() {
	modified := true
	for modified {
		modified = false

		coreKernels := []string{}
		mergeCandidates := map[string][]*ItemSet{}

		for _, state := range states.OrderedStates {
			core := itemSetBuilder{}

			for _, item := range state.Items {
				core.maybeAdd(item.ReplaceLookAhead(""))
			}

			coreState := core.finalize()

			_, ok := mergeCandidates[coreState.Kernel]
			if !ok {
				coreKernels = append(coreKernels, coreState.Kernel)
			}

			mergeCandidates[coreState.Kernel] = append(
				mergeCandidates[coreState.Kernel],
				state)
		}

		newStates := []*ItemSet{}

		// lr state kernel -> merged state
		stateMapping := map[string]*ItemSet{}

		// NOTE: iterate over coreKernels to preserve state ordering
		for _, coreKernel := range coreKernels {
			candidates := mergeCandidates[coreKernel]

			mergedStates := []*ItemSet{}
			for _, candidate := range candidates {
				merged := false
				for _, mergedState := range mergedStates {
					if mergedState.canMergeFrom(candidate) {
						mergedState.mergeFrom(candidate)

						stateMapping[candidate.Kernel] = mergedState
						merged = true
						break
					}
				}

				if !merged {
					newState := candidate.clone()

					mergedStates = append(mergedStates, newState)
					newStates = append(newStates, newState)

					stateMapping[candidate.Kernel] = newState
				}
			}
		}

		for idx, state := range newStates {
			newGoto := map[string]*ItemSet{}
			for symbol, next := range state.Goto {
				newGoto[symbol] = stateMapping[next.Kernel]
			}
			state.Goto = newGoto

			state.StateNum = idx
		}

		if len(states.OrderedStates) > len(newStates) {
			modified = true

			states.OrderedStates = newStates

			states.States = map[string]*ItemSet{}
			for _, state := range newStates {
				states.States[state.Kernel] = state
			}
		}
	}
}

func NewLRStates(grammar *Grammar) *LRStates {
	states := &LRStates{
		Grammar:  grammar,
		States:   map[string]*ItemSet{},
		ItemPool: newItemPool(),
	}

	states.computeFirstTerminals()
	states.populateStartStates()
	states.generateStates()

	states.mergeStates()

	for _, state := range states.OrderedStates {
		state.computeConflictSymbols()
		state.compress()
	}

	return states
}
