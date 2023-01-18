#!/usr/bin/env bash
set -e

CHANGE_ONLY="false"
GO_MOD_DIRS=$(find . -type f -name "go.mod" | xargs -n1 dirname)

print_usage() {
    echo "Runs unit-tests for all modules."
    echo
    echo "Syntax: ./test-unit.sh [-c]"
    echo "options:"
    echo "c     Run only on changed modules."
    echo "h     Print this Help."
    echo
}

while getopts 'cp:' flag; do
  case "${flag}" in
    c) CHANGE_ONLY='true' ;;
    *) print_usage
       exit 1 ;;
  esac
done

echo "Running test-unit with the following options:"
echo "CHANGE_ONLY: $CHANGE_ONLY"

echo $@

for f in $GO_MOD_DIRS; do
    if [ -d $f ]; then
        cd $f
        CHANGES=""
        if [ "$CHANGE_ONLY" = "true" ]; then
            CHANGES=$(git diff --name-only HEAD..origin/main ../)
            if [ "$CHANGES" = "" ]; then
                echo "No changes in $f, skipping."
                cd -
                continue
            fi
        fi
        echo "Running unit tests $f"
        go test -race -v ./...
        cd -
    fi
done