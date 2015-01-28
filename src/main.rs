mod lexer;
mod parser;

fn main() {
    let tokens = lexer::get_token("<html><p>{{value}}</p><p>{{another value}}</p></html>");
    let parsed = parser::map_tokens(tokens.clone());


    for t in tokens.map_in_place(|&: x: lexer::Token| -> String { x.stringify() }).iter() {
        println!("Got: {}", t)
    }
}
