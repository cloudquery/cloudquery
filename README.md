<!-- markdownlint-disable MD033 -->
<h1 align="center"><img alt="cloudquery logo" width=75% src="https://github.com/cloudquery/cloudquery/raw/main/cli/docs/images/logo.png"/></h1>
<!-- markdownlint-enable MD033 -->

The open source high performance data integration platform designed for security and infrastructure teams.

CloudQuery extracts, transforms, and loads your cloud assets. CloudQuery enables you to assess, audit, and monitor the configurations of your cloud assets.

CloudQuery key use-cases and features:

- **Search**: Use standard SQL to find any asset based on any configuration or relation to other assets.
- **Visualize**: Connect CloudQuery standard PostgreSQL database to your favorite BI/Visualization tool such as Grafana, QuickSight, etc.
- **Policy-as-Code**: Codify your security & compliance rules with SQL as the query engine.

### Links

- Homepage: https://www.cloudquery.io
- Releases: https://github.com/cloudquery/cloudquery/releases
- Documentation: https://www.cloudquery.io/docs
- Plugins: https://www.cloudquery.io/plugins

### Supported plugins (Actively expanding)

Visit https://www.cloudquery.io/plugins.

If you want us to add a new plugin or resource, please open an [Issue](https://github.com/cloudquery/cloudquery/issues).

See [our guide for developing a new plugin](https://www.cloudquery.io/docs/developers/developing-new-provider).

## Getting Started

Please check out our 'Getting Started' guides:

- [Getting Started with AWS](https://www.cloudquery.io/docs/getting-started/getting-started-with-aws)
- [Getting Started with GCP](https://www.cloudquery.io/docs/getting-started/getting-started-with-gcp)
- [Getting Started with Azure](https://www.cloudquery.io/docs/getting-started/getting-started-with-azure)

For other plugins, you can visit our [plugins directory](https://www.cloudquery.io/plugins), as well as reference the [Getting Started with AWS](https://www.cloudquery.io/docs/getting-started/getting-started-with-aws) for general installation and configuration tips.

## Compile and run CLI

```bash
make build-cli
./bin/cloudquery # --help to see all options
```

## Deployment via Helm

Check out [cloudquery/helm-charts](https://github.com/cloudquery/helm-charts)

## License

By contributing to CloudQuery you agree that your contributions will be licensed as defined on the LICENSE file.

## Hiring

If you are into Go, Backend, Cloud, GCP, AWS - ping us at jobs [at] our domain

## Contribution

Feel free to open a pull request for small fixes and changes. For bigger changes and new providers, please open an issue first to prevent double work and have relevant discussions.
