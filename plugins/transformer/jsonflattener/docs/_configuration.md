```yaml copy
kind: transformer
spec:
  name: "jsonflattener"
  path: "cloudquery/jsonflattener"
  version: VERSION_TRANSFORMER_JSONFLATTENER
  spec:
    tables: ["aws_ec2_instances"]
```