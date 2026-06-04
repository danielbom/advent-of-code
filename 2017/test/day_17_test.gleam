import gleeunit
import gleeunit/should

import gleam/list

import day_17

pub fn main() -> Nil {
  gleeunit.main()
}

pub fn part1_test() {
  let inputs = [
    // 
    #("3", 638),
  ]

  list.map(inputs, fn(p) {
    let #(input, expected) = p
    let result = day_17.part1(input)
    should.equal(result, expected)
  })
}

pub fn part2_test() {
  let inputs = []

  list.map(inputs, fn(p) {
    let #(input, expected) = p
    let result = day_17.part2(input)
    should.equal(result, expected)
  })
}
