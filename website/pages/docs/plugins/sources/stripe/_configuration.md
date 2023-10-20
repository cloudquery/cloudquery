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
  backend_options:
    table_name: "cq_state_stripe"
    connection: "@@plugins.DESTINATION_NAME.connection"
  # Stripe specific configuration
  spec:
    api_key: "<YOUR_SECRET_API_KEY_HERE>"
```
