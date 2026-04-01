FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /app/homelab

FROM alpine:latest AS runtime

RUN apk add --no-cache ca-certificates

COPY --from=0 /app/homelab /app/homelab

CMD ["/app/homelab"]
