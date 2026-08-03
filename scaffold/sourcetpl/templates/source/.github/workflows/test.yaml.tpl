name: test

on:
  push:
    branches:
      - main
  pull_request:
    branches: [main]

jobs:
  test:
    timeout-minutes: 30
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@3d3c42e5aac5ba805825da76410c181273ba90b1 # v7
        with:
          fetch-depth: 2
      - name: Set up Go 1.x
        uses: actions/setup-go@b7ad1dad31e06c5925ef5d2fc7ad053ef454303e # v7
        with:
          go-version-file: go.mod
      - name: golangci-lint
        uses: golangci/golangci-lint-action@ba0d7d2ec06a0ea1cb5fa41b2e4a3ab91d21278a # v9
        with:
          version: v2.12.2
          args: --timeout=10m
      - name: Get dependencies
        run: go get -t -d ./...
      - name: Build
        run: go build .
      - name: Test
        run: make test
      - name: gen
        if: github.event_name == 'pull_request'
        run: make gen
      - name: Fail if generation updated files
        if: github.event_name == 'pull_request'
        run: test "$(git status -s | wc -l)" -eq 0 || (git status -s; exit 1)
