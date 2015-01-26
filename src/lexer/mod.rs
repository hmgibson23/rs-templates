use std::io::{BufferedReader, File};
use std::char::CharExt;

mod lexer;

/* The lexer that lexes */
enum Token {
    Expr(String),
    Missing
}

impl Token {
    pub fn stringify(self) -> String {
        match self {
            Token::Expr(val) => val,
            Missing => format!("No token.")
        }
    }
}

pub fn get_token(input: &str) -> Token {

    let mut last_char = String::new();
    let mut token = Token::Missing;

    let mut tokenizing = 0; //state machine for tokenizing

    for c in input.as_slice().chars() {

        if CharExt::is_whitespace(c) {
            continue;
        }

        if tokenizing == 1 {
            // keep going until we get to the end of the token stream
            last_char.push(c);

            // check if we're at the end yet
            if complete_token(last_char.clone()) {
                token = get_expr(last_char.clone());
                break;
            }

            continue;
        }


        if c == '{' {
            // assume it's an expression and tokenize it
            last_char.push(c);
            // we are now tokenizing
            tokenizing = 1;
            continue;
        }

        }
        token
    }

    fn complete_token(input: String) -> bool {
        match input.find_str("}}") {
            Some(uint) => true,
            None => false
        }
    }

    fn get_expr(input: String) -> Token {
        Token::Expr(input)
    }
