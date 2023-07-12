

### Plugin Versioning

Changes to schema, configurations and required user permissions are all factors that go into the versioning of the AWS plugin. Any release that requires manual changes to a an existing deployment of the AWS plugin in order to retain the same functionality will be indicated by an increase to the major version. When support for additional resources is added it will result in a minor version bump. This is important to be aware of because if you are using `tables: ["*"]` to specify the set of tables to sync then in minor versions new resources that might require additional IAM permissions might result in errors being raised.

<Callout type="info">
It is recommended that you specify the exact tables you wish to sync rather than using wildcard characters as well as specifying `skip_dependent_tables: true` to limit risk while upgrading versions as new resources that can take significant time to sync can be added.
</Callout>


###  Breaking changes
The following examples are some of the most common examples of reasons for a major version change:

1. Changing a primary key for a table
2. Changing the name of a table
3. Changing the permissions required to sync a resource


All releases contain a release log that indicates all of the changes (and highlights the breaking changes), all changelogs are available [here](https://github.com/cloudquery/cloudquery/releases?q=plugins-source-aws-). If you are ever unsure about a change that is included feel free to reach out to the CloudQuery team on Discord to find out more.

### Preview features

Sometimes features or tables will be released and marked as `alpha`. This indicates that future minor versions might change, break or remove functionality. This enables the CloudQuery team to release functionality prior to it being fully stable so that the community can give feedback. Once a feature is released as Generally Available then all of the above rules for semantic versioning will apply.


#### Current Preview features

The following features are currently in `Preview`

- the `table_options` parameter in the [aws plugin spec](/docs/plugins/sources/aws/configuration#aws-spec)
- All tables that are prefixed with `aws_alpha_` including:
    - [aws_alpha_cloudwatch_metrics](tables/aws_alpha_cloudwatch_metrics)
        - [aws_alpha_cloudwatch_metric_statistics](tables/aws_alpha_cloudwatch_metric_statistics)
    - [aws_alpha_costexplorer_cost_custom](tables/aws_alpha_costexplorer_cost_custom)
