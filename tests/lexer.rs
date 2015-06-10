extern crate templates;

use templates::lexer;

#[test]
fn parse_tokens() {
    let test_str = "<html><p>{{value}}</p><p>{{another}}</p></html>";
    let tokens = lexer::get_tokens(test_str);
    assert_eq!(2, tokens.into_boxed_slice().len());
}

#[test]
fn parse_correct_tokens() {
    let test_str = "<html><p>{{value}}</p><p>{{another value}}</p></html>";
    let tokens = lexer::get_tokens(test_str);
    let first = lexer::Token::Expr("{{value}}".to_string());
    let second = lexer::Token::Expr("{{another value}}".to_string());
    assert!(tokens.contains(&first));
    assert!(tokens.contains(&second));
}
