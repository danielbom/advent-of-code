import gleam/dict.{type Dict}
import gleam/int
import gleam/io
import gleam/list
import gleam/string

import utils

pub type RegVal {
  Val(x: Int)
  Reg(r: String)
}

pub type Instruction {
  Set(a: String, b: RegVal)
  Jnz(a: RegVal, b: RegVal)
  Mul(a: String, b: RegVal)
  Sub(a: String, b: RegVal)
}

type Instructions =
  Dict(Int, Instruction)

type Registers =
  Dict(String, Int)

fn parse_regval(x: String) {
  case int.parse(x) {
    Error(_) -> Reg(x)
    Ok(val) -> Val(val)
  }
}

fn parse_instruction(instruction: String) -> Instruction {
  case string.split(instruction, on: " ") {
    ["set", a, b] -> Set(a, parse_regval(b))
    ["jnz", a, b] -> Jnz(parse_regval(a), parse_regval(b))
    ["mul", a, b] -> Mul(a, parse_regval(b))
    ["sub", a, b] -> Sub(a, parse_regval(b))
    _ -> panic as string.append("unknown instruction: ", instruction)
  }
}

fn parse_instructions(s: String) -> Instructions {
  string.split(s, on: "\n")
  |> list.map(parse_instruction)
  |> list.index_map(fn(instruction, index) { #(index, instruction) })
  |> dict.from_list()
}

fn get_val(registers, regval) {
  case regval {
    Val(val) -> Ok(val)
    Reg(reg) -> dict.get(registers, reg)
  }
}

/// Executes a single instruction and returns:
/// - the relative program counter offset to apply next
/// - the updated register state
///
/// The returned offset is usually `1`, but `jnz` may return any value.
fn run_instruction(
  registers: Registers,
  instruction: Instruction,
) -> #(Int, Registers) {
  case instruction {
    Set(a, b) -> {
      let assert Ok(val_b) = get_val(registers, b)
      let registers = dict.insert(registers, a, val_b)
      #(1, registers)
    }
    Jnz(a, b) -> {
      let assert Ok(val_a) = get_val(registers, a)
      let assert Ok(val_b) = get_val(registers, b)
      let jump = case val_a {
        0 -> 1
        _ -> val_b
      }
      #(jump, registers)
    }
    Mul(a, b) -> {
      let assert Ok(val_a) = get_val(registers, Reg(a))
      let assert Ok(val_b) = get_val(registers, b)
      let registers = dict.insert(registers, a, val_a * val_b)
      #(1, registers)
    }
    Sub(a, b) -> {
      let assert Ok(val_a) = get_val(registers, Reg(a))
      let assert Ok(val_b) = get_val(registers, b)
      let registers = dict.insert(registers, a, val_a - val_b)
      #(1, registers)
    }
  }
}

/// Executes the program until termination and returns the number
/// of times the `mul` instruction was executed.
fn run_and_count_mul(
  pc: Int,
  registers: Registers,
  instructions: Instructions,
  count: Int,
) {
  case dict.get(instructions, pc) {
    Ok(instruction) -> {
      let count = case instruction {
        Mul(_, _) -> count + 1
        _ -> count
      }
      let #(jump, registers) = run_instruction(registers, instruction)
      let pc = pc + jump
      run_and_count_mul(pc, registers, instructions, count)
    }
    Error(_) -> count
  }
}

/// Executes instructions starting at `pc` until the program counter
/// leaves the instruction space, returning the final register state.
///
/// NOTE: This interpreter was not used to solve part 2.
///
/// After reimplementing the program in C and debugging its execution,
/// I discovered that the expensive part of the computation is testing
/// whether the value in register b is prime.
///
/// The final result is the number of non-prime values between the
/// initial values of registers b and c.
pub fn run_program(pc: Int, registers: Registers, instructions: Instructions) {
  case dict.get(instructions, pc) {
    Ok(instruction) -> {
      let #(jump, registers) = run_instruction(registers, instruction)
      let pc = pc + jump
      run_program(pc, registers, instructions)
    }
    Error(_) -> registers
  }
}

fn get_registers() -> Registers {
  "abcdefgh"
  |> string.to_graphemes()
  |> list.map(fn(register) { #(register, 0) })
  |> dict.from_list()
}

pub fn part1(s: String) {
  let instructions = parse_instructions(s)
  let registers = get_registers()
  run_and_count_mul(0, registers, instructions, 0)
}

pub fn part2(_: String) {
  // let instructions =  parse_instructions(s)
  // let registers = get_registers() |> dict.insert("a", 1)
  // run_program(0, registers, instructions)
  //
  // Solved in C
  //   gcc -o day23.exe inputs/day-23.c
  //   ./day23.exe
  915
}

pub fn solve() {
  let input = utils.read_all_file("inputs/day-23.txt") |> string.trim()
  io.println("Day 23")
  utils.time_it("Part 1", fn() { part1(input) |> int.to_string() })
  utils.time_it("Part 2", fn() { part2(input) |> int.to_string() })
}
