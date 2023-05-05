#!/usr/bin/env bash

set -eu

# https://github.com/docker-library/healthcheck/tree/master/rabbitmq

host="$(hostname --short || echo 'localhost')"
export RABBITMQ_NODENAME="${RABBITMQ_NODENAME:-"rabbit@$host"}"

if rabbitmqctl status; then
	exit 0
fi

exit 1
