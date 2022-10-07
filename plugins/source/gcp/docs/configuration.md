# GCP Source Plugin Configuration Reference

## GCP Spec

This is the top level spec used by GCP Source Plugin

- `project_ids` ([]string) (default: empty. will use all available projects available to the current authenticated account)

  specify specific projects to connect to.

- `service_account_key_json` (string)

  GCP service account key content. Using service accounts is not recommended, but if it is used it is better to use env variable expansion
