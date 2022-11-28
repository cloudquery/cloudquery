# Heroku Source Plugin Configuration Reference

## Example

This example connects a Heroku account to a Postgres destination. The (top level) source spec section is described in the [Source Spec Reference](/docs/reference/source-spec).

```yaml
kind: source
spec: # Common source spec section
  name: heroku
  path: cloudquery/heroku
  version: "VERSION_SOURCE_HEROKU"
  tables: ["*"]
  destinations: ["postgresql"]

  spec: # Heroku specific section
    token: <YOUR_TOKEN_HERE>
```

## Heroku Spec

This is the (nested) spec used by the Heroku source plugin.

- `token` (string, **required**)

  Heroku API token. See the [Authentication section](overview#authentication) on how to generate it.
