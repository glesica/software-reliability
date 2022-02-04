# Unit Testing

  * [Assignment](assignment/)

As our software is expected to do more, it inevitably becomes more complex.
When this happens, adequate functional testing becomes significantly more
difficult. This doesn't mean that we shouldn't employ functional testing for
complex software (quite the opposite). But it does mean that we need a way to
reduce our dependence on functional testing for verification of the internal
details of our code. Let's work through an example. Say we have the program
below, we can use functional testing to verify it without too much trouble.

```python
import sys

def say_hello(first_name, last_name):
  print('Hello, ' + first_name + ' ' + last_name)

if __name__ == "__main__":
  if len(sys.argv) == 2:
    say_hello(sys.argv[1], '')
  else:
    say_hello(sys.argv[1], sys.argv[2])
```

There are only two execution paths here, and the choice happens immediately and
is based only on how the program was run (the number of command line arguments).
This program is almost trivial to test functionally. But what if we make the
program just a bit more complicated, like so?

```python
# hello.py

import sys

class Person:
  def __init__(self, first_name, last_name=''):
    self.first_name = first_name
    self.last_name = last_name

  def full_name(self):
    if self.last_name == '':
      return self.first_name
    else:
      return self.first_name + ' ' + self.last_name

def say_hello(person: Person):
  if person.last_name == '':
    # We're all casual and such
    print('Sup, ' + person.full_name())
  else:
    # Super formal and stuffy
    print('Greetings to you, ' + person.full_name())

if __name__ == "__main__":
  if len(self.argv) == 2:
    p = Person(self.argv[1])
    say_hello(p)
  elif len(self.argv) == 3:
    p = Person(self.argv[1], self.argv[2])
    say_hello(p)
```

So what's the difference? Well, it's a bit subtle, but now we have logic in two
different places. We have the same basic function we had before (`say_hello`),
but it no longer handles the name formatting. Instead, the name is formatted
inside the `Person` class.

Additionally, there are now simply more code paths (there are now three
branching statements instead of one), that means that we have to run the program
in more ways in order to exercise them all. This can be complicated and error
prone because we have to think about the entire program when choosing our test
inputs. It also tends to be slow since the entire program has to run for each
test, so if there are a large number of tests then the entire test suite may
take a long time to run.

We could certainly continue to use functional tests to verify the entire
program, but that will only get more difficult as the code complexity continues
to increase.

Instead, we could test the `Person` class by itself.  The new tests for the
`Person` class will be what are called "unit tests" because they will cover a
single, isolated "unit" (piece) of code.

## Example

In order to test the `Person` class by itself we will use a separate Python
script. A very simple version might look like this:

```python
# person_test.py

from hello import Person

p = Person('First', '')
assert p.first_name == 'First'
assert p.last_name == ''
assert p.full_name() == 'First'

p = Person('First', 'Last')
assert p.first_name == 'First'
assert p.last_name == 'Last'
assert p.full_name() == 'First Name'
```

So now we can run this script to verify that our `Person` class works as we
expect it to work. We no longer have to worry about whether the name will be
formatted correctly in our main program because we've tested it as part of its
unit, the `Person` class.

We could take this even further and write unit tests for the `say_hello`
function. Then, between our two unit test suites, we'll be confident that the
text output of this program works correctly.

The first thing we should do, however, is change `say_hello` so that it returns
its value instead of print it. This will make it easier to unit test since we
can "capture" the output for verification. The entire `hello.py` file is
reproduced below even though only the `say_hello` and script behaviors have
changed.

```python
# hello.py

import sys

class Person:
  def __init__(self, first_name, last_name=''):
    self.first_name = first_name
    self.last_name = last_name

  def full_name(self):
    if self.last_name == '':
      return self.first_name
    else:
      return self.first_name + ' ' + self.last_name

def say_hello(person: Person) -> str:
  if person.last_name == "":
    # We're all casual and such
    return "Sup, " + person.full_name()
  else:
    # Super formal and stuffy
    return "Greetings to you, " + person.full_name()

if __name__ == "__main__":
  if len(self.argv) == 2:
    p = Person(self.argv[1])
    print(say_hello(p))
  elif len(self.argv) == 3:
    p = Person(self.argv[1], self.argv[2])
    print(say_hello(p))
```

Now we can write some unit tests like the ones above.

```python
# say_hello_test.py

from hello import Person, say_hello

p0 = Person("A")
assert say_hello(p0) == f"Sup, {p0.full_name()}"

p1 = Person("A", "B")
assert say_hello(p1) == f"Greetings to you, {p1.full_name()}"
```

Notice that we don't hard code the full name because the `Person` class might
change in the future and the `say_hello` function doesn't really care how the
full name is computed. For example, the `full_name` method might change so that
it returns `Last, First` instead of `First Last`, but we don't want that change
to cause the `say_hello` tests to fail.

Also notice that our "unit" here is a function, not a class. Unit testing
doesn't prescribe what a unit is, it just means that you are testing a single,
independent chunk of your program.

## When to Unit Test

But how do we know when to use unit tests and when to stick to functional tests?
Unfortunately, that decision is really more of an art than a science. But your
goal, which can help guide your decisions, should be to have robust tests. This
means that they don't fail unless there's an actual problem.

In the example above, we can change either the `Person` class or the `say_hello`
function and only the relevant tests will fail. This is a desirable property as
it simplifies future maintenance on code.

Performance is another consideration. Unit tests are generally expected to run
quickly (so that they can be run often), so if a test requires time-consuming
setup it should probably be a functional test.

## Further Reading

  * [Lecture slides](https://docs.google.com/presentation/d/1BMIU80of7iG-OX2OtsjuKApYSkkapMwqJD_ZPvIN8hU/edit?usp=sharing)
