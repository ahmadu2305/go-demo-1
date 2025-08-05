# Start from a lightweight Go image
FROM golang:1.20-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy Go module files and source code
COPY go.mod ./
COPY main.go ./
COPY mascot ./mascot

# Build the Go binary
RUN go build -o app

# Run the binary as default command
CMD ["./app"]