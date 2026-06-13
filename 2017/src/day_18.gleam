import gleam/dict.{type Dict}
import gleam/int
import gleam/io
import gleam/list
import gleam/string

import utils

type RegVal {
  Register(String)
  Value(Int)
}

type Instruction {
  Set(String, RegVal)
  Add(String, RegVal)
  Multiply(String, RegVal)
  Modulus(String, RegVal)
  Jump(RegVal, RegVal)
  Snd(RegVal)
  Rcv(String)
}

fn parse_regval(arg: String) -> RegVal {
  case int.parse(arg) {
    Ok(value) -> Value(value)
    Error(_) -> Register(arg)
  }
}

/// Convert input text into a list of instructions
fn parse(s: String) {
  string.trim(s)
  |> string.split(on: "\n")
  |> list.map(fn(line) {
    let assert [command, ..args] = string.split(line, on: " ")
    case command {
      "set" -> {
        let assert [fst, snd] = args
        let snd = parse_regval(snd)
        Set(fst, snd)
      }
      "mod" -> {
        let assert [fst, snd] = args
        let snd = parse_regval(snd)
        Modulus(fst, snd)
      }
      "add" -> {
        let assert [fst, snd] = args
        let snd = parse_regval(snd)
        Add(fst, snd)
      }
      "mul" -> {
        let assert [fst, snd] = args
        let snd = parse_regval(snd)
        Multiply(fst, snd)
      }
      "jgz" -> {
        let assert [fst, snd] = args
        let fst = parse_regval(fst)
        let snd = parse_regval(snd)
        Jump(fst, snd)
      }
      "snd" -> {
        let assert [single] = args
        let single = parse_regval(single)
        Snd(single)
      }
      "rcv" -> {
        let assert [single] = args
        Rcv(single)
      }
      _ -> panic as string.append("unknown command: ", command)
    }
  })
}

