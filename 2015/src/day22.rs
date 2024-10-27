use crate::utils;

#[derive(Debug, Clone)]
struct Stats {
    pub hit_points: i32,
    pub damage: i32,
    pub armor: i32,
    pub mana: i32,
}

impl Stats {
    pub fn new_player() -> Self {
        Self {
            hit_points: 50,
            damage: 0,
            armor: 0,
            mana: 500,
        }
    }

    pub fn parse(text: &str) -> Option<Self> {
        let mut lines = text.lines();
        let hit_points = lines.next()?["Hit Points: ".len()..].parse().ok()?;
        let damage = lines.next()?["Damage: ".len()..].parse().ok()?;
        Some(Self {
            hit_points,
            damage,
            armor: 0,
            mana: 0,
        })
    }
}

struct Spell {
    mana: i32,
    timer: i32,
    apply_damage: i32,
    add_armor: i32,
    restore_life: i32,
    restore_mana: i32,
}

impl Spell {
    fn new(
        mana: i32,
        timer: i32,
        apply_damage: i32,
        add_armor: i32,
        restore_life: i32,
        restore_mana: i32,
    ) -> Self {
        Self {
            mana,
            timer,
            apply_damage,
            add_armor,
            restore_life,
            restore_mana,
        }
    }

    fn available(&self, player: &Stats) -> bool {
        self.mana <= player.mana
    }
}

struct Spells {
    magic_missile: Spell,
    drain: Spell,
    shield: Spell,
    poison: Spell,
    recharge: Spell,
}

impl Spells {
    fn new() -> Self {
        Self {
            magic_missile: Spell::new(53, 0, 4, 0, 0, 0),
            drain: Spell::new(73, 0, 2, 0, 2, 0),
            shield: Spell::new(113, 6, 0, 7, 0, 0),
            poison: Spell::new(173, 6, 3, 0, 0, 0),
            recharge: Spell::new(229, 5, 0, 0, 0, 101),
        }
    }
}

#[derive(Debug, PartialEq)]
enum Win {
    Player,
    Boss,
}

#[derive(Clone)]
struct Duel {
    player: Stats,
    boss: Stats,
    mana_spent: i32,
    timer_poison: i32,
    timer_shield: i32,
    timer_recharge: i32,
}

impl Duel {
    fn new(player: Stats, boss: Stats) -> Self {
        Self {
            player,
            boss,
            mana_spent: 0,
            timer_poison: -1,
            timer_shield: -1,
            timer_recharge: -1,
        }
    }

    fn cast_magic_missile(&mut self, spells: &Spells) {
        self.player.mana -= spells.magic_missile.mana;
        self.mana_spent += spells.magic_missile.mana;
        self.boss.hit_points -= spells.magic_missile.apply_damage;
    }

    fn cast_drain(&mut self, spells: &Spells) {
        self.player.mana -= spells.drain.mana;
        self.mana_spent += spells.drain.mana;
        self.boss.hit_points -= spells.drain.apply_damage;
        self.player.hit_points += spells.drain.restore_life;
    }

    fn cast_shield(&mut self, spells: &Spells) {
        self.player.mana -= spells.shield.mana;
        self.mana_spent += spells.shield.mana;
        self.timer_shield = spells.shield.timer;
        self.player.armor += spells.shield.add_armor;
    }

    fn cast_poison(&mut self, spells: &Spells) {
        self.player.mana -= spells.poison.mana;
        self.mana_spent += spells.poison.mana;
        self.timer_poison = spells.poison.timer;
    }

    fn cast_recharge(&mut self, spells: &Spells) {
        self.player.mana -= spells.recharge.mana;
        self.mana_spent += spells.recharge.mana;
        self.timer_recharge = spells.recharge.timer;
    }

    fn run_effects(&mut self, spells: &Spells) {
        if self.timer_poison > 0 {
            self.timer_poison -= 1;
            self.boss.hit_points -= spells.poison.apply_damage;
        }

        if self.timer_shield >= 0 {
            self.timer_shield -= 1;
        }
        if self.timer_shield == 0 {
            self.player.armor -= spells.shield.add_armor;
        }

        if self.timer_recharge > 0 {
            self.timer_recharge -= 1;
            self.player.mana += spells.recharge.restore_mana;
        }
    }

    fn run_boss(&mut self) {
        self.player.hit_points -= i32::max(1, self.boss.damage - self.player.armor);
    }

    fn running(&self) -> bool {
        self.player.hit_points > 0 && self.boss.hit_points > 0
    }

    fn step(&mut self, spells: &Spells) -> bool {
        if !self.running() {
            return false;
        }

        self.run_effects(spells);
        if !self.running() {
            return false;
        }

        self.run_boss();
        if !self.running() {
            return false;
        }
        self.run_effects(spells);
        self.running()
    }

