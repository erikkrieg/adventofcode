use std::{
    cmp::Ordering,
    fs::File,
    io::{self, BufRead, BufReader},
};

pub fn solve() -> io::Result<()> {
    let input = File::open("input/day-13.txt")?;
    test();
    part_one(BufReader::new(input));
    Ok(())
}

fn part_one(reader: BufReader<File>) {
    let signals: Vec<Token> = reader
        .lines()
        .filter_map(|l| l.ok())
        .filter(|line| !line.is_empty())
        .map(|l| as_token(&l))
        .collect();
    let sum = unordered_signals(&signals);
    println!("  - Part 1: {sum}");
}

fn unordered_signals(signals: &[Token]) -> usize {
    signals
        .array_chunks::<2>()
        .enumerate()
        .filter_map(|(i, [a, b])| match a.compare(b) {
            Ordering::Greater => None,
            _ => Some(i + 1),
        })
        .sum()
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

fn test() {
    let signals: Vec<Token> = "
        [1,1,3,1,1]
        [1,1,5,1,1]

        [[1],[2,3,4]]
        [[1],4]

        [9]
        [[8,7,6]]

        [[4,4],4,4]
        [[4,4],4,4,4]

        [7,7,7,7]
        [7,7,7]

        []
        [3]

        [[7,0,7,[4,[7,7],[2,1,9,6],8],10],[],[]]
        []

        [[[10],3,[],1,4],[1,7,[1,[9,1],[5,0]],[[],10],2],[[8],[0,[2,4,2,5],[7,3]],6],[[2,[8,4]],[8,[6,7,6,2]],1,1,[1,3]],[2,0,6]]
        [[],[[]],[],[0,10,[9,0,7],[3,9]],[]]

        [[[]]]
        [[]]

        [1,[2,[3,[4,[5,6,7]]]],8,9]
        [1,[2,[3,[4,[5,6,0]]]],8,9]"
        .to_string()
        .lines()
        .map(|l| l.trim())
        .filter(|l| !l.is_empty())
        .map(as_token)
        .collect();
    let sum = unordered_signals(&signals);
    println!("  - Part 0: {sum}");
}
