# Factory Pattern

A factory is a place where similar products are produced in mass quantities. For
example, a car factory might produce one model of car, but each car produced may
have different features. Some might have leather seats, others cloth. They may
come in different colors. Some of their engines might have six cylinders, others
four. Or, a particular factory might be able to produce both car model X and car
model Y. In this case, the factory might switch between the two based on the
prevailing situation (demand, corporate strategy, and so on).

In software, a "factory" is similar, though it produces instances of a class or
other data structure instead of physical objects. Let's look at an example.

```java
public abstract class Character {
  public void move(int direction);
  public void speak();

  public static Character makeCharacter(int level) {
    if (level % 2 == 0) {
      return new Mario();
    } else {
      return new Luigi();
    }
  }
}

public class Mario extends Character { ... }

public class Luigi extends Character { ... }
```

In this case, we want to use Mario for even-numbered levels, and Luigi for
odd-numbered levels (for whatever reason). So depending on which level we
specify, we get back either a Mario or a Luigi transparently from the factory.
Note that this probably seems similar to Dependency Injection and, in fact, in
many cases we treat a factory as a dependency, then we can choose which factory
to provide to a class or method based on some other criteria. For example:

```java
public abstract class Character {
  // ...

  public static Character makeEasyCharacter() {
    return new Mario();
  }

  public static Character makeDifficultCharacter() {
    return new Luigi();
  }
}
```

In this case, we want to supply a different character depending on the
difficulty level chosen by the player. But our game logic doesn't need to care
about the difficulty, or which character is assigned to which difficulty level.
Instead, we simply provide the game logic with the correct factory (through
injection) and everything works as it should.

This pattern is often used to provide a different implementation of something
like a database connection depending on the environment in which the program is
running (including a testing environment).

