# K8s Source Plugin Configuration Reference

## K8s Spec

This is the top level spec used by K8s Source Plugin

- `contexts` ([]string) (default: empty. Will use the default context from K8s's config file)

  Specify specific K8s contexts to connect to. Specifying `*` will connect to all contexts available in
  the K8s config file (usually `~/.kube/config`).