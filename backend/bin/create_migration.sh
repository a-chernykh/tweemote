#!/usr/bin/env bash -eu

TIMESTAMP=$(date +%s)

function createFile() {
  echo "create $1"
  touch $1
}

createFile migrations/${TIMESTAMP}_${NAME}.up.sql
createFile migrations/${TIMESTAMP}_${NAME}.down.sql
