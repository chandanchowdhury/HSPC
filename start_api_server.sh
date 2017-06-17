#!/bin/bash

export GOBIN=$GOPATH/bin

#go install ./cmd/hspc-server
#$GOBIN/hspc-server --port 8080

# reformat code
go fmt ./dbhandler
go fmt ./requesthandler

# export the environment variables
export POSTGRES_HOST="localhost"
export POSTGRES_DB="HSPC"
export POSTGRES_USER="hspc"
export POSTGRES_PASSWORD="HSPC-Password"

export MONGO_DB_HOST="localhost"
export MONGO_DB_USER="hspc"
export MONGO_DB_PASSWORD="HSPC-Password"
export MONGO_DB_AUTH_DB="admin"

# run the server
go run ./cmd/hspc-server/main.go --port 8080
