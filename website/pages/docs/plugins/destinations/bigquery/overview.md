# BigQuery Destination Plugin

import { Callout } from 'nextra-theme-docs'
import { getLatestVersion } from "../../../../../utils/versions";
import { Badge } from "../../../../../components/Badge";

<Badge text={"Latest: " + getLatestVersion("destination", "bigquery")}/>

The BigQuery plugin syncs data from any CloudQuery source plugin(s) to a BigQuery database running on Google Cloud Platform.

The plugin currently only supports a streaming mode through the legacy streaming API. This is suitable for small- to medium-sized datasets, and will stream the results directly to the BigQuery database. A batch mode of operation is being developed to support larger datasets, but this is not currently supported.

<Callout type="info">
Streaming is not available for the [Google Cloud free tier](https://cloud.google.com/bigquery/pricing#free-tier).
</Callout>

## Before you begin

1. Make sure that billing is enabled for your Cloud project. Learn how to [check if billing is enabled on a project](https://cloud.google.com/billing/docs/how-to/verify-billing-enabled).
2. Create a BigQuery dataset that will contain the tables synced by CloudQuery. CloudQuery will automatically create the tables as part of a migration run on the first `sync`.
3. Ensure that you have write access to the dataset. See [Required Permissions](https://cloud.google.com/bigquery/docs/streaming-data-into-bigquery) for details.

## Example config

The following config reads the values for `project_id` and `dataset_id` from environment variables:

```yaml
kind: destination
spec:
  name: bigquery
  path: cloudquery/bigquery
  version: "VERSION_DESTINATION_BIGQUERY"
  write_mode: "append"
  spec:
    project_id: ${PROJECT_ID}
    dataset_id: ${DATASET_ID}
```

Note that the BigQuery plugin only supports the `append` write mode.

## Authentication

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
- If not available you can always use service account keys and export the location of the key via `GOOGLE_APPLICATION_CREDENTIALS`. (**Not recommended as long-lived keys are a security risk**)

## BigQuery Spec

This is the top-level spec used by the BigQuery destination plugin.

- `project_id` (string) (required)

  The id of the project where the destination BigQuery database resides.


- `dataset_id` (string) (required)

  The name of the BigQuery dataset within the project, e.g. `my_dataset`. This dataset needs to be created before running a sync or migration.


- `dataset_location` (string) (optional)

  The data location of the BigQuery dataset. If set, will be used as the default location for job operations. Pro-tip: this can solve "dataset not found" issues for newly created datasets.


- `time_partitioning` (string) (options: `none`, `hour`, `day`) (default: `none`)

  The time partitioning to use when creating tables. The partition time column used will always be `_cq_sync_time` so that all rows for a sync run will be partitioned on the hour/day the sync started.


- `service_account_key_json` (string) (default: empty).

  GCP service account key content. This allows for using different service accounts for the GCP source and BigQuery destination. If using service account keys, it is best to use [environment or file variable substitution](/docs/advanced-topics/environment-variable-substitution).

- `batch_size` (int, optional. Default: 1000)

  Number of rows to insert in a single batch.

## Underlying library

We use the official [cloud.google.com/go/bigquery](https://pkg.go.dev/cloud.google.com/go/bigquery) package for database connection.
