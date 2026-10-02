# `cloudquery upgrade check` fixtures

`TestUpgradeCheckRFCCases` runs the RFC upgrade cases against the real PostgreSQL destination without downloading plugins.

- `<case>.yml`: the configuration for each case.
- `<case>.txt`: the expected report.
- `<case>.json`: the expected report with `--output json`. These are also the reference for the JSON shape.
- `tables/<source>-<version>.json`: the selected tables of each source version, as returned by `GetTables`, keyed by table name (base64 Arrow schemas).
- `postgresql.mod` and `postgresql.sum`: the `go.mod` and `go.sum` of a module that pins the PostgreSQL destination. The test copies them to a temporary directory, builds the destination there and runs it with `registry: local`, so the CLI module does not depend on the destination module. They are not named `go.mod` and `go.sum` because tools that walk the CLI directory for modules, like `make gen-licenses`, would pick them up.

To record the tables again from the real source plugins and rewrite the expected reports (requires `cloudquery login`):

```bash
go test ./cmd -run TestUpgradeCheckRFCCases -update-upgrade-check-fixtures
```

To test against another PostgreSQL destination version, from the `cli` directory:

```bash
fixtures="$PWD/cmd/testdata/upgrade-check"
cd "$(mktemp -d)"
cp "$fixtures/postgresql.mod" go.mod && cp "$fixtures/postgresql.sum" go.sum
GOWORK=off go get -tool github.com/cloudquery/cloudquery/plugins/destination/postgresql/v8@<version or commit>
GOWORK=off go mod tidy
cp go.mod "$fixtures/postgresql.mod" && cp go.sum "$fixtures/postgresql.sum"
```
