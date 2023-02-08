mod day01;
mod day02;
mod day03;
mod day04;
mod day05;
mod day06;
mod day07;
mod day08;
mod day09;
mod day10;
mod day11;
mod day12;

fn main() -> std::io::Result<()> {
    let day = std::env::args().nth(1).expect("Expect the [day] argument");
    let day = day
        .parse::<u32>()
        .expect("Expect the [day] must be a positive integer");
    match day {
        1 => day01::solve(),
        2 => day02::solve(),
        3 => day03::solve(),
        4 => day04::solve(),
        5 => day05::solve(),
        6 => day06::solve(),
        7 => day07::solve(),
        8 => day08::solve(),
        9 => day09::solve(),
        10 => day10::solve(),
        11 => day11::solve(),
        12 => day12::solve(),
        _ => panic!("Invalid [day] passed: {}", day),
    }
}
