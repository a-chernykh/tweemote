#!/usr/bin/env bash

set -eu

grunt prod
aws s3 sync build s3://www.reactiveboost.com
