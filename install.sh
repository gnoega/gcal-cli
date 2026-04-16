#!/bin/bash

set -e

REPO="gnoega/gcal-cli"
TOOL_NAME="gcal-cli"
LATEST_URL="https://api.github.com/repos/$REPO/releases/latest"

# detect OS
OS=$(uname | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

# Map architecture to Go-compatible names
if [[ "$ARCH" == "x86_64" ]]; then
  ARCH=amd64
elif [[ "$ARCH" == "arm" ]]; then
  ARCH="arm64"
else
  echo "Unsupported architecture $ARCH"
  exit 1
fi

BINARY_NAME="${TOOL_NAME}-${OS}-${ARCH}"

echo "installing $TOOL_NAME for $OS/$ARCH..."

# downloading the binary
DOWNLOAD_URL=$(curl -sL $LATEST_URL | grep "browser_download_url" | grep "$BINARY_NAME" | cut -d '"' -f 4)

if [ -z "$DOWNLOAD_URL" ]; then
  echo "Error: Unable to find a binary for $OS/$ARCH."
  exit 1
fi

echo "Downloading $DOWNLOAD_URL..."
curl -L -o "$BINARY_NAME" "$DOWNLOAD_URL"

# make binary executable
chmod +x "$BINARY_NAME"

echo "Installing to /usr/local/bin"
sudo mv "$BINARY_NAME" /usr/local/bin/$TOOL_NAME

echo "$TOOL_NAME installed successfully! you can now use it by running '$TOOL_NAME'."
