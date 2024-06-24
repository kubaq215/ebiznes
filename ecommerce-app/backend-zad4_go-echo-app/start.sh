#!/bin/sh

nginx -g "daemon off;" &

# Run the Go app
/bin/bash /app