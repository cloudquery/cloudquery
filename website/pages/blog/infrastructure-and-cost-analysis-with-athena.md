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

Hopefully you now have a directory called `cq_csv_output` that contains a number of CSV files. We need to upload these to S3 so that Athena can query them. To do so, we'll use the AWS CLI in this tutorial, but you can also use the AWS web console if you prefer.

First, we'll create a bucket to store the data:

```bash copy
aws s3 mb s3://cloudquery-athena-example
```

Now we can upload the files to the bucket. Athena requires every table to be in its own directory, so we'll use a little bash scripting to create upload each CSV file into its own directory:

```bash copy
for file in cq_csv_output/*.csv; do
    filename=$(basename "$file")
    foldername="${filename%.*}"
    aws s3 cp "$file" "s3://cloudquery-athena-example/$foldername/$filename"
done
```

You should now see a large number of objects in the bucket, which you can verify by listing the items:

``bash copy
aws s3 ls s3://cloudquery-athena-example
```

### 6. Create the Athena Database

Now that we have our data in S3, we need to create a database in Athena that will allow us to query it. To do so, we'll use the AWS CLI again:

