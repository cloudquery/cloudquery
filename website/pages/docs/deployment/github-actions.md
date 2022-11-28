---
title: Running CloudQuery with GitHub Actions
tag: tutorial
date: 2022/10/25
---

# Deploy-less Orchestration with GitHub Actions

In this tutorial, we will show you how to load AWS resources into a Postgres database by running CloudQuery as a [GitHub Action](https://github.com/features/actions), using the AWS source- and Postgresql destination plugins.

## Prerequisites

Since we'll be running CloudQuery in the context of a GitHub Action runner, we'll need to add AWS authentication.

To set up authentication with AWS from GitHub Actions you can follow the [Configuring OpenID Connect in Amazon Web Services blog](https://docs.github.com/en/actions/deployment/security-hardening-your-deployments/configuring-openid-connect-in-amazon-web-services) from GitHub.

> The role that you create will be used by CloudQuery to download your cloud configuration, so you should also grant it permissions to read from your cloud (e.g. ReadOnlyAccess permission policy)

## Creating the CloudQuery configuration file

Under the root of your repository, create a new `cloudquery.yml` file with the following content:

```yaml copy
kind: source
spec:
  name: 'aws'
  path: cloudquery/aws
  version: "VERSION_SOURCE_AWS"
  destinations: ['postgresql']
---
kind: destination
spec:
  name: 'postgresql'
  path: cloudquery/postgresql
  version: "VERSION_DESTINATION_POSTGRESQL"
  spec:
    connection_string: ${CQ_DSN} # The CQ_DSN environment variable will be set by GitHub Action workflow
```

> For more configuration options, [visit our docs](/docs/reference/source-spec)

## Creating the GitHub Action workflow

First we'll need [to create a GitHub secret](https://docs.github.com/en/actions/security-guides/encrypted-secrets#creating-encrypted-secrets-for-a-repository) with the name `CQ_DSN` and the value of the connection string to your PostgreSQL database.

Create a workflow file under `.github/workflows/cloudquery.yml` with the following content, and fill in `<role-arn>` and `<region>` according to the role you created in the prerequisites.

```yaml copy
name: CloudQuery
on:
  schedule:
    - cron: '0 3 * * *' # Run daily at 03:00 (3am)

jobs:
  cloudquery:
    permissions:
      id-token: write
      contents: read
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3 # Checkout the code so we have access to the config file
      - name: Configure AWS credentials # Setup AWS credentials (example)
        uses: aws-actions/configure-aws-credentials@v1
        with:
          role-to-assume: <role-arn> # based on the role you created in the prerequisites
          aws-region: <region> # based on the region you created the role in
      - uses: cloudquery/setup-cloudquery@v3
        name: Setup CloudQuery
        with:
          version: "VERSION_CLI"
      - name: Sync with CloudQuery
        run: cloudquery sync cloudquery.yml --log-console
        env:
          CQ_DSN: ${{ secrets.CQ_DSN }} # Connection string to a PostgreSQL database
```

Once committed to the default branch of the repository, the above workflow will run daily at 3 a.m. and will sync the AWS source plugin with the PostgreSQL destination plugin.

## Running CloudQuery in parallel to speed up sync time

By default, CloudQuery extracts all supported resources, which can take a bit of time, depending on the number of resources you have in your AWS account.

With the [GitHub Actions matrix configuration](https://docs.github.com/en/actions/using-jobs/using-a-matrix-for-your-jobs), you can split the sync process into multiple jobs and run them in parallel.

> In the example below, we'll split the sync process using regions, but you can split it by any other dimension, such as tables, accounts or any combination that makes sense for your use case.

First, we'll need to create a new `cloudquery-regions.yml` configuration file under the root of the repository:

```yaml copy
kind: source
spec:
  name: 'aws-REGION_PLACEHOLDER' # when splitting configurations, we need to keep the names unique
  path: cloudquery/aws
  version: "VERSION_SOURCE_AWS"
  destinations: ['postgresql-REGION_PLACEHOLDER']
  spec:
    regions:
      - REGION_PLACEHOLDER # This will be replaced by the matrix value
---
kind: destination
spec:
  name: 'postgresql-REGION_PLACEHOLDER' # when splitting configurations, we need to keep the names unique
  path: cloudquery/postgresql
  version: "VERSION_DESTINATION_POSTGRESQL"
  spec:
    connection_string: ${CQ_DSN} # The CQ_DSN environment variable will be set by GitHub Action workflow
```

To do so, create the following workflow file under `.github/workflows/cloudquery-parallel.yml`:

```yaml copy
name: CloudQuery Parallel
on:
  schedule:
    - cron: '0 3 * * *' # Run daily at 03:00 (3am)
jobs:
  cloudquery:
    permissions:
      id-token: write
      contents: read
    runs-on: ubuntu-latest
    strategy:
      matrix:
        region: [us-east-1, us-east-2, us-west-1, us-west-2, eu-west-1, eu-west-2] # List of regions to sync in parallel
    steps:
      - uses: actions/checkout@v3 # Checkout the code so we have access to the config file
      - name: Set region in config file
        uses: jacobtomlinson/gha-find-replace@657b0d1fe020da9943d1702b576f5d37d43b9c03
        with:
          include: cloudquery-regions.yml
          find: REGION_PLACEHOLDER
          replace: ${{ matrix.region }}
      - name: Configure AWS credentials # Setup AWS credentials (example)
        uses: aws-actions/configure-aws-credentials@v1
        with:
          role-to-assume: <role-arn> # based on the role you created in the prerequisites
          aws-region: <region> # based on the region you created the role in
      - uses: cloudquery/setup-cloudquery@v3
        name: Setup CloudQuery
        with:
          version: "VERSION_CLI"
      - name: Sync with CloudQuery
        run: cloudquery sync cloudquery-regions.yml --log-console
        env:
          CQ_DSN: ${{ secrets.CQ_DSN }} # Connection string to a PostgreSQL database
```

Once committed to the default branch of the repository, the above workflow will run daily at 3 a.m and will sync the AWS source plugin with the PostgreSQL destination plugin, in parallel, using the regions defined in the matrix.
