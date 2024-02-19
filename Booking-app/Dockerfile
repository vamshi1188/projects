# Start with a lightweight base image
FROM golang:1.22.0-alpine3.19 AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules manifests
COPY go.mod ./


# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o myapp .

# Start a new stage from scratch
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the executable from the previous stage
COPY --from=build /app/myapp .

# Command to run the executable
CMD ["./myapp"]
