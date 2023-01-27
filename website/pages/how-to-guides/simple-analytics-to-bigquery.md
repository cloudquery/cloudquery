---
title: How to Sync Simple Analytics Data to BigQuery using GitHub Actions
tag: tutorial
description: >-
  A guide on how to load Simple Analytics data into BigQuery using GitHub Actions
author: hermanschaaf
---

import { HowToGuideHeader } from "../../components/HowToGuideHeader"

<HowToGuideHeader/>

## Introduction

In this how-to guide, we will show you how you can load page views and events from Simple Analytics into BigQuery on regular schedule. We will use GitHub Actions to automate the process. This is a very convenient, simple and serverless way to load data into BigQuery.

## Step 0: Prerequisites

Before we start, to follow this guide you will need the following:

- a Simple Analytics account to load data from,
- a BigQuery dataset to sync to,
- a GitHub repository to store the workflow and config files, and
- the [`gcloud` CLI](https://cloud.google.com/sdk/gcloud) installed and configured. You may also choose to use the Google Cloud Console UI (or Terraform) instead, but we won't cover the steps for that here.  

Now, let's get started!

## Step 1: Set up the GitHub repository

### Step 1.1: Add the workflow file

Inside your GitHub repository, create a new file called `.github/workflows/sync_simple_analytics_to_bigquery.yml` and add the following content:

```yaml copy
name: Sync Simple Analytics to BigQuery

on:
  workflow_dispatch:
  schedule:
    - cron: "0 2 * * *"  # daily at 02:00 UTC

jobs:
  simple-analytics-to-bigquery:
    timeout-minutes: 60
    runs-on: ubuntu-latest
    permissions:
      id-token: 'write' # This required for OIDC
      contents: 'read' # This is required for actions/checkout@v3
    steps:
      - uses: actions/checkout@v3
        with:
          fetch-depth: 2
      - name: Authenticate to Google Cloud
        uses: 'google-github-actions/auth@v0'
        with:
          workload_identity_provider: 'projects/${{ secrets.SA_BIGQUERY_PROJECT_NUMBER }}/locations/global/workloadIdentityPools/cloudquery-pool/providers/cloudquery-provider'
          service_account: 'cloudquery-service-account@${{ secrets.SA_BIGQUERY_PROJECT_ID }}.iam.gserviceaccount.com'
      - uses: cloudquery/setup-cloudquery@v3
        name: Setup CloudQuery
        with:
          version: "v2.3.1"
      - name: CloudQuery Sync
        run: cloudquery sync --log-console simple-analytics.yml bigquery.yml
        env:
          SA_USER_ID: ${{ secrets.SA_USER_ID }}
          SA_API_KEY: ${{ secrets.SA_API_KEY }}
          SA_BIGQUERY_PROJECT_ID: ${{ secrets.SA_BIGQUERY_PROJECT_ID }}
          SA_BIGQUERY_DATASET_ID: ${{ secrets.SA_BIGQUERY_DATASET_ID }}
```

This configuration allows us to run the workflow manually or on a schedule. The workflow will run on a self-hosted runner, which is a virtual machine hosted by GitHub. The workflow will run the following steps:
 1. Check out the repository
 2. Authenticate to Google Cloud using the [google-github-actions/auth](https://github.com/google-github-actions/auth) Action
 3. Setup CloudQuery using the [cloudquery/setup-cloudquery](https://github.com/cloudquery/setup-cloudquery) Action
 4. Run the `cloudquery sync` command to load data from Simple Analytics into BigQuery

We will now configure the secrets and the configuration files.

### Step 1.2: Add the secrets

In the GitHub repository, go to `Settings > Secrets` and add the following repository secrets:

- `SA_USER_ID`: the user ID of your Simple Analytics account. You can find this in the Simple Analytics [Account Settings](https://simpleanalytics.com/account).
- `SA_API_KEY`: the API key of your Simple Analytics account. You can find this in the Simple Analytics [Account Settings](https://simpleanalytics.com/account).
- `SA_BIGQUERY_PROJECT_ID`: the ID of the Google Cloud project where your BigQuery dataset is located. You can find this in the Google Cloud Console home page when your project is selected.
- `SA_BIGQUERY_DATASET_ID`: the ID of the BigQuery dataset where you want to load the data. You can find this in the BigQuery UI.
- `SA_BIGQUERY_PROJECT_NUMBER`: the number of the Google Cloud project where your BigQuery dataset is located.  You can find this in the Google Cloud Console home page when your project is selected.

### Step 1.3: Add the Simple Analytics source plugin configuration file

In the GitHub repository, create a new file called `simple-analytics.yml` and add the following content:

```yaml copy
kind: source
spec:
  name: "simple-analytics"
  path: "simpleanalytics/simple-analytics"
  version: "v1.0.0"
  tables:
    ["*"]
  destinations:
    - "bigquery"
  spec:
    user_id: ${SA_USER_ID}
    api_key: ${SA_API_KEY}
    websites:
      - hostname: <your-website-hostname>
```

You may find the latest version on the releases page. Replace `<your-website-hostname>` with the hostname of your website as defined on Simple Analytics, e.g. `mywebsite.com`.

### Step 1.4: Add the BigQuery destination plugin configuration file

In the GitHub repository, create a new file called `bigquery.yml` and add the following content:

```yaml copy

TODO CONTINUE HERE

```

## Step 3: Set up Workload Identity

### Step 3.1: Set Environment Variables

First, let's set a few environment variables that we'll use throughout the guide. Set the following variables, replacing `<your-project-id>` and `<your-repo>` with your own values:

```bash copy
export PROJECT_ID=<your-project-id> # The ID of your Google Cloud project, e.g. "my-project"
export REPO=<your-repo> # e.g. cloudquery/simple-analytics-to-bigquery
```

### Step 3.2: Create a Service Account

We will need a service account that can write to BigQuery. We will use the `gcloud` CLI to create the service account and grant it the necessary permissions. Run the following commands:

```bash copy
gcloud iam service-accounts create "cloudquery-service-account" \
  --project "${PROJECT_ID}"
```

If you haven't used the IAM Credentials API in the project before, it will need to be enabled:

```bash copy
gcloud services enable iamcredentials.googleapis.com \
  --project "${PROJECT_ID}"
```

Now we will grant the Google Cloud Service Account permissions to access Google Cloud resources. In this case, BigQuery:

We will assign the service account the `BigQuery Data Editor` role: 

```bash copy
gcloud projects add-iam-policy-binding $PROJECT_ID --member="serviceAccount:cloudquery-service-account@${PROJECT_ID}.iam.gserviceaccount.com" --role=roles/bigquery.dataEditor
```

### Step 3.3: Create a Workload Identity Pool and Provider

We will now set up Workload Identity, which the GitHub action will rely on to authenticate itself without the need for keys. You can read more about keyless authentication with GCP and GitHub Actions [here](https://cloud.google.com/blog/products/identity-security/enabling-keyless-authentication-from-github-actions).

First we need to create a pool:

```bash copy
gcloud iam workload-identity-pools create "cloudquery-pool" \
  --project="${PROJECT_ID}" \
  --location="global" \
  --display-name="Cloudquery pool"
```

Now let's describe the pool and set an environment variable that we'll use in subsequent steps:

```bash copy
export WORKLOAD_IDENTITY_POOL_ID=$(gcloud iam workload-identity-pools describe "cloudquery-pool" \
  --project="${PROJECT_ID}" \
  --location="global" \
  --format="value(name)")
```

We can use `echo $WORKLOAD_IDENTITY_POOL_ID` to verify that the variable is set correctly:

```bash copy
echo $WORKLOAD_IDENTITY_POOL_ID
# projects/<your-project-id>/locations/global/workloadIdentityPools/cloudquery-pool
```

Now let's create an OIDC provider:

```bash copy
gcloud iam workload-identity-pools providers create-oidc "cloudquery-provider" \
  --project="${PROJECT_ID}" \
  --location="global" \
  --workload-identity-pool="cloudquery-pool" \
  --display-name="Cloudquery provider" \
  --attribute-mapping="google.subject=assertion.sub,attribute.actor=assertion.actor,attribute.repository=assertion.repository" \
  --issuer-uri="https://token.actions.githubusercontent.com"
```

Finally, we need to grant the service account the `Workload Identity User` role:

```bash copy
gcloud iam service-accounts add-iam-policy-binding "cloudquery-service-account@${PROJECT_ID}.iam.gserviceaccount.com" \
  --project="${PROJECT_ID}" \
  --role="roles/iam.workloadIdentityUser" \
  --member="principalSet://iam.googleapis.com/${WORKLOAD_IDENTITY_POOL_ID}/attribute.repository/${REPO}"
```
