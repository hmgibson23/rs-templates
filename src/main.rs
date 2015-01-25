

mod lexer {
    use std::io::{BufferedReader, File};
    use std::char::CharExt;

    /* The lexer that lexes */
    enum Token {
        tok_value(String),
        Missing
    }

    fn getToken(input: &str) -> Token {

        let mut lastChar = String::new();
        let mut token = Token::Missing;

        let mut tokenizing = 0; //state machine for tokenizing

        for c in input.as_slice().chars() {

            if(CharExt::is_whitespace(c)) {
                continue;
            }

            match lastChar.find_str("}}") {
                Some(uint) => { token = getExpr(lastChar) },
                Noe => continue

            }

            if(tokenizing == 1) {
                // keep going until we get to the end of the token stream
                lastChar.push(c);
                continue;
            }

            if(c == '{') {
                // assume it's an expression and tokenize it
                lastChar.push(c);
                // we are now tokenizing
                tokenizing = 1;
                continue;
            }
        }
        token
    }

    fn getExpr(input: String) -> Token {
        Token::Missing
    }

}

fn main() {
    println!("Hello, world!");
}
