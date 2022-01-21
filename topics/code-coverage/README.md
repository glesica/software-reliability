# Code Coverage

Code coverage is, put simply, a measure of how much of the code that makes up an
application is actively tested. In other words, how much of the code actually
executes when the tests are run.

When we talk about measuring "code coverage" we are interested in knowing what
fraction of our code is actually tested by our test suite. For example, let's
consider the following code:

```python
def make_greeting(name):
  if name == 'George':
    return 'Oh, hi there!'
  else:
    return 'Nice to meet you'
```

Now, imagine we were to write some unit tests for this simple function.

```python
import hello

assert say_hello('Megan') == 'Nice to meet you'
```

This is fine, the test calls the function and verifies the output. But the test
doesn't cause every line in the function under test to be run. Specifically, we
never hit the `true` branch of the if-statement. So out of five lines, our test
executes four, making our code coverage 80%. This is a very simple example, but
the idea should be clear.

## Types of Coverage

There are several ways to measure coverage. Above I demonstrated "line"
coverage, in other words, measuring coverage as a percentage of the total lines
in the program exercised by tests.

Another common metric is "branch coverage". Branch coverage essentially counts
the number of "paths" through the program that are covered. At each point where
a decision is made (if-statements, switch-statements, and so on) it counts how
many of the possibilities are run by the test suite, and how many are not. The
code coverage then becomes the fraction of covered paths. In the example above,
the branch coverage would be 50%.

It is also somewhat common to count the percentage of functions or classes
covered by a test suite.

## Chasing 100% Coverage

Coverage is a somewhat controversial concept. It is commonly misused, usually
because people misunderstand it. 

One of the traps that developers and organizations sometimes fall into is
"chasing" 100% code coverage. It seems strange to think that we might not want
100% coverage, why wouldn't we? Clearly our test suite above would have been
improved by adding another test that passed "George" as the name, right? Maybe
not. We have to consider the cost of that additional test, and this isn't always
simple to do. There are several costs we have to consider here:

  * direct cost of developer time
  * opportunity cost of developer time
  * long-term maintenance cost

The first is pretty simple, if the developer earns some amount of money `N` per
hour, and a test took `M` hours to write then the test cost `N * M`. If we know
about how much a test will cost, we need to ask ourselves (preferably before we
write it) how much it is worth. A test that verifies extremely simple code may
not be worth very much, and we should be cognizant of this fact.

The second cost listed above might sound familiar if you've taken an economics
course. The idea behind "opportunity cost" is that if you choose to do a thing
you give up the opportunity to do something else, at least during the time you
spend doing the thing you chose to do.

So, as a developer, if I choose to add another test for a section of code I am
giving up the opportunity to spend that time adding an additional feature to my
software, or adding a test to another section of code, etc. So, again, when we
consider adding additional tests to try to achieve 100% coverage we need to ask
ourselves if doing so is really the most valuable use of our time.

Finally, direct and opportunity costs don't end once the code (or test) is
written. Software must be maintained, bugs must be fixed, features must be
added. Every line of code we add today will continue to incur costs into the
future. So when we "chase" 100% coverage we might consider the costs acceptable
today, but if we fail to consider the costs we (or our organizations) will pay
tomorrow, we may still end up overpaying for test coverage.

## Further Reading

  * [Lecture slides](https://docs.google.com/presentation/d/1Gn8tg564q0imQb86PV8iMo7hQgMusLK17zgj8De9_m4/edit#slide=id.p)

