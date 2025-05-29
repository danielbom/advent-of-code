import gleeunit
import gleeunit/should

import gleam/list
import gleam/pair

import day_01

pub fn main() -> Nil {
  gleeunit.main()
}

pub fn part1_test() {
  let inputs = [#("1122", 3), #("1111", 4), #("1234", 0), #("91212129", 9)]

  list.map(inputs, fn(p) {
    let input = pair.first(p)
    let expected = pair.second(p)
    day_01.part1(input)
    |> should.equal(expected)
  })
}

pub fn part2_test() {
  let inputs = [
    #("1212", 6),
    #("1221", 0),
    #("123425", 4),
    #("123123", 12),
    #("12131415", 4),
  ]

  list.map(inputs, fn(p) {
    let input = pair.first(p)
    let expected = pair.second(p)
    day_01.part2(input)
    |> should.equal(expected)
  })
}
