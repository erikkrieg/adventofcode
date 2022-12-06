mod depth;

fn main() {
  let depths = depth::get_depths();
  // When referencing `depths` multiple times I hit a compiler error
  // that I don't really understand. Fixed it by "borrowing" depths.
  println!("Depth increased {} times.", depth::count_increases(&depths));
  println!("Depth increased {} times.", depth::count_sliding_window_increases(&depths));
}
