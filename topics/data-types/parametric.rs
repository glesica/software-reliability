use std::fmt::Display;

fn concrete(x: i64) {
    println!("concrete: {}", x);
}

fn parametric<T: Display>(x: T) {
    println!("parametric: {}", x);
}

fn main() {
    concrete(5);

    parametric(5);
    parametric("hello");

    let mut x: Vec<i64> = Vec::new();
    x.push(5);
    x.push(3);

    let y = x.get(0).unwrap();
    println!("{}", y);
}
