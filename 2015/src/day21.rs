use crate::utils;

mod game {
    #[derive(Debug, Clone)]
    pub struct Item {
        pub name: String,
        pub cost: i32,
        pub damage: i32,
        pub armor: i32,
    }

    impl Item {
        pub fn new(name: &str, cost: i32, damage: i32, armor: i32) -> Self {
            Self {
                name: name.to_string(),
                cost,
                damage,
                armor,
            }
        }
    }

    #[derive(Debug)]
    pub struct Shop {
        pub weapons: Vec<Item>,
        pub armors: Vec<Item>,
        pub rings: Vec<Item>,
    }

    impl Shop {
        pub fn new() -> Self {
            let weapons = vec![
                Item::new("Dagger", 8, 4, 0),
                Item::new("Shortsword", 10, 5, 0),
                Item::new("Warhammer", 25, 6, 0),
                Item::new("Longsword", 40, 7, 0),
                Item::new("Greataxe", 74, 8, 0),
            ];
            let armors = vec![
                Item::new("Leather", 13, 0, 1),
                Item::new("Chainmail", 31, 0, 2),
                Item::new("Splintmail", 53, 0, 3),
                Item::new("Bandedmail", 75, 0, 4),
                Item::new("Platemail", 102, 0, 5),
            ];
            let rings = vec![
                Item::new("Damage +1", 25, 1, 0),
                Item::new("Damage +2", 50, 2, 0),
                Item::new("Damage +3", 100, 3, 0),
                Item::new("Defense +1", 20, 0, 1),
                Item::new("Defense +2", 40, 0, 2),
                Item::new("Defense +3", 80, 0, 3),
            ];
            Self {
                weapons,
                armors,
                rings,
            }
        }
    }

    #[derive(Debug, Clone)]
    pub struct Stats {
        pub hit_points: i32,
        pub damage: i32,
        pub armor: i32,
    }

    impl Stats {
        pub fn new_player() -> Self {
            Self {
                hit_points: 100,
                armor: 0,
                damage: 0,
            }
        }

        pub fn parse(text: &str) -> Option<Self> {
            let mut lines = text.lines();
            let hit_points = lines.next()?["Hit Points: ".len()..].parse().ok()?;
            let damage = lines.next()?["Damage: ".len()..].parse().ok()?;
            let armor = lines.next()?["Armor: ".len()..].parse().ok()?;
            Some(Self {
                hit_points,
                damage,
                armor,
            })
        }
    }
}

use game::{Shop, Stats};

#[derive(Debug, Clone)]
struct Shopping {
    weapon_ix: i32,
    armor_ix: i32,
    ring1_ix: i32,
    ring2_ix: i32,

    armors_count: i32,
    weapons_count: i32,
    #[allow(unused)]
    rings_count: i32,
}

impl Shopping {
    pub fn new(shop: &Shop) -> Self {
        let weapons_count = shop.weapons.len() as i32;
        let armors_count = shop.armors.len() as i32;
        let rings_count = shop.rings.len() as i32;

        Self {
            weapon_ix: weapons_count - 1,
            armor_ix: armors_count,
            ring1_ix: rings_count,
            ring2_ix: rings_count,

            armors_count,
            weapons_count,
            rings_count,
        }
    }

    fn apply(&self, shop: &Shop, stats: &Stats) -> (Stats, i32) {
        let weapon = shop.weapons.get(self.weapon_ix as usize);
        let armor = shop.armors.get(self.armor_ix as usize);
        let ring1 = shop.rings.get(self.ring1_ix as usize);
        let ring2 = shop
            .rings
            .get(self.ring2_ix as usize)
            .filter(|_| self.ring1_ix != self.ring2_ix);

        let cost = armor.map_or(0, |item| item.cost)
            + weapon.map_or(0, |item| item.cost)
            + ring1.map_or(0, |item| item.cost)
            + ring2.map_or(0, |item| item.cost);

        let stats = Stats {
            armor: stats.armor
                + armor.map_or(0, |item| item.armor)
                + weapon.map_or(0, |item| item.armor)
                + ring1.map_or(0, |item| item.armor)
                + ring2.map_or(0, |item| item.armor),
            damage: stats.damage
                + armor.map_or(0, |item| item.damage)
                + weapon.map_or(0, |item| item.damage)
                + ring1.map_or(0, |item| item.damage)
                + ring2.map_or(0, |item| item.damage),
            hit_points: stats.hit_points,
        };

        (stats, cost)
    }
}

