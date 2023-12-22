# Use the official Golang image to build your application
FROM golang:latest


# Set necessary environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Set the working directory inside the container
WORKDIR /app


# Copy the Go modules and download dependencies
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the entire application code into the container
COPY . .

# Build the Go application
RUN go build -o main ./cmd/go_auth/main.go

# Expose the port your application uses
EXPOSE 8080

# Command to run the application when the container starts
CMD ["./main"]
