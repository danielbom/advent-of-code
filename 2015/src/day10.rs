use crate::utils;
struct CountGroupChars<'a> {
    chars: &'a str,
    index: usize,
}

impl<'a> CountGroupChars<'a> {
    fn new(chars: &'a str) -> CountGroupChars<'a> {
        CountGroupChars { chars, index: 0 }
    }
}

impl Iterator for CountGroupChars<'_> {
    type Item = (char, u32);

    fn next(&mut self) -> Option<Self::Item> {
        if self.index >= self.chars.len() {
            None
        } else {
            let mut count = 1;
            let mut index = self.index + 1;
            let bytes = self.chars.as_bytes();
            let byte = bytes[self.index];
            while index < self.chars.len() && bytes[index] == byte {
                count += 1;
                index += 1;
            }
            self.index = index;
            Some((byte as char, count))
        }
    }
}

fn transform(input: &str) -> String {
    CountGroupChars::new(input)
        .map(|(c, count)| format!("{}{}", count, c))
        .collect()
}

fn transform_times(input: &str, n: usize) -> String {
    (0..n).fold(input.to_string(), |input, _| transform(&input))
}

fn part1(content: &str) -> usize {
    transform_times(content, 40).len()
}

fn part2(content: &str) -> usize {
    transform_times(content, 50).len()
}

pub fn solve() -> std::io::Result<()> {
    let mut content = String::new();
    utils::read_file("inputs/day-10.txt", &mut content)?;

    println!("Day 10");
    time_it!("Part 1", part1(&content));
    time_it!("Part 2", part2(&content));

    Ok(())
}
