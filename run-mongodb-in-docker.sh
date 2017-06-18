#!/bin/bash
docker pull mongo:3

MONGO_DB=HSPC
MONGO_USER=hspc
# TODO: Get the password from command line or environment
MONGO_PASSWORD=HSPC-Password

PWD=`pwd`

# where the DB files will be stored
DBPATH=$PWD/mongodb_store

if [ "$1" = "clean" ]
then
    echo "Deleting direcotries..."
    rm -rf $DBPATH
fi

# make sure the DB directory exists
mkdir -p $DBPATH

# run MongoDB with required parameters
docker run -d -t --rm -p 27017:27017 \
    --name $MONGO_DB-Mongo \
    -v $DBPATH:/data/db \
    -e MONGO_INITDB_ROOT_USERNAME=$MONGO_USER \
    -e MONGO_INITDB_ROOT_PASSWORD=$MONGO_PASSWORD \
    mongo:3