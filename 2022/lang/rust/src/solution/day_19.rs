use std::{
    collections::VecDeque,
    fs::File,
    io::{self, BufRead, BufReader},
};

#[derive(Debug)]
struct Blueprint {
    ore_robot: Cost,
    clay_robot: Cost,
    obsidian_robot: Cost,
    geode_robot: Cost,
}

impl Blueprint {
    fn max_ore_cost(&self) -> usize {
        *[
            self.ore_robot.ore,
            self.clay_robot.ore,
            self.obsidian_robot.ore,
            self.geode_robot.ore,
        ]
        .iter()
        .max()
        .unwrap()
    }
    fn max_clay_cost(&self) -> usize {
        *[
            self.ore_robot.clay,
            self.clay_robot.clay,
            self.obsidian_robot.clay,
            self.geode_robot.clay,
        ]
        .iter()
        .max()
        .unwrap()
    }
}

#[derive(Debug)]
struct Cost {
    ore: usize,
    clay: usize,
    obsidian: usize,
}

fn parse_blueprints(reader: BufReader<File>) -> Vec<Blueprint> {
    reader
        .lines()
        .filter_map(|line| line.ok())
        .map(|line| {
            let mut nums = line
                .split_whitespace()
                .filter_map(|s| s.parse::<usize>().ok());
            Blueprint {
                ore_robot: Cost {
                    ore: nums.next().unwrap(),
                    clay: 0,
                    obsidian: 0,
                },
                clay_robot: Cost {
                    ore: nums.next().unwrap(),
                    clay: 0,
                    obsidian: 0,
                },
                obsidian_robot: Cost {
                    ore: nums.next().unwrap(),
                    clay: nums.next().unwrap(),
                    obsidian: 0,
                },
                geode_robot: Cost {
                    ore: nums.next().unwrap(),
                    clay: 0,
                    obsidian: nums.next().unwrap(),
                },
            }
        })
        .collect()
}

#[derive(Default, Debug, Clone)]
struct SimulationState {
    ore_robots: usize,
    clay_robots: usize,
    obsidian_robots: usize,
    geode_robots: usize,
    ore_count: usize,
    clay_count: usize,
    obsidian_count: usize,
    geode_count: usize,
    minutes_passed: usize,
}

impl SimulationState {
    fn new() -> Self {
        Self {
            ore_robots: 1,
            ..Default::default()
        }
    }
}

enum Action {
    Wait,
    BuyOreBot,
    BuyClayBot,
    BuyObsidianBot,
    BuyGeodeBot,
}

fn simulate(bp: &Blueprint, duration: usize) -> usize {
    let mut sims = VecDeque::from([SimulationState::new()]);
    let mut max_geode_count = 0;
    while let Some(mut state) = sims.pop_front() {
        if state.minutes_passed == duration
            || (state.ore_robots + state.clay_robots + state.obsidian_robots + state.geode_robots)
                + 1
                < state.minutes_passed.div_floor(2)
        {
            /*if state.minutes_passed == duration && state.geode_count > max_geode_count {
                println!("{state:?}");
            }*/
            max_geode_count = max_geode_count.max(state.geode_count);
            continue;
        }

        use Action::*;
        let next_ore = state.ore_count + state.ore_robots;
        let next_clay = state.clay_count + state.clay_robots;
        let next_obsidian = state.obsidian_count + state.obsidian_robots;
        let next_geode = state.geode_count + state.geode_robots;
        let mut available_actions = Vec::new();
        if state.ore_count >= bp.geode_robot.ore && state.obsidian_count >= bp.geode_robot.obsidian
        {
            available_actions.push(BuyGeodeBot);
        } else if state.ore_count >= bp.obsidian_robot.ore
            && state.clay_count >= bp.obsidian_robot.clay
        {
            if next_ore >= bp.geode_robot.ore && next_obsidian >= bp.geode_robot.obsidian {
                available_actions.push(Wait);
            }
            available_actions.push(BuyObsidianBot);
            if state.ore_count >= bp.clay_robot.ore
                && state.clay_robots < bp.max_clay_cost().div_floor(2)
            {
                available_actions.push(BuyClayBot);
            }
        } else {
            if state.ore_count >= bp.ore_robot.ore && state.ore_robots < bp.max_ore_cost() {
                available_actions.push(BuyOreBot);
            }
            if state.ore_count >= bp.clay_robot.ore && state.clay_robots < bp.max_clay_cost() {
                available_actions.push(BuyClayBot);
            }
            available_actions.push(Wait);
        }

        // update state
        state.ore_count = next_ore;
        state.clay_count = next_clay;
        state.obsidian_count = next_obsidian;
        state.geode_count = next_geode;
        state.minutes_passed += 1;

        available_actions.iter().for_each(|a| match a {
            BuyOreBot => {
                let mut state = state.clone();
                state.ore_robots += 1;
                state.ore_count -= bp.ore_robot.ore;
                sims.push_back(state);
            }
            BuyClayBot => {
                let mut state = state.clone();
                state.clay_robots += 1;
                state.ore_count -= bp.clay_robot.ore;
                sims.push_back(state);
            }
            BuyObsidianBot => {
                let mut state = state.clone();
                state.obsidian_robots += 1;
                state.ore_count -= bp.obsidian_robot.ore;
                state.clay_count -= bp.obsidian_robot.clay;
                sims.push_back(state);
            }
            BuyGeodeBot => {
                let mut state = state.clone();
                state.geode_robots += 1;
                state.ore_count -= bp.geode_robot.ore;
                state.obsidian_count -= bp.geode_robot.obsidian;
                sims.push_back(state);
            }
            Wait => {
                sims.push_back(state.clone());
            }
        });
    }
    println!("{max_geode_count}");
    max_geode_count
}

pub fn solve() -> io::Result<()> {
    println!("- Day 19:");
    let input = File::open("input/day-19.txt")?;
    //let input = File::open("input/sample-19.txt")?;
    let blueprints = parse_blueprints(BufReader::new(input));
    part_one(&blueprints);
    part_two(&blueprints);
    Ok(())
}

fn part_one(blueprints: &[Blueprint]) {
    let duration = 24;
    let quality_level_sum: usize = blueprints
        .iter()
        .enumerate()
        .map(|(i, b)| (i + 1) * simulate(b, duration))
        .sum();
    println!("  - Part 1: {quality_level_sum}");
}

fn part_two(blueprints: &[Blueprint]) {
    let duration = 32;
    let product: usize = blueprints
        .iter()
        .take(3)
        .map(|b| simulate(b, duration))
        .product();
    println!("  - Part 2: {product}");
}
