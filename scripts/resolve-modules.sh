#!/bin/bash
# Recursively finds all directories with a go.mod file and creates
# a GitHub Actions JSON output option. This is used by several actions.

# Based on https://github.com/golangci/golangci-lint/issues/828#issuecomment-658207652
echo "Resolving modules in $(pwd)"

PATHS=$(find . -mindepth 2 -type f -name go.mod -printf '{"workdir":"%h"},' | sed 's/\.\///g')
echo "::set-output name=matrix::{\"include\":[${PATHS%?}]}"