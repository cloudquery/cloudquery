# `cloudquery upgrade check` fixtures

`TestUpgradeCheckRFCCases` runs the RFC upgrade cases against the real PostgreSQL destination without downloading plugins.

- `<case>.yml`: the configuration for each case.
- `<case>.txt`: the expected report.
- `tables/<source>-<version>.json`: the selected tables of each source version, as returned by `GetTables`, keyed by table name (base64 Arrow schemas).
- `postgresql/`: a module that pins the PostgreSQL destination. The test builds it and runs it with `registry: local`, so the CLI module does not depend on the destination module.

To record the tables again from the real source plugins and rewrite the expected reports (requires `cloudquery login`):

```bash
go test ./cmd -run TestUpgradeCheckRFCCases -update-upgrade-check-fixtures
```

To test against another PostgreSQL destination version:

```bash
cd cmd/testdata/upgrade-check/postgresql
GOWORK=off go get -tool github.com/cloudquery/cloudquery/plugins/destination/postgresql/v8@<version or commit>
GOWORK=off go mod tidy
```
