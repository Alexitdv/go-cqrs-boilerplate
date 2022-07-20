#!/bin/bash

# shellcheck disable=SC2046
export $(grep -v '^#' ./configs/env/.env."${1:-.env.dev}" | xargs) > /dev/null

$2