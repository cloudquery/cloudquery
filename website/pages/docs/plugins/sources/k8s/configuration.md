# K8s Source Plugin Configuration Reference

The K8s source plugin connects to a Kubernetes cluster, fetches resources and loads it into any supported CloudQuery destination (e.g. PostgreSQL, BigQuery, Snowflake, and [more](/docs/plugins/destinations/overview)).

## Example

This example connects a single k8s context to a Postgres destination. The (top level) source spec section is described in the [Source Spec Reference](/docs/reference/source-spec).

import Configuration from "./_configuration.mdx";
<Configuration/>

## K8s Spec

This is the (nested) spec used by K8s Source Plugin

- `contexts` ([]string) (default: empty. Will use the default context from K8s's config file)

  Specify specific K8s contexts to connect to. Specifying `*` will connect to all contexts available in
  the K8s config file (usually `~/.kube/config`).