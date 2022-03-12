#!/usr/bin/env bash

set -e

for f in terraform/*; do
    if [ -d $f ]; then
        cd $f/prod
        terraform init
        terraform apply -no-color
        cd -
    fi
done
