<!-- markdownlint-disable MD033 -->
<h1 align="center"><img alt="cloudquery logo" width=75% src="https://github.com/cloudquery/cloudquery/raw/main/cli/docs/images/logo.png"/></h1>
<!-- markdownlint-enable MD033 -->

[![License: MPL 2.0](https://img.shields.io/badge/License-MPL_2.0-brightgreen.svg)](https://opensource.org/licenses/MPL-2.0)  [![Go Report Card](https://goreportcard.com/badge/github.com/cloudquery/cloudquery)](https://goreportcard.com/report/github.com/cloudquery/cloudquery)  [![CLI Workflow](https://github.com/cloudquery/cloudquery/actions/workflows/cli.yml/badge.svg)](https://github.com/cloudquery/cloudquery/actions/workflows/cli.yml)  

CloudQuery is a versatile [open-source](https://github.com/cloudquery/cloudquery) data movement tool built for developers that allows you to sync data from any [source](https://hub.cloudquery.io/plugins/source) to any [destination](https://hub.cloudquery.io/plugins/destination).

### **Installation**

Check out the [quickstart guide](https://www.cloudquery.io/docs/quickstart) for step-by-step instructions on completing your first sync with CloudQuery.

## **Why CloudQuery?**

- **High-performance data ingestion and processing:** Using Go's concurrency model and Apache Arrow, CloudQuery can quickly stream large amounts of data using [GRPC](https://docs.cloudquery.io/docs/developers/architecture).
- **Sync your data to any data destination:** You can move your data to any data source
- **Deploy anywhere:** CloudQuery can be run as a single-binary executable and deployed and run anywhere. This means you can run it in your [CI/CD pipelines](https://docs.cloudquery.io/docs/deployment/github-actions), inside your application, locally, or in the cloud.
- **Unlimited scale:** CloudQuery integrations are completely stateless and can be scaled horizontally on any platform, such as [VMs](https://docs.cloudquery.io/docs/deployment/google-cloud-vm), [Kubernetes](https://docs.cloudquery.io/docs/deployment/kubernetes), or batch jobs.
- **Security and compliance:** Reliable security measures protect sensitive data, and compliance features help meet industry standards.
- **Blazing fast**: CloudQuery is optimized for performance, utilizing the excellent Go concurrency model with lightweight goroutines and [streaming your data over GRPC](https://docs.cloudquery.io/docs/developers/architecture).
- **Open source framework**: Develop integrations in [Go](https://docs.cloudquery.io/docs/developers/creating-new-integration/go-source), [Python](https://docs.cloudquery.io/docs/developers/creating-new-integration/python-source), [Java](https://docs.cloudquery.io/docs/developers/creating-new-integration/java-source), or [JavaScript](https://docs.cloudquery.io/docs/developers/creating-new-integration/javascript-source) using the [open source CloudQuery SDK](https://github.com/cloudquery).

## **Use Cases**

- [**Cloud Security Posture Management (CSPM)**](https://www.cloudquery.io/blog/how-to-build-a-cspm-with-grafana-and-cloudquery): Use as a [CSPM](https://www.cloudquery.io/blog/how-to-build-a-cspm-with-grafana-and-cloudquery) solution to monitor and enforce security policies across your cloud infrastructure for [AWS](https://hub.cloudquery.io/plugins/source/cloudquery/aws/latest/docs), [GCP](https://hub.cloudquery.io/plugins/source/cloudquery/gcp/latest/docs), [Azure](https://hub.cloudquery.io/plugins/source/cloudquery/azure/latest/docs) and many more.
- [**Cloud Asset Inventory**](https://www.cloudquery.io/blog/what-is-a-cloud-asset-inventory): First-class support for [all major cloud infrastructure providers](https://hub.cloudquery.io/plugins/source?categories=cloud-infrastructure) such as [AWS](https://www.cloudquery.io/blog/building-cloud-asset-inventory-with-aws), [GCP](https://www.cloudquery.io/blog/building-cloud-asset-inventory-with-gcp), and [Azure](https://www.cloudquery.io/blog/how-to-build-a-cloud-asset-inventory-for-azure) allows you to [collect and unify your cloud configuration data](https://www.cloudquery.io/blog/how-to-build-a-multi-cloud-asset-inventory).
- **Cloud FinOps**: Collect and unify billing data from cloud providers to save money on your cloud expenses.
- **ELT Platform**: With hundreds of integration combinations and extensible architecture, CloudQuery can be used for reliable, efficient export from any API to any database or from one database to another.
- **Attack Surface Management**: [Solution](https://www.cloudquery.io/how-to-guides/attack-surface-management-with-graph) for continuous discovery, analysis, and monitoring of potential attack vectors that make up your organization's attack surface.
- **Eliminate data silos**: Eliminate data silos across your organization, unifying data between [security](https://hub.cloudquery.io/plugins/source?categories=security), [infrastructure](https://hub.cloudquery.io/plugins/source?categories=cloud-infrastructure), [marketing](https://hub.cloudquery.io/plugins/source?categories=marketing-analytics), and [finance](https://hub.cloudquery.io/plugins/source?categories=finance) teams.

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

Feel free to open a pull request for small fixes and changes. For bigger changes and new integrations, please open an issue first to prevent duplicated work and to have the relevant discussions first.

## Open source

The CloudQuery framework, SDK, CLI and some integrations are open source - please [file an issue](https://github.com/cloudquery/cloudquery/issues/new/choose) before opening a PR.
