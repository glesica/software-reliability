# Data Types

In programming, a "data type" (or just "type") is a collection of one or more
pieces of data that supports a particular set of operations. For example, an
integer type might consist of 32 bits of storage and support addition,
subtraction, and other arithmetic operations.

Another way of looking at types is that they give meaning to raw data. For
example, say we have the value `65` stored at some location in memory, what does
it represent? It could be a number, or it could be a capital "A" (ASCII
character), or it could be interpreted as `true` (boolean) since it is non-zero.
Going a step further, it could be a temperature in celsius, or a distance in
miles, or a frequency in hertz.

Of course, the computer itself, the CPU, doesn't care about types like integers
and booleans, and most certainly doesn't care about celcius or miles. To the
computer, there are only strings of binary digits (bits) and the operations that
can be performed on them. But that's why virtually all software, aside from
device drivers and the like, which need to interface with the hardware directly,
is written in high-level programming languages that supply additional
abstractions and allow us to add our own.

The meaning behind the value, the type, tells us what we can logically do with
it and how we ought to present it to the user. If we interpret that `65` as an
ASCII character, then dividing it by `2` probably doesn't make any sense. On the
other hand, if `65` represents a distance between two points on a map, then
dividing it in half could make perfect sense!

Different programming languages handle types differently. Some require
variables, and even literal expressions, to be annotated with their types,
others do not. Some will automatically convert data between types when needed,
others will not.

There are many ways to implement types in a programming language, and there are
many different opinions about what is most productive or effective. The most
important thing for programmers is to form opinions from a place of knowledge
and understanding rather than habit and prejudice.

To that end, we are going to think a bit about what types are and what they
mean, how they can implemented in programming languages, and then explore how
several different real languages handle them.

## Prerequisites

We're going to use several different tools and languages for this little
adventure. I've created a Docker image that contains all of them, so you
shouldn't need to install anything additional. However, if you'd like to install
things locally, feel free.

## Types in the World

Believe it or not, computer scientists didn't invent types, they have always
existed separately as an abstract concept and we interact with them every day of
our lives. Everything in the world can be thought of as having one or more
types. In fact, it would only be a little dramatic to say that reality itself is
strongly, if not statically, typed.

Think about your intuitive interactions with physical objects. You wouldn't
expect to drink your morning coffee from your phone, and you wouldn't expect a
coffee mug to run Instagram.

This is because you implicitly understand that those objects possess certain
capabilities, and lack others. Your phone can't hold a substantial quantity of
liquid, and your coffee mug probably doesn't have a microprocessor and screen,
let alone the necessary software to install and run a social media app.

This is how types work in programming languages as well. Certain data types are
useful in certain situations, and not in others and we, as programmers, are
tasked with choosing the correct types for each circumstance we encounter.

## Untyped Languages

Some languages have little to no concept of types at all. In general, these
languages are either very low-level or widely considered "toys".

Assembly languages, which are about as low-level as programming languages get,
for example, generally have only a vague notion of types, usually limited to the
behavior of the CPU itself.

At the other end of the spectrum,
[Brainfuck](https://en.wikipedia.org/wiki/Brainfuck), a "toy" language, has no
concept of types beyond the fact that memory consists of integer bytes.

In both cases, it is entirely up to the programmer to give meaning to a
particular value stored in memory. For example, if a particular byte is used to
store an ASCII character, then the programmer must remember not to assign it a
value that is unavailable for ASCII encoding.

## Strong and Weak Typing

Strongly typed languages, such as Python, associate a specific type with each
piece of data that determines how the data can be used. For example, a string
can be "added" (concatenated) to another string, but it cannot be "added" to
an integer:

```python
>>> "hello" + "5"
'hello5'
>>> "hello" + 5
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
TypeError: can only concatenate str (not "int") to str
```

On the other hand, weakly typed languages still associate a type with each piece
of data, but they will coerce data between types (sometimes even in surprising
and unintuitive ways) in order to avoid crashes. For example, compare the
JavaScript interpreter session below to the Python session above:

```javascript
> "hello" + "5"
'hello5'
> "hello" + 5
'hello5'
```

These results shouldn't surprise us because JavaScript is weakly typed (note
that the term "weak" is not meant as a pejorative, it is simply descriptive).
We still have types in JavaScript, but the language may change those types at
runtime without telling anyone about it. As programmers, when we make mistakes
in a weakly typed language, we often get incorrect results rather than errors.

The "line" between strong and weak typing isn't terribly bright or well-defined.
In fact, both terms are a bit fuzzy. For this reason, it is probably better to
think of the difference between strong and weak typing as a continuum rather
than a dichotomy. So, while we may not always be able to decide whether a given
language should be considered "strongly" or "weakly" typed, we can usually
decide, between any two languages, which is *more* strongly (or weakly) typed.

### Benefits of Strong Typing

### Benefits of Weak Typing

## Static and Dynamic Typing

Static type systems embed additional metadata within the source code of our
programs, which can reduce the cognitive overhead of maintenance and
refactoring and prevent certain bugs. Dynamic type systems, on the other hand,
allow for incredibly flexible program designs.

