package yacc

import (
	"fmt"

	"github.com/pattyshack/abc/src/lr/internal/parser"
)

func LRSymbolTypeToYaccTokenNum(tt parser.LRSymbolId) int {
	switch tt {
	case parser.LRTokenSymbol:
		return TOKEN
	case parser.LRTypeSymbol:
		return TYPE
	case parser.LRStartSymbol:
		return START
	case parser.Arrow:
		return ARROW
	case parser.Colon:
		return COLON
	case parser.LRRuleDefSymbol:
		return RULE_DEF
	case parser.LRLabelSymbol:
		return LABEL
	case parser.LRLtSymbol:
		return LT
	case parser.LRGtSymbol:
		return GT
	case parser.LROrSymbol:
		return OR
	case parser.LRSemicolonSymbol:
		return SEMICOLON
	case parser.LRIdentifierSymbol:
		return IDENTIFIER
	case parser.LRSectionMarkerSymbol:
		return SECTION_MARKER
	case parser.LRSectionContentSymbol:
		return SECTION_CONTENT
	}

	panic(fmt.Sprintf("Unexpected LRSymbolId: %v", tt))
}
