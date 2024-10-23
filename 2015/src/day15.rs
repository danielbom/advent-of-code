use crate::utils;
#[derive(Debug, Clone)]
struct Ingredient {
    #[allow(dead_code)]
    name: String,
    capacity: i64,
    durability: i64,
    flavor: i64,
    texture: i64,
    calories: i64,
}

impl Ingredient {
    fn parse(line: &str) -> Option<Ingredient> {
        // Sprinkles: capacity 5, durability -1, flavor 0, texture 0, calories 5
        let mut parts = line.split(": ");
        let name = parts.next()?.to_string();
        let mut parts = parts.next()?.split(", ").flat_map(|it| -> Option<i64> {
            let mut tuple = it.split_whitespace();
            let _name = tuple.next()?;
            let value = tuple.next()?.parse::<i64>().ok()?;
            Some(value)
        });
        let capacity = parts.next()?;
        let durability = parts.next()?;
        let flavor = parts.next()?;
        let texture = parts.next()?;
        let calories = parts.next()?;
        Some(Ingredient {
            name,
            capacity,
            durability,
            flavor,
            texture,
            calories,
        })
    }

    fn parse_lines(content: &str) -> Vec<Ingredient> {
        content.split('\n').flat_map(Ingredient::parse).collect()
    }
}

fn quantities_of_ingredients(quantities: &[i64], ingredients: &[Ingredient]) -> Vec<i64> {
    let mut scores = vec![0_i64; 5];

    ingredients
        .iter()
        .zip(quantities.iter())
        .map(|(ingredient, quantity)| {
            vec![
                ingredient.capacity * quantity,
                ingredient.durability * quantity,
                ingredient.flavor * quantity,
                ingredient.texture * quantity,
                ingredient.calories * quantity,
            ]
        })
        .for_each(|props| {
            scores
                .iter_mut()
                .zip(props.iter())
                .for_each(|(v, p)| *v += p);
        });

    scores.iter_mut().for_each(|x| *x = *x.max(&mut 0_i64));

    scores
}

fn create_quantities(n: usize, i: i64) -> Vec<i64> {
    let mut quantities = vec![0; n];
    let mut index = 0;
    let mut x = i;
    while x > 0 {
        quantities[index] = x % 101;
        x /= 101;
        index += 1;
    }
    quantities
}

fn compute_total_score(scores: &[i64]) -> i64 {
    scores.iter().rev().skip(1).product()
}

fn quantities_constraint(quantities: &[i64]) -> bool {
    quantities.iter().sum::<i64>() == 100
}

fn part1(input: &str) -> i64 {
    let ingredients = Ingredient::parse_lines(input);
    let top = 100_i64.pow(ingredients.len() as u32);
    let n = ingredients.len();

    (0..top)
        .step_by(100)
        .map(|i| create_quantities(n, i))
        .filter(|quantities| quantities_constraint(quantities))
        .map(|quantities| quantities_of_ingredients(&quantities, &ingredients))
        .map(|scores| compute_total_score(&scores))
        .max()
        .unwrap()
}

fn calories_constraint(scores: &[i64]) -> bool {
    scores.last().unwrap() == &500_i64
}

fn part2(input: &str) -> i64 {
    let ingredients = Ingredient::parse_lines(input);
    let top = 100_i64.pow(ingredients.len() as u32);
    let n = ingredients.len();

    (0..top)
        .step_by(100)
        .map(|i| create_quantities(n, i))
        .filter(|quantities| quantities_constraint(quantities))
        .map(|quantities| quantities_of_ingredients(&quantities, &ingredients))
        .filter(|scores| calories_constraint(scores))
        .map(|scores| compute_total_score(&scores))
        .max()
        .unwrap()
}

pub fn solve() -> std::io::Result<()> {
    let mut content = String::new();
    utils::read_file("inputs/day-15.txt", &mut content)?;

    println!("Day 15");
    time_it!("Part 1", part1(&content));
    time_it!("Part 2", part2(&content));

    Ok(())
}

#[cfg(test)]
mod tests {
    use super::*;

    #[inline]
    fn input_sample<'a>() -> &'a str {
        "Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8
Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3"
    }

    #[test]
    fn parse_lines_must_works() {
        let content = input_sample();
        let result = Ingredient::parse_lines(content).len();
        let expected = 2;
        assert_eq!(expected, result);
    }

    #[test]
    fn quantities_of_ingredients_must_works() {
        let content = input_sample();
        let ingredients = Ingredient::parse_lines(content);
        let quantities = vec![44, 56];
        let result = quantities_of_ingredients(&quantities, &ingredients);
        let expected = vec![68, 80, 152, 76, 520];
        assert_eq!(expected, result);
    }
}
