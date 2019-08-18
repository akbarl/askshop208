#!/bin/bash
IMAGE_NAME = "askshop208"
EXTERNAL_PORT = 8181
INTERNAL_PORT = 8181
echo "Building $IMAGE_NAME image"

sudo docker build -t $IMAGE_NAME .


