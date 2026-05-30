import gleam/dict.{type Dict}
import gleam/int
import gleam/io
import gleam/list
import gleam/string

import utils

/// Adjacency lookup table keyed by program ID.
type Graph =
  Dict(Int, Program)

/// A program and the IDs of programs directly connected to it.
type Program {
  Program(id: Int, connections: List(Int))
}

/// Parses the input into a list of programs and their connections.
fn parse(s: String) -> List(Program) {
  string.trim(s)
  |> string.split(on: "\n")
  |> list.map(fn(line) {
    let assert Ok(#(id, connections)) = string.split_once(line, on: " <-> ")
    let assert Ok(id) = int.parse(id)
    let connections =
      connections
      |> string.split(on: ", ")
      |> list.map(fn(value) {
        let assert Ok(value) = int.parse(value)
        value
      })
    Program(id:, connections:)
  })
}

/// Builds a program lookup table keyed by program ID.
fn build_graph(programs: List(Program)) -> Graph {
  list.fold(programs, dict.new(), fn(graph, program) {
    dict.insert(graph, program.id, program)
  })
}

/// Performs a depth-first traversal from the IDs in `stack`,
/// returning both the number of visited programs and the updated
/// visited set.
fn visit_component(
  graph: Graph,
  stack: List(Int),
  visited: Dict(Int, Bool),
  count: Int,
) {
  case stack {
    [] -> #(count, visited)
    [current, ..stack] -> {
      case dict.has_key(visited, current) {
        True -> visit_component(graph, stack, visited, count)
        False -> {
          let assert Ok(program) = dict.get(graph, current)
          let visited = dict.insert(visited, current, True)
          let stack = list.append(stack, program.connections)
          visit_component(graph, stack, visited, count + 1)
        }
      }
    }
  }
}

/// Counts how many programs are reachable from `start`,
/// including `start` itself.
fn reachable_count(graph: Graph, start: Int) -> Int {
  let #(count, _) = visit_component(graph, [start], dict.new(), 0)
  count
}

pub fn part1(s: String) {
  parse(s)
  |> build_graph()
  |> reachable_count(0)
}

/// Iterates over all program IDs, discovering connected components
/// and counting how many distinct groups exist in the graph.
fn count_components_loop(
  graph: Graph,
  ids: List(Int),
  visited: Dict(Int, Bool),
  count: Int,
) -> Int {
  case ids {
    [] -> count
    [current, ..ids] -> {
      case dict.has_key(visited, current) {
        True -> count_components_loop(graph, ids, visited, count)
        False -> {
          let #(_, visited) = visit_component(graph, [current], visited, 0)
          count_components_loop(graph, ids, visited, count + 1)
        }
      }
    }
  }
}

/// Counts the number of connected components in the graph.
fn count_components(graph: Graph) {
  count_components_loop(graph, dict.keys(graph), dict.new(), 0)
}

pub fn part2(s: String) {
  parse(s)
  |> build_graph()
  |> count_components()
}

pub fn solve() {
  let input = utils.read_all_file("inputs/day-12.txt") |> string.trim()
  io.println("Day 12")
  utils.time_it("Part 1", fn() { part1(input) |> int.to_string() })
  utils.time_it("Part 2", fn() { part2(input) |> int.to_string() })
}
