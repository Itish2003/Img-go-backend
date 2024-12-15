# Start from the Go base image
FROM golang:latest AS build
WORKDIR /app

# Copy Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the application source code
COPY . .

# Install the primitive image processing tool
RUN go install github.com/fogleman/primitive@latest

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server main.go


# Final container
FROM alpine:latest

# Install necessary runtime dependencies, such as certificates
RUN apk --no-cache add ca-certificates
WORKDIR /root/

# Copy the built server and the primitive binary from the build stage
COPY --from=build /app/server .
COPY --from=build /go/bin/primitive /usr/local/bin/primitive

# Expose the backend port
EXPOSE 8080

# Run the server
CMD ["./server"]
