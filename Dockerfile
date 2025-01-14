# Base image
FROM golang:1.23-alpine

# Set working directory in the container
WORKDIR /app

# Copy the rest of the application code
COPY . .

#Copy the .env file into the container
COPY .env .env

# Download dependencies
RUN go mod tidy

# Build the Go application
RUN go build -o main ./cmd/main.go

EXPOSE 8080

# Command to run the executable // Or CMD ["app/main"]
CMD ["./main"]
