FROM golang:1.26-alpine AS builder

WORKDIR /app

RUN go install github.com/a-h/templ/cmd/templ@latest && \
    go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN templ generate && \
    sqlc generate && \
    go build -o /app/homelab

FROM alpine:latest AS runtime

RUN apk add --no-cache ca-certificates

COPY --from=0 /app/homelab /app/homelab

CMD ["/app/homelab"]
