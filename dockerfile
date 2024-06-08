# Stage 1: Build the Go app
FROM golang:1.20-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app/

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

COPY * ./

# Copy the source from the current directory to the Working Directory inside the container
COPY vendor/ ./vendor/
COPY . .

# Copy the assets
COPY assets/ /app/assets/

# Copy the .env file
COPY .env .env
# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping


# Expose port 8080 to the outside world
EXPOSE 8090

# Command to run the executable
CMD ["./docker-gs-ping"]
