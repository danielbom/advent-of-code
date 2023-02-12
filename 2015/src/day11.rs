use std::collections::HashSet;

fn count_contiguous_by<A>(xs: &[A], p: fn(&A, &A) -> bool) -> Vec<(&A, u32)> {
    xs.iter().fold(Vec::new(), |mut counter, x| {
        match counter.last_mut() {
            Some(last) => {
                if p(last.0, x) {
                    last.1 += 1;
                } else {
                    counter.push((x, 1));
                }
            }
            None => counter.push((x, 1)),
        };
        counter
    })
}

fn is_invalid_char(ch: char) -> bool {
    ch == 'i' || ch == 'o' || ch == 'l'
}

fn contains_invalid_chars(input: &str) -> bool {
    input.chars().any(is_invalid_char)
}

fn contains_two_pairs(input: &str) -> bool {
    let chars = input.chars().collect::<Vec<_>>();
    let chars_counter = count_contiguous_by(chars.as_slice(), |&a, &b| a == b);

    let chars_paired_unique: HashSet<char> = chars_counter
        .iter()
        .map(|&it| (*it.0, it.1 / 2))
        .filter(|&it| it.1 >= 1)
        .map(|it| it.0)
        .collect();

    chars_paired_unique.len() >= 2
}

fn contains_increasing_three_letters(input: &str) -> bool {
    let increasing_letters = input
        .chars()
        .zip(input.chars().skip(1))
        .map(|it| (it.0 as u16, it.1 as u16))
        .map(|it| it.0 == it.1 - 1)
        .collect::<Vec<_>>();

    let increasing_letters_counter =
        count_contiguous_by(increasing_letters.as_slice(), |&a, &b| a == b);

    let increasing_letters_max = increasing_letters_counter
        .iter()
        .filter(|&&it| *it.0)
        .map(|&it| it.1)
        .max()
        .unwrap_or(0);

    increasing_letters_max >= 2
}

fn increment(input: &str) -> String {
    input
        .chars()
        .rev()
        .scan(true, |carry, c| {
            let mut c = c as u8;
            if *carry {
                c += 1;
            }
            if c > b'z' {
                c = b'a';
                *carry = true;
            } else {
                *carry = false;
            }
            Some(c as char)
        })
        .collect::<Vec<char>>()
        .iter()
        .rev()
        .collect::<String>()
}

fn increment_invalid_chars(input: &str) -> String {
    input
        .chars()
        .scan(false, |changed, ch| {
            if *changed {
                Some('a')
            } else if is_invalid_char(ch) {
                *changed = true;
                Some(((ch as u8) + 1) as char)
            } else {
                Some(ch)
            }
        })
        .collect()
}

fn part1(input: &str) -> String {
    let mut password = input.to_string();

    loop {
        password = increment(&password);
        password = increment_invalid_chars(&password);

        if contains_invalid_chars(&password) {
            continue;
        }
        if !contains_two_pairs(&password) {
            continue;
        }
        if !contains_increasing_three_letters(&password) {
            continue;
        }

        return password;
    }
}

fn part2(input: &str) -> String {
    part1(&part1(input))
}

pub fn solve() -> std::io::Result<()> {
    let mut content = String::new();
    aoc2015::read_file("inputs/day-11.txt", &mut content)?;

    println!("Day 11");
    println!("Part 1: {}", part1(&content));
    println!("Part 2: {}", part2(&content));

    Ok(())
}

#[cfg(test)]
mod tests {
    use super::*;

    fn contains_two_pairs_must_returns(expected: bool, input: &str) {
        let result = contains_two_pairs(input);
        assert_eq!(expected, result);
    }

    fn contains_increasing_letters_must_returns(expected: bool, input: &str) {
        let result = contains_increasing_three_letters(input);
        assert_eq!(expected, result);
    }

    #[test]
    fn contains_two_pairs_must_works() {
        contains_two_pairs_must_returns(true, "aabb");
        contains_two_pairs_must_returns(true, "aasdfgbbasdf");

        contains_two_pairs_must_returns(false, "aaaa");
        contains_two_pairs_must_returns(false, "aaa");
        contains_two_pairs_must_returns(false, "aa");
        contains_two_pairs_must_returns(false, "abab");
    }

    #[test]
    fn contains_increasing_letters_must_works() {
        contains_increasing_letters_must_returns(true, "abc");
        contains_increasing_letters_must_returns(true, "ghjaabcc");

        contains_increasing_letters_must_returns(false, "abd");
        contains_increasing_letters_must_returns(false, "hepxddaa");
    }

    #[test]
    fn part1_must_works() {
        assert_eq!(part1("abcdefgh"), "abcdffaa");
        assert_eq!(part1("ghijklmn"), "ghjaabcc");
    }
}
