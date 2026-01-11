#!/bin/bash

set -e

# Installation script for envcheck

REPO="LordCodex164/envcheck"
INSTALL_DIR="/usr/local/bin"
BINARY_NAME="envcheck"

# Detect OS and architecture
OS="$(uname -s)"
ARCH="$(uname -m)"

case "$OS" in
  Darwin)
    OS_NAME="darwin"
    ;;
  Linux)
    OS_NAME="linux"
    ;;
  *)
    echo "❌ Unsupported operating system: $OS"
    exit 1
    ;;
esac

case "$ARCH" in
  x86_64)
    ARCH_NAME="amd64"
    ;;
  arm64|aarch64)
    ARCH_NAME="arm64"
    ;;
  *)
    echo "❌ Unsupported architecture: $ARCH"
    exit 1
    ;;
esac

# envcheck-darwin-amd64
BINARY_FILENAME="${BINARY_NAME}-${OS_NAME}-${ARCH_NAME}"
DOWNLOAD_URL="https://github.com/${REPO}/releases/latest/download/${BINARY_FILENAME}"

echo "🔍 Detected: ${OS_NAME}-${ARCH_NAME}"
echo "📥 Downloading envcheck from ${DOWNLOAD_URL}..."

# Download binary
if command -v curl &> /dev/null; then
  curl -sSL "$DOWNLOAD_URL" -o "$BINARY_NAME"
elif command -v wget &> /dev/null; then
  wget -q "$DOWNLOAD_URL" -O "$BINARY_NAME"
else
  echo "❌ Neither curl nor wget found. Please install one of them."
  exit 1
fi

# Make executable
chmod +x "$BINARY_NAME"

# Move to install directory
echo "📦 Installing to ${INSTALL_DIR}/${BINARY_NAME}..."

if [ -w "$INSTALL_DIR" ]; then
  mv "$BINARY_NAME" "$INSTALL_DIR/"
else
  sudo mv "$BINARY_NAME" "$INSTALL_DIR/"
fi

echo "✅ envcheck installed successfully!"
echo ""
echo "Try it:"
echo "  envcheck --help"
echo "  envcheck create"
echo "  envcheck validate"