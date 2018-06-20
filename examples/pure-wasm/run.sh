#!/usr/bin/env bash

cd examples/pure-wasm

GOARCH=wasm GOOS=js /usr/local/Cellar/go/dev/bin/go build -o example.wasm github.com/Chyroc/webapi/examples/pure-wasm && serve