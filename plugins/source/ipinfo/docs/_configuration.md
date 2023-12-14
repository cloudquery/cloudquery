```yaml copy
kind: source
spec:
  # Source spec section
  name: ipinfo
  path: cloudquery/ipinfo
  registry: cloudquery
#   registry: "grpc"
  version: "VERSION_SOURCE_IPINFO"
  tables: ["*"]
  destinations: ["DESTINATION_NAME"]
  spec:
    # Firestore specific configuration goes here
    ip: ${IP}
    token: ${TOKEN}
```