use crate::utils;
use std::collections::{HashMap, VecDeque};

#[derive(Debug)]
enum Atom<'a> {
    Key(&'a str),
    Value(u16),
}

#[derive(Debug)]
enum Kind1 {
    Assign,
    Not,
}

#[derive(Debug)]
struct Op1<'a> {
    kind: Kind1,
    value: Atom<'a>,
}

#[derive(Debug)]
enum Kind2 {
    And,
    Or,
    RShift,
    LShift,
}

#[derive(Debug)]
struct Op2<'a> {
    kind: Kind2,
    lhs: Atom<'a>,
    rhs: Atom<'a>,
}

#[derive(Debug)]
enum Op<'a> {
    Unary(Op1<'a>),
    Binary(Op2<'a>),
}

#[derive(Debug)]
struct Gate<'a> {
    op: Op<'a>,
    output: &'a str,
}

impl<'a> Atom<'a> {
    fn parse(atom: &'a str) -> Atom<'a> {
        match atom.parse() {
            Ok(value) => Atom::Value(value),
            Err(_) => Atom::Key(atom),
        }
    }

    fn key(&self) -> Option<&str> {
        match self {
            Atom::Key(key) => Some(key),
            _ => None,
        }
    }
}

impl Kind1 {
    fn resolve(&self, val: u16) -> u16 {
        match self {
            Kind1::Assign => val,
            Kind1::Not => !val,
        }
    }
}

impl Kind2 {
    fn resolve(&self, lhs: u16, rhs: u16) -> u16 {
        match self {
            Kind2::And => lhs & rhs,
            Kind2::Or => lhs | rhs,
            Kind2::RShift => lhs >> rhs,
            Kind2::LShift => lhs << rhs,
        }
    }

    fn parse(value: &str) -> Option<Kind2> {
        match value {
            "AND" => Some(Kind2::And),
            "OR" => Some(Kind2::Or),
            "RSHIFT" => Some(Kind2::RShift),
            "LSHIFT" => Some(Kind2::LShift),
            _ => None,
        }
    }
}

impl<'a> Op<'a> {
    fn parse(expr: &'a str) -> Option<Op<'a>> {
        let mut parts = expr.split(' ');
        match (parts.next(), parts.next(), parts.next()) {
            (Some(lhs), Some(op), Some(rhs)) => Some(Op::Binary(Op2 {
                kind: Kind2::parse(op)?,
                lhs: Atom::parse(lhs),
                rhs: Atom::parse(rhs),
            })),
            (Some("NOT"), Some(var), _) => Some(Op::Unary(Op1 {
                kind: Kind1::Not,
                value: Atom::parse(var),
            })),
            (Some(var), None, _) => Some(Op::Unary(Op1 {
                kind: Kind1::Assign,
                value: Atom::parse(var),
            })),
            _ => None,
        }
    }
}

impl<'a> Gate<'a> {
    fn parse(line: &'a str) -> Option<Gate<'a>> {
        let mut parts = line.split(" -> ");
        let expr = parts.next()?;
        let output = parts.next()?;
        let op = Op::parse(expr)?;
        Some(Gate { op, output })
    }
}

struct Machine<'a> {
    values: HashMap<&'a str, u16>,
    gates: HashMap<&'a str, Gate<'a>>,
}

impl<'a> Machine<'a> {
    fn new() -> Self {
        let values = HashMap::new();
        let gates = HashMap::new();
        Self { values, gates }
    }

    fn insert(&mut self, gate: Gate<'a>) {
        self.gates.insert(gate.output, gate);
    }

    fn parse(content: &'a str) -> Self {
        let machine = Machine::new();

        content
            .lines()
            .map(Gate::parse)
            .fold(machine, |mut machine, gate| {
                if let Some(gate) = gate {
                    machine.insert(gate);
                }
                machine
            })
    }

    fn get_atom(&self, atom: &Atom) -> Option<u16> {
        match atom {
            Atom::Value(value) => Some(*value),
            Atom::Key(key) => self.values.get(key).copied(),
        }
    }

    fn compute(&mut self, key: &str) -> Option<u16> {
        if let Some(value) = self.values.get(key) {
            return Some(*value);
        }

        let mut stack = VecDeque::new();
        stack.push_back(key);

        while let Some(key) = stack.pop_back() {
            let gate = self.gates.get(key)?;

            match &gate.op {
                Op::Unary(op) => match self.get_atom(&op.value) {
                    Some(value) => {
                        self.values.insert(gate.output, op.kind.resolve(value));
                    }
                    None => {
                        let key1 = op.value.key()?;
                        stack.push_back(key);
                        stack.push_back(key1);
                    }
                },
                Op::Binary(op) => match (self.get_atom(&op.lhs), self.get_atom(&op.rhs)) {
                    (Some(lhs), Some(rhs)) => {
                        self.values.insert(gate.output, op.kind.resolve(lhs, rhs));
                    }
                    _ => {
                        stack.push_back(key);
                        if let Some(key2) = op.rhs.key() {
                            stack.push_back(key2);
                        }
                        if let Some(key1) = op.lhs.key() {
                            stack.push_back(key1);
                        }
                    }
                },
            }
        }

        self.values.get(key).copied()
    }
}

fn part1(content: &str) -> u16 {
    let mut machine = Machine::parse(content);
    machine.compute("a").unwrap_or(0)
}

fn part2(content: &str) -> u16 {
    let part1_result = part1(content);
    let mut machine = Machine::parse(content);
    machine.insert(Gate {
        output: "b",
        op: Op::Unary(Op1 {
            kind: Kind1::Assign,
            value: Atom::Value(part1_result),
        }),
    });
    machine.compute("a").unwrap_or(0)
}

pub fn solve() -> std::io::Result<()> {
    let mut content = String::new();
    utils::read_file("inputs/day-07.txt", &mut content)?;

    println!("Day 07");
    time_it!("Part 1", part1(&content));
    time_it!("Part 2", part2(&content));

    Ok(())
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn base_test() {
        let mut expected = HashMap::new();
        expected.insert("d", 72);
        expected.insert("e", 507);
        expected.insert("f", 492);
        expected.insert("g", 114);
        expected.insert("h", 65412);
        expected.insert("i", 65079);
        expected.insert("x", 123);
        expected.insert("y", 456);

        let input = "\
        123 -> x\n\
        456 -> y\n\
        x AND y -> d\n\
        x OR y -> e\n\
        x LSHIFT 2 -> f\n\
        y RSHIFT 2 -> g\n\
        NOT x -> h\n\
        NOT y -> i";

        let mut machine = Machine::parse(input);
        let keys = machine
            .gates
            .keys()
            .map(|it| it.to_string())
            .collect::<Vec<String>>();
        for key in keys {
            machine.compute(&key);
        }

        let result = machine.values;

        assert_eq!(expected, result);
    }
}
