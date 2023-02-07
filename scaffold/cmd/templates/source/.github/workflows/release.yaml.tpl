name: release
on:
  push:
    tags:
      - 'v*.*.*'
env:
  CGO_ENABLED: 0

jobs:
  release-binary:
    runs-on: ubuntu-latest
    steps:
      # This fails for invalid semver strings
      - name: Parse semver string
        id: semver_parser
        uses: booxmedialtd/ws-action-parse-semver@966a26512c94239a00aa10b1b0c196906f7e1909
        with:
          input_string: ${{"{{"}}github.ref_name{{"}}"}}
      - name: Checkout
        uses: actions/checkout@v3
        with:
          fetch-depth: 0
      - name: Set up Go
        uses: actions/setup-go@v3
        with:
          go-version-file: go.mod
          cache: true
      - name: Run GoReleaser Dry-Run
        uses: goreleaser/goreleaser-action@v3
        with:
          version: latest
          args: release --clean --skip-validate --skip-publish --skip-sign
      - name: Run GoReleaser
        uses: goreleaser/goreleaser-action@v3
        with:
          version: latest
          args: release --clean --skip-sign
        env:
          GITHUB_TOKEN: ${{"{{"}} secrets.GITHUB_TOKEN {{"}}"}}
