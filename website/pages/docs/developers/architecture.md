---
title: CloudQuery Architecture
description: Learn about the architecture of CloudQuery.
---

# Architecture

This is an advanced section describing the inner workings and design of CloudQuery. \(It might be useful when developing new integrations\).

CloudQuery has a plugin-based architecture and uses the [gRPC](https://grpc.io/docs/languages/go/basics/) to communicate between integrations.

![high-level architecture of CloudQuery](/images/cloudquery-architecture.png)

## CloudQuery CLI Responsibilities

- Main entry point and CLI for the user.
- Reading CloudQuery configuration.
- Downloading, verifying, and running sync from source to destination integrations

## CloudQuery Integration Responsibilities

- Intended to be run only by CloudQuery CLI.
- Communicates with CloudQuery CLI over gRPC to receive commands and actions.
- Source integrations: Initialization, authentication, and fetching data via third-party cloud/SaaS API.
- Destination integrations: Authentication, Database migrations, Data Insertion.

## SDK

CloudQuery integrations utilize `plugin-sdk`, which abstracts most of the TL \(in ETL, extract-transform-load\). So, as a developer, you will only have to implement the \("E" in "ETL"\) initializing, authentication, and fetching of the data via the third-party APIs — the SDK will take care of transforming the data and loading it into the database. Also, your integration will get support out-of-the-box for new features and things like other database support as `cloudquery-core` progresses.

## Resources

- [Creating a new integration Tutorial](./creating-new-integration).
- [cloudquery/plugin-sdk documentation](https://pkg.go.dev/github.com/cloudquery/plugin-sdk)
