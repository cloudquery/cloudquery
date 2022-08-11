# Architecture

This is an advanced section describing the inner workings and design of CloudQuery. \(It might be useful when developing new providers\).

CloudQuery has a pluggable architecture and uses the [go-plugin](https://github.com/hashicorp/go-plugin) to load, run, and communicate between providers via gRPC. To develop a new provider for CloudQuery, you don’t need to understand the inner workings of go-plugin, as those are abstracted away by [cq-provider-sdk](https://github.com/cloudquery/cq-provider-sdk).

![cloudquery high-level architecture](/images/cloudquery-architecture.png)

Similarly to any application utilizing the [go-plugin](https://github.com/hashicorp/go-plugin) framework, CloudQuery is split into [CloudQuery Core](https://github.com/cloudquery/cloudquery) and [CloudQuery Providers](https://github.com/orgs/cloudquery/repositories?language=&q=cloudquery-provider&sort=&type=).

## CloudQuery Core Responsibilities

- Main entry point and CLI for the user.
- Reading CloudQuery configuration.
- Downloading, verifying, and running providers.
- Running policy packs.

## CloudQuery Provider Responsibilities

- Intended to be run only by cloudquery-core.
- Communicates with cloudquery-core over gRPC to receive commands and actions.
- Initialization, authentication, and fetching data via third-party cloud/SaaS API.
