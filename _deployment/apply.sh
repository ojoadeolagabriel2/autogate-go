#!/bin/sh

set -ex

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

kubectl delete namespace autogate-go-ns || true
kubectl apply -f "${DIR}/apply.yaml"