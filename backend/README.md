# Development

## Setup project

### Configure git to use SSH urls for bitbucket

Run:

    git config --global url."git@bitbucket.org:".insteadOf "https://bitbucket.org/"

### Install project

Run:

    go get bitbucket.org/andreychernih/tweemote

Project directory will be located at:

    $GOPATH/src/bitbucket.org/andreychernih/tweemote

### Install dependencies

Run:

    govendor sync

### Start all services

Run:

    make linux

## Dependencies

Dependencies are controlled by `govendor`. They reside in [vendor/](/vendor) directory and defined in [vendor/vendor.json](vendor/vendor.json) file. To fetch all dependencies from [vendor/vendor.json](vendor/vendor.json), run:

    govendor sync

It will download all dependencies.

After introducing new dependency, add import statement as usual and run:

    go get <dep_url>
    govendor add +external

This will fetch it from the remote and add it to [vendor/vendor.json](vendor/vendor.json)

Update package to the latest:

    govendor fetch github.com/jinzhu/gorm

Install updated dependency:

    govendor install


## REPL console

You can start [gore](https://github.com/motemen/gore) by running:

    docker-compose run console

And then run gore inside it:

    gore

## Testing

Create test database:

    docker-compose exec postgres psql -U postgres -c 'CREATE DATABASE tweemote_test'

## Database migrations

Create new migration by running `bin/create_migration.sh`. Example:

    NAME=create_twitter_applications bin/create_migration.sh

Migrate databae:

    make migrate

Seed database:

    make
    tweemote seed

# Building

See [Makefile](Makefile) for details.

## Local

If you want to start it locally, use `bin/run.sh` script. It will run `make local` to build local binary.

## Linux

`make linux` will build a Linux binary using `Dockerfile.build`.

# Testing

## Create test DB

Run:

    make testdb

It will create new PostgreSQL database `tweemote_test` using database named `tweemote` as a schema template.
