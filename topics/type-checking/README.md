# Types

Static type systems embed additional metadata within the source code of our
programs, which can reduce the cognitive overhead of maintenance and
refactoring.

## Introduction

In programming, a type, or data type, is a collection of one or more pieces of
data that supports a particular set of operations. For example, an integer type
might consist of 32 bits of storage and support addition, subtraction, and other
arithmetic operations.

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
