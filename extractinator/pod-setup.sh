#!/bin/sh
set -eux
# Requires following env variables
# QUAY_TOKEN
# SOURCE_URL

podman pod rm -f csvpod || true

USER="rhceph-dev+readonly"
SOURCE_URL="quay.io/rhceph-dev/ocs-registry:4.10.0-110"

TAG=$(echo $SOURCE_URL | cut -d: -f2)

podman login -u="$USER" -p="$QUAY_TOKEN" quay.io

PORT=50051
podman run --rm -d -i --pod new:csvpod --name csv-container -p $PORT:$PORT "${SOURCE_URL}"
#podman run --rm -d -i --pod=csvpod --name=extractinator-container extractinator
#podman logs extractinator-container



echo "finished!"