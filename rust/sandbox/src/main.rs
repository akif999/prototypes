extern crate rand;

use rand::random;
use std::collections::HashMap;
use std::ops::{Add, Mul};
use std::rc::Rc;

type Table = HashMap<String, Vec<String>>;

fn main() {
    {
        let big_val = std::i32::MAX;
        // let x = big_val + 1; // panic by overflow
        let x = big_val.wrapping_add(1);

        let vec = build_vector();

        println!("Hello, world!");
        println!("{}", &vec[0]);
        println!("{}", &vec[1]);
        println!("{}", big_val);
        println!("{}", x);
    }

    {
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
    }

    {
        struct Person {
            name: Option<String>,
            birth: i32,
        }

        let mut composers = Vec::new();
        composers.push(Person {
            name: Some("Palestrina".to_string()),
            birth: 1525,
        });
        println!("{:?}", composers[0].name);
        // let first_name = composers[0].name // ownership error
        // let first_name = std::mem::replace(&mut composers[0].name, None);
        let first_name = composers[0].name.take();
        println!("{:?}", composers[0].name);
        println!("{:?}", first_name);
    }

    {
        #[derive(Copy, Clone)] // define as Copy type
        struct Label {
            number: u32,
        }
        fn printl(l: Label) {
            println!("STAMP: {}", l.number);
        }
        let l = Label { number: 3 };
        printl(l);
        println!("My label number is: {}", l.number);
    }

    {
        // jointing ownership by 'Rc' type
        let s: Rc<String> = Rc::new("shirataki".to_string());
        let t: Rc<String> = s.clone();
        let u: Rc<String> = s.clone();
        println!("{}", s);
        println!("{}", t);
        println!("{:?}", u);
        println!();

        let mut table = Table::new();
        table.insert(
            "Gesualdo".to_string(),
            vec![
                "many madrigals".to_string(),
                "Tenebrae Responsoria".to_string(),
            ],
        );
        table.insert(
            "Caravaggio".to_string(),
            vec![
                "The Musicians".to_string(),
                "The Calling of St. Matthew".to_string(),
            ],
        );
        table.insert(
            "Cellini".to_string(),
            vec![
                "Perseus with the head of Medusa".to_string(),
                "a salt cellar".to_string(),
            ],
        );
        show(&table);
        println!();
    }

    {
        // borrowing
        let r;
        {
            let x = 1;
            r = &x;
        }
        // println!("{}", *r); // bad: reads memory 'x' used to occupy
    }

    {
        static mut STASH: &i32 = &128;
        fn f(p: &'static i32) {
            unsafe {
                STASH = p;
            }
        }
        static WORTH_POINTING_AT: i32 = 1000;
        f(&WORTH_POINTING_AT);
        unsafe {
            println!("{}", STASH);
        }

        fn g<'a>(p: &'a i32) -> i32 {
            let t = *p + 10;
            t
        }
        let x = 10;
        let y = g(&x);
        println!("{}", y);
    }

    {
        struct S<'a, 'b> {
            x: &'a i32,
            y: &'b i32,
        }

        let x = 10;
        let r;
        {
            let y = 20;
            {
                let s = S { x: &x, y: &y };
                r = s.x;
            }
        }
    }

    {
        // let score = match card.rank {
        //     Jack => 10,
        //     Queen => 10,
        //     Ace => 11
        // }; error: nonexhaustive patterns
    }

    {
        let mut v = vec![5, 1, 2, 4, 3];
        let vc = v.clone();
        println!("before v:");
        for elt in vc {
            println!("{}", elt);
        }
        quicksort(&mut v);
        println!("after v:");
        for elt in v {
            println!("{}", elt);
        }
        let mut va = vec![3, 1, 2, 4];
        let vac = va.clone();
        println!("before va:");
        for elt in vac {
            println!("{}", elt);
        }
        quicksort(&mut va);
        println!("after va:");
        for elt in va {
            println!("{}", elt);
        }
    }

    {
        let is_even = |x| x % 2 == 0;
        // let is_even = |x: u64| -> { x % 2 == 0 };
        println!("is_even(2) is {}", is_even(2));
        println!("is_even(3) is {}", is_even(3));
    }

    {
        let x = random::<i32>();
        let y = random::<char>();
        println!("random value: {}", x);
        println!("random value: {}", y);
    }
}

fn dot<N>(v1: &[N], v2: &[N]) -> N
where
    N: Add<Output = N> + Mul<Output = N> + Default + Copy,
{
    let mut total = N::default();
    for i in 0..v1.len() {
        total = total + v1[i] * v2[i];
    }
    total
}

#[test]
fn test_dot() {
    assert_eq!(dot(&[1, 2, 3, 4], &[1, 1, 1, 1]), 10);
    assert_eq!(dot(&[53.0, 7.0], &[1.0, 5.0]), 88.0);
}

fn show(table: &Table) {
    for (artist, works) in table {
        println!("works by {}:", artist);
        for work in works {
            println!("  {}", work);
        }
    }
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

fn quicksort<T: Ord>(slice: &mut [T]) {
    if slice.len() <= 1 {
        return;
    }
    let pivot_index = partition(slice);
    quicksort(&mut slice[..pivot_index]);
    quicksort(&mut slice[pivot_index + 1..]);
}

fn partition<T: Ord>(slice: &[T]) -> usize {
    // println!("fn partition -> slice.len(): {}", slice.len());
    return slice.len() / 2;
}

#[test]
fn test_partition() {
    let mut t1 = vec![1];
    let mut t2 = vec![1, 2];
    let mut t3 = vec![1, 2, 3];
    let mut t4 = vec![1, 2, 3, 4];
    let mut t5 = vec![1, 2, 3, 4, 5];
    assert_eq!(partition(&mut t1), 0);
    assert_eq!(partition(&mut t2), 1);
    assert_eq!(partition(&mut t3), 1);
    assert_eq!(partition(&mut t4), 2);
    assert_eq!(partition(&mut t5), 2);
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
