use std::fs::File;
use std::io::BufReader;
use std::io::prelude::*;

pub fn read_file(filename: &str, mut content: &mut String) -> std::io::Result<()> {
    let file = File::open(filename)?;
    let mut buf_reader = BufReader::new(file);
    buf_reader.read_to_string(&mut content)?;
    Ok(())
}

