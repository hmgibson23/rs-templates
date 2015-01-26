mod lexer;

fn main() {
    let token = lexer::get_token("<html>{{value}}</html>");
    println!("Got: {}", token.stringify())
}
