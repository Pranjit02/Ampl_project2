# Step 1: Build the Go binary
FROM golang:1.19 AS builder

WORKDIR /app

# Copy Go Modules manifests
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod tidy

# Copy the source code
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Step 2: Create a smaller image to run the Go application
FROM alpine:latest

# Install necessary libraries (for MySQL client if required)
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Go binary from the builder image
COPY --from=builder /app/main .

# Expose the application port
EXPOSE 8080

# Command to run the app
CMD ["./main"]
