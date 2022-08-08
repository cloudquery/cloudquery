#!/usr/bin/env bash
set -e

CHANGE_ONLY="false"
GO_MOD_DIRS=$(find . -type f -name "go.mod" | xargs -n1 dirname)

print_usage() {
    echo "Runs go build on all modules."
    echo
    echo "Syntax: ./build.sh [-c]"
    echo "options:"
    echo "c     Run only on changed modules."
    echo "h     Print this Help."
    echo
}

while getopts 'c:' flag; do
  case "${flag}" in
    c) CHANGE_ONLY='true' ;;
    *) print_usage
       exit 1 ;;
  esac
done

echo "Running go build with the following options:"
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
        echo "Running go build $f"
        go build -v .
        cd -
    fi
done