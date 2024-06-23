# Use the official Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
FROM golang:bookworm AS builder

# Copy the local package files to the container's workspace.
WORKDIR /backend
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the command inside the container.
# (You may need to modify the path based on your app structure)
WORKDIR /backend/cmd/echo-app
RUN go build -o /bin/app

# Use a Docker multi-stage build to create a lean production image.
# https://docs.docker.com/develop/develop-images/multistage-build/
FROM gcr.io/distroless/base-debian12:latest

# Install Nginx
RUN apt-get update && apt-get install -y nginx && apt-get clean && rm -rf /var/lib/apt/lists/*

# Copy the Nginx configuration file
COPY configs/nginx.conf /etc/nginx/nginx.conf

# Copy the built Go application
COPY --from=builder /bin/app /

# Expose port 80
EXPOSE 80

# Start Nginx and the Go application
CMD ["sh", "-c", "nginx && /app"]
