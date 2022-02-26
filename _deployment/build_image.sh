#!/usr/bin/env sh

set -ex

MY_PATH=$(dirname "$0")
MY_PATH=$(cd "$MY_PATH" && cd .. && pwd)
CONTAINER_NAME=test_autogate

export PATH=$PATH:/usr/local/go/bin

go env
go mod download
# go test -v "$MY_PATH"/...
go build

docker system prune -af
docker build -t app/autogate:latest .

docker stop "$CONTAINER_NAME" || true && docker rm "$CONTAINER_NAME" || true
docker run --name "$CONTAINER_NAME" -p 12345:12345 -d app/autogate:latest