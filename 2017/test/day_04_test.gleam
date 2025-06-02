import gleeunit
import gleeunit/should

import gleam/list

import day_04

pub fn main() -> Nil {
  gleeunit.main()
}

pub fn part1_test() {
  let input =
    "aa bb cc dd ee
aa bb cc dd aa
aa bb cc dd aaa"
  let inputs = [#(input, 2)]

  list.map(inputs, fn(p) {
    let #(input, expected) = p
    let result = day_04.part1(input)
    should.equal(result, expected)
  })
}

pub fn part2_test() {
  let input =
    "abcde fghij
abcde xyz ecdab
a ab abc abd abf abj
iiii oiii ooii oooi oooo
oiii ioii iioi iiio"
  let inputs = [#(input, 3)]

  list.map(inputs, fn(p) {
    let #(input, expected) = p
    let result = day_04.part2(input)
    should.equal(result, expected)
  })
}
