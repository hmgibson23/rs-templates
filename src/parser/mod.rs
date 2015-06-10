use lexer::Token;
use lexer::tokenize_file;

/* take the stream of tokens from the lexer and parse them */

// source is the original value -> retained to do a string replace
// value is the parsed value
#[derive(Clone, PartialEq)]
pub struct Value {
    pub source: String,
    pub value: String
}

pub fn parse_token(expr: &Token) -> Option<Value> {
    match expr {
        &Token::Expr(ref input) => {
            Some(prepare_expr(input))
        },
        &Token::Missing => None
    }
}

fn prepare_expr(input: &String) -> Value {
    // remove surrounding {} and whitespace
    // and you've got the value
    let raw = input.clone().replace("{", "").replace("}", "");
    Value {source: input.clone(), value: raw}
}

pub fn parse_file(file: &str) -> Vec<Option<Value>> {
    let tokens = tokenize_file(file);
    tokens.iter().map(parse_token).collect::<Vec<_>>()
}
