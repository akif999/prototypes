use std::rc::Rc;

fn main() {
    let big_val = std::i32::MAX;
    // let x = big_val + 1; // panic by overflow
    let x = big_val.wrapping_add(1);

    let vec = build_vector();

    println!("Hello, world!");
    println!("{}", &vec[0]);
    println!("{}", &vec[1]);
    println!("{}", big_val);
    println!("{}", x);

    let text = "I see the eigenvalue in thine eye";
    let (head, tail) = text.split_at(21);
    println!("{}", head);
    println!("{}", tail);

    let arr = [0; 5];
    for e in arr.iter() {
        println!("arr: {}", e);
    }

    let mut v = vec![2, 3, 5, 7];
    println!("folded v :{}", v.iter().fold(1, |a, b| a * b));

    let mut va = Vec::new();
    println!("va cap: {}", va.capacity());
    // println!("va[2] :{}", va[2]); // in this time, va.len() == 0. so it will panic.
    va.push(1);
    va.push(2);
    va.push(3);
    println!("va cap: {}", va.capacity());
    println!("va[2] :{}", va[2]);
    println!();

    // v(vector) to slice
    print(&arr);
    print(&arr[0..2]);
    // arr(array) to slice
    print(&v);
    print(&v[..2]);

    struct Person { name: Option<String>, birth: i32 }

    let mut composers = Vec::new();
    composers.push(Person { name: Some("Palestrina".to_string()),
                            birth: 1525});
    println!("{:?}", composers[0].name);
    // let first_name = composers[0].name // ownership error
    // let first_name = std::mem::replace(&mut composers[0].name, None);
    let first_name = composers[0].name.take();
    println!("{:?}", composers[0].name);
    println!("{:?}", first_name);

    #[derive(Copy, Clone)]  // define as Copy type
    struct Label { number: u32 }
    fn printl(l: Label) { println!("STAMP: {}", l.number); }
    let l = Label { number: 3 };
    printl(l);
    println!("My label number is: {}", l.number);

    // jointing ownership by 'Rc' type
    let s: Rc<String> = Rc::new("shirataki".to_string());
    let t: Rc<String> = s.clone();
    let u: Rc<String> = s.clone();
    println!("{}", s);
    println!("{}", t);
    println!("{:?}", u);
}

/// iterator for slice
/// it apply for both array and vector
fn print(n: &[u32]) {
    for elt in n {
        println!("{}", elt)
    }
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

#[test]
fn test_integer_method() {
    assert_eq!(2u16.pow(4), 16);
    assert_eq!((-4i32).abs(), 4);
    assert_eq!(0b101101u8.count_ones(), 4);
}

#[test]
fn test_array_definition() {
    let lazy_caterer: [u32; 6] = [1, 2, 3, 4, 5, 6];

    assert_eq!(lazy_caterer[3], 4);
}
