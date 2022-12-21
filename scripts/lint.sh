#!/usr/bin/env bash
MAKEFILE_DIRS=$(find . -type f -name "Makefile" | xargs -n1 dirname)

for f in $MAKEFILE_DIRS; do
    if [ -d $f ]; then
        cd "$f"
        echo "Running lint $f"
        make lint
        cd -
    fi
done