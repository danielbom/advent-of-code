import gleeunit
import gleeunit/should

import gleam/list

import day_05

pub fn main() -> Nil {
  gleeunit.main()
}

pub fn part1_test() {
  let input = "0\n3\n0\n1\n-3"
  let inputs = [#(input, 5)]

  list.map(inputs, fn(p) {
    let #(input, expected) = p
    let result = day_05.part1(input)
    should.equal(result, expected)
  })
}

pub fn part2_test() {
  let input = "0\n3\n0\n1\n-3"
  let inputs = [#(input, 10)]

  list.map(inputs, fn(p) {
    let #(input, expected) = p
    let result = day_05.part2(input)
    should.equal(result, expected)
  })
}
