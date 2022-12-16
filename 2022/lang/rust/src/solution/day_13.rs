use std::{
    cmp::Ordering,
    fs::File,
    io::{self, BufRead, BufReader},
};

pub fn solve() -> io::Result<()> {
    let input = File::open("input/day-13.txt")?;
    let mut signals = parse_signals(BufReader::new(input));
    part_one(&signals);
    part_two(&mut signals);
    Ok(())
}

fn parse_signals(reader: BufReader<File>) -> Vec<Token> {
    reader
        .lines()
        .filter_map(|l| l.ok())
        .filter(|line| !line.is_empty())
        .map(|l| as_token(&l))
        .collect()
}

fn part_one(signals: &[Token]) {
    let sum: usize = signals
        .array_chunks::<2>()
        .enumerate()
        .filter_map(|(i, [a, b])| match a.compare(b) {
            Ordering::Greater => None,
            _ => Some(i + 1),
        })
        .sum();
    println!("  - Part 1: {sum}");
}

fn part_two(signals: &mut Vec<Token>) {
    signals.push(as_token("[[2]]"));
    signals.push(as_token("[[6]]"));
    signals.sort_by(|a, b| a.compare(b));
    let mut si = signals.iter().enumerate();
    let (start, _) = si
        .find(|(_, s)| matches!(s.compare(&as_token("[[2]]")), Ordering::Equal))
        .expect("Start tracer wasn't found");
    let (end, _) = si
        .find(|(_, s)| matches!(s.compare(&as_token("[[6]]")), Ordering::Equal))
        .expect("End tracer wasn't found");
    println!("  - Part 2: {}", (start + 1) * (end + 1));
}

#[derive(Debug)]
enum Token {
    List(Vec<Token>),
    Int(usize),
}

impl Token {
    fn compare(&self, rhs: &Token) -> Ordering {
        if matches!(self, Token::Int(_)) && matches!(rhs, Token::Int(_)) {
            let Token::Int(lhs) = self else { todo!() };
            let Token::Int(rhs) = rhs else { todo!() };
            lhs.cmp(rhs)
        } else if matches!(self, Token::List(_)) && matches!(rhs, Token::List(_)) {
            let Token::List(lhs) = self else { todo!() };
            let Token::List(rhs) = rhs else { todo!() };
            lhs.iter()
                .zip(rhs.iter())
                .find_map(|(a, b)| match a.compare(b) {
                    Ordering::Greater => Some(Ordering::Greater),
                    Ordering::Equal => None,
                    Ordering::Less => Some(Ordering::Less),
                })
                .unwrap_or_else(|| lhs.len().cmp(&rhs.len()))
        } else if let Token::Int(i) = self {
            Token::List(vec![Token::Int(*i)]).compare(rhs)
        } else if let Token::Int(i) = rhs {
            self.compare(&Token::List(vec![Token::Int(*i)]))
        } else {
            panic!("This shouldn't happen");
        }
    }
}

// There is actually a bug with this where some list tokens are nested in an
// additional list that interestingly doesn't affect the comparisons.
fn as_token(raw: &str) -> Token {
    if let Ok(i) = raw.parse() {
        return Token::Int(i);
    }
    let mut tokens: Vec<Token> = Vec::new();
    let mut depth = 0;
    let mut cursor = 0;
    for (i, c) in raw.chars().enumerate() {
        match c {
            '[' => {
                if depth == 0 {
                    cursor = i;
                }
                depth += 1
            }
            ']' => {
                depth -= 1;
                if depth == 0 {
                    tokens.push(as_token(&raw[cursor + 1..i]));
                    cursor = i + 2;
                }
            }
            ',' => {
                if depth == 0 && cursor < i {
                    tokens.push(as_token(&raw[cursor..i]));
                    cursor = i + 1;
                }
            }
            _ => {
                // check if end of a list
                if i == raw.len() - 1 {
                    tokens.push(as_token(&raw[cursor..=i]));
                }
            }
        }
    }
    Token::List(tokens)
}
