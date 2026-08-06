import gleam/int
import gleam/io
import gleam/list
import gleam/order
import gleam/string

import utils

type Port {
  Port(fst: Int, snd: Int)
}

fn parse(s: String) -> List(Port) {
  string.split(s, on: "\n")
  |> list.map(fn(item) {
    let assert Ok(#(fst, snd)) = string.split_once(item, on: "/")
    let assert Ok(fst) = int.parse(fst)
    let assert Ok(snd) = int.parse(snd)
    Port(fst, snd)
  })
}

/// Folds over a list while maintaining the already-visited elements
/// separately from the remaining elements.
fn zipper_fold(
  left: List(a),
  right: List(a),
  state: b,
  step: fn(b, a, List(a), List(a)) -> b,
) {
  case left {
    [] -> state
    [current, ..left] -> {
      let next = step(state, current, left, right)
      zipper_fold(left, [current, ..right], next, step)
    }
  }
}

fn zipper_partition_loop(
  left: List(a),
  right: List(a),
  predicate: fn(a) -> Bool,
  left_result: List(a),
  right_result: List(a),
) -> #(List(a), List(a)) {
  case left, right {
    [], [] -> #(left_result, right_result)
    [], _ ->
      zipper_partition_loop(right, left, predicate, left_result, right_result)
    [value, ..left], _ -> {
      case predicate(value) {
        True ->
          zipper_partition_loop(
            //
            left,
            right,
            predicate,
            [value, ..left_result],
            right_result,
          )
        False ->
          zipper_partition_loop(
            //
            left,
            right,
            predicate,
            left_result,
            [value, ..right_result],
          )
      }
    }
  }
}

fn zipper_partition(
  left: List(a),
  right: List(a),
  condition: fn(a) -> Bool,
) -> #(List(a), List(a)) {
  zipper_partition_loop(left, right, condition, [], [])
}

/// Determines how completed bridges are ranked.
type Objective {
  /// Selects the bridge with the greatest strength.
  Strongest
  /// Selects the longest bridge, using strength to break ties.
  LongestThenStrongest
}

type BridgeScore {
  BridgeScore(length: Int, strength: Int)
}

/// Explores all bridges that can be built from the available ports and
/// returns the best length and strength found according to the objective.
fn find_best_bridge(
  objective: Objective,
  left_ports: List(Port),
  right_ports: List(Port),
  open_pin: Int,
  strength: Int,
  length: Int,
  best_score: BridgeScore,
) -> BridgeScore {
  let #(matching_ports, remaining_ports) =
    zipper_partition(left_ports, right_ports, fn(port) {
      port.fst == open_pin || port.snd == open_pin
    })

  case matching_ports {
    [] -> {
      case objective {
        LongestThenStrongest -> {
          case int.compare(length, best_score.length) {
            order.Gt -> BridgeScore(length, strength)
            order.Eq ->
              BridgeScore(length, int.max(strength, best_score.strength))
            order.Lt -> best_score
          }
        }
        Strongest -> BridgeScore(length, int.max(strength, best_score.strength))
      }
    }
    _ -> {
      zipper_fold(
        matching_ports,
        remaining_ports,
        best_score,
        fn(best_score, port, left, right) {
          let next_pin = case port.fst == open_pin {
            True -> port.snd
            False -> port.fst
          }
          find_best_bridge(
            objective,
            left,
            right,
            next_pin,
            strength + open_pin + next_pin,
            length + 1,
            best_score,
          )
        },
      )
    }
  }
}

pub fn part1(s: String) {
  let ports = parse(s)
  let best_score =
    find_best_bridge(Strongest, ports, [], 0, 0, 0, BridgeScore(0, 0))
  best_score.strength
}

pub fn part2(s: String) {
  let ports = parse(s)
  let best_score =
    find_best_bridge(
      LongestThenStrongest,
      ports,
      [],
      0,
      0,
      0,
      BridgeScore(0, 0),
    )
  best_score.strength
}

pub fn solve() {
  let input = utils.read_all_file("inputs/day-24.txt") |> string.trim()
  io.println("Day 24")
  utils.time_it("Part 1", fn() { part1(input) |> int.to_string() })
  utils.time_it("Part 2", fn() { part2(input) |> int.to_string() })
}
