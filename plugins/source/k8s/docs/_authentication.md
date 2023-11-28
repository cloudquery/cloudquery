Similar to how `kubectl` works, `cloudquery` depends on a Kubernetes configuration file to connect to a Kubernetes cluster and `sync` its information.
By default, `cloudquery` uses the default Kubernetes configuration file (`~/.kube/config`).
You can also specify a different configuration by setting the `KUBECONFIG` environment variable before running `cloudquery sync`.

```bash copy
export KUBECONFIG="<PATH_TO_YOUR_CONFIG_FILE>"
```

### Kubernetes Service Account

If `cloudquery` is running in a pod of the Kubernetes cluster, the Kubernetes Service Account can be used for direct authentication. To use the Kubernetes Service Account for direct authentication, a cluster role with all get and list privileges will need to be used.

The below command creates a new cluster role with `get` and `list` privileges.

```bash copy
kubectl apply -f - <<EOF
apiVersion: rbac.authorization.k8s.io/v1
kind:       ClusterRole
metadata:
  name: cloudquery-cluster-read
rules:
- apiGroups:
  - '*'
  resources:
  - '*'
  verbs:
  - get
  - list
- nonResourceURLs:
  - '*'
  verbs:
  - get
  - list
EOF
```

Next, the cluster role and service account will need to be linked via a cluster role binding.
The following creates a cluster role binding for the role we created above and the service account for the `cloudquery` pod.

```bash copy
kubectl apply -f - <<EOF
apiVersion: rbac.authorization.k8s.io/v1
kind:       ClusterRoleBinding
metadata:
  name: cloudquery-cluster-read-binding
subjects:
- kind: ServiceAccount
  name: cloudquery-sa
roleRef:
  kind: ClusterRole
  name: cloudquery-cluster-read
EOF
```
