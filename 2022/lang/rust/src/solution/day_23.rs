use std::{
    collections::HashMap,
    fs::File,
    io::{self, BufRead, BufReader},
    ops::{Add, Sub},
};

#[derive(Debug, Eq, PartialEq, Hash, Clone, Copy)]
struct Point {
    x: i32,
    y: i32,
}

impl Point {
    fn abs(self) -> Self {
        Point {
            x: self.x.abs(),
            y: self.y.abs(),
        }
    }
}

impl Add for Point {
    type Output = Self;
    fn add(self, other: Self) -> Self {
        Point {
            x: self.x + other.x,
            y: self.y + other.y,
        }
    }
}

impl Sub for Point {
    type Output = Self;
    fn sub(self, other: Self) -> Self {
        Point {
            x: self.x - other.x,
            y: self.y - other.y,
        }
    }
}

#[derive(Debug, Clone)]
enum Cell {
    Occupied,
    Targeted(usize),
}

#[derive(Debug, Clone)]
struct Elf {
    pos: Point,
    proposed_pos: Option<Point>,
    next_dir: usize,
}

impl Elf {
    fn alone(&self, cells: &HashMap<Point, Cell>) -> bool {
        (-1..=1)
            .map(|i| {
                [
                    self.pos + Point { x: i, y: -1 },
                    self.pos + Point { x: i, y: 0 },
                    self.pos + Point { x: i, y: 1 },
                ]
                .iter()
                .filter_map(|p| cells.get(p))
                .filter(|c| matches!(c, Cell::Occupied))
                .count()
            })
            .sum::<usize>()
            == 1
    }

    fn propose_move(&mut self, cells: &mut HashMap<Point, Cell>) {
        let directions: [Point; 4] = [
            Point { x: 0, y: -1 },
            Point { x: 0, y: 1 },
            Point { x: -1, y: 0 },
            Point { x: 1, y: 0 },
        ];
        if self.alone(cells) {
            self.proposed_pos = None;
        } else if let Some((proposed_pos, _dir)) = directions
            .iter()
            .cycle()
            .skip(self.next_dir)
            .take(directions.len())
            .map(|&d| (self.pos + d, d))
            .find(|&(p, d)| {
                let adj = Point { x: 1, y: 1 } - d.abs();
                ![p, p + adj, p - adj]
                    .iter()
                    .filter_map(|p| cells.get(p))
                    .any(|c| matches!(c, Cell::Occupied))
            })
        {
            self.proposed_pos = Some(proposed_pos);
        } else {
            self.proposed_pos = None;
        }
        if let Some(pos) = self.proposed_pos {
            cells
                .entry(pos)
                .and_modify(|c| match c {
                    Cell::Targeted(i) => *i += 1,
                    Cell::Occupied => panic!("Can't propose moving into an occupied cell!"),
                })
                .or_insert(Cell::Targeted(1));
        }
        self.next_dir = (self.next_dir + 1) % directions.len();
    }
    fn try_move(&mut self, cells: &mut HashMap<Point, Cell>) {
        if let Some(pos) = self.proposed_pos {
            if let Some(Cell::Targeted(i)) = cells.get(&pos) {
                if i > &1 {
                    // between rounds need to remove these entries from cells
                    // else they will build up
                    return;
                }
            }
            cells.remove(&self.pos);
            cells.insert(pos, Cell::Occupied);
            self.proposed_pos = None;
            self.pos = pos;
        }
    }
}

fn parse_input(input: BufReader<File>) -> (Vec<Elf>, HashMap<Point, Cell>) {
    let elves = Vec::new();
    let cells = HashMap::new();
    input
        .lines()
        .filter_map(|line| line.ok())
        .enumerate()
        .fold((elves, cells), |acc, (y, line)| {
            let y = y as i32;
            line.chars().enumerate().filter(|(_, c)| c == &'#').fold(
                acc,
                |(mut elves, mut cells), (x, _c)| {
                    let point = Point { x: x as i32, y };
                    elves.push(Elf {
                        pos: point,
                        proposed_pos: None,
                        next_dir: 0,
                    });
                    cells.insert(point, Cell::Occupied);
                    (elves, cells)
                },
            )
        })
}

pub fn solve() -> io::Result<()> {
    println!("- Day 23:");
    let input = File::open("input/day-23.txt")?;
    let (elves, cells) = parse_input(BufReader::new(input));
    part_one(elves.clone(), cells.clone());
    part_two(elves, cells);
    Ok(())
}

fn part_one(elves: Vec<Elf>, cells: HashMap<Point, Cell>) {
    let (elves, _cells) = (0..10).fold((elves, cells), |(mut elves, mut cells), _i| {
        elves.iter_mut().for_each(|e| e.propose_move(&mut cells));
        elves.iter_mut().for_each(|e| e.try_move(&mut cells));
        cells.retain(|_, cell| matches!(cell, Cell::Occupied));
        (elves, cells)
    });
    let max_x = elves.iter().map(|e| e.pos.x).max().unwrap();
    let min_x = elves.iter().map(|e| e.pos.x).min().unwrap();
    let max_y = elves.iter().map(|e| e.pos.y).max().unwrap();
    let min_y = elves.iter().map(|e| e.pos.y).min().unwrap();
    let open_cell_count = (max_x + 1 - min_x) * (max_y + 1 - min_y) - elves.len() as i32;
    println!("  - Part 1: {open_cell_count}");
}

fn part_two(elves: Vec<Elf>, cells: HashMap<Point, Cell>) {
    let mut elves = elves;
    let mut cells = cells;
    let mut rounds = 0;
    loop {
        elves.iter_mut().for_each(|e| e.propose_move(&mut cells));
        if !elves.iter().any(|e| e.proposed_pos.is_some()) {
            break;
        }
        elves.iter_mut().for_each(|e| e.try_move(&mut cells));
        cells.retain(|_, cell| matches!(cell, Cell::Occupied));
        rounds += 1;
    }
    println!("  - Part 2: {}", rounds + 1);
}
