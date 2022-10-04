use std::iter::Iterator;
use std::cell::RefCell;
use std::rc::Rc;
use regex::{Regex, Captures};

struct Grid<T : Copy> {
    grid: Vec<Vec<T>>
}

struct Parser {
    re: Regex,
}

enum ParseMessage {
    Toggle,
    TurnOn,
    TurnOff
}

struct ParseItem {
    message: ParseMessage,
    x1: usize,
    y1: usize,
    x2: usize,
    y2: usize,
}

impl ParseMessage {
    fn parse<'a>(message: &'a str) -> Self {
        match message {
            "toggle"   => ParseMessage::Toggle,
            "turn on"  => ParseMessage::TurnOn,
            "turn off" => ParseMessage::TurnOff,
            _          => panic!("Invalid message: {}", message)
        }
    }
}

impl ParseItem {
    fn from<'a>(it: Captures<'a>) -> Self {
        let message = ParseMessage::parse(&it[1]);
        let &x1 = &it[2].parse::<usize>().unwrap();
        let &y1 = &it[3].parse::<usize>().unwrap();
        let &x2 = &it[4].parse::<usize>().unwrap();
        let &y2 = &it[5].parse::<usize>().unwrap();
        Self { message, x1, y1, x2, y2 }
    }
}

impl Parser {
    fn new() -> Self {
        let re = Regex::new("(?m)^(turn on|turn off|toggle) (\\d+),(\\d+) through (\\d+),(\\d+)$").unwrap();
        Parser { re }
    }

    fn parse<T>(&self, content: &String, initial: T, fold: fn(T, ParseItem) -> T) -> T {
        self.re.captures_iter(&content).map(ParseItem::from).fold(initial, fold)
    }
}

impl <T : Copy> Grid<T> {
    fn new(initial: T) -> Self {
        Self { grid: vec![vec![initial; 1000]; 1000] }
    }

    fn slice_mut(&mut self, item: &ParseItem, mut map: impl FnMut(&T) -> T) {
        for x in item.x1..item.x2 + 1 {
            for y in item.y1..item.y2 + 1 {
                self.grid[x][y] = map(&self.grid[x][y]);
            }
        }
    }
}

impl Grid<bool> {
    fn count(&self) -> u32 {
        let mut count = 0;
        for row in self.grid.iter() {
            for it in row.iter() {
                if *it {
                    count += 1;
                }
            }
        }
        count
    }
}

impl Grid<u32> {
    fn sum(&self) -> u32 {
        let mut sum = 0;
        for row in self.grid.iter() {
            for it in row.iter() {
                sum += *it;
            }
        }
        sum
    }
}

fn part1(content: &String) -> u32 {
    let rc_grid = Rc::new(RefCell::new(Grid::new(false)));

    let grid = rc_grid.borrow_mut();
    Parser::new().parse(content, grid, |mut grid, item| {
        match item.message {
            ParseMessage::TurnOn => grid.slice_mut(&item, |_| true),
            ParseMessage::TurnOff => grid.slice_mut(&item, |_| false),
            ParseMessage::Toggle => grid.slice_mut(&item, |it| !it),
        };
        grid
    });

    let grid = rc_grid.borrow();
    grid.count()
}

fn part2(content: &String) -> u32 {
    let rc_grid = Rc::new(RefCell::new(Grid::new(0)));

    let grid = rc_grid.borrow_mut();
    Parser::new().parse(content, grid, |mut grid, item| {
        match item.message {
            ParseMessage::TurnOn => grid.slice_mut(&item, |it| it + 1),
            ParseMessage::TurnOff => grid.slice_mut(&item, |it| if *it == 0 { 0 } else { it - 1 }),
            ParseMessage::Toggle => grid.slice_mut(&item, |it| it + 2),
        };
        grid
    });
    
    let grid = rc_grid.borrow();
    grid.sum()
}

pub fn solve() -> std::io::Result<()> {
    let mut content = String::new();
    aoc2015::read_file("inputs/day-06.txt", &mut content)?;

    println!("Day 06");
    println!("Part 1: {}", part1(&content));
    println!("Part 2: {}", part2(&content));

    Ok(())
}

