use std::{
    collections::{HashSet, VecDeque},
    fs::File,
    io::{self, BufRead, BufReader},
};

#[derive(Debug, Eq, PartialEq, Hash, Clone)]
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

fn measure_exterior(cubes: &HashSet<Cube>) -> usize {
    let origin = cubes.iter().min_by_key(|c| c.z).unwrap().shift(0, 0, -1);
    let mut checked = HashSet::new();
    let mut perimeter = HashSet::new();
    let mut queue = VecDeque::from([origin]);
    while let Some(queued_cube) = queue.pop_front() {
        if checked.contains(&queued_cube) {
            continue;
        }
        checked.insert(queued_cube.clone());
        // Seach adjacent cubes of adjacent cubes to better creep along the
        // more extreme edges.
        let adjacent_cubes = queued_cube
            .adjacent()
            .iter()
            .filter(|c| !cubes.contains(c))
            .fold(Vec::new(), |mut v, c| {
                let mut adj = c
                    .adjacent()
                    .iter()
                    .filter(|c| !cubes.contains(c))
                    .cloned()
                    .collect();
                v.append(&mut adj);
                v.push(c.clone());
                v
            });
        let mut queue_next: VecDeque<Cube> = adjacent_cubes
            .iter()
            .filter(|c| c.adjacent().iter().any(|a| cubes.contains(a)))
            .cloned()
            .collect();
        if !queue_next.is_empty() {
            perimeter.insert(queued_cube.clone());
            queue.append(&mut queue_next);
        }
    }
    perimeter.iter().fold(0, |area, c| {
        area + c.adjacent().iter().filter(|c| cubes.contains(c)).count()
    })
}

pub fn solve() -> io::Result<()> {
    println!("- Day 18:");
    let input = File::open("input/day-18.txt")?;
    let cubes = parse_cubes(BufReader::new(input));
    part_one(&cubes);
    part_two(&cubes);
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

fn part_two(cubes: &HashSet<Cube>) {
    let perimeter = measure_exterior(cubes);
    println!("  - Part 2: {perimeter}");
}
