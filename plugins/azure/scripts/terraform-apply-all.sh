#!/usr/bin/env bash

for f in terraform/*; do
    if [ -d $f ]; then
        cd $f/prod
        terraform apply -no-color
        cd -
    fi
done
