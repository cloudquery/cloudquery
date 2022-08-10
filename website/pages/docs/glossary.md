# Glossary

## CLI

The CloudQuery CLI (command-line interface) is the core [open-source](https://github.com/cloudquery/cloudquery) project that expects to be run on your local machine as CLI or on a remote machine.

## SDK

The [CloudQuery SDK](https://github.com/cloudquery/cq-provider-sdk) is the open-source SDK library used by official and community providers to integrate with the CloudQuery ecosystem. The SDK makes it easy to write new providers and takes care of the TL in ETL (extract-transform-load).

## Fetch

Fetch is both the CLI command and the process when CloudQuery extracts all the configured resources in `cloudquery.yml`, transforms them, and loads them into the database.

## Policy

Policy compliance is a broad term and can refer to any kind of policy, from internal standards to regulatory requirements. A CloudQuery Policy is a codified form of this that is written with HCL as the logic layer and SQL as the query layer.

## Query

SQL query, usually targeting the CloudQuery database.

## Provider

CloudQuery Provider is a plugin responsible for extracting information/configuration from a specific cloud infrastructure provider SaaS application or literally anything else that is accessible via API (Rest, GRPC, GraphQL).

Currently, all providers are listed in <https://hub.cloudquery.io>.

Developing new provider

## HCL

[HashiCorp Configuration Language](https://github.com/hashicorp/hcl), which is used by CloudQuery to write configuration and policies.

## Resource

The fetch command is working on a list of resources defined in each provider (in `cloudquery.yml`).

For example, `ec2_instances` is a resource in `aws` provider.
