#!/usr/bin/env bash

set -e

docker run \
    --detach \
    --rm \
    --name supersecret-mariadb \
    --env MARIADB_USER=supersecret \
    --env MARIADB_PASSWORD=supersecret \
    --env MARIADB_ROOT_PASSWORD=supersecret \
    --env MARIADB_DATABASE=supersecret_db \
    -v "$PWD/fixtures":/docker-entrypoint-initdb.d \
    -p 3306:3306 \
    mariadb:10.7.3

