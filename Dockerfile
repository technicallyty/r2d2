# Start from the official minimal Go image
FROM golang:1.24-alpine

# Set environment variables
ENV CGO_ENABLED=0 \
    GO111MODULE=on \
    GOOS=linux \
    GOARCH=amd64

# Create a working directory
WORKDIR /app

# Copy Go module files and download dependencies early to leverage Docker cache
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy the rest of the code
COPY . .

# Build the Go binary
RUN go build -o /go-action .

# Run the binary when the container starts
ENTRYPOINT ["/go-action"]
