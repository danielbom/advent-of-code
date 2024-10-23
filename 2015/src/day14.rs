#[derive(Debug)]
struct Reindeer {
    #[allow(dead_code)]
    name: String,
    acceleration: i32,
    time: i32,
    rest_time: i32,
}

#[derive(Debug)]
struct ReindeerPoint {
    reindeer: Reindeer,
    points: i32,
    distance: i32,
}

impl Reindeer {
    fn parse(line: &str) -> Option<Reindeer> {
        let mut parts = line.split(" can fly ");
        let name = parts.next()?.to_string();
        let mut parts = parts.next()?.split(" km/s for ");
        let acceleration = parts.next()?.parse::<i32>().unwrap();
        let mut parts = parts.next()?.split(" seconds, but then must rest for ");
        let time = parts.next()?.parse::<i32>().unwrap();
        let mut parts = parts.next()?.split(" seconds.");
        let rest_time = parts.next()?.parse::<i32>().unwrap();
        Some(Reindeer {
            name,
            acceleration,
            time,
            rest_time,
        })
    }

    fn parse_lines(content: &str) -> Vec<Reindeer> {
        content.split('\n').flat_map(Reindeer::parse).collect()
    }

    fn total_move_in(&self, time_available: i32) -> i32 {
        let time_to_move_and_rest = self.time + self.rest_time;
        let times = time_available / time_to_move_and_rest;
        let missing_time = time_available - times * time_to_move_and_rest;
        let total_time = times * self.time + missing_time.min(self.time);
        total_time * self.acceleration
    }

    fn is_resting(&self, time: i32) -> bool {
        time % (self.time + self.rest_time) >= self.time
    }
}

fn compute_points_by_distance(reindeers: Vec<Reindeer>, time_available: i32) -> i32 {
    reindeers
        .iter()
        .map(|it| it.total_move_in(time_available))
        .max()
        .unwrap()
}

fn part1(input: &str) -> i32 {
    let reindeers = Reindeer::parse_lines(input);
    let time_available = 2503;
    compute_points_by_distance(reindeers, time_available)
}

fn compute_points_by_leader_time(reindeers: Vec<Reindeer>, time_available: i32) -> i32 {
    let mut input = reindeers
        .into_iter()
        .map(|it| ReindeerPoint {
            points: 0,
            distance: 0,
            reindeer: it,
        })
        .collect::<Vec<_>>();

    let mut time = 0;
    while time <= time_available {
        input
            .iter_mut()
            .filter(|it| !it.reindeer.is_resting(time))
            .for_each(|it| it.distance += it.reindeer.acceleration);

        let winner_points = input.iter().map(|it| it.distance).max().unwrap();

        input
            .iter_mut()
            .filter(|it| it.distance == winner_points)
            .for_each(|it| it.points += 1);

        time += 1;
    }

    input.iter().map(|it| it.points).max().unwrap()
}

fn part2(input: &str) -> i32 {
    let reindeers = Reindeer::parse_lines(input);
    let time_available = 2503;
    compute_points_by_leader_time(reindeers, time_available)
}

pub fn solve() -> std::io::Result<()> {
    let mut content = String::new();
    aoc2015::read_file("inputs/day-14.txt", &mut content)?;

    println!("Day 14");
    aoc2015::time_it!("Part 1", part1(&content));
    aoc2015::time_it!("Part 2", part2(&content));

    Ok(())
}

#[cfg(test)]
mod tests {
    use super::*;

    #[inline]
    fn input_sample<'a>() -> &'a str {
        "Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.
Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds."
    }

    #[test]
    fn compute_points_by_distance_must_works() {
        let input = input_sample();
        let reindeers = Reindeer::parse_lines(input);
        let result = compute_points_by_distance(reindeers, 1000);
        let expected = 1120;
        assert_eq!(expected, result);
    }

    #[test]
    fn compute_points_by_leader_time_must_works() {
        let input = input_sample();
        let reindeers = Reindeer::parse_lines(input);
        let result = compute_points_by_leader_time(reindeers, 1000);
        let expected = 689;
        assert_eq!(expected, result);
    }
}
