#[derive(Default, PartialEq, Debug)]
struct AuntSue {
    id: u32,
    children: u32,
    cats: u32,
    samoyeds: u32,
    pomeranians: u32,
    akitas: u32,
    vizslas: u32,
    goldfish: u32,
    trees: u32,
    cars: u32,
    perfumes: u32,
}

impl AuntSue {
    fn parse(line: &str) -> Option<AuntSue> {
        // Sue 1: goldfish: 6, trees: 9, akitas: 0
        let mut parts = line.split(' ');
        let _sue = parts.next()?;
        let id = parts.next()?.strip_suffix(':')?.parse().ok()?;
        let mut result = AuntSue::default();

        while let Some(key) = parts.next() {
            let value = parts.next()?;
            let value = value.strip_suffix(',').unwrap_or(value);
            let value = value.parse().ok()?;
            match key {
                "children:" => result.children = value,
                "cats:" => result.cats = value,
                "samoyeds:" => result.samoyeds = value,
                "pomeranians:" => result.pomeranians = value,
                "akitas:" => result.akitas = value,
                "vizslas:" => result.vizslas = value,
                "goldfish:" => result.goldfish = value,
                "trees:" => result.trees = value,
                "cars:" => result.cars = value,
                "perfumes:" => result.perfumes = value,
                _ => panic!("Invalid key: {}", key),
            }
        }

        result.id = id;
        Some(result)
    }

    fn parse_lines(input: &str) -> Vec<Self> {
        input.split('\n').flat_map(Self::parse).collect()
    }
}

fn sue_attributes() -> AuntSue {
    let sue = "Sue 0: \
    children: 3, \
    cats: 7, \
    samoyeds: 2, \
    pomeranians: 3, \
    akitas: 0, \
    vizslas: 0, \
    goldfish: 5, \
    trees: 3, \
    cars: 2, \
    perfumes: 1";
    AuntSue::parse(sue).unwrap()
}

fn part1(aunts: &[AuntSue]) -> u32 {
    let other = sue_attributes();

    let sue_founded = aunts.iter().max_by_key(|it| {
        (it.id == other.id) as u32
            + (it.children == other.children) as u32
            + (it.cats == other.cats) as u32
            + (it.samoyeds == other.samoyeds) as u32
            + (it.pomeranians == other.pomeranians) as u32
            + (it.akitas == other.akitas) as u32
            + (it.vizslas == other.vizslas) as u32
            + (it.goldfish == other.goldfish) as u32
            + (it.trees == other.trees) as u32
            + (it.cars == other.cars) as u32
            + (it.perfumes == other.perfumes) as u32
    });

    sue_founded.map(|it| it.id).unwrap_or(0)
}

fn part2(aunts: &[AuntSue]) -> u32 {
    let other = sue_attributes();

    let sue_founded = aunts.iter().max_by_key(|it| {
        (it.id == other.id) as u32
            + (it.children == other.children) as u32
            + (it.cats > other.cats) as u32
            + (it.samoyeds == other.samoyeds) as u32
            + (it.pomeranians < other.pomeranians) as u32
            + (it.akitas == other.akitas) as u32
            + (it.vizslas == other.vizslas) as u32
            + (it.goldfish < other.goldfish) as u32
            + (it.trees > other.trees) as u32
            + (it.cars == other.cars) as u32
            + (it.perfumes == other.perfumes) as u32
    });

    sue_founded.map(|it| it.id).unwrap_or(0)
}

pub fn solve() -> std::io::Result<()> {
    let mut content = String::new();
    aoc2015::read_file("inputs/day-16.txt", &mut content)?;
    let content = AuntSue::parse_lines(&content);

    println!("Day 16");
    aoc2015::time_it!("Part 1", part1(&content));
    aoc2015::time_it!("Part 2", part2(&content));

    Ok(())
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn parse_lines_must_works() {
        let input = "Sue 1: goldfish: 6, trees: 9, akitas: 0";
        let mut expected = AuntSue::default();
        expected.id = 1;
        expected.goldfish = 6;
        expected.trees = 9;
        expected.akitas = 0;
        let actual = AuntSue::parse(input);
        assert_eq!(actual, Some(expected));
    }
}
