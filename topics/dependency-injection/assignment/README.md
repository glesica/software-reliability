# Dependency Injection Homework

This assignment deals with an implementation of the game Tic-Tac-Toe. It is
contained in the `tictactoe` directory, or you can [download](tictactoe.zip) it
as a ZIP file.

To run, you can use `python3 -m tictactoe`.

Our goal is to figure out how we can rewrite this program to be easier to test.
Right now, the entire thing is implemented as one big method. We could use some
kind of automation testing, but it would be much simpler if we could use unit
tests for most of the code.

In order to improve it, you'll need to figure out which dependencies you can
extract and replace with test doubles through dependency injection. You might
also think about how to extract helper methods that better align with the single
responsibility principle.

Rewrite the code and add appropriate tests. Submit your rewritten module as a
ZIP file.
