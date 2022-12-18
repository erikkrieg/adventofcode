use std::{
    collections::{BTreeSet, HashMap},
    fs::File,
    io::{self, BufRead, BufReader},
};

type Point = (usize, usize);

#[derive(Debug, Clone)]
enum Cell {
    Sand,
    Rock,
}

#[derive(Debug, Clone)]
struct Bounds {
    max_x: usize,
    min_x: usize,
    max_y: usize,
    min_y: usize,
}

impl Bounds {
    fn contains(&self, (x, y): &Point) -> bool {
        x <= &self.max_x && x >= &self.min_x && y <= &self.max_y && y >= &self.min_y
    }
}

fn simulate(source: Point, bounds: Bounds, cells: HashMap<Point, Cell>) -> usize {
    let rock_count = cells.len();
    let mut cells = cells;
    let mut cursor = source;
    while bounds.contains(&cursor) {
        let next = (cursor.0, cursor.1 + 1);
        if !cells.contains_key(&next) {
            cursor = next;
            continue;
        }
        let next_left = (next.0 - 1, next.1);
        let next_right = (next.0 + 1, next.1);
        if !cells.contains_key(&next_left) {
            cursor = next_left;
        } else if !cells.contains_key(&next_right) {
            cursor = next_right;
        } else {
            cells.insert(cursor, Cell::Sand);
            if cursor == source {
                break;
            }
            cursor = source;
        }
    }
    cells.len() - rock_count
}

fn parse_cells(reader: BufReader<File>) -> (Bounds, HashMap<Point, Cell>) {
    let mut bounds = Bounds {
        max_x: 500,
        min_x: 500,
        min_y: 0,
        max_y: 0,
    };
    let cells =
        reader
            .lines()
            .filter_map(|l| l.ok())
            .fold(HashMap::<Point, Cell>::new(), |m, l| {
                let mut prev_point: Option<Point> = None;
                l.split(" -> ")
                    .map(|s| {
                        let [x, y] = s
                            .split(',')
                            .filter_map(|p| p.parse().ok())
                            .collect::<Vec<usize>>()[..] else { panic!("Point should have x and y") };
                        bounds.min_x = bounds.min_x.min(x);
                        bounds.max_x = bounds.max_x.max(x);
                        bounds.max_y = bounds.max_y.max(y);
                        (x, y)
                    })
                    .fold(m, |mut m, p| {
                        if let Some(pp) = prev_point {
                            let px = BTreeSet::from([p.0, pp.0]);
                            let py = BTreeSet::from([p.1, pp.1]);
                            for x in *px.first().unwrap()..=*px.last().unwrap() {
                                for y in *py.first().unwrap()..=*py.last().unwrap() {
                                    m.entry((x, y)).or_insert(Cell::Rock);
                                }
                            }
                        }
                        prev_point = Some(p);
                        m
                    })
            });
    (bounds, cells)
}

pub fn solve() -> io::Result<()> {
    println!("- Day 14:");
    let input = File::open("input/day-14.txt")?;
    let (bounds, cells) = parse_cells(BufReader::new(input));
    part_one(bounds.clone(), cells.clone());
    part_two(bounds, cells);
    Ok(())
}

fn part_one(bounds: Bounds, cells: HashMap<Point, Cell>) {
    let pouring_point = (500, 0);
    let sand_count = simulate(pouring_point, bounds, cells);
    println!("  - Part 1: {sand_count}");
}

fn part_two(bounds: Bounds, cells: HashMap<Point, Cell>) {
    let mut cells = cells;
    let mut bounds = bounds;
    let pouring_point = (500, 0);
    bounds.min_x = 250;
    bounds.max_x = 750;
    bounds.max_y += 2;
    for x in bounds.min_x..=bounds.max_x {
        cells.insert((x, bounds.max_y), Cell::Rock);
    }
    let sand_count = simulate(pouring_point, bounds, cells);
    println!("  - Part 2: {sand_count}");
}
