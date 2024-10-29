---
title: Using an Offline License
description: Short walkthrough on how to use an offline license with CloudQuery.
---

# Offline Licensing

If you're using CloudQuery in an environment that doesn't have internet access, you can use an offline license with CloudQuery.
This guide will walk you through the process of obtaining and using an offline license.

## Obtaining an Offline License

To obtain an offline license, you'll need to contact our sales team by filling out the form in our [pricing page](https://www.cloudquery.io/pricing).

The license will be assigned to your organization and will be valid for a specific period of time.
If you need to use CloudQuery in a different organization or after the license has expired, you'll need to obtain a new license.

## Offline License File

Once you've obtained an offline license, you'll receive a file with a `.cqlicense` extension.  Using this file you will be able to run [cloudquery sync](/docs/reference/cli/cloudquery_sync) and [cloudquery migrate](/docs/reference/cli/cloudquery_migrate) commands without a connection to the CloudQuery API. 

## Using the Offline License

To use the offline license, you'll need to place the `.cqlicense` file (or files, as you may have more than one) into a common directory.  Then, when running `migrate` or `sync` commands, include the `--license` flag with the path to the directory containing the license files, or directly point it to a single license file.

For example:

```bash
mv mycompany.cqlicense /path/to/license/directory/
cloudquery sync --license /path/to/license/directory ./aws.yml ./pg.yml
```

or:

```bash
mv mycompany.cqlicense ~
cloudquery sync --license ~/mycompany.cqlicense ./aws.yml ./pg.yml
````

## Limitations of Using an Offline License

The offline license may be used only for [sync](/docs/reference/cli/cloudquery_sync) and [migrate](/docs/reference/cli/cloudquery_migrate) commands.
If you are setting up a new environment, you will need to have the integrations downloaded into a `.cq` directory: automatic integration downloads will **not** work with an offline license.
You will need to enable internet access and [login](/docs/reference/cli/cloudquery_login) (or, [generate](/docs/deployment/generate-api-key) and use an API key) and run `cloudquery plugin install` manually to install the integrations first. You may also just run a `sync` or `migrate` which will download the integrations, and then you can use the offline license for subsequent runs.

## Inspecting the License File

The license file is a JSON object and will contain a `license` field, which is base64 encoded data, and a `signature` field to ensure the authenticity of the file contents.
You can check the details of the license file using this command (requires the `jq` tool to decode JSON):

```bash
cat mycompany.cqlicense | jq -r .license | base64 -d | jq .
```

The result will look like this:

```json
{
  "licensed_to": "Your Company Name",
  "plugins": [
    "cloudquery/*"
  ],
  "issued_at": "2024-03-06T12:00:00Z",
  "valid_from": "2024-03-06T12:00:00Z",
  "expires_at": "2024-09-06T12:00:00Z"
}
```

If you have any questions about the contents of the license file, please contact our support team.
