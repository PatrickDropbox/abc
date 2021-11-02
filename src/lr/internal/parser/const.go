package parser

const (
	TokenKeyword = "%token"
	TypeKeyword  = "%type"

	// intermediate tokens used by the lexer, not directly consumed by the
	// parser.
	Arrow = "ARROW"
	Colon = "COLON"
)
