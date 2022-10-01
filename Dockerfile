## Stage 1 - build
FROM golang:1.19-alpine3.16 AS builder

WORKDIR /build

COPY main.go go.mod go.sum ./

RUN go mod download
RUN go build main.go


## Stage 2 - run
FROM alpine:3.16

LABEL maintainer='Ilias Mavridis'

RUN adduser -S -D -H -u 12222 -h /app appuser
USER appuser

WORKDIR /app

COPY --from=builder /build/main .

ENTRYPOINT ["./main"]