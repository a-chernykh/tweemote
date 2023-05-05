#!/usr/bin/env bash

set -eu

export DOMAIN=lvh.me
openssl req -subj "/CN=${DOMAIN}/C=US/L=Mountain View/ST=CA/O=Localhost" -newkey rsa:2048 -nodes -keyout ssl/${DOMAIN}.key -out ssl/${DOMAIN}.csr
openssl x509 -req -sha256 -days 3650 -in ssl/${DOMAIN}.csr -signkey ssl/${DOMAIN}.key -out ssl/${DOMAIN}.crt
