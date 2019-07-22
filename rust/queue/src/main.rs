fn main() {
    // let mut q = Queue::<char>::new();
    let mut q = Queue::new();

    q.push('0');
    q.push('1');
    assert_eq!(q.pop(), Some('0'));
    q.push('2');
    assert_eq!(q.pop(), Some('1'));
    assert_eq!(q.pop(), Some('2'));
    assert_eq!(q.pop(), None);
}

pub struct Queue<T> {
    older: Vec<T>,
    younger: Vec<T>,
}

impl<T> Queue<T> {
    pub fn new() -> Queue<T> {
        Queue {
            older: Vec::new(),
            younger: Vec::new(),
        }
    }
    pub fn push(&mut self, t: T) {
        self.younger.push(t)
    }
    pub fn pop(&mut self) -> Option<T> {
        if self.older.is_empty() {
            if self.younger.is_empty() {
                return None;
            }
            use std::mem::swap;
            swap(&mut self.older, &mut self.younger);
            self.older.reverse();
        }
        self.older.pop()
    }
}
