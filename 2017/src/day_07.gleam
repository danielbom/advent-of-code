import gleam/dict.{type Dict}
import gleam/int
import gleam/io
import gleam/list
import gleam/result
import gleam/string

import utils

type Program {
  Program(name: String, weight: Int, next: List(String))
}

fn parse(s: String) -> List(Program) {
  string.split(s, on: "\n")
  |> list.map(fn(line) {
    let #(name_weight, next_programs) = case string.split(line, on: " -> ") {
      [fst] -> #(fst, [])
      [fst, snd] -> #(fst, string.split(snd, on: ", "))
      _ -> panic as string.append("invalid line: ", line)
    }
    let assert Ok(#(name, weight)) = string.split_once(name_weight, " ")
    let assert Ok(weight) =
      weight
      |> string.drop_start(1)
      |> string.drop_end(1)
      |> int.parse()
    Program(name, weight, next_programs)
  })
}

/// Finds the root program in a program dependency tree.
///
/// The root program is the only node that has no parent.
fn find_root_program(programs: List(Program)) -> List(Program) {
  let parents =
    list.fold(programs, dict.new(), fn(graph, program) {
      list.fold(program.next, graph, fn(parents, next) {
        let values = dict.get(parents, next)
        let values = result.lazy_unwrap(values, fn() { dict.new() })
        let values = dict.insert(values, program.name, True)
        dict.insert(parents, next, values)
      })
    })
  list.filter(programs, fn(p) { !dict.has_key(parents, p.name) })
}

pub fn part1(s: String) -> String {
  let programs = parse(s)
  case find_root_program(programs) {
    [root] -> root.name
    _ -> ""
  }
}

fn catch_missmatch(weights: Dict(String, Int), edges: List(String)) {
  case edges {
    [] -> Error(Nil)
    [to, ..] -> {
      let assert Ok(weight) = dict.get(weights, to)
      list.find_map(edges, fn(other) {
        let assert Ok(other_weight) = dict.get(weights, other)
        case weight != other_weight {
          True -> Ok(#(weights, edges))
          False -> Error(Nil)
        }
      })
    }
  }
}

/// Finds the unbalanced program in the dependency tree.
///
/// A program total weight is its own weight plus the total
/// weight of all its children.
///
/// Traversal proceeds bottom-up, updating parent weights
/// from already processed child nodes.
fn find_unbalanced_program(
  visited: Dict(String, Bool),
  weights: Dict(String, Int),
  current: List(#(String, List(String))),
  next: List(#(String, List(String))),
) {
  case current, next {
    [], [] -> Error(Nil)
    [#(from, edges), ..rest], _ -> {
      case dict.has_key(visited, from) {
        True -> find_unbalanced_program(visited, weights, rest, next)
        False -> {
          let completed =
            list.all(edges, fn(edge) { dict.has_key(visited, edge) })
          case completed {
            True -> {
              use <- result.lazy_or(catch_missmatch(weights, edges))
              let visited = dict.insert(visited, from, True)
              let assert Ok(weight) = dict.get(weights, from)
              let assert [child, ..] = edges
              let assert Ok(child_weight) = dict.get(weights, child)
              let new_weight = list.length(edges) * child_weight + weight
              let weights = dict.insert(weights, from, new_weight)
              find_unbalanced_program(visited, weights, rest, next)
            }
            False -> {
              let next = [#(from, edges), ..next]
              find_unbalanced_program(visited, weights, rest, next)
            }
          }
        }
      }
    }
    [], _ -> find_unbalanced_program(visited, weights, next, [])
  }
}

/// Computes the corrected weight for the unbalanced program.
///
/// Finds the unbalanced program and its sibling weights,
/// then infers the weight needed to balance the tree.
fn compute_corrected_weight(programs: List(Program)) -> Int {
  let weights =
    list.fold(programs, dict.new(), fn(weights, p) {
      dict.insert(weights, p.name, p.weight)
    })
  let dependency_tree =
    list.map(programs, fn(program) { #(program.name, program.next) })
  let visited =
    list.fold(programs, dict.new(), fn(visited, program) {
      case program.next {
        [] -> dict.insert(visited, program.name, True)
        _ -> visited
      }
    })
  let assert Ok(#(new_weights, nodes)) =
    find_unbalanced_program(visited, weights, dependency_tree, [])
  let nodes_weight =
    list.map(nodes, fn(edge) {
      let assert Ok(weight) = dict.get(weights, edge)
      let assert Ok(new_weight) = dict.get(new_weights, edge)
      #(weight, new_weight)
    })
  let assert [#(_, new_weight), ..] = nodes_weight
  let #(fst, snd) =
    list.partition(nodes_weight, fn(values) {
      let #(_, other_new_weight) = values
      new_weight == other_new_weight
    })
  let #(#(weight, new_weight), #(_, expected_new_weight)) = case fst, snd {
    [single], [head, ..] -> #(single, head)
    [head, ..], [single] -> #(single, head)
    _, _ -> panic as "unexpected"
  }
  weight + expected_new_weight - new_weight
}

pub fn part2(s: String) -> Int {
  let programs = parse(s)
  compute_corrected_weight(programs)
}

pub fn solve() {
  let input = utils.read_all_file("inputs/day-07.txt") |> string.trim()
  io.println("Day 07")
  utils.time_it("Part 1", fn() { part1(input) })
  utils.time_it("Part 2", fn() { part2(input) |> int.to_string() })
}
