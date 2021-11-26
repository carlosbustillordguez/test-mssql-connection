# syntax=docker/dockerfile:1

##
## Builder Image
##
FROM golang:1.17 AS builder

LABEL org.opencontainers.image.authors="Carlos Bustillo"

WORKDIR /build

COPY go.mod .
COPY go.sum .

# Install dependencies
RUN go mod download

# Copy the app code
COPY *.go .

# Build the app
RUN CGO_ENABLED=0 GOOS=linux go build -a -o test-mssql-connection

##
## Final Image
##
FROM alpine:3.15.0

WORKDIR /app

COPY --from=builder /build/test-mssql-connection .

# Create a group and user
RUN addgroup -S app && adduser -S app -G app -h /app

# Run the commands as app user
USER app

CMD [ "/app/test-mssql-connection" ]
