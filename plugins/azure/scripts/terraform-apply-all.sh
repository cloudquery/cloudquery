#!/usr/bin/env bash

set -e
set -x

for f in terraform/*; do
    if [ -d $f ]; then
        cd $f/prod
        terraform init
        terraform apply -no-color
        cd -
    fi
done