Students sometimes lament that statically typed languages like Java are
difficult to use. However, this difficulty usually stems from an understandable,
but incorrect, belief that types limit what a programmer can do with a language.
On the contrary, static types allow the programmer to embed a great deal of
additional information in the code itself, often with a minimal increase in
verbosity.

## Type Inference

## Compound Types

## Union and Intersection Types

## Interfaces

In the physical world, when you walk into a room you've never been in before,
you can recognize a lot of objects by their general shape. For instance, you can
usually tell if something is a chair, or a bowl, or a broom.

## Optional Typing

## Top and Bottom Types

## Structural and Nominal Typing

## Parametric Types

Sometimes called "generics" or "generic types", parametric types simply allow us
to add variables to our types. These work just like function or method
parameters and, in fact, we call these variables "type parameters".  The
difference is that whereas function parameters end up replaced by concrete data,
type parameters end up replaced by concrete types.

Parametric type syntax varies from language to language. We will use Rust in the
examples below because its syntax is relatively standard and it's just a fun
language.

Imagine a function with a single parameter, declared with a particular type:

```rust
fn foo(x: i64) {
    // ...
}
```

Within the body of this function, we know `x` holds a 64 bit integer because its
type is `i64`. We don't know which integer it holds, however, and that is only
determined when the function is called and passed an argument.

In the meantime, however, we are perfectly free to manipulate `x` in any way
that makes sense for its type. For example, we could double it and print the
result.

```rust
fn foo(x: i64) {
    // ...
    println!("{}", x * x);
}
```

This will work just fine because we know that `x` is an integer, and we also
know that integers can be multiplied together.

Type parameters work the same way. Below, we have a similar function that has,
effectively, two parameters. One is, again, `x`. The other is the type of `x`,
called (by convention) `T` (for "type").

```rust
fn foo<T>(x: T) {
    // ...
}
```

Now there are two things we don't know within the body of this function: the
value of `x` (as before) *and* the type of `x`. This means that we can no longer
assume that `x` can be multiplied since `T` could be, say, `String`, which would
mean that `x` is a string (which can't be multiplied).

```rust
fn foo<T>(x: T) {
    println!("{}", x * x); // <-- ERROR
}
```

If we try to compile this code we will get an error because the Rust compiler
knows that there are potential values of `T` that won't work (which would cause
the program to crash).

```
➜ rustc parametric.rs
error[E0369]: cannot multiply `T` by `T`
 --> parametric.rs:2:22
  |
2 |     println!("{}", x * x);
  |                    - ^ - T
  |                    |
  |                    T
  |
help: consider restricting type parameter `T`
  |
1 | fn foo<T: std::ops::Mul<Output = T>>(x: T) {
  |         +++++++++++++++++++++++++++
```

This means that we are limited to operations supported by every possible type,
and there aren't many of these. In fact, it isn't even possible to print any value,
so the version below won't compile either.

```rust
fn foo<T>(x: T) {
    println!("{}", x); // <-- ERROR
}
```

The most obvious thing we can do here is restrict (or "bound") the value of the
type paramter. As the syntax suggests, this is a little like specifying a type
for the type. In this case, we require that the type passed for `T` is
compatible with `std::fmt::Display`{.rust}, which is a "trait" (like an
interface) that allows a piece of data to be printed.

```rust
use std::fmt::Display;

fn foo<T: Display>(x: T) {
    println!("{}", x);
}
```

Now we can call `foo` with any value that is compatible with `std::fmt::Display`
and the code will compile and print the value when run.

Bounded type parameters are often what we want. However, there are uses for
unbounded type parameters as well. For example, the Rust standard library
contains a type called `Vec<T>`{.rust} with an unbounded type parameter. This
type implements a "vector" (similar to an array or list). The type parameter `T`
specifies the type of the elements stored in a given vector. Since the vector
doesn't actually *do* anything with the elements it stores, there's no need to
bound the type paramter.

For example, below we create a vector to store integers and store a value in it.
When we extract that value back out, it has the correct type even though
`Vec<T>`{.rust} itself doesn't know anything about the type of data we stored
(since the type parameter is unbounded). This means that from "inside" the
vector, no operations can be performed on the data stored in the vector, but
from "outside" (where we know the type), we can treat the values like the
integers they are.

```rust
  let mut numbers: Vec<i64> = Vec::new();
  numbers.push(5);
  let x = numbers.get(0).unwrap(); // x is an i64
```

## Liskov Substitution Principle

The [Liskov Substitution Principle](https://doi.org/10.1145%2F197320.197383)
(LSP) is a well-known principle in object-oriented programming that, put simply,
states that it should be possible to replace an instance of a particular type
with an instance of any of its subtypes without changing the behavior of the
program in question.

Most statically typed languages enforce some level of compliance with the LSP,
at least purely as a matter of types (behavior aside). But it is still important
for us, as programmers, to bear it in mind. Additionally, it is easier to
violate the LSP by accident in dynamically typed languages.

## Type Relationships
