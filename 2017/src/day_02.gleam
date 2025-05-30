import gleam/int
import gleam/io
import gleam/list
import gleam/result
import gleam/string

import utils

fn parse(s: String) -> Result(List(List(Int)), Nil) {
  let sep = case string.contains(s, "\t") {
    True -> "\t"
    False -> " "
  }
  string.trim(s)
  |> string.split("\n")
  |> list.map(fn(line) {
    string.trim(line)
    |> string.split(sep)
    |> list.map(int.parse)
    |> result.all()
  })
  |> result.all()
}

fn mapped_rows_sum(s: String, func: fn(List(Int)) -> Int) -> Int {
  parse(s)
  |> result.map(fn(xss) {
    xss
    |> list.map(func)
    |> int.sum()
  })
  |> result.unwrap(-1)
}

fn minmax_difference(xs: List(Int)) -> Int {
  case list.first(xs) {
    Error(_) -> 0
    Ok(head) -> {
      let #(min, max) =
        list.fold(xs, #(head, head), fn(p, x) {
          let #(min, max) = p
          #(int.min(min, x), int.max(max, x))
        })
      max - min
    }
  }
}

pub fn part1(s: String) -> Int {
  mapped_rows_sum(s, minmax_difference)
}

// adaptation of gleam/list function
// https://github.com/gleam-lang/stdlib/blob/v0.60.0/src/gleam/list.gleam#L2159-L2159
fn evenly_divisible_loop(items: List(Int)) -> Result(Int, Nil) {
  case items {
    [] -> Error(Nil)
    [first, ..rest] -> {
      let evenly_divisor =
        list.find_map(rest, fn(other) {
          case other % first == 0 || first % other == 0 {
            True -> Ok(other)
            False -> Error(Nil)
          }
        })
      case evenly_divisor {
        Ok(other) ->
          Ok(case other >= first {
            True -> other / first
            False -> first / other
          })
        Error(_) -> evenly_divisible_loop(rest)
      }
    }
  }
}

fn evenly_divisible(xs: List(Int)) -> Int {
  case evenly_divisible_loop(xs) {
    Ok(value) -> value
    Error(_) -> panic as "none evenly divisible founded"
  }
}

pub fn part2(s: String) -> Int {
  mapped_rows_sum(s, evenly_divisible)
}

pub fn solve() {
  let input = utils.read_all_file("inputs/day-02.txt")
  io.println("Day 02")
  utils.time_it("Part 1", fn() { part1(input) |> int.to_string() })
  utils.time_it("Part 2", fn() { part2(input) |> int.to_string() })
}
