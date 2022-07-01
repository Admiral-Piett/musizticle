#!/bin/bash

CONTAINER="musizticle"
IMAGE_NAME="admiralpiett/${CONTAINER}"

# Stop any affected running containers and remove the images before pulling latest
docker stop $CONTAINER
docker rmi $(docker images --format "{{ .Repository }}:{{ .Tag }}" $IMAGE_NAME)

docker pull "${IMAGE_NAME}:latest"
