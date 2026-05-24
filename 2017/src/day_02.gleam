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

fn sum_by(xs: List(a), by: fn(a) -> Int) -> Int {
  list.fold(xs, 0, fn(curr, x) { curr + by(x) })
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

/// Computes the spreadsheet checksum.
///
/// For each line, finds the difference between the largest
/// and smallest whitespace-separated numbers, then returns
/// the sum of all line differences.
pub fn part1(s: String) -> Int {
  case parse(s) {
    Ok(xss) -> sum_by(xss, minmax_difference)
    Error(_) -> -1
  }
}

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

/// Computes the spreadsheet checksum using evenly divisible values.
///
/// For each line, finds the only pair of numbers where one evenly
/// divides the other, then adds the division result to the total.
pub fn part2(s: String) -> Int {
  case parse(s) {
    Ok(xss) -> sum_by(xss, evenly_divisible)
    Error(_) -> -1
  }
}

pub fn solve() {
  let input = utils.read_all_file("inputs/day-02.txt")
  io.println("Day 02")
  utils.time_it("Part 1", fn() { part1(input) |> int.to_string() })
  utils.time_it("Part 2", fn() { part2(input) |> int.to_string() })
}
