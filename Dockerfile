# Start from golang base image
FROM golang:alpine as builder

# ENV GO111MODULE=on

# Add Maintainer info
LABEL maintainer="Tikno <tikno.edu@gmail.com>"

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

# Set the current working directory inside the container 
WORKDIR /app

# Copy go mod and sum files 
COPY go.mod go.sum ./

# RUN go mod tidy

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed 
RUN go mod download 

# Copy the source from the current directory to the working Directory inside the container 
COPY . .
# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo cmd/monitor/main.go

# Start a new stage from scratch
FROM alpine:latest
# RUN apk --no-cache add ca-certificates

RUN mkdir /home/db-monitor

WORKDIR /home/db-monitor

RUN mkdir config

RUN mkdir cmd

RUN mkdir cmd/monitor

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/config config
COPY --from=builder /app/main cmd/monitor
COPY --from=builder /app/.env .

WORKDIR cmd/monitor

RUN pwd

# Expose port 8080 to the outside world
EXPOSE 8080

#Command to run the executable
CMD ["./main"]