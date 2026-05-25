import gleeunit
import gleeunit/should

import gleam/list

import day_07

pub fn main() -> Nil {
  gleeunit.main()
}

pub fn part1_test() {
  let input =
    "pbga (66)
xhth (57)
ebii (61)
havc (66)
ktlj (57)
fwft (72) -> ktlj, cntj, xhth
qoyq (66)
padx (45) -> pbga, havc, qoyq
tknk (41) -> ugml, padx, fwft
jptl (61)
ugml (68) -> gyxo, ebii, jptl
gyxo (61)
cntj (57)"
  let inputs = [#(input, "tknk")]

  list.map(inputs, fn(p) {
    let #(input, expected) = p
    let result = day_07.part1(input)
    should.equal(result, expected)
  })
}

pub fn part2_test() {
  let input =
    "pbga (66)
xhth (57)
ebii (61)
havc (66)
ktlj (57)
fwft (72) -> ktlj, cntj, xhth
qoyq (66)
padx (45) -> pbga, havc, qoyq
tknk (41) -> ugml, padx, fwft
jptl (61)
ugml (68) -> gyxo, ebii, jptl
gyxo (61)
cntj (57)"
  let inputs = [#(input, 60)]

  list.map(inputs, fn(p) {
    let #(input, expected) = p
    let result = day_07.part2(input)
    should.equal(result, expected)
  })
}
