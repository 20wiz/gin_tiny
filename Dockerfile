# Stage 1: Build the Go application
FROM golang:1.23-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY ./go.mod ./
COPY ./go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN go build -o gin_tiny main.go

# Stage 2: Run the application
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/gin_tiny .

# Expose port 8080
EXPOSE 8080

# Command to run the executable
CMD ["./gin_tiny"]