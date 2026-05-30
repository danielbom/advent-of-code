import gleeunit
import gleeunit/should

import gleam/list

import day_12

pub fn main() -> Nil {
  gleeunit.main()
}

pub fn part1_test() {
  let input =
    "0 <-> 2
1 <-> 1
2 <-> 0, 3, 4
3 <-> 2, 4
4 <-> 2, 3, 6
5 <-> 6
6 <-> 4, 5"
  let inputs = [
    //
    #(input, 6),
  ]

  list.map(inputs, fn(p) {
    let #(input, expected) = p
    let result = day_12.part1(input)
    should.equal(result, expected)
  })
}

pub fn part2_test() {
  let input =
    "0 <-> 2
1 <-> 1
2 <-> 0, 3, 4
3 <-> 2, 4
4 <-> 2, 3, 6
5 <-> 6
6 <-> 4, 5"
  let inputs = [
    //
    #(input, 2),
  ]

  list.map(inputs, fn(p) {
    let #(input, expected) = p
    let result = day_12.part2(input)
    should.equal(result, expected)
  })
}
