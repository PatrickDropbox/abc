package yacc

import (
	"fmt"

	"github.com/pattyshack/abc/src/lr/internal/parser"
)

func LRSymbolTypeToYaccTokenNum(tt parser.LRSymbolId) int {
	switch tt {
	case parser.LRTokenToken:
		return TOKEN
	case parser.LRTypeToken:
		return TYPE
	case parser.LRStartToken:
		return START
	case parser.Arrow:
		return ARROW
	case parser.Colon:
		return COLON
	case parser.LRRuleDefToken:
		return RULE_DEF
	case parser.LRLabelToken:
		return LABEL
	case parser.LRLtToken:
		return LT
	case parser.LRGtToken:
		return GT
	case parser.LROrToken:
		return OR
	case parser.LRSemicolonToken:
		return SEMICOLON
	case parser.LRIdentifierToken:
		return IDENTIFIER
	case parser.LRSectionMarkerToken:
		return SECTION_MARKER
	case parser.LRSectionContentToken:
		return SECTION_CONTENT
	}

	panic(fmt.Sprintf("Unexpected LRSymbolId: %v", tt))
}
