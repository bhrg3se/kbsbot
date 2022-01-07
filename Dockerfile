FROM golang:1.16 as builder

WORKDIR /go/src/kbsbot

COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

ENV CGO_ENABLED=0
RUN go build -o kbsbot

FROM alpine:latest
WORKDIR /kbsbot

COPY --from=builder /go/src/kbsbot/kbsbot .
COPY setup/config.toml /etc/kbsbot/config.toml

CMD ["./kbsbot"]
