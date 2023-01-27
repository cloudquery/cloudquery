# LaunchDarkly Source Plugin

import { getLatestVersion } from "../../../../../utils/versions";
import { Badge } from "../../../../../components/Badge";

<Badge text={"Latest: " + getLatestVersion("source", `launchdarkly`)}/>

The CloudQuery LaunchDarkly plugin pulls data from LaunchDarkly and loads it into any supported CloudQuery destination (e.g. PostgreSQL).

## Authentication

In order to fetch information from LaunchDarkly, `cloudquery` needs to be authenticated using a [Service Account](https://developer.launchdarkly.com/reference/service-accounts) from your LaunchDarkly account. You can view and manage your access tokens in LaunchDarkly `Account settings > Authorization`. Refer to [REST API Authentication](https://apidocs.launchdarkly.com/#section/Overview/Authentication) for more information.

### Access Token API Version

Creating the access token you should select at least `20220603` as the API version. This is the minimum API version supported by the LaunchDarkly plugin. Refer to [API version changelog](https://apidocs.launchdarkly.com/#section/Overview/Versioning) for more information.

## Incremental Syncing

The LaunchDarkly plugin supports incremental syncing for audit log data. This means that only new audit log entries will be fetched from LaunchDarkly and loaded into your destination. This is done by keeping track of the entry fetched and only fetching entries that has been created since then.
To enable this, `backend` option must be set in the spec (as shown below). This is documented in the [Managing Incremental Tables](/docs/advanced-topics/managing-incremental-tables) section.

### Example

This example syncs from LaunchDarkly to a Postgres destination. The (top level) source spec section is described in the [Source Spec Reference](/docs/reference/source-spec).

```yaml
kind: source
# Common source-plugin configuration
spec:
  name: launchdarkly
  path: cloudquery/launchdarkly
  version: "VERSION_SOURCE_LAUNCHDARKLY"
  tables: ["*"]
  destinations: ["postgresql"]
  backend: local
  # LaunchDarkly specific configuration
  spec:
    access_token: "<YOUR_ACCESS_TOKEN_HERE>"
```

## Configuration Reference

This is the (nested) spec used by the LaunchDarkly source plugin:

- `access_token` (string, required):
  Your access token from LaunchDarkly.
- `timeout_secs` (integer in seconds, optional. Default: 30):
  Timeout for requests against the LaunchDarkly API.
