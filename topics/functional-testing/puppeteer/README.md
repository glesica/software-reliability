# Puppeteer Example

The script [umt.edu.js](umt.edu.js) in this directory uses the
[Puppeteer](https://pptr.dev) library to automate verification of some of the
main navigation links on the [UM website](https://umt.edu).

To do this, it loads the page in a browser and then finds the links it is
supposed to click, clicks on each one, and, finally, verifies that the main
heading on the resulting page is correct.

You must have [Node.js](https://nodejs.org/en/) installed on your machine. We're
not going to use a Docker image for this, because Puppeteer has the option to
display the browser while it runs so that you can watch what is happening and
Docker doesn't work well for graphical software.

To run the script, you can use the simple command below:

```output
node umt.edu.js
```

If the program exits successfully, the tests passed.

