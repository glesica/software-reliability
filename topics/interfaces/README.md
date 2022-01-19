# Interfaces

Anyone who has been exposed to Java should have some familiarity with
interfaces. Put simply, an interface is a description of the methods that must
be available on a class instance that implements the interface. I might help to
think about an interface as a description of the "shape" of an object.

```java
public interface Sender {
  void send(String message);
}
```

But how do interfaces help us with testing? Largely by letting us create simple
test doubles more easily. This is going to look very similar to dependency
injection, and that's because it is. Interfaces are one way to enable dependency
injection. First, let's look at a class that we might use to talk to a chat
service, like Slack or Discourse.

```java
public class Client {
  void send(String message) {
    // verify server connection
    // send message
    // check for errors
  }
}
```

Looks pretty reasonable. Now let's say we have bot that interacts with our chat
client in one way or another (maybe it tells a joke whenever someone says "tell
me a joke" in chat).

```java
public class JokeBot {
  private Client client;

  public JokeBot() {
    this.client = new Client();
  }

  public void tellJoke() {
    String nextJoke = "i just flew in, boy are my arms tired";
    client.send(nextJoke);
  }
}
```

This looks great. But now how do we test our JokeBot class? We can't just create
a new instance and use it, because it will create a client, which will attempt
to connect to a chat server (presumably). We can fix this problem using
dependency injection.

```java
public class JokeBot {
  private Client client;

  public JokeBot(Client client) {
    this.client = client;
  }

  public void tellJoke() {
    String nextJoke = "i just flew in, boy are my arms tired";
    client.send(nextJoke);
  }
}
```

OK, so now the JokeBot constructor doesn't attempt to create a new Client, but
we still have to give it a Client to use, and if we try to create one we'll end
up in the same place we were a minute ago because the Client will try to talk to
a chat server. Enter interfaces...

```java
public interface Client {
  void send(String message);
}
```

We change Client from a concrete class to an interface. Now we can rename our
previous Client class to be more specific and implement Client:

```java
public class NetworkClient implements Client {
  void send(String message) {
    // ...
  }
}
```

Finally, we can define another Client implementation that doesn't try to connect
to a server but instead simply tracks the messages it has been asked to send
(just like in the dependency injection example):

```java
public class TestClient implements Client {
  String lastMessage;

  void send(String message) {
    lastMessage = message;
  }
}
```

Now we can use TestClient for our JokeBot tests, which allows us to make sure
that JokeBot sends the jokes we expect it to send, when we expect it to send
them, but without having to connect to a chat server or use a real chat room.

## Further Reading

  * <https://docs.microsoft.com/en-us/dotnet/csharp/programming-guide/interfaces/>
  * <https://www.codeproject.com/Articles/702246/Program-to-Interface-not-Implementation-Beginners>

