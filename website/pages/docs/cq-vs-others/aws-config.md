# CloudQuery vs AWS Config

AWS Config is the native AWS asset inventory provided by AWS.

**Key Differences:**

- **Resource Types**: CloudQuery [supports](https://github.com/cloudquery/cloudquery/tree/main/plugins/source/aws/docs/tables) more than 250 types of resources while AWS Config currently supports about 120 types of resources (tables). Being an open-source project you can easily add the missing resources without being blocked by a vendor.
- **Database Agnostic and Raw Access to data**: CloudQuery supports multiple databases such as PostgreSQL, BigQuery and others. This makes it play nicely with the whole SQL eco-system and gives you the ability to re-use other tools like Grafana/BI. AWS Config is using a proprietary subset of SQL and database and thus doesn't give you the ability to re-use other tools easily.
- **Cloud Agnostic**: CloudQuery gives you the ability to assess, audit and monitor [multi-cloud and SaaS infrastructure](https://hub.cloudquery.io).
- **Pricing**: CloudQuery is open-source and thus you will pay only for the hosting of your PostgreSQL (you can use RDS, or any other managed version) and the compute for running [CQ binary](../deployment/overview). AWS Config charges both per item recorded and rule evaluated.
- **Limits**: See limits section below. AWS Config has a lot of soft and hard limit (that cannot be increased). CloudQuery doesn't impose any of those limits and is bound mainly by the PostgreSQL instance and compute.
- **Policy Language**: CloudQuery uses [standard SQL](https://github.com/cloudquery/cloudquery/tree/main/plugins/source/aws/policies) which you can then customize and monitor with your existing tools, stack. AWS Config uses JSON/YAML and subset SQL as the query engine.

## Limits Comparison

This will only list hard limits as soft limits can be increased (though it is not documented what is their max limit). Some limits are quite serious for large accounts (especially around number of rules/policies).

| Limit                                                                                                 | AWS Config | CloudQuery |
| ----------------------------------------------------------------------------------------------------- | ---------- | ---------- |
| Max Accounts                                                                                          | 10000      | No Limit   |
| Maximum number of conformance packs per account (CQ Policies)                                         | 50         | No Limit   |
| Maximum number of AWS Config Rules per conformance pack (CQ Queries)                                  | 130        | No Limit   |
| Maximum number of AWS Config Rules per account across all conformance packs (Total Queries)           | 130        | No Limit   |
| Maximum number of conformance packs per organization (CQ Policies)                                    | 50         | No Limit   |
| Maximum AWS Config Rules per organization conformance pack (CQ Queries)                               | 130        | No Limit   |
| Maximum number of AWS Config Rules per account across all organization conformance packs (CQ Queries) | 150        | No Limit   |
| Maximum number of organization AWS Config rules per organization (CQ Queries)                         | 150        | No Limit   |
