# PagerDuty Source Plugin

import { getLatestVersion } from "../../../../../utils/versions";
import { Badge } from "../../../../../components/Badge";

<Badge text={"Latest: " + getLatestVersion("source", "pagerduty")}/>

The CloudQuery PagerDuty plugin extracts PagerDuty resources. It is based on [The PagerDuty Go SDK](https://github.com/PagerDuty/go-pagerduty) and the [PagerDuty REST API](https://developer.pagerduty.com/docs/ZG9jOjExMDI5NTUw-rest-api-v2-overview).

## Authentication

In order to authenticate with your PagerDuty account, you will need a [PagerDuty authorization token](https://support.pagerduty.com/docs/api-access-keys#section-generating-a-general-access-rest-api-key).
CloudQuery supports two methods of reading the authorization token:
- From a `~/.pd.yml` file, such as:
  ```yaml
  authtoken: <YOUR_AUTH_TOKEN>
  ```
- From an environment variable `PAGERDUTY_AUTH_TOKEN`.

## Configuration

In order to get started with the PagerDuty plugin, you need to create a YAML file in your CloudQuery configuration directory (e.g. named `pagerduty.yml`).

The following example sets up the PagerDuty plugin, and connects it to a postgresql destination:

```yaml
kind: source
spec:
  name: "pagerduty"
  path: cloudquery/pagerduty
  version: "VERSION_SOURCE_PAGERDUTY" 
  destinations: ["postgresql"]
  tables: ["*"]

  spec:
```