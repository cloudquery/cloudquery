# GCP + Kubernetes (GKE) + PostgreSQL

```yaml
kind: source
spec:
  name: gcp
  path: "cloudquery/gcp"
  version: "v1.0.1"
  destinations: ["postgresql"]
---
kind: source
spec:
  name: k8s
  path: "cloudquery/k8s"
  version: "v2.0.1"
  destinations: ["postgresql"]
---
kind: destination
spec:
  name: "postgresql"
  path: "cloudquery/postgresql"
  version: "v0.3.0"
  write_mode: "overwrite" # overwrite, append
  spec:
    connection_string: "postgresql://{CQ_PG_USER}:{CQ_PG_PASS}@localhost:5432/postgres?sslmode=disable"
```

Kubernetes users may see the following message when running the K8s plugin on GKE Clusters:
```
WARNING: the gcp auth plugin is deprecated in v1.22+, unavailable in v1.26+; use gcloud instead.
```

As part of an initiative to remove platform specific code from Kubernetes, authentication will begin to be delegated to authentication plugins, starting in version 1.26.

## What does this mean for CloudQuery users?

CloudQuery does not use any specific resources which hinder the upgrade. 

### Install
The easiest way to upgrade, is to install `gke-gcloud-auth-plugin` from `gcloud components` on Mac or Windows:

```zsh
gcloud components install gke-gcloud-auth-plugin
```

and apt on Deb based systems:
```bash
sudo apt-get install google-cloud-sdk-gke-gcloud-auth-plugin
```

### Verify

Mac or Linux:
```
gke-gcloud-auth-plugin --version 
```

Windows:
```
gke-gcloud-auth-plugin.exe --version
```

### Switch authentication methods
Set the flag:
```sh
export USE_GKE_GCLOUD_AUTH_PLUGIN=True
```

Update components:
```sh
gcloud components update
```

Force credential update:
```
gcloud container clusters get-credentials {$CLUSTER_NAME}
```

Now you should be able to use `kubectl` as normal, and you
should no longer see the warning in the CloudQuery output.

For more information, read [Google's press release](https://cloud.google.com/blog/products/containers-kubernetes/kubectl-auth-changes-in-gke).
