use crate::utils;

mod factors {
    pub struct FactorsNaive {
        i: usize,
        value: usize,
    }

    #[allow(unused)]
    impl FactorsNaive {
        pub fn new(value: usize) -> Self {
            Self { i: 1, value }
        }
    }

    impl Iterator for FactorsNaive {
        type Item = usize;

        fn next(&mut self) -> Option<Self::Item> {
            if self.i > self.value {
                return None;
            }

            while self.i < self.value {
                if self.value % self.i == 0 {
                    let result = self.value / self.i;
                    self.i += 1;
                    return Some(result);
                }
                self.i += 1;
            }

            if self.i <= self.value {
                self.i = self.value + 1;
                return Some(1);
            }

            return None;
        }
    }

    pub struct FactorsOptimized {
        i: usize,
        value: usize,
        factor: bool,
    }

    impl FactorsOptimized {
        #[allow(unused)]
        pub fn new(value: usize) -> Self {
            Self {
                i: 1,
                value,
                factor: true,
            }
        }
    }

    impl Iterator for FactorsOptimized {
        type Item = usize;

        fn next(&mut self) -> Option<Self::Item> {
            if self.i > self.value {
                return None;
            }

            if self.value == 1 {
                self.i = 2;
                return Some(1);
            }

            while self.i * self.i < self.value {
                if self.value % self.i == 0 {
                    if self.factor {
                        self.factor = false;
                        return Some(self.i);
                    }

                    self.factor = true;
                    let result = self.value / self.i;
                    self.i += 1;
                    return Some(result);
                }
                self.i += 1;
            }

            if self.i * self.i == self.value {
                let i = self.i;
                self.i = self.value + 1;
                return Some(i);
            }

            return None;
        }
    }

    pub type Factors = FactorsOptimized;
}

fn parse_input(content: &str) -> usize {
    content.trim().parse().unwrap()
}

mod part1 {
    use super::factors::Factors;
    use super::parse_input;

    pub const PRESENTS_PER_HOUSE: usize = 10;

    pub fn compute_presents(house: usize) -> usize {
        Factors::new(house).into_iter().sum::<usize>() * PRESENTS_PER_HOUSE
    }

    pub fn closest_house(presents: usize) -> usize {
        let mut house = 1;
        loop {
            let house_presents = compute_presents(house);
            let distance = presents as i64 - house_presents as i64;
            if distance <= 0 {
                return house;
            }

            if house < 10 {
                house += 1;
            } else if house < 66 {
                house += 2;
            } else {
                house += 6;
            }
        }
    }

    pub fn solve(content: &str) -> usize {
        let presents = parse_input(content);
        closest_house(presents)
    }
}

mod part2 {
    use super::factors::Factors;
    use super::parse_input;

    pub const PRESENTS_PER_HOUSE: usize = 11;
    pub const HOUSES_LIMIT: usize = 50;

    pub fn compute_presents(house: usize) -> usize {
        Factors::new(house)
            .into_iter()
            .filter(|&x| x >= house / HOUSES_LIMIT)
            .sum::<usize>()
            * PRESENTS_PER_HOUSE
    }

    pub fn closest_house(presents: usize) -> usize {
        let mut house = 1;
        loop {
            let house_presents = compute_presents(house);
            let distance = presents as i64 - house_presents as i64;
            if distance <= 0 {
                return house;
            }

            if house < 10 {
                house += 1;
            } else if house < 66 {
                house += 2;
            } else {
                house += 6;
            }
        }
    }

    pub fn solve(content: &str) -> usize {
        let presents = parse_input(content);
        // 718200
        closest_house(presents)
    }
}

/*
House : 1  2  3  4  5  6  7  8  9  10
Elf 1 : 1  1  1  1  1  1  1  1  1  1  ...
ELf 2 :    2     2     2     2     2  ...
ELf 3 :       3        3        3     ...
Elf 4 :          4           4        ...
Elf 5 :             5              5  ...
Elf 6 :                6              ...
Elf 7 :                   7           ...
Elf 8 :                      8        ...
Elf 9 :                         9     ...
Elf 10:                            10 ...
Prs   : 1  3  4  7  6  12 8  15 12 18 ...

*/

pub fn solve() -> std::io::Result<()> {
    let mut content = String::new();
    utils::read_file("inputs/day-20.txt", &mut content)?;

    println!("Day 20");
    time_it!("Part 1", part1::solve(&content));
    time_it!("Part 2", part2::solve(&content));

    Ok(())
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1_compute_presents() {
        let expected = vec![10, 30, 40, 70, 60, 120, 80, 150, 130];
        let result = (1..=9)
            .into_iter()
            .map(part1::compute_presents)
            .collect::<Vec<_>>();
        assert_eq!(expected, result);
    }

    #[test]
    fn test_part1_closest_house() {
        let expected = vec![1, 2, 3, 4, 4, 6, 6, 8, 8];
        let result = vec![10, 30, 40, 70, 60, 120, 80, 150, 130]
            .into_iter()
            .map(part1::closest_house)
            .collect::<Vec<_>>();
        assert_eq!(expected, result);
    }

    #[test]
    fn test_part2_closest_house() {
        let expected = vec![1, 2, 3, 4, 4, 6, 6, 8, 6];
        let result = vec![10, 30, 40, 70, 60, 120, 80, 150, 130]
            .into_iter()
            .map(part2::closest_house)
            .collect::<Vec<_>>();
        assert_eq!(expected, result);
    }

    pub fn part1_compute_factors(house: usize) -> Vec<usize> {
        let mut fs = super::factors::Factors::new(house)
            .into_iter()
            .collect::<Vec<usize>>();
        fs.sort();
        fs
    }

    pub fn part2_compute_factors(house: usize) -> Vec<usize> {
        let mut fs = super::factors::Factors::new(house)
            .into_iter()
            .filter(|&x| x >= house / part2::HOUSES_LIMIT)
            .collect::<Vec<usize>>();
        fs.sort();
        fs
    }

    #[test]
    fn test_compute_factors() {
        let fs1 = part1_compute_factors(500);
        let fs2 = part2_compute_factors(500);
        assert_eq!(fs1, vec![1, 2, 4, 5, 10, 20, 25, 50, 100, 125, 250, 500]);
        assert_eq!(fs2, vec![            10, 20, 25, 50, 100, 125, 250, 500]);
    }

    #[test]
    fn test_divisors() {
        use super::factors::{FactorsNaive, FactorsOptimized};

        for i in 1..10000 {
            let mut s1 = FactorsNaive::new(i).collect::<Vec<usize>>();
            let mut s2 = FactorsOptimized::new(i).collect::<Vec<usize>>();
            s1.sort();
            s2.sort();
            assert_eq!(s1, s2, "Divisors {}", i);
        }
    }
}
