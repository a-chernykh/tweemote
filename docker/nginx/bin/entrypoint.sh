#!/usr/bin/env sh

set -eu

export NAMESERVERS=$(cat /etc/resolv.conf | grep "nameserver" | awk '{print $2}' | tr '\n' ' ')
envsubst '$NAMESERVERS' < /etc/nginx/nginx.conf.tmpl > /etc/nginx/nginx.conf

exec "$@"
