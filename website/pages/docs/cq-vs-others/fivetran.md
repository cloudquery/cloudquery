# CloudQuery Vs Fivetran

Fivetran is a closed source managed ELT SaaS solution with a few key differences.

**Key Differences:**

- **Open Source and extendable**: CloudQuery is open-source and has a [pluggable architecture](/docs/developers/architecture) which means you can contribute missing resources to existing plugins or you can easily create your own plugins (Using CloudQuery SDK) to grab data from in-house data stores or other SaaS applications.
- **High Performance** - CloudQuery SDK and connectors are written in Go utilizing excellent support of go-routines which enables high-performance and low memory usage.
- **Type System** - CloudQuery SDK supports a rich type system for connectors which enables richer schemas and more accurate data.
- **Self-hosted** - You can run CloudQuery on your own infrastructure, ensuring that your data doesn't leave your infrastructure. Fivetran only offers a managed solution at the moment.
- **Connectors** - Fivetran has more source connectors than CloudQuery at the moment. The CloudQuery team adds and maintains official, widely-used connectors, while also developing the CloudQuery SDK to enable developers to write their own high-performance connectors.