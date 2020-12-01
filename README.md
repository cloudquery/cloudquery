# cloudquery

<p align="center">
<img alt="cloudquery logo" width="200px" height="200px" src="https://github.com/cloudquery/cloudquery/raw/main/docs/images/logo.png" />
</p>

cloudquery exposes your cloud configuration and metadata as sql tables,
providing powerful analysis and monitoring without writing code.

### Links
* Homepage: https://cloudquery.io
* Releases: https://github.com/cloudquery/cloudquery/releases
* Documentation: https://docs.cloudquery.io

### Supported providers (Actively expanding)

Currently we support: [AWS](https://docs.cloudquery.io/aws), [GCP](https://docs.cloudquery.io/gcp), [Okta](https://docs.cloudquery.io/okta/table-reference) (Azure and DigitalOcean are on the roadmap)
If you want to us to add new provider please open an [Issue](https://github.com/cloudquery/cloudquery/issues).

## Download & install

You can download the precompiled binary from [releases](https://github.com/cloudquery/cloudquery/releases), or using CLI:

```bash
export VERSION=v0.2.0
export OS=Darwin # Possible values: Linux,Windows,Darwin
curl -L https://github.com/cloudquery/cloudquery/releases/download/${VERSION}/cloudquery_${OS}_x86_64 -o cloudquery
chmod a+x cloudquery
./cloudquery --help
```

## Quick Start

#### AWS 
You should be authenticated with an AWS account with correct permission with either option (see full [documentation](https://docs.aws.amazon.com/sdk-for-java/v1/developer-guide/credentials.html)):
 * `AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`
 * `~/.aws/credentials` created via `aws configure`
 
#### GCP

You should be authenticated with a GCP that has correct permissions for the data you want to pull.
You should set `GOOGLE_APPLICATION_CREDENTIALS` to point to your downloaded credential file.

#### Running

 Run the following commands:
 
```bash
cp example.config.yml config.yml
# uncomment resource of interest in config.yml
./cloudquery
 ```

If you uncommented images (it may take up to 30 seconds to fetch all images),
you can now run the following basic query in your sqlite3 console (`sqlite3 ./cloudquery.db`):

```sql
SELECT * FROM aws_ec2_images;
```

Full Documentation, resources and SQL schema definitions are available [here](https://docs.cloudquery.io)

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

## License

By contributing to cloudquery you agree that your contributions will be licensed as defined on the LICENSE file.

## Compile and run

```
go build .
./cloudquery # --help to see all options
```

## Contribution

Feel free to open Pull-Request for small fixes and changes. For bigger changes and new providers please open an issue first to prevent double work and discuss relevant stuff.
