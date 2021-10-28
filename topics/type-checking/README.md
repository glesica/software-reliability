# Types

Static type systems embed additional metadata within the source code of our
programs, which can reduce the cognitive overhead of maintenance and
refactoring.

## Introduction

In programming, a type, or data type, is a collection of one or more pieces of
data that supports a particular set of operations. For example, an integer type
might consist of 32 bits of storage and support addition, subtraction, and other
arithmetic operations.

Another way of looking at types is that they give meaning to raw data. For
example, say we have the value `65` stored at some location in memory, what does
it represent? It could be a number, or it could be a capital "A" (ASCII
character), or it could be interpreted as `true` (boolean) since it is non-zero.
Going a step further, it could be a temperature in celcius, or a distance in
miles, or a frequency in hertz.

Of course, the computer itself, the CPU, doesn't care about types like integers
and booleans, and most certainly doesn't care about celcius or miles. To the
computer, there are only strings of binary digits (bits) and the operations that
can be performed on them. But that's why virtually all software, aside from
device drivers and the like which need to interface with the hardware directly,
is written in high-level programming languages that supply additional
abstractions and allow us to add our own.

The meaning behind the value, the type, tells us what we can logically do with
it and how we ought to present it to the user. If we interpret `65` as a
boolean, for instance, then negating the value isn't a very useful thing to do,
and would probably represent a programming error. On the other hand, negating a
distance in miles might make perfect sense (depending on what we're doing).

Different programming languages handle types differently. Some require
variables, and even literal expressions, to be annotated with their types,
others do not. Some will automatically convert data between types when needed,
others will not.

Students sometimes lament that statically typed languages like Java are
difficult to use. However, this difficulty usually stems from an understandable,
but incorrect, belief that types limit what a programmer can do with a language.
On the contrary, types allow the programmer to embed a great deal of additional
information in the code itself, often with a minimal increase in verbosity.

That being said, there are many ways to implement types in a programming
language, and there are many different opinions about what is most productive or
effective. The most important thing for programmers is to form opinions from a
place of knowledge and understanding rather than habit and prejudice.

To that end, we are going to think a bit about what types are and what they
mean, and then explore how several different languages handle types.

## Prerequisites

We're going to use several different tools and languages for this little
adventure. I've created a Docker image that contains all of them, so you
shouldn't need to install anything additional. However, if you'd like to install
things locally, feel free.

## Types in the World

Believe it or not, computer science didn't invent types. In reality, everything
in the world can be thought of as belonging to one or more types. In fact, it
would only be a little dramatic to say that reality itself is strongly, if not
statically, typed.

Think about your intuitive interactions with physical objects. You wouldn't
expect to drink your morning coffee from your phone, and you wouldn't expect a
coffee mug to run Instagram.

This is because you implicitly understand that those objects possess certain
capabilities, and lack others. Your phone can't hold a substantial quantity of
liquid, and your coffee mug probably doesn't have a microprocessor and screen,
let alone the necessary software to install and run a social media app.

This is how types work in programming languages as well. Certain data types are
useful in certain situations, and not in others.

Let's look at some code. Open the file `examples.js`, you can also run it with
the command `node examples.js` (from within the course Docker container). Even
though there aren't any types evident in the source code, they are certainly
present. Some of the function invocations do surprising things. Some of them
"work" even though they shouldn't, and some seem like they should cause a crash,
but don't.

These results shouldn't surprise us because JavaScript uses what is typically
called a "weak" type system (the term "weak" is not meant as a pejorative, it is
simply descriptive). This means that the language does whatever it can to avoid
type errors, even if that means doing implicit conversions that don't
necessarily make a lot of sense.

For example, in JavaScript, the expression `"" + null` results in the string
`"null"`, not an error as we might have assumed. The inverse is also true,
`null + ""` evaluates to `"null"`.

So we still have types in JavaScript, but the language may change types at
runtime without telling anyone about it. As programmers, when we make mistakes
in a weakly typed language, we often get incorrect results rather than errors.

## Untyped Languages

Some languages have little to no concept of types at all. In general, these
languages are either very low-level or are mostly considered "toys".

Assembly languages, for example, generally have only a vague notion of types,
usually limited to the behavior of the CPU itself.

At the other end of the spectrum,
[Brainfuck](https://en.wikipedia.org/wiki/Brainfuck), a "toy" language, has no
concept of types beyond the fact that memory consists of integer bytes. In this
case, it is entirely up to the programmer to give meaning to a given value
stored in memory. For example, if a particular byte is used to store an ASCII
character, then the programmer must remember not to assign it a value that is
unused for ASCII encoding.

## Strong and Weak Typing

## Static and Dynamic Typing

## Type Inference

## Compound Types

## Interfaces

## Optional Typing

## Top and Bottom Types

## Structural and Nominal Typing

## Generic Types

## Liskov Substitution Principle

## Type Relationships
