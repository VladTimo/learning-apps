# Use the official Go image as the build environment
FROM golang:1.23.3 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go files into the container
COPY . .

# Run go mod tidy to ensure dependencies are fetched
RUN go mod tidy

# Build the Go binary statically linked
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/hello .

# Create a minimal runtime image (alpine)
FROM alpine:latest

# Set the working directory for the runtime container
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/hello .

# Expose the application port
EXPOSE 8089

# Run the binary when the container starts
CMD ["./hello"]
