#!/bin/bash
docker pull postgres

POSTGRES_DB=HSPC
POSTGRES_USER=hspc
# TODO: Get the password from command line or environment
POSTGRES_PASSWORD=HSPC-Password

PWD=`pwd`

# where the DB files will be stored
DBPATH=$PWD/postgres_db

if [ "$1" = "clean" ]
then
    echo "Deleting direcotries..."
    rm -rf $DBPATH
fi

# make sure the DB directory exists
mkdir -p $DBPATH

# run postgres with required parameters
docker run -d -t -p 5432:5432 --rm \
    --name HSPC-postgres \
    -e POSTGRES_DB=$POSTGRES_DB \
    -e POSTGRES_USER=$POSTGRES_USER \
    -e POSTGRES_PASSWORD=$POSTGRES_PASSWORD \
    -v $DBPATH:/var/lib/postgresql/data \
    postgres
