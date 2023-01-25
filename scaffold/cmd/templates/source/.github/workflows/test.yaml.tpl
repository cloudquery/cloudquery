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
      - uses: actions/checkout@v3
        with:
          fetch-depth: 2
      - name: Set up Go 1.x
        uses: actions/setup-go@v3
        with:
          go-version-file: go.mod
          cache: true
      - name: golangci-lint
        uses: golangci/golangci-lint-action@v3
        with:
          version: v1.50.1
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
