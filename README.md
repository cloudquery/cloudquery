# cloudquery

<p align="center">
<img alt="cloudquery logo" width="200px" height="200px" src="https://github.com/cloudquery/cloudquery/raw/main/docs/images/logo.png" />
</p>

cloudquery transforms your cloud infrastructure into queryable SQL tables for easy monitoring, governance and security.

### What is cloudquery and use it?

cloudquery pulls, normalize, expose and monitor your cloud infrastructure and SaaS apps as SQL database.
This abstracts various scattered APIs enabling you to define security,governance,cost and compliance policies with SQL.

cloudquery can be easily extended to more resources and SaaS providers (open an [Issue](https://github.com/cloudquery/cloudquery/issues)). 

cloudquery comes with built-in policy packs such as: [AWS CIS](#running-policy-packs) (more is coming!).

### Links
* Homepage: https://cloudquery.io
* Releases: https://github.com/cloudquery/cloudquery/releases
* Documentation: https://docs.cloudquery.io

### Supported providers (Actively expanding)

Currently, we support: [AWS](https://docs.cloudquery.io/aws), [GCP](https://docs.cloudquery.io/gcp), [Okta](https://docs.cloudquery.io/okta/table-reference) (Azure and DigitalOcean are on the roadmap)
If you want to us to add a new provider please open an [Issue](https://github.com/cloudquery/cloudquery/issues).

## Download & install

You can download the precompiled binary from [releases](https://github.com/cloudquery/cloudquery/releases), or using CLI:

```bash
export OS=Darwin # Possible values: Linux,Windows,Darwin
curl -L https://github.com/cloudquery/cloudquery/releases/latest/download/cloudquery_${OS}_x86_64 -o cloudquery
chmod a+x cloudquery
./cloudquery --help

# if you want to download a specific version and not latest use the following endpoint
export VERSION= # specifiy a version
curl -L https://github.com/cloudquery/cloudquery/releases/download/${VERSION}/cloudquery_${OS}_x86_64 -o cloudquery
```

## Quick Start

### Running

First generate a `config.yml` file that will describe which resources you want cloudquery to pull, normalize
and transform resources to the specified SQL database by running the following command:
 
```shell script
./cloudquery gen config aws
# ./cloudquery gen config gcp okta # This will generate a config containing gcp and okta providers
# ./cloudquery gen config --help # Show all possible auto generated configs and flags
 ```

Once your `config.yml` is generated run the following command to fetch the resources:

```shell script
./cloudquery fetch
# ./cloudquery fetch --help # Show all possible fetch flags
```

If you used the default `sqlite` provider you run the following example queries

List ec2_images
```sql
SELECT * FROM aws_ec2_images;
```

Find all public facing AWS load balancers
```sql
SELECT * FROM aws_elbv2_load_balancers WHERE scheme = 'internet-facing';
```

#### Running policy packs

cloudquery comes with some ready compliance policy pack which you can use as is or modify to fit your use-case.

Currently, cloudquery support [AWS CIS](https://d0.awsstatic.com/whitepapers/compliance/AWS_CIS_Foundations_Benchmark.pdf)
policy pack (it is under active development, so it doesn't cover the whole spec yet).

To run AWS CIS pack enter the following commands (make sure you fetched all the resources beforehand by the `fetch` command):

```shell script
./cloudquery gen policy aws_cis
./cloudquery query 
``` 

Full Documentation, resources and SQL schema definitions are available [here](https://docs.cloudquery.io)

### Providers Authentication

#### AWS 
You should be authenticated with an AWS account with correct permission with either option (see full [documentation](https://docs.aws.amazon.com/sdk-for-java/v1/developer-guide/credentials.html)):
 * `AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`
 * `~/.aws/credentials` created via `aws configure`
 
Multi-account AWS support is available by using an account which can [AssumeRole](https://docs.aws.amazon.com/STS/latest/APIReference/API_AssumeRole.html) to other accounts.

In your config.yml you need to specify role_arns if you want to query multiple accounts in the following way:
```yaml
 accounts:
     - role_arn: <arn>
```
 
#### GCP

You should be authenticated with a GCP that has correct permissions for the data you want to pull.
You should set `GOOGLE_APPLICATION_CREDENTIALS` to point to your downloaded credential file.

#### Okta

You need to set `OKTA_TOKEN` environment variable

#### Query Examples

##### Find GCP buckets with public facing read permissions:

```sql
SELECT gcp_storage_buckets.name
FROM gcp_storage_buckets
         JOIN gcp_storage_bucket_policy_bindings ON gcp_storage_bucket_policy_bindings.bucket_id = gcp_storage_buckets.id
         JOIN gcp_storage_bucket_policy_bindings_members ON gcp_storage_bucket_policy_bindings_members.bucket_policy_binding_id = gcp_storage_bucket_policy_bindings.id
WHERE gcp_storage_bucket_policy_bindings_members.name = 'allUsers' AND gcp_storage_bucket_policy_bindings.role = 'roles/storage.objectViewer';
```

##### Find all public facing AWS load balancers

```sql
SELECT * FROM aws_elbv2_load_balancers WHERE scheme = 'internet-facing';
```

##### Find all unencrypted RDS instances

```sql
SELECT * from aws_rds_clusters where storage_encrypted = 0;
```

##### Find all unencrypted AWS buckets

```sql
SELECT * from aws_s3_buckets
    JOIN aws_s3_bucket_encryption_rules ON aws_s3_buckets.id != aws_s3_bucket_encryption_rules.bucket_id;
```

More examples are available [here](https://docs.cloudquery.io)

#### AWS Compliance Pack

There is a growing pack of compliance queries under [examples/aws_compliance_packs](https://github.com/cloudquery/cloudquery/tree/main/examples/aws_compliance_packs).
You can run them with the following commands:

```shell script
cp example.aws.config.yml
./cloudquery
./examples/aws_compliance_packs/compliance_test.sh
```

## License

By contributing to cloudquery you agree that your contributions will be licensed as defined on the LICENSE file.

## Compile and run

```
go build .
./cloudquery # --help to see all options
```

## Contribution

Feel free to open Pull-Request for small fixes and changes. For bigger changes and new providers please open an issue first to prevent double work and discuss relevant stuff.
