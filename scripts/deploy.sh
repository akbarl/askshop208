#!/bin/bash
IMAGE_NAME="askshop208"
EXTERNAL_PORT=8181
INTERNAL_PORT=8181
echo "Building $IMAGE_NAME image"

sudo docker build -f ../deployment/-t $IMAGE_NAME .

echo "Created $IMAGE_NAME successfully"

echo "Starting a container with $IMAGE_NAME"
sudo docker run -d -p $EXTERNAL_PORT:$INTERNAL_PORT $IMAGE_NAME

sudo docker ps
echo "Started the container  successfully"
