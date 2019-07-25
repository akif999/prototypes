use std::ops::{Deref, DerefMut};

struct Selector<T> {
    elements: Vec<T>,
    current: usize,
}

impl<T> Deref for Selector<T> {
    type Target = T;
    fn deref(&self) -> &T {
        &self.elements[self.current]
    }
}

impl<T> DerefMut for Selector<T> {
    fn deref_mut(&mut self) -> &mut T {
        &mut self.elements[self.current]
    }
}

fn main() {
    let mut s = Selector {
        elements: vec!['x', 'y', 'z'],
        current: 2,
    };
    assert_eq!(*s, 'z');

    assert!(s.is_alphabetic());

    *s = 'w';
    assert_eq!(s.elements, ['x', 'y', 'w']);
}
