# Mixpanel Source Plugin

import { getLatestVersion } from "../../../../../utils/versions";
import { Badge } from "../../../../../components/Badge";

<Badge text={"Latest: " + getLatestVersion("source", `mixpanel`)}/>

The CloudQuery Mixpanel plugin pulls data from Mixpanel and loads it into any supported CloudQuery destination (e.g. PostgreSQL).

## Authentication

In order to fetch information from Mixpanel, `cloudquery` needs to be authenticated using a [Service Account](https://developer.mixpanel.com/reference/service-accounts) from your Mixpanel account. You can view and manage your Service Accounts in Mixpanel `Organization Settings > Service Accounts`.

### Example

This example syncs from Mixpanel to a Postgres destination. The (top level) source spec section is described in the [Source Spec Reference](/docs/reference/source-spec).

```yaml
kind: source
# Common source-plugin configuration
spec:
  name: mixpanel
  path: cloudquery/mixpanel
  version: "VERSION_SOURCE_MIXPANEL"
  tables: ["*"]
  destinations: ["postgresql"]
  # Mixpanel specific configuration
  spec:
    username: "<YOUR_SERVICE_ACCOUNT_USER_HERE>"
    secret: "<YOUR_SERVICE_ACCOUNT_SECRET_HERE>"
    project_id: 12345 # Your project ID 
```

## Configuration Reference

This is the (nested) spec used by the Mixpanel source plugin:

- `username` (string, required):
  Your Service Account username from Mixpanel.

- `secret` (string, required):
  Service Account secret.

- `project_id` (integer, required):
  ID of the Mixpanel project to sync data from.

- `workspace_id` (integer, optional):
  ID of the Mixpanel workspace to sync data from.

- `start_date` (date in YYYY-MM-DD format, optional):
  Start date to sync data from. Defaults to 30 days ago.

- `end_date` (date in YYYY-MM-DD format, optional):
  End date to sync data until. Defaults to yesterday.

- `region` (string, optional):
  Your Mixpanel region. Defaults to "US". Possible values are "US" and "EU".

- `timeout_secs` (integer in seconds, optional. Default: 30):
  Timeout for requests against the Mixpanel API.

- `max_retries` (integer, optional. Default: 5):
  Number of retries if a request was rate limited at the API endpoint.
