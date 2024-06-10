#!/bin/bash

# Start the Play application
sudo docker run -p 9000:9000 scala-app &

# Start ngrok to expose the Play application
ngrok http 9000
