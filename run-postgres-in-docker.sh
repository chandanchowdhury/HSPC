#docker pull postgres

POSTGRES_DB=HSPC
POSTGRES_USER=hspc
# TODO: Get the password from command line or environment
POSTGRES_PASSWORD=password

PWD=`pwd`

# where the DB files will be stored
db_dir=$PWD/postgres_db

# make sure the DB directory exists
mkdir -p db_dir

# run postgres with required parameters
docker run -d -t -p 5432:5432 --rm \
    --name HSPC-postgres \
    -e POSTGRES_DB=$POSTGRES_DB \
    -e POSTGRES_USER=$POSTGRES_USER \
    -e POSTGRES_PASSWORD=POSTGRES_PASSWORD \
    -v $PWD/postgres_db:/var/lib/postgresql/data \
    postgres
