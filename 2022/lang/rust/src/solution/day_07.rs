use std::collections::HashMap;
use std::fs::File;
use std::io::{self, prelude::*, BufReader};
use std::path::PathBuf;

const TOTAL_SPACE: u32 = 70_000_000;
const REQUIRED_SPACE: u32 = 30_000_000;
const MAX_SIZE: u32 = 100_000;

type Dirs = HashMap<String, u32>;

pub fn solve() -> io::Result<()> {
    println!("- Day 07");
    let input = File::open("input/day-07.txt")?;
    let dirs = dir_sizes(BufReader::new(input));

    let part_one: u32 = dirs.values().filter(|&v| v <= &MAX_SIZE).sum();
    println!("  - Part 1: {part_one}");

    let root_size = dirs.get("/").unwrap();
    let available_space = TOTAL_SPACE - root_size;
    let space_to_free = REQUIRED_SPACE - available_space;
    let rm_dir_size = dirs.values().filter(|&d| d > &space_to_free).min();
    println!("  - Part 2: {}", rm_dir_size.unwrap());

    Ok(())
}

fn dir_sizes(input: BufReader<File>) -> Dirs {
    let (_, dirs) = input.lines().filter_map(|l| l.ok()).fold(
        (PathBuf::from(""), Dirs::new()),
        |(mut path, mut dirs), x| {
            let s: Vec<&str> = x.split_whitespace().collect();
            if s[1] == "cd" {
                if s[2] == ".." {
                    path.pop();
                } else {
                    path.push(s[2]);
                }
            } else if let Ok(size) = s[0].parse::<u32>() {
                let mut p = Some(path.as_path());
                while let Some(dir) = p {
                    let d = dir.to_str().unwrap().to_string();
                    dirs.entry(d).and_modify(|v| *v += size).or_insert(size);
                    p = dir.parent();
                }
            }
            (path, dirs)
        },
    );
    dirs
}
