import gleam/dict.{type Dict}
import gleam/int
import gleam/io
import gleam/list
import gleam/string

import iv.{type Array}

import utils

type DanceMove {
  Spin(Int)
  Exchange(Int, Int)
  Partner(String, String)
}

fn parse_moves(s: String) -> List(DanceMove) {
  string.trim(s)
  |> string.split(on: ",")
  |> list.map(fn(move) {
    let value = string.drop_start(move, 1)

    case string.first(move) {
      Ok("s") -> {
        let assert Ok(index) = int.parse(value)
        Spin(index)
      }

      Ok("x") -> {
        let assert Ok(#(a, b)) = string.split_once(value, on: "/")
        let assert Ok(a) = int.parse(a)
        let assert Ok(b) = int.parse(b)
        Exchange(a, b)
      }

      Ok("p") -> {
        let assert Ok(#(a, b)) = string.split_once(value, on: "/")
        Partner(a, b)
      }

      _ -> panic as string.append("invalid move: ", move)
    }
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
  moves: List(DanceMove),
  programs: Array(String),
  seen: Dict(Array(String), Int),
  iteration: Int,
  remaining: Int,
) -> Array(String) {
  case remaining == 0 {
    True -> programs
    False -> {
      case dict.has_key(seen, programs) {
        True ->
          dance_loop(moves, programs, dict.new(), 0, { remaining % iteration })
        False -> {
          let seen = dict.insert(seen, programs, iteration)
          let programs = list.fold(moves, programs, apply_move)
          dance_loop(moves, programs, seen, iteration + 1, remaining - 1)
        }
      }
    }
  }
}

fn dance(moves: List(DanceMove), programs: Array(String), repeat: Int) {
  dance_loop(moves, programs, dict.new(), 0, repeat)
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
