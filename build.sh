#!/bin/bash

# Build script for envcheck

VERSION="1.0.0"
COMMIT=$(git rev-parse --short HEAD 2>/dev/null || echo "dev")
DATE=$(date -u +"%Y-%m-%dT%H:%M:%SZ")

# Build for current platform
echo "Building envcheck v$VERSION...".env
go build -ldflags="-X 'env-checker/cmd.version=$VERSION' -X 'env-checker/cmd.commit=$COMMIT' -X 'env-checker/cmd.date=$DATE'" -o envcheck .

echo "✓ Built: ./envcheck"
echo ""
echo "Test it:"
echo "  ./envcheck --help"
echo "  ./envcheck version"