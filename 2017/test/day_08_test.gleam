import gleeunit
import gleeunit/should

import gleam/list

import day_08

pub fn main() -> Nil {
  gleeunit.main()
}

pub fn part1_test() {
  let input =
    "b inc 5 if a > 1
a inc 1 if b < 5
c dec -10 if a >= 1
c inc -20 if c == 10"
  let inputs = [#(input, 1)]

  list.map(inputs, fn(p) {
    let #(input, expected) = p
    let result = day_08.part1(input)
    should.equal(result, expected)
  })
}

pub fn part2_test() {
  let input =
    "b inc 5 if a > 1
a inc 1 if b < 5
c dec -10 if a >= 1
c inc -20 if c == 10"
  let inputs = [#(input, 10)]

  list.map(inputs, fn(p) {
    let #(input, expected) = p
    let result = day_08.part2(input)
    should.equal(result, expected)
  })
}
