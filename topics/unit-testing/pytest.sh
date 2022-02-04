#!/usr/bin/env sh

set -e

docker run -u "$(id -u):$(id -g)" -v "$PWD":/data glesica/pytest pytest "$@"

