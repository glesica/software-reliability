#!/usr/bin/env python3

print("PYTHAGORAS")

def pythagoras(a: int, b: int) -> None:
    from math import sqrt
    print(sqrt((a * a) + (b * b)))

pythagoras(2, 3)
pythagoras("2", 3)
pythagoras(True + True, 3)
pythagoras("shoe", 3)
pythagoras(None, 3)

print("\nREPEAT")

def repeat(msg: str, count: int) -> None:
    output = None
    for i in range(count):
        if output is None:
            output = ""
        else:
            output += " "
        output += msg
    print(output)

repeat("hello", 3)
repeat("hello", "3")
repeat("hello", True + True + True)
repeat("hello", "duck")
repeat(None, 3)
