FROM debian:bullseye-slim as build

ARG COMMIT=b5954caff90df8dcd35e45e103ddd2f8ad99208b

RUN apt-get update && apt-get -y install \
        git \
        golang \
    && rm -rf /var/lib/apt/lists/*
RUN git clone https://github.com/commander-cli/commander.git \
    && cd commander \
    && git checkout ${COMMIT} \
    && go build cmd/commander/* \
    && ls

FROM debian:bullseye-slim

COPY --from=build /commander/commander /usr/local/bin/

VOLUME /data
WORKDIR /data

