#!/usr/bin/env bash

set -e

for f in terraform/*; do
    if [ -d $f ]; then
        cd $f/prod
        echo "Running terraform init,validate in $f"
        terraform init
        terraform validate
        cd -
    fi
done
