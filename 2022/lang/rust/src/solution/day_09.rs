use std::fs::File;
use std::io::{self, prelude::*, BufReader};

use advent::grid::point::Point;

pub fn solve() -> io::Result<()> {
    println!("- Day 09");
    let file = File::open("input/day-09.txt")?;
    part_one(BufReader::new(file));
    let file = File::open("input/day-09.txt")?;
    part_two(BufReader::new(file));
    Ok(())
}

fn part_one(reader: BufReader<File>) {
    println!("  - Part 1: {}", tail_trail_size(reader, 2));
}

fn part_two(reader: BufReader<File>) {
    println!("  - Part 2: {}", tail_trail_size(reader, 10));
}

type PointTrails = Vec<Vec<Point>>;

fn tail_trail_size(reader: BufReader<File>, knots: usize) -> usize {
    let segments: PointTrails = vec![0; knots].iter().map(|_| vec![Point::new()]).collect();
    let mut segments: PointTrails =
        reader
            .lines()
            .filter_map(|l| l.ok())
            .fold(segments, |mut segments, cmd| {
                let cmd: Vec<&str> = cmd.split_whitespace().collect();
                let steps = cmd[1].parse::<i32>().unwrap();
                for _ in 0..steps {
                    let prev_head = *segments[0].last().unwrap();
                    let next_head = prev_head
                        + match cmd[0] {
                            "U" => (0, -1),
                            "R" => (1, 0),
                            "D" => (0, 1),
                            "L" => (-1, 0),
                            _ => panic!("Unknown direction: {}", cmd[0]),
                        };
                    segments[0].push(next_head);
                    let mut leader = next_head;
                    segments.iter_mut().skip(1).for_each(|s| {
                        let last = s.last().unwrap();
                        if !leader.is_adjacent(last) {
                            let next = *last
                                + ((leader.x - last.x).signum(), (leader.y - last.y).signum());
                            s.push(next);
                        }
                        leader = *s.last().unwrap();
                    });
                }
                segments
            });
    let tail_trail = segments.last_mut().unwrap();
    tail_trail.sort();
    tail_trail.dedup();
    tail_trail.len()
}
