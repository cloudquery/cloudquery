#!/usr/bin/env bash

set -e

for f in terraform/*; do
    if [ -d $f ]; then
        cd $f/prod
        echo "Running terraform init,apply in $f"
        terraform init -no-color
        terraform apply -no-color
        cd -
    fi
done
