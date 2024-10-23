
use std::sync::{mpsc, Arc, Mutex};
use std::thread;

const THREADS_COUNT: i32 = 8;
const TRIES: i32 = 1000;

fn compute(content: &str, start: &str) -> i32 {
    let (result_tx, result_rx) = mpsc::sync_channel(THREADS_COUNT as usize);
    let (range_tx, range_rx) = mpsc::sync_channel(THREADS_COUNT as usize);
    let result_tx = Arc::new(result_tx);
    let range_rx = Arc::new(Mutex::new(range_rx));
    let content = Arc::new(content.to_string());
    let start = Arc::new(start.to_string());
    let mut count = 0;
    let produce = &mut || {
        range_tx.send(Some(count)).unwrap();
        count += 1;
    };

    for _ in 0..THREADS_COUNT {
        produce();
        let content = content.clone();
        let start = start.clone();
        let result_tx = result_tx.clone();
        let range_rx = range_rx.clone();
        thread::spawn(move || loop {
            let count = {
                match range_rx.lock().unwrap().recv() {
                    Ok(Some(from)) => from,
                    Ok(None) => return,
                    Err(_) => return,
                }
            };
            let from = count * TRIES;

            let result = (from..from + TRIES).find(|j| {
                let input = format!("{}{}", content, j);
                let digest = md5::compute(input);
                let digest = format!("{:x}", digest);
                digest.starts_with(&format!("{}", start))
            });

            result_tx.send(result).ok().unwrap_or_default()
        });
    }

    loop {
        match result_rx.recv().unwrap() {
            Some(result) => return result,
            None => produce(),
        }
    }
}

#[allow(dead_code)]
fn compute1(content: &str, start: &str) -> u32 {
    let mut count = 0;
    loop {
        let input = format!("{}{}", content, count);
        let digest = md5::compute(input);
        let digest = format!("{:x}", digest);
        if digest.starts_with(start) {
            break;
        }
        count += 1;
    }
    count
}

fn part1(content: &str) -> i32 {
    compute(content, "00000")
}

fn part2(content: &str) -> i32 {
    compute(content, "000000")
}

pub fn solve() -> std::io::Result<()> {
    let mut content = String::new();
    aoc2015::read_file("inputs/day-04.txt", &mut content)?;
    let content = content.trim_end();

    println!("Day 04");
    aoc2015::time_it!("Part 1", part1(content));
    aoc2015::time_it!("Part 2", part2(content));

    Ok(())
}
