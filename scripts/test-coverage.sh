#!/usr/bin/env bash
MAKEFILE_DIRS=$(find . -type f -name "Makefile" ! -path "**/node_modules/*" ! -path "**/vendor/*" | xargs -n1 dirname)

for f in $MAKEFILE_DIRS; do
    if [ -d $f ]; then
        cd "$f"
        echo "Running coverage $f"
        exit_code=0
        make coverage || exit_code=$?
        if [ $exit_code -ne 0 ]; then
            echo "Coverage failed for $f"
            exit $exit_code
        fi
        cd -
    fi
done