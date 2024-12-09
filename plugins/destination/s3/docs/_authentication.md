The plugin needs to be authenticated with your account(s) in order to sync information from your cloud setup.

The plugin requires only `PutObject` permissions (we will never make any changes to your cloud setup), so, following the principle of least privilege, it's recommended to grant it `PutObject` permissions.

There are multiple ways to authenticate with AWS, and the plugin respects the AWS credential provider chain. This means that CloudQuery will follow the following priorities when attempting to authenticate:

- The `AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`, `AWS_SESSION_TOKEN` environment variables.
- The `credentials` and `config` files in `~/.aws` (the `credentials` file takes priority).
- You can also use `aws sso` to authenticate cloudquery - [you can read more about it here](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-sso.html).
- IAM roles for AWS compute resources (including EC2 instances, Fargate and ECS containers).

You can read more about AWS authentication [here](https://aws.github.io/aws-sdk-go-v2/docs/configuring-sdk/#specifying-credentials) and [here](https://docs.aws.amazon.com/sdkref/latest/guide/creds-config-files.html).

### Environment Variables

CloudQuery can use the credentials from the `AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`, and
`AWS_SESSION_TOKEN` environment variables (`AWS_SESSION_TOKEN` can be optional for some accounts). For information on obtaining credentials, see the [AWS guide](https://aws.github.io/aws-sdk-go-v2/docs/getting-started/#get-your-aws-access-keys).

To export the environment variables (On Linux/Mac - similar for Windows):

```bash copy
export AWS_ACCESS_KEY_ID='{Your AWS Access Key ID}'
export AWS_SECRET_ACCESS_KEY='{Your AWS secret access key}'
export AWS_SESSION_TOKEN='{Your AWS session token}'
```

### Shared Configuration files

The plugin can use credentials from your `credentials` and `config` files in the `.aws` directory in your home folder.
The contents of these files are practically interchangeable, but CloudQuery will prioritize credentials in the `credentials` file.

For information about obtaining credentials, see the
[AWS guide](https://aws.github.io/aws-sdk-go-v2/docs/getting-started/#get-your-aws-access-keys).

Here are example contents for a `credentials` file:

```toml copy filename="~/.aws/credentials"
[default]
aws_access_key_id = YOUR_ACCESS_KEY_ID
aws_secret_access_key = YOUR_SECRET_ACCESS_KEY
```

You can also specify credentials for a different profile, and instruct CloudQuery to use the credentials from this profile instead of the default one.

For example:

```toml copy filename="~/.aws/credentials"
[myprofile]
aws_access_key_id = YOUR_ACCESS_KEY_ID
aws_secret_access_key = YOUR_SECRET_ACCESS_KEY
```

Then, you can either export the `AWS_PROFILE` environment variable (On Linux/Mac, similar for Windows):

```bash copy
export AWS_PROFILE=myprofile
```

You can also use the `local_profile` field in plugin configuration (can be helpful for syncing between different accounts).

### IAM Roles for AWS Compute Resources

The plugin can use IAM roles for AWS compute resources (including EC2 instances, Fargate and ECS containers).
If you configured your AWS compute resources with IAM, the plugin will use these roles automatically.
For more information on configuring IAM, see the AWS docs [here](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/iam-roles-for-amazon-ec2.html) and [here](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-iam-roles.html).

### User Credentials with MFA

In order to leverage IAM User credentials with MFA, the STS "get-session-token" command may be used with the IAM User's long-term security credentials (Access Key and Secret Access Key). For more information, see [here](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/sts/get-session-token.html).

```bash copy
aws sts get-session-token --serial-number <YOUR_MFA_SERIAL_NUMBER> --token-code <YOUR_MFA_TOKEN_CODE> --duration-seconds 3600
```

Then export the temporary credentials to your environment variables.

```bash copy
export AWS_ACCESS_KEY_ID=<YOUR_ACCESS_KEY_ID>
export AWS_SECRET_ACCESS_KEY=<YOUR_SECRET_ACCESS_KEY>
export AWS_SESSION_TOKEN=<YOUR_SESSION_TOKEN>
```

### Using a Custom S3 Endpoint

If you are using a custom S3 endpoint, you can specify it using the `endpoint` spec option. If you're using authentication, the `region` option in the spec determines the signing region used.
