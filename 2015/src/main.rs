#[macro_use]
mod utils;

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
mod day13;
mod day14;
mod day15;
mod day16;
mod day17;
mod day18;
mod day19;
mod day20;
mod day21;
mod day22;
mod day23;
mod day24;
mod day25;

fn run_day(day: u32) -> std::io::Result<()> {
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
        13 => day13::solve(),
        14 => day14::solve(),
        15 => day15::solve(),
        16 => day16::solve(),
        17 => day17::solve(),
        18 => day18::solve(),
        19 => day19::solve(),
        20 => day20::solve(),
        21 => day21::solve(),
        22 => day22::solve(),
        23 => day23::solve(),
        24 => day24::solve(),
        25 => day25::solve(),
        _ => panic!("Invalid [day] passed: {}", day),
    }
}

fn main() -> std::io::Result<()> {
    let day = std::env::args()
        .nth(1)
        .expect("Expect the <day> argument: day [0..25]");
    let day = day
        .parse::<u32>()
        .expect("Expect the [day] must be a positive integer");
    if day == 0 {
        for day in 1..=25 {
            run_day(day)?;
        }
    } else {
        run_day(day)?;
    }
    Ok(())
}
