# Salesforce Source Plugin Configuration Reference

## Example

This example syncs from Salesforce to a specified CQ destination.

```yaml
kind: source
# Common source-plugin configuration
spec:
  name: slack
  path: cloudquery/salesforce
  version: "VERSION_SOURCE_SALESFORCE"
  tables: ["*"]
  destinations: ["destination"]
  # Salesforce specific configuration
  spec:
	  client_id: "${SF_CLIENT_ID}"
    client_secret: "${SF_CLIENT_SECRET}"
    username: "${SF_USERNAME}"
    password: "${SF_PASSWORD}"
    include_objects: ["Account", "Contact"]
    exclude_objects: ["Account"]
```

## Salesforce Spec

This is the (nested) spec used by the Slack source plugin.

- `client_id` (string, required):
   
  An API token to access Slack resources. This can be obtained by [creating a Slack app](/docs/plugins/sources/slack/overview#step-1).

- `debug` (boolean, optional. Default: false):

  Turn on to activate debug logging from the Slack SDK.
