# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:latest

# Add Maintainer Info
LABEL maintainer="akbarl@live.com"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy everything from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main cmd/lazada_service/app.go

# Expose port 8080 to the outside world
EXPOSE 8181

# Command to run the executable
CMD ["./main"]