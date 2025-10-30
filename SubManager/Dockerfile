FROM golang:1.25 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./cmd ./cmd
COPY ./internal ./internal
COPY ./migrations ./migrations
COPY ./docs ./docs
COPY ./entrypoint.sh .

ADD https://github.com/pressly/goose/releases/download/v3.14.0/goose_linux_x86_64 ./bin/goose

RUN chmod +x ./entrypoint.sh
RUN chmod +x ./bin/goose

RUN go build -o subManager ./cmd/app

FROM gcr.io/distroless/base-debian12:debug

WORKDIR /app

COPY --from=builder /app/subManager ./
COPY --from=builder /app/migrations ./migrations
COPY --from=builder /app/bin/goose ./goose
COPY --from=builder /app/entrypoint.sh ./entrypoint.sh

CMD ["/app/entrypoint.sh"]