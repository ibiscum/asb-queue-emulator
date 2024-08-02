#!/bin/bash

# Get the directory of the current script
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

# Use the Python script to merge the YAML files
swagger -q mixin $DIR/../swagger/base.yaml $DIR/../swagger/api.yaml -o $DIR/../swagger/azure-servicebus-spec.yaml

# Generate the server using the merged YAML
swagger generate server -f "$DIR/../swagger/azure-servicebus-spec.yaml" -t "$DIR/../swagger/gen" --exclude-main
