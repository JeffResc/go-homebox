#!/bin/bash
set -e

# Read the OpenAPI URL from .openapi-url file
OPENAPI_URL=$(cat .openapi-url)

echo "Fetching OpenAPI spec from: $OPENAPI_URL"

# Create temp directory for the spec
mkdir -p tmp

# Download the OpenAPI spec
curl -L -o tmp/openapi.json "$OPENAPI_URL"

echo "Generating Go client..."

# Generate the client using oapi-codegen
go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest -config .oapi-codegen.yaml tmp/openapi.json

echo "Generation complete!"
echo "Generated files:"
echo "  - client/client.go"
