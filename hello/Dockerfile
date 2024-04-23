ARG GO_VERSION=1.21

# First stage: build the executable.
FROM golang:${GO_VERSION}-alpine AS builder

WORKDIR /src

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Import the code from the context.
COPY ./ ./
RUN ls -la
# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /hello ./main.go

# Final stage: the running container.
FROM gcr.io/distroless/static-debian12:nonroot as final

COPY --from=builder /hello /hello

ENTRYPOINT [ "/hello" ]
