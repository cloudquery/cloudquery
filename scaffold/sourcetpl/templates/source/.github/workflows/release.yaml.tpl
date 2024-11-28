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
        uses: booxmedialtd/ws-action-parse-semver@7784200024d6b3fc01253e617ec0168daf603de3
        with:
          input_string: ${{"{{"}}github.ref_name{{"}}"}}
          # If you're tagging releases without `v` prefix, remove `version_extractor_regex` param.
          # See doc for more details:
          # https://github.com/marketplace/actions/parse-semver#version_extractor_regex
          version_extractor_regex: '^v(.*)$'

      - name: Checkout
        uses: actions/checkout@v4
        with:
          fetch-depth: 0

      - name: Set up Go
        uses: actions/setup-go@v5
        with:
          go-version-file: go.mod

      # Needed for shell escape
      - name: Use Node.js LTS
        uses: actions/setup-node@v4
        with:
          node-version: 'lts/*'

      - name: Install shell escape
        run: |
              npm install shell-escape@0.2.0

      - name: Get Release Notes
        id: release-notes
        uses: actions/github-script@v7
        env:
          PRERELEASE: ${{"{{"}} steps.semver_parser.outputs.prerelease {{"}}"}}
        with:
          result-encoding: string
          script: |
            const shellescape = require('shell-escape');
            const { PRERELEASE } = process.env;
            if (PRERELEASE) {
             return shellescape(["This is a pre-release version of the plugin and should be used for testing purposes only"])
            }
            const { data } = await github.rest.repos.getReleaseByTag({
              owner: "{{.Org}}",
              repo: context.repo.repo,
              tag: context.ref.replace('refs/tags/', ''),
            });
            return shellescape([data.body]);

      - name: Find and Replace
        uses: jacobtomlinson/gha-find-replace@f1069b438f125e5395d84d1c6fd3b559a7880cb5
        with:
          find: "(?i)version_source_{{.Name}}"
          replace: ${{"{{"}} steps.semver_parser.outputs.fullversion {{"}}"}}
          include: ./docs/*.md

      - name: Run package command
        run: |
          rm -rf docs/tables.md
          go run main.go package -m ${{"{{"}} steps.release-notes.outputs.result {{"}}"}} ${{"{{"}} steps.semver_parser.outputs.fullversion {{"}}"}} .

      - name: Setup CloudQuery
        uses: cloudquery/setup-cloudquery@v4
        with:
          version: v5.0.1

      - name: Publish plugin to hub
        # See https://docs.cloudquery.io/docs/deployment/generate-api-key for instructions how to generate this key.
        env:
          CLOUDQUERY_API_KEY: ${{"{{"}} secrets.CLOUDQUERY_API_KEY {{"}}"}}
        run: |
          cloudquery plugin publish --finalize