# Functional Testing Homework

There are three components to this assignment, but they will all be submitted
as a single assignment. Don't forget to complete all three!

## 1. Manual Functional Testing

Write a functional test procedure for three Linux commands chosen from the list
below. Your procedure should verify some aspect of the command's functionality
and must clearly specify any assumptions, such as files or directories that must
exist.

  * `cat`
  * `find`
  * `grep`
  * `ls`
  * `sort`
  * `tr`

## 2. Automated Command Line Testing

Using [Commander](https://github.com/commander-cli/commander) (discussed in
class), automate the manual tests you created above. All of the commands above
are already installed in the container image provided, so you don't have to
worry about installing them.

## 3. Automated Browser Testing

Choose a website and use Puppeteer to verify three things about it, such as page
headers, links, or other content. Your script should resemble the example we
discussed in class and you can even borrow the `assertEqual` function, if you'd
like.

