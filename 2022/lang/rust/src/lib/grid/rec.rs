use super::point::Point;

#[derive(Debug, Clone, Copy)]
pub struct Rec {
    points: [Point; 2],
    //origin: Point,
    // Height is a distance along the y-axis
    //height: i32,
    // Width is a distance along the x-axis
    //width: i32,
}

impl Rec {
    pub fn from(a: Point, b: Point) -> Self {
        let mut points = [a, b];
        // TODO: test to make sure that this sort works as expected
        points.sort();
        Self { points }
    }

    pub fn min(&self) -> Point {
        self.points[0]
    }

    pub fn min_mut(&mut self) -> &mut Point {
        &mut self.points[0]
    }

    pub fn max(&self) -> Point {
        self.points[1]
    }

    pub fn max_mut(&mut self) -> &mut Point {
        &mut self.points[1]
    }

    pub fn contains(&self, point: &Point) -> bool {
        let min = self.min();
        let max = self.max();
        min.x <= point.x && point.x <= max.x && min.y <= point.y && point.y <= max.y
    }
}
