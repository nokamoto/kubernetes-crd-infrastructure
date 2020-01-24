FROM golang:1.13.6-alpine3.11 AS build

RUN apk update && apk add git

WORKDIR /src

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

WORKDIR /src/api

COPY api .

WORKDIR /src

RUN go install ./api/server

FROM alpine:3.11

RUN apk update && apk add --no-cache ca-certificates

COPY --from=build /go/bin/server /usr/local/bin/server

ENTRYPOINT [ "server" ]
