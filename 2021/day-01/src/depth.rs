use std::fs::File;
use std::io::BufReader;
use std::io::BufRead;

pub fn get_depths () -> Vec<u16> {
  let inputs = File::open("./input").expect("File failed to open");
  let input_reader = BufReader::new(inputs);
  // will implicitly return last value if there is no `;`.
  input_reader
    .lines()
    .map(|l| l.unwrap().parse::<u16>().unwrap())
    .collect()
}

// Was first written as `depths: &Vec<u16>` but linter suggested alternative.
pub fn count_increases (depths: &[u16]) -> u32 {
  let mut increases = 0;
  let mut last: u16 = u16::MAX;  
  for d in depths {
    if *d > last {
      increases += 1
    }
    last = *d;
  }
  increases
}

pub fn count_sliding_window_increases (depths: &[u16]) -> u32 {
  let d_size = depths.len();
  let mut increases = 0;
  let radius = 1;
  let mut i = radius;
  let mut last_window_sum = u16::MAX;
  while i + radius < d_size {
    let window_sum = depths[i - radius] + depths[i] + depths[i + radius];
    if window_sum > last_window_sum {
      increases += 1;
    }
    last_window_sum = window_sum;
    i += 1;
  }
  increases
}
