# Step 1: Build the Go app
FROM golang:1.25 AS build

WORKDIR /app

# Copy go.mod and go.sum first for caching
COPY go.mod go.sum ./
RUN go mod download

# Copy entire project
COPY . .

# Build the binary
RUN go build -o petclinic main.go

# Step 2: Run the Go app
FROM debian:bookworm-slim

WORKDIR /app

# Copy binary from builder
COPY --from=build /app/petclinic .

# Copy uploads folder
RUN mkdir -p uploads

# Expose API port
EXPOSE 8080

# Start the app
CMD ["./petclinic"]
