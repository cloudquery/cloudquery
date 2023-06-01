# GCP Source Plugin Configuration Reference

## Example

This example connects a single GCP project to a Postgres destination. The (top level) source spec section is described in the [Source Spec Reference](/docs/reference/source-spec).

```yaml copy
kind: source
spec:
  # Source spec section
  name: "gcp"
  path: "cloudquery/gcp"
  version: "VERSION_SOURCE_GCP"
  tables: ["gcp_storage_buckets"]
  destinations: ["postgresql"]

  spec:
    # GCP Spec section described below
    project_ids: ["my-project"]
```

## GCP Spec

This is the (nested) spec used by GCP Source Plugin

- `project_ids` ([]string) (default: empty. will use all projects available to the current authenticated account)

  Specify projects to connect to. If either `folder_ids` or `project_filter` is specified, these projects will be synced in addition
  to the projects from the folder/filter.

- `service_account_key_json` (string) (default: empty).

  GCP service account key content. Using service accounts is not recommended, but if it is used it is better to use [environment or file variable substitution](/docs/advanced-topics/environment-variable-substitution).

- `folder_ids` ([]string) (default: empty).
  
  cloudquery will `sync` from all the projects in the specified folders, recursively. `folder_ids` must be of the format
  `folders/<folder_id>` or `organizations/<organization_id>`. This feature requires the `resourcemanager.folders.list` permission. 
  By default cloudquery will also `sync` from subfolders recursively (up to depth 100) - to reduce this, set `folder_recursion_depth` to a lower value (or 0 to disable recursion completely).
  Mutually exclusive with `project_filter`.

- `folder_recursion_depth` (int) (default: 100).
  
  the maximum depth to recurse into subfolders. 0 means no recursion (only the top-level projects in folders will be used for `sync`).

- `project_filter` (string) (default: empty).

  A filter to determine the projects that are synced, mutually exclusive with `folder_ids`. For instance, to only sync projects where the name starts with `how-`,
  set `project_filter` to `name:how-*`. 
  
  More examples:
    - `"name:how-* OR name:test-*"` matches projects starting with `how-` or `test-`
    - `"NOT name:test-*"` matches all projects *not* starting with `test-`
  
  For syntax and example queries refer to API References [here](https://cloud.google.com/resource-manager/reference/rest/v1/projects/list#google.cloudresourcemanager.v1.Projects.ListProjects) and [here](https://cloud.google.com/sdk/gcloud/reference/topic/filters).
  

- `organization_ids` ([]string) (default: empty. will use all organizations available to the current authenticated account)

  Specify organizations to use when syncing organization level resources (e.g. [folders](https://github.com/cloudquery/cloudquery/blob/0e384a84d1c9545b24c2eda9af00f111bab79c36/plugins/source/gcp/resources/services/resourcemanager/folders_fetch.go#L23) or [security findings](https://github.com/cloudquery/cloudquery/blob/0e384a84d1c9545b24c2eda9af00f111bab79c36/plugins/source/gcp/resources/services/securitycenter/organization_findings.go#L43)). If `organization_filter` is specified, these organizations will be used in addition to the organizations from the filter.

- `organization_filter` (string) (default: empty).

  A filter to determine the organizations to use when syncing organization level resources (e.g. [folders](https://github.com/cloudquery/cloudquery/blob/0e384a84d1c9545b24c2eda9af00f111bab79c36/plugins/source/gcp/resources/services/resourcemanager/folders_fetch.go#L23) or [security findings](https://github.com/cloudquery/cloudquery/blob/0e384a84d1c9545b24c2eda9af00f111bab79c36/plugins/source/gcp/resources/services/securitycenter/organization_findings.go#L43)). For instance, to use only organizations from the `cloudquery.io` domain, set `organization_filter` to `domain:cloudquery.io`. For syntax and example queries refer to API Reference [here](https://cloud.google.com/resource-manager/reference/rest/v1/organizations/search#google.cloudresourcemanager.v1.SearchOrganizationsRequest).

- `backoff_delay` (int) (default: 30).
  If specified APIs will be retried with exponential backoff if they are rate limited. This is the max delay (in seconds) between retries.

- `backoff_retries` (int) (default: 0).
  If specified APIs will be retried with exponential backoff if they are rate limited. This is the max number of retries.

- `enabled_services_only` (bool) (default: false).
If enabled CloudQuery will skip any resources that belong to a service that has been disabled or not been enabled. If you use this option on a large organization (with more than 500 projects) you should also set the `backoff_retries` to a value greater than `0` otherwise you may hit the API rate limits. In `v9.0.0` and greater if an error is returned then CloudQuery will assume that all services are enabled and will continue to attempt to sync all specified tables rather than just ending the sync.

- `discovery_concurrency` (int) (default: 100).
  The number of concurrent requests that CloudQuery will make to resolve the enabled services. This is only used when `enabled_services_only` is enabled.

## GCP + Kubernetes (GKE)

```yaml copy
kind: source
spec:
  name: gcp
  path: "cloudquery/gcp"
  version: "VERSION_SOURCE_GCP"
  tables: ["gcp_container_clusters"]
  destinations: ["<destination>"]
---
kind: source
spec:
  name: k8s
  path: "cloudquery/k8s"
  version: "VERSION_SOURCE_K8S"
  tables: ["*"]
  destinations: ["<destination>"]
```

Kubernetes users may see the following message when running the K8s plugin on GKE Clusters:

```bash copy
WARNING: the gcp auth plugin is deprecated in v1.22+, unavailable in v1.26+; use gcloud instead.
```

As part of an initiative to remove platform specific code from Kubernetes, authentication will begin to be delegated to authentication plugins, starting in version 1.26.

### What does this mean for CloudQuery users?

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
