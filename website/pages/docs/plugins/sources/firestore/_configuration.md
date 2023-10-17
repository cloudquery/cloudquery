```yaml copy
kind: source
spec:
  # Source spec section
  name: firestore
  path: cloudquery/firestore
  version: "VERSION_SOURCE_FIRESTORE"
  tables: ["*"]
  destinations: ["DESTINATION_NAME"]
  spec:
    # Firestore specific configuration goes here
```
