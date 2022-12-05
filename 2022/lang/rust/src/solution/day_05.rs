use std::{fs, io};

// This isn't very necessary, I was just interested in trying to use some different
// language features.
type Stack = Vec<char>;
type Stacks = Vec<Stack>;

fn stacks_new(initial_state: String) -> Stacks {
    let mut stacks: Stacks = Vec::new();
    initial_state.lines().rev().for_each(|row| {
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

pub fn solve() -> io::Result<()> {
    println!("- Day 05");

    let input = fs::read_to_string("input/day-05.txt")?;
    part_one(&input);

    Ok(())
}

fn part_one(input: &str) {
    let splits: Vec<&str> = input.split("\n\n").collect();
    let mut stacks = stacks_new(splits[0].to_string());

    splits[1]
        .lines()
        .into_iter()
        .map(|m| {
            m.split_whitespace()
                .filter_map(|word| word.parse::<usize>().ok())
                .collect::<Vec<usize>>()
        })
        .for_each(|m| {
            for _ in 0..m[0] {
                if let Some(last) = stacks[m[1] - 1].pop() {
                    stacks[m[2] - 1].push(last);
                }
            }
        });

    let top = stacks.iter().filter_map(|s| s.last()).collect::<String>();
    println!("  - Part 1: {top}");
}
