#!/usr/bin/env bash

set -e

docker compose --env-file env.dev up --build
