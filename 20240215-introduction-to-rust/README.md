---
marp: true
class: invert
---

# **Introduction to Rust**

**Ravern Koh**
NUS Hackers
Hackerschool
7 February 2024

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

## Exercise: Beer Song

Let's get Rust to print out the lyrics to the classic song.

```
99 bottles of beer on the wall, 99 bottles of beer.
Take one down and pass it around, 98 bottles of beer on the wall.

98 bottles of beer on the wall, 98 bottles of beer.
Take one down and pass it around, 97 bottles of beer on the wall.

97 bottles of beer on the wall, 97 bottles of beer.
Take one down and pass it around, 96 bottles of beer on the wall.

...

2 bottles of beer on the wall, 2 bottles of beer.
Take one down and pass it around, 1 bottle of beer on the wall.

1 bottle of beer on the wall, 1 bottle of beer.
Take it down and pass it around, no more bottles of beer on the wall.

No more bottles of beer on the wall, no more bottles of beer.
Go to the store and buy some more, 99 bottles of beer on the wall.
```

---

## Solution: Beer Song

```rust
fn main() {
    for i in (0..100).rev() {
        verse(i);
    }
}

fn verse(n: u32) {
    match n {
        0 => println!("No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n"),
        1 => println!("1 bottle of beer on the wall, 1 bottle of beer.\nTake it down and pass it around, no more bottles of beer on the wall.\n"),
        2 => println!("2 bottles of beer on the wall, 2 bottles of beer.\nTake one down and pass it around, 1 bottle of beer on the wall.\n"),
        n => println!("{} bottles of beer on the wall, {} bottles of beer.\nTake one down and pass it around, {} bottles of beer on the wall.\n", n, n, n - 1),
    }
}
```

---

## Ownership and memory model

* Rust handles memory management in a unique way
* No need for garbage collector, no need for manual free-ing of memory
* It is termed the "ownership model" and has vast implications on how you write Rust

*If there's one thing to takeaway from this workshop, it would be an understanding of the ownership model of Rust*

---

## Ownership model example: `String`

Let's try this seemingly innocent example.

```rust
let s1 = "Hello".to_string();
let s2 = s1;

println!("{}, world!", s1);
```

---

## Compile error!

```
error[E0382]: borrow of moved value: `s1`
 --> src/main.rs:5:28
  |
2 |     let s1 = "Hello".to_string();
  |         -- move occurs because `s1` has type `String`, which does not implement the `Copy` trait
3 |     let s2 = s1;
  |              -- value moved here
4 |
5 |     println!("{}, world!", s1);
  |                            ^^ value borrowed here after move
```

So it seems like `s1` is *moved*, but we're trying to *borrow* it?

---

## Values and variables

We focus here on the distinction between *values* and *variables*.

```rust
let s1 = "Hello".to_string();
```

Here, we see that the variable `s1` *owns* the value `"Hello"`.

```rust
let s2 = s1;
```

We're not really assigning `s1` to `s2`, we're *moving* the value of `s1` into `s2`.

```rust
println!("{}, world!", s1);
```

`println!` wants to use the value `s1`, but it's no longer valid. It has been *moved* into `s2`.

---

## Rules of ownership

