// Auto-generated from source: ../piecemeal-parsing/grammar.lr

use std::error;
use std::fmt;

use super::node::{Block, Err, Expr, Id};

#[derive(Debug, Clone)]
pub enum SymbolKind {
    //
    // Token symbols.
    //
    EofToken,
    AsciiCharToken(char),

    IdToken,
    ErrorToken,

    //
    // Type symbols.
    //
    ExprListType,
    AtomType,
    ExprType,
    OpType,
    BlockType,

    //
    // For internal use only.
    //
    _StartParseMarker,
    _WildcardMarker,
}

impl SymbolKind {
    fn to_string(&self) -> String {
        match self {
            SymbolKind::_StartParseMarker => "^",
            SymbolKind::_WildcardMarker => "*",
            SymbolKind::EofToken => "$",

            SymbolKind::AsciiCharToken('+') => "'+'",
            SymbolKind::AsciiCharToken('-') => "'-'",
            SymbolKind::AsciiCharToken('{') => "'{'",
            SymbolKind::AsciiCharToken('}') => "'}'",
            SymbolKind::IdToken => "ID",
            SymbolKind::ErrorToken => "ERROR",

            SymbolKind::ExprListType => "expr_list",
            SymbolKind::AtomType => "atom",
            SymbolKind::ExprType => "expr",
            SymbolKind::OpType => "op",
            SymbolKind::BlockType => "block",
            SymbolKind::AsciiCharToken(c) => panic!("Unexpected token '{}'", c),
        }
        .to_string()
    }
}

#[derive(Debug)]
pub enum SymbolData {
    // Note: %token without value declaration must have Nil data.
    Nil,
    // Note: %type without value declaration must have Any data.
    Any(Box<Symbol>),

    Block(Block),
    Err(Err),
    Expr(Expr),
    ExprList(Vec<Expr>),
    Ident(Id),
}

#[derive(Debug)]
pub struct Symbol {
    pub kind: SymbolKind,
    pub data: SymbolData,
}

impl Symbol {
    pub fn validate(&self) -> Result<(), Box<dyn error::Error>> {
        match self {
            Symbol {
                kind: SymbolKind::EofToken,
                data: _,
            } => (),
            Symbol {
                kind: SymbolKind::AsciiCharToken('+'),
                data: SymbolData::Nil,
            } => (),
            Symbol {
                kind: SymbolKind::AsciiCharToken('-'),
                data: SymbolData::Nil,
            } => (),
            Symbol {
                kind: SymbolKind::AsciiCharToken('{'),
                data: SymbolData::Nil,
            } => (),
            Symbol {
                kind: SymbolKind::AsciiCharToken('}'),
                data: SymbolData::Nil,
            } => (),
            Symbol {
                kind: SymbolKind::IdToken,
                data: SymbolData::Ident(_),
            } => (),
            Symbol {
                kind: SymbolKind::ErrorToken,
                data: SymbolData::Err(_),
            } => (),
            Symbol {
                kind: SymbolKind::ExprListType,
                data: SymbolData::ExprList(_),
            } => (),
            Symbol {
                kind: SymbolKind::AtomType,
                data: SymbolData::Expr(_),
            } => (),
            Symbol {
                kind: SymbolKind::ExprType,
                data: SymbolData::Expr(_),
            } => (),
            Symbol {
                kind: SymbolKind::OpType,
                data: SymbolData::Any(_),
            } => (),
            Symbol {
                kind: SymbolKind::BlockType,
                data: SymbolData::Block(_),
            } => (),
            _ => {
                return Err(Box::new(Error::new(format!(
                    "Unexpected symbol {:?}",
                    self
                ))))
            }
        };

        Ok(())
    }
}

pub trait Lexer {
    fn next(&mut self) -> Result<Symbol, Box<dyn error::Error>>;
}

