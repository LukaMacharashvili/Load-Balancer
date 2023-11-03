# Use an official Go runtime as a parent image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the local directory to the container at /app
ADD . /app

# Build the Go application inside the container
RUN go build -o main cmd/lb/main.go

# Expose port 3002 for the application
EXPOSE 3002

# Command to run the application with the TARGET_URLS environment variable
CMD ["./main"]
