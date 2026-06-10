import gleeunit

import gleam/list

import day_16

pub fn main() -> Nil {
  gleeunit.main()
}

pub fn parse_and_dance_test() {
  let inputs = [
    //
    #("s1,x3/4,pe/b", 1, "baedc"),
    #("s1,x3/4,pe/b", 2, "ceadb"),
    #("s1,x3/4,pe/b", 3, "ecbda"),
    #("s1,x3/4,pe/b", 4, "abcde"),
    // repeat
    #("s1,x3/4,pe/b", 5, "baedc"),
    #("s1,x3/4,pe/b", 6, "ceadb"),
    #("s1,x3/4,pe/b", 7, "ecbda"),
    #("s1,x3/4,pe/b", 8, "abcde"),
  ]

  list.map(inputs, fn(p) {
    let #(input, repeat, expected) = p
    let result = day_16.parse_and_dance(input, "abcde", repeat)
    assert result == expected
  })
}
