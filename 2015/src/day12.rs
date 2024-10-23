use crate::utils;
use serde_json::Value;

fn json_sum_numbers(value: &Value) -> i64 {
    if let Some(val) = value.as_i64() {
        val
    } else if let Some(array) = value.as_array() {
        array.iter().map(json_sum_numbers).sum()
    } else if let Some(object) = value.as_object() {
        object.values().map(json_sum_numbers).sum()
    } else {
        0
    }
}

fn json_is_red(value: &Value) -> bool {
    value.as_str().map(|it| it == "red").unwrap_or(false)
}

fn json_sum_numbers_without_red_value(value: &Value) -> i64 {
    if let Some(val) = value.as_i64() {
        val
    } else if let Some(object) = value.as_object() {
        if object.values().any(json_is_red) {
            0
        } else {
            object
                .values()
                .map(json_sum_numbers_without_red_value)
                .sum()
        }
    } else if let Some(array) = value.as_array() {
        array.iter().map(json_sum_numbers_without_red_value).sum()
    } else {
        0
    }
}

fn part1(input: &str) -> i64 {
    let value: Value = serde_json::from_str(input).unwrap();
    json_sum_numbers(&value)
}

fn part2(input: &str) -> i64 {
    let value: Value = serde_json::from_str(input).unwrap();
    json_sum_numbers_without_red_value(&value)
}

pub fn solve() -> std::io::Result<()> {
    let mut content = String::new();
    utils::read_file("inputs/day-12.txt", &mut content)?;

    println!("Day 12");
    time_it!("Part 1", part1(&content));
    time_it!("Part 2", part2(&content));

    Ok(())
}