pub trait Reducer {
    // 18:4: expr_list -> add: ...
    fn add_to_expr_list(
        &self,
        expr_list_: Vec<Expr>,
        expr_: Expr,
    ) -> Result<Vec<Expr>, Box<dyn error::Error>>;
    // 19:4: expr_list -> nil: ...
    fn nil_to_expr_list(&self) -> Result<Vec<Expr>, Box<dyn error::Error>>;

    // 22:4: atom -> id: ...
    fn id_to_atom(&self, id_: Id) -> Result<Expr, Box<dyn error::Error>>;
    // 23:4: atom -> error: ...
    fn error_to_atom(&self, error_: Err) -> Result<Expr, Box<dyn error::Error>>;
    // 24:4: atom -> block: ...
    fn block_to_atom(&self, block_: Block) -> Result<Expr, Box<dyn error::Error>>;

    // 27:4: expr -> atom: ...
    fn atom_to_expr(&self, atom_: Expr) -> Result<Expr, Box<dyn error::Error>>;
    // 28:4: expr -> binary: ...
    fn binary_to_expr(
        &self,
        expr_: Expr,
        op_: Symbol,
        atom_: Expr,
    ) -> Result<Expr, Box<dyn error::Error>>;

    // 31:4: op -> plus: ...
    fn plus_to_op(&self, char_: Symbol) -> Result<Symbol, Box<dyn error::Error>>;
    // 32:4: op -> minus: ...
    fn minus_to_op(&self, char_: Symbol) -> Result<Symbol, Box<dyn error::Error>>;

    // 34:9: block -> ...
    fn to_block(
        &self,
        char_: Symbol,
        expr_list_: Vec<Expr>,
        char_2: Symbol,
    ) -> Result<Block, Box<dyn error::Error>>;
}

#[derive(Debug)]
pub struct Error {
    msg: String,
}

impl Error {
    pub fn new(msg: String) -> Self {
        Self { msg: msg }
    }
}

impl fmt::Display for Error {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(f, "{}", self.msg)
    }
}

impl error::Error for Error {}

pub fn parse_expr_list<L: Lexer, R: Reducer>(
    lexer: &mut L,
    reducer: &R,
) -> Result<Vec<Expr>, Box<dyn error::Error>> {
    let result = parse(lexer, reducer, StateId::State1)?;

    match result {
        SymbolData::ExprList(val) => return Ok(val),
        _ => (),
    }

    panic!("Invalid symbol data type. This should never happen");
}
pub fn parse_block<L: Lexer, R: Reducer>(
    lexer: &mut L,
    reducer: &R,
) -> Result<Block, Box<dyn error::Error>> {
    let result = parse(lexer, reducer, StateId::State2)?;

    match result {
        SymbolData::Block(val) => return Ok(val),
        _ => (),
    }

    panic!("Invalid symbol data type. This should never happen");
}

// ==============================
// Parser internal implementation
// ==============================

fn parse<L: Lexer, R: Reducer>(
    lexer: &mut L,
    reducer: &R,
    start_state: StateId,
) -> Result<SymbolData, Box<dyn error::Error>> {
    let mut state_stack = vec![ParseStackFrame {
        state_id: start_state,
        symbol: Symbol {
            kind: SymbolKind::_StartParseMarker,
            data: SymbolData::Nil,
        },
    }];

    let mut symbol_stack = SymbolStack::new(lexer);

    loop {
        let next_symbol_kind = symbol_stack.peek()?;

        let current_state_id = state_stack[state_stack.len() - 1].state_id.clone();

        match lookup_action(current_state_id, next_symbol_kind) {
            Action::Goto(next_state_id) => state_stack.push(ParseStackFrame {
                state_id: next_state_id,
                symbol: symbol_stack.pop(),
            }),
            Action::Reduce(reduce_kind) => {
                symbol_stack.push(reduce_symbol(reduce_kind, reducer, &mut state_stack)?)?;
            }
            Action::Accept => {
                assert_eq!(state_stack.len(), 2, "This should never happen");
                return Ok(state_stack.pop().unwrap().symbol.data);
            }
            Action::Error => {
                return Err(new_syntax_error(
                    symbol_stack.pop(),
                    state_stack.pop().unwrap().state_id,
                ))
            }
        }
    }
}

