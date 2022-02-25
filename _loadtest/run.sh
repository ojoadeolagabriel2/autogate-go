#!/usr/bin/env sh

# see: https://k6.io/docs/using-k6/metrics/

MY_PATH=$(dirname "$0")
MY_PATH=$(cd "$MY_PATH" && cd .. && pwd)
echo "running from... $MY_PATH"

k6 run "$MY_PATH"/_loadtest/test.js