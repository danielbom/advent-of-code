import gleam/int
import gleam/list
import gleam/result

import argv

import day_01
import day_02
import day_03
import day_04
import day_05
import day_06
import day_07
import day_08
import day_09
import day_10
import day_11
import day_12
import day_13

fn foreach(begin: Int, end: Int, func: fn(Int) -> Nil) -> Nil {
  case begin <= end {
    True -> {
      func(begin)
      foreach(begin + 1, end, func)
    }
    _ -> Nil
  }
}

fn run(day: Int) -> Nil {
  case day {
    0 -> foreach(1, 25, run)
    1 -> day_01.solve()
    2 -> day_02.solve()
    3 -> day_03.solve()
    4 -> day_04.solve()
    5 -> day_05.solve()
    6 -> day_06.solve()
    7 -> day_07.solve()
    8 -> day_08.solve()
    9 -> day_09.solve()
    10 -> day_10.solve()
    11 -> day_11.solve()
    12 -> day_12.solve()
    13 -> day_13.solve()
    14 -> Nil
    15 -> Nil
    16 -> Nil
    17 -> Nil
    18 -> Nil
    19 -> Nil
    20 -> Nil
    21 -> Nil
    22 -> Nil
    23 -> Nil
    24 -> Nil
    25 -> Nil
    _ -> panic as "invalid day"
  }
}

pub fn main() -> Nil {
  let args =
    argv.load().arguments
    |> list.map(int.parse)
    |> result.values()
  case args {
    [day] -> run(day)
    _ -> panic as "invalid amount of arguments"
  }
}
