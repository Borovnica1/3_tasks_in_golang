# Use the official Go image to build the application
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go source code into the container
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the port your application listens on (assuming it's port 8080)
EXPOSE 8080

# Command to run the application when the container starts
CMD ["./main"]