fn index_instructions(instructions: List(Instruction)) {
  instructions
  |> list.index_map(fn(value, index) { #(index, value) })
  |> dict.from_list()
}

fn get_register_value(registers, register) -> Int {
  case dict.get(registers, register) {
    Ok(value) -> value
    Error(_) -> 0
  }
}

fn get_regval_value(registers, regval) -> Int {
  case regval {
    Register(register) -> get_register_value(registers, register)
    Value(value) -> value
  }
}

type SndRcv {
  OnSnd(RegVal)
  OnRcv(String)
  OnNone
}

/// Execute a single instruction
fn execute_once(registers, instruction) {
  case instruction {
    Set(register, regval) -> {
      let value = get_regval_value(registers, regval)
      let registers = dict.insert(registers, register, value)
      #(registers, 1, OnNone)
    }
    Add(register, regval) -> {
      let lhs = get_register_value(registers, register)
      let rhs = get_regval_value(registers, regval)
      let registers = dict.insert(registers, register, lhs + rhs)
      #(registers, 1, OnNone)
    }
    Multiply(register, regval) -> {
      let lhs = get_register_value(registers, register)
      let rhs = get_regval_value(registers, regval)
      let registers = dict.insert(registers, register, lhs * rhs)
      #(registers, 1, OnNone)
    }
    Modulus(register, regval) -> {
      let lhs = get_register_value(registers, register)
      let rhs = get_regval_value(registers, regval)
      let registers = dict.insert(registers, register, lhs % rhs)
      #(registers, 1, OnNone)
    }
    Jump(value, jump) -> {
      let value = get_regval_value(registers, value)
      let jump = get_regval_value(registers, jump)
      let jump = case value <= 0 {
        True -> 1
        False -> jump
      }
      #(registers, jump, OnNone)
    }
    Snd(regval) -> #(registers, 1, OnSnd(regval))
    Rcv(register) -> #(registers, 1, OnRcv(register))
  }
}

/// Executes instructions in a loop until receiving a value
fn execute_loop(
  instructions: Dict(Int, Instruction),
  pc: Int,
  registers: Dict(String, Int),
  played: List(Int),
) {
  case dict.get(instructions, pc) {
    Error(_) -> Error(Nil)
    Ok(instruction) -> {
      let #(registers, jump, snd_rcv) = execute_once(registers, instruction)
      let pc = pc + jump
      case snd_rcv {
        OnNone -> {
          execute_loop(instructions, pc, registers, played)
        }
        OnSnd(regval) -> {
          let value = get_regval_value(registers, regval)
          let played = [value, ..played]
          execute_loop(instructions, pc, registers, played)
        }
        OnRcv(register) -> {
          let value = get_register_value(registers, register)
          let continue = case played, value == 0 {
            _, True -> Error(Nil)
            [], False -> Error(Nil)
            [last, ..], _ if last != 0 -> Ok(last)
            _, _ -> Error(Nil)
          }
          case continue {
            Ok(result) -> Ok(result)
            Error(_) -> execute_loop(instructions, pc + 1, registers, played)
          }
        }
      }
    }
  }
}

fn execute(instructions: List(Instruction)) {
  let instructions = index_instructions(instructions)
  let assert Ok(result) = execute_loop(instructions, 0, dict.new(), [])
  result
}

pub fn part1(s: String) {
  parse(s)
  |> execute()
}

type ProcId {
  Proc0
  Proc1
}

type Proc {
  Proc(
    id: ProcId,
    pc: Int,
    registers: Dict(String, Int),
    waiting: List(String),
    queue: List(Int),
    count: Int,
  )
}

/// Execute one instruction of a process, handling message passing
fn execute_duo_once(instructions: Dict(Int, Instruction), proc: Proc) {
  case proc.waiting, proc.queue {
    [_], [] -> #([], proc)

    [register], [value, ..queue] -> {
      let waiting = []
      let registers = dict.insert(proc.registers, register, value)
      let next_proc = Proc(..proc, registers:, queue:, waiting:)
      #([], next_proc)
    }

    [], _ -> {
      case dict.get(instructions, proc.pc) {
        Error(_) -> #([], proc)

        Ok(instruction) -> {
          let #(registers, jump, effect) =
            execute_once(proc.registers, instruction)

          case effect {
            OnNone -> {
              let pc = proc.pc + jump
              let next_proc = Proc(..proc, registers:, pc:)
              #([], next_proc)
            }
            OnSnd(regval) -> {
              let value = get_regval_value(proc.registers, regval)
              let pc = proc.pc + jump
              let count = proc.count + 1
              let next_proc = Proc(..proc, registers:, pc:, count:)
              #([value], next_proc)
            }
            OnRcv(register) -> {
              let pc = proc.pc + jump
              let next_proc = Proc(..proc, pc:, waiting: [register])
              #([], next_proc)
            }
          }
        }
      }
    }

    _, _ -> {
      echo #("invalid waiting", proc)
      panic as "waiting.length > 1"
    }
  }
}

/// Executes both processes in lockstep until termination or deadlock.
///
/// Each process runs one step, producing outbound messages that are
/// appended to the peer queue after both executions complete.
///
/// Deadlock occurs when:
/// - both processes are waiting; and
/// - both message queues are empty.
///
/// NOTE:
/// I previously attempted to model this with actors (`gleam_otp`), but
/// synchronization issues introduced race conditions that caused
/// nondeterministic results.
///
/// I would like to revisit that approach in the future, possibly with a
/// stronger synchronization primitive or explicit barrier between the
/// "waiting for input" states of both processes.
fn execute_duo_loop(
  instructions: Dict(Int, Instruction),
  proc0: Proc,
  proc1: Proc,
) {
  let #(messages1, proc0) = execute_duo_once(instructions, proc0)
  let #(messages0, proc1) = execute_duo_once(instructions, proc1)
  let proc0 = Proc(..proc0, queue: list.append(proc0.queue, messages0))
  let proc1 = Proc(..proc1, queue: list.append(proc1.queue, messages1))
  case proc0.waiting, proc0.queue, proc1.waiting, proc1.queue {
    [_], [], [_], [] -> {
      // deadlock
      proc1.count
    }
    _, _, _, _ -> {
      case
        dict.has_key(instructions, proc0.pc)
        || dict.has_key(instructions, proc1.pc)
      {
        True -> {
          execute_duo_loop(instructions, proc0, proc1)
        }
        False -> {
          proc1.count
        }
      }
    }
  }
}

fn execute_duo(instructions: List(Instruction)) -> Int {
  let instructions = index_instructions(instructions)
  let proc0 = Proc(Proc0, 0, dict.from_list([#("p", 0)]), [], [], 0)
  let proc1 = Proc(Proc1, 0, dict.from_list([#("p", 1)]), [], [], 0)
  execute_duo_loop(instructions, proc0, proc1)
}

pub fn part2(s: String) {
  parse(s)
  |> execute_duo()
}

pub fn solve() {
  let input = utils.read_all_file("inputs/day-18.txt") |> string.trim()
  io.println("Day 18")
  utils.time_it("Part 1", fn() { part1(input) |> int.to_string() })
  utils.time_it("Part 2", fn() { part2(input) |> int.to_string() })
}
