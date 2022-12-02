use std::{env, io};

mod solution;

fn main() -> io::Result<()> {
    println!("Advent of Code 2022");
    let day = env::args()
        .nth(1)
        .unwrap_or_else(|| "0".to_string())
        .parse::<u8>()
        .unwrap_or(0);

    match day {
        1 => solution::day_01::solve()?,
        _ => panic!("Solution not found"),
    };

    Ok(())
}
