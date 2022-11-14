# Heroku Source Plugin Configuration Reference

## Example

This example connects a Heroku account to a Postgres destination. The (top level) source spec section is described in the [Source Spec Reference](https://www.cloudquery.io/docs/reference/source-spec).

```yml
kind: source
spec: # Common source spec section
  name: heroku
  path: cloudquery/heroku
  version: "v3.0.5" # latest version of heroku plugin
  tables: ["*"]
  destinations: ["postgresql"]

  spec: # Heroku specific section
    token: <YOUR_TOKEN_HERE>
```

## Heroku Spec

This is the (nested) spec used by the Heroku source plugin.

- `token` (string, **required**)

  Heroku API token. See the [Authentication section](../README.md#Authentication) on how to generate it.
