# Best Practices

## Read-Only Access to Cloud Providers

Each provider has its own way of authenticating (which is described in the hub), but a rule of thumb for all of them is that in 100% (if possible, of course) of the cases, CloudQuery requires only a read-only key to fetch the information. CloudQuery **does not** make any changes to your infrastructure or SaaS applications.

## Cross-Account Access

If possible, you should use one read-only account per cloud provider that has access to all your accounts/projects. For example, in AWS, you can use the [AssumeRole](https://docs.aws.amazon.com/STS/latest/APIReference/API_AssumeRole.html) capability. In GCP, an account should be able to access all relevant projects (read-only). See the appropriate documentation on that in [CloudQuery Hub](https://hub.cloudquery.io).

If this is not possible, you can use multiple accounts by specifying multiple [provider blocks](./cli/fetch/configuration.md).

## Daily Infrastructure Snapshot

It is advised to run `cloudquery fetch` via cron on a daily basis (on lambda or any other secure place that has access to the required infrastructure). This, of course, varies highly depending on your needs, and can run even more frequently.