fn reduce_symbol<R: Reducer>(
    reduce_kind: ReduceKind,
    reducer: &R,
    state_stack: &mut Vec<ParseStackFrame>,
) -> Result<Symbol, Box<dyn error::Error>> {
    let reduced = match reduce_kind {
        ReduceKind::AddToExprList => {
            let symbol1 = state_stack.pop().unwrap().symbol;
            let arg1 = match symbol1 {
                Symbol {
                    data: SymbolData::Expr(val),
                    ..
                } => val,
                _ => panic!("Failed to extract argument.  This should never happen"),
            };

            let symbol0 = state_stack.pop().unwrap().symbol;
            let arg0 = match symbol0 {
                Symbol {
                    data: SymbolData::ExprList(val),
                    ..
                } => val,
                _ => panic!("Failed to extract argument.  This should never happen"),
            };

            let result = reducer.add_to_expr_list(arg0, arg1)?;
            Symbol {
                kind: SymbolKind::ExprListType,
                data: SymbolData::ExprList(result),
            }
        }
        ReduceKind::NilToExprList => {
            let result = reducer.nil_to_expr_list()?;
            Symbol {
                kind: SymbolKind::ExprListType,
                data: SymbolData::ExprList(result),
            }
        }
        ReduceKind::IdToAtom => {
            let symbol0 = state_stack.pop().unwrap().symbol;
            let arg0 = match symbol0 {
                Symbol {
                    data: SymbolData::Ident(val),
                    ..
                } => val,
                _ => panic!("Failed to extract argument.  This should never happen"),
            };

            let result = reducer.id_to_atom(arg0)?;
            Symbol {
                kind: SymbolKind::AtomType,
                data: SymbolData::Expr(result),
            }
        }
        ReduceKind::ErrorToAtom => {
            let symbol0 = state_stack.pop().unwrap().symbol;
            let arg0 = match symbol0 {
                Symbol {
                    data: SymbolData::Err(val),
                    ..
                } => val,
                _ => panic!("Failed to extract argument.  This should never happen"),
            };

            let result = reducer.error_to_atom(arg0)?;
            Symbol {
                kind: SymbolKind::AtomType,
                data: SymbolData::Expr(result),
            }
        }
        ReduceKind::BlockToAtom => {
            let symbol0 = state_stack.pop().unwrap().symbol;
            let arg0 = match symbol0 {
                Symbol {
                    data: SymbolData::Block(val),
                    ..
                } => val,
                _ => panic!("Failed to extract argument.  This should never happen"),
            };

            let result = reducer.block_to_atom(arg0)?;
            Symbol {
                kind: SymbolKind::AtomType,
                data: SymbolData::Expr(result),
            }
        }
        ReduceKind::AtomToExpr => {
            let symbol0 = state_stack.pop().unwrap().symbol;
            let arg0 = match symbol0 {
                Symbol {
                    data: SymbolData::Expr(val),
                    ..
                } => val,
                _ => panic!("Failed to extract argument.  This should never happen"),
            };

            let result = reducer.atom_to_expr(arg0)?;
            Symbol {
                kind: SymbolKind::ExprType,
                data: SymbolData::Expr(result),
            }
        }
        ReduceKind::BinaryToExpr => {
            let symbol2 = state_stack.pop().unwrap().symbol;
            let arg2 = match symbol2 {
                Symbol {
                    data: SymbolData::Expr(val),
                    ..
                } => val,
                _ => panic!("Failed to extract argument.  This should never happen"),
            };

            let symbol1 = state_stack.pop().unwrap().symbol;
            let arg1 = symbol1;

            let symbol0 = state_stack.pop().unwrap().symbol;
            let arg0 = match symbol0 {
                Symbol {
                    data: SymbolData::Expr(val),
                    ..
                } => val,
                _ => panic!("Failed to extract argument.  This should never happen"),
            };

            let result = reducer.binary_to_expr(arg0, arg1, arg2)?;
            Symbol {
                kind: SymbolKind::ExprType,
                data: SymbolData::Expr(result),
            }
        }
        ReduceKind::PlusToOp => {
            let symbol0 = state_stack.pop().unwrap().symbol;
            let arg0 = symbol0;

            let result = reducer.plus_to_op(arg0)?;
            Symbol {
                kind: SymbolKind::OpType,
                data: SymbolData::Any(Box::new(result)),
            }
        }
        ReduceKind::MinusToOp => {
            let symbol0 = state_stack.pop().unwrap().symbol;
            let arg0 = symbol0;

            let result = reducer.minus_to_op(arg0)?;
            Symbol {
                kind: SymbolKind::OpType,
                data: SymbolData::Any(Box::new(result)),
            }
        }
        ReduceKind::ToBlock => {
            let symbol2 = state_stack.pop().unwrap().symbol;
            let arg2 = symbol2;

            let symbol1 = state_stack.pop().unwrap().symbol;
            let arg1 = match symbol1 {
                Symbol {
                    data: SymbolData::ExprList(val),
                    ..
                } => val,
                _ => panic!("Failed to extract argument.  This should never happen"),
            };

            let symbol0 = state_stack.pop().unwrap().symbol;
            let arg0 = symbol0;

            let result = reducer.to_block(arg0, arg1, arg2)?;
            Symbol {
                kind: SymbolKind::BlockType,
                data: SymbolData::Block(result),
            }
        }
    };

    Ok(reduced)
}

