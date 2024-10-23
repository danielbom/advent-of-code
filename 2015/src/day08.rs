use crate::utils;
use regex::Regex;

fn tokenizer_codes() -> Regex {
    Regex::new(r#"\\x[\da-f]{1,2}|\\["\\]|\w"#).unwrap()
}

fn tokenizer_special_chars() -> Regex {
    Regex::new(r#""|\\"#).unwrap()
}

fn count_regex(re: &Regex, input: &str) -> usize {
    re.captures_iter(input).count()
}

fn count_characters(input: &str) -> usize {
    input.chars().count()
}

fn part1(content: &str) -> i32 {
    let re = tokenizer_codes();
    let (total_characters, total_codes) = content
        .lines()
        .map(|line| (count_characters(line), count_regex(&re, line)))
        .fold((0, 0), |acc, curr| (acc.0 + curr.0, acc.1 + curr.1));

    total_characters as i32 - total_codes as i32
}

fn part2(content: &str) -> i32 {
    let scapes_adjust = 2;
    let re = tokenizer_special_chars();
    content
        .lines()
        .map(|it| count_regex(&re, it))
        .map(|it| it + scapes_adjust)
        .sum::<usize>() as i32
}

pub fn solve() -> std::io::Result<()> {
    let mut content = String::new();
    utils::read_file("inputs/day-08.txt", &mut content)?;

    println!("Day 08");
    time_it!("Part 1", part1(&content));
    time_it!("Part 2", part2(&content));

    Ok(())
}

#[cfg(test)]
mod tests {
    use super::*;

    fn sample_fixture() -> String {
        "\"\"\n\"abc\"\n\"aaa\\\"aaa\"\n\"\\x27\"".to_string()
    }

    #[test]
    fn count_codes_test() {
        let re = tokenizer_codes();
        let input = sample_fixture();
        let inputs = input.lines().collect::<Vec<_>>();
        let expected = vec![0, 3, 7, 1];
        let result = inputs
            .iter()
            .map(|it| count_regex(&re, it))
            .collect::<Vec<_>>();

        assert_eq!(expected, result);
    }

    #[test]
    fn part1_test() {
        let expected = 12;
        let input = sample_fixture();
        let result = part1(&input);

        assert_eq!(expected, result);
    }

    #[test]
    fn count_special_characters() {
        let re = tokenizer_special_chars();
        let input = sample_fixture();
        let inputs = input.lines().collect::<Vec<_>>();
        let expected = vec![2, 2, 4, 3];
        let result = inputs
            .iter()
            .map(|it| count_regex(&re, it))
            .collect::<Vec<_>>();

        assert_eq!(expected, result);
    }

    #[test]
    fn part2_test() {
        let input = sample_fixture();
        let expected = 19;
        let result = part2(&input);

        assert_eq!(expected, result);
    }
}
