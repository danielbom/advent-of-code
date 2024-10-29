use crate::utils;

#[derive(Debug, Clone, Eq, PartialEq)]
struct SheetIterator {
    row: u32,
    col: u32,
    ix: usize,
}

impl SheetIterator {
    fn new() -> Self {
        Self::with(1, 1, 1)
    }

    fn with(row: u32, col: u32, ix: usize) -> Self {
        Self { row, col, ix }
    }
}

impl Iterator for SheetIterator {
    type Item = Self;

    fn next(&mut self) -> Option<Self::Item> {
        let next = self.clone();

        self.row -= 1;
        self.col += 1;
        self.ix += 1;

        if self.row == 0 {
            self.row = self.col;
            self.col = 1;
        }
        
        Some(next)
    }
}

fn parse_input(content: &str) -> Option<(u32, u32)> {
    let row = content.rfind("row ")?;
    let row = content[row + "row ".len()..]
        .chars()
        .take_while(|ch| ch.is_ascii_digit())
        .collect::<String>();
    let col = content.rfind("column ")?;
    let col = content[col + "column ".len()..]
        .chars()
        .take_while(|ch| ch.is_ascii_digit())
        .collect::<String>();
    let row = row.parse().ok()?;
    let col = col.parse().ok()?;
    Some((row, col))
}

fn find_code_at(row: u32, col: u32) -> u64 {
    SheetIterator::new()
        .take_while(|s| !(s.row == row && s.col == col))
        .fold(20151125, |x, _| (x * 252533) % 33554393)
}

fn part1(content: &str) -> u64 {
    let (row, col) = parse_input(content).expect("invalid input");
    find_code_at(row, col)
}

pub fn solve() -> std::io::Result<()> {
    let mut content = String::new();
    utils::read_file("inputs/day-25.txt", &mut content)?;

    println!("Day 25");
    time_it!("Part 1", part1(&content));

    Ok(())
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_input() {
        let input = "Some text before. row 123 and column 312";
        let expected = (123, 312);
        let result = parse_input(input);
        assert!(result.is_some());
        let result = result.unwrap();
        assert_eq!(expected, result);
    }

    #[test]
    fn test_sheet_iterator() {
        let mut it = SheetIterator::new();

        assert_eq!(it.next(), Some(SheetIterator::with(1, 1, 1)));
        assert_eq!(it.next(), Some(SheetIterator::with(2, 1, 2)));
        assert_eq!(it.next(), Some(SheetIterator::with(1, 2, 3)));
        assert_eq!(it.next(), Some(SheetIterator::with(3, 1, 4)));
    }

    #[test]
    fn test_find_code_at() {
        let table = vec![
            vec![20151125, 18749137, 17289845, 30943339, 10071777, 33511524],
            vec![31916031, 21629792, 16929656, 7726640, 15514188, 4041754],
            vec![16080970, 8057251, 1601130, 7981243, 11661866, 16474243],
            vec![24592653, 32451966, 21345942, 9380097, 10600672, 31527494],
            vec![77061, 17552253, 28094349, 6899651, 9250759, 31663883],
            vec![33071741, 6796745, 25397450, 24659492, 1534922, 27995004],
        ];

        assert_eq!(find_code_at(1, 1), 20151125);

        for (i, row) in table.iter().enumerate() {
            for (j, x) in row.iter().enumerate() {
                let i = i as u32 + 1;
                let j = j as u32 + 1;
                assert_eq!((i, j, find_code_at(i, j)), (i, j, *x));
            }
        }
    }
}
