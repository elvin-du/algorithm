use std::cell::RefCell;
use std::rc::Rc;

fn main() {
    let mut queue = Queue::new();
    queue.push(1);
    queue.push(2);
    queue.push(3);

    println!("{:?}", queue);
    println!("{:?}", queue.pop());
    println!("{:?}", queue);
}

#[derive(Debug)]
struct Queue {
    head: Option<Rc<RefCell<QueueNode>>>,
    tail: Option<Rc<RefCell<QueueNode>>>,
}

impl Queue {
    fn new() -> Queue {
        Queue { head: None, tail: None }
    }

    fn push(&mut self, v: i32) {
        let node = Rc::new(RefCell::new(QueueNode::new(v)));
        match self.tail {
            None => {
                self.tail = Some(node.clone());
                self.head = Some(node.clone());
            }

            Some(ref mut x) => {
                x.borrow_mut().next = Some(node.clone());
                self.tail = Some(node.clone());
            }
        }
    }

    fn pop(&mut self)->Option<i32>{
        match self.head {
            None => None,
            Some(ref mut x)=>{
                let val = x.borrow().val;
                let node = x.borrow_mut().next.take();
                self.head = node;
                Some(val)
            }
        }
    }
}

#[derive(Debug, PartialEq)]
struct QueueNode {
    val: i32,
    next: Option<Rc<RefCell<QueueNode>>>,
}

impl QueueNode {
    fn new(v: i32) -> QueueNode {
        QueueNode { val: v, next: None }
    }
}

