# K8s Source Plugin

The K8s Source plugin for CloudQuery extracts configuration from a variety of K8s APIs.

## Libraries in Use

- https://pkg.go.dev/k8s.io/api

## Authentication

Similar to how `kubectl` works, `cloudquery` depends on a Kubernetes configuration file to connect to a Kubernetes cluster and `sync` its information. By default, `cloudquery` uses the default Kubernetes configuration
file (`~/.kube/config`). You can also specify a different configuration by setting the `KUBECONFIG` environment variable before running `cloudquery sync`.

```bash
export KUBECONFIG=<PATH_TO_YOUR_CONFIG_FILE>
```