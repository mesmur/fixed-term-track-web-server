ARG GO_VERSION=1.22.5

FROM golang:${GO_VERSION}-bookworm AS builder

LABEL author="Maahir <msmaahirur@gmail.com>"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . ./

RUN go build -tags netgo -ldflags '-s -w' -o app ./cmd

FROM debian:bookworm

WORKDIR /app

# https://github.com/google/go-github/issues/1049
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --chown=app:app --from=builder /app/app ./app

EXPOSE 3000

CMD ["/app/app"]
