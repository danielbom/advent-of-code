use crate::utils;

fn part1(text: &str) -> i32 {
    text.chars().fold(0, |floor, ch| match ch {
        '(' => floor + 1,
        ')' => floor - 1,
        _ => floor,
    })
}

fn part2(text: &str) -> i32 {
    let mut floor: i32 = 0;
    for (i, ch) in text.chars().enumerate() {
        match ch {
            '(' => floor += 1,
            ')' => floor -= 1,
            _ => {}
        }
        if floor == -1 {
            return i as i32 + 1;
        }
    }
    -1
}

pub fn solve() -> std::io::Result<()> {
    let mut content = String::new();
    utils::read_file("inputs/day-01.txt", &mut content)?;

    println!("Day 01");
    time_it!("Part 1", part1(&content));
    time_it!("Part 2", part2(&content));

    Ok(())
}
