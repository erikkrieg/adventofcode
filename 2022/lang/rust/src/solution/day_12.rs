use std::{
    collections::{HashMap, VecDeque},
    fs::File,
    io::{self, BufRead, BufReader},
};

pub fn solve() -> io::Result<()> {
    println!("- Day 12:");
    let input = File::open("input/day-12.txt")?;

    // todo: get parser to return start and end points
    let height_map = parse_height_map(BufReader::new(input));
    let starting_point = (0, 20);
    let end_point = (148, 20);
    let distances = find_distances(starting_point, end_point, &height_map);
    println!("{}", distances.get(&end_point).expect("should find"));
    Ok(())
}

type HeightMap = Vec<Vec<u8>>;
type Point = (usize, usize);

fn find_distances(start: Point, end: Point, height_map: &HeightMap) -> HashMap<Point, usize> {
    let max = (height_map[0].len() - 1, height_map.len() - 1);
    let mut distance_map = HashMap::new();
    let mut queue: VecDeque<Point> = VecDeque::new();
    distance_map.insert(start, 0);
    queue.push_back(start);
    while let Some(p) = queue.pop_front() {
        let distance = &distance_map
            .get(&p)
            .expect("Points popped from queue should have a distance already set")
            .clone();
        let height = &height_map[p.1][p.0];
        [
            (p.0, p.1 + 1),
            (p.0 + 1, p.1),
            (p.0.saturating_sub(1), p.1),
            (p.0, p.1.saturating_sub(1)),
        ]
        .iter()
        .filter(|(x, y)| x <= &max.0 && y <= &max.1)
        .filter(|point| height_map[point.1][point.0] <= height + 1)
        .for_each(|point| {
            let next_dist = *distance + 1;
            if let Some(d) = distance_map.get(point) {
                if d > &(next_dist) {
                    distance_map.insert(*point, next_dist);
                    queue.push_back(*point);
                }
            } else {
                distance_map.insert(*point, next_dist);
                queue.push_back(*point);
            }
        });
    }
    distance_map
}

fn parse_height_map(reader: BufReader<File>) -> Vec<Vec<u8>> {
    reader
        .lines()
        .filter_map(|line| line.ok())
        .map(|line| {
            line.chars()
                .map(|c| {
                    if c.is_lowercase() {
                        c as u8 % 97
                    } else {
                        match c {
                            'S' => 0,
                            _ => 25,
                        }
                    }
                })
                .collect()
        })
        .collect()
}
