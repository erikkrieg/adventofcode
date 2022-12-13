use std::{
    fs::File,
    io::{self, BufRead, BufReader},
};

pub fn solve() -> io::Result<()> {
    println!("- Day 11:");
    let input = File::open("input/day-11.txt")?;
    // todo: part one solution isn't available until this is refactored
    part_two(BufReader::new(input));
    Ok(())
}

#[derive(Debug, Clone)]
struct Monkey {
    items: Vec<i64>,
    inspect_op: String,
    inspect_val: Option<i64>,
    test_div: i64,
    targets: (usize, usize),
    inspected: usize,
}

impl Monkey {
    fn get_inspect_val(&self, i: i64) -> i64 {
        match self.inspect_val {
            Some(v) => v,
            None => i,
        }
    }
    fn inspect(&self, i: i64) -> i64 {
        let val = self.get_inspect_val(i);
        match self.inspect_op.as_str() {
            "*" => i * val,
            "+" => i + val,
            _ => panic!("Not an available operation: {}", self.inspect_op),
        }
    }
}

fn part_two(reader: BufReader<File>) {
    println!("  - Part 1: ");
    let mut monkeys: Vec<Monkey> = reader
        .lines()
        .filter_map(|l| l.ok())
        .fold(
            (None, Vec::<Monkey>::new()),
            |(cur_monkey, mut monkeys), line| {
                let mut m = cur_monkey.unwrap_or(Monkey {
                    items: Vec::new(),
                    inspect_op: "+".to_string(),
                    inspect_val: None,
                    test_div: 0.into(),
                    targets: (0, 0),
                    inspected: 0,
                });
                if line.is_empty() {
                    monkeys.push(m);
                    return (None, monkeys);
                }
                let s: Vec<&str> = line.split(':').collect();
                match s[0].trim() {
                    "Starting items" => {
                        m.items = s[1]
                            .split(", ")
                            .filter_map(|i| i.trim().parse::<i64>().ok())
                            .collect();
                    }
                    "Operation" => {
                        let op: Vec<&str> = s[1].split_whitespace().skip(3).collect();
                        let val = op[1].parse::<i64>().ok();
                        m.inspect_op = op[0].to_string();
                        m.inspect_val = val;
                    }
                    "Test" => {
                        m.test_div = s[1]
                            .split_whitespace()
                            .last()
                            .unwrap()
                            .parse()
                            .expect("Last string in test should be parseable integer");
                    }
                    "If true" => {
                        m.targets.0 = s[1]
                            .split_whitespace()
                            .last()
                            .unwrap()
                            .parse()
                            .expect("Last string in target should be parseable integer");
                    }
                    "If false" => {
                        m.targets.1 = s[1]
                            .split_whitespace()
                            .last()
                            .unwrap()
                            .parse()
                            .expect("Last string in target should be parseable integer");
                    }
                    _ => (),
                }
                (Some(m), monkeys)
            },
        )
        .1;

    // todo: calculate lcm dynamically
    // Currently this only solves part 2 because the div by 3 was removed.
    // I'd like for this to solve both parts.
    let lcm = 9_699_690;
    let mut monkey_biz: Vec<usize> = (0..(monkeys.len() * 10_000))
        .fold(&mut monkeys, |monkeys, i| {
            let i = i % monkeys.len();
            if monkeys[i].items.is_empty() {
                return monkeys;
            }
            let mut m = monkeys[i].clone();
            let mut items: Vec<i64> = m.items.to_vec();
            m.inspected += items.len();
            items.drain(..).for_each(|item| {
                let worry = m.inspect(item);
                let target = match worry % m.test_div == 0.into() {
                    true => m.targets.0,
                    false => m.targets.1,
                };
                monkeys[target].items.push(worry % lcm)
            });
            m.items.clear();
            monkeys[i] = m;
            monkeys
        })
        .iter()
        .map(|m| m.inspected)
        .collect();

    monkey_biz.sort();
    monkey_biz.reverse();
    println!("{:?}", &monkey_biz[0..=1]);
}
