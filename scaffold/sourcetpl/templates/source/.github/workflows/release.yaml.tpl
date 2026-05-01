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
        uses: actions/checkout@de0fac2e4500dabe0009e67214ff5f5447ce83dd # v6
        with:
          fetch-depth: 0

      - name: Set up Go
        uses: actions/setup-go@4a3601121dd01d1626a1e23e37211e3254c1c06c # v6
        with:
          go-version-file: go.mod

      # Needed for shell escape
      - name: Use Node.js LTS
        uses: actions/setup-node@53b83947a5a98c8d113130e565377fae1a50d02f # v6
        with:
          node-version: 'lts/*'

      - name: Install shell escape
        run: |
              npm install shell-escape@0.2.0

      - name: Get Release Notes
        id: release-notes
        uses: actions/github-script@3a2844b7e9c422d3c10d287c895573f7108da1b3 # v9
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
        uses: jacobtomlinson/gha-find-replace@b19bfcb2a015af55fd6e160d1d1987e887f8c163
        with:
          find: "(?i)version_source_{{.Name}}"
          replace: ${{"{{"}} steps.semver_parser.outputs.fullversion {{"}}"}}
          include: ./docs/*.md

      - name: Run package command
        run: |
          rm -rf docs/tables.md
          go run main.go package -m ${{"{{"}} steps.release-notes.outputs.result {{"}}"}} ${{"{{"}} steps.semver_parser.outputs.fullversion {{"}}"}} .

      - name: Setup CloudQuery
        uses: cloudquery/setup-cloudquery@4a3af61f7d8c362d8d152a4a17053ed2aaa7180d # v4
        with:
          version: v5.0.1

      - name: Publish plugin to hub
        # See https://www.cloudquery.io/docs/deployment/generate-api-key for instructions how to generate this key.
        env:
          CLOUDQUERY_API_KEY: ${{"{{"}} secrets.CLOUDQUERY_API_KEY {{"}}"}}
        run: |
          cloudquery plugin publish --finalize