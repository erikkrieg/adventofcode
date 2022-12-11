use core::ops::RangeInclusive;
use std::{
    fs::File,
    io::{self, BufRead, BufReader},
};

enum Overlap {
    Full,
    Partial,
    None,
}

trait Overlaps {
    fn overlaps(&self, other: &Self) -> Overlap;
}

impl Overlaps for RangeInclusive<i32> {
    fn overlaps(&self, other: &Self) -> Overlap {
        let overlap_size = self.end().min(other.end()) - self.start().max(other.start());
        let min_range_size = (self.end() - self.start()).min(other.end() - other.start());
        if overlap_size == min_range_size {
            Overlap::Full
        } else if overlap_size > 0 {
            Overlap::Partial
        } else {
            Overlap::None
        }
    }
}

pub fn solve() -> io::Result<()> {
    println!("- Day 04:");
    let input = File::open("input/day-04.txt")?;
    part_one(BufReader::new(input));
    Ok(())
}

fn part_one(reader: BufReader<File>) {
    let contained_assignments = reader
        .lines()
        .filter_map(|l| l.ok())
        .filter(|line| {
            let ranges = line
                .split(',')
                .map(|s| {
                    let r = s
                        .split('-')
                        .map(|i| i.parse::<i32>().expect("Assignments must be ints"))
                        .collect::<Vec<i32>>();
                    r[0]..=r[1]
                })
                .collect::<Vec<RangeInclusive<i32>>>();
            matches!(ranges[0].overlaps(&ranges[1]), Overlap::Full)
        })
        .count();
    println!("  - Part 1: {contained_assignments}");
}
