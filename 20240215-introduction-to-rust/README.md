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

## What _won't_ we cover today

_If I discover we have enough time, we'll cover them in order until we run out of time._

* Generics, more traits and lifetimes
* Iterators and closures, functional programming techniques
* Smart pointers
* Threading and synchronisation
* Asynchronous Rust (it's a whole other beast)