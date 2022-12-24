use std::{
    char,
    collections::HashSet,
    fs::File,
    io::{self, BufRead, BufReader},
};

pub fn solve() -> io::Result<()> {
    println!("- Day 24:");
    let input = File::open("input/day-24.txt")?;
    //let input = File::open("input/sample-24.txt")?;
    let valley = Valley::from(input);
    part_one(valley.clone());
    Ok(())
}

fn part_one(valley: Valley) {
    let mut valley = valley;
    let mut minutes = 0;
    loop {
        let mut reached_exit = false;
        minutes += 1;
        valley.elves = valley.elves.iter().fold(HashSet::new(), |mut elves, elf| {
            [
                elf,
                &elf.shift(1, 0),
                &elf.shift(-1, 0),
                &elf.shift(0, 1),
                &elf.shift(0, -1),
            ]
            .iter()
            .for_each(|&new_elf_point| {
                if valley.walls.get(new_elf_point).is_none() {
                    elves.insert(*new_elf_point);
                    // TODO: if the exit is reached, stop checking movements for other elves
                    reached_exit = reached_exit || valley.is_exit(new_elf_point);
                }
            });
            elves
        });
        if reached_exit {
            break;
        }
        valley.blizzards.iter_mut().for_each(|b| {
            let (x, y) = match b.direction {
                '>' => (1, 0),
                '<' => (-1, 0),
                '^' => (0, -1),
                'v' => (0, 1),
                _ => panic!("Invalid direction: {}", b.direction),
            };
            b.position = b.position.shift(x, y);
            if b.position.x == 0 {
                b.position.x = valley.max_x - 1;
            } else if b.position.x == valley.max_x {
                b.position.x = 1;
            } else if b.position.y == 0 {
                b.position.y = valley.max_y - 1;
            } else if b.position.y == valley.max_y {
                b.position.y = 1;
            }
            valley.elves.remove(&b.position);
        });
    }
    println!("  - Part 1: {minutes}");
}

#[derive(Debug, Eq, Hash, PartialEq, Clone, Copy)]
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
    max_x: usize,
    max_y: usize,
}

impl Valley {
    fn new() -> Self {
        let mut valley = Valley::default();
        valley.elves.insert(Point { x: 1, y: 0 });
        valley
    }

    fn from(file: File) -> Self {
        let reader = BufReader::new(file);
        reader
            .lines()
            .filter_map(|l| l.ok())
            .enumerate()
            .fold(Self::new(), |mut valley, (y, l)| {
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
            })
    }

    fn is_exit(&self, point: &Point) -> bool {
        point.x == self.max_x - 1 && point.y == self.max_y
    }
}
