#!/usr/bin/env bash

CHROOT="/usr/x86_64-pc-linux-gnu-static"

export CC=gcc
export PKG_CONFIG_PATH="$CHROOT/usr/lib/pkgconfig"
export PKG_CONFIG_LIBDIR="$CHROOT/usr/lib/pkgconfig"
export LIBRARY_PATH="$CHROOT/usr/lib:$CHROOT/lib"

CGO_LDFLAGS="-L$CHROOT/usr/lib -L$CHROOT/lib -lasound" \
CGO_ENABLED=1 GOOS=linux GOARCH=amd64 \
go build -v -x -ldflags "-linkmode external -s -w" github.com/gen2brain/flite-go/cmd/say
