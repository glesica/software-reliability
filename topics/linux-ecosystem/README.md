# The Linux Ecosystem

We're going to be using Linux for this course and, in particular, the Linux
command line. Many of the tools we will use are also available for Windows in
various forms, and macOS. In many cases, in fact, the Linux commands we use will
work verbatim on a Mac.

While there is nothing wrong with developing software in a Windows environment,
for the sake of this course it is more efficient if we all use the same basic
tools. Additionally, Unix and Linux environments are incredibly common in modern
software development (most of the web runs on them) and many of the concepts are
transferable to modern Windows.

## Directory Structure

Unix / Linux systems arrange their files quite differently from Windows systems.
In Windows, the top-most node in the file and directory hierarchy is the disk
drive (`C`, `D`, and so on) and the directories are arranged into a forest of trees.
There is a hierarchy under each drive and the drives themselves do not share a
common parent, although you might think of "My Computer" or "Computer" (in more
recent versions of Windows) as the shared parent.

On the other hand, the Unix / Linux directory structure looks something like
this. The path to the top level is `/`, and it does not correspond to any
particular disk drive. In fact, any disk drive can be "mounted" at any point
under `/` in the hierarchy. For example, it is common to store user files
("home") on their own partition or drive. It is also common for external storage
devices to appear under the "mnt" directory. This might seem confusing at first,
but it is actually quite nice.

There are two special directory names. The first is `.` (a single period), which
refers to the current working directory. The second is `..` (two periods) which
refers to the directory above the current working directory or above its current
position in a path. For example, the path `/home/george/..` is equivalent to the
path `/home/`.

## Permissions

File permissions, that is, who can access which files, are often confusing for
new users of Unix / Linux systems. Historically, consumer Windows systems
haven't had much in the way of file permissions. NTFS (the filesystem used by
modern versions of Windows) changed this, but most users still don't have to
worry much about file permissions because Windows includes a feature called User
Account Control (UAC) that largely abstracts permissions away from the user.

Modern Unix and Linux systems also generally allow the user to ignore file
permissions, but there are times when they become important, so they are worth
understanding, at least at a basic level. You can view permissions for a file
using the "ls" command (see below) like so:

```
$ ls -l
total 12K
-rw-r--r-- 1 george george 244 Aug  4 18:50 README.md
-rw-r--r-- 1 george george 344 Aug  4 19:24 tags
-rwxr-xr-x 1 george george 703 Aug  4 19:24 test-bc.bats
```

The first part of permissions is the concept of ownership. Each file has an
owner, a user who is associated with that file. Additionally, each file has a
group. A group can have one or more users as members. Groups allow roles to be
assigned to different users. For example, a system might have a "web" group for
users who are allowed to modify files served by a web server. Those files would
then be assigned to the "web" group. In the example above, both the owner and
group are "george". Each user has their own group default and files in a user's
home directory have this group assigned to them.

There are three permission types: read, write, and execute. The "read"
permission allows the contents of a file to be inspected. The "write" permission
allows a file to be written to or deleted. The "execute" permission allows a
file to be run as a program. The first two are probably pretty clear, but why a
separate permission for executing a file? It's partly for security, so that you
can't run something that isn't intended to be run and might have negative
consequences if it were to be run. Additionally, whether a file is executable or
not can determine what happens when you, say, double-click on it in a file
manager.

Each of these permission types can be applied to one or more of three different
scopes: owner, group, and other. The first determines what the owner of the file
can do, the second determines what members of the file's group can do, and the
third determines what everyone else can do. In the example above, the owner can
read ("r") and write ("w") to all three files and execute ("x") the third file.
The group can read all three and execute the third, and others can do the same.

Further reading:

  * https://en.wikipedia.org/wiki/File_system_permissions#Traditional_Unix_permissions
  * https://kb.iu.edu/d/abdb
  * https://help.unc.edu/help/how-to-use-unix-and-linux-file-permissions/

## Basic Command Line Usage

A command line is a user interface metaphor, just like a window or a button. The
user types a command, and the computer executes it and prints out the results
and any error messages. The handy thing about a command line interface is that
common sequences of commands can be strung together into scripts for
convenience. Another nice thing is that command line interfaces are largely
unambiguous, a complex command can be copy-pasted or sent via email very easily.

