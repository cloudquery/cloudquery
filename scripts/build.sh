#!/usr/bin/env bash
set -e

CHANGE_ONLY="false"
GO_MOD_DIRS=$(find . -type f -name "go.mod" | xargs -n1 dirname)

print_usage() {
    echo "Runs go build on all modules. Output is written to ./bin"
    echo
    echo "Syntax: ./build.sh [-c]"
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

echo "Running go build with the following options:"
echo "CHANGE_ONLY: $CHANGE_ONLY"
echo "OUTPUT_DIR: $PWD/bin"

echo $@

for f in $GO_MOD_DIRS; do
    if [ -d $f ]; then
        BASEDIR=$PWD
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
        echo "Running go build on $f"
        go build -v -o $BASEDIR/bin/ .
        cd -
    fi
done