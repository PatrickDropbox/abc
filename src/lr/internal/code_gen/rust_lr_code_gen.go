package code_gen

import (
	"fmt"
	"io"

	"github.com/pattyshack/abc/src/lr/codegenutil"
	lr "github.com/pattyshack/abc/src/lr/internal"
	"github.com/pattyshack/abc/src/lr/internal/code_gen/rust_template"
	"github.com/pattyshack/abc/src/lr/internal/parser"
)

func populateRustCodeGenVariables(
	grammar *lr.Grammar,
	states *lr.LRStates) error {

	for _, term := range grammar.Terminals {
		if term.SymbolId == parser.LRIdentifierToken {
			term.CodeGenSymbolConst = codegenutil.SnakeToCamel(term.Name)
		}

		term.CodeGenType = term.ValueType
	}

	for _, term := range grammar.NonTerminals {
		term.CodeGenSymbolConst = codegenutil.SnakeToCamel(term.Name)
		term.CodeGenType = term.ValueType

		for _, clause := range term.Clauses {
			clause.CodeGenReducerName = clause.Label + "_to_" + term.Name
			if clause.Label == "" {
				clause.CodeGenReducerName = clause.CodeGenReducerName[1:]
			}

			clause.CodeGenReducerNameConst = codegenutil.SnakeToCamel(
				clause.CodeGenReducerName)
		}
	}

	for _, state := range states.OrderedStates {
		state.CodeGenConst = fmt.Sprintf("State%d", state.StateNum)
	}

	return nil
}

func GenerateRustLRCode(
	grammar *lr.Grammar,
	states *lr.LRStates) (
	io.WriterTo,
	error) {

	cfg := grammar.LangSpecs.Rust
	if cfg == nil {
		return nil, fmt.Errorf("rust configuration not specified in lang_specs")
	}

	for _, term := range grammar.Terms {
		if term.ValueType == lr.Generic {
			continue
		}

		valueType := cfg.ValueTypes[term.ValueType]
		if valueType == "" {
			return nil, fmt.Errorf(
				"Undefined value type for <%s> %s %v",
				term.ValueType,
				term.LRLocation,
				cfg.ValueTypes)
		}
	}

	err := populateRustCodeGenVariables(grammar, states)
	if err != nil {
		return nil, err
	}

	orderedSymbolNames := []string{lr.StartMarker, lr.Wildcard, lr.EndMarker}
	for _, term := range grammar.Terminals {
		orderedSymbolNames = append(orderedSymbolNames, term.Name)
	}
	for _, term := range grammar.NonTerminals {
		orderedSymbolNames = append(orderedSymbolNames, term.Name)
  }

  startMarker := &lr.Term{
      Name:               lr.StartMarker,
      IsTerminal:         true,
      ValueType:          lr.Generic,
    }
  wildcard := &lr.Term{
      Name:               lr.Wildcard,
      IsTerminal:         true,
      ValueType:          lr.Generic,
    }
  endMarker := &lr.Term{
      Name:               lr.EndMarker,
      IsTerminal:         true,
      ValueType:          lr.Generic,
    }

  grammar.Terms[startMarker.Name] = startMarker
  grammar.Terms[wildcard.Name] = wildcard
  grammar.Terms[endMarker.Name] = endMarker

	return &rust_template.File{
		Grammar:            grammar,
		States:             states,
		Cfg:                cfg,
		OrderedSymbolNames: orderedSymbolNames,
	}, nil
}
