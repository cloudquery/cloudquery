# K8s Source Plugin Configuration Reference

## Example

This example connects a single k8s context to a Postgres destination. The (top level) source spec section is described in the [Source Spec Reference](https://www.cloudquery.io/docs/reference/source-spec).

```yml
kind: source
spec:
  # Source spec section
  name: k8s
  path: cloudquery/k8s
  version: "v2.4.1" # latest version of k8s plugin
  tables: ["*"]
  destinations: ["postgresql"]

  spec:
    contexts: ["context"]
```

## K8s Spec

This is the (nested) spec used by K8s Source Plugin

- `contexts` ([]string) (default: empty. Will use the default context from K8s's config file)

  Specify specific K8s contexts to connect to. Specifying `*` will connect to all contexts available in
  the K8s config file (usually `~/.kube/config`).