import gleam/bool
import gleam/dict.{type Dict}
import gleam/int
import gleam/io
import gleam/list
import gleam/string

import iv.{type Array}

import utils

type DanceMove {
  Spin(index: Int)
  Exchange(index_a: Int, index_b: Int)
  Partner(name_a: String, name_b: String)
}

fn parse_moves(s: String) -> List(DanceMove) {
  string.trim(s)
  |> string.split(on: ",")
  |> list.map(fn(item) {
    let value = string.drop_start(item, 1)
    use <- bool.lazy_guard(string.starts_with(item, "s"), fn() {
      let assert Ok(value) = int.parse(value)
      Spin(index: value)
    })
    use <- bool.lazy_guard(string.starts_with(item, "x"), fn() {
      let assert Ok(#(a, b)) = string.split_once(value, on: "/")
      let assert Ok(index_a) = int.parse(a)
      let assert Ok(index_b) = int.parse(b)
      Exchange(index_a:, index_b:)
    })
    use <- bool.lazy_guard(string.starts_with(item, "p"), fn() {
      let assert Ok(#(name_a, name_b)) = string.split_once(value, on: "/")
      Partner(name_a:, name_b:)
    })
    panic as string.append("invalid item: ", item)
  })
}

fn apply_move(programs: Array(String), move: DanceMove) -> Array(String) {
  case move {
    Spin(index) -> {
      let size = iv.size(programs)
      let #(init, tail) = iv.split(programs, size - index)
      iv.concat(tail, init)
    }
    Exchange(index_a, index_b) -> {
      let assert Ok(a) = iv.get(programs, index_a)
      let assert Ok(b) = iv.get(programs, index_b)
      programs
      |> iv.try_set(index_a, b)
      |> iv.try_set(index_b, a)
    }
    Partner(a, b) -> {
      let assert Ok(index_a) = iv.find_index(programs, fn(x) { x == a })
      let assert Ok(index_b) = iv.find_index(programs, fn(x) { x == b })
      programs
      |> iv.try_set(index_a, b)
      |> iv.try_set(index_b, a)
    }
  }
}

fn dance_loop(
  seen: Dict(Array(String), Bool),
  moves: List(DanceMove),
  programs: Array(String),
  count: Int,
  repeat: Int,
) {
  case repeat > 0, dict.has_key(seen, programs) {
    False, _ -> programs
    _, True -> dance_loop(dict.new(), moves, programs, 0, { repeat % count })
    _, False -> {
      let seen = dict.insert(seen, programs, True)
      let programs = list.fold(moves, programs, apply_move)
      dance_loop(seen, moves, programs, count + 1, repeat - 1)
    }
  }
}

fn dance(
  moves: List(DanceMove),
  programs: Array(String),
  repeat: Int,
) -> Array(String) {
  dance_loop(dict.new(), moves, programs, 0, repeat)
}

pub fn parse_and_dance(moves: String, input: String, repeat: Int) -> String {
  input
  |> string.to_graphemes()
  |> iv.from_list()
  |> dance(parse_moves(moves), _, repeat)
  |> iv.to_list()
  |> string.concat()
}

pub fn part1(s: String) {
  parse_and_dance(s, "abcdefghijklmnop", 1)
}

pub fn part2(s: String) {
  parse_and_dance(s, "abcdefghijklmnop", 1_000_000_000)
}

pub fn solve() {
  let input = utils.read_all_file("inputs/day-16.txt") |> string.trim()
  io.println("Day 16")
  utils.time_it("Part 1", fn() { part1(input) })
  utils.time_it("Part 2", fn() { part2(input) })
}
