use std::{
    collections::HashSet,
    fs::File,
    io::{self, BufRead, BufReader},
};

#[derive(Debug, Eq, PartialEq, Hash)]
struct Cube {
    x: i32,
    y: i32,
    z: i32,
}

impl Cube {
    fn shift(&self, x: i32, y: i32, z: i32) -> Self {
        Cube {
            x: self.x + x,
            y: self.y + y,
            z: self.z + z,
        }
    }

    fn adjacent(&self) -> [Cube; 6] {
        [
            self.shift(1, 0, 0),
            self.shift(-1, 0, 0),
            self.shift(0, 1, 0),
            self.shift(0, -1, 0),
            self.shift(0, 0, 1),
            self.shift(0, 0, -1),
        ]
    }
}

fn parse_cubes(reader: BufReader<File>) -> HashSet<Cube> {
    reader
        .lines()
        .filter_map(|line| line.ok())
        .fold(HashSet::new(), |mut cubes, line| {
            let [x, y, z] = line
                .split(',')
                .filter_map(|s| s.parse::<i32>().ok())
                .collect::<Vec<i32>>()[..] else { panic!("Expected to parse 3 ints from line") };
            cubes.insert(Cube { x, y, z });
            cubes
        })
}

pub fn solve() -> io::Result<()> {
    println!("- Day 18:");
    let input = File::open("input/day-18.txt")?;
    let cubes = parse_cubes(BufReader::new(input));
    part_one(&cubes);
    Ok(())
}

fn part_one(cubes: &HashSet<Cube>) {
    let lava_surface_area: usize = cubes.iter().fold(0, |area, cube| {
        area + cube
            .adjacent()
            .iter()
            .filter(|c| !cubes.contains(c))
            .count()
    });
    println!("  - Part 1: {lava_surface_area}");
}
