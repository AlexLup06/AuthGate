FROM golang:1.25-alpine as base

RUN apk add --no-cache postgresql-client ca-certificates
RUN go install github.com/rubenv/sql-migrate/...@latest

WORKDIR /migrations

COPY . .

ENTRYPOINT ["sql-migrate"]
