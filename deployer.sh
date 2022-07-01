#!/bin/bash

# TODO - https://www.docker.com/blog/faster-multi-platform-builds-dockerfile-cross-compilation-guide/
# Build for both linux/amd64 (mac) and linux/arm/v7 (alpine)
docker buildx build --platform linux/amd64,linux/arm/v7 --push -t admiralpiett/musizticle .
#docker push admiralpiett/musizticle
