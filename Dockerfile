# Step 1. Build stage
FROM golang:1.25.4 as builder
WORKDIR /app

# Download modules before copying full source for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server ./cmd/main
# Step 2. Run stage
FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 8080
CMD ["./server"]
 
