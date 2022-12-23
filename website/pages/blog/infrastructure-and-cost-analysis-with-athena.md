# Infrastructure and Cost Analysis with Athena

## Introduction

Athena is a serverless query service by AWS that allows you to query data in S3 using standard SQL. In this post, we will show you how to load your cloud infrastructure data into S3 using CloudQuery, combine it with cost reporting data, and query these together using Athena. This is a powerful combination that allows you to get fine-grained insight into the resources that cost you the most, all from a convenient serverless query environment. Let's get started!

## Steps

### 1. Install CloudQuery

To sync infrastructure data to S3, you will first need an installation of the CloudQuery CLI (or Docker image). See our [Quickstart](/docs/quickstart) for detailed instructions.

### 2. Configure the AWS Source Plugin 

We will configure the CloudQuery AWS source plugin to sync data from your AWS account(s). To do so, we'll create a file called `aws.yml` with the following config:

```yaml title="aws.yml"
kind: source
spec:
  name: aws
  path: cloudquery/aws
  version: "VERSION_SOURCE_AWS"
  tables: ["*"]
  destinations: ["csv"]
```

This is the most basic configuration of the AWS source plugin. It will work as long as the AWS credentials you have configured in your environment have the appropriate permissions (e.g. via `aws sso login`). For more information on configuring the AWS source plugin, see the [AWS Source Plugin](/docs/plugins/source/aws) documentation.

### 3. Configure the CSV Destination Plugin

Similar to the config we created for the AWS plugin, we also need a destination config that is configured to write the AWS data to CSV files inside a local directory called `cq_csv_output`:

```yaml title="csv.yml"
kind: destination
spec:
  name: "csv"
  path: "cloudquery/csv"
  version: "VERSION_DESTINATION_CSV"
  write_mode: "append"
  spec:
    directory: './`cq_csv_output`'
```

For more information on configuring the CSV destination plugin, see the [CSV Destination Plugin](/docs/plugins/destination/csv) documentation.

### 4. Run CloudQuery sync

With the CLI installed, these two files in place, and an enviroment authenticated to AWS, we can now run the following command to sync our AWS data to CSV files:

```bash copy
cloudquery sync aws.yml csv.yml
```

This will write the AWS data to CSV files in the `cq_csv_output` directory. You can see the files that were created by running `ls -al cq_csv_output`.

### 5. Upload the files to S3

You should now have a directory called `cq_csv_output` that contains a number of CSV files. We need to upload these to S3 so that Athena can query them. We'll use the AWS CLI to do this in this tutorial, but you can also use the AWS web console or Terraform/CloudFormation if you prefer.

First, we'll create a bucket to store the data:

```bash copy
export BUCKET_NAME=cloudquery-athena-example
aws s3 mb s3://$BUCKET_NAME
```

Now we can upload the files to the bucket. Athena requires every table to be in its own directory, so we'll use a little bash scripting to upload each CSV file into its own directory:

```bash copy
for file in cq_csv_output/*.csv; do
    filename=$(basename "$file")
    foldername="${filename%.*}"
    aws s3 cp "$file" "s3://$BUCKET_NAME/$foldername/$filename"
done
```

You should now see a large number of objects in the bucket, which we can verify by listing the items:

```bash copy
aws s3 ls s3://$BUCKET_NAME
```

### 7. Create a Glue Crawler

Athena can query data in S3, but it needs to know the schema of the data in order to do so. We can use a Glue Crawler to automatically infer the schema of the data and create a table in the Athena database we created in the previous step. We'll use the AWS CLI again.

First we'll create a role for the crawler to use. It will need access to the S3 bucket we created in the previous step. The following trust policy will give the role what it needs:

```json copy title="crawler-trust-policy.json"
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Principal": {
                "Service": "glue.amazonaws.com"
            },
            "Action": "sts:AssumeRole"
        }
    ]
}
```

With the above policy saved as `crawler-trust-policy.json`, we can now create the role (make a note of the ARN for the newly created role):

```bash copy
aws iam create-role \
    --role-name cloudquery-athena-example-crawler \
    --assume-role-policy-document file://crawler-trust-policy.json
```

We should also attach a policy to the role that gives it access to the S3 bucket we created in the previous step. The following policy will give the role everything it needs, including the ability to write CloudWatch logs. You should review and fine-tune these permissions before applying them. Also make sure to update the bucket name to the value of `BUCKET_NAME` you chose earlier (in our example, `cloudquery-athena-example`):

```json copy title="crawler-policy.json"
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "s3:GetObject",
                "s3:PutObject"
            ],
            "Resource": [
                "arn:aws:s3:::cloudquery-athena-example*"
            ]
        }
    ]
}
```

Let's attach this policy to the role:

```bash
aws iam put-role-policy \
    --role-name cloudquery-athena-example-crawler \
    --policy-name cloudquery-athena-example-crawler-s3-access \
    --policy-document file://crawler-policy.json
```

Now we can create the crawler, making sure to reference the ARN for the role we created above:

```bash copy
aws glue create-crawler \
    --name cloudquery-athena-example \
    --database-name cloudquery-athena-example \
    --role arn:aws:iam::123456789012:role/cloudquery-athena-example-crawler \
    --targets "S3Targets=[{Path=s3://$BUCKET_NAME}]" \
    --schema-change-policy "UpdateBehavior=UPDATE_IN_DATABASE,DeleteBehavior=DEPRECATE_IN_DATABASE"
```

With our crawler created, we can run it any time:

```text
aws glue start-crawler --name cloudquery-athena-example
```

### 6. Create the Athena Database

Now that we have our data in S3, we need to create a database in Athena that will allow us to query it. We'll use the AWS CLI again:

```bash copy
aws athena create-data-catalog --name cloudquery --type GLUE
```
