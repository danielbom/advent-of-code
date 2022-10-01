use md5;

fn compute(content: &str, start: &str) -> u32 {
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
    let result = compute(content, "00000");
    println!("Part 1: {}", result);
}

fn part2(content: &str) {
    let result = compute(content, "000000");
    println!("Part 2: {}", result);
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
