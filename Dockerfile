# Stage 1: Build
FROM golang:1.22 AS builder

WORKDIR /app

# Copy Go modules and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code 
COPY . .

# Build the Go app 
RUN CGO_ENABLED=0 GOOS=linux go build -o app .

# Stage 2: Minimal
FROM alpine:lastest

# Add CA certificates (needed if your app makes HTTPS request)
RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy the built Go binary from the builder Stage
COPY --from=builder /app/app .

# Expose port
EXPOSE 8080

# Mount point for audio recordings 
VOLUME ["/app/recordings"]

# Run the binary
ENTRYPOINT ["./app"]
