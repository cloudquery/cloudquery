---
title: Generate an API Key
description: Generate an API Key for CloudQuery Hub and Premium integrations
---

# Generate a CloudQuery API Key

In automated environments, an API Key can be used with CloudQuery Hub and premium integrations instead of `cloudquery login`. 

Create an API key by following these steps:

1. Go to [CloudQuery Cloud](https://cloud.cloudquery.io) and log in or register for a new account.
2. If you are not already in a team, you will be prompted to create a new team.
3. Once you are logged in, click `Team Settings` in the left sidebar.
4. Go to the `API Keys` tab.
   ![API Keys](/images/docs/deployment/generate-api-key/01-api-keys.png)
5. Click `Generate new key`
   ![Generate new key](/images/docs/deployment/generate-api-key/02-generate-new-key.png)
6. Choose a key name for you to identify this key by, and an expiration date. Then click `Generate new key`.
   ![Generate key](/images/docs/deployment/generate-api-key/03-generate-key.png)
7. Copy the key and store it in a safe place. You will not be able to see it again after this step.
   ![Copy key](/images/docs/deployment/generate-api-key/04-save-key.png)

To use the API key with the CloudQuery CLI, set the `CLOUDQUERY_API_KEY` environment variable to the value of the key.

```bash
export CLOUDQUERY_API_KEY=<your-api-key>
```