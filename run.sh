#!/bin/bash

export GOBIN=$GOPATH/bin

#go install ./cmd/hspc-server
#$GOBIN/hspc-server --port 8080

go run ./cmd/hspc-server/main.go --port 80