use std::{
    collections::HashSet,
    fs::File,
    io::{self, BufRead, BufReader},
};

type Point = (i32, i32);
type Intersections = Vec<(Point, Point)>;

fn parse_sensors(reader: BufReader<File>) -> Vec<(Point, Point, i32)> {
    reader
        .lines()
        .filter_map(|line| line.ok())
        .map(|line| {
            let line: String = line
                .matches(|c| char::is_numeric(c) || c == '=' || c == '-')
                .collect();
            let pos: Vec<i32> = line.split('=').filter_map(|i| i.parse().ok()).collect();
            let sensor = (pos[0], pos[1]);
            let beacon = (pos[2], pos[3]);
            let radius = sensor.0.abs_diff(beacon.0) + sensor.1.abs_diff(beacon.1);
            (sensor, beacon, radius as i32)
        })
        .collect()
}

fn find_intersections(sensors: &[(Point, Point, i32)], intersect_y: i32) -> Intersections {
    sensors
        .iter()
        .filter_map(|(p, _, r)| {
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

fn join_overlapping(mut intersects: Intersections, limit: Option<(i32, i32)>) -> Intersections {
    let no_overlap = |a: Point, b: Point| a.1 < b.0 || a.0 > b.1;
    let mut combined = Vec::new();
    while let Some(mut i) = intersects.pop() {
        if let Some(limit) = limit {
            i.0 .0 = i.0 .0.clamp(limit.0, limit.1);
            i.1 .0 = i.1 .0.clamp(limit.0, limit.1);
        }
        if combined.is_empty() {
            combined.push(i);
            continue;
        }
        let mut overlaps: Intersections = combined
            .drain_filter(|(s, e)| !no_overlap((i.0 .0, i.1 .0), (s.0, e.0)))
            .collect();
        overlaps.push(i);
        let start = overlaps.iter().min_by(|a, b| a.0 .0.cmp(&b.0 .0)).unwrap();
        let end = overlaps.iter().max_by(|a, b| a.1 .0.cmp(&b.1 .0)).unwrap();
        combined.push((start.0, end.1));
    }
    combined
}

fn sum_intersections(joined_intersects: &Intersections) -> usize {
    joined_intersects
        .iter()
        .map(|((sx, _), (ex, _))| (ex - sx + 1) as usize)
        .sum()
}

pub fn solve() -> io::Result<()> {
    println!("- Day 15:");
    let input = File::open("input/day-15.txt")?;
    let sensors = parse_sensors(BufReader::new(input));
    part_one(&sensors);
    part_two(&sensors);
    Ok(())
}

fn part_one(sensors: &[(Point, Point, i32)]) {
    let intersect_y = 2_000_000;
    let intersects = find_intersections(sensors, intersect_y);
    let beacons: HashSet<Point> = sensors
        .iter()
        .filter_map(|(_, b, _)| match b.1.cmp(&intersect_y) {
            std::cmp::Ordering::Equal => Some(*b),
            _ => None,
        })
        .fold(HashSet::new(), |mut set, b| {
            set.insert(b);
            set
        });
    let joined_intersects = join_overlapping(intersects, None);
    let sum = sum_intersections(&joined_intersects);
    let points_without_beacons = sum - beacons.len();
    println!("  - Part 1: {points_without_beacons}");
}

fn part_two(sensors: &[(Point, Point, i32)]) {
    let limit = (0, 4_000_000);
    for y in limit.0..=limit.1 {
        let intersects = find_intersections(sensors, y);
        let joined_intersects = join_overlapping(intersects, Some(limit));
        let sum = sum_intersections(&joined_intersects);
        if (sum as i32) < limit.1 + 1 {
            let mut x_ranges: Vec<Point> =
                joined_intersects.iter().map(|(s, e)| (s.0, e.0)).collect();
            let x = if x_ranges.len() == 1 {
                match x_ranges[0].0.cmp(&limit.0) {
                    std::cmp::Ordering::Greater => limit.0,
                    _ => limit.1,
                }
            } else {
                x_ranges.sort_by(|a, b| a.0.cmp(&b.0));
                let (_, ex) = x_ranges[0];
                ex + 1
            };
            let beacon_tuning_frequency = x as u64 * limit.1 as u64 + y as u64;
            println!("  - Part 2: {beacon_tuning_frequency}");
            break;
        }
    }
}
