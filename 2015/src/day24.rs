use crate::utils;

mod priority_queue {
    // Based on https://github.com/Kezzryn/Advent-of-Code/blob/main/2015/Day%2024/Program.cs
    use std::cmp::Ordering;
    use std::collections::{BinaryHeap, HashSet};

    #[derive(Debug, Clone, Eq, PartialEq)]
    struct Package {
        bag: usize,
        weight: usize,
    }

    impl Package {
        fn new(bag: usize, weight: usize) -> Self {
            Self {
                bag,
                weight,
            }
        }

        fn contains(&self, bit: usize) -> bool {
            self.bag & bit > 0
        }

        fn union(&self, other: &Self) -> Self {
            Self {
                bag: self.bag | other.bag,
                weight: self.weight + other.weight,
            }
        }

        fn len(&self) -> u32 {
            self.bag.count_ones()
        }
    }

    impl Ord for Package {
        fn cmp(&self, other: &Self) -> Ordering {
            self.weight.cmp(&other.weight)
        }
    }

    impl PartialOrd for Package {
        fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
            Some(self.cmp(other))
        }
    }

    pub fn compute(values: &[usize], total_parts: usize) -> u64 {
        let total_weight = values.iter().sum::<usize>();
        let part_weight = total_weight / total_parts;

        let packages = values
            .iter()
            .enumerate()
            .map(|(i, x)| Package::new(1 << i, *x))
            .collect::<Vec<Package>>();

        let mut queue = BinaryHeap::new();
        let mut seen = HashSet::new();
        let mut best_quantum = u64::MAX;
        let mut best_len = u32::MAX;

        for p in packages.iter() {
            queue.push(p.clone());
        }

        while let Some(p1) = queue.pop() {
            if seen.contains(&p1.bag) {
                continue;
            }
            seen.insert(p1.bag);

            if p1.weight == part_weight {
                let quantum = values
                    .iter()
                    .enumerate()
                    .filter(|(i, _)| p1.contains(1 << *i))
                    .map(|(_, x)| *x)
                    .fold(1, |a, b| a * b as u64);
                best_len = p1.len();
                best_quantum = u64::min(best_quantum, quantum);
                continue;
            }

            for p2 in packages.iter() {
                if !p1.contains(p2.bag) {
                    let p3 = p1.union(&p2);
                    if p3.weight <= part_weight && p3.len() <= best_len {
                        queue.push(p3);
                    }
                }
            }
        }

        best_quantum
    }
}

mod greedy {
    #[derive(Debug, Clone, Eq, PartialEq)]
    struct Package {
        bag: usize,
        weight: usize,
    }

    impl Package {
        fn new(bag: usize, weight: usize) -> Self {
            Self { bag, weight }
        }

        fn contains(&self, bit: usize) -> bool {
            self.bag & bit > 0
        }

        fn union(&self, other: &Self) -> Self {
            Self {
                bag: self.bag | other.bag,
                weight: self.weight + other.weight,
            }
        }

        fn len(&self) -> u32 {
            self.bag.count_ones()
        }
    }

    pub fn compute(values: &[usize], total_parts: usize) -> u64 {
        let values = {
            let mut values = values.to_vec();
            values.sort();
            values
        };
        let total_weight = values.iter().sum::<usize>();
        let part_weight = total_weight / total_parts;

        let mut queue = values
            .iter()
            .enumerate()
            .map(|(i, x)| (Package::new(1 << i, *x), i))
            .collect::<Vec<_>>();

        let mut best_quantum = u64::MAX;
        let mut best_count = u32::MAX;

        while let Some((p1, start)) = queue.pop() {
            if p1.weight > part_weight {
                return 0;
            }

            if p1.weight == part_weight {
                let quantum = values
                    .iter()
                    .enumerate()
                    .filter(|(i, _)| p1.contains(1 << *i))
                    .map(|(_, x)| *x)
                    .fold(1, |a, b| a * b as u64);
                best_count = p1.len();
                best_quantum = u64::min(best_quantum, quantum);
            } else {
                for i in (start + 1)..values.len() {
                    let p2 = Package::new(1 << i, values[i]).union(&p1);
                    if p2.weight <= part_weight && p2.len() <= best_count {
                        queue.push((p2, i));
                    }
                }
            }
        }

        best_quantum
    }
}

fn parse_values(content: &str) -> Vec<usize> {
    content
        .lines()
        .map(|line| line.parse().ok())
        .flatten()
        .collect::<Vec<usize>>()
}

const GREEDY: bool = true;

fn part1(content: &str) -> u64 {
    let values = parse_values(content);
    if GREEDY {
        greedy::compute(&values, 3)
    } else {
        priority_queue::compute(&values, 3)
    }
}

fn part2(content: &str) -> u64 {
    let values = parse_values(content);
    if GREEDY {
        greedy::compute(&values, 4)
    } else {
        priority_queue::compute(&values, 4)
    }
}

pub fn solve() -> std::io::Result<()> {
    let mut content = String::new();
    utils::read_file("inputs/day-24.txt", &mut content)?;

    println!("Day 24");
    time_it!("Part 1", part1(&content));
    time_it!("Part 2", part2(&content));

    Ok(())
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        let values = (1..=11).filter(|x| *x != 6).collect::<Vec<usize>>();
        let expect = 99;
        let result = greedy::compute(&values, 3);
        assert_eq!(expect, result);
    }
}
