# Logging

When we deploy software it can feel a little like releasing a balloon into the
air, it's gone forever. Sometimes, if we're very lucky, users will tell us when
things break. We might get bug reports or, more likely, negative reviews or
calls to customer support (depending on the type of product). But more often
than not, if there is something wrong with our software, users will simply find
another way to do whatever it is they need, or want, to do.

Even when users try to be helpful, they are almost always unfamiliar with the
inner workings of our software. In most cases, a bug reports saying that the
software "crashed" is unhelpful. Unfortunately, though, that is usually about
all we can expect from an average user.

So how do we solve this problem?

Obviously, our software needs to be able to tell us what's happening on its own.
At the very least, we need to be able to glean additional information from a
sparse user report. At best, we shouldn't need the user reports at all.

## Log Files

The simplest way to track what's going on with our software is logging. And the
simplest way to collect logs is with log files. Think of this as "print"
debugging after a good night's sleep and a healthy breakfast.

We will look at logging to files using Python.

## Telemetry

Telemetry is data collected by software and transmitted to a server. Think of it
as remote logging, though it's a bit more than that. Generally, telemetry
involves user-facing code, since we're interested in not just how data are
flowing through the system, but what the user is actually doing, how long is
takes, and so on.

