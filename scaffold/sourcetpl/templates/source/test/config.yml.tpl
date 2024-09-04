kind: source
spec:
  name: '{{.Name}}'
  version: v1.0.0
  destinations: [test]
  path: ./cq-source-{{.Name}}
  registry: local
  tables: ["*"]
  spec:
---
kind: destination
spec:
  name: test
  path: cloudquery/test
  version: "v2.2.3" # latest version of test plugin