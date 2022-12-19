#![feature(array_chunks, drain_filter, iter_array_chunks, int_roundings)]

use std::{env, io};

mod solution;

fn main() -> io::Result<()> {
    println!("Advent of Code 2022");
    let day = env::args()
        .nth(1)
        .unwrap_or_else(|| "0".to_string())
        .parse::<u8>()
        .expect("Must provide a number for corresponding day");

    match day {
        1 => solution::day_01::solve()?,
        2 => solution::day_02::solve()?,
        3 => solution::day_03::solve()?,
        4 => solution::day_04::solve()?,
        5 => solution::day_05::solve()?,
        7 => solution::day_07::solve()?,
        9 => solution::day_09::solve()?,
        10 => solution::day_10::solve()?,
        11 => solution::day_11::solve()?,
        12 => solution::day_12::solve()?,
        13 => solution::day_13::solve()?,
        14 => solution::day_14::solve()?,
        15 => solution::day_15::solve()?,
        18 => solution::day_18::solve()?,
        _ => panic!("Solution not found"),
    };

    Ok(())
}
