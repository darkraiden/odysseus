FROM golang:alpine3.11 AS builder

ENV GO111MODULE=on

WORKDIR /go/src/github.com/darkraiden/odysseus

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o odysseus ./cmd/odysseus/main.go

FROM alpine:3.11.3

RUN adduser -S -D -H -h /app -u 1001 odysseus

USER odysseus

COPY --from=builder /go/src/github.com/darkraiden/odysseus/odysseus /app/odysseus

WORKDIR /app

CMD ["./odysseus"]
