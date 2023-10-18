```yaml copy
kind: source
# Common source-plugin configuration
spec:
  name: stripe
  path: cloudquery/stripe
  version: "VERSION_SOURCE_STRIPE"
  tables: ["*"]
  destinations: ["DESTINATION_NAME"]
  skip_tables:
    - stripe_issuing*  # Needs sign-up at https://stripe.com/issuing
    - stripe_treasury* # Needs sign-up at https://stripe.com/treasury
    - stripe_sigma_scheduled_query_runs # Live keys only
  # Stripe specific configuration
  spec:
    backend_options:
      table_name: "test_state_table"
      connection: "@@plugins.DESTINATION_NAME.connection"
    api_key: "<YOUR_SECRET_API_KEY_HERE>"
```

The Stripe plugin supports incremental syncing. This means that only new data will be fetched from Stripe and loaded into your destination for supported tables (support depending on API endpoint). This is done by keeping track of the last item fetched and only fetching data that has been created since then.
To enable this, `backend` option must be set in the spec (as shown below). This is documented in the [Managing Incremental Tables](/docs/advanced-topics/managing-incremental-tables) section.
