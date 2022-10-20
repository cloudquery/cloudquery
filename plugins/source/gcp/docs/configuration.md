# GCP Source Plugin Configuration Reference

## GCP Spec

This is the top level spec used by GCP Source Plugin

- `project_ids` ([]string) (default: empty. will use all available projects available to the current authenticated account)

  specify specific projects to connect to.

- `service_account_key_json` (string)

  GCP service account key content. Using service accounts is not recommended, but if it is used it is better to use env variable expansion

- `project_filter` (string)

  Query to limit the projects that are synced. Default value is `lifecycleState=ACTIVE`. For syntax and example queries refer to API Reference [here](https://cloud.google.com/resource-manager/reference/rest/v1/projects/list#google.cloudresourcemanager.v1.Projects.ListProjects)
