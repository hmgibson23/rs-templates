extern crate rustc_serialize;

use std::fs::File;
use std::io::Read;
use parser::Value;

pub fn parse_json_file(file: &str) -> rustc_serialize::json::Json {
    let mut file = File::open(file).unwrap();
    let mut data = String::new();

    file.read_to_string(&mut data).unwrap();

    let json = rustc_serialize::json::Json::from_str(&data).unwrap();
    json
}

pub fn do_replace(values: &Vec<Option<Value>>, json: &rustc_serialize::json::Json, input: &String) -> String {

    values.iter().fold(input.clone(), |acc, item| {
        let val = item.clone().unwrap();
        let replace = json.find(&val.value).unwrap();

        let ret = acc.replace(&val.source, &replace.pretty().to_string());
        ret
    })
}
