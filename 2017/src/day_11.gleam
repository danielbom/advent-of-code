import gleam/int
import gleam/io
import gleam/list
import gleam/string

import utils

type Dir {
  N
  NE
  NW
  S
  SW
  SE
}

/// Position in a hex grid represented by three movement axes.
///
/// forward: N ↔ S
/// left:    NE ↔ SW
/// right:   NW ↔ SE
type Hex {
  Hex(left: Int, forward: Int, right: Int)
}

const origin = Hex(0, 0, 0)

/// Returns the shortest path length from the origin to the given position.
///
/// The position is represented as offsets along three hex-grid axes.
fn distance_from_origin(hex: Hex) -> Int {
  let left = int.absolute_value(hex.left)
  let forward = int.absolute_value(hex.forward)
  let right = int.absolute_value(hex.right)
  case left, forward, right {
    0, a, b | a, 0, b | a, b, 0 -> int.max(a, b)
    _, _, _ -> {
      let #(sides, forward) = case left > right {
        True -> #(left - right, forward + right)
        False -> #(right - left, forward + left)
      }
      sides + forward
    }
  }
}

/// Moves the position one step in the given direction.
fn move(hex: Hex, dir: Dir) -> Hex {
  case dir {
    N -> Hex(..hex, forward: hex.forward + 1)
    NE -> Hex(..hex, left: hex.left + 1)
    NW -> Hex(..hex, right: hex.right + 1)
    S -> Hex(..hex, forward: hex.forward - 1)
    SE -> Hex(..hex, right: hex.right - 1)
    SW -> Hex(..hex, left: hex.left - 1)
  }
}

/// Walks the path from the origin and updates a state value after each step.
fn walk_path(
  path: List(Dir),
  position: Hex,
  state: a,
  step: fn(a, Hex) -> a,
) -> a {
  case path {
    [] -> state
    [dir, ..path] -> {
      let position = move(position, dir)
      let state = step(state, position)
      walk_path(path, position, state, step)
    }
  }
}

/// Returns the position reached after following the entire path.
fn final_position(path: List(Dir)) -> Hex {
  list.fold(path, origin, move)
}

/// Parses a comma-separated path of hex-grid directions.
fn parse(s: String) -> List(Dir) {
  case string.trim(s) {
    "" -> []
    s -> string.split(s, on: ",")
  }
  |> list.map(fn(value) {
    case value {
      "n" -> N
      "ne" -> NE
      "nw" -> NW
      "s" -> S
      "se" -> SE
      "sw" -> SW
      _ -> panic as string.append("invalid dir: ", value)
    }
  })
}

/// Returns the distance from the origin after following the path.
pub fn part1(s: String) {
  parse(s)
  |> final_position()
  |> distance_from_origin()
}

/// Returns the maximum distance from the origin reached while following
/// the path.
pub fn part2(s: String) {
  walk_path(parse(s), origin, 0, fn(max_distance, position) {
    int.max(max_distance, distance_from_origin(position))
  })
}

pub fn solve() {
  let input = utils.read_all_file("inputs/day-11.txt") |> string.trim()
  io.println("Day 11")
  utils.time_it("Part 1", fn() { part1(input) |> int.to_string() })
  utils.time_it("Part 2", fn() { part2(input) |> int.to_string() })
}
