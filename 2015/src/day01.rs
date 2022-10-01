use std::fs::File;
use std::io::BufReader;
use std::io::prelude::*;

fn read_file(filename: &str, mut content: &mut String) -> std::io::Result<()> {
    let file = File::open(filename)?;
    let mut buf_reader = BufReader::new(file);
    buf_reader.read_to_string(&mut content)?;
    Ok(())
}

fn part1(text: &str) -> i32 {
    text.chars().fold(0, |floor, c| match c {
        '(' => floor + 1,
        ')' => floor - 1,
        _   => floor
    })
}

fn part2(text: &str) -> i32 {
    let mut floor: i32 = 0;
    for (i, ch) in text.chars().enumerate() {
        match ch {
            '(' => floor += 1,
            ')' => floor -= 1,
            _   => {}
        }
        if floor == -1 {
            return i as i32 + 1;
        }
    }
    -1
}

fn solve() -> std::io::Result<()> {
    let mut content = String::new();
    read_file("inputs/day-01.txt", &mut content)?;

    println!("Part 1: {}", part1(&content.as_str()));
    println!("Part 2: {}", part2(&content.as_str()));

    Ok(())
}
