# GCP Source Plugin Configuration Reference

## GCP Spec

This is the top level spec used by GCP Source Plugin

- `project_ids` ([]string) (default: empty. will use all available projects available to the current authenticated account)

  specify specific projects to connect to.

- `service_account_key_json` (string)

  GCP service account key content. Using service accounts is not recommended, but if it is used it is better to use env variable expansion

- `folder_ids` ([]string) (default: empty).
  
  cloudquery will `sync` from all the projects in the specified folders, recursively. `folder_ids` must be of the format
  `folders/<folder_id>` or `organizations/<organization_id>`. This feature requires the `resourcemanager.folders.list` permission. 
  By default cloudquery will also `sync` from subfolders recursively (up to depth 100) - to reduce this, set `folder_recursion_depth` to a lower value (or 0 to disable recursion completely).

- `folder_recursion_depth` (int) (default: 100).
  
  the maximum depth to recurse into subfolders. 0 means no recursion (only the top-level projects in folders will be used for `sync`).