struct SymbolStack<'a, L: Lexer> {
    lexer: &'a mut L,
    top: Vec<Symbol>,
}

impl<'a, L: Lexer> SymbolStack<'a, L> {
    fn new(lexer: &'a mut L) -> Self {
        Self {
            lexer: lexer,
            top: Vec::new(),
        }
    }

    fn peek(&mut self) -> Result<SymbolKind, Box<dyn error::Error>> {
        if self.top.is_empty() {
            let symbol = self.lexer.next()?;
            let _ = symbol.validate()?;
            self.top.push(symbol);
        }

        Ok(self.top[self.top.len() - 1].kind.clone())
    }

    fn push(&mut self, symbol: Symbol) -> Result<(), Box<dyn error::Error>> {
        symbol.validate()?;
        self.top.push(symbol);
        Ok(())
    }

    fn pop(&mut self) -> Symbol {
        if self.top.is_empty() {
            panic!("Cannot pop an empty symbol stack. This should never happen");
        }

        self.top.pop().unwrap()
    }
}

#[derive(Clone, Debug)]
enum StateId {
    State1,
    State2,
    State3,
    State4,
    State5,
    State6,
    State7,
    State8,
    State9,
    State10,
    State11,
    State12,
    State13,
    State14,
    State15,
    State16,
}

struct ParseStackFrame {
    state_id: StateId,
    symbol: Symbol,
}

#[derive(Clone, Debug)]
enum ReduceKind {
    AddToExprList,
    NilToExprList,
    IdToAtom,
    ErrorToAtom,
    BlockToAtom,
    AtomToExpr,
    BinaryToExpr,
    PlusToOp,
    MinusToOp,
    ToBlock,
}

#[derive(Clone, Debug)]
enum Action {
    Goto(StateId),
    Reduce(ReduceKind),
    Accept,
    Error,
}

