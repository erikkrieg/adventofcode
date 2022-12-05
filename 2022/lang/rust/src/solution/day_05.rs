use std::{fs, io};

// Using type alias and a custom trait just for practice with more language features
type Stack = Vec<char>;
type Stacks = Vec<Stack>;

trait Top {
    fn top(&self) -> String;
}

impl Top for Stacks {
    fn top(&self) -> String {
        self.iter().filter_map(|s| s.last()).collect::<String>()
    }
}

pub fn solve() -> io::Result<()> {
    println!("- Day 05");

    let input = fs::read_to_string("input/day-05.txt")?;
    let splits: Vec<&str> = input.split("\n\n").collect();
    let stacks = parse_stacks(splits[0]);
    let moves = parse_moves(splits[1]);
    part_one(&stacks, &moves);
    part_two(&stacks, &moves);

    Ok(())
}

fn part_one(stacks: &Stacks, moves: &[Vec<usize>]) {
    let mut stacks = stacks.clone();
    moves.iter().for_each(|m| {
        for _ in 0..m[0] {
            if let Some(last) = stacks[m[1] - 1].pop() {
                stacks[m[2] - 1].push(last);
            }
        }
    });
    println!("  - Part 1: {}", stacks.top());
}

fn part_two(stacks: &Stacks, moves: &[Vec<usize>]) {
    let mut stacks = stacks.clone();
    moves.iter().for_each(|m| {
        let split_at = stacks[m[1] - 1].len() - m[0];
        let mut move_stacks = stacks[m[1] - 1].split_off(split_at);
        stacks[m[2] - 1].append(&mut move_stacks);
    });
    println!("  - Part 2: {}", stacks.top());
}

fn parse_stacks(raw_stacks: &str) -> Stacks {
    let mut stacks: Stacks = Vec::new();
    raw_stacks.lines().rev().for_each(|row| {
        row.chars()
            .enumerate()
            .filter(|(i, _)| (*i as i32 - 1) % 4 == 0)
            .for_each(|(i, c)| {
                if c.is_uppercase() {
                    let i = i.div_floor(4);
                    stacks[i].push(c);
                } else if c.is_numeric() {
                    stacks.push(Vec::new());
                }
            })
    });
    stacks
}

fn parse_moves(moves: &str) -> Vec<Vec<usize>> {
    moves
        .lines()
        .into_iter()
        .map(|m| {
            m.split_whitespace()
                .filter_map(|word| word.parse::<usize>().ok())
                .collect::<Vec<usize>>()
        })
        .collect()
}
