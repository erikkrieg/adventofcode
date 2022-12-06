use std::fs;

fn main() {
    let inputs = get_inputs();
    println!("Part 1: {}", part_1(&inputs));
    println!("Part 2: {}", part_2(&inputs));
}

fn get_inputs() -> String {
    fs::read_to_string("./input").expect("input not found")
}

fn part_1(inputs: &str) -> i32 {
    // Could get size of bin dynamically or hardcode
    // let bin_size = inputs.find('\n').unwrap();
    let mut ones: [u16; 12] = [0; 12];
    let mut zeros: [u16; 12] = [0; 12];

    inputs.lines().for_each(|input| {
        input.chars().enumerate().for_each(|(i, c)| {
            if c == '1' {
                ones[i] += 1;
            } else {
                zeros[i] += 1;
            }
        })
    });

    let mut gamma = 0;
    let mut epsilon = 0;

    for i in 0..12_usize {
        if ones[i] > zeros[i] {
            gamma = gamma * 2 + 1; // mores 1s
            epsilon *= 2; // less 0s
        } else {
            gamma *= 2; // more 0s
            epsilon = epsilon * 2 + 1; // less 1s
        }
    }

    gamma * epsilon
}

fn part_2(_inputs: &str) {}
