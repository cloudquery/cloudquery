# Stripe Source Plugin

import { getLatestVersion } from "../../../../../utils/versions";
import { Badge } from "../../../../../components/Badge";

<Badge text={"Latest: " + getLatestVersion("source", `stripe`)}/>

The CloudQuery Stripe plugin pulls data from Stripe and loads it into any supported CloudQuery destination (e.g. PostgreSQL).

## Authentication

In order to fetch information from Stripe, `cloudquery` needs to be authenticated. The Stripe API uses [API keys](https://stripe.com/docs/keys) to authenticate requests. You can view and manage your API keys in [the Stripe Dashboard](https://stripe.com/login?redirect=/account/apikeys).

# Configuration Reference

This is the (nested) spec used by the Stripe source plugin:

- `api_key` (string, required):
  Your API key from the Stripe Dashboard.

- `max_retries` (integer, optional. Default: 2):
  Number of retries if a request was rate limited.

## Example

This example syncs from Stripe to a Postgres destination. The (top level) source spec section is described in the [Source Spec Reference](https://www.cloudquery.io/docs/reference/source-spec).

```yaml
kind: source
# Common source-plugin configuration
spec:
  name: stripe
  path: cloudquery/stripe
  version: "VERSION_SOURCE_STRIPE"
  tables: ["*"]
  destinations: ["postgresql"]

  # Stripe specific configuration
  spec:
    api_key: "<YOUR_API_KEY_HERE>"
```
