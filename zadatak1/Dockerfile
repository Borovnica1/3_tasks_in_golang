# Use the official Golang image as the base image
FROM golang:1.17-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the necessary files to the container
COPY . .

# Build the Go binary
RUN go build -o microservice

# Expose the port on which the microservice will listen
EXPOSE 8080

# Set the entry point for the container
CMD ["./microservice"]