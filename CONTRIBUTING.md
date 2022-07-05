
# Contributing to CloudQuery

:+1::tada: First off, thanks for taking the time to contribute! :tada::+1:

The following is a set of guidelines for contributing to this repository.


#### Table of Contents

[Code of Conduct](#code-of-conduct)

[I don't want to read this whole thing, I just have a question!!!](#i-dont-want-to-read-this-whole-thing-i-just-have-a-question)

[What should I know before I get started?](#what-should-i-know-before-i-get-started)
  * [Core](#cq-core-repo)
  * [SDK](#cq-provider-sdk-repo)
  * [Provider](#cq-provider-repos)

[How Can I Contribute?](#how-can-i-contribute)
  * [Reporting Bugs](#reporting-bugs)
  * [Suggesting Enhancements](#suggesting-enhancements)
  * [Your First Code Contribution](#your-first-code-contribution)


## Code of Conduct

This project and everyone participating in it is governed by the [CloudQuery Code of Conduct](https://github.com/cloudquery/cloudquery/blob/main/CODE_OF_CONDUCT.md). By participating, you are expected to uphold this code. To report inappropriate behavior in violation of the code, please start by reaching out to us on our [Discord channel](https://cloudquery.io/discord).


## I don't want to read this whole thing I just have a question!!!

> **Note:** Please don't file an issue to ask a question. You'll get faster results by reaching out to the community on our [Discord channel](https://cloudquery.io/discord)


## What To Know Before Getting Started

### CloudQuery Architecture

CloudQuery has a pluggable architecture and is using the go-plugin to load, run, and communicate between providers via gRPC. To develop a new provider for CloudQuery, you don’t need to understand the inner workings of go-plugin as those are abstracted away via the [provider-sdk](#cq-provider-sdk-repo).

![architecture overview](https://docs.cloudquery.io/assets/images/cloudquery-architecture-02b1d1162883cd02510db2cb75c29d27.png)

### Breakdown of Responsibilities and Repositories

#### CQ Core [Repo](https://github.com/cloudquery/cloudquery):
- Main entry point and CLI for the user
- Reading CloudQuery configuration
- Downloading, verifying, and running providers
- Running policy packs
- Repository for CQ Core can be found 

#### CQ Provider SDK [Repo](https://github.com/cloudquery/cq-provider-sdk):
- Interacting with CQ Core for initialization and configuration
- Helper functions for defining table schemas
- Methods for testing the resource
- Framework for running and building a provider locally

#### CQ Provider [Repos](https://github.com/search?q=org%3Acloudquery+cq-provider-&type=repositories):
- Executed only by CQ-Core via gRPC 
- Interaction with remote data sources:
    - Initialization of clients
    - Authentication
    - Fetching of configuration information
- More information about developing your own provider can be found [here](https://docs.cloudquery.io/docs/developers/developing-new-provider)


## How Can I Contribute?

### Reporting Bugs

This section guides you through submitting a bug report for the AWS Provider for CloudQuery. Following these guidelines helps maintainers and the community understand your report :pencil:, reproduce the behavior :computer: :cloud:, and find related reports :mag_right:.

Before creating bug reports, please check [this list](#before-submitting-a-bug-report) as you might find out that you don't need to create one. When you are creating a bug report, please [include as many details as possible](#how-do-i-submit-a-good-bug-report). Fill out [the required template](.github/ISSUE_TEMPLATE/bug_report.md), the information it asks for helps us resolve issues faster.

> **Note:** If you find a **Closed** issue that seems like it is the same thing that you're experiencing, open a new issue and include a link to the original issue in the body of your new one

#### Before Submitting a Bug Report
* **Determine [which repository the problem should be reported in](#break-down-of-responsibilities-and-repositories)**
* **Perform a [cursory search](https://github.com/cloudquery/cloudquery/issues)** to see if the problem has already been reported. If it has **and the issue is still open**, add a comment to the existing issue instead of opening a new one

#### How Do I Submit a (Good) Bug Report?

Bugs are tracked as [GitHub issues](https://guides.github.com/features/issues/). After you've determined [which repository](#break-down-of-responsibilities-and-repositories) your bug is related to, create an issue on that repository and provide the following information by filling in [the template](.github/ISSUE_TEMPLATE/bug_report.md).

Explain the problem and include additional details to help maintainers reproduce the problem:

* **Use a clear and descriptive title** for the issue to identify the problem
* **Describe the Bug** in as many details as possible. For example, start by explaining how and where you are running CloudQuery (local machine, cloud service, docker, k8s, CI Pipeline, etc)
* **Provide specific examples to demonstrate the steps**. Include links to gists and or files, or copy/pasteable snippets to help give context to the issue. If you're providing snippets in the issue, use [Markdown code blocks](https://help.github.com/articles/markdown-basics/#multiple-lines)
* **Explain which behavior you expected to see instead and why.**
* **Include (sanitized) log output** execute CloudQuery with the `--enable-console-log` and `-v` flags to get all of the debug information

Provide more context by answering these questions:

Include details about your configuration and environment:

* **Which version of CloudQuery are you using?** You can get the exact version by running `cloudquery version`
* **What's in your config.hcl**? Include as much of the `config.hcl` as possible. This will allow the community to work to reproduce the issue and identify workarounds and/or create fixes to your issues


### Suggesting Enhancements

This section guides you through submitting an enhancement suggestion for CloudQuery, including completely new features, minor improvements to existing functionality and new providers. Following these guidelines helps maintainers and the community understand your suggestion :pencil: and find related suggestions :mag_right:.

Before creating enhancement suggestions, please check [this list](#before-submitting-an-enhancement-suggestion) as you might find out that you don't need to create one. When you are creating an enhancement suggestion, please [include as many details as possible](#how-do-i-submit-a-good-enhancement-suggestion). Fill in [the template](.github/ISSUE_TEMPLATE/feature_request.md), including the steps that you imagine you would take if the feature you're requesting existed.

#### Before Submitting an Enhancement Suggestion

* **Determine [which repository the enhancement should be suggested in](#break-down-of-responsibilities-and-repositories)**
* **Perform a [cursory search](https://github.com/search?q=is%3Aopen+label%3Aenhancement+org%3Acloudquery)** to see if the enhancement has already been suggested. If it has, add a comment to the existing issue instead of opening a new one

#### How Do I Submit a (Good) Enhancement Suggestion?

Enhancement suggestions are tracked as [GitHub issues](https://guides.github.com/features/issues/). After you've determined [which repository](#break-down-of-responsibilities-and-repositories) your enhancement suggestion is related to, create an issue on that repository and provide the following information:

* **Use a clear and descriptive title** for the issue to identify the suggestion
* **Describe the problem** In detail please try and convey the workflow or functionality you are trying to implement. This will help the community design and implement tooling that is both intuitive to use across many different domains as well as applicable to you and your specific challenge
* **Describe the use case** for this feature in as much detail as possible. Be sure to include any relevant information inluding links or other implementations

### Your First Code Contribution

Unsure where to begin contributing to CloudQuery? You can start by looking through these `beginner` and `help-wanted` issues:

* [Beginner issues][beginner] - issues which should only require a few lines of code, and a test or two
* [Help wanted issues][help-wanted] - issues which should be a bit more involved than `beginner` issues


If you don't see any issues that you think you can help with reach out to the community on Discord and we would be happy to work with you!

#### Local Development

CloudQuery has the ability to be run locally with a corresponding local postgres database. To get it up and running follow the following instructions:
* [Building and Running the Provider Locally](../docs/index.md)
* [Connecting to a database](https://docs.cloudquery.io/docs/getting-started#spawn-or-connect-to-a-database)
* [Debugging a Provider](https://docs.cloudquery.io/docs/developers/debugging)
* [Developing a New Provider](https://docs.cloudquery.io/docs/developers/developing-new-provider)

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