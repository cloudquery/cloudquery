# CloudQuery Vs Airbyte

Both CloudQuery and Airbyte are open source ELT frameworks with a few key differences.

**Key Differences:**

- **High Performance** - CloudQuery SDK and connectors are written in Go utilizing excellent support of go-routines which enables high-performance and low memory usage.
- **Type System** - CloudQuery SDK supports a rich type system for connectors which enables richer schemas and more accurate data.
- **Stateless** - CloudQuery architecture is execution agnostic, and can be deployed in any environment such as Kubernetes, ECS, Google Cloud Run, Batch Jobs, etc. Airbyte needs an Airbyte database, backend, workflows engine and other components that need to run and maintained.
- **Native as code approach** - CloudQuery was designed from the ground up to be configurable and run as code while in Airbyte it came later as a way to configure Airbyte UI.
- **Connectors** - Airbyte has more source connectors then CloudQuery at the moment. CloudQuery team is maintaining a quality set of official widely used connectors while mostly developing CloudQuery SDK to enable developers to write their own high performance connectors.
