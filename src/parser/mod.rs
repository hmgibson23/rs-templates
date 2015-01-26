mod parser;

use std::collections::HashMap;

/* take the stream of tokens from the lexer and parser them */



/* eventually we'll need an AST for this */
enum Expr {
    /* for now this only contains simple values but it should be extended for
     * loops etc.
     */
    Value(HashMap), // string -  complete lexed string -> string - parsed value inside
    None
}


pub fn map_tokens(tokens: Vec<Token>) -> Vec<Expr> {
    tokens.map_in_place()
}

fn parse_token(expr: token) -> Expr {
    // remove surrounding {} and whitespace
    // and you've got the value
    let (input, _) = expr;
    let raw = input.clone();
    raw.remove(0);
    raw.remove(1);

    Expr::Value(temp, input)
}
