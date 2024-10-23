use crate::utils;

struct GameOfLife {
    map: Vec<Vec<bool>>,
    update: Vec<Vec<bool>>,
    rows: usize,
    cols: usize,
    skip_corner: bool,
}

impl GameOfLife {
    fn new(rows: usize, cols: usize) -> Self {
        let map = vec![vec![false; cols]; rows];
        let update = vec![vec![false; cols]; rows];
        Self {
            map,
            update,
            rows,
            cols,
            skip_corner: false,
        }
    }

    fn is_corner(&self, row: usize, col: usize) -> bool {
        (row == 0 || row == self.rows - 1) && (col == 0 || col == self.cols - 1)
    }

    fn in_bounds(&self, row: i32, col: i32) -> bool {
        0 <= row && row < self.cols as i32 && 0 <= col && col < self.rows as i32
    }

    fn foreach_neighbors<F>(&self, row: usize, col: usize, mut f: F)
    where
        F: FnMut(bool),
    {
        let row = row as i32;
        let col = col as i32;
        for i in -1..=1 {
            for j in -1..=1 {
                let new_row = row + i;
                let new_col = col + j;
                if self.in_bounds(new_row, new_col) && (new_row != row || new_col != col) {
                    f(self.map[new_row as usize][new_col as usize]);
                }
            }
        }
    }

    fn turn_on_corners(&mut self) {
        self.skip_corner = true;
    
        self.map[0][0] = true;
        self.map[0][self.cols - 1] = true;
        self.map[self.rows - 1][0] = true;
        self.map[self.rows - 1][self.cols - 1] = true;
        
        self.update[0][0] = true;
        self.update[0][self.cols - 1] = true;
        self.update[self.rows - 1][0] = true;
        self.update[self.rows - 1][self.cols - 1] = true;
    }

    fn count_lights(&self) -> i32 {
        self.map.iter().flatten().filter(|&&x| x).count() as i32
    }

    fn rule(on: bool, neighbors_count: usize) -> bool {
        if on {
            neighbors_count == 2 || neighbors_count == 3
        } else {
            neighbors_count == 3
        }
    }

    fn steps(&mut self, n: usize) {
        for _ in 0..n {
            self.step();
        }
    }

    fn step(&mut self) {
        for row in 0..self.rows {
            for col in 0..self.cols {
                if self.skip_corner && self.is_corner(row, col) {
                    continue;
                }

                let mut neighbors_count = 0;
                self.foreach_neighbors(row, col, |nbor| {
                    neighbors_count += nbor as usize;
                });

                self.update[row][col] = GameOfLife::rule(self.map[row][col], neighbors_count);
            }
        }

        for i in 0..self.rows {
            for j in 0..self.cols {
                self.map[i][j] = self.update[i][j];
            }
        }
    }

    fn parse(content: &str, rows: usize, cols: usize) -> Self {
        let mut grid = GameOfLife::new(rows, cols);

        for (row, line) in content.lines().enumerate() {
            for (col, ch) in line.char_indices() {
                let on = ch == '#';
                grid.map[row][col] = on;
            }
        }

        grid
    }

    #[allow(unused)]
    fn print(&self) {
        for row in self.map.iter() {
            for on in row.iter() {
                if *on {
                    print!("#");
                } else {
                    print!(".");
                }
            }
            println!();
        }
        println!();
    }
}

fn part1(content: &str) -> i32 {
    let mut grid = GameOfLife::parse(content, 100, 100);
    grid.steps(100);
    grid.count_lights()
}

fn part2(content: &str) -> i32 {
    let mut grid = GameOfLife::parse(content, 100, 100);
    grid.turn_on_corners();
    grid.steps(100);
    grid.count_lights()
}

pub fn solve() -> std::io::Result<()> {
    let mut content = String::new();
    utils::read_file("inputs/day-18.txt", &mut content)?;

    println!("Day 18");
    time_it!("Part 1", part1(&content));
    time_it!("Part 2", part2(&content));

    Ok(())
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn step_rule1_must_works() {
        let content = ".#.#.#
...##.
#....#
..#...
#.#..#
####..";
        let mut grid = GameOfLife::parse(content, 6, 6);
        assert_eq!(15, grid.count_lights());
        grid.step();
        assert_eq!(11, grid.count_lights());
        grid.step();
        assert_eq!(8, grid.count_lights());
        grid.step();
        assert_eq!(4, grid.count_lights());
        grid.step();
        assert_eq!(4, grid.count_lights());
    }

    #[test]
    fn step_rule2_must_works() {
        let content = "##.#.#
...##.
#....#
..#...
#.#..#
####.#";
        let mut grid: GameOfLife = GameOfLife::parse(content, 6, 6);
        grid.turn_on_corners();
        assert_eq!((line!(), 17), (line!(), grid.count_lights()));
        grid.step();
        assert_eq!((line!(), 18), (line!(), grid.count_lights()));
        grid.step();
        assert_eq!((line!(), 18), (line!(), grid.count_lights()));
        grid.step();
        assert_eq!((line!(), 18), (line!(), grid.count_lights()));
        grid.step();
        assert_eq!((line!(), 14), (line!(), grid.count_lights()));
        grid.step();
        assert_eq!((line!(), 17), (line!(), grid.count_lights()));
    }
}
