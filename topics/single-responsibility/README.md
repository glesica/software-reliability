# Single Responsibility Principle

This principle basically states that a given module, or other unit of code,
should do one thing in the context of the overall software package. For example,
if you have a class called SpellChecker, it should do one thing (probably check
spelling). It shouldn't participate in the layout of a document, or even perhaps
grammar checking. It probably also shouldn't be responsible for creating the
interface elements shown to the user during the spell checking process. This
idea is sometimes also called "cohesion".

This principle applies to software testing because a module that has a single
responsibility will often be easier to test, and easier to change in the future.
This is partially because dependencies are minimized, and partially because
simpler, more cohesive modules are less likely to change frequently. Once they
work they tend to continue to work, which is really our ultimate goal.

## Further Reading

  * https://en.wikipedia.org/wiki/Single_responsibility_principle
  * https://en.wikipedia.org/wiki/Cohesion_(computer_science)