There is nothing that intrinsically connects command line interfaces to Unix /
Linux. Nor is there any reason that a command line can only be used for certain
system-related tasks. In fact, many popular Windows applications have their own,
specialized, command line interfaces. Some examples include the popular design
application AutoCAD, and MS SQL Server, and of course Windows has always had an
emulated DOS prompt, though most people don't use it very often. However, it is
very common for Unix / Linux users to work at a command line to this day.

### Getting Started

First, open up a terminal. On a Mac, there is an application called "Terminal"
installed by default, but you might want to try iTerm2, a popular alternative,
instead. On most Linux distributions there is a program called "Terminal",
although it might also be called "Konsole" or "Gnome Terminal". These
applications don't provide a command line by themselves. Instead, they run, and
provide an interface for, a program called a "shell" that actually accepts
commands and prints output and errors.

There are many shells, some old and some new. Your terminal probably uses a
shell called "Bash". You'll find that many shells have names that end in "sh",
which is an abbreviation for "shell". Some others are Zsh, KornShell, C Shell,
and Fish (which is quite different from the rest). I suggest that you stick with
Bash for now.

Below, you can see what my shell looks like when it first starts. It is common
for shells to display a "$" character just before the user's input, this is
called the "prompt". It is customary for this prompt to use the "#" character
for the root (super) user. This reminds you to be careful when logged in as
root. My prompt is heavily customized because I spend a lot of time in a
terminal, this is pretty common among software developers and the like.

### Files and Directories

Using our shell we can manipulate files and directories. Type `pwd` (which
stands for "print working directory") and press the Enter key. This is an
example of a "command" (hence "command line"). This will print the path to the
current directory, called the "working directory".

Next, it would be helpful to see a list of files and directories inside of the
current directory. To do this, use the `ls` (short for `list`) command. If you'd
like to see more detail about the contents of the current directory, try `ls -l`.
The `-l` is called a `flag` or an `option`. These are used to change the
behavior of commands and they vary from command to command.

If you need to know which options are available for a given command, you can
generally use the `-h` or `--help` options to find out, most commands support
these. Let's try it with the `ls` command: `ls --help` (on my machine, this
command accepts the latter option). Here's the output (I've truncated it a bit
since it's so long):

```
Usage: ls [OPTION]... [FILE]...
List information about the FILEs (the current directory by default).
Sort entries alphabetically if none of -cftuvSUX nor --sort is specified.

Mandatory arguments to long options are mandatory for short options too.
  -a, --all                  do not ignore entries starting with .

...
```

The first line tells us how to use this command. In this case, we type "ls",
followed by zero or more options (flags) and zero or more files or directories.
The flags available to us are then listed, the "-a" flag stands for "all" and
shows all files, even hidden files (those preceded by a period).

Here are a few more file and directory commands you might find useful. We'll
talk about each of these in class, and you can always search for them online if
you want to learn more.

  * `cd` - change the current working directory to the current user's home directory
  * `cd foo` - change the current working directory to the directory called
    `foo` in the current working directory
  * `mkdir foo` - creates a new directory called `foo` in the current working directory
  * `rmdir foo` - removes the directory called `foo` (assuming it is empty) from
    the current working directory
  * `touch foo.txt` - creates an empty file called `foo.txt` in the current working directory
  * `rm foo.txt` - removes the file called `foo.txt` from the current working directory
  * `cat foo.txt` - print out the contents of the file called `foo.txt` in the
    current working directory (works best if it is text)
  * `cp foo.txt bar.txt` - copy the file `foo.txt` to `bar.txt`
  * `mv foo.txt bar.txt` - move the file `foot.txt` to `bar.txt` (rename it, essentially)
  * `chmod +x foo` - make `foo` executable for all users (common when writing a script)

### Command Completion

It might occur to you that typing all these commands would get tedious pretty
quickly. For this reason, most shells support command completion. For example,
type `cd ` (note the space) and then press the "Tab" key on your keyboard, you
should see a list of all available directories within the current working
directory. Type the first couple characters of one of them and then hit "Tab"
again, the directory name should either complete automatically (if it is
unambiguous) or you should see a shorter list of possible directories, those
that start with what you have typed. This same pattern works for commands
themselves, many command line options and flags, and many arguments to various
commands.

### System Information

