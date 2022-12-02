use std::fs::File;
use std::io::{self, prelude::*, BufReader};

pub fn solve() -> io::Result<()> {
    println!("- Day 02");
    let score = calculate_score()?;
    println!("{score}");
    Ok(())
}

// TODO: update this to be flexible enough to solve part 1 and 2
fn calculate_score() -> Result<u32, io::Error> {
    let file = File::open("input/day-02.txt")?;
    let reader = BufReader::new(file);
    Ok(reader.lines().into_iter().fold(0, |acc, x| {
        let x = x.expect("Line should be string");
        let res: Vec<&str> = x.split_whitespace().collect();
        let mut score_change = 0;
        match res[1] {
            // lose
            "X" => {
                score_change += match res[0] {
                    "A" => 3,
                    "B" => 1,
                    "C" => 2,
                    _ => 0,
                };
            }
            // draw
            "Y" => {
                score_change += 3 + match res[0] {
                    "A" => 1,
                    "B" => 2,
                    "C" => 3,
                    _ => 0,
                };
            }
            // win
            "Z" => {
                score_change += 6 + match res[0] {
                    "A" => 2,
                    "B" => 3,
                    "C" => 1,
                    _ => 0,
                };
            }
            _ => panic!("Unknown move: {}", res[1]),
        };
        /*
        match res[1] {
            "X" => {
                score_change += 1;
                if res[0] == "A" {
                    score_change += 3;
                } else if res[0] == "C" {
                    score_change += 6;
                }
            }
            "Y" => {
                score_change += 2;
                if res[0] == "B" {
                    score_change += 3;
                } else if res[0] == "A" {
                    score_change += 6;
                }
            }
            "Z" => {
                score_change += 3;
                if res[0] == "C" {
                    score_change += 3;
                } else if res[0] == "B" {
                    score_change += 6;
                }
            }
            _ => panic!("Unknown move: {}", res[1]),
        };
        */
        acc + score_change
    }))
}
