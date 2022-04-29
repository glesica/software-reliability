import logging

l = logging.getLogger(__name__)

h = logging.FileHandler("output.log")
f = logging.Formatter("%(asctime)s - %(message)s")
h.setFormatter(f)
l.addHandler(h)


l.setLevel(logging.INFO)

def double(x):
    l.info(f"doubling value {x}")
    return 2 * x

def triple(x):
    l.info(f"tripling value {x}")
    return 3 * x

def divide(x, y):
    l.info(f"dividing value {x} by {y}")
    try:
        return x / y
    except ZeroDivisionError as e:
        l.error(e)
        raise

import sys
value = int(sys.argv[1])
print(double(value))
print(triple(value))
print(divide(10, value))

