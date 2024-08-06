use std::error::Error;

fn indent(i: i32) -> String {
    let mut result = "".to_string();
    for _ in 0..i {
        result += "    ";
    }

    result
}

pub trait Node {
    fn to_string(&self) -> String {
        self.to_indent_string(0)
    }

    fn to_indent_string(&self, i: i32) -> String;
}

#[derive(Debug)]
pub struct Id {
    pub value: String,
}

impl Node for Id {
    fn to_indent_string(&self, i: i32) -> String {
        format!("{}ID={}", indent(i), self.value)
    }
}

#[derive(Debug)]
pub struct Err {
    pub error: Box<dyn Error>,
}

impl Node for Err {
    fn to_indent_string(&self, i: i32) -> String {
        format!("{}ERROR={}", indent(i), self.error)
    }
}

#[derive(Debug)]
pub struct Binary {
    pub rhs: Box<Expr>,
    pub op: char,
    pub lhs: Box<Expr>,
}

impl Node for Binary {
    fn to_indent_string(&self, i: i32) -> String {
        let mut result = format!("{}BINARY=\n", indent(i));
        result.push_str(&self.rhs.to_indent_string(i + 1));
        result.push_str(&format!("\n{}OP={}\n", indent(i + 1), self.op));
        result.push_str(&self.lhs.to_indent_string(i + 1));
        result
    }
}

#[derive(Debug)]
pub struct Block {
    pub list: Vec<Expr>,
}

impl Node for Block {
    fn to_indent_string(&self, i: i32) -> String {
        let indent_str = indent(i);
        let mut result = format!("{}BLOCK={{\n", indent_str);

        for expr in &self.list {
            result.push_str(&expr.to_indent_string(i + 1));
            result.push('\n');
        }

        result.push_str(&indent_str);
        result.push('}');
        result
    }
}

#[derive(Debug)]
pub enum Expr {
    Id(Id),
    Err(Err),
    Block(Block),
    Binary(Binary),
}

impl Node for Expr {
    fn to_indent_string(&self, i: i32) -> String {
        match self {
            Expr::Id(id) => id.to_indent_string(i),
            Expr::Err(err) => err.to_indent_string(i),
            Expr::Block(block) => block.to_indent_string(i),
            Expr::Binary(binary) => binary.to_indent_string(i),
        }
    }
}
