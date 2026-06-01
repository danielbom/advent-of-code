import gleeunit
import gleeunit/should

import gleam/list

import day_09

pub fn main() -> Nil {
  gleeunit.main()
}

pub fn part1_test() {
  let inputs = [
    // {}, score of 1.
    #("{}", 1),
    // {{{}}}, score of 1 + 2 + 3 = 6.
    #("{{{}}}", 6),
    // {{},{}}, score of 1 + 2 + 2 = 5.
    #("{{},{}}", 5),
    // {{{},{},{{}}}}, score of 1 + 2 + 3 + 3 + 3 + 4 = 16.
    #("{{{},{},{{}}}}", 16),
    // {<a>,<a>,<a>,<a>}, score of 1.
    #("{<a>,<a>,<a>,<a>}", 1),
    // {{<ab>},{<ab>},{<ab>},{<ab>}}, score of 1 + 2 + 2 + 2 + 2 = 9.
    #("{{<ab>},{<ab>},{<ab>},{<ab>}}", 9),
    // {{<!!>},{<!!>},{<!!>},{<!!>}}, score of 1 + 2 + 2 + 2 + 2 = 9.
    #("{{<!!>},{<!!>},{<!!>},{<!!>}}", 9),
    // {{<a!>},{<a!>},{<a!>},{<ab>}}, score of 1 + 2 = 3.
    #("{{<a!>},{<a!>},{<a!>},{<ab>}}", 3),
  ]

  list.map(inputs, fn(p) {
    let #(input, expected) = p
    let result = day_09.part1(input)
    should.equal(result, expected)
  })
}

pub fn part2_test() {
  let inputs = [
    // <>, 0 characters.
    #("<>", 0),
    // <random characters>, 17 characters.
    #("<random characters>", 17),
    // <<<<>, 3 characters.
    #("<<<<>", 3),
    // <{!>}>, 2 characters.
    #("<{!>}>", 2),
    // <!!>, 0 characters.
    #("<!!>", 0),
    // <!!!>>, 0 characters.
    #("<!!!>>", 0),
    // <{o"i!a,<{i<a>, 10 characters.
    #("<{o\"i!a,<{i<a>", 10),
  ]

  list.map(inputs, fn(p) {
    let #(input, expected) = p
    let result = day_09.part2(input)
    should.equal(result, expected)
  })
}
