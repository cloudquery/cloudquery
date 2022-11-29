# GCP Source Plugin Configuration Reference

## Example

This example connects a single GCP project to a Postgres destination. The (top level) source spec section is described in the [Source Spec Reference](/docs/reference/source-spec).

```yaml
kind: source
spec:
  # Source spec section
  name: "gcp"
  path: "cloudquery/gcp"
  version: "VERSION_SOURCE_GCP"
  destinations: ["postgresql"]

  spec:
    # GCP Spec section described below
    project_ids: ["my-project"]
```

## GCP Spec

This is the (nested) spec used by GCP Source Plugin

- `project_ids` ([]string) (default: empty. will use all available projects available to the current authenticated account)

  Specify specific projects to connect to. If either `folder_ids` or `project_filter` is specified, these projects will be fetched in addition
  to the projects from the folder/filter.

- `service_account_key_json` (string) (default: empty).

  GCP service account key content. Using service accounts is not recommended, but if it is used it is better to use env variable expansion

- `folder_ids` ([]string) (default: empty).
  
  cloudquery will `sync` from all the projects in the specified folders, recursively. `folder_ids` must be of the format
  `folders/<folder_id>` or `organizations/<organization_id>`. This feature requires the `resourcemanager.folders.list` permission. 
  By default cloudquery will also `sync` from subfolders recursively (up to depth 100) - to reduce this, set `folder_recursion_depth` to a lower value (or 0 to disable recursion completely).
  Mutually exclusive with `project_filter`.

- `folder_recursion_depth` (int) (default: 100).
  
  the maximum depth to recurse into subfolders. 0 means no recursion (only the top-level projects in folders will be used for `sync`).

- `project_filter` (string) (default: empty).

  A filter to determine the projects that are synced. For instance, to only sync projects where the name starts with `how-`,
  set `project_filter` to `name:how-*`. Another example is: `"name:how-* OR name:test-*"`. For syntax and example queries refer to API Reference [here](https://cloud.google.com/resource-manager/reference/rest/v1/projects/list#google.cloudresourcemanager.v1.Projects.ListProjects).
  Mutually exclusive with `folder_ids`.