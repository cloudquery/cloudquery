# Architecture

This is an advanced section describing the inner workings and design of CloudQuery. \(It might be useful when developing new providers\).

CloudQuery has a pluggable architecture and uses the [gRPC](https://grpc.io/docs/languages/go/basics/) to communicate between plugins.

![cloudquery high-level architecture](/images/cloudquery-architecture.png)

## CloudQuery Core Responsibilities

- Main entry point and CLI for the user.
- Reading CloudQuery configuration.
- Downloading, verifying, and running providers.
- Running policy packs.

## CloudQuery Provider Responsibilities

- Intended to be run only by cloudquery-core.
- Communicates with cloudquery-core over gRPC to receive commands and actions.
- Initialization, authentication, and fetching data via third-party cloud/SaaS API.
