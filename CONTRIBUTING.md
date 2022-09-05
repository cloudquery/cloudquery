
# Contributing to CloudQuery

:+1::tada: First off, thanks for taking the time to contribute! :tada::+1:

The following is a set of guidelines for contributing to this repository.


#### Table of Contents

[Code of Conduct](#code-of-conduct)

[I don't want to read this whole thing, I just have a question](#i-dont-want-to-read-this-whole-thing-i-just-have-a-question)

[What should I know before I get started?](#what-to-know-before-getting-started)
  * [Core](#cq-cli-directoryhttpsgithubcomcloudquerycloudquerytreemaincli)
  * [SDK](#cq-provider-sdk-repohttpsgithubcomcloudquerycq-provider-sdk)
  * [Plugins](#cq-plugins)

[How Can I Contribute?](#how-can-i-contribute)
  * [Reporting Bugs and Suggesting Enhancements](#reporting-bugs-and-suggesting-enhancements)
  * [Your First Code Contribution](#your-first-code-contribution)

## Code of Conduct

This project and everyone participating in it is governed by the [CloudQuery Code of Conduct](https://github.com/cloudquery/cloudquery/blob/main/CODE_OF_CONDUCT.md). By participating, you are expected to uphold this code. To report inappropriate behavior in violation of the code, please start by reaching out to us on our [Discord channel](https://cloudquery.io/discord).


## I don't want to read this whole thing, I just have a question

> **Note:** Please don't file an issue to ask a question. You'll get faster results by reaching out to the community on our [Discord channel](https://cloudquery.io/discord)


## What To Know Before Getting Started

### CloudQuery Architecture

CloudQuery has a pluggable architecture and is using the go-plugin to load, run, and communicate between providers via gRPC. To develop a new provider for CloudQuery, you don’t need to understand the inner workings of go-plugin as those are abstracted away via the [provider-sdk](#cq-provider-sdk-repo).

![architecture overview](https://www.cloudquery.io/images/cloudquery-architecture.png)

### Breakdown of Responsibilities and Repositories

#### CQ CLI [Directory](https://github.com/cloudquery/cloudquery/tree/main/cli):
- Main entry point and CLI for the user
- Reading CloudQuery configuration
- Downloading, verifying, and running plugins

#### CQ Provider SDK [Repository](https://github.com/cloudquery/cq-provider-sdk):
- Interacting with CQ CLI for initialization and configuration
- Helper functions for defining table schemas
- Methods for testing resources
- Framework for running and building a plugin locally

#### CQ Plugins
- [Officially-supported Plugins](https://github.com/cloudquery/cloudquery/tree/main/plugins) and [Community plugins](https://github.com/search?p=1&q=cq-provider-&type=Repositories)
- Previously known as Providers
- Executed by CQ CLI via gRPC 
- Interaction with remote data sources:
    - Initialization of clients
    - Authentication
    - Fetching of configuration information
- More information about developing your own plugin can be found [here](https://docs.cloudquery.io/docs/developers/developing-new-provider)


## How Can I Contribute?

### Reporting Bugs and Suggesting Enhancements

See [issue_reporting.md](issue_reporting.md) for more details on how to do this.

### Your First Code Contribution

Unsure where to begin contributing to CloudQuery? You can start by looking through these `beginner` and `help-wanted` issues:

* [Beginner issues][beginner] - issues which should only require a few lines of code, and a test or two
* [Help wanted issues][help-wanted] - issues which should be a bit more involved than `beginner` issues


If you don't see any issues that you think you can help with reach out to the community on Discord. We would be happy to work with you!

#### Local Development

CloudQuery has the ability to be run locally with a corresponding local Postgres database. To get it up and running follow the following instructions:
* [Development Environment](contributing/development_environment.md)
* [Connecting to a database](https://docs.cloudquery.io/docs/getting-started#spawn-or-connect-to-a-database)
* [Debugging a Provider](https://docs.cloudquery.io/docs/developers/debugging)
* [Developing a New Provider](https://docs.cloudquery.io/docs/developers/developing-new-provider)

#### Further guides

* [Creating a new plugin](contributing/creating_a_new_plugin.md)
* [Adding a new resource](contributing/adding_a_new_resource.md)

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