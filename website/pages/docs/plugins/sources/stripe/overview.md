# Stripe Source Plugin

import { getLatestVersion } from "../../../../../utils/versions";
import { Badge } from "../../../../../components/Badge";

<Badge text={"Latest: " + getLatestVersion("source", `stripe`)}/>

The CloudQuery Stripe plugin pulls data from Stripe and loads it into any supported CloudQuery destination (e.g. PostgreSQL).

## Authentication

In order to fetch information from Stripe, `cloudquery` needs to be authenticated using a [secret API key](https://stripe.com/docs/keys) from your Stripe account. You can view and manage your API keys in [the Stripe Dashboard](https://stripe.com/login?redirect=/account/apikeys).

## Incremental Syncing

The Stripe plugin supports incremental syncing. This means that only new data will be fetched from Stripe and loaded into your destination for supported tables (support depending on API endpoint). This is done by keeping track of the last item fetched and only fetching data that has been created since then.
To enable this, `backend` option must be set in the spec (as shown below). This is documented in the [Managing Incremental Tables](/docs/advanced-topics/managing-incremental-tables) section.

### Example

This example syncs from Stripe to a Postgres destination. The (top level) source spec section is described in the [Source Spec Reference](/docs/reference/source-spec). Incremental syncing is enabled and will be saved to a `.cq/state/` directory by default.

```yaml
kind: source
# Common source-plugin configuration
spec:
  name: stripe
  path: cloudquery/stripe
  version: "VERSION_SOURCE_STRIPE"
  tables: ["*"]
  destinations: ["postgresql"]
  skip_tables:
    - stripe_issuing*  # Needs sign-up at https://stripe.com/issuing
    - stripe_treasury* # Needs sign-up at https://stripe.com/treasury
    - stripe_sigma_scheduled_query_runs # Live keys only
  backend: local
  # Stripe specific configuration
  spec:
    api_key: "<YOUR_SECRET_API_KEY_HERE>"
```

## Configuration Reference

This is the (nested) spec used by the Stripe source plugin:

- `api_key` (string, required):
  Your secret API key from the Stripe Dashboard.

- `rate_limit` (integer, optional. Default: varies):
  Used to override number of requests allowed per second. Defaults to 90 req/sec for production environment keys, otherwise 20 req/sec.

- `max_retries` (integer, optional. Default: 2):
  Number of retries if a request was rate limited at the API endpoint.

- `stripe_debug` (boolean, optional. Default: false):
  Enables verbose logging on the Stripe client.
