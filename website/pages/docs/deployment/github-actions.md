---
title: Running CloudQuery with GitHub Actions
description: Learn how to use GitHub Actions to run CloudQuery syncs.
tag: tutorial
date: 2022/10/25
---

# Deploy-less Orchestration with GitHub Actions

In this tutorial, we will show you how to load AWS resources into a Postgres database by running CloudQuery as a [GitHub Action](https://github.com/features/actions), using the AWS source- and Postgresql destination integrations.

## Prerequisites

### Generating a CloudQuery API key

Downloading integrations requires users to be authenticated, normally this means running `cloudquery login` but that is not doable in a CI environment like GitHub Actions. The recommended way to handle this is to use an API key. More information on generating an API Key can be found [here](/docs/deployment/generate-api-key).

### AWS Authentication

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
  registry: cloudquery
  version: "VERSION_SOURCE_AWS"
  tables: ['*']
  destinations: ['postgresql']
---
kind: destination
spec:
  name: 'postgresql'
  path: cloudquery/postgresql
  registry: cloudquery
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
          version: "vVERSION_CLI"
      - name: Sync with CloudQuery
        run: cloudquery sync cloudquery.yml --log-console
        env:
          CLOUDQUERY_API_KEY: ${{ secrets.CLOUDQUERY_API_KEY }} # See https://docs.cloudquery.io/docs/deployment/generate-api-key
          CQ_DSN: ${{ secrets.CQ_DSN }} # Connection string to a PostgreSQL database
```

Once committed to the default branch of the repository, the above workflow will run daily at 3 a.m. and will sync the AWS source integration with the PostgreSQL destination integration.

> **Warning**
> GitHub automatically disables workflows on public repositories [if no repository activity has occurred for 60 days](https://docs.github.com/en/actions/using-workflows/disabling-and-enabling-a-workflow). This may impact your sync if the repository does not receive regular commits.

## Running CloudQuery in parallel to speed up sync time

By default, CloudQuery extracts all supported resources, which can take a bit of time, depending on the number of resources you have in your AWS account.

With the [GitHub Actions matrix configuration](https://docs.github.com/en/actions/using-jobs/using-a-matrix-for-your-jobs), you can split the sync process into multiple jobs and run them in parallel.

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
        shard: [1/4, 2/4, 3/4, 4/4] # Split the sync into 4 parts
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
          version: "vVERSION_CLI"
      - name: Sync with CloudQuery
        run: cloudquery sync cloudquery.yml --log-console --shard ${{ matrix.shard }}
        env:
          CLOUDQUERY_API_KEY: ${{ secrets.CLOUDQUERY_API_KEY }} # See https://docs.cloudquery.io/docs/deployment/generate-api-key
          CQ_DSN: ${{ secrets.CQ_DSN }} # Connection string to a PostgreSQL database
```

Once committed to the default branch of the repository, the above workflow will run daily at 3 a.m and will sync the AWS source integration with the PostgreSQL destination integration, in parallel.
