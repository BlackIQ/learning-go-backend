#!/usr/bin/env bash

set -e

APP="todogo"
PACKAGE="."

mkdir -p out

build() {
    OS=$1
    ARCH=$2
    EXT=""

    if [ "$OS" = "windows" ]; then
        EXT=".exe"
    fi

    OUTPUT="out/${APP}-${OS}-${ARCH}${EXT}"

    echo "Building ${OS}/${ARCH}..."

    CGO_ENABLED=0 GOOS="$OS" GOARCH="$ARCH" \
        go build -trimpath -ldflags="-s -w" \
        -o "$OUTPUT" "$PACKAGE"
}

build linux amd64
build linux arm64
build windows amd64
build windows arm64
build darwin amd64
build darwin arm64

echo "Build completed!"
