#!/usr/bin/env bash

if [[ $(git --no-pager log --no-merges --first-parent --patch-with-stat HEAD^..HEAD -- . ../.github/workflows/broken_links.yml) ]]; then
    echo "there are changes to the website or the broken links workflow in the last commit"
    exit 1
else
    echo "there are no changes to the website or the broken links workflow in the last commit"
    exit 0
fi