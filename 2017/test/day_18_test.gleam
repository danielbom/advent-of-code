import gleeunit

import gleam/list

import day_18

pub fn main() -> Nil {
  gleeunit.main()
}

pub fn part1_test() {
  let input =
    "set a 1
add a 2
mul a a
mod a 5
snd a
set a 0
rcv a
jgz a -1
set a 1
jgz a -2"
  let inputs = [#(input, 4)]

  list.map(inputs, fn(p) {
    let #(input, expected) = p
    let result = day_18.part1(input)
    assert result == expected
  })
}

pub fn part2_test() {
  let input =
    "snd 1
snd 2
snd p
rcv a
rcv b
rcv c
rcv d"
  let inputs = [#(input, 3)]

  list.map(inputs, fn(p) {
    let #(input, expected) = p
    let result = day_18.part2(input)
    assert result == expected
  })
}
