before:
  hooks:
    - go mod download
builds:
  - flags:
      - -buildmode=exe
    env:
      - CGO_ENABLED=0
      - GO111MODULE=on
    ldflags:
      - -s -w -X github.com/{{.Org}}/cq-source-{{.Name}}/plugin.Version={{"{{"}}.Version{{"}}"}}
    goos:
      - windows
      - linux
      - darwin
    goarch:
      - amd64
      - arm64
    ignore:
      - goos: windows
        goarch: arm64
archives:
  - name_template: "{{"{{"}} .Binary {{"}}"}}_{{"{{"}} .Os {{"}}"}}_{{"{{"}} .Arch {{"}}"}}"
    format: zip
checksum:
  name_template: "checksums.txt"
changelog:
  sort: asc
  filters:
    exclude:
      - "^docs:"
      - "^test:"

release:
  prerelease: auto
