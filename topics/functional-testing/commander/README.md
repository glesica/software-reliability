# Commander Examples

Use the Docker [image](Dockerfile) to create a container suitable for running
Commander tests. Take a look at [build.sh](build.sh) for a command to build the
image. The script [commander.sh](commander.sh) is a useful helper. You can treat
this script as though it were the Commander executable itself.

For example, to run the test file ([commander.yaml](commander.yaml)):

```output
./commander.sh test
```

Note, however, that this won't work until you've built the container image.