fn lookup_action(current_state: StateId, next_symbol: SymbolKind) -> Action {
    match current_state {
        StateId::State1 => match next_symbol {
            SymbolKind::ExprListType => Action::Goto(StateId::State3),
            _ => Action::Reduce(ReduceKind::NilToExprList),
        },
        StateId::State2 => match next_symbol {
            SymbolKind::AsciiCharToken('{') => Action::Goto(StateId::State5),
            SymbolKind::BlockType => Action::Goto(StateId::State4),
            _ => Action::Error,
        },
        StateId::State3 => match next_symbol {
            SymbolKind::EofToken => Action::Accept,
            SymbolKind::AsciiCharToken('{') => Action::Goto(StateId::State5),
            SymbolKind::IdToken => Action::Goto(StateId::State7),
            SymbolKind::ErrorToken => Action::Goto(StateId::State6),
            SymbolKind::AtomType => Action::Goto(StateId::State8),
            SymbolKind::ExprType => Action::Goto(StateId::State10),
            SymbolKind::BlockType => Action::Goto(StateId::State9),
            _ => Action::Error,
        },
        StateId::State4 => match next_symbol {
            SymbolKind::EofToken => Action::Accept,
            _ => Action::Error,
        },
        StateId::State5 => match next_symbol {
            SymbolKind::ExprListType => Action::Goto(StateId::State11),
            _ => Action::Reduce(ReduceKind::NilToExprList),
        },
        StateId::State6 => match next_symbol {
            _ => Action::Reduce(ReduceKind::ErrorToAtom),
        },
        StateId::State7 => match next_symbol {
            _ => Action::Reduce(ReduceKind::IdToAtom),
        },
        StateId::State8 => match next_symbol {
            _ => Action::Reduce(ReduceKind::AtomToExpr),
        },
        StateId::State9 => match next_symbol {
            _ => Action::Reduce(ReduceKind::BlockToAtom),
        },
        StateId::State10 => match next_symbol {
            SymbolKind::AsciiCharToken('+') => Action::Goto(StateId::State12),
            SymbolKind::AsciiCharToken('-') => Action::Goto(StateId::State13),
            SymbolKind::OpType => Action::Goto(StateId::State14),
            _ => Action::Reduce(ReduceKind::AddToExprList),
        },
        StateId::State11 => match next_symbol {
            SymbolKind::AsciiCharToken('{') => Action::Goto(StateId::State5),
            SymbolKind::AsciiCharToken('}') => Action::Goto(StateId::State15),
            SymbolKind::IdToken => Action::Goto(StateId::State7),
            SymbolKind::ErrorToken => Action::Goto(StateId::State6),
            SymbolKind::AtomType => Action::Goto(StateId::State8),
            SymbolKind::ExprType => Action::Goto(StateId::State10),
            SymbolKind::BlockType => Action::Goto(StateId::State9),
            _ => Action::Error,
        },
        StateId::State12 => match next_symbol {
            _ => Action::Reduce(ReduceKind::PlusToOp),
        },
        StateId::State13 => match next_symbol {
            _ => Action::Reduce(ReduceKind::MinusToOp),
        },
        StateId::State14 => match next_symbol {
            SymbolKind::AsciiCharToken('{') => Action::Goto(StateId::State5),
            SymbolKind::IdToken => Action::Goto(StateId::State7),
            SymbolKind::ErrorToken => Action::Goto(StateId::State6),
            SymbolKind::AtomType => Action::Goto(StateId::State16),
            SymbolKind::BlockType => Action::Goto(StateId::State9),
            _ => Action::Error,
        },
        StateId::State15 => match next_symbol {
            SymbolKind::EofToken => Action::Reduce(ReduceKind::ToBlock),
            _ => Action::Error,
        },
        StateId::State16 => match next_symbol {
            _ => Action::Reduce(ReduceKind::BinaryToExpr),
        },
    }
}
fn new_syntax_error(next_symbol: Symbol, current_state_id: StateId) -> Box<dyn error::Error> {
    let expected_terminals = match current_state_id {
        StateId::State1 => "*, expr_list",
        StateId::State2 => "'{', block",
        StateId::State3 => "$, '{', ID, ERROR, atom, expr, block",
        StateId::State4 => "$",
        StateId::State5 => "*, expr_list",
        StateId::State6 => "*",
        StateId::State7 => "*",
        StateId::State8 => "*",
        StateId::State9 => "*",
        StateId::State10 => "*, '+', '-', op",
        StateId::State11 => "'{', '}', ID, ERROR, atom, expr, block",
        StateId::State12 => "*",
        StateId::State13 => "*",
        StateId::State14 => "'{', ID, ERROR, atom, block",
        StateId::State15 => "$",
        StateId::State16 => "*",
    };

    Box::new(Error::new(format!(
        "Syntax error: unexpected symbol {}. Expecting [{}]",
        next_symbol.kind.to_string(),
        expected_terminals
    )))
}

