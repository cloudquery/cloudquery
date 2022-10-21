# GCP + Kubernetes (GKE) + PostgreSQL

```yaml
kind: source
spec:
  name: gcp
  version: "v1.0.1"
  destinations: ["postgresql"]
---
kind: source
spec:
  name: k8s
  version: "v2.0.1"
  destinations: ["postgresql"]
---
kind: destination
spec:
  name: "postgresql"
  version: "v0.3.0"
  write_mode: "overwrite" # overwrite, append
  spec:
    connection_string: "postgresql://{CQ_PG_USER}:{CQ_PG_PASS}@localhost:5432/postgres?sslmode=disable"
```

