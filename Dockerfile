# Dockerfile References: https://docs.docker.com/engine/reference/builder/

### STAGE 1: BUILD ###
FROM golang:1.19.0-alpine as builder
# Create app directory
RUN mkdir /app
# Set the Current Working Directory inside the container
WORKDIR /app

ADD . /app
# Download all dependancies. Dependencies will be cached if the go.mod and go.sum files are not changed
CMD PWD
RUN go mod download
# Build the Go api
RUN go build -o ./api ./cmd/api

### STAGE 2: RUN ###
FROM golang:1.19.0-alpine
COPY --from=builder /app/api /go/bin/api
# Run the executable
CMD ["api"]