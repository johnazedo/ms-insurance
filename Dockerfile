FROM golang:1.19-alpine

RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /app

RUN apk add --no-cache \
    # Important: required for go-sqlite3
    gcc \
    # Required for Alpine
    musl-dev

# Copy go mod and sum files
COPY go.mod ./
COPY go.sum ./

ENV GOPROXY="https://proxy.golang.org,direct"
# Download all dependancies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o bin/main

# Expose port 8080 to the outside world
EXPOSE 8000

## Run the executable
CMD ["./bin/main"]