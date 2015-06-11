extern crate templates;

use templates::{parser, util};

use std::env;
use std::fs::{File, OpenOptions};
use std::io::{Read, Write};


fn main() {
    let mut env = env::args();

    if env.len() != 2 {
        println!("Usage: templates <input-file> <json-file>");
        std::process::exit(1);
    }

    let input_file = env.nth(0).unwrap();
    let data_file = env.nth(1);

    let values = parser::parse_file(&input_file);
    let json = util::parse_json_file(&data_file.unwrap());

    let mut file = File::open(&input_file.clone()).unwrap();
    let mut data = String::new();

    file.read_to_string(&mut data).unwrap();


    //map the values
    let done = util::do_replace(&values, &json, &data);

    let mut out_file = OpenOptions::new().write(true).open(&input_file.clone()).unwrap();
    out_file.write_all(&done.into_bytes()).unwrap()
}