These are lifted from [the book](https://doc.rust-lang.org/book/ch04-01-what-is-ownership.html#ownership-rules).

* Each value in Rust has an *owner*.
* There can only be **one** owner at a time.
* When the owner goes *out of scope*, the value will be *dropped*.

---

## Variable scope

What does it mean to go *out of scope*?

```rust
{                                // open a new scope, s is not valid yet
    let s = "Some string slice"; // s is now valid from this point
    
    // do stuff with s
    // e.g. println!("{}", s);

}                                // scope is over, s is no longer valid
```

---

## Are all values moved?

* Some values aren't moved but copied instead.
* *Most* of the types that start with lowercase letters are copied by default.
* This compiles, because the value of `n1` is copied into `n2` instead of being moved.

  ```rust
  let n1: u32 = 324;
  let n2 = n1;  
  println!("{}", n1);
  ```

There is no assignment, there is only *moving* or *copying*.

---

## Ownership and functions

```rust
fn foo(baz: String) {
    println!("I have {}!", baz);
}

fn main() {
    let bar = "a pen".to_string();
    foo(bar);

    println!("I don't have {} :(", bar);
}
```

* The *value* of `bar` is *moved* into the `foo` function call here
* We can't use the variable `bar` in `main` anymore
* The new owner of `"a pen"` is the `baz` parameter/variable in `foo`

---

## Moving with return values

You can move values back to the caller by returning them.

```rust
fn foo(baz: String) -> String {
    println!("I have {}!", baz);
    baz
}

fn main() {
    let bar = "a pen".to_string();
    let qux = foo(bar);

    println!("I have {} :)", qux);
}
```

---

## References and borrowing

Do I really have to keep passing and returning all my values?

```rust
fn foo(baz: &String) {
    println!("I have {}!", baz);
}

fn main() {
    let bar = "a pen".to_string();
    foo(&bar);

    println!("I still have {} :)", bar);
}
```

* Here, we borrow the value of `bar` (instead of moving from it) 
* A reference to the value `"a pen"` is passed into `foo`
* `&String` means a reference to a string
* `&bar` means to take a reference of `bar`

---

## Mutation of owned values

Let's talk about mutation.

```rust
let mut foo = "Hello".to_string();
foo.push_str(", world!");

println!("{}", foo); // Hello, world!
```

---

## Mutation of borrowed values

How about when values are borrowed?

```rust
let foo: String = "Hello".to_string();  // do we need `mut` here?
let bar: &String = &foo;                // maybe a `mut` here instead?
bar.push_str(", world!");

println!("{}", foo);

// Compile error! Try to fix this.
```

---

## Rule of mutation and borrowing

You can either have *multiple immutable borrows*, or *one mutable borrow*.

Let's test this out!

```rust
let mut the_string: String = "Hello".to_string();
let foo: &String = &the_string;         // this is fine
let bar: &String = &the_string;         // this is fine too
let baz: &mut String = &mut the_string; // this is fine too...?
let qux: &String = &the_string;         // this is also fine???
```

Oops! The Rust compiler is really too smart for us. What do you think is happening here?

---

## Rule of mutation and borrowing

Let's try again.

```rust
let mut the_string: String = "Hello".to_string();
let foo: &String = &the_string;         // this is fine
let bar: &String = &the_string;         // this is fine too
let baz: &mut String = &mut the_string; // Compile error!

println!("{}", foo);
println!("{}", bar);
println!("{}", baz);
```

---

## Exercise: Fix the bugs!

Fix the bugs by changing the way we handle value e.g. owned value, reference or mutable reference.

```rust
fn main() {
    let data = "Rust is great!".to_string();
    display_to_user(data);
    add_cheer(&data);
    let uppercase_data = convert_to_uppercase(&data);
    println!("Final form: {}", uppercase_data);
}

fn display_to_user(data: String) {
    println!("Current form: {}", data);
}

fn add_cheer(data: &String) {
    data.push_str(" And so are you!");
}

fn convert_to_uppercase(data: &String) -> &String {
    &data.to_uppercase()
}
```

---

## Solution: Fix the bugs!

```rust
fn main() {
    let mut data = "Rust is great!".to_string();     // Added a `mut` here
    display_to_user(&data);                          // Borrow immutably
    add_cheer(&mut data);                            // Borrow mutably
    let uppercase_data = convert_to_uppercase(data); // Move and return
    println!("Final form: {}", uppercase_data);
}

// Take in an immutable reference
fn display_to_user(data: &String) {
    println!("Current form: {}", data);
}

// Take in a mutable reference
fn add_cheer(data: &mut String) {
    data.push_str(" And so are you!");
}

// Take in an owned `String`
fn convert_to_uppercase(data: String) -> String {
    data.to_uppercase()
}
```

---

# Let's take a break.

Feel free to ask me some questions!

---

## Composing data in Rust

We're now ready to introduce more complex data structures!

* Structs (product types)
* Enums (sum types)
* `Option` and `Result`
* Traits and polymorphism

---

## Grouping data with structs

```rust
struct Rectangle {
    width: u32,
    height: u32,
}

fn main() {
    let rectangle = Rectangle {
        width: 100,
        height: 50,
    };

    println!("width: {}, height: {}", rectangle.width, rectangle.height);

    // Can we print out the whole `Rectangle` at once?
    // e.g. println!("rectangle: {}", rectangle);
}
```

---

## Ownership model with structs

```rust
struct Rectangle {
    width: u32,
    height: u32,
}

fn expand_width(rectangle: &mut Rectangle) {
    rectangle.width += 50;
}

fn main() {
    let mut rectangle = Rectangle {
        width: 100,
        height: 50,
    };

    expand_width(&mut rectangle);
}
```

---

## Methods on structs

```rust
impl Rectangle {
    // This is called an "associated function"
    fn new(width: u32, height: u32) -> Rectangle {
        Rectangle {
            width,
            height,
        }
    }

    // This is your regular old method
    fn area(&self) -> u32 {
        self.width * self.height
    }

    // This is also a method, notice `&mut self`
    fn expand_width(&mut self) {
        self.width += 50;
    }
}

fn main() {
    let mut rectangle = Rectangle::new(100, 50);
    println!("area: {}", rectangle.area());
    rectangle.expand_width();
    println!("expanded area: {}", Rectangle::area(&rectangle)); // same thing
}
```

---

## Different way of grouping data with enums

```rust
enum Color {
    Red,
    Green,
    Blue,
    Yellow,
}

fn main() {
    let red = Color::Red;
    let green = Color::Green;
    let blue = Color::Blue;
    let yellow = Color::Yellow;
}
```

---

## Enums can store data

```rust
enum Message {
    Move { x: i32, y: i32 },
    Write(String),
    ChangeColor(i32, i32, i32),
    Quit,
}

fn send_message(message: Message) {
    // Send it off to the stars...
}

fn main() {
    let move_up = Message::Move { x: 0, y: -32 };
    send_message(move_up);

    let write = Message::Write("I call upon you!".to_string());
    send_message(write);

    let change_color = Message::ChangeColor(42, 32, 45);
    send_message(change_color);

    let quit = Message::Quit;
    send_message(quit);
}
```

---

## Getting data out of an enum

Let's re-use `send_message` from earlier

```rust
enum Message {
    Move { x: i32, y: i32 },
    Write(String),
    ChangeColor(i32, i32, i32),
    Quit,
}

fn send_message(message: Message) {
    // We can use an `if let` expression
    if let Message::Write(text) = &message { // why do we need `&` here?
        println!("The message is: {}", text);
    }

    // Or a `match` expression
    match message {
        Message::Move { x, y } => println!("Position moved to x: {}, y: {}.", x, y),
        Message::Write(_) => println!("Message written."),
        Message::ChangeColor(_, _, _) => println!("Color changed."),
        Message::Quit => println!("BYE BYE."),
    }
}
```

---

## Methods on enums

```rust
impl Message {
    fn send(&self) {
        match self {
            Message::Move { x, y } => println!("Position moved to x: {}, y: {}.", x, y),
            Message::Write(_) => println!("Message written."),
            Message::ChangeColor(_, _, _) => println!("Color changed."),
            Message::Quit => println!("BYE BYE."),
        }
    }
}

fn main() {
    let move_up = Message::Move { x: 0, y: -32 };
    move_up.send();

    let write = Message::Write("I call upon you!".to_string());
    write.send();

    let change_color = Message::ChangeColor(42, 32, 45);
    change_color.send();

    let quit = Message::Quit;
    quit.send();
}
```

---

## Generic types

* Sometimes, we have types that are *parameterized* by other types
* Like when you say "array" in Java, the reply would be "well, array of what?"

  ```rust
  // THIS DOESN'T COMPILE, JUST AN EXAMPLE

  struct Array<T> {
    // some stuff inside
  }

  fn main() {
    let array_of_ints: Array<u32> = Array::new();
    let array_of_strings: Array<String> = Array::new();
    let array_of_arrays_of_ints: Array<Array<u32>> = Array::new();
  }
  ```

---

## `Option`

* Rust has no `null` or `nil`, so you'll never face that dreaded `NullPointerException` (it has no exceptions either)
* That means you, the programmer, has to handle all possible cases of "emptiness" explictly
* We use the `Option` enum for that

  ```rust
  enum Option<T> {
      Some(T),
      None,
  }
  ```

* So here, we say you either have `Some`thing, or `None` at all

---

## Using `Option`

```rust
struct Error {
    message: Option<String>,
    line: u32,
}

fn main() {
    let compile_error_with_message = Error {
        message: Some("You forgot your semicolon, noob.".to_string()),
        line: 32,
    };

    let compile_error_no_message = Error {
        message: None,
        line: 35,
    };

    let compile_error_numerical_message = Error {
        message: Some(524), // Compile error!
        line: 35,
    };
}
```

---

## `Result`

* Remember how Rust doesn't have exceptions? Well what if an operation fails e.g. a file isn't found in the filesystem
* For that, we have `Result`

  ```rust
  enum Result<T, E> {
    Ok(T),
    Err(E),
  }
  ```

---

## Using `Result`

```rust
// If we can divide by 5, return `Ok` with the answer, otherwise return
// an `Err` with the remainder.
fn divide_by_5(n: u32) -> Result<u32, u32> {
    if n % 5 == 0 {
        Ok(n / 5)
    } else {
        Err(n % 5)
    }
}

fn main() {
    let result = divide_by_5(32);
    match result {
        Ok(answer) => println!("Success! The answer is {}.", answer),
        Err(remainder) => println!("Failure. The remainder is {}.", remainder),
    }
}
```

---

## Traits and polymorphism

* Traits are like the `interface` (Java, TypeScript) or `protocol` (Swift) of Rust
* You use them to group common behaviour between different types
* For example, you'd say that both arrays and linked lists are `Iterator`s (a commonly used Rust trait)
* They are instrumental for polymorphism in Rust

---

## Defining a trait

```rust
trait Walk {
    // Returns the new position.
    fn walk(&self, position: u32) -> u32;
}
```

---

## Implementing a trait

```rust
struct Person {
    speed: u32,
}

impl Walk for Person {
    fn walk(&self, position: u32) -> u32 {
        position + self.speed
    }
}

struct Animal {
    is_facing_forward: bool,
}

impl Walk for Animal {
    fn walk(&self, position: u32) -> u32 {
        if self.is_facing_forward {
            position + 5
        } else {
            position - 5
        }
    }
}
```

---

## Using a trait

```rust
fn perform_walk<T: Walk>(walker: T, position: u32) -> u32 {
    walker.walk(position)
}

fn main() {
    let slow_person = Person { speed: 20 };
    perform_walk(slow_person, 32);
    let fast_person = Person { speed: 55 };
    perform_walk(fast_person, 32);
    let animal = Animal { is_facing_forward: false };
    perform_walk(animal, 32);
}
```

---

## Exercise: Tic-Tac-Toe

```rust
enum Player {
    // TODO
}

enum Cell {
    // TODO
}

struct Board {
    cells: [[Cell; 3]; 3],
}

impl Board {
    fn empty() -> Board {
        Board {
            cells: [
                [Cell::Empty, Cell::Empty, Cell::Empty],
                [Cell::Empty, Cell::Empty, Cell::Empty],
                [Cell::Empty, Cell::Empty, Cell::Empty],
            ],
        }
    }

    fn cell(&self, col: usize, row: usize) -> &Cell {
        &self.cells[row][col]
    }

    fn place(/* TODO */) {
        // TODO
    }
}

// ---------------------------------------------------------------

trait Line: Sized {
    fn cells(self, board: &Board) -> (&Cell, &Cell, &Cell);
}

fn check_line_winner</* TODO */>(/* TODO */) -> Option<Player> {
    // TODO
}

struct Row {
    row: usize,
}

impl Line for Row {
    fn cells(self, board: &Board) -> (&Cell, &Cell, &Cell) {
        (
            board.cell(self.row, 0),
            board.cell(self.row, 1),
            board.cell(self.row, 2),
        )
    }
}

struct Col {
    // TODO
}

impl Line for Col {
    // TODO
}

enum Diag {
    TopLeftToBottomRight,
    BottomLeftToTopRight,
}

impl Line for Diag {
    // TODO
}

// --------------------------------------------------------------

fn check_winner(board: &Board) -> Option<Player> {
    // TODO
}

fn main() {
    let mut board = Board::empty();
  
    board.place(1, 1, Player::X);
    board.place(0, 1, Player::O);
    board.place(0, 0, Player::X);
    board.place(2, 2, Player::O);
    board.place(1, 0, Player::X);
    board.place(1, 2, Player::O);
    board.place(2, 0, Player::X);
  
    match check_winner(&board) {
        Some(Player::X) => println!("Player X has won!"),
        Some(Player::O) => println!("Player O has won!"),
        None => println!("No one has won yet."),
    }
}
```

---

## Solution: Tic-Tac-Toe

```rust
enum Player {
    O,
    X,
}

enum Cell {
    Empty,
    Occupied(Player),
}

struct Board {
    cells: [[Cell; 3]; 3],
}

impl Board {
    fn empty() -> Board {
        Board {
            cells: [
                [Cell::Empty, Cell::Empty, Cell::Empty],
                [Cell::Empty, Cell::Empty, Cell::Empty],
                [Cell::Empty, Cell::Empty, Cell::Empty],
            ],
        }
    }

    fn cell(&self, col: usize, row: usize) -> &Cell {
        &self.cells[row][col]
    }

    fn place(&mut self, col: usize, row: usize, player: Player) {
        self.cells[row][col] = Cell::Occupied(player);
    }
}

// ---------------------------------------------------------------

trait Line: Sized {
    fn cells(self, board: &Board) -> (&Cell, &Cell, &Cell);
}

fn check_line_winner<T: Line>(line: T, board: &Board) -> Option<Player> {
    let cells = line.cells(board);
    match cells {
        (Cell::Occupied(Player::X), Cell::Occupied(Player::X), Cell::Occupied(Player::X)) => {
            Some(Player::X)
        }
        (Cell::Occupied(Player::O), Cell::Occupied(Player::O), Cell::Occupied(Player::O)) => {
            Some(Player::O)
        }
        _ => None,
    }
}

struct Row {
    row: usize,
}

impl Line for Row {
    fn cells(self, board: &Board) -> (&Cell, &Cell, &Cell) {
        (
            board.cell(self.row, 0),
            board.cell(self.row, 1),
            board.cell(self.row, 2),
        )
    }
}

struct Col {
    col: usize,
}

impl Line for Col {
    fn cells(self, board: &Board) -> (&Cell, &Cell, &Cell) {
        (
            board.cell(0, self.col),
            board.cell(1, self.col),
            board.cell(2, self.col),
        )
    }
}

enum Diag {
    TopLeftToBottomRight,
    BottomLeftToTopRight,
}

impl Line for Diag {
    fn cells(self, board: &Board) -> (&Cell, &Cell, &Cell) {
        match self {
            Diag::TopLeftToBottomRight => (board.cell(0, 0), board.cell(1, 1), board.cell(2, 2)),
            Diag::BottomLeftToTopRight => (board.cell(0, 2), board.cell(1, 1), board.cell(2, 0)),
        }
    }
}

// --------------------------------------------------------------

fn check_winner(board: &Board) -> Option<Player> {
    for row in 0..3 {
        let row = Row { row };
        if let Some(winner) = check_line_winner(row, board) {
            return Some(winner);
        }
    }
    for col in 0..3 {
        let col = Col { col };
        if let Some(winner) = check_line_winner(col, board) {
            return Some(winner);
        }
    }
    if let Some(winner) = check_line_winner(Diag::TopLeftToBottomRight, board) {
        return Some(winner);
    }
    if let Some(winner) = check_line_winner(Diag::BottomLeftToTopRight, board) {
        return Some(winner);
    }
    None
}

fn main() {
    let mut board = Board::empty();
  
    board.place(1, 1, Player::X);
    board.place(0, 1, Player::O);
    board.place(0, 0, Player::X);
    board.place(2, 2, Player::O);
    board.place(1, 0, Player::X);
    board.place(1, 2, Player::O);
    board.place(2, 0, Player::X);
  
    match check_winner(&board) {
        Some(Player::X) => println!("Player X has won!"),
        Some(Player::O) => println!("Player O has won!"),
        None => println!("No one has won yet."),
    }
}
```

---

## Thank you for coming!