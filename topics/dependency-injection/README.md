# Dependency Injection

  * [Homework](assignment/)

When a function or class requires access to another function or class instance
in order to do its work, we say that the latter is a "dependency" of the former.
For example, the function below depends on the `print` function.

```python
def send_message(message):
  print("Message: " + message)
```

This might seem like an unusual way to think about our code. Why does it matter
what we depend on? Considering our dependencies explicitly has a couple
potential benefits, but the most notable is probably that it can simplify
automated testing.

Let's say you wanted to test the function above. You could read from stdout to
make sure your message got printed, but that's clunky (and if this were a
network request or a database query, perhaps quite unreasonable).

Instead, we can allow a print function (which is really just any function that
accepts a string, in this case) to be "injected". Within our `send_message`
function, then, any time we want to print something, we use the injected print
function instead of the built-in version.

```python
def send_message(message, print_func):
  print_func("Message: " + message)
```

If we wanted to get fancy we could give `print_func` a default value,
perhaps setting it to the standard print function by default.

```python
def send_message(message, print_func=print):
  print_func("Message: " + message)
```

Now, in our test we can give it a function that stores whatever arguments it is
called with so that we can verify that they are correct without having to mess
around with stdout.

```python
class MessageTrap:
  def __init__():
    self.messages = []

  def __call__(self, message):
    self.messages.add(message)

def test_send_message():
  trap = MessageTrap()
  send_message("hello there", trap)
  assert len(trap.messages) == 1
  assert "hello there" in trap.messages[0]
  # verify trap.messages
```

There are other benefits to treating the print function as a dependency as well.
For instance, we might want to wrap or format the output in some way, depending
on who is running the program and in what environment. This is a common use-case
for message logging, where we want to include things like the date and time of a
message alongside the content of the message. We might also want to publish the
message to some kind of network service or even send the message over email,
perhaps depending on its content.

## Advanced Example

TODO

## DI Frameworks

TODO

## Further Reading

  * https://www.youtube.com/watch?v=IKD2-MAkXyQ
  * https://stackoverflow.com/questions/130794/what-is-dependency-injection
  * https://en.wikipedia.org/wiki/Dependency_injection
