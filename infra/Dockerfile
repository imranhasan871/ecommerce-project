# Use official Golang image as base
FROM golang:1.25-alpine

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy application code
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port (change if your app uses a different port)
EXPOSE 8000

# Start the application
CMD ["./main"]