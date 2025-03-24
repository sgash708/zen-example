FROM golang:1.24-alpine AS builder

# Install git for dependency downloads
RUN apk add --no-cache git

# Configure git to use HTTP instead of HTTPS
RUN git config --global url."https://".insteadOf git://

WORKDIR /app

# Copy go mod and sum files
COPY go.mod ./

# Download dependencies with GOPROXY to avoid direct GitHub access
RUN GOPROXY=https://proxy.golang.org go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o apiserver main.go

# Use a smaller image for the final stage
FROM alpine:latest

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/apiserver .

# Expose port
EXPOSE 8080

# Run the application
CMD ["./apiserver"]