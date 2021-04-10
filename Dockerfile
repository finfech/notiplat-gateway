FROM golang:1.15.11-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN apk add --no-cache git && \
    go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build cmd/email-gw/main.go cmd/email-gw/wire_gen.go

FROM scratch

COPY --from=builder /app/main /app/main

ENTRYPOINT [ "/app/main" ]