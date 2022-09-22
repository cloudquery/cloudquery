# GCP Source Plugin

The GCP Source plugin for CloudQuery extracts configuration from a variety of GCP APIs and loads them into any supported CloudQuery destinations (e.g. PostgreSQL).

## Links

- Tables: [./docs/tables/README.md](./docs/tables/README.md)
- Homepage: [https://cloudquery.io](https://cloudquery.io)
- Documentation: [https://docs.cloudquery.io](https://docs.cloudquery.io)
- Discord: [https://cloudquery.io/discord](https://cloudquery.io/discord)

## Install

```bash
cloudquery gen source gcp
```

## Libraries in Use

- https://cloud.google.com/go/docs/reference
- https://github.com/googleapis/google-api-go-client

### Authentication

The GCP plugin authenticates using your [Application Default Credentials](https://cloud.google.com/sdk/gcloud/reference/auth/application-default). Available options are all the same options described [here](https://cloud.google.com/docs/authentication/provide-credentials-adc) in detail:

Local Environment:

- `gcloud auth application-default login` (recommended when running locally)

Google Cloud cloud-based development environment:

- When you run on Cloud Shell or Cloud Code credentials are already available.

Google Cloud containerized environment:

- When running on GKE use [workload identity](https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity).

[Google Cloud services that support attaching a service account](https://cloud.google.com/docs/authentication/provide-credentials-adc#attached-sa):

- Services such as Compute Engine, App Engine and functions supporting attaching a user-managed service account which will CloudQuery will be able to utilize.

On-premises or another cloud provider

- The suggested way is to use [Workload identity federation](https://cloud.google.com/iam/docs/workload-identity-federation)
- If not available you can always use service account keys and export the location of the key via `GOOGLE_APPLICATION_CREDENTIALS`. **Highly not recommended as long lived keys are a security risk**

### Configuration

Following are the options available for the gcp plugin

```yaml
spec:
  # Optional. If not specified either using all projects accessible.
  # project_ids:
  # - "<CHANGE_THIS_TO_YOUR_PROJECT_ID>"
  # Optional. ServiceAccountKeyJSON passed as value instead of a file path
  # service_account_key_json: "<YOUR_JSON_SERVICE_ACCOUNT_KEY_DATA>"
  # Can also use env variable
  # service_account_key_json: ${MY_ENV_VAR}
```

By default cloudquery will fetch all configurations from **all** resources in **all** regions in the **default** project. You can customize this behaviour with the following arguments:

- `project_ids` - Specify multiple projects that you want to sync from.

### Query Examples:

#### Find all buckets without uniform bucket-level access

```sql
select project_id, name from gcp_storage_buckets where uniform_bucket_level_access->>'Enabled' = 'true';
```
