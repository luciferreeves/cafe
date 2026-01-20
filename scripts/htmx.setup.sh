#!/usr/bin/env bash
set -e

STATIC_JS_DIR="static/js"
HTMX_FILE="$STATIC_JS_DIR/htmx.min.js"

echo "Fetching latest HTMX release..."
LATEST_RELEASE=$(curl -s https://api.github.com/repos/bigskysoftware/htmx/releases/latest | grep '"tag_name":' | sed -E 's/.*"v([^"]+)".*/\1/')

if [ -z "$LATEST_RELEASE" ]; then
  echo "Failed to fetch latest release, using default version 2.0.8"
  LATEST_RELEASE="2.0.8"
fi

echo "Latest version: $LATEST_RELEASE"
DOWNLOAD_URL="https://unpkg.com/htmx.org@${LATEST_RELEASE}/dist/htmx.min.js"

echo "Downloading from: $DOWNLOAD_URL"
echo ""

mkdir -p "$STATIC_JS_DIR"

if curl -fSL "$DOWNLOAD_URL" -o "$HTMX_FILE"; then
  echo ""
  echo "HTMX installed successfully!"
  echo "Location: $HTMX_FILE"
  echo "Version: $LATEST_RELEASE"
else
  echo ""
  echo "Failed to download HTMX"
  exit 1
fi