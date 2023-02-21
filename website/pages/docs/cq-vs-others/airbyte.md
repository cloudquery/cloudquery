# CloudQuery Vs Airbyte

Both CloudQuery and Airbyte are open source ELT frameworks with a few key differences.

**Key Differences:**

- **High Performance** - CloudQuery SDK and connectors are optimized for performance, utilizing Go's support for goroutines to enable high performance and low memory usage.
- **Type System** - CloudQuery SDK supports a rich type system for connectors which enables richer schemas and more accurate data.
- **Stateless** - CloudQuery architecture is execution agnostic, and can be deployed in any environment such as Kubernetes, ECS, Google Cloud Run, Batch Jobs, etc. Airbyte needs an Airbyte database, backend, workflows engine and other components that need to be run and maintained.
- **Code-first approach** - CloudQuery was designed from the ground up to be configurable and run as code, while in Airbyte it was added later as a way to configure the Airbyte UI. With CloudQuery, it is easy to store and deploy your ELT configuration as code.
- **Connectors** - Airbyte has more source connectors than CloudQuery at the moment. The CloudQuery team adds and maintains official, widely-used connectors, while also developing the CloudQuery SDK to enable developers to write their own high-performance connectors.
