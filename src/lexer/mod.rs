use std::char;
use std::fs::File;
use std::io::Read;


/* The lexer that lexes */
#[derive(Clone, PartialEq)]
pub enum Token {
    Expr(String),
    Missing
}

impl Token {
    pub fn stringify(self) -> String {
        match self {
            Token::Expr(val) => val,
            Token::Missing => format!("No token.")
        }
    }
}

pub fn get_tokens(input: &str) -> Vec<Token> {

    let mut last_char = String::new();

    let mut tokens = Vec::new();


    let mut tokenizing = 0; //state machine for tokenizing

    for c in input.chars() {

        if char::is_whitespace(c) && tokenizing != 1 {
            continue;
        }

        if tokenizing == 1 {
            // keep going until we get to the end of the token stream
            last_char.push(c);

            // check if we're at the end yet
            if complete_token(last_char.clone()) {
                tokens.push(get_expr(last_char.clone()));
                tokenizing = 0;
                last_char = String::new()
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
    tokens
}

fn complete_token(input: String) -> bool {
    match input.find("}}") {
        Some(_) => true,
        None => false
    }
}

fn get_expr(input: String) -> Token {
    Token::Expr(input)
}

pub fn tokenize_file(file: &str) -> Vec<Token> {
    let mut file = File::open(file).unwrap();
    let mut data = String::new();

    file.read_to_string(&mut data).unwrap();
    get_tokens(&data)
}
