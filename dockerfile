# Stage 1: Build the Go app
FROM golang:1.20-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source files from the current directory to the Working Directory inside the container
COPY . .

# Copy the assets
COPY assets/ /app/assets/

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/pos-backend

# Stage 2: Run the Go app
FROM alpine:latest

WORKDIR /app

# Copy the Pre-built binary file and assets from the previous stage
COPY --from=builder /app/pos-backend .
COPY --from=builder /app/assets/ ./assets/
COPY --from=builder /app/. .

# Copy the .env file
COPY .env .env

# Expose port 8080 to the outside world
EXPOSE 8090

# Command to run the executable
CMD ["./pos-backend"]
