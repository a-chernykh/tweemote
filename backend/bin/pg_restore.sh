#!/usr/bin/env bash

set -euo pipefail

RESTORE_DIR=$(pwd)/backup

echo "Are you sure you want to restore database from ${RESTORE_DIR}? Type 'yes' and press ENTER to confirm."
read confirmation

if [ "$confirmation" == "yes" ]; then
  docker-compose stop postgres
  echo "Restoring database"
  docker run -it -v tweemote_postgres:/volume -v ${RESTORE_DIR}:/backup alpine sh -c "rm -rf /volume/* ; tar -C /volume/ -xjf /backup/pg.tar.bz2"
  echo "Restoring finished"
  docker-compose up -d postgres
fi
