#!/bin/bash

# export GOPATH=/Users/chandan/GOPATH

# export GOBIN=$GOPATH/bin
# export $PATH=$GOBIN:$PATH

swagger validate api_spec.yaml

if [ $? -eq 0 ]
then
    # rm -rf cmd restapi models
    swagger generate server -A hspc -f api_spec.yaml

    go get -u -f ./...
fi