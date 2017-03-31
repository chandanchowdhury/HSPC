#!/bin/bash
docker pull mongodb

MONGO_DB=HSPC
MONGO_USER=hspc
# TODO: Get the password from command line or environment
MONGO_PASSWORD=HSPC-Password

PWD=`pwd`

# where the DB files will be stored
DBPATH=$PWD/mongodb_store

# make sure the DB directory exists
mkdir -p $DBPATH

# run MongoDB with required parameters
docker run -d -t --rm -p 27017:27017 \
    --name $MONGO_DB \
    -v $DBPATH:/data/db \
    -e MONGO_INITDB_ROOT_USERNAME=$MONGO_USER \
    -e MONGO_INITDB_ROOT_PASSWORD=$MONGO_PASSWORD \
    mongo