# CloudQuery Onfleet Source Plugin

[![test](https://github.com/cloudquery/cq-source-onfleet/actions/workflows/test.yaml/badge.svg)](https://github.com/cloudquery/cq-source-onfleet/actions/workflows/test.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/cloudquery/cq-source-onfleet)](https://goreportcard.com/report/github.com/cloudquery/cq-source-onfleet)

A Onfleet source plugin for CloudQuery that loads data from Onfleet to any database, data warehouse or data lake supported by [CloudQuery](https://www.cloudquery.io/), such as PostgreSQL, BigQuery, Athena, and many more.

## Links

 - [CloudQuery Quickstart Guide](https://www.cloudquery.io/docs/quickstart)
 - [Supported Tables](docs/tables/README.md)


## Configuration

The following source configuration file will sync to a PostgreSQL database. See [the CloudQuery Quickstart](https://www.cloudquery.io/docs/quickstart) for more information on how to configure the source and destination.

```yaml
kind: source
spec:
  name: "onfleet"
  path: "cloudquery/onfleet"
  version: "${VERSION}"
  destinations:
    - "postgresql"
  spec:
    # plugin spec section
```

## Development

### Run tests

```bash
make test
```

### Run linter

```bash
make lint
```

### Generate docs

```bash
make gen-docs
```