We can also use commands to find information about our computers. For example,
the `ps` command lists processes (programs) running on your computer. There are
numerous options to control which programs are shown. One example is `ps -e`,
which displays all processes current running on the machine.

Another fun command is `free -m` which shows information about your computer's
memory (RAM). The `-m` flag tells the program to display memory in megabytes,
which is easier to read than the default bytes. Try running `free` by itself to
see the difference.

Last, try `df -h`. This command provides information about the amount of space
available on your disk drives. The `-h` option instructs the command to present
the information in a human-readable format (a little like `-m` above).

## Advanced Command Line Usage

### Searching and Finding

Finding specific files can be challenging for beginners. There are two tools
that can help you with this: `grep` and `find`. The first searches for text
within one or more files, the second searches for files by name or type (among
other things). Some examples will probably be most useful. Let's assume we have
a directory that contains two files, `file0.py` and `file1.py`, and a
sub-directory, `subdir`. The sub-directory, in turn, contains two files,
`subfile0.py` and `subfile1.py`.

```
.
├── file0.py
├── file1.py
└── subdir
    ├── subfile0.py
    └── subfile1.py
```

The files above contain some Python code:

```
$ cat file0.py 
def foo():
    print('foo')

$ cat file1.py 
def bar():
    print('bar')

$ cat subdir/subfile0.py 
def foo():
    print('sub foo')

$ cat subdir/subfile1.py 
def bar():
    print('sub bar')
```

What if we want to find all the files that contain "foo"? We can use "grep":

```
$ grep -r foo .
./file0.py:def foo():
./file0.py:    print('foo')
./subdir/subfile0.py:def foo():
./subdir/subfile0.py:    print('sub foo')
```

The `-r` flag tells grep to search recursively, which causes it to descend into
the sub-directory `subdir`. The general form for this command is `grep TARGET
LOCATIONS` to search for the string given by TARGET in the files or directories
given by LOCATIONS. We can also include the `-n` flag to include the line number
of each match in the output:

```
$ grep -rn foo .
./file0.py:1:def foo():
./file0.py:2:    print('foo')
./subdir/subfile0.py:1:def foo():
./subdir/subfile0.py:2:    print('sub foo')
```

Next, what if we want to search by filename? This is where the `find` command
comes in handy. I don't use this one very often, personally, and I always forget
how it works. I suggest making yourself a cheat sheet of commands you forget, or
searching online for a comprehensive one.

```
$ find . -name 'sub*'
./subdir
./subdir/subfile0.py
./subdir/subfile1.py
```

Let's break this down a bit. First, note that `grep` took the search term first,
then the locations to look in. This one is the other way around, the location to
look in comes first ("."), and the search terms come next. In this case I'm
looking for any file (or directory, directories are just weird files) with a
name that matches a particular pattern. In this case the pattern is "sub"
followed by anything else. I have to surround the search expression with single
quotes to prevent the shell from interpreting the "*" character. Always use
single quotes with find and you won't have to worry about it. But what if we
wanted only files, not directories?

```
$ find . -name 'sub*' -type f
./subdir/subfile0.py
./subdir/subfile1.py
```

We can use the `-type` filter to find only regular files (which is what the `f`
stands for). We could replace the `f` with a `d` to find only directories.

### Pipes and Redirection

One of the most powerful features of the Unix / Linux command line interface is
the ability to compose small commands into larger functions. This idea is part
of what is generally referred to as the Unix Philosophy. Two features present in
most shells support this kind of composition. The first is pipes. A pipe is
basically just the ability to feed the output of one command directly into the
input of another command. A simple and useful example is using the `grep`
command to filter the results of the `ls` command. This ends up being similar to
the behavior of the `find` command, and I do this sometimes when I can't
remember how to use `find` off the top of my head:

```
$ ls | grep usb
usb_modeswitch.conf
usb_modeswitch.d
```

I hopped over to the `/etc/` directory because it's got a lot of files in it.
Running "ls" results in 243 files and directories. What if I wanted to find out
which ones have the string "usb" in their names? This is what I did above. The
"|" character (which is pronounced "pipe", conveniently) instructs the shell to
take the output of the command on the left and feed it into the command on the
right as its input. Notice that I didn't give `grep` any locations to search,
that's because it will automatically search through whatever text it receives
through the pipe.

