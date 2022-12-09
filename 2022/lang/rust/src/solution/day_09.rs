use std::fs::File;
use std::io::{self, prelude::*, BufReader};

pub fn solve() -> io::Result<()> {
    println!("- Day 03");

    let file = File::open("input/day-09.txt")?;
    //part_one(BufReader::new(file));
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

type Point = (i32, i32);
type PointTrails = Vec<Vec<Point>>;

fn in_contact(a: Point, b: Point) -> bool {
    let max_distance = a.0.abs_diff(b.0).max(a.1.abs_diff(b.1));
    max_distance <= 1
}

fn tail_trail_size(reader: BufReader<File>, knots: usize) -> usize {
    let segments: PointTrails = vec![0; knots].iter().map(|_| vec![(0, 0)]).collect();
    let mut segments: PointTrails =
        reader
            .lines()
            .filter_map(|l| l.ok())
            .fold(segments, |mut segments, cmd| {
                let cmd: Vec<&str> = cmd.split_whitespace().collect();
                let steps = cmd[1].parse::<i32>().unwrap();
                for _ in 0..steps {
                    let prev_head = *segments[0].last().unwrap();
                    let next_head = match cmd[0] {
                        "U" => (prev_head.0, prev_head.1 - 1),
                        "R" => (prev_head.0 + 1, prev_head.1),
                        "D" => (prev_head.0, prev_head.1 + 1),
                        "L" => (prev_head.0 - 1, prev_head.1),
                        _ => panic!("Unknown direction: {}", cmd[0]),
                    };
                    segments[0].push(next_head);
                    let mut leader = next_head;
                    segments.iter_mut().skip(1).for_each(|s| {
                        let last = s.last().unwrap();
                        if !in_contact(leader, *last) {
                            let next = (
                                last.0 + (leader.0 - last.0).signum(),
                                last.1 + (leader.1 - last.1).signum(),
                            );
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
