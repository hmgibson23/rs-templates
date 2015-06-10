extern crate templates;

use templates::parser;
use templates::lexer;

#[test]
fn parse_tokens() {
    let tokens = lexer::get_tokens("<html><p>{{value}}</p><p>{{another value}}</p></html>");

    let parsed = tokens.iter().map(parser::parse_token).collect::<Vec<_>>();
    assert_eq!(2, parsed.len());
}

#[test]
fn parse_tokens_correct() {
    let tokens = lexer::get_tokens("<html><p>{{value}}</p><p>{{another value}}</p></html>");

    let parsed = tokens.iter().map(parser::parse_token).collect::<Vec<_>>();
    let first = Some(parser::Value {value: "value".to_string(),
                                      source: "{{value}}".to_string()});
    assert!(parsed.contains(&first));
}
