# Start from the latest golang base image
FROM golang:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . ./

# Declare DISCORD_BOT_TOKEN as a build argument
ARG DISCORD_BOT_TOKEN

ARG SQL_TOKEN

# Set the DISCORD_BOT_TOKEN environment variable
ENV DISCORD_BOT_TOKEN=$DISCORD_BOT_TOKEN

ENV SQL_TOKEN=$SQL_TOKEN

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]