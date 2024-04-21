#!/bin/bash

# Check if Docker daemon is running
if ! docker info > /dev/null 2>&1; then
  echo "Docker daemon is not running. Please start the Docker daemon."
  exit 1
else
  echo "Docker daemon is running."
fi

# Check if docker-compose.yml exists in the current directory
if [ ! -f "./docker-compose.yml" ]; then
  echo "docker-compose.yml does not exist in the current directory."
  exit 1
else
  echo "docker-compose.yml found. Starting services..."
  # Run docker-compose to build and start your containers
  docker compose up --build
fi
