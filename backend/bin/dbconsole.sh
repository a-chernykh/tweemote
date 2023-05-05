#!/usr/bin/env bash

set -eu

docker-compose exec postgres psql -U postgres tweemote
