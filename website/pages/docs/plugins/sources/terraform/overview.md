# Terraform Plugin

The CloudQuery Terraform plugin extracts terraform state and loads it into any supported CloudQuery destination (e.g. PostgreSQL).

## Configuration

To configure terraform to read a tfstate file, you need to create a `.yml` file in your cloudquery directory (e.g. named `terraform.yml`):

```yaml
kind: source
spec:
  # Source spec section
  name: terraform
  path: cloudquery/terraform
  version: "VERSION_SOURCE_TERRAFORM"
  tables: ["*"]
  destinations: ["postgresql"]
```

You can have multiple backends at the same time, by describing them in the configuration. Every configuration block describes one backend to handle.
CloudQuery currently supports LOCAL and S3 backends.

## Authentication for S3 backends

To authenticate CloudQuery with your Terraform state in S3 you can use any of the following options (see full documentation at [AWS SDK V2](https://aws.github.io/aws-sdk-go-v2/docs/configuring-sdk/#specifying-credentials)):

- Static Credentials: `AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`, `AWS_SESSION_TOKEN`
- Shared configuration files (via `aws configure`).
  - SDK defaults to `credentials` file under `.aws` folder that is placed in the home folder on your computer
  - SDK defaults to `config` file under `.aws` folder that is placed in the home folder on your computer
- If your application uses an ECS task definition or RunTask API operation, IAM role for tasks
- If your application is running on an Amazon EC2 instance, IAM role for Amazon EC2