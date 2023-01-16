# CloudQuery Vs Airbyte

Both CloudQuery and Airbyte are open source ELT frameworks with a few key differences.

**Key Differences:**

- **High Performance** - CloudQuery SDK and connectors are written in Go utilizing excelent support of go-routines which enables high-performance and low memory usage.
- **Type System** - CloudQuery SDK supports a rich type system for connectors which enables richer schemas and more accurate data.
- **Statless** - CloudQuery architecture is execution agnostic, and can be deployed in any environment such as Kubernetes, ECS, Google Cloud Run, Batch Jobs, etc. Airbyte needs a Airbyte database to be run and maintained.
- **Connectors** - Airbyte has more source connectors then CloudQuery at the moment. CloudQuery team is maintaing a quality set of official widely used connectors while mostly developing CloudQuery SDK to enable developers to write their own high performance connectors.
