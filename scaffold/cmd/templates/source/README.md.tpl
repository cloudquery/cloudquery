# CloudQuery {{.Name}} Source Plugin

[![test](https://github.com/{{.Org}}/cq-source-{{.Name}}/actions/workflows/test.yaml/badge.svg)](https://github.com/{{.Org}}/cq-source-{{.Name}}/actions/workflows/test.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/{{.Org}}/cq-source-{{.Name}})](https://goreportcard.com/report/github.com/{{.Org}}/cq-source-{{.Name}})

A {{.Name}} source plugin for CloudQuery that loads data from {{.Name}} to any database, data warehouse or data lake supported by [CloudQuery](https://www.cloudquery.io/), such as PostgreSQL, BigQuery, Athena, and many more.

## Links

 - [CloudQuery Quickstart Guide](https://www.cloudquery.io/docs/quickstart)
 - [Supported Tables](docs/tables/README.md)


## Configuration

The following source configuration file will sync to a PostgreSQL database. See [the CloudQuery Quickstart](https://www.cloudquery.io/docs/quickstart) for more information on how to configure the source and destination.

```yaml
kind: source
spec:
  name: "{{.Name}}"
  path: "{{.Org}}/{{.Name}}"
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