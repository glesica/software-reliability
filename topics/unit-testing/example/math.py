def double(x: int) -> int:
    """
    >>> double(0)
    0
    >>> double(2)
    4
    """
    return x * 2

def average(x) -> float:
    s = x[0]
    for v in x[1:]:
        s += v
    return s / (len(x) - 1)

