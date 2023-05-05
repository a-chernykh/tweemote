#!/usr/bin/env bash

set -eu

CMDS=("DROP DATABASE IF EXISTS tweemote"
      "CREATE DATABASE tweemote")
IFS=""

for cmd in ${CMDS[*]}; do
  echo "Running $cmd"
  docker-compose exec postgres psql -U postgres -c "$cmd" postgres
done
