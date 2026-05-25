import gleam/dict.{type Dict}
import gleam/int
import gleam/io
import gleam/list
import gleam/string

import utils

type Operation {
  Increment
  Decrement
}

type Action {
  Action(register: String, operation: Operation, value: Int)
}

type Compare {
  LessThan
  LessOrEqual
  Equal
  NotEqual
  GranterOrEqual
  GranterThan
}

type Condition {
  Condition(register: String, compare: Compare, value: Int)
}

type Instruction {
  Instruction(action: Action, condition: Condition)
}

fn parse(s: String) -> List(Instruction) {
  string.split(s, on: "\n")
  |> list.map(fn(line) {
    let assert [a, b, c, "if", d, e, f] = string.split(line, on: " ")
    let register = a
    let operation = case b {
      "inc" -> Increment
      "dec" -> Decrement
      _ -> panic as string.append("invalid operation: ", b)
    }
    let assert Ok(value) = int.parse(c)
    let action = Action(register:, operation:, value:)
    let register = d
    let compare = case e {
      "==" -> Equal
      "!=" -> NotEqual
      ">" -> GranterThan
      ">=" -> GranterOrEqual
      "<" -> LessThan
      "<=" -> LessOrEqual
      _ -> panic as string.append("invalid compare: ", e)
    }
    let assert Ok(value) = int.parse(f)
    let condition = Condition(register:, compare:, value:)
    Instruction(action:, condition:)
  })
}

fn read(
  registers: Dict(String, Int),
  register: String,
) -> #(Dict(String, Int), Int) {
  case dict.get(registers, register) {
    Ok(value) -> #(registers, value)
    Error(_) -> #(dict.insert(registers, register, 0), 0)
  }
}

fn check(
  registers: Dict(String, Int),
  condition: Condition,
) -> #(Dict(String, Int), Bool) {
  let #(registers, value) = read(registers, condition.register)
  let result = case condition.compare {
    Equal -> value == condition.value
    NotEqual -> value != condition.value
    LessOrEqual -> value <= condition.value
    LessThan -> value < condition.value
    GranterOrEqual -> value >= condition.value
    GranterThan -> value > condition.value
  }
  #(registers, result)
}

fn act(registers: Dict(String, Int), action: Action) -> Dict(String, Int) {
  let #(registers, value) = read(registers, action.register)
  let result = case action.operation {
    Increment -> value + action.value
    Decrement -> value - action.value
  }
  dict.insert(registers, action.register, result)
}

fn largest_register_value_after_execution(
  instructions: List(Instruction),
) -> Int {
  let registers =
    list.fold(instructions, dict.new(), fn(registers, instruction) {
      case check(registers, instruction.condition) {
        #(registers, True) -> act(registers, instruction.action)
        #(registers, False) -> registers
      }
    })
  dict.fold(registers, 0, fn(curr, _, value) { int.max(curr, value) })
}

/// Returns the largest register value after all instructions finish executing.
pub fn part1(s: String) -> Int {
  parse(s)
  |> largest_register_value_after_execution()
}

fn highest_register_value_seen_during_execution(
  instructions: List(Instruction),
) -> Int {
  let #(_, maximum) =
    list.fold(instructions, #(dict.new(), 0), fn(state, instruction) {
      let #(registers, maximum) = state
      case check(registers, instruction.condition) {
        #(registers, True) -> {
          let registers = act(registers, instruction.action)
          let #(registers, value) = read(registers, instruction.action.register)
          #(registers, int.max(maximum, value))
        }
        #(registers, False) -> #(registers, maximum)
      }
    })
  maximum
}

/// Returns the highest register value observed during execution.
pub fn part2(s: String) -> Int {
  parse(s)
  |> highest_register_value_seen_during_execution()
}

pub fn solve() {
  let input = utils.read_all_file("inputs/day-08.txt") |> string.trim()
  io.println("Day 08")
  utils.time_it("Part 1", fn() { part1(input) |> int.to_string() })
  utils.time_it("Part 2", fn() { part2(input) |> int.to_string() })
}
