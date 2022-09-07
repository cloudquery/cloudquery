# Configuration

CloudQuery, just like [terraform](https://www.terraform.io/language/providers), relies on "providers" to extract, transform and load cloud assets configuration from Cloud Providers, SaaS Providers and other APIs into PostgreSQL.

This section is similar to the terraform providers section due to the similar design of CloudQuery providers, but their purpose and implementation is completely different:

CloudQuery providers are read-only providers that extract, transform and load cloud assets configuration while Terraform providers interact with the cloud to provision assets.

## Configuration

Each provider is configured by a `provider "provider_name"` that can include general options, and a set of `resources` this provider will extract data from.

Each provider CloudQuery supports can be found on [hub.cloudquery.io](https://hub.cloudquery.io).

Each provider defines a set of relational tables that can also be found on the [hub](/plugins/aws).

Inside the cloudquery main block you need to add

```yaml
cloudquery:
    providers:
        - name: aws
          version: latest
```

Default configuration for a specific provider can always be generated via `cloudquery init [provider]`
