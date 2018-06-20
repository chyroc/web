#!/usr/bin/env bash

cd examples/webapi

GOARCH=wasm GOOS=js /usr/local/Cellar/go/dev/bin/go build -o example.wasm github.com/Chyroc/webapi/examples/webapi && serve