    fn win(&self) -> Win {
        if self.player.hit_points > 0 {
            Win::Player
        } else {
            Win::Boss
        }
    }
}

enum Mode {
    Easy,
    Hard,
}

fn minimum_mana_spent(player: &Stats, boss: &Stats, mode: Mode) -> i32 {
    let mut min_mana_spent = i32::MAX;
    let spells = &Spells::new();
    let penalty = match mode {
        Mode::Easy => 0,
        Mode::Hard => 1,
    };
    let mut stack = vec![Duel::new(player.clone(), boss.clone())];

    while let Some(mut new_duel) = stack.pop() {
        new_duel.player.hit_points -= penalty;

        if new_duel.running() {
            if spells.recharge.available(&new_duel.player) {
                let mut new_duel = Duel::clone(&new_duel);
                new_duel.cast_recharge(spells);
                new_duel.step(spells);
                if new_duel.mana_spent < min_mana_spent {
                    stack.push(new_duel);
                }
            }

            if spells.shield.available(&new_duel.player) {
                let mut new_duel = Duel::clone(&new_duel);
                new_duel.cast_shield(spells);
                new_duel.step(spells);
                if new_duel.mana_spent < min_mana_spent {
                    stack.push(new_duel);
                }
            }

            if spells.poison.available(&new_duel.player) {
                let mut new_duel = Duel::clone(&new_duel);
                new_duel.cast_poison(spells);
                new_duel.step(spells);
                if new_duel.mana_spent < min_mana_spent {
                    stack.push(new_duel);
                }
            }

            if spells.drain.available(&new_duel.player) {
                let mut new_duel = Duel::clone(&new_duel);
                new_duel.cast_drain(spells);
                new_duel.step(spells);
                if new_duel.mana_spent < min_mana_spent {
                    stack.push(new_duel);
                }
            }

            if spells.magic_missile.available(&new_duel.player) {
                let mut new_duel = Duel::clone(&new_duel);
                new_duel.cast_magic_missile(spells);
                new_duel.step(spells);
                if new_duel.mana_spent < min_mana_spent {
                    stack.push(new_duel);
                }
            }
        } else if new_duel.win() == Win::Player {
            min_mana_spent = new_duel.mana_spent;
        }
    }

    min_mana_spent
}

fn part1(content: &str) -> i32 {
    let player = Stats::new_player();
    let boss = Stats::parse(content).expect("invalid input");
    minimum_mana_spent(&player, &boss, Mode::Easy)
}

fn part2(content: &str) -> i32 {
    let player = Stats::new_player();
    let boss = Stats::parse(content).expect("invalid input");
    minimum_mana_spent(&player, &boss, Mode::Hard)
}

pub fn solve() -> std::io::Result<()> {
    let mut content = String::new();
    utils::read_file("inputs/day-22.txt", &mut content)?;

    println!("Day 22");
    time_it!("Part 1", part1(&content));
    time_it!("Part 2", part2(&content));

    Ok(())
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1_01() {
        let boss = Stats {
            hit_points: 13,
            damage: 8,
            armor: 0,
            mana: 0,
        };
        let player = Stats {
            hit_points: 10,
            damage: 0,
            armor: 0,
            mana: 250,
        };
        let result = minimum_mana_spent(&player, &boss, Mode::Easy);
        let expected = 226;
        assert_eq!(expected, result);
    }

    #[test]
    fn test_part1_simulation1() {
        let boss = Stats {
            hit_points: 13,
            damage: 8,
            armor: 0,
            mana: 0,
        };
        let player = Stats {
            hit_points: 10,
            damage: 0,
            armor: 0,
            mana: 250,
        };
        let spells = Spells::new();
        let mut duel = Duel::new(player.clone(), boss.clone());

        duel.cast_poison(&spells);
        assert_eq!(duel.player.hit_points, 10);
        assert_eq!(duel.player.mana, 77);
        assert_eq!(duel.boss.hit_points, 13);

        assert!(duel.step(&spells), "1nd");
        assert_eq!(duel.player.hit_points, 2);
        assert_eq!(duel.player.mana, 77);
        assert_eq!(duel.boss.hit_points, 7);

        duel.cast_magic_missile(&spells);
        assert_eq!(duel.player.hit_points, 2);
        assert_eq!(duel.player.mana, 24);
        assert_eq!(duel.boss.hit_points, 3);

        assert!(!duel.step(&spells), "2nd");
        assert_eq!(duel.player.hit_points, 2);
        assert_eq!(duel.player.mana, 24);
        assert_eq!(duel.boss.hit_points, 0);

        assert_eq!(duel.win(), Win::Player);
    }
}
