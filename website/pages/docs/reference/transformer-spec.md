---
title: Transformer Spec Reference
description: Reference for the transformer spec CloudQuery configuration object.
---

# Transformer Spec Reference

This goes through all the available options for the transformer integration `spec` object.

## Example

This example configures the `basic` integration to rename all tables by adding a `cq_` prefix before reaching the destination.

```yaml copy
kind: transformer
spec:
  name: "basic"
  path: "cloudquery/basic"
  registry: "cloudquery"
  version: "VERSION_TRANSFORMER_BASIC"

  spec:
    transformations:
      - kind: change_table_names
        tables: ["*"]
        new_table_name_template: "cq_{{.OldName}}"
```

## Spec

### name

(`string`, required)

Name of the integration. If you have multiple transformer integrations, this must be unique.

The name field may be used to uniquely identify a particular transformer configuration. For example, if you have two configs for the basic integration for transforming a source table differently in each of two different destination databases, one may be named `basic-1` and the other `basic-2`. In this case, the `path` option below must be used to specify the download path for the integration.

### registry

(`string`, optional, default: `cloudquery`, available: `github`, `cloudquery`, `local`, `grpc`, `docker`)

- `cloudquery`: CloudQuery will look for and download the integration from the official CloudQuery registry, and then execute it.
- `local`: CloudQuery will execute the integration from a local path.
- `grpc`: mostly useful in debug mode when integration is already running in a different terminal, CloudQuery will connect to the gRPC integration server directly without spawning the process.

### path

(`string`, required)

Configures how to retrieve the integration. The contents depend on the value of `registry` (`github` by default).

- For integrations hosted on GitHub, `path` should be of the form `"<org>/<repository>"`. For official integrations, should be `cloudquery/<integration-name>`.
- For integrations that are located in the local filesystem, `path` should a filesystem path to the integration binary.
- To connect to a running integration via `grpc` (mostly useful for debugging), `path` should be the host-port of the integration (e.g. `localhost:7777`).

### version

(`string`, required)

`version` must be a valid [SemVer](https://semver.org/), e.g. `vMajor.Minor.Patch`. You can find all official integration versions under [our GitHub releases page](https://github.com/cloudquery/cloudquery/releases), and for community integrations you can find it in the relevant community repository.

### spec

(`object`, optional)

Plugin specific configurations. Visit [transformers](https://hub.cloudquery.io/addons/transformation) documentation for more information.
