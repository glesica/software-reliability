#!/usr/bin/env sh

set -e

docker run -u $UID -v $PWD:/data glesica/commander commander "$@"

