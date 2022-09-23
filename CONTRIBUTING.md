
# Contributing to CloudQuery

:+1::tada: First off, thanks for taking the time to contribute! :tada::+1:

The following is a set of guidelines for contributing to this repository.

## Code of Conduct

This project and everyone participating in it is governed by the [CloudQuery Code of Conduct](./CODE_OF_CONDUCT.md). By participating, you are expected to uphold this code. To report inappropriate behavior in violation of the code, please start by reaching out to us on our [Discord channel](https://cloudquery.io/discord).

## I don't want to read this whole thing, I just have a question

> **Note:** Please don't file an issue to ask a question. You'll get faster results by reaching out to the community on our [Discord channel](https://cloudquery.io/discord)

## What To Know Before Getting Started

### CloudQuery Architecture

CloudQuery has a pluggable architecture and is using [gRPC](https://grpc.io/) to communicate between source plugins, CLI and destination plugins. To develop a new plugin for CloudQuery, you don’t need to understand the inner workings of gRPC as those are abstracted away via the [plugin-sdk](#cloudquery-plugin-sdk-repository).

### Breakdown of Responsibilities and Repositories

#### CloudQuery CLI [Directory](./cli)

* Main entry point and CLI for the user
* Reading CloudQuery configuration
* Downloading, verifying, and running plugins

#### CloudQuery Plugin SDK [Repository](https://github.com/cloudquery/plugin-sdk)

* Interacting with CloudQuery CLI for initialization and configuration
* Helper functions for defining table schemas
* Methods for testing the resource
* Framework for running and building a plugin locally

#### CloudQuery Plugins

* [Officially-supported Plugins](./plugins) and [Community plugins](https://github.com/search?p=1&q=cq-plugin-&type=Repositories)
* Previously known as Providers
* Executed by CloudQuery CLI via gRPC
* Interaction with remote data sources:
  * Initialization of clients
  * Authentication
  * Fetching of configuration information
* More information about developing your own plugin can be found [here](https://cloudquery.io/docs/developers/developing-new-provider)

## How Can I Contribute?

### Reporting Bugs and Requesting Feature Requests

Follow our [bug reporting template](https://github.com/cloudquery/cloudquery/issues/new?assignees=&labels=bug&template=bug_report.md) or [feature request template](https://github.com/cloudquery/cloudquery/issues/new?assignees=&labels=enhancement&template=feature_request.md) to ensure you provide all the necessary information for us to either reproduce and fix the bug or implement the feature.

### Your First Code Contribution

Unsure where to begin contributing to CloudQuery? You can start by looking through these [`good first issue` issues](https://github.com/cloudquery/cloudquery/issues?q=is%3Aopen+is%3Aissue+label%3A%22good+first+issue%22).
If you don't see any issues that you think you can help with reach out to the community on Discord and we would be happy to work with you!

#### Local Development

CloudQuery has the ability to be run locally with a corresponding local Postgres database. To get it up and running follow the following instructions:

* [Development Environment](./contributing/development_environment.md)
* [Connecting to a database](https://docs.cloudquery.io/docs/getting-started#spawn-or-connect-to-a-database)
* [Debugging a plugin](https://docs.cloudquery.io/docs/developers/debugging)
* [Developing a new plugin](https://docs.cloudquery.io/docs/developers/developing-new-provider)

#### Further guides

* [Creating a new plugin](./contributing/creating_a_new_plugin.md)
* [Adding a new resource](./contributing/adding_a_new_resource.md)

#### Commit Messages

We make use of the [Conventional Commits specification](https://www.conventionalcommits.org/en/v1.0.0/) for pull request titles. This allows us to categorize contributions and automate versioning for releases. Pull request titles should start with one of the prefixes specified in the table below:

| Title      | Message | Action |
| ----------- | ----------- |----------- |
| `chore: <Message>`      |  `<String>`       | patch release|
| `fix: <Message>`      |  `<String>`      | patch release|
| `feat: <Message>`      |  `<String>`       | patch release|
| `refactor: <Message>`      |  `<String>`       | patch release|
| `test: <Message>`      |  `<String>`       | patch release|

Additional context can be provided in parentheses, e.g. `fix(docs): Fix typo`. Breaking changes should be suffixed with `!`, e.g. `feat!: Drop support for X`. This will always result in a minor release.
