use std::{
    fs::File,
    io::{self, BufRead, BufReader},
};

pub fn solve() -> io::Result<()> {
    println!("- Day 10:");
    let input = File::open("input/day-10.txt")?;
    let reader = BufReader::new(input);
    part_one(reader)?;
    Ok(())
}

fn part_one(reader: BufReader<File>) -> io::Result<()> {
    let (cycles, _): (Vec<i32>, i32) =
        reader
            .lines()
            .filter_map(|l| l.ok())
            .fold((vec![1], 0), |(mut cycles, add), l| {
                cycles.push(*cycles.last().unwrap() + add);
                let add = match l.split_whitespace().nth(1) {
                    Some(addend) => {
                        // addx uses an extra cycle
                        cycles.push(*cycles.last().unwrap());
                        addend.parse().expect("addx command only accepts integers")
                    }
                    None => 0,
                };
                (cycles, add)
            });
    let signal_strength_sum: i32 = cycles
        .iter()
        .enumerate()
        .skip(20)
        .filter_map(|(i, v)| {
            if (i as i32 - 20) % 40 == 0 {
                println!("{i} * {v} = {}", i as i32 * v);
            }
            match (i as i32 - 20) % 40 {
                0 => Some(i as i32 * v),
                _ => None,
            }
        })
        .sum();
    println!("{signal_strength_sum}");
    Ok(())
}
