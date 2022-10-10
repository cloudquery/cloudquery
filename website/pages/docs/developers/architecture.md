# Architecture

This is an advanced section describing the inner workings and design of CloudQuery. \(It might be useful when developing new plugins\).

CloudQuery has a pluggable architecture and uses the [gRPC](https://grpc.io/docs/languages/go/basics/) to communicate between plugins.

![cloudquery high-level architecture](/images/cloudquery-architecture.png)

## CloudQuery CLI Responsibilities

- Main entry point and CLI for the user.
- Reading CloudQuery configuration.
- Downloading, verifying, and running sync from source to destination plugins

## CloudQuery Plugin Responsibilities

- Intended to be run only by CloudQuery CLI.
- Communicates with CloudQuery CLI over gRPC to receive commands and actions.
- Source Plugins: Initialization, authentication, and fetching data via third-party cloud/SaaS API.
- Destination Plugins: Authentication, Database migrations, Data Insertion.

## SDK

CloudQuery plugins utilize `plugin-sdk`, which abstracts most of the TL \(in ETL, extract-transform-load\). So, as a developer, you will only have to implement the \("E" in "ETL"\) initializing, authentication, and fetching of the data via the third-party APIs — the SDK will take care of transforming the data and loading it into the database. Also, your plugin will get support out-of-the-box for new features and things like other database support as cloudquery-core progresses.

## Resources

- [Creating a new plugin Tutorial](./creating-new-plugin).
- [cloudquery/plugin-sdk documentation](https://pkg.go.dev/github.com/cloudquery/plugin-sdk)
