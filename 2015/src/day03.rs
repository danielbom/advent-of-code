use std::sync::{Arc, Mutex};
use std::thread;

use std::fs::File;
use std::io::BufReader;
use std::io::prelude::*;

use std::collections::HashSet;

#[derive(Hash, Copy, Clone, Default, Eq, PartialEq, Debug)]
struct SantaMove {
    x: i32,
    y: i32,
}

impl SantaMove {
    fn move_to(&mut self, ch: char) {
        match ch {
            '>' => self.x += 1,
            '<' => self.x -= 1,
            '^' => self.y += 1,
            'v' => self.y -= 1,
            _   => {}
        }
    }
}

fn read_file_async(mutex: Arc<Mutex<String>>) -> std::io::Result<()> {
    let file = File::open("inputs/day-03.txt")?;
    let mut buf_reader = BufReader::new(file);
    let mut content = mutex.lock().unwrap();
    buf_reader.read_to_string(&mut content)?;
    Ok(())
}

fn part1(mutex: Arc<Mutex<String>>) -> std::io::Result<()> {
    let mut set = HashSet::new();
    let mut santa = SantaMove::default();
    set.insert(santa);

    let content = mutex.lock().unwrap();
    for ch in content.chars() {
        santa.move_to(ch);
        set.insert(santa);
    }

    println!("Part 1: {}", set.len());
    Ok(())
}

fn part2(mutex: Arc<Mutex<String>>) -> std::io::Result<()> {
    let mut set = HashSet::new();
    let mut santa = SantaMove::default();
    let mut robot = SantaMove::default();
    set.insert(santa);

    let content = mutex.lock().unwrap();
    for (i, ch) in content.chars().enumerate() {
        if i % 2 == 0 {
            santa.move_to(ch);
            set.insert(santa);
        } else {
            robot.move_to(ch);
            set.insert(robot);
        }
    }

    println!("Part 2: {}", set.len());
    Ok(())
}

fn solve() -> std::io::Result<()> {
    let content = Arc::new(Mutex::new(String::new()));
    let mut threads = Vec::new();
    {
        let content = Arc::clone(&content);
        threads.push(thread::spawn(move || part1(content).unwrap()));
    }
    {
        let content = Arc::clone(&content);
        threads.push(thread::spawn(move || part2(content).unwrap()));
    }
    {
        let content = Arc::clone(&content);
        threads.push(thread::spawn(move || read_file_async(content).unwrap()));
    }

    threads.into_iter().for_each(|thread| {
        thread.join().unwrap();
    });

    Ok(())
}
