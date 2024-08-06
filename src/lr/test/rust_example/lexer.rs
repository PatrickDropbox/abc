use std::collections::VecDeque;
use std::error;

use super::node::{Err, Id, Node};
use super::parser;
use super::parser::{Error, Lexer, Reducer, Symbol, SymbolData, SymbolKind};

pub struct BasicLexer {
    file_name: String,
    content: String,

    tokens: VecDeque<Symbol>,
}

impl BasicLexer {
    pub fn new(file_name: String, content: String) -> Self {
        let mut tokens = VecDeque::new();

        for (column, character) in content.chars().enumerate() {
            if character == ' ' {
                continue;
            }

            tokens.push_back(match character {
                '{' | '}' | '+' | '-' => Symbol {
                    kind: SymbolKind::AsciiCharToken(character),
                    data: SymbolData::Nil,
                },
                _ => Symbol {
                    kind: SymbolKind::IdToken,
                    // XXX: not utf8 safe!!! should use the char's byte size instead of +1
                    data: SymbolData::Ident(Id {
                        value: content[column..column + 1].to_string(),
                    }),
                },
            })
        }

        Self {
            file_name: file_name,
            content: content,
            tokens: tokens,
        }
    }
}

impl Lexer for BasicLexer {
    fn next(&mut self) -> Result<Symbol, Box<dyn error::Error>> {
        match self.tokens.pop_front() {
            Some(symbol) => Ok(symbol),
            _ => Ok(Symbol {
                kind: SymbolKind::EofToken,
                data: SymbolData::Nil,
            }),
        }
    }
}

pub struct BufferedLexer<'a, L: Lexer> {
    base: &'a mut L,
    buffer_stack: Vec<Result<Symbol, Box<dyn error::Error>>>,
}

impl<'a, L: Lexer> BufferedLexer<'a, L> {
    fn next(&mut self) -> (Result<Symbol, Box<dyn error::Error>>, bool) {
        if let Some(s) = self.buffer_stack.pop() {
            return (s, true);
        }

        (self.base.next(), false)
    }

    fn push_front(&mut self, prev: Result<Symbol, Box<dyn error::Error>>) {
        self.buffer_stack.push(prev);
    }
}

pub struct ScopedLexer<'a, L: Lexer, R: Reducer> {
    base: BufferedLexer<'a, L>,
    reducer: &'a R,

    depth: i32,
    scope_ended: bool,
}

impl<'a, L: Lexer, R: Reducer> ScopedLexer<'a, L, R> {
    pub fn new(base: &'a mut L, reducer: &'a R) -> Self {
        Self {
            base: BufferedLexer {
                base: base,
                buffer_stack: vec![],
            },
            reducer: reducer,
            depth: 0,
            scope_ended: false,
        }
    }
}

impl<'a, L: Lexer, R: Reducer> Lexer for ScopedLexer<'a, L, R> {
    fn next(&mut self) -> Result<Symbol, Box<dyn error::Error>> {
        if self.scope_ended {
            self.scope_ended = false;
            return Ok(Symbol {
                kind: SymbolKind::EofToken,
                data: SymbolData::Nil,
            });
        }

        let (result, buffered) = self.base.next();
        if buffered {
            return result;
        }

        let symbol = result?;
        match symbol.kind.clone() {
            SymbolKind::AsciiCharToken('{') => {
                let block_depth = self.depth;
                self.depth += 1;

                self.base.push_front(Ok(symbol));

                let result = parser::parse_block(self, self.reducer);
                self.scope_ended = false;

                match result {
                    Ok(block) => {
                        return Ok(Symbol {
                            kind: SymbolKind::BlockType,
                            data: SymbolData::Block(block),
                        })
                    }
                    Err(e) => {
                        while self.depth > block_depth {
                            let (next, _) = self.base.next();
                            if let Ok(next_symbol) = next {
                                match next_symbol.kind.clone() {
                                    SymbolKind::AsciiCharToken('{') => self.depth += 1,
                                    SymbolKind::AsciiCharToken('}') => self.depth -= 1,
                                    SymbolKind::EofToken => {
                                        self.base.push_front(Ok(next_symbol));
                                        break;
                                    }
                                    _ => (),
                                }
                            } else {
                                self.base.push_front(next);
                                break;
                            }
                        }

                        return Ok(Symbol {
                            kind: SymbolKind::ErrorToken,
                            data: SymbolData::Err(Err { error: e }),
                        });
                    }
                }
            }
            SymbolKind::AsciiCharToken('}') => {
                assert!(self.depth >= 0);
                if self.depth == 0 {
                    let e = Box::new(Error::new(
                        "Unexpected RBRACE. No matching LBRACE".to_string(),
                    ));
                    return Ok(Symbol {
                        kind: SymbolKind::ErrorToken,
                        data: SymbolData::Err(Err { error: e }),
                    });
                }

                self.depth -= 1;
                self.scope_ended = true;
                return Ok(symbol);
            }
            _ => return Ok(symbol),
        }
    }
}
