```yaml copy
kind: source
spec: # Common source spec section
  name: heroku
  path: /path/to/downloaded/plugin # Buy from here: https://cloudquery.io/integrations/heroku
  registry: local
  version: "PREMIUM"
  tables: ["*"]
  destinations: ["DESTINATION_NAME"]

  spec: # Heroku specific section
    token: <YOUR_TOKEN_HERE>
```
