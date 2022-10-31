#!/bin/env sh

if [ ! gccgo ]; then
    echo "GCCGo doesn't seems to be installed or in PATH. Falling back with default go compiler"
    if [ ! go ]; then
        echo "Go isn't installed in your system."
        exit
    fi
    go build
else
    gccgo *.go -O2 -march=native -s -o netspeed
fi
