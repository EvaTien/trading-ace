# Use the official Golang image as the base image
FROM golang:1.23.2-alpine
# Set the Current Working Directory inside the container
WORKDIR /app
# Copy go.mod and go.sum files
COPY go.mod ./
# Install dependencies
RUN go mod tidy
# Copy the source code
COPY . .
# Expose port 8081 to the outside world
EXPOSE 8080
# Command to run the executable
CMD ["go", "run", "main.go"]
