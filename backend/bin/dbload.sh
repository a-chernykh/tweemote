#!/usr/bin/env bash

set -euo pipefail

echo "Loading ${DUMP_PATH}"
docker-compose exec postgres psql --set ON_ERROR_STOP=on -U postgres tweemote < ${DUMP_PATH} > /dev/null
