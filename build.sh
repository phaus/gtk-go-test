#!/bin/bash

go get -u github.com/jteeuwen/go-bindata/...

go-bindata data/

CGO_ENABLED=1 go build -ldflags="-s -w" -o bin/gogtk
