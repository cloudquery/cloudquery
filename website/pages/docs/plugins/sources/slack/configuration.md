# Slack Source Plugin Configuration Reference

## Example

This example syncs from Slack to a Postgres destination, using bot `token` authentication. The (top level) source spec section is described in the [Source Spec Reference](/docs/reference/source-spec).

```yaml copy
kind: source
# Common source-plugin configuration
spec:
  name: slack
  path: cloudquery/slack
  version: "VERSION_SOURCE_SLACK"
  tables: ["*"]
  destinations: ["postgresql"]

  # Slack specific configuration
  spec:
    token: "<YOUR_BOT_TOKEN_HERE>"
```

## Slack Spec

This is the (nested) spec used by the Slack source plugin.

- `token` (string, required):
   
  An API token to access Slack resources. This can be obtained by [creating a Slack app](/docs/plugins/sources/slack/overview#step-1).

- `debug` (boolean, optional. Default: false):

  Turn on to activate debug logging from the Slack SDK.
