use std::time;
use std::thread;
use std::sync::{mpsc,Arc};
use md5;

const THREADS_COUNT: i32 = 12;
const TRIES: i32 = 1000;

fn compute(content: &str, start: &str) -> i32 {
    let (tx, rx) = mpsc::sync_channel(THREADS_COUNT as usize);
    let tx = Arc::new(tx);
    let content = Arc::new(content.to_string());
    let start = Arc::new(start.to_string());

    let mut result: Option<i32> = None;
    let mut count = 0;
    while result.is_none() {
        let mut threads = Vec::with_capacity(THREADS_COUNT as usize);

        for _ in 0..THREADS_COUNT {
            let content = content.clone();
            let start = start.clone();
            let tx = tx.clone();
            threads.push(thread::spawn(move || {
                let begin = count * TRIES;
                let end = begin + TRIES;
                for j in begin..end {
                    let input = format!("{}{}", content, j);
                    let digest = md5::compute(input);
                    let digest = format!("{:x}", digest);
                    if digest.starts_with(&format!("{}", start)) {
                        tx.send(Some(j)).unwrap();
                        return;
                    }
                }
                tx.send(None).unwrap();
            }));
            count += 1;
        }

        threads.into_iter().for_each(|it| {
            if let Some(thread_result) = rx.recv().unwrap() {
                if let Some(final_result) = result {
                    if final_result > thread_result {
                        result = Some(thread_result);
                    }
                } else {
                    result = Some(thread_result);
                }
            }

            it.join().unwrap();
        });
    }

    result.unwrap_or(0)
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

fn part1(content: &str) {
    let now = time::Instant::now();
    let result = compute(content, "00000");
    println!("Part 1: {} [{} s]", result, now.elapsed().as_secs());
}

fn part2(content: &str) {
    let now = time::Instant::now();
    let result = compute(content, "000000");
    println!("Part 2: {} [{} s]", result, now.elapsed().as_secs());
}

pub fn solve() -> std::io::Result<()> {
    let mut content = String::new();
    aoc2015::read_file("inputs/day-04.txt", &mut content)?;
    let content = content.trim_end();

    println!("Day 04");
    part1(content);
    part2(content);

    Ok(())
}
