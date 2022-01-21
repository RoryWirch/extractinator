#!/bin/sh
set -eux

#
# extractinator.sh is a simple tool that extracts data from a container image using grpcurl.
# It requires the following positional arguments:
#   1.) username for quay.io authentication
#   2.) source URL for the container image
# Dependendies: grpcurl, podman, jq
#

USER=$1
SOURCE_URL=$2

TAG=$(echo $SOURCE_URL | cut -d: -f2)

# Authenticate to quay.io
sudo podman login --authfile /root/.docker/config.json -u="$USER" -p="$QUAY_TOKEN" quay.io

PORT=50051
podman run --rm -d --pod new:csvpod --name csv-container -p $PORT:$PORT "${SOURCE_URL}"

HOST=localhost
GRPCURL="podman run --rm --pod csvpod fullstorydev/grpcurl:latest"
$GRPCURL -plaintext $HOST:50051 api.Registry/ListBundles > bundles.json
CHANNEL=$(cat bundles.json | jq -r \
    ". | select(.packageName == \"$CSV_NAME\") | select(.version as \$version | \"$TAG\" | test(\$version)) | .channelName"
)
$GRPCURL \
    -plaintext \
    -d "{\"pkgName\": \"$CSV_NAME\", \"channelName\": \"$CHANNEL\"}" \
    $HOST:$PORT \
    api.Registry/GetBundleForChannel > bundle.json
cat bundle.json | jq '.csvJson | fromjson' > $CSV_NAME.json
