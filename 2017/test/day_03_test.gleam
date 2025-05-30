import gleeunit
import gleeunit/should

import gleam/list

import day_03

pub fn main() -> Nil {
  gleeunit.main()
}

pub fn part1_test() {
  let inputs = [
    //
    #("1", 0),
    #("12", 3),
    #("23", 2),
    #("1024", 31),
  ]

  list.map(inputs, fn(p) {
    let #(input, expected) = p
    let result = day_03.part1(input)
    should.equal(result, expected)
  })
}

pub fn part2_test() {
  let inputs = [
    //
    #("1", 2),
    #("12", 23),
    #("806", 880),
    #("1024", 1968),
  ]

  list.map(inputs, fn(p) {
    let #(input, expected) = p
    let result = day_03.part2(input)
    should.equal(result, expected)
  })
}
