mod lib;
use md5;

fn solve(content: &str, start: &str) -> u32 {
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
    let result = solve(content, "00000");
    println!("Part 1: {}", result);
}

fn part2(content: &str) {
    let result = solve(content, "000000");
    println!("Part 2: {}", result);
}

fn main() -> std::io::Result<()> {
    let mut content = String::new();
    lib::read_file("inputs/day-04.txt", &mut content)?;
    let content = content.trim_end();

    part1(content);
    part2(content);

    Ok(())
}
