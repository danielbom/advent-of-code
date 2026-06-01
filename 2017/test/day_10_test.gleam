import gleeunit
import gleeunit/should

import gleam/list

import day_10

pub fn main() -> Nil {
  gleeunit.main()
}

pub fn knot_hash_check_test() {
  let inputs = [
    //
    #([3, 4, 1, 5], 5, 12),
  ]

  list.map(inputs, fn(p) {
    let #(input, size, expected) = p
    let result = day_10.knot_hash_check(input, size)
    should.equal(result, expected)
  })
}

pub fn part2_test() {
  let inputs = [
    #("", "a2582a3a0e66e6e86e3812dcb672a272"),
    #("AoC 2017", "33efeb34ea91902bb2f59c9920caa6cd"),
    #("1,2,3", "3efbe78a8d82f29979031a4aa0b16a9d"),
    #("1,2,4", "63960835bcdc130f0b66d7ff4f6a5a8e"),
  ]

  list.map(inputs, fn(p) {
    let #(input, expected) = p
    let result = day_10.part2(input)
    should.equal(result, expected)
  })
}
