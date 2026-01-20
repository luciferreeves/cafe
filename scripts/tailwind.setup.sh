#!/usr/bin/env bash
set -e

TOOLCHAIN_DIR="toolchain"
TAILWIND_BIN="$TOOLCHAIN_DIR/tailwind"

echo "Detecting platform..."

OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case "$OS" in
  linux*)
    PLATFORM="linux"
    ;;
  darwin*)
    PLATFORM="macos"
    ;;
  msys*|mingw*|cygwin*)
    PLATFORM="windows"
    ;;
  *)
    echo "Unsupported OS: $OS"
    exit 1
    ;;
esac

case "$ARCH" in
  x86_64|amd64)
    ARCHITECTURE="x64"
    ;;
  arm64|aarch64)
    ARCHITECTURE="arm64"
    ;;
  armv7l)
    ARCHITECTURE="armv7"
    ;;
  *)
    echo "Unsupported architecture: $ARCH"
    exit 1
    ;;
esac

BINARY_NAME="tailwindcss-${PLATFORM}-${ARCHITECTURE}"

echo "Platform: $PLATFORM"
echo "Architecture: $ARCHITECTURE"
echo "Binary: $BINARY_NAME"
echo ""

echo "Fetching latest Tailwind CSS release..."
LATEST_RELEASE=$(curl -s https://api.github.com/repos/tailwindlabs/tailwindcss/releases/latest | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')

if [ -z "$LATEST_RELEASE" ]; then
  echo "Failed to fetch latest release"
  exit 1
fi

echo "Latest version: $LATEST_RELEASE"
DOWNLOAD_URL="https://github.com/tailwindlabs/tailwindcss/releases/download/${LATEST_RELEASE}/${BINARY_NAME}"

echo "Downloading from: $DOWNLOAD_URL"
echo ""

mkdir -p "$TOOLCHAIN_DIR"

if curl -fSL "$DOWNLOAD_URL" -o "$TAILWIND_BIN"; then
  chmod +x "$TAILWIND_BIN"
  echo ""
  echo "Tailwind CSS installed successfully!"
  echo "Location: $TAILWIND_BIN"
  echo "Version: $($TAILWIND_BIN --help | head -n 1 || echo $LATEST_RELEASE)"
else
  echo ""
  echo "Failed to download Tailwind CSS binary"
  echo "Please check if the binary exists for your platform:"
  echo "$DOWNLOAD_URL"
  exit 1
fi