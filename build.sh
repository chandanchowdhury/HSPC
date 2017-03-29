#!/bin/bash

# export GOPATH=/Users/chandan/GOPATH

# export GOBIN=$GOPATH/bin
# export $PATH=$GOBIN:$PATH


swagger validate api_spec.yaml


if [ $? -eq 0 ]
then
    if [ "$1" = "clean" ]
    then
        echo "Deleting direcotries..."
        rm -rf cmd restapi models
    fi
    swagger generate server -A hspc -f api_spec.yaml

    go get -u -f ./...
fi