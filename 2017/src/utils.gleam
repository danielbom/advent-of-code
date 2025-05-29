import gleam/float
import gleam/int
import gleam/io
import gleam/string

import gleam/time/duration
import gleam/time/timestamp

import simplifile

pub fn read_all_file(filepath: String) -> String {
  let assert Ok(content) = simplifile.read(from: filepath)
  content
}

pub fn time_it(name: String, func: fn() -> String) {
  let start_time = timestamp.system_time()
  let result = func()
  let end_time = timestamp.system_time()
  let #(duration_secs, duration_nanos) =
    start_time
    |> timestamp.difference(end_time)
    |> duration.to_seconds_and_nanoseconds()
  let milis =
    { int.to_float(duration_secs) *. 1.0e3 }
    +. { int.to_float(duration_nanos) /. 1.0e6 }
  let message =
    string.concat([name, ": ", result, " [", float.to_string(milis), "ms]"])
  io.println(message)
}
