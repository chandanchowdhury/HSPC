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
    swagger generate server -A hspc -P models.Principal -f api_spec.yaml

    #go get -u -f ./...
fi


#sed -n s/"api.ServerShutdown = func() {}"/"requesthandler.Override_configure_hspc(api); api.ServerShutdown = func() {}"/1 restapi/configure_hspc.go #&2 > restapi/configure_hspc.go

#$LINE=`grep -n -F 'api.ServerShutdown = func() {}' restapi/configure_hspc.go | cut -d':' -f1`
#sed '/api.ServerShutdown = func() {}/i\requesthandler.Override_configure_hspc(api)\' restapi/configure_hspc.go
#grep -n -F 'api.ServerShutdown = func() {}' restapi/configure_hspc.go | cut -d':' -f1 | sed $line' a \requesthandler.Override_configure_hspc(api)' < restapi/configure_hspc.go

# add our code hook
#sed -i '' '
#/github\.com\/chandanchowdhury\/HSPC\/models/ i\
#\"github.com/chandanchowdhury/HSPC/requesthandler\"' restapi/configure_hspc.go
#
#sed -i '' '
#/api.ServerShutdown = func() {}/ i\
#\requesthandler.Override_configure_hspc(api)' restapi/configure_hspc.go
#
#gofmt -w restapi/configure_hspc.go

# reformat code
go fmt ./dbhandler
go fmt ./requesthandler
go fmt ./restapi


# Check code for suspecious constructs
go tool vet ./dbhandler
go tool vet ./requesthandler
go tool vet ./restapi
go tool vet ./models
