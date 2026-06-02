import gleeunit
import gleeunit/should

import gleam/list

import day_15

pub fn main() -> Nil {
  gleeunit.main()
}

pub fn part1_test() {
  // Too timing consuming - uncomment to test it
  let inputs = [
    // 
  // #("a with 65\nb with 8921", 588),
  ]

  list.map(inputs, fn(p) {
    let #(input, expected) = p
    let result = day_15.part1(input)
    should.equal(result, expected)
  })
}

pub fn part2_test() {
  // Too timing consuming - uncomment to test it
  let inputs = [
    // 
  // #("a with 65\nb with 8921", 309),
  ]

  list.map(inputs, fn(p) {
    let #(input, expected) = p
    let result = day_15.part2(input)
    should.equal(result, expected)
  })
}
