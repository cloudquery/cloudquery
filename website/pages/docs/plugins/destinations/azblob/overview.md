# Azure Blob Storage Destination Plugin

import { getLatestVersion } from "../../../../../utils/versions";
import { Badge } from "../../../../../components/Badge";

<Badge text={"Latest: " + getLatestVersion("destination", "azblob")}/>

This destination plugin lets you sync data from a CloudQuery source to remote Azure Blob Storage storage in various formats such as CSV, JSON.

This is useful in various use-cases, especially in data lakes where you can query the data direct from Athena or load it to various data warehouses such as BigQuery, RedShift, Snowflake and others.

## Authentication

Authenitcation is similar to Azure CLI. See also [azure source plugin](../../sources/azure/overview#authentication) for more information.

## Example

This example configures a Azure blob storage destination, to create CSV files in `https://cqdestinationazblob.blob.core.windows.net/test/path/to/files`. Note that the S3 plugin only supports `append` write-mode.

The (top level) spec section is described in the [Destination Spec Reference](/docs/reference/destination-spec).

```yaml
kind: destination
spec:
  name: "azblob"
  path: "cloudquery/azblob"
  version: "VERSION_DESTINATION_AZBLOB"
  write_mode: "append" # this plugin only supports 'append' mode

  spec:
    storage_account: "cqdestinationazblob"
    container: "test"
    path: "path/to/files"
    format: "csv"
```

## S3 Spec

This is the (nested) spec used by the CSV destination Plugin.

- `storage_account` (string) (required)

  Storage account where to sync the files.

- `container` (string) (required)

  Storage container inside the storage account where to sync the files.

- `path` (string) (required)

  Path to where the files will be uploaded in the above bucket.

- `format` (string) (required)

  Format of the output file. `json` and `csv` are supported.
