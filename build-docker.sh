#!/bin/bash

export GOPATH=$(pwd)/../../../../



go get
go build -ldflags "-linkmode external -extldflags -static -s"

docker build -t homeip .