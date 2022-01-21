# extractinator

Extractinator is a simple tool for extracting information about a container's bundles using grpcurl. The tool requires the following as positional arguments:
1. A username for quay.io
2. A source URL for a container image

Before using extractinator, you must have the following dependencies:
- podman
- grpcurl
- jq

To authenticate to quay.io there must be a valid quay.io token in $QUAY_TOKEN