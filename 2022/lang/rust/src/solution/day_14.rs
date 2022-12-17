use std::{
    collections::{BTreeSet, HashMap},
    fs::File,
    io::{self, BufRead, BufReader},
};

pub fn solve() -> io::Result<()> {
    println!("- Day 14:");
    let input = File::open("input/day-14.txt")?;
    //let input = File::open("input/sample-14.txt")?;
    part_one(BufReader::new(input));
    Ok(())
}

type Point = (usize, usize);

#[derive(Debug)]
enum Cell {
    Sand,
    Rock,
}

#[derive(Debug)]
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
            cursor = source;
        }
    }
    cells.len() - rock_count
}

fn part_one(reader: BufReader<File>) {
    let pouring_point = (500, 0);
    let mut bounds = Bounds {
        max_x: 500,
        min_x: 500,
        min_y: 0,
        max_y: 0,
    };
    let cell_map =
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
    let sand_count = simulate(pouring_point, bounds, cell_map);
    println!("  - Part 1: {sand_count}");
}
