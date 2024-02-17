# Use the official Golang image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod .
COPY go.sum .

# Download Go dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Copy the wait-for-it.sh script from the root directory to the working directory inside the container
COPY wait-for-it.sh /app/wait-for-it.sh

# Set execute permissions for the script
RUN chmod +x /app/wait-for-it.sh

# Build the Go application
RUN go build -o main .

# Expose the port that your application runs on
EXPOSE 8080

# Command to run the executable
CMD ["./wait-for-it.sh", "db:5432", "-t", "600", "--", "./main"]