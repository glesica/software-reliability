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

See the sequence diagram below for an illustration of how a message is created.
This is the process that students will test as a homework assignment.

![Creating a secret message](create-message.png)

The process for viewing a message is show below. There is an example test for
this process in the `tests/` directory.

![Viewing a secret message](view-message.png)

