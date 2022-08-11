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
By default cloudquery fetches data from default context of the kubernetes config. Context to fetch can be selected by setting contexts variable of provider's `configuration` block in `config.hcl`. 
Example of context selection:

```yml title="cloudquery.yml"
providers:
  - name: k8s
   #  configuration:
      # Optional. Set contexts that you want to fetch. If it is not given then all contexts from config are iterated over.
      # contexts:
        # - "<YOUR_CONTEXT_NAME1>"
        # - "<YOUR_CONTEXT_NAME2>"
    resources:
      - "*"
```

To fetch all the contexts set `contexts: "*"`
