# CloudQuery Vs Fivetran

Fivetran is a closed source managed ELT SaaS solution with a few key differences.

**Key Differences:**

**Open Source and Extensible**: CloudQuery is open-source and has a [pluggable architecture](https://hub.cloudquery.io) which means you can contribute missing resources to existing plugins or you can create your own plugins in order to grab data from proprietary APIs or other SaaS applications.
- **High Performance** - CloudQuery SDK and connectors are written in Go utilizing excelent support of go-routines which enables high-performance and low memory usage.
- **Type System** - CloudQuery SDK supports a rich type system for connectors which enables richer schemas and more accurate data.
- **Connectors** - Fivetran has more source connectors then CloudQuery at the moment. CloudQuery team is maintaing official widely used connectors while mostly developing CloudQuery SDK to enable developer to write their own high performance connectors.