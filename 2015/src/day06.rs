use std::cell::RefCell;
use std::rc::Rc;
use regex::Regex;

struct Grid<T : Copy> {
    grid: Vec<Vec<T>>
}

impl <T : Copy> Grid<T> {
    fn new(initial: T) -> Self {
        Self { grid: vec![vec![initial; 1000]; 1000] }
    }

    fn slice_mut(&mut self, x1: usize, y1: usize, x2: usize, y2: usize, mut map: impl FnMut(&T) -> T) {
        for x in x1..x2 + 1 {
            for y in y1..y2 + 1 {
                self.grid[x][y] = map(&self.grid[x][y]);
            }
        }
    }
}

fn parse_with(content: &String, mut callback: impl FnMut(&str, usize, usize, usize, usize)) {
    let re: Regex = Regex::new("(?m)^(turn on|turn off|toggle) (\\d+),(\\d+) through (\\d+),(\\d+)$").unwrap();
    for it in re.captures_iter(&content) {
        let &x1 = &it[2].parse::<usize>().unwrap();
        let &y1 = &it[3].parse::<usize>().unwrap();
        let &x2 = &it[4].parse::<usize>().unwrap();
        let &y2 = &it[5].parse::<usize>().unwrap();
        callback(&it[1], x1, y1, x2, y2);
    }
}

fn part1(content: &String) -> u32 {
    let rc_grid = Rc::new(RefCell::new(Grid::new(false)));
    let mut grid = rc_grid.borrow_mut();

    parse_with(&content, move |message, x1, y1, x2, y2| {
        match message {
            "turn on" => grid.slice_mut(x1, y1, x2, y2, |_| true),
            "turn off" => grid.slice_mut(x1, y1, x2, y2, |_| false),
            "toggle" => grid.slice_mut(x1, y1, x2, y2, |it| !it),
            _ => panic!("Unexpected input: {}", message),
        }
    });

    let grid = rc_grid.borrow();
    let mut count = 0;
    for x in 0..1000 {
        for y in 0..1000 {
            if grid.grid[x][y] {
                count += 1
            } 
        } 
    } 

    count
}

fn part2(content: &String) -> i32 {
    let rc_grid = Rc::new(RefCell::new(Grid::new(0)));
    let mut grid = rc_grid.borrow_mut();

    parse_with(&content, move |message, x1, y1, x2, y2| {
        match message {
            "turn on" => grid.slice_mut(x1, y1, x2, y2, |it| it + 1),
            "turn off" => grid.slice_mut(x1, y1, x2, y2, |it| if *it == 0 { 0 } else { it - 1 }),
            "toggle" => grid.slice_mut(x1, y1, x2, y2, |it| it + 2),
            _ => panic!("Unexpected input: {}", message),
        }
    });

    let grid = rc_grid.borrow();
    let mut count = 0;
    for x in 0..1000 {
        for y in 0..1000 {
            count += grid.grid[x][y];
        } 
    } 

    count
}

pub fn solve() -> std::io::Result<()> {
    let mut content = String::new();
    aoc2015::read_file("inputs/day-06.txt", &mut content)?;

    println!("Day 06");
    println!("Part 1: {}", part1(&content));
    println!("Part 2: {}", part2(&content));

    Ok(())
}

