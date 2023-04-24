# CloudQuery Policies
CloudQuery SQL Policies for AWS

## Policies and Compliance Frameworks Available

- [AWS CIS V1.2.0](./cis_v1.2.0/policy.sql)
- [AWS PCI DSS v3.2.1](./pci_dss_v3.2.1/policy.sql)
- [AWS Foundational Security Best Practices](./foundational_security/policy.sql)
- [AWS Public Egress](./public_egress/policy.sql)
- [AWS Publicly Available](./publicly_available/policy.sql)
- [AWS Unused Resources](./unused_resources/policy.sql)

## Installing

Clone this repository locally:

```bash
git clone https://github.com/cloudquery/cloudquery.git cloudquery
```

Check out the tag matching the AWS Source plugin version you are using:

```bash
cd cloudquery
git checkout plugins-source-aws-v16.3.0  # Example. Change to match the AWS version in your CloudQuery source config
```

Change directory into `plugins/source/aws/policies`:

```bash
cd plugins/source/aws/policies
```

## Running

You can execute policies with `psql`. For example:

```bash
# Set DSN to your PostgreSQL populated by CloudQuery
export DSN=postgres://postgres:pass@localhost:5432/postgres
# Execute CIS V1.2.0 Policy
psql ${DSN} -f  ./cis_v1.2.0/policy.sql
```

This will create all the results in `aws_policy_results` table which you can query directly, connect to any BI system (Grafana, Preset, AWS QuickSight, PowerBI, …).

You can also output it into CSV or HTML with the following built-in `psql` commands:

```bash
# Set DSN to your PostgreSQL populated by CloudQuery
export DSN=postgres://postgres:pass@localhost:5432/postgres
# default tabular output
psql ${DSN} -c "select * from aws_policy_results"
# CSV output
psql ${DSN} -c "select * from aws_policy_results" --csv
# HTML output
psql ${DSN} -c "select * from aws_policy_results" --html
```

## Dashboards

Currently we have a pre-built compliance dashboard on top of the `aws_policy_results` table which is available [here](../dashboards/grafana/compliance.json)

<img alt="Azure Asset Inventory Grafana Dashboard" src="../dashboards/grafana/compliance.png" width=50% height=50%>

See [installation instructions](../dashboards/README.md)
