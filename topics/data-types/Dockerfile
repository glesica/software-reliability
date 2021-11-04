FROM ubuntu:21.10

RUN apt-get update \
    && apt-get install -y \
        build-essential \
        fsharp \
        mypy \
        node-typescript \
        nodejs \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

VOLUME /work
