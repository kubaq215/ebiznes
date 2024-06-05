#!/bin/bash

# Function to forward port 80 to port 3000
forward_port() {
    echo "Forwarding port 80 to port 3000..."
    sudo iptables -t nat -A PREROUTING -p tcp --dport 80 -j REDIRECT --to-port 3000
    echo "Port forwarding added."
}

# Function to remove the forwarding rule
unforward_port() {
    echo "Removing port forwarding from port 80 to port 3000..."
    sudo iptables -t nat -D PREROUTING -p tcp --dport 80 -j REDIRECT --to-port 3000
    echo "Port forwarding removed."
}

# Check the argument provided
if [ "$1" == "forward" ]; then
    forward_port
elif [ "$1" == "unforward" ]; then
    unforward_port
else
    echo "Usage: $0 {forward|unforward}"
    exit 1
fi
