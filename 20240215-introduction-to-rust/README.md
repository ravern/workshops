---
marp: true
class: invert
---

# **Introduction to Rust**

**Ravern Koh**
NUS Hackers
Hackerschool
15 February 2024

---

![w:256](./images/profile.png)

## About Me

* Year 2 studying Computer Science
* Enjoys webdev and proglangs
* Been writing Rust on/off for ~4 years

---

![w:128](./images/rust-logo.png)

## The Rust Programming Language

_Lifted shamlessly off the [Rust website](https://www.rust-lang.org/)_

* Performance — no runtime or garbage collector, embedded support
* Reliability — strong type system, memory and thread safety guarantees
* Productivity — actually good compiler errors, tooling, editor support

---

## What we'll cover

* Much of today's content structure is taken from ["The Book"](https://doc.rust-lang.org/book/)
* I aim to summarise it in a succint way, and how I understand it
* Hopefully that's helpful to you

---

## Okay, so what _will_ we cover?

* Data types, functions, control flow
* Ownership and memory model
* Composing data with structs
* Composing more data with enums, pattern matching
* "Null" and error handling
* Polymorphism with traits (somewhat)
* Crates and modules

---

## What we _won't_ cover today

_If I discover we have enough time, we'll cover them in order until we run out of time._

* Generics, more traits and lifetimes
* Iterators and closures, functional programming techniques
* Smart pointers
* Threading and synchronisation
* Asynchronous Rust (it's a whole other beast)

---

## Hello, world!

We'll be using [Rust playground](https://play.rust-lang.org) for convenience today.

```rust
fn main() {
    println!("Hello, world!");
}
```

For the sake of brevity, unless I explicitly write `fn main...` and the some code outside of it, most of the code in the rest of this workshop should be written within the `main` function.

---

## Variables and mutability

```rust
let x = 5;
x = 6;
// Compile error! Try adding `mut` after `let`.

let y = 7;
let y = 8;
// Successfully compiles. Why?

let mut z = 9;
z = true;
// Compile error! Try to fix it.
```

---

## Data types

```rust
let unsigned_integer: u32 = 5; // also u8, u16, u64, u128, usize
let signed_integer: i32 = -63; // also i8, i16, i64, i128, usize
let float: f32 = 3.14159;      // also f64
let boolean: bool = true;      // and false

let character: char = '💨';
let tuples: (u32, i32) = (9, -3);
let unit: () = ();
```

---

## Slices and strings

```rust
let array: [u32; 5] = [1, 2, 3, 4, 5]; // array of u32 of length 5
let slice: &[u32] = &array[1..3];      // slice referencing original array

let string_slice: &str = "A UTF-8 encoded slice of bytes 🤡";
let string: String = "An owned string".to_string(); // more on "own"ing later
```

---

## Functions

```rust
fn foo(number: u32) {
    println!("Hello from `foo`!");
}

fn bar(condition: bool) -> u32 {
    println!("Hello from `bar`!");
    3433 + 343 // Notice the lack of `;` here.
}

fn main() {
    foo(6);
    let number = bar(false);
    println!("The number is {}", number);
    // Check the output.
}
```

---

## Control flow

```rust
if true && false || true && true {
    println!("It's true :)");
} else {
    println!("It's false :(");
}

match 6 + 32 - 28 {
    3 => {
        println!("It's three!");
    }
    5 => {
        println!("It's five!");
    }
    _ => {
        println!("It's none of the above numbers...");
    }
}

let result = if true && true {
    34
} else {
    438
};
```

---

## Control flow

```rust
for index in 0..6 {
    println!("The index is {}", index);
}

while true && false || true && false {
    println!("Looping...");
}
```

---

## Ownership and memory model

* Rust handles memory management in a unique way
* No need for garbage collector, no need for manual free-ing of memory
* It is termed the "ownership model" and has vast implications on how you write Rust

*If there's one thing to takeaway from this workshop, it would be an understanding of the ownership model of Rust*

---

## Rules of ownership

These are lifted from [the book](https://doc.rust-lang.org/book/ch04-01-what-is-ownership.html#ownership-rules).

* Each value in Rust has an *owner*.
* There can only be **one** owner at a time.
* When the owner goes out of scope, the value will be *dropped*.

---

## Variable scope

```rust
{                                // open a new scope, s is not valid yet
    let s = "Some string slice"; // s is now valid from this point
    
    // do stuff with s
    // e.g. println!("{}", s);

}                                // scope is over, s is no longer valid
```

---

## Ownership model example: `String`

```rust
```