# Choosing Test Values

Most software testing involves a basic pattern: provide an input that should
generate a known output and then verify that the code generates that output. But
how do we choose good input values?

One option is to test all possible program inputs, and sometimes this is a
reasonable option. For example, if a function takes a boolean value, we should
certainly test both `true` and `false`.

However, in most cases there are simply too many possibilities for this to be
feasible. So, a good rule of thumb is to test what are known as "boundary
values". These are values that occur at the boundary between groups of input
values that have different characteristics. The exact definition of these groups
depends on the type of data and the algorithm under test.

For example, if your program accepts an integer as input, a reasonable way to
split up the input into groups is:

  1. `i > 0`
  2. `i = 0`
  3. `i < 0`

This is because these groups tend to behave differently when various numerical
calculations are applied. For example, you can't divide by zero. Off-by-one
errors are also relevant here since they essentially mean that an algorithm
shifts between groups when it shouldn't.

A wise starting point is to test 1, 0, and -1 (boundary values) as well as the
maximum and minimum integers that can be stored by the type in question. These
are boundary values because integers roll over.

Invalid values are also important to test. For example, `null` is a common
source of errors in programming languages like Java. In dynamically typed
languages you may also want to test values that don't actually make sense, since
they will cause runtime errors. For example, passing a string to a square root
function is probably a good idea, to verify that it produces the correct error.

That being said, it is necessary to think about the problem your code solves in
order to choose groups that will result in robust, useful tests.

## Further Reading

  * [Lecture slides](https://docs.google.com/presentation/d/1sUZlgu-rgpYVtdkJluwh7oh03NS3hu-ksjTV-mAIQ-8/edit?usp=sharing)
  * <http://web.mit.edu/6.005/www/sp16/classes/03-testing/> - includes an
    excellent discussion about how to choose good test values
