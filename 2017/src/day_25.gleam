import gleam/dict.{type Dict}
import gleam/int
import gleam/io
import gleam/list
import gleam/string

import utils

/// The direction in which the Turing-machine head moves.
type Direction {
  Left
  Right
}

/// The operation performed for a particular tape value:
/// write a value, move the head, and transition to another state.
type Instruction {
  Instruction(write_value: Bool, direction: Direction, next_state: String)
}

/// The pair of instructions selected according to whether the current
/// tape value is 0 or 1.
type StateTransition {
  StateTransition(on_one: Instruction, on_zero: Instruction)
}

/// The initial state, diagnostic step count, and state transition rules
/// defining the machine.
type TuringMachine {
  TuringMachine(
    initial_state: String,
    diagnostic_steps: Int,
    transitions: Dict(String, StateTransition),
  )
}

fn parse_write_value(write_value: String) {
  case write_value {
    "1" -> True
    "0" -> False
    _ -> panic as string.append("invalid write value: ", write_value)
  }
}

fn parse_direction(direction: String) {
  case direction {
    "right" -> Right
    "left" -> Left
    _ -> panic as string.append("invalid direction: ", direction)
  }
}

fn parse_transitions(
  lines: List(String),
  transitions: Dict(String, StateTransition),
) -> Dict(String, StateTransition) {
  case lines {
    [] -> transitions
    _ -> {
      let assert [
        state,
        if0,
        write0,
        direction0,
        next0,
        if1,
        write1,
        direction1,
        next1,
        ..lines
      ] = lines
      let assert "In state " <> state = state
      let state = string.remove_suffix(state, ":")
      //
      let assert "If the current value is 0:" = if0
      let assert "- Write the value " <> write0 = write0
      let write0 = string.remove_suffix(write0, ".") |> parse_write_value()
      let assert "- Move one slot to the " <> direction0 = direction0
      let direction0 =
        string.remove_suffix(direction0, ".") |> parse_direction()
      let assert "- Continue with state " <> next0 = next0
      let next0 = string.remove_suffix(next0, ".")
      let on_zero = Instruction(write0, direction0, next0)
      //
      let assert "If the current value is 1:" = if1
      let assert "- Write the value " <> write1 = write1
      let write1 = string.remove_suffix(write1, ".") |> parse_write_value()
      let assert "- Move one slot to the " <> direction1 = direction1
      let direction1 =
        string.remove_suffix(direction1, ".") |> parse_direction()
      let assert "- Continue with state " <> next1 = next1
      let next1 = string.remove_suffix(next1, ".")
      let on_one = Instruction(write1, direction1, next1)
      //
      let result =
        dict.insert(transitions, state, StateTransition(on_one:, on_zero:))
      parse_transitions(lines, result)
    }
  }
}

fn parse_instructions(lines: List(String)) {
  parse_transitions(lines, dict.new())
}

fn parse_tm(s: String) -> TuringMachine {
  let lines =
    string.split(s, on: "\n")
    |> list.map(string.trim)
    |> list.filter(fn(line) { string.length(line) > 0 })

  let assert [initial_state, diagnostic_after, ..lines] = lines

  let assert "Begin in state " <> initial_state = initial_state
  let initial_state = string.remove_suffix(initial_state, ".")

  let assert "Perform a diagnostic checksum after " <> steps = diagnostic_after
  let assert Ok(diagnostic_steps) =
    string.remove_suffix(steps, " steps.") |> int.parse()

  let transitions = parse_instructions(lines)
  TuringMachine(initial_state, diagnostic_steps, transitions)
}

fn count_ones(acc, left, right) {
  case left, right {
    [], [] -> acc
    [], _ -> count_ones(acc, right, left)
    [value, ..left], _ -> {
      let acc = case value {
        True -> acc + 1
        False -> acc
      }
      count_ones(acc, left, right)
    }
  }
}

fn move_head(left: List(Bool), right: List(Bool), direction: Direction) {
  case direction {
    Left -> {
      let #(previous, right) = case right {
        [] -> #(False, right)
        [previous, ..right] -> #(previous, right)
      }
      let left = [previous, ..left]
      #(left, right)
    }
    Right -> {
      let #(head, left) = case left {
        [] -> #(False, [False])
        [head] -> #(head, [False])
        [head, ..left] -> #(head, left)
      }
      let right = [head, ..right]
      #(left, right)
    }
  }
}

/// Executes the Turing machine recursively until the diagnostic
/// step count is reached, then returns the number of 1 values on the tape.
fn execute(
  tm: TuringMachine,
  state: String,
  left: List(Bool),
  right: List(Bool),
  step_count: Int,
) {
  case step_count < tm.diagnostic_steps {
    False -> count_ones(0, left, right)
    True -> {
      let assert Ok(transition) = dict.get(tm.transitions, state)
      let assert [head, ..left] = left
      let instruction = case head {
        False -> transition.on_zero
        True -> transition.on_one
      }
      let left = [instruction.write_value, ..left]
      let #(left, right) = move_head(left, right, instruction.direction)
      let current_state = instruction.next_state
      let steps = step_count + 1
      execute(tm, current_state, left, right, steps)
    }
  }
}

/// Runs a Turing machine from its initial state on a blank tape.
fn run(tm: TuringMachine) {
  execute(tm, tm.initial_state, [False], [], 0)
}

/// Parses and executes a Turing machine described in text, then
/// returns the diagnostic checksum after the specified number of steps.
pub fn part1(s: String) {
  run(parse_tm(s))
}

pub fn solve() {
  let input = utils.read_all_file("inputs/day-25.txt") |> string.trim()
  io.println("Day 25")
  utils.time_it("Part 1", fn() { part1(input) |> int.to_string() })
}
