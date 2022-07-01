#!/bin/bash

function on_death() {
  echo "Stopping container"
  docker stop musizticle
}
trap on_death INT

IMAGE_NAME="admiralpiett/musizticle"
CONTAINER_NAME="musizticle"

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"
DATA_VOLUME=${DIR}/data/:/opt/svc/data/
MUSIC_VOLUME=${DIR}/music/:/opt/svc/music/

echo docker run -d --rm --name=$CONTAINER_NAME -p=9000:9000 --env-file=.env -v $DATA_VOLUME -v $MUSIC_VOLUME $IMAGE_NAME
docker run -d --rm --name=$CONTAINER_NAME -p=9000:9000 --env-file=.env -v $DATA_VOLUME -v $MUSIC_VOLUME $IMAGE_NAME

docker exec -it $CONTAINER_NAME sh -c "tail -F /var/log/musizticle.log"

echo "Stopping container"
docker stop $CONTAINER_NAME
