#!/bin/bash

VERSION="1.0.0"
COMMIT=$(git rev-parse --short HEAD 2>/dev/null || echo "dev")
DATE=$(date -u +"%Y-%m-%dT%H:%M:%SZ")

LDFLAGS="-X 'env-checker/cmd.version=$VERSION' -X 'env-checker/cmd.commit=$COMMIT' -X 'env-checker/cmd.date=$DATE'"

# Create dist directory
mkdir -p dist

# Build for multiple platforms
echo "Building envcheck v$VERSION for multiple platforms..."

# macOS (Intel)
GOOS=darwin GOARCH=amd64 go build -ldflags="$LDFLAGS" -o dist/envcheck-darwin-amd64 .
echo "✓ Built: dist/envcheck-darwin-amd64"

# macOS (Apple Silicon)
GOOS=darwin GOARCH=arm64 go build -ldflags="$LDFLAGS" -o dist/envcheck-darwin-arm64 .
echo "✓ Built: dist/envcheck-darwin-arm64"

# Linux (amd64)
GOOS=linux GOARCH=amd64 go build -ldflags="$LDFLAGS" -o dist/envcheck-linux-amd64 .
echo "✓ Built: dist/envcheck-linux-amd64"

# Windows
GOOS=windows GOARCH=amd64 go build -ldflags="$LDFLAGS" -o dist/envcheck-windows-amd64.exe .
echo "✓ Built: dist/envcheck-windows-amd64.exe"

echo ""
echo "All builds complete! Check ./dist/"