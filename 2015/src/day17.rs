fn parse_input(content: &str) -> Vec<i32> {
    content.lines().map(|line| line.parse().unwrap()).collect()
}

fn count_containers(containers: &Vec<i32>, expected: i32) -> i32 {
    let mut count: i32 = 0;
    let mut stack = vec![(0, 0)];

    while let Some((index, value)) = stack.pop() {
        if value == expected {
            count += 1;
            continue;
        }
        if value > expected {
            continue;
        }
        if let Some(current) = containers.get(index) {
            stack.push((index + 1, value + current));
            stack.push((index + 1, value));
        }
    }

    count
}

fn count_unique_containers(containers: &Vec<i32>, expected: i32) -> i32 {
    let mut stack = vec![(0, 0, vec![])];
    let mut unique = vec![];
    let mut min_size = containers.len();

    while let Some((index, value, sequence)) = stack.pop() {
        if value == expected {
            min_size = usize::min(min_size, sequence.len());
            unique.push(sequence);
            continue;
        }
        if value > expected {
            continue;
        }
        if sequence.len() > min_size {
            continue;
        }

        if let Some(current) = containers.get(index) {
            let mut next_sequence = vec![*current];
            next_sequence.extend(sequence.iter());
            stack.push((index + 1, value + current, next_sequence));
            stack.push((index + 1, value, sequence));
        }
    }

    unique.iter().filter(|it| it.len() == min_size).count() as i32
}

fn part1(content: &str) -> i32 {
    let containers = parse_input(content);
    count_containers(&containers, 150) // 654
}

fn part2(content: &str) -> i32 {
    let containers = parse_input(content);
    count_unique_containers(&containers, 150) // not: 402, 398, 1, 143, 231, 4, 49
}

pub fn solve() -> std::io::Result<()> {
    let mut content = String::new();
    aoc2015::read_file("inputs/day-17.txt", &mut content)?;

    println!("Day 17");
    aoc2015::time_it!("Part 1", part1(&content));
    aoc2015::time_it!("Part 2", part2(&content));

    Ok(())
}

#[cfg(test)]
mod tests {
    use super::*;

    fn sample_input() -> Vec<i32> {
        let content =  "20
15
10
5
5";
        parse_input(content)
    }

    #[test]
    fn part1_must_works() {
        let containers = sample_input();
        let result = count_containers(&containers, 25);
        assert_eq!(result, 4);
    }

    #[test]
    fn part2_must_works() {
        let containers = sample_input();
        let result = count_unique_containers(&containers, 25);
        assert_eq!(result, 3);
    }
}
