FROM golang:1.15-buster

ENV GIN_MODE=release

WORKDIR /api
COPY . /api

RUN make build

EXPOSE 3000

ENTRYPOINT ["./kubernetes-hello-world-api"]