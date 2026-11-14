#!/usr/bin/env bash
set -euo pipefail
cd "$(dirname "$0")"
go version
go mod tidy
go test -count=1 -v ./...
