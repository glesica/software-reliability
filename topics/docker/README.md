# Docker

Docker is a Linux container runtime. A container is a little like a virtual
machine in the sense that it allows you to run one version of Linux (complete
with various software installed) inside of a (potentially different) version of
Linux.

This comes in handy for many reasons, but for the purposes of this course we
will use Docker to ensure we all have the same versions of the code and tools we
use for assignments.

  * [Docker Desktop](https://www.docker.com/products/docker-desktop)

If you are using a Linux distribution, you can probably install Docker with a
command like `sudo apt-get install docker.io` (Debian or Ubuntu). On a Mac or
Windows computer, you can install "Docker Desktop" from the link above. This may
require you to sign up for a free account.

We can use Docker from the command line. Try the command below to drop into a
minimal Debian installation.

```output
$ docker run -it debian:bullseye-slim bash
```

Use `ls` and `cd` to look around and convince yourself that you're using a
totally isolated copy of Linux. You can install software inside the container
with `apt-get`, as you would if you had Debian installed directly on your
laptop. The only difference is that the commands you run won't have any
permanent effect on your computer.

You might be wondering, then, what good is this if I have to reinstall all of my
software every time I use a container? Great question! As it turns out, we can
customize the container before we even run it by creating what is known as a
"container image" (or "Docker image", colloquially). To do this, we use a
Dockerfile.

## Dockerfiles

  * [Dockerfile Reference](https://docs.docker.com/engine/reference/builder/)

You can think of a Dockerfile like a small program or script that defines what
the container image will contain when it is run. It does this mostly by copying
files into the container and running shell commands.

A container image, defined by a Dockerfile, can be based on another container
image. This lets us, for example, start with a basic Debian system and then use
its package manager and other tools in our customization. For example, the
Dockerfile below simply installs a single Debian package, which is then
available in the resulting image (it was unavailable in the original Debian
image).

```dockerfile
FROM debian:bullseye-slim
RUN apt-get update && apt-get install -y cowsay
ENV PATH=${PATH}:/usr/games
```

We build the container image, passing along a "tag" (basically a name) with the
`-t` flag. In this case, since the Dockerfile isn't called `Dockerfile`, we also
need to pass a path to the file using the `-f` flag.

The final argument (`.`) is the "build context". This is basically the directory
to make available for copying files into the container. We use the current
directory in this case, which is common.

```output
$ docker build -f Dockerfile_simple -t docker-simple .
```

Now we can run the container, passing along a command to be run inside of the
container.

```output
$ docker run docker-simple cowsay "hello there"
 _____________
< hello there >
 -------------
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||
```

Notice that, at least in this case, we didn't drop into an interactive shell
within the container. In most cases this is how we want to interact with
containers: by using them almost like applications installed on our machines as
opposed to separate operating systems.

In fact, it is common to create a shell alias to make running a container
simpler. For instance, run:

```output
alias containersay=docker run dockerfile-simple cowsay
```

Now we can run something like `containersay "hello there"` and get the same
result as above.

## Use in Class

We will talk more about how to use the Docker images distributed for class when
the time comes. For now, feel free to explore Docker further and think about how
you might make use of it for your other classes.

