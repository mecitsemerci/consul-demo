# Dockerfile References: https://docs.docker.com/engine/reference/builder/

### STAGE 1: BUILD ###
FROM golang:1.19.0-alpine as builder
# The latest alpine images don't have some tools like (`git` and `bash`).
# Adding git, bash and openssh to the image
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
# Create app directory
RUN mkdir /app
# Set the Current Working Directory inside the container
WORKDIR /app
ADD ../.. /app
# Download all dependancies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download
# Build the Go api
RUN go build -o ./api ./cmd/api

### STAGE 2: RUN ###
FROM golang:1.19.0-alpine
COPY --from=builder /app/api /go/bin/api
# Run the executable
CMD ["api"]