use super::lexer::{BasicLexer, ScopedLexer};
use super::node::Node;
use super::parser;
use super::parser::{Lexer, Symbol, SymbolKind};
use super::reducer::Reducer;

fn print_tree(input: &str) {
    println!("======");
    println!("Input: {}", input);
    println!("======");
    println!("Basic:");

    let mut lexer = BasicLexer::new("source.txt".to_string(), input.to_string());

    let reducer = Reducer::new();

    let result = parser::parse_expr_list(&mut lexer, &reducer);

    match result {
        Ok(list) => {
            for item in list.iter() {
                println!("{}", item.to_string())
            }
        }
        Err(e) => println!("parse error: {}", e),
    };

    println!("------");
    println!("Scoped:");

    let mut base = BasicLexer::new("source.txt".to_string(), input.to_string());
    let mut lexer = ScopedLexer::new(&mut base, &reducer);

    let result = parser::parse_expr_list(&mut lexer, &reducer);

    match result {
        Ok(list) => {
            for item in list.iter() {
                println!("{}", item.to_string())
            }
        }
        Err(e) => println!("parse error: {}", e),
    };
}

#[test]
fn test_main() {
    print_tree("a + b");
    print_tree("a + b + c - d");
    print_tree("{ }");
    print_tree("a + { }");
    print_tree("a { }");
    print_tree("{ } { }");
    print_tree("{ } + { }");
    print_tree("{ { } }");
    print_tree("{ { 1 } }");
    print_tree("{ } a b + c d { 1 2 3 4 + { 5 + 6 } + 7 }");
    print_tree("{ } } a b + c d { 1 2 3 4 + { { a + + b }  + 6 } + 7 } f { a + } b {");
    print_tree("{ } a b + c d { 1 2 3 4 + { { a + + b }  + 6 } + 7 } f { a + } b {");
}
