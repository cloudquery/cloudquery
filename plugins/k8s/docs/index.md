## K8S Provider

The CloudQuery K8S provider pulls configuration out of Kubernetes, normalizes them and stores them in PostgreSQL database.

### Install

```shell
cloudquery init k8s
```

### Authentication

Similar to how `kubectl` works, `cloudquery` depends on a kubernetes configuration file to connect to a 
kubernetes cluster and `fetch` its information. By default, `cloudquery` uses the default kubernetes configuration
file (`~/.kube/config`). You can also specify a different configuration by setting the `KUBECONFIG` environment variable before running `cloudquery fetch`.

```bash
export KUBECONFIG=<PATH_TO_YOUR_CONFIG_FILE>
```

### Configuration

At the moment provider does not have any configuration.
