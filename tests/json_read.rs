extern crate templates;
extern crate rustc_serialize;

use templates::util;
use templates::lexer;
use templates::parser;

#[test]
fn should_read_json_file() {
    let json = util::parse_json_file("/Users/hgibson/git/rs-templates/tests/test.json");
    assert!(json.is_object());
}

#[test]
fn should_replace_all_values() {
    let input = "<html><p>{{another}}</p><p></p></html>";
    let json = rustc_serialize::json::Json::from_str("{\"another\": 20}").unwrap();
    let tokens = lexer::get_tokens(&input);
    let parsed = tokens.iter().map(parser::parse_token).collect::<Vec<_>>();
    let result = util::do_replace(&parsed, &json, &input.to_string());
    assert!(result.contains("20"))
}
