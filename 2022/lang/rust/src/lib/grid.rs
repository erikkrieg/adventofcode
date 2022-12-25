use std::ops::{Add, Sub};

#[derive(PartialEq, Eq, Hash, Debug, Default, Ord, PartialOrd, Clone, Copy)]
pub struct Point {
    pub x: i32,
    pub y: i32,
}

impl Point {
    pub fn new() -> Self {
        Point::default()
    }

    pub fn from(x: i32, y: i32) -> Self {
        Point { x, y }
    }

    pub fn is_adjacent(&self, other: &Point) -> bool {
        self.x.abs_diff(other.x).max(self.y.abs_diff(other.y)) <= 1
    }
}

impl Add<(i32, i32)> for Point {
    type Output = Self;
    fn add(self, other: (i32, i32)) -> Self {
        Point {
            x: self.x + other.0,
            y: self.y + other.1,
        }
    }
}

impl Sub<(i32, i32)> for Point {
    type Output = Self;
    fn sub(self, other: (i32, i32)) -> Self {
        Point {
            x: self.x - other.0,
            y: self.y - other.1,
        }
    }
}
