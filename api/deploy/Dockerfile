FROM golang:1.14

ADD . /go/src/github.com/replicatedhq/kots-bots/api
WORKDIR /go/src/github.com/replicatedhq/kots-bots/api
RUN make build


FROM debian:stretch-slim
RUN apt-get update && apt-get install -y --no-install-recommends \
    curl ca-certificates git \
    && rm -rf /var/lib/apt/lists/*

COPY --from=0 /go/src/github.com/replicatedhq/kots-bots/api/bin/* /
COPY ./web/dist-kots /web/dist

EXPOSE 3000
