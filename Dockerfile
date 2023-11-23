# Start from the latest golang base image
FROM golang:latest

EXPOSE 8080

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY *.go go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .
RUN go build -o ./src ./main.go


# Expose port 8080 to the outside


# This command runs your application, represented here as `uwu-bot`
ENTRYPOINT ["./trader-bot"]