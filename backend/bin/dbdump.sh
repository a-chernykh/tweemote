#!/usr/bin/env bash

set -euo pipefail

docker-compose exec postgres pg_dump "$@" -U postgres tweemote
