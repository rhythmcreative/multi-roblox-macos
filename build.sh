#!/bin/bash
set -e

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd "$DIR"

echo "==> Building native arm64 binary..."
GOOS=darwin GOARCH=arm64 CGO_ENABLED=1 go build -ldflags="-s -w" -o build/bin-arm64 .

echo "==> Building x86_64 binary..."
GOOS=darwin GOARCH=amd64 CGO_ENABLED=1 go build -ldflags="-s -w" -o build/bin-amd64 .

echo "==> Creating Universal Binary (arm64 + x86_64)..."
mkdir -p build
lipo -create -output build/multi-roblox-macos build/bin-arm64 build/bin-amd64

echo "==> Installing into multiroblox.app bundle..."
mkdir -p multiroblox.app/Contents/MacOS
cp build/multi-roblox-macos multiroblox.app/Contents/MacOS/multi-roblox-macos
chmod +x multiroblox.app/Contents/MacOS/multi-roblox-macos

echo "==> Ad-hoc signing multiroblox.app..."
codesign --force --deep --sign - multiroblox.app

echo "==> Verifying codesign..."
codesign -v multiroblox.app

echo "==> Copying standalone universal binary to root..."
cp build/multi-roblox-macos multi-roblox-macos

echo "==> Cleaning build artifacts..."
rm -rf build

echo "==> Build complete successfully!"
