# Secret Messages Example

We're going to end-to-end test a system for sending encrypted messages between
users. A user creates a message, chooses a password for that message, and
submits them through a web interface. The system encrypts the message using the
given password, stores the encrypted message in a database, then displays a link
where anyone with the password can read the message.

**Important note:** this isn't a *good* system, it's just an example. This is
actually a pretty bad way to send secret messages, not least of all because the
password and the unencrypted message have to leave the user's computer. It would
be better to encrypt the message on the user's device and send only the
encrypted version to the server. But that wouldn't be as much fun to test!

## Tests

See the sequence diagram below for an illustration of how a message is created.
This is the process that students will test as a homework assignment.

![Creating a secret message](create-message.png)

The process for viewing a message is show below. There is an example test for
this process in the `tests/` directory. This test uses
[Puppeteer](https://pptr.dev) to automate the browser and
[Chai](https://www.chaijs.com) for assertions. See the documentation for more
information. You can use the existing test as a starting point for your own.

![Viewing a secret message](view-message.png)

To run a test script, you must have [NodeJS](http://nodejs.org) installed. Then,
use the `npm` tool to install dependencies: `npm install --include=dev`. Once
that is done, you can run a script with the `node` tool: `node read-message.js`.

## Services

Each of the boxes in the diagrams above represents a service. The left-most two
are custom, they run code that I've written. The one on the right, on the other
hand, is a third-party component (a database system). This mixture is common in
real-world applications.

As an aside, why would we want to structure an application like this? The
database is probably easy to understand, but what about the "Encryption
Service"? This is a small example of a common architectural pattern called
"micro-services". Put simply, this involves building applications out of small,
independent services that communicate using HTTP or some other network protocol.

The benefit, in this case, comes from the fact that the encryption service is
separate from the web server.  This means a couple things:

  1. the encryption service can be developed by a complete separate team
  2. the web server will stay up even if the encryption service crashes

The second point might seems a little dubious since, in our case, the
application is useless if the encryption service goes down. However, there are
plenty of examples where applications can maintain at least partial
functionality even if individual services go down. Further, even in our case, it
is nice that a bug in our encryption logic can't take down the whole application
since, this way, at least our users can still see that their existing messages
are safe, even if they can't be decrypted.

## Environments

There are two environments that our code can run in. The first is a developer's
computer (we might call this the "development" environment).

In this environment, most of the code runs on the same machine, we're not
worried about long-term data persistence (if your test messages get lost when
you restart, it's not a big deal), we don't particularly care about using "real"
encryption, and, ideally, we'd actually like for there to be some "fake" data in
the database when we start the whole thing up.

The other environment is commonly called "production", or "prod" for short. This
is where the application is available to users. In this environment we want
reliable data persistence, real encryption, and we don't need any "fake" data to
test with.

In order to control how our application works when we run it, we can use
environment variables. For example, when the `SUPERSECRET_FAKE_ENCRYPT`
variable is set, the application won't actually use the encryption service to
encrypt the data and will instead prepend messages with "encrypted: " so that
we can tell they have been through encryption for testing purposes.

A list of environment variables that we can use to control the application are
listed below.

  * `SUPERSECRET_FAKE_ENCRYPT` - use fake encryption when set to "true"
  * `SUPERSECRET_FAKE_DB` - use a fake, in-memory database when set to "true"

## Running

The most obvious way to run the application is to start the web and encryption
servers in separate terminals, run a MySQL database (perhaps in a third
terminal), and then open a web browser and point it at the appropriate URL.
In fact, this works fine. However, it would probably get pretty annoying after
awhile, and it doesn't really help us at all with the production environment.
