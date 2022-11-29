# GCP + Kubernetes (GKE)

```yaml copy
kind: source
spec:
  name: gcp
  path: "cloudquery/gcp"
  version: "VERSION_SOURCE_GCP"
  destinations: ["<destination>"]
---
kind: source
spec:
  name: k8s
  path: "cloudquery/k8s"
  version: "VERSION_SOURCE_K8S"
  destinations: ["<destination>"]
```

Kubernetes users may see the following message when running the K8s plugin on GKE Clusters:

```bash copy
WARNING: the gcp auth plugin is deprecated in v1.22+, unavailable in v1.26+; use gcloud instead.
```

As part of an initiative to remove platform specific code from Kubernetes, authentication will begin to be delegated to authentication plugins, starting in version 1.26.

## What does this mean for CloudQuery users?

CloudQuery does not use any specific resources which hinder the upgrade.

### Install

The easiest way to upgrade, is to install `gke-gcloud-auth-plugin` from `gcloud components` on Mac or Windows:

```bash copy
gcloud components install gke-gcloud-auth-plugin
```

and apt on Deb based systems:

```bash copy
sudo apt-get install google-cloud-sdk-gke-gcloud-auth-plugin
```

### Verify

Mac or Linux:

```bash copy
gke-gcloud-auth-plugin --version
```

Windows:

```bash copy
gke-gcloud-auth-plugin.exe --version
```

### Switch authentication methods

Set the flag:

```bash copy
export USE_GKE_GCLOUD_AUTH_PLUGIN=True
```

Update components:

```bash copy
gcloud components update
```

Force credential update:

```bash copy
gcloud container clusters get-credentials {$CLUSTER_NAME}
```

Now you should be able to use `kubectl` as normal, and you
should no longer see the warning in the CloudQuery output.

For more information, read [Google's press release](https://cloud.google.com/blog/products/containers-kubernetes/kubectl-auth-changes-in-gke).