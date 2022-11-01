#!/bin/env sh

arch=$ARCHFLAGS | sed -E 's/-arch|[ ,]//g'

if [ ! -x "$(command -v gccgo)" ]; then
    echo "GCCGo doesn't seems to be installed or in PATH. Falling back with default go compiler"
    if [ ! -x "$(command -v go)" ]; then
        echo "Go isn't installed in your system."
        exit
    fi
    env GOARCH=$arch go build
else
    gccgo *.go -O2 -march=native -s -o netspeed
fi