Let's look at another example. How did I know how many files and directories
were in the `/etc/` directory? I'm sure there are dozens of ways to discover
this information, but the easiest, in my opinion, is to pipe the output of `ls`
through a program called `wc` (word count) which can count characters, words,
and lines in text. The output of `ls` is passed through the pipe one file per
line, so by counting the lines we can determine how many files there are:

```
$ ls | wc -l
243
```

The `-l` flag on the `wc` command tells it to provide only a line count. This
pattern is available for many, many commands. I suggest seeing if you can come
up with your own combinations. Also note that you can use more than one pipe in
one command. For example, there is a silly program installed on many systems
called `cowsay` that draws a little cow on the screen and makes it "say" (with a
word bubble) whatever was piped into the command. So the command below actually
uses two pipes, one to get the line count (same as above) and one to send the
resulting output to `cowsay`:

```
$ ls | wc -l | cowsay
 _____
< 243 >
 -----
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||
```

## Version Control With Git

I'm not going to go into much detail here because there are hundreds of good Git
tutorials online. Instead, I'm going to try to explain why we would want to use
a version control system and link to a couple good tutorials.

### Why Use Version Control

For better or worse, programming is an experimental process. We change our code,
then we try to verify that it works (and possibly validate that it is what the
customer wants). This process is difficult to get right and, at its core, this
is why we use version control. Version control lets us manage changes to our
source code (which is particularly useful when we work in large teams), and it
gives us a way to undo what we've done when we inevitably mess up.

Version control gives us a way to organize and manage changes to our code. The
most naive way to share code might be to email "versions" back and forth. Many
people do this with Word documents in school. Have you ever seen someone name
files using a pattern like "report1.docx", "report2.docx", "report final.docx",
"report_final2.docx", and so on? This gets confusing pretty quickly. It is even
worse when multiple people are involved. How do I know if "report2.docx" is the
one I created, or the one you created and emailed to me?

Naive versioning schemes like this one also make it very difficult to merge two
versions of a document. If you and I both modify the same line of text, how do
we combine them? Should we keep both my version and your version, or should we
choose one? What if I miss one of your changes while merging them into my copy
of the document, will the combined document still make sense? Version control
systems help us solve this problem.

Git and other version control systems track versions of our source code and
allow us to combine versions automatically. If there are differences between
versions that can't be combined automatically they are clearly labeled so that
we don't miss any changes and we can combine them ourselves.

Version control also allows us to revert our changes easily and without losing
track of what we had added since the version we revert to. Have you ever changed
a bunch of code and then realized that your changes didn't make as much sense as
you thought they would when you started? It's annoying. In the professional
world it can also cost money if your customers run into bugs you didn't notice
until too late.

The ability to revert our code to some past state can help in both these cases
and version control systems are what allow us to revert in an orderly manner.

### Tutorials and Resources

  * https://blog.udemy.com/git-tutorial-a-comprehensive-guide/ - a detailed
    guide, all in one page. Useful as a reference as well.
  * https://git-scm.com/book/en/v1/Getting-Started - the official tutorial /
    book. This one has excellent depth, though you might find yourself skipping
    certain sections, that's fine.
  * https://www.udacity.com/course/how-to-use-git-and-github--ud775 - a Udacity
    course that introduces both Git and GitHub.
  * https://www.udacity.com/course/version-control-with-git--ud123 - another,
    more introductory, Udacity course.

## Suggested Editors

I personally use JetBrains IDEs for most languages and projects. They are
full-featured, professional quality tools with support for code intelligence,
refactoring, and testing. There are free versions for a couple languages, and as
a student you can get free academic licenses for the rest. I also use Vim for
quick edits and some programming.

I don't require that you use what I use, but if you do I'll be able to help you
troubleshoot more easily. That being said, here are some other popular editors:

  * https://code.visualstudio.com/ - a powerful editor available for free from
    Microsoft that runs on Windows, Mac, and Linux.
  * https://atom.io/ - a similarly powerful editor available for free from
    GitHub that runs on all major platforms.

I do not personally recommend Eclipse because I don't like it. However, if you
are comfortable with it then by all means, use it. The choice of tools is a
highly personal one, and it is important to use what you are comfortable and
productive with rather than what you think someone else wants you to use for one
reason or another. That being said, you should also take the time to try new
things occasionally, you never know what you might be missing.

