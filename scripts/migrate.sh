#!/bin/bash
./scripts/env.sh "${1:-dev}" "go run ./tools/migration/main.go -dir ./migrations/sql $2"
