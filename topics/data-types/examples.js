#!/usr/bin/env node

console.log("PYTHAGORAS");

function pythagoras(a, b) {
    console.log(Math.sqrt((a * a) + (b * b)));
}

pythagoras(2, 3);
pythagoras("2", 3);
pythagoras(true + true, 3);
pythagoras("shoe", 3);
pythagoras(null, 3);

console.log("\nREPEAT");

function repeat(msg, count) {
    var output = null;
    for (i = 0; i < count; i++) {
        if (output === null) {
            output = "";
        } else {
            output += " ";
        }
        output += msg;
    }
    console.log(output);
}

repeat("hello", 3);
repeat("hello", "3");
repeat("hello", true + true + true);
repeat("hello", "duck");
repeat(null, 3);
