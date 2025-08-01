<!-- markdownlint-disable MD033 -->
<h1 align="center"><img alt="cloudquery logo" width=75% src="https://github.com/cloudquery/cloudquery/raw/main/cli/docs/images/logo.png"/></h1>
<!-- markdownlint-enable MD033 -->

[![License: MPL 2.0](https://img.shields.io/badge/License-MPL_2.0-brightgreen.svg)](https://opensource.org/licenses/MPL-2.0)  [![Go Report Card](https://goreportcard.com/badge/github.com/cloudquery/cloudquery)](https://goreportcard.com/report/github.com/cloudquery/cloudquery)  [![CLI Workflow](https://github.com/cloudquery/cloudquery/actions/workflows/cli.yml/badge.svg)](https://github.com/cloudquery/cloudquery/actions/workflows/cli.yml)  

[CloudQuery](https://cloudquery.io) is a high-performance data movement that runs entirely on your infrastructure. Extract from any source, from cloud infrastructure to SaaS, powering AI applications with CloudQuery’s flexible, composable data movement framework.

## Installation

```bash
brew install cloudquery/tap/cloudquery
```

Check out the [quickstart guide](https://cli-docs.cloudquery.io/docs/quickstart/) for install guides for Linux and Windows and for step-by-step instructions on completing your first sync with CloudQuery.

## **Why CloudQuery?**

- **Composable and flexible** - Use the languages, destinations, and orchestrators you want. CloudQuery is built to fit into your stack, not the other way around.
- **Runs on your infrastructure** - Your cloud data never touches CloudQuery's servers. Full privacy, built for regulated, secure, and performance-critical environments.
- **Built for developers** - Code-first, extensible plugins, multi-language, open plugin system, no lock-in. Write it, extend it, ship it. No black boxes, no unexplained failures.
- **Fast, powerful data movement** - Move large volumes of data with high performance and fine-grained control, powered by Apache Arrow. Perfect for feeding AI models, LLM pipelines, or large-scale data stores.
- **Specialized plugin coverage** - Support for complex, unique data sources such as cloud infrastructure, security, and FinOps data.

## **Use Cases**

- [**Cloud Security Posture Management (CSPM)**](https://www.cloudquery.io/blog/how-to-build-a-cspm-with-grafana-and-cloudquery): Use as a [CSPM](https://www.cloudquery.io/blog/how-to-build-a-cspm-with-grafana-and-cloudquery) solution to monitor and enforce security policies across your cloud infrastructure for [AWS](https://hub.cloudquery.io/plugins/source/cloudquery/aws/latest/docs), [GCP](https://hub.cloudquery.io/plugins/source/cloudquery/gcp/latest/docs), [Azure](https://hub.cloudquery.io/plugins/source/cloudquery/azure/latest/docs) and many more.
- [**Cloud Asset Inventory**](https://www.cloudquery.io/blog/what-is-a-cloud-asset-inventory): First-class support for [all major cloud infrastructure providers](https://hub.cloudquery.io/plugins/source?categories=cloud-infrastructure) such as [AWS](https://www.cloudquery.io/blog/building-cloud-asset-inventory-with-aws), [GCP](https://www.cloudquery.io/blog/building-cloud-asset-inventory-with-gcp), and [Azure](https://www.cloudquery.io/blog/how-to-build-a-cloud-asset-inventory-for-azure) allows you to [collect and unify your cloud configuration data](https://www.cloudquery.io/blog/how-to-build-a-multi-cloud-asset-inventory).
- **Cloud FinOps**: Collect and unify billing data from cloud providers to save money on your cloud expenses.
- **ELT Platform**: With hundreds of integration combinations and an extensible architecture, CloudQuery can be used for reliable, efficient export from any API to any database or from one database to another.
- **Attack Surface Management**: [Solution](https://www.cloudquery.io/how-to-guides/attack-surface-management-with-graph) for continuous discovery, analysis, and monitoring of potential attack vectors that make up your organization's attack surface.
- **Eliminate data silos**: Eliminate data silos across your organization, unifying data between [security](https://hub.cloudquery.io/plugins/source?categories=security), [infrastructure](https://hub.cloudquery.io/plugins/source?categories=cloud-infrastructure), [marketing](https://hub.cloudquery.io/plugins/source?categories=marketing-analytics), and [finance](https://hub.cloudquery.io/plugins/source?categories=finance) teams.

### Links

- Homepage: [https://www.cloudquery.io](https://www.cloudquery.io)
- Documentation: [https://platform-docs.cloudquery.io/](https://cli-docs.cloudquery.io/)
- Integrations: [https://hub.cloudquery.io](https://hub.cloudquery.io)
- Plugin SDK: [https://github.com/cloudquery/plugin-sdk](https://github.com/cloudquery/plugin-sdk)

## License

By contributing to CloudQuery, you agree that your contributions will be licensed as defined in the LICENSE file.

## Contribution

Feel free to open a pull request for small fixes and changes. For bigger changes and new integrations, please open an issue first to prevent duplicated work and to have the relevant discussions first.

## Open source

The CloudQuery framework, SDK, CLI, and some integrations are open source - please [file an issue](https://github.com/cloudquery/cloudquery/issues/new/choose) before opening a PR.
