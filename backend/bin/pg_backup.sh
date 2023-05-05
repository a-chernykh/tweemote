#!/usr/bin/env bash

set -euo pipefail

echo "Backing up DB to ./backup/pg.tar.bz2"
docker run -it -v tweemote_postgres:/volume -v $(pwd)/backup:/backup alpine tar -cjf /backup/pg.tar.bz2 -C /volume ./
echo "Backup finished"
