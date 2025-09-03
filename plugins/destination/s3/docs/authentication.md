---
hub-order: 1
---

## Quickstart

The plugin needs to be authenticated with your AWS account(s) in order to sync information from your cloud setup.

The plugin requires only `PutObject` permissions (we will never make any changes to your cloud setup), so, following the principle of least privilege, it's recommended to grant it `PutObject` permissions.

If you are running CloudQuery CLI locally, and have AWS CLI installed, [sign in with AWS CLI](https://docs.aws.amazon.com/signin/latest/userguide/command-line-sign-in.html).

Test that your AWS CLI is working:

```shell
aws account list-regions
```

### Non-interactive Authentication

There are multiple ways to authenticate with AWS, and the plugin respects the AWS credential provider chain. This means that AWS plugin will follow the following priorities when attempting to authenticate:

1. The `AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`, `AWS_SESSION_TOKEN` environment variables - see the [AWS guide](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_access-keys.html).
2. The `credentials` and `config` files in `~/.aws` (the `credentials` file takes priority)
3. A session created using the `aws sso` to authenticate the plugin - see [Configuring IAM Identity Center authentication with the AWS CLI
   ](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-sso.html)
4. IAM roles for AWS compute resources (including EC2 instances, Fargate and ECS containers)

For details about configuring the individual authentication options, see [Advanced Authentication Configuration](#advanced-authentication-configuration).

There are multiple ways to authenticate with AWS, and the plugin respects the AWS credential provider chain. This means that CloudQuery will follow the following priorities when attempting to authenticate:

- The `AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`, `AWS_SESSION_TOKEN` environment variables.
- The `credentials` and `config` files in `~/.aws` (the `credentials` file takes priority).
- You can also use `aws sso` to authenticate cloudquery - [you can read more about it here](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-sso.html).
- IAM roles for AWS compute resources (including EC2 instances, Fargate and ECS containers).
