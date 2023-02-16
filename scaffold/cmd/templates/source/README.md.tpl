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

### Release a new version

1. Follow [this link](https://github.com/{{.Org}}/cq-source-{{.Name}}/releases/new) to draft a new release.
2. Click `Choose a tag` and enter the new version number:
   ![image](https://user-images.githubusercontent.com/26760571/219360662-0ad1f83d-84c9-47c8-afb9-fe774ce03dcc.png)
3. Click `Create new tag: <version> on publish` assuming it's a new tag.
4. Click `Generate release notes` to automatically generate release notes.
5. Click `Publish release` to publish the release.

> Once the tag is pushed, a new GitHub Actions workflow will be triggered to build and upload the release binaries to the release
