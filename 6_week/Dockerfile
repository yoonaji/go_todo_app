# Build stage
FROM golang:1.23-bullseye as deploy-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -trimpath -ldflags "-w -s" -o app

# Final stage
FROM debian:bullseye-slim as deploy
RUN apt-get update
COPY --from=deploy-builder /app/app .
CMD ["./app"]

# Development stage
# Development stage
FROM golang:1.23 as dev
WORKDIR /app
RUN go install github.com/air-verse/air@latest
CMD ["air"]

