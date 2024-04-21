# Use the official Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
FROM golang:bookworm as builder

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
FROM gcr.io/distroless/base-debian12

COPY --from=builder /bin/app /

# Run the outyet command by default when the container starts.
ENTRYPOINT ["/app"]

# Document that the service listens on port 8080.
EXPOSE 8080
