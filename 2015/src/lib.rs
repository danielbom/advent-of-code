use std::fs::File;
use std::io::prelude::*;
use std::io::BufReader;

pub fn read_file(filename: &str, content: &mut String) -> std::io::Result<()> {
    let file = File::open(filename)?;
    let mut buf_reader = BufReader::new(file);
    buf_reader.read_to_string(content)?;
    Ok(())
}
