use std::{
    collections::{BTreeSet, VecDeque},
    fs::File,
    io::{self, BufRead, BufReader},
};

pub fn solve() -> io::Result<()> {
    println!("- Day 06");
    let file = File::open("input/day-06.txt")?;
    part_one(BufReader::new(file));
    Ok(())
}

fn part_one(reader: BufReader<File>) {
    let result: usize = reader
        .lines()
        .filter_map(|l| l.ok())
        .map(|l| {
            let mut i = 4;
            let mut window = l.chars().take(i).collect::<VecDeque<char>>();
            let mut rest = l.chars().skip(i);
            while !is_uniq(&window) {
                window.push_back(rest.next().unwrap());
                window.pop_front();
                i += 1;
            }
            i
        })
        .sum();
    println!("  - Part one: {result}");
}

fn is_uniq(v: &VecDeque<char>) -> bool {
    BTreeSet::<&char>::from_iter(v.iter()).len() == v.len()
}
