```yaml copy
kind: source
spec: # Common source spec section
  name: heroku
  path: cloudquery/heroku
  registry: cloudquery
  version: "VERSION_SOURCE_HEROKU"
  tables: ["*"]
  destinations: ["DESTINATION_NAME"]

  spec: # Heroku specific section
    token: <YOUR_TOKEN_HERE>
```
