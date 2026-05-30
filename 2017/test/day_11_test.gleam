import gleeunit
import gleeunit/should

import gleam/list

import day_11

pub fn main() -> Nil {
  gleeunit.main()
}

pub fn part1_test() {
  let inputs = [
    //
    #("", 0),
    #("s,s,s", 3),
    #("n,n,n", 3),
    #("se,se,se", 3),
    #("sw,sw,sw", 3),
    #("ne,ne,ne", 3),
    #("nw,nw,nw", 3),
    #("n,n,s,s", 0),
    #("ne,ne,sw,sw", 0),
    #("nw,nw,se,se", 0),
    #("ne,ne,s,s", 2),
    #("nw,nw,s,s", 2),
    #("se,se,n,n", 2),
    #("sw,sw,n,n", 2),
    #("se,sw,se,sw,se,sw", 3),
    #("ne,nw,ne,nw,ne,nw", 3),
    #("se,sw,se,sw,sw", 3),
    #("ne,ne,ne,nw,nw,nw", 3),
    #("s,s,s,n,n,n,se,sw,se,sw,se,sw", 3),
    #("s,s,s,s,n,n,n,se,sw,se,sw,se,sw", 4),
    #("n,s,s,s,n,n,n,se,sw,se,sw,se,sw", 4),
    #("sw,s,s,s,n,n,n,se,sw,se,sw,se,sw", 4),
    #("se,s,s,s,n,n,n,se,sw,se,sw,se,sw", 4),
    #("nw,s,s,s,n,n,n,se,sw,se,sw,se,sw", 3),
    #("ne,s,s,s,n,n,n,se,sw,se,sw,se,sw", 3),
    #("s,s,s,n,n,n,ne,nw,ne,nw,ne,nw", 3),
    #("s,s,s,s,n,n,n,ne,nw,ne,nw,ne,nw", 4),
    #("n,s,s,s,n,n,n,ne,nw,ne,nw,ne,nw", 4),
    #("sw,s,s,s,n,n,n,ne,nw,ne,nw,ne,nw", 3),
    #("se,s,s,s,n,n,n,ne,nw,ne,nw,ne,nw", 3),
    #("nw,s,s,s,n,n,n,ne,nw,ne,nw,ne,nw", 4),
    #("ne,s,s,s,n,n,n,ne,nw,ne,nw,ne,nw", 4),
    #("n,s,ne,nw,se,sw", 0),
    #("n,n,s,ne,nw,se,sw", 1),
    #("s,n,s,ne,nw,se,sw", 1),
    #("ne,n,s,ne,nw,se,sw", 1),
    #("se,n,s,ne,nw,se,sw", 1),
    #("nw,n,s,ne,nw,se,sw", 1),
    #("sw,n,s,ne,nw,se,sw", 1),
  ]

  list.map(inputs, fn(p) {
    let #(input, expected) = p
    let result = day_11.part1(input)
    should.equal(result, expected)
  })
}

pub fn part2_test() {
  let inputs = [
    //
    #("", 0),
    #("s,s,s", 3),
    #("n,n,n", 3),
    #("se,se,se", 3),
    #("sw,sw,sw", 3),
    #("ne,ne,ne", 3),
    #("nw,nw,nw", 3),
    #("n,n,s,s", 2),
    #("ne,ne,sw,sw", 2),
    #("nw,nw,se,se", 2),
    #("ne,ne,s,s", 2),
    #("nw,nw,s,s", 2),
    #("se,se,n,n", 2),
    #("sw,sw,n,n", 2),
    #("se,sw,se,sw,se,sw", 3),
    #("ne,nw,ne,nw,ne,nw", 3),
    #("se,sw,se,sw,sw", 3),
    #("ne,ne,ne,nw,nw,nw", 3),
    #("s,s,s,n,n,n,se,sw,se,sw,se,sw", 3),
    #("s,s,s,s,n,n,n,se,sw,se,sw,se,sw", 4),
    #("n,s,s,s,n,n,n,se,sw,se,sw,se,sw", 4),
    #("sw,s,s,s,n,n,n,se,sw,se,sw,se,sw", 4),
    #("se,s,s,s,n,n,n,se,sw,se,sw,se,sw", 4),
    #("nw,s,s,s,n,n,n,se,sw,se,sw,se,sw", 3),
    #("ne,s,s,s,n,n,n,se,sw,se,sw,se,sw", 3),
    #("s,s,s,n,n,n,ne,nw,ne,nw,ne,nw", 3),
    #("s,s,s,s,n,n,n,ne,nw,ne,nw,ne,nw", 4),
    #("n,s,s,s,n,n,n,ne,nw,ne,nw,ne,nw", 4),
    #("sw,s,s,s,n,n,n,ne,nw,ne,nw,ne,nw", 3),
    #("se,s,s,s,n,n,n,ne,nw,ne,nw,ne,nw", 3),
    #("nw,s,s,s,n,n,n,ne,nw,ne,nw,ne,nw", 4),
    #("ne,s,s,s,n,n,n,ne,nw,ne,nw,ne,nw", 4),
    #("n,s,ne,nw,se,sw", 1),
    #("n,n,s,ne,nw,se,sw", 2),
    #("s,n,s,ne,nw,se,sw", 2),
    #("ne,n,s,ne,nw,se,sw", 2),
    #("se,n,s,ne,nw,se,sw", 1),
    #("nw,n,s,ne,nw,se,sw", 2),
    #("sw,n,s,ne,nw,se,sw", 1),
  ]

  list.map(inputs, fn(p) {
    let #(input, expected) = p
    let result = day_11.part2(input)
    should.equal(result, expected)
  })
}
