use std::collections::HashSet;

trait Rule {
    fn reset(&mut self);
    fn is_valid(&self) -> bool;
    fn feed(&mut self, ch: char);
}

struct HasForbidSeq {
    last: Option<char>,
    sequences: Vec<String>,
    valid: bool,
}

impl HasForbidSeq {
    fn new() -> Self {
        let sequences = vec!["ab", "cd", "pq", "xy"];
        let sequences = sequences.iter().map(|it| it.to_string()).collect();
        Self {
            last: None,
            valid: true,
            sequences,
        }
    }
}

impl Rule for HasForbidSeq {
    fn reset(&mut self) {
        self.last = None;
        self.valid = true;
    }
    fn is_valid(&self) -> bool {
        self.valid
    }
    fn feed(&mut self, ch: char) {
        if !self.is_valid() {
            return;
        }

        if let Some(lch) = self.last {
            let seq = vec![lch, ch];
            let seq = String::from_iter(seq);
            self.valid = !self.sequences.contains(&seq);
        }
        self.last = Some(ch);
    }
}

struct Has3Vowels {
    vowels: String,
    count: u32,
}

impl Has3Vowels {
    fn new() -> Self {
        let vowels = "aeiou".to_string();
        Self { vowels, count: 0 }
    }
}

impl Rule for Has3Vowels {
    fn reset(&mut self) {
        self.count = 0;
    }
    fn is_valid(&self) -> bool {
        self.count >= 3
    }
    fn feed(&mut self, ch: char) {
        if self.is_valid() {
            return;
        }

        if self.vowels.contains(ch) {
            self.count += 1;
        }
    }
}

struct HasSeq2 {
    last: Option<char>,
    valid: bool,
}

impl HasSeq2 {
    fn new() -> Self {
        Self {
            last: None,
            valid: false,
        }
    }
}

impl Rule for HasSeq2 {
    fn reset(&mut self) {
        self.last = None;
        self.valid = false;
    }
    fn is_valid(&self) -> bool {
        self.valid
    }
    fn feed(&mut self, ch: char) {
        if self.is_valid() {
            return;
        }

        if let Some(lch) = self.last {
            self.valid = lch == ch;
        }

        self.last = Some(ch);
    }
}

struct HasSeq2With1Gap {
    last2: Option<char>,
    last: Option<char>,
    valid: bool,
}

impl HasSeq2With1Gap {
    fn new() -> Self {
        Self {
            last2: None,
            last: None,
            valid: false,
        }
    }
}

impl Rule for HasSeq2With1Gap {
    fn reset(&mut self) {
        self.last2 = None;
        self.last = None;
        self.valid = false;
    }
    fn is_valid(&self) -> bool {
        self.valid
    }
    fn feed(&mut self, ch: char) {
        if self.is_valid() {
            return;
        }

        if self.last.is_some() {
            if let Some(l2ch) = self.last2 {
                self.valid = l2ch == ch;
            }
        }

        self.last2 = self.last;
        self.last = Some(ch);
    }
}

struct ContainsTwicePair {
    valid: bool,
    last: Option<char>,
    last_pair: Option<(char, char)>,
    last_pairs: HashSet<(char, char)>,
}

impl ContainsTwicePair {
    fn new() -> Self {
        Self {
            valid: false,
            last: None,
            last_pair: None,
            last_pairs: HashSet::new(),
        }
    }
}

impl Rule for ContainsTwicePair {
    fn reset(&mut self) {
        self.valid = false;
        self.last = None;
        self.last_pair = None;
        self.last_pairs.clear();
    }
    fn is_valid(&self) -> bool {
        self.valid
    }
    fn feed(&mut self, ch: char) {
        if self.is_valid() {
            return;
        }

        if let Some(lch) = self.last {
            let pair = (lch, ch);
            self.valid = self.last_pairs.contains(&pair);

            if let Some(lpair) = self.last_pair {
                self.last_pairs.insert(lpair);
            }

            self.last_pair = Some(pair);
        }

        self.last = Some(ch);
    }
}

struct IsNice {
    rules: Vec<Box<dyn Rule>>,
}

impl IsNice {
    fn new1() -> Self {
        Self {
            rules: vec![
                Box::new(HasSeq2::new()),
                Box::new(Has3Vowels::new()),
                Box::new(HasForbidSeq::new()),
            ],
        }
    }

    fn new2() -> Self {
        Self {
            rules: vec![
                Box::new(HasSeq2With1Gap::new()),
                Box::new(ContainsTwicePair::new()),
            ],
        }
    }

    fn validate(&mut self, content: &str) -> bool {
        self.rules.iter_mut().for_each(|it| it.reset());

        for ch in content.chars() {
            self.rules.iter_mut().for_each(|it| it.feed(ch));
        }

        self.rules.iter().fold(true, |acc, it| acc && it.is_valid())
    }
}

fn compute(content: &String, is_nice: &mut IsNice) -> u32 {
    content.lines().filter(|it| is_nice.validate(it)).count() as u32
}

fn part1(content: &String) -> u32 {
    compute(content, &mut IsNice::new1())
}

fn part2(content: &String) -> u32 {
    compute(content, &mut IsNice::new2())
}

pub fn solve() -> std::io::Result<()> {
    let mut content = String::new();
    aoc2015::read_file("inputs/day-05.txt", &mut content)?;

    println!("Day 05");
    println!("Part 1: {}", part1(&content));
    println!("Part 2: {}", part2(&content));

    Ok(())
}
