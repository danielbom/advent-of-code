import gleeunit
import gleeunit/should

import gleam/list

import day_13

pub fn main() -> Nil {
  gleeunit.main()
}

pub fn part1_test() {
  let input =
    "0: 3
1: 2
4: 4
6: 4"
  let inputs = [
    //
    #(input, 24),
  ]

  list.map(inputs, fn(p) {
    let #(input, expected) = p
    let result = day_13.part1(input)
    should.equal(result, expected)
  })
}

pub fn part2_test() {
  let input =
    "0: 3
1: 2
4: 4
6: 4"
  let inputs = [
    //
    #(input, 10),
  ]

  list.map(inputs, fn(p) {
    let #(input, expected) = p
    let result = day_13.part2(input)
    should.equal(result, expected)
  })
}
