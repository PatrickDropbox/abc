package parser

import (
	"bytes"
	"fmt"
	"io"
	"regexp"

	"github.com/pattyshack/abc/src/lr/parseutil"
)

var (
	keywordsAndSymbols = map[string]LRSymbolId{
		TokenKeyword: LRTokenToken,
		TypeKeyword:  LRTypeToken,
		"%start":     LRStartToken,
		"<":          '<',
		">":          '>',
		"|":          '|',
		";":          ';',

		"%%": LRSectionMarkerToken,

		"->": Arrow,
		":":  ':',
	}

	whitespaces = map[string]struct{}{
		" ":  struct{}{},
		"\n": struct{}{},
		"\t": struct{}{},
		"\r": struct{}{},
	}
)

type rawLexer struct {
	reader *parseutil.LocationReader
}

func newRawLexer(filename string, reader io.Reader) *rawLexer {
	return &rawLexer{parseutil.NewLocationReader(filename, reader)}
}

func (lexer *rawLexer) Next() (LRToken, error) {
	err := lexer.stripLeadingWhitespacesAndComments()
	if err != nil {
		return nil, err
	}

	_, err = lexer.reader.Peek(1)
	if err != nil {
		return nil, err
	}

	var token LRToken
	token, err = lexer.maybeTokenizeKeywordOrSymbol()
	if token != nil || err != nil {
		return token, err
	}
	token, err = lexer.maybeTokenizeCharacter()
	if token != nil || err != nil {
		return token, err
	}

	token, err = lexer.maybeTokenizeIdentifier()
	if token != nil || err != nil {
		return token, err
	}

	token, err = lexer.maybeTokenizeSectionContent()
	if token != nil || err != nil {
		return token, err
	}

	return nil, fmt.Errorf("Unexpected character at %s", lexer.reader.Location)
}

func (lexer *rawLexer) stripLeadingWhitespacesAndComments() error {
	modified := true
	for modified {
		modified = false

		bytes, err := lexer.reader.Peek(1)
		if err != nil {
			return err
		}

		_, ok := whitespaces[string(bytes)]
		if ok {
			_, err = lexer.reader.ReadByte()
			if err != nil {
				panic(err) // should never happen
			}

			modified = true
			continue
		}

		bytes, _ = lexer.reader.Peek(2)

		if string(bytes) == "//" {
			for {
				char, err := lexer.reader.ReadByte()
				if err != nil {
					return err
				}

				if char == '\n' {
					break
				}
			}

			modified = true
			continue
		}

		if string(bytes) == "/*" {
			n, err := lexer.reader.Read(bytes)
			if n != 2 || err != nil {
				panic(err) // should never happen
			}

			for {
				bytes, err = lexer.reader.Peek(2)
				if err != nil {
					return err
				}

				if string(bytes) == "*/" {
					n, err := lexer.reader.Read(bytes)
					if n != 2 || err != nil {
						panic(err) // should never happen
					}

					break
				}

				_, err = lexer.reader.ReadByte()
				if err != nil {
					panic(err)
				}
			}

			modified = true
		}
	}

	return nil
}

func (lexer *rawLexer) maybeTokenizeKeywordOrSymbol() (LRToken, error) {
	for str, ttype := range keywordsAndSymbols {
		bytes, _ := lexer.reader.Peek(len(str))

		if string(bytes) == str {
			token := &LRGenericSymbol{ttype, LRLocation(lexer.reader.Location)}

			n, err := lexer.reader.Read(bytes)
			if len(bytes) != n || err != nil {
				panic(err) // should never happen
			}

			return token, nil
		}
	}

	return nil, nil
}

func (lexer *rawLexer) maybeTokenizeCharacter() (LRToken, error) {
	bytes, loc, err := parseutil.MaybeTokenizeCharacter(lexer.reader)
	if err != nil {
		return nil, err
	}

	if bytes == nil {
		return nil, nil
	}

	return &Token{
		LRLocation: LRLocation(loc),
		LRSymbolId: LRCharacterToken,
		Value:      string(bytes),
	}, nil
}

func (lexer *rawLexer) maybeTokenizeIdentifier() (LRToken, error) {
    bytes, loc, err := parseutil.MaybeTokenizeIdentifier(lexer.reader)
    if err != nil {
        return nil, err
    }

    if bytes == nil {
        return nil, nil
    }

    return &Token{
		LRLocation: LRLocation(loc),
		LRSymbolId: LRIdentifierToken,
		Value:      string(bytes),
    },nil
}

func (lexer *rawLexer) maybeTokenizeSectionContent() (LRToken, error) {
	peek, _ := lexer.reader.Peek(1)
	if string(peek) != "{" {
		return nil, nil
	}

	token := &Token{
		LRLocation: LRLocation(lexer.reader.Location),
		LRSymbolId: LRSectionContentToken,
		Value:      "",
	}

	n, err := lexer.reader.Read(peek)
	if n != 1 || err != nil {
		panic(err) // should never happen
	}

	buffer := bytes.NewBuffer(nil)

	for {
		peek, err = lexer.reader.Peek(3)
		if err != nil {
			return nil, err
		}

		if string(peek) == "}%%" {
			n, err := lexer.reader.Read(peek)
			if n != 3 || err != nil {
				panic(err) // should never happen
			}

			break
		}

		char, err := lexer.reader.ReadByte()
		if err != nil {
			panic(err)
		}

		buffer.WriteByte(char)
	}

	token.Value = string(buffer.Bytes())
	return token, nil
}

// This merges LRIdentifierSymbol Arrow token pairs into a single RULE_DEF token and
// LRIdentifierSymbol Colon token pairs into a single LABEL token.
type tokenPairLexer struct {
	base *rawLexer

	nextToken LRToken
	nextErr   error
}

func (lexer *tokenPairLexer) Next() (LRToken, error) {
	if lexer.nextErr != nil {
		err := lexer.nextErr
		lexer.nextErr = nil

		return nil, err
	}

	curr := lexer.nextToken
	lexer.nextToken = nil

	var err error
	if curr == nil {
		curr, err = lexer.base.Next()
		if err != nil {
			return nil, err
		}
	}

	if curr.Id() != LRIdentifierToken {
		return curr, nil
	}

	next, err := lexer.base.Next()
	if err != nil {
		lexer.nextErr = err
		return curr, nil
	}

	if next.Id() == Arrow {
		curr.(*Token).LRSymbolId = LRRuleDefToken
		return curr, nil
	}

	if next.Id() == ':' {
		curr.(*Token).LRSymbolId = LRLabelToken
		return curr, nil
	}

	lexer.nextToken = next
	return curr, nil
}

func (lexer *tokenPairLexer) CurrentLocation() LRLocation {
	if lexer.nextToken != nil {
		return lexer.nextToken.Loc()
	}

	return LRLocation(lexer.base.reader.Location)
}

func NewLexer(filename string, reader io.Reader) LRLexer {
	return &tokenPairLexer{
		base:      newRawLexer(filename, reader),
		nextToken: nil,
		nextErr:   nil,
	}
}
