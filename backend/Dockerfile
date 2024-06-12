# Use a Go base image with the same version as your local environment
FROM golang:1.22.3-alpine

# Set the current working directory inside the container
WORKDIR /app

# Copy the Go modules files
COPY go.mod ./
COPY go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o main ./cmd/scheduler

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
