use std::{
    char,
    collections::HashSet,
    fs::File,
    io::{self, BufRead, BufReader},
};

pub fn solve() -> io::Result<()> {
    println!("- Day 24:");
    let input = File::open("input/day-24.txt")?;
    let valley = Valley::from(input);
    part_one(valley.clone());
    part_two(valley);
    Ok(())
}

fn part_one(valley: Valley) {
    let mut valley = valley;
    let minutes = valley.time_to(&valley.exit.clone());
    println!("  - Part 1: {minutes}");
}

fn part_two(valley: Valley) {
    let mut valley = valley;
    let mut minutes = 0;
    minutes += valley.time_to(&valley.exit.clone());
    valley.elves = HashSet::from([valley.exit]);
    minutes += valley.time_to(&valley.start.clone()) - 1;
    valley.elves = HashSet::from([valley.start]);
    minutes += valley.time_to(&valley.exit.clone()) - 1;
    println!("  - Part 2: {minutes}");
}

#[derive(Debug, Default, Eq, Hash, PartialEq, Clone, Copy)]
struct Point {
    x: usize,
    y: usize,
}

impl Point {
    fn shift(&self, x: i32, y: i32) -> Self {
        let x = ((self.x as i32) + x).max(0) as usize;
        let y = ((self.y as i32) + y).max(0) as usize;
        Point { x, y }
    }
}

#[derive(Debug, Clone)]
struct Blizzard {
    position: Point,
    direction: char,
}

#[derive(Default, Debug, Clone)]
struct Valley {
    walls: HashSet<Point>,
    elves: HashSet<Point>,
    blizzards: Vec<Blizzard>,
    start: Point,
    exit: Point,
    max_x: usize,
    max_y: usize,
}

impl Valley {
    fn new() -> Self {
        let mut valley = Valley::default();
        valley.start.x = 1;
        valley.elves.insert(valley.start);
        valley
    }

    fn from(file: File) -> Self {
        let reader = BufReader::new(file);
        let mut valley = reader.lines().filter_map(|l| l.ok()).enumerate().fold(
            Self::new(),
            |mut valley, (y, l)| {
                valley.max_y = y;
                l.chars().enumerate().fold(valley, |mut v, (x, c)| {
                    v.max_x = v.max_x.max(x);
                    match c {
                        '#' => {
                            v.walls.insert(Point { x, y });
                            v
                        }
                        '.' => v,
                        _ => {
                            v.blizzards.push(Blizzard {
                                position: Point { x, y },
                                direction: c,
                            });
                            v
                        }
                    }
                })
            },
        );
        valley.exit = Point {
            x: valley.max_x - 1,
            y: valley.max_y,
        };
        valley
    }

    fn time_to(&mut self, target_point: &Point) -> usize {
        let mut minutes = 0;
        loop {
            let mut reached_point = false;
            minutes += 1;
            self.elves = self.elves.iter().fold(HashSet::new(), |mut elves, elf| {
                if reached_point {
                    return elves;
                }
                reached_point = [
                    elf,
                    &elf.shift(1, 0),
                    &elf.shift(-1, 0),
                    &elf.shift(0, 1),
                    &elf.shift(0, -1),
                ]
                .iter()
                .any(|&new_elf_point| {
                    let mut at_end = false;
                    if self.walls.get(new_elf_point).is_none() {
                        at_end = new_elf_point == target_point;
                        if at_end {
                            elves.clear();
                        }
                        elves.insert(*new_elf_point);
                    }
                    at_end
                });
                elves
            });
            if reached_point {
                break;
            }
            self.blizzards.iter_mut().for_each(|b| {
                let (x, y) = match b.direction {
                    '>' => (1, 0),
                    '<' => (-1, 0),
                    '^' => (0, -1),
                    'v' => (0, 1),
                    _ => panic!("Invalid direction: {}", b.direction),
                };
                b.position = b.position.shift(x, y);
                if b.position.x == 0 {
                    b.position.x = self.max_x - 1;
                } else if b.position.x == self.max_x {
                    b.position.x = 1;
                } else if b.position.y == 0 {
                    b.position.y = self.max_y - 1;
                } else if b.position.y == self.max_y {
                    b.position.y = 1;
                }
                self.elves.remove(&b.position);
            });
        }
        minutes
    }
}
