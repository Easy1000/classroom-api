# Use an official Go image
FROM golang:1.23.5

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod tidy

# Copy the source code into the container
COPY . .

# Expose the port used by the API
EXPOSE 8080

# Default command (overridden by docker-compose)
CMD ["go", "run", "."]
