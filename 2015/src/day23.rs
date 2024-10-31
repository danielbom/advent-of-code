use crate::utils;

#[derive(Debug)]
enum R {
    A,
    B,
}

impl R {
    fn parse(s: &str) -> Option<R> {
        if s.starts_with('a') {
            Some(R::A)
        } else if s.starts_with('b') {
            Some(R::B)
        } else {
            None
        }
    }
}

#[derive(Debug)]
enum I {
    Half(R),
    Triple(R),
    Increment(R),
    Jump(i64),
    JumpIfEven(R, i64),
    JumpIfOne(R, i64),
}

impl I {
    fn parse(line: &str) -> Option<I> {
        let mut parts = line.split(' ');
        let instruction = parts.next()?;
        match instruction {
            "hlf" => {
                let register = R::parse(parts.next()?)?;
                Some(I::Half(register))
            }
            "tpl" => {
                let register = R::parse(parts.next()?)?;
                Some(I::Triple(register))
            }
            "inc" => {
                let register = R::parse(parts.next()?)?;
                Some(I::Increment(register))
            }
            "jmp" => {
                let offset = parts.next()?.parse().ok()?;
                Some(I::Jump(offset))
            }
            "jie" => {
                let register = R::parse(parts.next()?)?;
                let offset = parts.next()?.parse().ok()?;
                Some(I::JumpIfEven(register, offset))
            }
            "jio" => {
                let register = R::parse(parts.next()?)?;
                let offset = parts.next()?.parse().ok()?;
                Some(I::JumpIfOne(register, offset))
            }
            _ => None,
        }
    }
}

fn parse_instructions(content: &str) -> Vec<I> {
    content.lines().filter_map(I::parse).collect()
}

struct Program {
    ra: i64,
    rb: i64,
    pc: i64,
}

impl Program {
    fn new() -> Self {
        Self {
            ra: 0,
            rb: 0,
            pc: 0,
        }
    }

    fn register(&self, r: &R) -> i64 {
        match r {
            R::A => self.ra,
            R::B => self.rb,
        }
    }

    fn register_mut(&mut self, r: &R) -> &mut i64 {
        match r {
            R::A => &mut self.ra,
            R::B => &mut self.rb,
        }
    }

    fn run(&mut self, instruction: &[I]) {
        while let Some(i) = instruction.get(self.pc as usize) {
            let mut offs = 1;

            match i {
                I::Half(r) => {
                    let r = self.register_mut(r);
                    *r /= 2;
                }
                I::Triple(r) => {
                    let r = self.register_mut(r);
                    *r *= 3;
                }
                I::Increment(r) => {
                    let r = self.register_mut(r);
                    *r += 1;
                }
                I::Jump(new_offs) => {
                    offs = *new_offs;
                }
                I::JumpIfEven(r, new_offs) => {
                    let r = self.register(r);
                    if r % 2 == 0 {
                        offs = *new_offs;
                    }
                }
                I::JumpIfOne(r, new_offs) => {
                    let r = self.register(r);
                    if r == 1 {
                        offs = *new_offs;
                    }
                }
            }

            self.pc += offs;
        }
    }
}

fn part1(content: &str) -> i64 {
    let instructions = parse_instructions(content);
    let mut p = Program::new();
    p.run(&instructions);
    p.rb
}

fn part2(content: &str) -> i64 {
    let instructions = parse_instructions(content);
    let mut p = Program::new();
    p.ra = 1;
    p.run(&instructions);
    p.rb
}

pub fn solve() -> std::io::Result<()> {
    let mut content = String::new();
    utils::read_file("inputs/day-23.txt", &mut content)?;

    println!("Day 23");
    time_it!("Part 1", part1(&content));
    time_it!("Part 2", part2(&content));

    Ok(())
}
