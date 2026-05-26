import gleam/int
import gleam/io
import gleam/list
import gleam/string

import utils

type Group {
  Group(items: List(Group))
  Value(value: String)
  Garbage(value: String)
}

fn parse_groups_loop(
  grouping: Bool,
  remaining: List(String),
  pending: List(String),
  result: List(List(Group)),
) {
  let b = fn(values) { Garbage(string.concat(list.reverse(values))) }
  let v = fn(values) { Value(string.concat(list.reverse(values))) }
  let g = fn(values) { Group(list.reverse(values)) }
  case grouping, remaining {
    True, [] -> {
      let result = case pending, result {
        [], _ -> result
        _, [] -> [[v(pending)]]
        _, [top, ..tail] -> [[v(pending), ..top], ..tail]
      }
      list.fold(result, [], list.append)
    }
    True, ["{", ..rest] -> {
      let #(others, result) = case pending, result {
        [], _ -> #([], result)
        _, [] -> #([], [[v(pending)]])
        _, [top, ..tail] -> #([], [[v(pending), ..top], ..tail])
      }
      parse_groups_loop(grouping, rest, others, [[], ..result])
    }
    True, ["}", ..rest] -> {
      let #(others, result) = case result {
        [] -> #(["{", ..pending], [])
        [top, ..tail] -> {
          case tail, pending {
            [], [] -> #([], [[g(top)]])
            [], _ -> #([], [[g([v(pending), ..top])]])
            [snd, ..rest], [] -> #([], [[g(top), ..snd], ..rest])
            [snd, ..rest], _ -> #([], [[g([v(pending), ..top]), ..snd], ..rest])
          }
        }
      }
      parse_groups_loop(grouping, rest, others, result)
    }
    True, ["<", ..rest] -> {
      let result = case pending, result {
        [], _ -> result
        _, [] -> [[v(pending)]]
        _, [top, ..tail] -> [[v(pending), ..top], ..tail]
      }
      parse_groups_loop(False, rest, [], result)
    }
    True, [ch, ..rest] -> {
      parse_groups_loop(grouping, rest, [ch, ..pending], result)
    }
    False, [">", ..rest] -> {
      let result = case result {
        [] -> [[b(pending)]]
        [top, ..tail] -> [[b(pending), ..top], ..tail]
      }
      parse_groups_loop(True, rest, [], result)
    }
    False, ["!", _, ..rest] -> {
      parse_groups_loop(False, rest, pending, result)
    }
    False, [ch, ..rest] -> {
      parse_groups_loop(False, rest, [ch, ..pending], result)
    }
    False, [] -> {
      parse_groups_loop(True, [], pending, result)
    }
  }
}

/// Builds the nested group structure from the input stream.
/// Groups are delimited by `{}` and garbage sections by `<>`.
fn parse_groups(input_stream: List(String)) {
  parse_groups_loop(True, input_stream, [], [])
}

fn groups_score_loop(groups: List(Group), depth: Int, score: Int) {
  case groups {
    [] -> score
    [Group(children), ..rest] -> {
      let score = groups_score_loop(children, depth + 1, score)
      groups_score_loop(rest, depth, score + depth + 1)
    }
    [_, ..rest] -> groups_score_loop(rest, depth, score)
  }
}

/// Returns the sum of all group nesting depths.
fn groups_score(groups: List(Group)) {
  groups_score_loop(groups, 0, 0)
}

pub fn part1(s: String) {
  string.to_graphemes(s)
  |> parse_groups()
  |> groups_score()
}

fn count_garbage_loop(groups: List(Group), count: Int) {
  case groups {
    [] -> count
    [Group(children), ..rest] -> {
      let count = count_garbage_loop(children, count)
      count_garbage_loop(rest, count)
    }
    [Garbage(value), ..rest] -> {
      count_garbage_loop(rest, count + string.length(value))
    }
    [_, ..rest] -> {
      count_garbage_loop(rest, count)
    }
  }
}

/// Returns the total number of garbage characters.
fn count_garbage(groups: List(Group)) {
  count_garbage_loop(groups, 0)
}

pub fn part2(s: String) {
  string.to_graphemes(s)
  |> parse_groups()
  |> count_garbage()
}

pub fn solve() {
  let input = utils.read_all_file("inputs/day-09.txt") |> string.trim()
  io.println("Day 09")
  utils.time_it("Part 1", fn() { part1(input) |> int.to_string() })
  utils.time_it("Part 2", fn() { part2(input) |> int.to_string() })
}
