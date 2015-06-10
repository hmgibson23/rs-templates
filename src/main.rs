extern crate templates;

use templates::{parser, util};

use std::env;


fn main() {
    let mut env = env::args();

    if env.len() != 2 {
        println!("Usage: templates <input-file> <json-file>");
        std::process::exit(1);
    }

    let input_file = env.nth(0);
    let data_file = env.nth(1);

    let values = parser::parse_file(&input_file.unwrap());
    let json = util::parse_json_file(&data_file.unwrap());

    //map the values

}
