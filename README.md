<!-- markdownlint-disable MD033 -->
<h1 align="center"><img alt="cloudquery logo" width=75% src="https://github.com/cloudquery/cloudquery/raw/main/cli/docs/images/logo.png"/></h1>
<!-- markdownlint-enable MD033 -->

[![License: MPL 2.0](https://img.shields.io/badge/License-MPL_2.0-brightgreen.svg)](https://opensource.org/licenses/MPL-2.0) [![Go Report Card](https://goreportcard.com/badge/github.com/cloudquery/cloudquery)](https://goreportcard.com/report/github.com/cloudquery/cloudquery) [![CLI Workflow](https://github.com/cloudquery/cloudquery/actions/workflows/cli.yml/badge.svg)](https://github.com/cloudquery/cloudquery/actions/workflows/cli.yml)

<p align="center"><strong>CloudOps, without the chaos</strong></p>

[CloudQuery](https://cloudquery.io) is a cloud asset inventory built for platform teams. Sync your cloud infrastructure metadata into your data warehouse, powering insights and automation. Unify data and provide context across AWS, Azure, GCP, and 70+ cloud and SaaS sources such as Wiz, Finout, and GitHub.

## Installation

```bash
brew install cloudquery/tap/cloudquery
```

Check out the [quickstart guide](https://cli-docs.cloudquery.io/docs/quickstart/) for install guides for Linux and Windows and for step-by-step instructions on completing your first sync with CloudQuery.

## **Why CloudQuery?**

- **Specialized plugin coverage** - Support for cloud infrastructure, security, and FinOps sources with normalization, rate limit handling, and more.
- **Normalized data** - 
- **Queryable with SQL** - No more writing scripts to hit poorly documented APIs: CloudQuery makes it easy to unify and make your cloud asset data accessible.
- **Integrate with anything** - Turn your cloud data into useful signals by connecting with BI tools, Slack alerts, Jira ticketing, and more.
- **Fast, powerful syncs** - Move large volumes of data with high performance and fine-grained control, powered by Apache Arrow.
- **Runs on your infrastructure** - Your cloud data never touches CloudQuery's servers. Full privacy, built for regulated, secure, and performance-critical environments.
- **Composable and flexible** - Use the languages, destinations, and orchestrators you want. CloudQuery is built to fit into your stack, not the other way around.
- **Built for developers** - Code-first, extensible plugins, multi-language, open plugin system, no lock-in. Write it, extend it, ship it.

## **Use Cases**

- [**Cloud Asset Inventory**](https://www.cloudquery.io/blog/what-is-a-cloud-asset-inventory): First-class support for [all major cloud infrastructure providers](https://hub.cloudquery.io/plugins/source?categories=cloud-infrastructure) such as [AWS](https://www.cloudquery.io/blog/building-cloud-asset-inventory-with-aws), [GCP](https://www.cloudquery.io/blog/building-cloud-asset-inventory-with-gcp), and [Azure](https://www.cloudquery.io/blog/how-to-build-a-cloud-asset-inventory-for-azure) allows you to [collect and unify your cloud configuration data](https://www.cloudquery.io/blog/how-to-build-a-multi-cloud-asset-inventory).
- [**Cloud Security Posture Management (CSPM)**](https://www.cloudquery.io/blog/how-to-build-a-cspm-with-grafana-and-cloudquery): Use as a [CSPM](https://www.cloudquery.io/blog/how-to-build-a-cspm-with-grafana-and-cloudquery) solution to monitor and enforce security policies across your cloud infrastructure for [AWS](https://hub.cloudquery.io/plugins/source/cloudquery/aws/latest/docs), [GCP](https://hub.cloudquery.io/plugins/source/cloudquery/gcp/latest/docs), [Azure](https://hub.cloudquery.io/plugins/source/cloudquery/azure/latest/docs) and many more.
- **Cloud FinOps**: Collect and unify billing data from cloud providers to save money on your cloud expenses.

### Links

- Homepage: [https://www.cloudquery.io](https://www.cloudquery.io)
- Documentation: [https://docs.cloudquery.io/](https://docs.cloudquery.io/docs)
- Integrations: [https://hub.cloudquery.io](https://hub.cloudquery.io)
- Plugin SDK: [https://github.com/cloudquery/plugin-sdk](https://github.com/cloudquery/plugin-sdk)

## License

By contributing to CloudQuery, you agree that your contributions will be licensed as defined in the LICENSE file.

## Contribution

Feel free to open a pull request for small fixes and changes. For bigger changes and new integrations, please open an issue first to prevent duplicated work and to have the relevant discussions first.

## Open source

The CloudQuery framework, SDK, CLI, and some integrations are open source - please [file an issue](https://github.com/cloudquery/cloudquery/issues/new/choose) before opening a PR.
Any code that was open source and moved to closed source can be found in the `git` history of this repo, or via the following links:

- [Closed source files under MPL 2.0 license (zip file)](https://mozilla-public-license-assets.cloudquery.io/MPL%202.0%20Assets/cloudquery-private-2025-10-15.zip)
- [Closed source files under MPL 2.0 license (CSV file)](https://mozilla-public-license-assets.cloudquery.io/MPL%202.0%20Assets/cloudquery-private-2025-10-15.csv)
