use std::collections::{HashMap, HashSet};

#[derive(Debug, Clone)]
struct Road<'a> {
    city1: &'a str,
    city2: &'a str,
    distance: i32,
}

impl<'a> PartialEq for Road<'a> {
    fn eq(&self, other: &Self) -> bool {
        self.city1.eq(other.city1) && self.city2.eq(other.city2)
    }
}

impl<'a> Road<'a> {
    fn parse(line: &str) -> Option<Road> {
        let mut parts = line.split(" = ");
        let cities = parts.next()?;
        let distance = parts.next()?.parse::<i32>().ok()?;
        let mut cities = cities.split(" to ");
        let city1 = cities.next()?;
        let city2 = cities.next()?;
        Some(Road {
            city1,
            city2,
            distance,
        })
    }
}

type Graph<'a> = HashMap<&'a str, HashMap<&'a str, i32>>;

#[derive(Debug)]
struct DfsNode<'a> {
    key: &'a str,
    distance: i32,
    depth: u16,
}

impl DfsNode<'_> {
    fn new(node: &str) -> DfsNode {
        DfsNode {
            key: node,
            distance: 0,
            depth: 0,
        }
    }
}

fn dfs<'a, Consume>(
    graph: &'a Graph<'a>,
    node: &DfsNode<'a>,
    visited: &HashSet<&'a str>,
    consume: &mut Consume,
) where
    Consume: FnMut(&DfsNode<'a>),
{
    let current_visited = {
        let mut current_visited = visited.clone();
        current_visited.insert(node.key);
        current_visited
    };

    for (&child, &next_distance) in graph.get(node.key).unwrap() {
        if !visited.contains(child) {
            let current_distance = node.distance + next_distance;
            let new_node = DfsNode {
                key: child,
                distance: current_distance,
                depth: node.depth + 1,
            };
            consume(&new_node);
            dfs(graph, &new_node, &current_visited, consume);
        }
    }
}

fn parse_content(content: &str) -> Vec<Road> {
    content.lines().flat_map(Road::parse).collect()
}

fn collect_cities<'a>(roads: &'a [Road<'a>]) -> Vec<&'a str> {
    roads
        .iter()
        .flat_map(|it| vec![it.city1, it.city2])
        .collect::<HashSet<_>>()
        .into_iter()
        .collect()
}

fn build_graph<'a>(roads: &'a Vec<Road<'a>>) -> Graph<'a> {
    let mut graph = HashMap::new();
    for road in roads {
        graph
            .entry(road.city1)
            .or_insert(HashMap::new())
            .insert(road.city2, road.distance);
        graph
            .entry(road.city2)
            .or_insert(HashMap::new())
            .insert(road.city1, road.distance);
    }
    graph
}

fn part1(content: &str) -> i32 {
    let roads = parse_content(content);
    let cities = collect_cities(&roads);
    let graph = build_graph(&roads);
    let cities_count = cities.len() as u16;
    let mut min_distance = i32::MAX;
    let find_min_distance = &mut |it: &DfsNode| {
        if it.depth == cities_count - 1 && it.distance < min_distance {
            min_distance = it.distance;
        }
    };

    cities.iter().for_each(|city| {
        dfs(
            &graph,
            &DfsNode::new(city),
            &HashSet::new(),
            find_min_distance,
        )
    });

    min_distance
}

fn part2(content: &str) -> i32 {
    let roads = parse_content(content);
    let cities = collect_cities(&roads);
    let graph = build_graph(&roads);
    let cities_count = cities.len() as u16;
    let mut max_distance = 0;
    let find_min_distance = &mut |it: &DfsNode| {
        if it.depth == cities_count - 1 && it.distance > max_distance {
            max_distance = it.distance;
        }
    };

    cities.iter().for_each(|city| {
        dfs(
            &graph,
            &DfsNode::new(city),
            &HashSet::new(),
            find_min_distance,
        )
    });

    max_distance
}

pub fn solve() -> std::io::Result<()> {
    let mut content = String::new();
    aoc2015::read_file("inputs/day-09.txt", &mut content)?;

    println!("Day 09");
    println!("Part 1: {}", part1(&content));
    println!("Part 2: {}", part2(&content));

    Ok(())
}
