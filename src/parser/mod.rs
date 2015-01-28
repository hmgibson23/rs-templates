use std::collections::HashMap;
use lexer::Token;

mod parser;


/* take the stream of tokens from the lexer and parse them */

/* eventually we'll need an AST for this */
enum Expr {
    /* for now this only contains simple values but it should be extended for
     * loops etc.
     */
    // string -  complete lexed string -> string - parsed value inside
    Value(HashMap<String, String>),
    Empty
}


pub fn map_tokens(tokens: Vec<Token>) -> Vec<Expr> {
    tokens.map_in_place(|&: x: Token| -> Expr { parse_token(x)})
}

fn parse_token(expr: Token) -> Expr {

    let val = match expr {
        Token::Expr(input) => {
            let f = prepare_expr(input);
            Expr::Value(f)
        },
        Token::Missing => Expr::Empty
    };
    val
}

fn prepare_expr(input: String) -> HashMap<String, String> {
    // remove surrounding {} and whitespace
    // and you've got the value
    let mut raw = input.clone();
    raw.remove(0);
    raw.remove(1);
    raw.pop();
    raw.pop();
    let mut h = HashMap::new();
    h.insert(raw, input);
    h
}
