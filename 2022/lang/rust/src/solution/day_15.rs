use std::{
    cmp::Ordering,
    fs::File,
    io::{self, BufRead, BufReader},
};

type Point = (i32, i32);
type Intersections = Vec<(Point, Point)>;

fn parse_intersections(reader: BufReader<File>, intersect_y: i32) -> Intersections {
    reader
        .lines()
        .filter_map(|line| line.ok())
        .map(|line| {
            let line: String = line.matches(|c| char::is_numeric(c) || c == '=').collect();
            let pos: Vec<i32> = line.split('=').filter_map(|i| i.parse().ok()).collect();
            let sensor = (pos[0], pos[1]);
            let beacon = (pos[2], pos[3]);
            let radius = sensor.0.abs_diff(beacon.0) + sensor.1.abs_diff(beacon.1);
            (sensor, radius as i32)
        })
        .filter_map(|(p, r)| {
            let min_y = p.1 - r;
            let max_y = p.1 + r;
            if min_y <= intersect_y && max_y >= intersect_y {
                let diff = intersect_y.abs_diff(min_y).min(intersect_y.abs_diff(max_y)) as i32;
                Some(((p.0 - diff, intersect_y), (p.0 + diff, intersect_y)))
            } else {
                None
            }
        })
        .collect()
}

// Since the number of intersections was very small I did the sums manually,
// but it would certainly be nicer to do it programmatically.
fn sum_intersections(intersects: Intersections) -> usize {
    todo!("Iterate and add intersections to determine the total area in range of sensors")
}

pub fn solve() -> io::Result<()> {
    println!("- Day 15:");
    let input = File::open("input/day-15.txt")?;
    let mut intersects = parse_intersections(BufReader::new(input), 2_000_000);
    intersects.sort_by(|a, b| match a.0 .0.cmp(&b.0 .0) {
        Ordering::Equal => a.1 .0.cmp(&b.1 .0),
        Ordering::Greater => Ordering::Greater,
        Ordering::Less => Ordering::Less,
    });
    println!("  - Part 1:");
    intersects.iter().for_each(|i| println!("    - {i:?}"));
    Ok(())
}
