# Integration Testing

Sometimes called an "end-to-end" test, an integration test verifies that two or
more separate components of a software system work correctly together. This kind
of testing can be used in two different situations:

  * integrating a custom component with off-the-shelf components that will exist
    in production, like a database
  * integrating two or more custom components with one another, and optionally
    off-the-shelf components

Let's say that we've written a web app in PHP, a common web scripting language.
It is very common to deploy PHP applications using a separate web server like
Apache or Nginx. The web server receives requests from the network, then runs
the correct PHP code for requests that match rules defined by the developer or
system administrator.

It is also common for web applications in general to rely on a database such as
MySQL. During development, however, developers often choose to use a simpler
database system such as SQLite that allows them to iterate more quickly (make a
change, manually test the change).

An end-to-end test is, at its most basic, an integration that includes an entire
"trip" through the system, from input to output, and whatever comes in between.

![Integration Testing](integration-test.jpg)

