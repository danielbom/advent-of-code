use crate::utils;
use std::collections::HashSet;

#[derive(Debug)]
struct Replacement(String, String);

impl Replacement {
    pub fn replace_each(&self, content: &str) -> Vec<String> {
        let parts = content.split(&self.0).collect::<Vec<_>>();
        let mut result = vec![];
        for i in 1..parts.len() {
            let mut sb = String::with_capacity(content.len() + self.1.len());
            sb.push_str(parts[0]);
            for j in 1..parts.len() {
                if j == i {
                    sb.push_str(&self.1);
                } else {
                    sb.push_str(&self.0);
                }
                sb.push_str(parts[j]);
            }
            result.push(sb);
        }
        result
    }

    pub fn replace_all(&self, content: &str) -> (String, usize) {
        let parts = content.split(&self.0).collect::<Vec<_>>();
        let mut sb = String::with_capacity(content.len() + self.1.len());
        sb.push_str(parts[0]);
        for i in 1..parts.len() {
            sb.push_str(&self.1);
            sb.push_str(parts[i]);
        }
        (sb, parts.len() - 1)
    }

    pub fn swap(self) -> Replacement {
        Replacement(self.1, self.0)
    }
}

fn parse_input(content: &str) -> (Vec<Replacement>, String) {
    let mut replacements = vec![];

    let mut lines = content.lines();
    for line in lines.by_ref() {
        if line.is_empty() {
            break;
        }

        let mut parts = line.split(" => ");
        let molecule_from = parts.next().unwrap().to_string();
        let molecule_to = parts.next().unwrap().to_string();
        replacements.push(Replacement(molecule_from, molecule_to));
    }

    let molecule = lines.next().unwrap_or("").to_string();

    (replacements, molecule)
}

fn part1(content: &str) -> i32 {
    let (replacements, molecule) = parse_input(content);
    let mut set = HashSet::<String>::new();

    for replacement in replacements {
        for new_molecule in replacement.replace_each(&molecule) {
            set.insert(new_molecule);
        }
    }

    set.len() as i32
}

fn part2(content: &str) -> i32 {
    let (replacements, molecule) = {
        let mut input = parse_input(content);
        input.0 = input.0.into_iter().map(|r| r.swap()).collect();
        input.0.sort_by(|a, b| a.0.len().cmp(&b.0.len()).reverse());
        input
    };

    let mut count = 0;
    let mut current_molecule = molecule.clone();
    let mut prev_molecule = current_molecule.clone();

    while !current_molecule.chars().all(|ch| ch == 'e') {
        for replacement in replacements.iter() {
            let (next_molecule, transforms) = replacement.replace_all(&current_molecule);
            count += transforms;
            current_molecule = next_molecule;
        }
        if current_molecule == prev_molecule {
            break;
        }
        prev_molecule = current_molecule.clone();
    }

    count as i32
}

pub fn solve() -> std::io::Result<()> {
    let mut content = String::new();
    utils::read_file("inputs/day-19.txt", &mut content)?;

    println!("Day 19");
    time_it!("Part 1", part1(&content));
    time_it!("Part 2", part2(&content));

    Ok(())
}

#[cfg(test)]
mod tests {
    use super::*;

    fn test1_input() -> String {
        "H => HO
H => OH
O => HH

"
        .to_string()
    }
    fn test2_input() -> String {
        "e => H
e => O
H => HO
H => OH
O => HH

"
        .to_string()
    }

    #[test]
    fn test_part1_01() {
        let input = test1_input();
        let expected = 4;
        let result = part1(&format!("{}{}", input, "HOH"));
        assert_eq!(expected, result);
    }

    #[test]
    fn test_part1_02() {
        let input = test1_input();
        let expected = 7;
        let result = part1(&format!("{}{}", input, "HOHOHO"));
        assert_eq!(expected, result);
    }

    #[test]
    fn test_part2_01() {
        let input = test2_input();
        let expected = 3 + 1;
        let result = part2(&format!("{}{}", input, "HOH"));
        assert_eq!(expected, result);
    }

    #[test]
    fn test_part2_02() {
        let input = test2_input();
        let expected = 6 + 1;
        let result = part2(&format!("{}{}", input, "HOHOHO"));
        assert_eq!(expected, result);
    }
}
