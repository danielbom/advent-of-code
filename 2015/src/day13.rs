use crate::utils;
use itertools::Itertools;
use std::collections::{HashMap, HashSet};

#[derive(Debug, PartialEq, Eq)]
struct SittingNext {
    who: String,
    next: String,
    happiness: i32,
}

impl SittingNext {
    fn parse(input: &str) -> Option<SittingNext> {
        let mut parts = input.split(" would ");
        let who = parts.next()?.to_string();
        let mut parts = parts.next()?.split(' ');
        let gain = parts.next()?.eq("gain");
        let happiness = parts.next()?.parse::<i32>().ok()?;
        let happiness = if gain { happiness } else { -happiness };
        let mut parts = parts.skip(6);
        let next = parts.next()?.strip_suffix('.')?.to_string();

        Some(SittingNext {
            who,
            next,
            happiness,
        })
    }

    fn parse_lines(input: &str) -> Vec<SittingNext> {
        input.split('\n').flat_map(SittingNext::parse).collect()
    }
}

fn get_peoples(paths: &[SittingNext]) -> Vec<&str> {
    paths
        .iter()
        .map(|it| it.who.as_str())
        .collect::<HashSet<_>>()
        .into_iter()
        .collect::<Vec<_>>()
}

fn compute(paths: &[SittingNext]) -> i32 {
    let peoples = get_peoples(paths);
    let paths = paths
        .iter()
        .map(|it| ((it.who.as_str(), it.next.as_str()), it))
        .collect::<HashMap<_, _>>();
    let peoples_count = peoples.len();

    peoples
        .into_iter()
        .permutations(peoples_count)
        .map(|sitting| {
            (0..peoples_count)
                .map(|i| {
                    let left = if i == 0 { peoples_count - 1 } else { i - 1 };
                    let right = if i == peoples_count - 1 { 0 } else { i + 1 };
                    let left = paths[&(sitting[i], sitting[left])];
                    let right = paths[&(sitting[i], sitting[right])];
                    left.happiness + right.happiness
                })
                .sum::<i32>()
        })
        .max()
        .unwrap()
}

fn part1(input: &str) -> i32 {
    let paths = SittingNext::parse_lines(input);
    compute(&paths)
}

fn part2(input: &str) -> i32 {
    let paths = {
        let paths = SittingNext::parse_lines(input);

        let my_paths: Vec<SittingNext> = get_peoples(&paths)
            .into_iter()
            .flat_map(|next| {
                vec![
                    SittingNext {
                        who: "Self".to_string(),
                        next: next.to_string(),
                        happiness: 0,
                    },
                    SittingNext {
                        who: next.to_string(),
                        next: "Self".to_string(),
                        happiness: 0,
                    },
                ]
            })
            .collect();

        let mut result: Vec<SittingNext> = Vec::new();
        result.extend(paths);
        result.extend(my_paths);
        result
    };

    compute(&paths)
}

pub fn solve() -> std::io::Result<()> {
    let mut content = String::new();
    utils::read_file("inputs/day-13.txt", &mut content)?;

    println!("Day 13");
    time_it!("Part 1", part1(&content));
    time_it!("Part 2", part2(&content));

    Ok(())
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn row_parse_must_works() {
        let input = "Alice would gain 54 happiness units by sitting next to Bob.";
        let expected = SittingNext {
            who: "Alice".to_string(),
            next: "Bob".to_string(),
            happiness: 54,
        };
        let result = SittingNext::parse(input);
        assert!(result.is_some());
        let result = result.unwrap();
        assert_eq!(expected, result);
    }

    #[test]
    fn part1_must_works() {
        let input = "Alice would gain 54 happiness units by sitting next to Bob.
Alice would lose 79 happiness units by sitting next to Carol.
Alice would lose 2 happiness units by sitting next to David.
Bob would gain 83 happiness units by sitting next to Alice.
Bob would lose 7 happiness units by sitting next to Carol.
Bob would lose 63 happiness units by sitting next to David.
Carol would lose 62 happiness units by sitting next to Alice.
Carol would gain 60 happiness units by sitting next to Bob.
Carol would gain 55 happiness units by sitting next to David.
David would gain 46 happiness units by sitting next to Alice.
David would lose 7 happiness units by sitting next to Bob.
David would gain 41 happiness units by sitting next to Carol.";
        let expected = 330;
        let result = part1(input);
        assert_eq!(expected, result);
    }
}
