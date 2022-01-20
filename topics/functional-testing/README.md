# Functional Testing

When we sit down to a build a piece of software, how do we know what to build?
In most cases, we know what to build because we've gathered a set of
requirements, characteristics our software must have in order to be considered
"complete" and "correct". The process of gathering requirements can be as
detailed as we want to make it.

Sometimes, if we're building software for ourselves, we may skip this step
since the requirements are already in our heads (or we might quickly jot them
down so we don't forget).

In other cases it may take weeks or even months to identify what our software
must do, particularly when we are unfamiliar with the problem to be solved (the
domain). For example, I worked for a company that makes software for
accountants. I know very little about accounting, so when I was asked to build
some feature I had to ask a lot of questions first to make sure that what I
ended up with was useful to accountants.

## Requirements

There are two kinds of requirements, functional and non-functional. They are
described below. We will, however, focus on function requirements for now.

### Functional Requirements

Functional requirements deal with what the software in question does. You can
think of a functional requirement as a "feature" without much loss of accuracy.

Programming assignments in your CS course mostly specify functional
requirements. For example, your program must print "Hello, World" when run. Or,
"given a set of numbers, your program must print the numbers from the set that
are prime". In a larger application, a functional requirement would probably
include background information and additional detail, but the idea is the same:
functional requirements are concerned with what the program does.

### Non-functional Requirements

We will discuss non-functional requirements more later in the semester, but for
now you can think of these as requirements that deal with how the program does
what it does.

For example, performance is generally a non-functional requirement.
In the second example above, your assignment might specify that your program
must be able to identify all the prime numbers in a set of 100 values within
10ms, or that the algorithm must have O(n) performance for larger input sets.

These are non-functional requirements because they do not affect the answers or
output produced by the program. However, this doesn't make them less important,
non-functional requirements can be deal breakers in many cases (for instance,
performance is incredibly important for games).

## Functional Testing

Functional testing is probably the most obvious form of testing (doing it well
requires some subtlety, though). In fact, every single one of us probably
invented it independently the day we wrote our first program.

Think of the first program you wrote, what did you do after you wrote it? You
ran it, of course! And you probably verified that the output was what you
expected (likely "Hello, world"). That is functional testing in a nutshell.

A functional test attempts to verify one or more functional requirements of the
software in question. For example, a functional test for a "Hello, world"
program simply verifies that, when run, the software prints the string "Hello,
world".

Historically, functional testing was largely a manual process. Armies of
software testers would interact with the application under test, completing
predetermined sets of tasks and reporting bugs as they were encountered. While
manual functional testing is still used (and is likely to continue to be
necessary), automated functional testing has replaced it for simpler tasks.  In
some cases, manual testing has also been moved to the users since the internet
enables frequent bug fixes.

### Manual Functional Testing

Manual functional testing requires that the tester be provided with unambiguous
instructions that describe the task to be completed and the expected result,
this is known as a test specification or "spec" for short.

One benefit of manual testing is that these test specifications can often assume
a certain baseline level of knowledge, although doing so isn't always wise. For
example, a tester working on a word processor might be assumed to know what to
do when told to "open" a document. However, this can become problematic. Should
the tester use a keyboard shortcut, the File menu, or even a toolbar button?
What about just double-clicking on the file to be opened? In some cases it might
not matter, but in others the tester's choice could allow a bug to slip through.

A reasonable test specification for the `ls` command is shown below.

  1. Open a terminal in a directory that contains one sub-directory, one symlink
     to that sub-directory, and one regular file
  1. Run `ls`
  1. Observe the list of files printed on the screen and verify that it contains
     each of the three items described above

### Automated Functional Testing

Many functional tests can be automated to improve consistency and speed and
reduce costs. Unfortunately, automating functional tests can be somewhat complex
because the tests must be "driven" by a piece of software that controls the
computer in the same way that a human would. Furthermore, different drivers are
required for different interface types.

We will discuss two drivers. The first, called
[Commander](https://github.com/commander-cli/commander) is used to automate
command line applications. The second, called [Puppeteer](https://pptr.dev) is
widely used to test web applications using the Chrome or Chromium web browsers.

#### Commander

There are a number of ways to test command line applications. The simplest way
is just to capture the output by redirecting it to a file and compare that file
to a known "good" output. This style of testing is commonly known as "gold file"
testing, the known good output is the "gold" file.

This style of testing has drawbacks, though. Gold file tests can be brittle;
small changes to the program can still cause an entire gold file test to fail.
Additionally, the gold files themselves can be cumbersome to create and update.

We can also test program output in smaller chunks. We can do this manually, of
course, but it's easier if we use a tool to make it easier.

As an example, we will create a couple tests for a command line utility included
by default on most systems using
[Commander](https://github.com/commander-cli/commander), a handy tool for
testing command line applications. There is a trivial example below.

```output
tests:
  echo hello world:
    stdout: hello world
    exit-code: 0
```

We can run this with the following command, assuming the file above is saved as
`echo_test.yaml`:

```output
commander test echo_test.yaml
```

Check out the [example](commander/) we will use in class, complete with a Docker
image.

#### Puppeteer



## Non-functional Testing

We will discuss various forms of non-functional testing later in the semester.
These will include performance testing and various forms of security analysis.








Command Line

Web Browser
We can also functionally test web applications. Gold file tests are possible using images or raw HTML, but we can also automate the web browser itself in order to create more advanced tests that interact with and verify properties of individual HTML elements.
Puppeteer Examples
Puppeteer is a tool for automating a web browser. It is relatively recent and relies on a protocol that is currently only implemented in Google Chrome, but support is planned for all three major web browsers. As its name implies, Puppeteer allows us to interact with a web browser using code that can be assembled into scripts for repeatability.
We will use Puppeteer to verify some characteristics of some popular web sites as though we were developers testing them. Puppeteer is distributed as a JavaScript library and it can be installed using NPM.
You can find examples that use Puppeteer here:
https://github.com/um-software-testing/reddit-puppeteer
