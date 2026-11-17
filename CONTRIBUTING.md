
# Contributing to CloudQuery

Thank you for your interest in CloudQuery. This document tells you how to contribute to this repository.

## Code of Conduct

This project and everyone participating in it is governed by the [CloudQuery Code of Conduct](./CODE_OF_CONDUCT.md). By participating, you are expected to uphold this code. To report inappropriate behavior in violation of the code, please start by reaching out to us on our [Community](https://community.cloudquery.io).

## I don't want to read this whole thing, I just have a question

> **Note:** Please don't file an issue to ask a question. You'll get faster results by reaching out to the [community](https://community.cloudquery.io)

## We do not accept pull requests from outside the CloudQuery team

Contribute with an issue that has a reproduction. Do not open a pull request.

We changed this policy because:

- External pull requests are a supply chain risk.
- When an issue has a good reproduction, our team can make a fix that is tested and released faster than a review of an external pull request.
- The most valuable contribution is a clear report and a reproduction, not the code change.

Pull requests from outside the CloudQuery team are closed without review.

## How Can I Contribute?

### Report a bug with a reproduction

Use the [bug report template](https://github.com/cloudquery/cloudquery/issues/new?template=bug_report.yml). A good reproduction has:

- The CloudQuery version (`cloudquery --version`) and the plugin versions.
- A minimal configuration, with sensitive information removed.
- The exact commands that you ran, and the logs.
- The current behavior and the expected behavior.
- If possible, a failing test or a small script that shows the problem.

We add the `needs-reproduction` label to issues that we cannot reproduce. When you add the missing information, we triage the issue again.

### Request a feature

Use the [feature request template](https://github.com/cloudquery/cloudquery/issues/new?template=feature_request.yml). Tell us the problem that you want to solve, not only the solution.

## CloudQuery Architecture

CloudQuery has a pluggable architecture and is using [gRPC](https://grpc.io/) to communicate between source plugins, CLI and destination plugins. To develop a new plugin for CloudQuery, you don’t need to understand the inner workings of gRPC as those are abstracted away via the [plugin-sdk](#cloudquery-plugin-sdk-repository).

### Breakdown of Responsibilities and Repositories

### CloudQuery CLI [Directory](./cli)

* Main entry point and CLI for the user
* Reading CloudQuery configuration
* Downloading, verifying, and running plugins

### CloudQuery Plugin SDK [Repository](https://github.com/cloudquery/plugin-sdk)

* Interacting with CloudQuery CLI for initialization and configuration
* Helper functions for defining table schemas
* Methods for testing the resource
* Framework for running and building a plugin locally

### CloudQuery Plugins

* [Officially-supported Plugins](./plugins) and [Community plugins](https://github.com/search?p=1&q=cq-plugin-&type=Repositories)
* Previously known as Providers
* Executed by CloudQuery CLI via gRPC
* Interaction with remote data sources:
  * Initialization of clients
  * Authentication
  * Fetching of configuration information
* More information about developing your own plugin can be found [here](https://cloudquery.io/docs/developers/developing-new-provider)

## Develop Locally

Use a local environment to make a reproduction or a failing test.

CloudQuery has the ability to be run locally with a corresponding local Postgres database. To get it up and running follow the following instructions:

* [Development Environment](./contributing/development_environment.md)
* [Connecting to a database](https://www.cloudquery.io/docs/getting-started#spawn-or-connect-to-a-database)
* [Running plugins locally](https://www.cloudquery.io/docs/developers/running-locally)
* [Developing a new plugin](https://www.cloudquery.io/docs/developers/developing-new-provider)

### Further guides

* [Creating a new plugin](./contributing/creating_a_new_integration.md)
* [Adding a new resource](./contributing/adding_a_new_resource.md)
* [Adding New Plugin To CQ Monorepo](./contributing/adding_a_new_integration_to_cq_monorepo.md)

