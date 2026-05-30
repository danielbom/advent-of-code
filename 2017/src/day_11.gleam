import gleam/int
import gleam/io
import gleam/list
import gleam/pair
import gleam/string

import utils

type Dir {
  North
  Northeast
  Northwest
  South
  Southwest
  Southeast
}

type Hex {
  Hex(n: Int, ne: Int, nw: Int, s: Int, sw: Int, se: Int)
}

fn hex_add(a: Hex, b: Hex) -> Hex {
  Hex(
    n: a.n + b.n,
    ne: a.ne + b.ne,
    nw: a.nw + b.nw,
    s: a.s + b.s,
    sw: a.sw + b.sw,
    se: a.se + b.se,
  )
}

/// Returns the minimum number of steps from the origin to the position.
fn hex_distance_from_origin(hex: Hex) -> Int {
  let Hex(n, ne, nw, s, sw, se) = hex
  let top = int.absolute_value(n - s)
  let left = int.absolute_value(nw - se)
  let right = int.absolute_value(ne - sw)
  case top == 0 || left == 0 || right == 0 {
    True -> int.max(right, left) |> int.max(top)
    False -> {
      let #(sides, forward) = case left > right {
        True -> #(left - right, top + right)
        False -> #(right - left, top + left)
      }
      sides + forward
    }
  }
}

const zero = Hex(0, 0, 0, 0, 0, 0)

/// Converts a direction into a unit movement on the hex grid.
fn dir_to_hex(dir: Dir) -> Hex {
  case dir {
    North -> Hex(..zero, n: 1)
    Northeast -> Hex(..zero, ne: 1)
    Northwest -> Hex(..zero, nw: 1)
    South -> Hex(..zero, s: 1)
    Southeast -> Hex(..zero, se: 1)
    Southwest -> Hex(..zero, sw: 1)
  }
}

/// Parses a comma-separated path of hex-grid directions.
fn parse(s: String) -> List(Dir) {
  string.trim(s)
  |> string.split(on: ",")
  |> list.filter_map(fn(value) {
    case value {
      "n" -> Ok(North)
      "ne" -> Ok(Northeast)
      "nw" -> Ok(Northwest)
      "s" -> Ok(South)
      "se" -> Ok(Southeast)
      "sw" -> Ok(Southwest)
      "" -> Error(Nil)
      _ -> panic as string.append("invalid dir: ", value)
    }
  })
}

pub fn part1(s: String) {
  parse(s)
  |> list.fold(zero, fn(acc, dir) { hex_add(acc, dir_to_hex(dir)) })
  |> hex_distance_from_origin()
}

pub fn part2(s: String) {
  parse(s)
  |> list.fold(#(zero, 0), fn(state, dir) {
    let #(acc, max) = state
    let acc = hex_add(acc, dir_to_hex(dir))
    #(acc, int.max(max, hex_distance_from_origin(acc)))
  })
  |> pair.second
}

pub fn solve() {
  let input = utils.read_all_file("inputs/day-11.txt") |> string.trim()
  io.println("Day 11")
  utils.time_it("Part 1", fn() { part1(input) |> int.to_string() })
  utils.time_it("Part 2", fn() { part2(input) |> int.to_string() })
}