/*
Parser Debug States:
  State 1:
    Kernel Items:
      #accept: ^.expr_list
    Reduce:
      * -> [expr_list]
    Goto:
      expr_list -> State 3

  State 2:
    Kernel Items:
      #accept: ^.block
    Reduce:
      (nil)
    Goto:
      '{' -> State 5
      block -> State 4

  State 3:
    Kernel Items:
      #accept: ^ expr_list., $
      expr_list: expr_list.expr
    Reduce:
      $ -> [#accept]
    Goto:
      '{' -> State 5
      ID -> State 7
      ERROR -> State 6
      atom -> State 8
      expr -> State 10
      block -> State 9

  State 4:
    Kernel Items:
      #accept: ^ block., $
    Reduce:
      $ -> [#accept]
    Goto:
      (nil)

  State 5:
    Kernel Items:
      block: '{'.expr_list '}'
    Reduce:
      * -> [expr_list]
    Goto:
      expr_list -> State 11

  State 6:
    Kernel Items:
      atom: ERROR., *
    Reduce:
      * -> [atom]
    Goto:
      (nil)

  State 7:
    Kernel Items:
      atom: ID., *
    Reduce:
      * -> [atom]
    Goto:
      (nil)

  State 8:
    Kernel Items:
      expr: atom., *
    Reduce:
      * -> [expr]
    Goto:
      (nil)

  State 9:
    Kernel Items:
      atom: block., *
    Reduce:
      * -> [atom]
    Goto:
      (nil)

  State 10:
    Kernel Items:
      expr_list: expr_list expr., *
      expr: expr.op atom
    Reduce:
      * -> [expr_list]
    Goto:
      '+' -> State 12
      '-' -> State 13
      op -> State 14

  State 11:
    Kernel Items:
      expr_list: expr_list.expr
      block: '{' expr_list.'}'
    Reduce:
      (nil)
    Goto:
      '{' -> State 5
      '}' -> State 15
      ID -> State 7
      ERROR -> State 6
      atom -> State 8
      expr -> State 10
      block -> State 9

  State 12:
    Kernel Items:
      op: '+'., *
    Reduce:
      * -> [op]
    Goto:
      (nil)

  State 13:
    Kernel Items:
      op: '-'., *
    Reduce:
      * -> [op]
    Goto:
      (nil)

  State 14:
    Kernel Items:
      expr: expr op.atom
    Reduce:
      (nil)
    Goto:
      '{' -> State 5
      ID -> State 7
      ERROR -> State 6
      atom -> State 16
      block -> State 9

  State 15:
    Kernel Items:
      block: '{' expr_list '}'., $
    Reduce:
      $ -> [block]
    Goto:
      (nil)

  State 16:
    Kernel Items:
      expr: expr op atom., *
    Reduce:
      * -> [expr]
    Goto:
      (nil)

Number of states: 16
Number of shift actions: 25
Number of reduce actions: 13
Number of shift/reduce conflicts: 0
Number of reduce/reduce conflicts: 0
]
*/