impl Iterator for Shopping {
    type Item = Shopping;

    fn next(&mut self) -> Option<Self::Item> {
        if self.ring2_ix == -1 {
            return None;
        }

        let result = self.clone();

        self.weapon_ix -= 1;
        if self.weapon_ix == -1 {
            self.weapon_ix = self.weapons_count - 1;
            self.armor_ix -= 1;
        }

        if self.armor_ix == -1 {
            self.armor_ix = self.armors_count;
            self.ring1_ix -= 1;
        }

        if self.ring1_ix == -1 {
            self.ring2_ix -= 1;
            self.ring1_ix = self.ring2_ix - 1;
        }

        if self.ring2_ix == 0 {
            self.ring2_ix -= 1;
        }

        Some(result)
    }
}

#[derive(Debug, PartialEq)]
pub enum Win {
    Player,
    Boss,
}

#[allow(unused)]
pub fn fight_iterative(player: &Stats, boss: &Stats) -> Win {
    let mut player = player.clone();
    let mut boss = boss.clone();
    let mut player_turn = true;
    while player.hit_points > 0 && boss.hit_points > 0 {
        if player_turn {
            boss.hit_points -= i32::max(1, player.damage - boss.armor);
        } else {
            player.hit_points -= i32::max(1, boss.damage - player.armor);
        }
        player_turn = !player_turn;
    }
    if player.hit_points > 0 {
        Win::Player
    } else {
        Win::Boss
    }
}

pub fn fight_math(player: &Stats, boss: &Stats) -> Win {
    let player_damage = i32::max(1, player.damage - boss.armor);
    let boss_damage = i32::max(1, boss.damage - player.armor);

    let player_hits = boss.hit_points as f32 / player_damage as f32;
    let boss_hits = player.hit_points as f32 / boss_damage as f32;

    if boss_hits.ceil() >= player_hits.ceil() {
        Win::Player
    } else {
        Win::Boss
    }
}

fn minimum_cost_to_win(player: &Stats, boss: &Stats) -> i32 {
    let shop = Shop::new();

    Shopping::new(&shop)
        .map(|ix| ix.apply(&shop, player))
        .filter(|(player, _)| fight_math(player, boss) == Win::Player)
        .map(|(_, cost)| cost)
        .min()
        .unwrap_or(i32::MAX)
}

fn part1(content: &str) -> i32 {
    let player = Stats::new_player();
    let boss = Stats::parse(content).expect("invalid input");
    minimum_cost_to_win(&player, &boss)
}

fn maximum_cost_to_lose(player: &Stats, boss: &Stats) -> i32 {
    let shop = Shop::new();

    Shopping::new(&shop)
        .map(|ix| ix.apply(&shop, player))
        .filter(|(player, _)| fight_math(player, boss) == Win::Boss)
        .map(|(_, cost)| cost)
        .max()
        .unwrap_or(i32::MIN)
}

fn part2(content: &str) -> i32 {
    let player = Stats::new_player();
    let boss = Stats::parse(content).expect("invalid input");
    maximum_cost_to_lose(&player, &boss)
}

pub fn solve() -> std::io::Result<()> {
    let mut content = String::new();
    utils::read_file("inputs/day-21.txt", &mut content)?;

    println!("Day 21");
    time_it!("Part 1", part1(&content));
    time_it!("Part 2", part2(&content));

    Ok(())
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1_fight_01() {
        let player = Stats {
            hit_points: 8,
            damage: 5,
            armor: 5,
        };
        let boss = Stats {
            hit_points: 12,
            damage: 7,
            armor: 2,
        };

        assert_eq!(Win::Player, fight_math(&player, &boss));
        assert_eq!(Win::Player, fight_iterative(&player, &boss));
    }

    #[test]
    fn test_part1_minimum_cost_to_win() {
        let player = Stats {
            hit_points: 8,
            damage: 0,
            armor: 0,
        };
        let boss = Stats {
            hit_points: 12,
            damage: 7,
            armor: 2,
        };

        assert_eq!(65, minimum_cost_to_win(&player, &boss), "Player must wins");
    }
}
