import gleam/int
import gleam/list
import gleam/result

import argv

import day_01
import day_02
import day_03

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
    0 -> {
      foreach(1, 25, fn(i) { run(i) })
    }
    1 -> {
      day_01.solve()
    }
    2 -> {
      day_02.solve()
    }
    3 -> {
      day_03.solve()
    }
    4 -> {
      Nil
    }
    5 -> {
      Nil
    }
    6 -> {
      Nil
    }
    7 -> {
      Nil
    }
    8 -> {
      Nil
    }
    9 -> {
      Nil
    }
    10 -> {
      Nil
    }
    11 -> {
      Nil
    }
    12 -> {
      Nil
    }
    13 -> {
      Nil
    }
    14 -> {
      Nil
    }
    15 -> {
      Nil
    }
    16 -> {
      Nil
    }
    17 -> {
      Nil
    }
    18 -> {
      Nil
    }
    19 -> {
      Nil
    }
    20 -> {
      Nil
    }
    21 -> {
      Nil
    }
    22 -> {
      Nil
    }
    23 -> {
      Nil
    }
    24 -> {
      Nil
    }
    25 -> {
      Nil
    }
    _ -> {
      panic as "invalid day"
    }
  }
}

pub fn main() -> Nil {
  let args =
    argv.load().arguments
    |> list.map(int.parse)
    |> result.values()
  case args {
    [] -> run(0)
    [day] -> run(day)
    _ -> panic as "invalid amount of arguments"
  }
}
