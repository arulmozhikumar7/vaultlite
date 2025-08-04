#!/bin/bash

set -e

REPO="arulmozhikumar7/vaultlite"
BIN_NAME="vaultlite"

# Detect OS and ARCH
OS=$(uname | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

# Normalize architecture
case "$ARCH" in
  x86_64) ARCH="amd64" ;;
  arm64|aarch64) ARCH="arm64" ;;
  *) echo "Unsupported architecture: $ARCH" && exit 1 ;;
esac

# Get latest version tag from GitHub
LATEST_TAG=$(curl -s "https://api.github.com/repos/$REPO/releases/latest" | grep '"tag_name":' | cut -d '"' -f 4)

if [[ -z "$LATEST_TAG" ]]; then
  echo "Failed to get latest release version"
  exit 1
fi

# Compose download URL
FILENAME="${BIN_NAME}-${LATEST_TAG}-${OS}-${ARCH}"
[[ "$OS" == "windows" ]] && FILENAME="${FILENAME}.exe"

DOWNLOAD_URL="https://github.com/${REPO}/releases/download/${LATEST_TAG}/${FILENAME}"

# Download binary
echo "Downloading $FILENAME from $DOWNLOAD_URL..."
curl -L "$DOWNLOAD_URL" -o "$BIN_NAME"

# Make it executable
chmod +x "$BIN_NAME"

# Move to /usr/local/bin (requires sudo)
echo "Installing to /usr/local/bin..."
sudo mv "$BIN_NAME" /usr/local/bin/$BIN_NAME

echo "✅ Installed $BIN_NAME version $LATEST_TAG"
echo "Run with: vaultlite"
