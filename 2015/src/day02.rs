use crate::utils;

use std::cmp::min;

struct Dim {
    l: u32,
    w: u32,
    h: u32,
}

impl Dim {
    fn parse(text: &str) -> Self {
        let parts = text
            .split('x')
            .map(|it| it.parse().unwrap())
            .collect::<Vec<_>>();
        Self {
            l: parts[0],
            w: parts[1],
            h: parts[2],
        }
    }

    fn surface_area(&self) -> u32 {
        let a = self.l * self.w;
        let b = self.w * self.h;
        let c = self.h * self.l;
        2 * (a + b + c) + min(a, min(b, c))
    }

    fn smallest_perimeter(&self) -> u32 {
        let mut xs = [self.l, self.w, self.h];
        xs.sort_unstable();
        2 * (xs[0] + xs[1])
    }

    fn volume(&self) -> u32 {
        self.l * self.w * self.h
    }
}

fn part1(content: &str) -> u32 {
    content
        .lines()
        .map(Dim::parse)
        .map(|it| it.surface_area())
        .sum()
}

fn part2(content: &str) -> u32 {
    content
        .lines()
        .map(Dim::parse)
        .map(|it| it.smallest_perimeter() + it.volume())
        .sum()
}

pub fn solve() -> std::io::Result<()> {
    let mut content = String::new();
    utils::read_file("inputs/day-02.txt", &mut content)?;

    println!("Day 02");
    time_it!("Part 1", part1(&content));
    time_it!("Part 2", part2(&content));

    Ok(())
}
