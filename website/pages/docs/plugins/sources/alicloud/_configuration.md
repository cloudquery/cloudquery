```yaml copy
kind: source
spec:
  name: "alicloud"
  path: "cloudquery/alicloud"
  registry: "cloudquery"
  version: "VERSION_SOURCE_ALICLOUD"
  tables: ["*"]
  destinations: ["DESTINATION_NAME"]
  spec:
    accounts:
      - name: my_account
        regions:
        - cn-hangzhou
        - cn-beijing
        - eu-west-1
        - us-west-1
        # ...
        access_key: ${ALICLOUD_ACCESS_KEY}
        secret_key: ${ALICLOUD_SECRET_KEY}
    # Optional parameters
    # concurrency: 50000
    # bill_history_months: 12
```
