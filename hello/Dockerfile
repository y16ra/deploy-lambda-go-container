ARG GO_VERSION=1.21

# First stage: build the executable.
FROM golang:${GO_VERSION}-alpine AS builder

WORKDIR /src

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Import the code from the context.
COPY ./ ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /app ./hello/main.go

