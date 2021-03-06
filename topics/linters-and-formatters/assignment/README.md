# Linters Homework

We're going to add a "lint" to an existing linter. The lint can do anything you
like, but some suggestions are provided below. Don't worry if your lint isn't
actually helpful. The goal is to gain some experience working with a real-world
linter and adding code to an existing software project. That being said, your
lint does need to *work* (in other words, your code needs to run and do what it
is intended to do).

## Linters

You can choose which linter to modify, but keep in mind that I need to be able
to evaluate your code easily. Pylint is probably a good choice because it is
well-documented and adding a lint ("checker") is straightforward. I will also
(probably) be more familiar with Pylint than any arbitrary linter you might
choose, and therefore more able to help you if you get stuck.

  * [How to Write a Checker](https://pylint.pycqa.org/en/latest/how_tos/custom_checkers.html)

## Sample Ideas

  * Identify variables / functions with curse words in their names.
  * Find comments with your name in them and suggest giving you a compliment
  * Suggest renaming function parameters after different Pokemon
  * Look for "TODO" comments and produce an error if there are too many

But feel free to come up with your own idea!
