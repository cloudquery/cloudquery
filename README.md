<!-- markdownlint-disable MD033 -->
<h1 align="center"><img alt="cloudquery logo" width=75% src="https://github.com/cloudquery/cloudquery/raw/main/cli/docs/images/logo.png"/></h1>
<!-- markdownlint-enable MD033 -->

[![License: MPL 2.0](https://img.shields.io/badge/License-MPL_2.0-brightgreen.svg)](https://opensource.org/licenses/MPL-2.0)  [![Go Report Card](https://goreportcard.com/badge/github.com/cloudquery/cloudquery)](https://goreportcard.com/report/github.com/cloudquery/cloudquery)  [![CLI Workflow](https://github.com/cloudquery/cloudquery/actions/workflows/cli.yml/badge.svg)](https://github.com/cloudquery/cloudquery/actions/workflows/cli.yml)  

CloudQuery is an [open-source](https://github.com/cloudquery/cloudquery),
high-performance data integration framework built for developers, with support for a wide range of plugins.

CloudQuery extracts, transforms, and loads configuration from cloud APIs, files or databases to
variety of supported destinations such as databases, data lakes, or streaming platforms
for further analysis.

### Installation

See the **[Quickstart guide](https://www.cloudquery.io/docs/quickstart)** for instructions how to start syncing data with CloudQuery.

## Why CloudQuery?

- **Blazing fast**: CloudQuery is optimized for performance, utilizing the excellent Go concurrency model with light-weight goroutines.
- **Deploy anywhere**: CloudQuery plugins are single-binary executables and can be deployed and run anywhere.
- **Open source framework**: Language-agnostic, extensible plugin architecture using [Apache Arrow](https://arrow.apache.org/): develop your own plugins in Go, Python, Java or JavaScript using the [CloudQuery SDK](https://docs.cloudquery.io/docs/developers/creating-new-plugin).
- **Pre-built queries**: CloudQuery maintains a number of out-of-the-box security and compliance policies for cloud infrastructure.
- **Unlimited scale**: CloudQuery plugins are stateless and can be scaled horizontally on any platform, such as EC2, Kubernetes, batch jobs or any other compute.

## Use Cases

- **Cloud Security Posture Management**: Use as a CSPM solution to monitor and enforce security policies across your cloud infrastructure for AWS, GCP, Azure and many more.
- **Cloud Asset Inventory**: First-class support for major cloud infrastructure providers such as AWS, GCP and Azure allow you to collect and unify configuration data.
- **Cloud FinOps**: Collect and unify billing data from cloud providers to drive financial accountability.
- **ELT Platform**: With hundreds of plugin combinations and extensible architecture, CloudQuery can be used for reliable, efficient export from any API to any database, or from one database to another.
- **Attack Surface Management**: [solution](https://www.cloudquery.io/how-to-guides/attack-surface-management-with-graph) for continuous discovery, analysis and monitoring of potential attack vectors that make up your organization's attack surface.
- **Eliminate data silos**: Eliminate data silos across your organization, unifying data between security, infrastructure, marketing and finance teams.

### Links

- Homepage: [https://www.cloudquery.io](https://www.cloudquery.io)
- Documentation: [https://www.cloudquery.io/docs](https://www.cloudquery.io/docs)
- Integrations: [https://hub.cloudquery.io](https://hub.cloudquery.io)
- Open Source Releases: [https://github.com/cloudquery/cloudquery/releases](https://github.com/cloudquery/cloudquery/releases)
- Plugin SDK: [https://github.com/cloudquery/plugin-sdk](https://github.com/cloudquery/plugin-sdk)

## License

By contributing to CloudQuery you agree that your contributions will be licensed as defined on the LICENSE file.

## Hiring

If you are into Go, Backend, Cloud, GCP, AWS - ping us at jobs [at] our domain

## Contribution

Feel free to open a pull request for small fixes and changes. For bigger changes and new plugins, please open an issue first to prevent duplicated work and to have the relevant discussions first.

## Open source

The CloudQuery framework, SDK, CLI and some plugins are open source - please [file an issue](https://github.com/cloudquery/cloudquery/issues/new/choose) before opening a PR.
