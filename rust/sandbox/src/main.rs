fn main() {
    println!("Hello, world!");
}

fn build_vector() -> Vec<i16> {
    let mut v: Vec<i16> = Vec::<i16>::new();
    v.push(10i16);
    v.push(10i16);
    v
}

fn build_vector_alt() -> Vec<i16> {
    let mut v = Vec::new();
    v.push(10);
    v.push(10);
    v
}

#[test]
fn test_build_vector() {
    let vec = build_vector();
    let vec_alt = build_vector_alt();

    assert_eq!(vec, vec_alt);
}
