# Hacker News Source Plugin

import { getLatestVersion } from "../../../../../utils/versions";
import { Badge } from "../../../../../components/Badge";

<Badge text={"Latest: " + getLatestVersion("source", `hackernews`)}/>

The Hacker News Source plugin for CloudQuery extracts configuration from the [Hacker News API](https://github.com/HackerNews/API) and loads it into any supported CloudQuery destination (e.g. PostgreSQL).

It can be used for real applications, but is mainly intended to serve as an example of an incremental CloudQuery Source plugin. 

## Configuration

The following configuration syncs from Hacker News to a Postgres destination. The (top level) source spec section is described in the [Source Spec Reference](https://www.cloudquery.io/docs/reference/source-spec). The `postgresql` destination is not shown here and needs to be separately defined.

```yaml
kind: source
spec:
  name: "hackernews"
  path: "cloudquery/hackernews"
  version: "VERSION_SOURCE_HACKERNEWS"
  tables: ["*"]
  destinations: 
    - "postgresql"
  spec:
    item_concurrency: 100
```

- `item_concurrency` (int, optional):
    The number of items to fetch concurrently. Defaults to 100.

- `start_time` (string, optional):
    A date-time string in RFC3339 format. For example, `"2023-01-01T00:00:00Z"` will sync all items created on or after January 1, 2023. If not specified, the plugin will fetch all items. Note that because this is an incremental table, a previous cursor position will take precedence over this setting, unless the given start time is after the last cursor position.