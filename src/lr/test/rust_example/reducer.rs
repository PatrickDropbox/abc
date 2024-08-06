use std::error::Error;

use super::node::{Binary, Block, Err, Expr, Id};
use super::parser;
use super::parser::{Symbol, SymbolData, SymbolKind};

pub struct Reducer {}

impl Reducer {
    pub fn new() -> Self {
        Self {}
    }
}

impl parser::Reducer for Reducer {
    fn add_to_expr_list(
        &self,
        mut expr_list_: Vec<Expr>,
        expr_: Expr,
    ) -> Result<Vec<Expr>, Box<dyn Error>> {
        expr_list_.push(expr_);
        Ok(expr_list_)
    }

    fn nil_to_expr_list(&self) -> Result<Vec<Expr>, Box<dyn Error>> {
        Ok(vec![])
    }

    fn id_to_atom(&self, id_: Id) -> Result<Expr, Box<dyn Error>> {
        Ok(Expr::Id(id_))
    }

    fn error_to_atom(&self, error_: Err) -> Result<Expr, Box<dyn Error>> {
        Ok(Expr::Err(error_))
    }

    fn block_to_atom(&self, block_: Block) -> Result<Expr, Box<dyn Error>> {
        Ok(Expr::Block(block_))
    }

    fn atom_to_expr(&self, atom_: Expr) -> Result<Expr, Box<dyn Error>> {
        Ok(atom_)
    }

    fn binary_to_expr(
        &self,
        expr_: Expr,
        op_: Symbol,
        atom_: Expr,
    ) -> Result<Expr, Box<dyn Error>> {
        if let Symbol {
            data: SymbolData::Any(token),
            ..
        } = op_
        {
            if let Symbol {
                kind: SymbolKind::AsciiCharToken(op_char),
                ..
            } = *token
            {
                return Ok(Expr::Binary(Binary {
                    rhs: Box::new(expr_),
                    op: op_char,
                    lhs: Box::new(atom_),
                }));
            } else {
                panic!("Unexpected binary op: {:?}", token)
            }
        } else {
            panic!("Unexpected binary op: {:?}", op_)
        }
    }

    fn plus_to_op(&self, char_: Symbol) -> Result<Symbol, Box<dyn Error>> {
        Ok(char_)
    }

    fn minus_to_op(&self, char_: Symbol) -> Result<Symbol, Box<dyn Error>> {
        Ok(char_)
    }

    fn to_block(
        &self,
        _: Symbol,
        expr_list_: Vec<Expr>,
        _: Symbol,
    ) -> Result<Block, Box<dyn Error>> {
        Ok(Block { list: expr_list_ })
    }
}
