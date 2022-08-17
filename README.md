<p align="center">
<a href="https://cloudquery.io">
<img alt="cloudquery logo" width=75% src="https://github.com/cloudquery/cloudquery/raw/main/cli/docs/images/logo.png" />
</a>
</p>

The open-source cloud asset inventory powered by SQL.

CloudQuery extracts, transforms, and loads your cloud assets into [normalized](https://hub.cloudquery.io) PostgreSQL tables. CloudQuery enables you to assess, audit, and monitor the configurations of your cloud assets.

CloudQuery key use-cases and features:

- **Search**: Use standard SQL to find any asset based on any configuration or relation to other assets.
- **Visualize**: Connect CloudQuery standard PostgreSQL database to your favorite BI/Visualization tool such as Grafana, QuickSight, etc...
- **Policy-as-Code**: Codify your security & compliance rules with SQL as the query engine.

### Links

- Homepage: https://cloudquery.io
- Releases: https://github.com/cloudquery/cloudquery/releases
- Documentation: https://docs.cloudquery.io
- Hub (Provider and schema docs): https://hub.cloudquery.io/

### Supported providers (Actively expanding)

Check out https://hub.cloudquery.io.

If you want us to add a new provider or resource, please open an [Issue](https://github.com/cloudquery/cloudquery/issues).

See [docs](https://docs.cloudquery.io/docs/developers/developing-new-provider) for developing a new provider.

## Getting Started

Please check out our 'Getting Started' guides:

- [Getting Started with AWS](https://docs.cloudquery.io/docs/getting-started/getting-started-with-aws)
- [Getting Started with GCP](https://docs.cloudquery.io/docs/getting-started/getting-started-with-gcp)
- [Getting Started with Azure](https://docs.cloudquery.io/docs/getting-started/getting-started-with-azure)

For other providers, you can follow their specific guide on [Cloudquery Hub](https://hub.cloudquery.io/providers), as well as reference the [Getting Started with AWS](https://docs.cloudquery.io/docs/getting-started/getting-started-with-aws) for general installation and configuration tips.

## Compile and run CLI

```
make build-cli
./bin/cli/cloudquery # --help to see all options
```

## Deployment via Helm

Check out [cloudquery/helm-charts](https://github.com/cloudquery/helm-charts)

## License

By contributing to CloudQuery you agree that your contributions will be licensed as defined on the LICENSE file.

## Hiring

If you are into Go, Backend, Cloud, GCP, AWS - ping us at jobs [at] our domain

## Contribution

Feel free to open a pull request for small fixes and changes. For bigger changes and new providers, please open an issue first to prevent double work and have relevant discussions.
