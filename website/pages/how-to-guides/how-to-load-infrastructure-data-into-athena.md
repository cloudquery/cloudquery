---
title: How to Load Infrastructure Data into Athena
tag: tutorial
description: Load CloudQuery source data into AWS Athena.  
author: hermanschaaf
---

import { HowToGuideHeader } from "../../components/HowToGuideHeader"

<HowToGuideHeader/>

# How to Load Infrastructure Data into Athena  

## Introduction

Athena is a serverless query service by AWS that allows you to query data in S3 using standard SQL. In this tutorial, we will show you how to load your cloud infrastructure data into S3 using CloudQuery and query it using Athena. This allows you to get fine-grained insights into your infrastructure data, all from the convenience of a serverless query environment running on AWS. 

By the end of this post, you will be able to query your AWS infrastructure data using Athena:

![Athena query editor](/images/how-to-guides/how-to-load-infrastructure-data-into-athena/athena-query-editor.png)

Let's get started!

## Steps

### Step 1: Install CloudQuery

To sync infrastructure data to S3, you will first need an installation of the CloudQuery CLI (or Docker image). See our [Quickstart](/docs/quickstart) for detailed instructions.

### Step 2: Create a Bucket for the Data

We will need to upload the sync results to S3 so that Athena can query them. We'll use the AWS CLI to do this in this tutorial, but you can also use the AWS web console or Terraform/CloudFormation if you prefer.

First, we'll create a bucket to store the data:

```bash copy
export BUCKET_NAME=cloudquery-athena-example
aws s3 mb s3://$BUCKET_NAME
```

### Step 3: Configure the AWS Source Plugin 

We will configure the CloudQuery AWS source plugin to sync data from your AWS account(s). To do so, we'll create a file called `aws.yml` with the following config:

```yaml copy title="aws.yml"
kind: source
spec:
  name: aws
  path: cloudquery/aws
  version: "VERSION_SOURCE_AWS"
  tables: ["*"]
  destinations: ["s3"]
```

This is the most basic configuration of the AWS source plugin. It will work as long as the AWS credentials you have configured in your environment have the appropriate permissions (e.g. via `aws sso login`). For more information on configuring the AWS source plugin, see the [AWS Source Plugin](/docs/plugins/sources/aws/overview) documentation.

### Step 4: Configure the S3 Destination Plugin

Similar to the config we created for the AWS plugin, we also need a destination config that is configured to write the AWS data to JSON files in S3:

```yaml copy title="s3.yml"
kind: destination
spec:
  name: "s3"
  path: "cloudquery/s3"
  version: "VERSION_DESTINATION_S3"
  write_mode: "append"
  spec:
    bucket: "${BUCKET_NAME}"
    path: "cloudquery/{{TABLE}}/{{UUID}}.json"
    format: "json"
```

For more information on configuring the S3 destination plugin, see the [S3 Destination Plugin](/docs/plugins/destinations/s3/overview) documentation.

### Step 5: Run CloudQuery sync

With the CLI installed, these two files in place, and an environment authenticated to AWS, we can now run the following command to sync our AWS infrastructure data to S3:

```bash copy
cloudquery sync aws.yml s3.yml
```

This will write the AWS data as JSON files to the S3 bucket. You can see the files that were created by running `aws s3 ls s3://$BUCKET_NAME`.

### Step 6: Create a Glue Crawler

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

We should also attach a policy to the role that gives it access to the S3 bucket we created in the previous step. The following policy will give the crawler access to the S3 bucket. Make sure to update the bucket name to the value of `BUCKET_NAME` you chose earlier (in our example, `cloudquery-athena-example`):

```json copy title="crawler-policy-s3-access.json"
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
    --policy-document file://crawler-policy-s3-access.json
```

The crawler will also need additional permissions to perform all its tasks, such as writing CloudWatch logs. We attach another policy to the role to give it these permissions. This is the same as the default policy created by AWS when creating a crawler through the console. You should review these permissions to ensure they are appropriate for your use case:

```json copy title="crawler-policy.json"
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "glue:*",
                "s3:GetBucketLocation",
                "s3:ListBucket",
                "s3:ListAllMyBuckets",
                "s3:GetBucketAcl",
                "ec2:DescribeVpcEndpoints",
                "ec2:DescribeRouteTables",
                "ec2:CreateNetworkInterface",
                "ec2:DeleteNetworkInterface",
                "ec2:DescribeNetworkInterfaces",
                "ec2:DescribeSecurityGroups",
                "ec2:DescribeSubnets",
                "ec2:DescribeVpcAttribute",
                "iam:ListRolePolicies",
                "iam:GetRole",
                "iam:GetRolePolicy",
                "cloudwatch:PutMetricData"
            ],
            "Resource": [
                "*"
            ]
        },
        {
            "Effect": "Allow",
            "Action": [
                "s3:CreateBucket"
            ],
            "Resource": [
                "arn:aws:s3:::aws-glue-*"
            ]
        },
        {
            "Effect": "Allow",
            "Action": [
                "s3:GetObject",
                "s3:PutObject",
                "s3:DeleteObject"
            ],
            "Resource": [
                "arn:aws:s3:::aws-glue-*/*",
                "arn:aws:s3:::*/*aws-glue-*/*"
            ]
        },
        {
            "Effect": "Allow",
            "Action": [
                "s3:GetObject"
            ],
            "Resource": [
                "arn:aws:s3:::crawler-public*",
                "arn:aws:s3:::aws-glue-*"
            ]
        },
        {
            "Effect": "Allow",
            "Action": [
                "logs:CreateLogGroup",
                "logs:CreateLogStream",
                "logs:PutLogEvents"
            ],
            "Resource": [
                "arn:aws:logs:*:*:/aws-glue/*"
            ]
        },
        {
            "Effect": "Allow",
            "Action": [
                "ec2:CreateTags",
                "ec2:DeleteTags"
            ],
            "Condition": {
                "ForAllValues:StringEquals": {
                    "aws:TagKeys": [
                        "aws-glue-service-resource"
                    ]
                }
            },
            "Resource": [
                "arn:aws:ec2:*:*:network-interface/*",
                "arn:aws:ec2:*:*:security-group/*",
                "arn:aws:ec2:*:*:instance/*"
            ]
        }
    ]
}
```

```bash copy
aws iam put-role-policy \
    --role-name cloudquery-athena-example-crawler \
    --policy-name cloudquery-athena-example-crawler \
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

With our crawler created, we can run it on demand like this:

```bash copy
aws glue start-crawler --name cloudquery-athena-example
```

(You can also run the crawler on a schedule, but we won't cover that here.)

### Step 7: Query the data

The crawler should have created a database and tables in the Glue Data Catalog. Now we can query the data using Athena! Let's use the AWS Console for this step. Navigate to the Athena service in the AWS Console, go to the Query Editor page, and select the database we created earlier. You should see a list of tables in the database. Let's run a simple query to find all S3 buckets that are permitted to be public:

```sql
-- Find all S3 buckets that are permitted to be public
SELECT "arn", "region"
FROM "cloudquery-athena-example"."aws_s3_buckets"
WHERE "block_public_acls" <> TRUE
    OR "block_public_policy" <> TRUE
    OR "ignore_public_acls" <> TRUE
    OR "restrict_public_buckets" <> TRUE
```

![Athena query editor](/images/how-to-guides/how-to-load-infrastructure-data-into-athena/athena-query-editor.png)

Wait, isn't that the bucket we just created for this guide? Better go lock it down 😉

## Conclusion

In this tutorial, we showed how to use CloudQuery to load infrastructure data into an S3 bucket, and then use Glue Crawler and Athena to query the data. This allows you to use the power of Athena to query your infrastructure data, and use the results to inform your security and compliance decisions.

One thing we haven't covered here is how data will be updated the next time you run a sync. The S3 destination plugin only supports the "append" write mode, which means it will keep adding new files and never remove the old ones. When querying data from the second sync onward, there will be duplicate entries. To deal with this, most queries should make use of the latest `_cq_sync_time`, a unique timestamp assigned to every sync:

```sql
-- Find all S3 buckets that are permitted to be public (updated to only use data from latest sync)
SELECT "arn", "region"
FROM "cloudquery-athena-example"."aws_s3_buckets"
WHERE ("block_public_acls" <> TRUE
    OR "block_public_policy" <> TRUE
    OR "ignore_public_acls" <> TRUE
    OR "restrict_public_buckets" <> TRUE)
    AND (_cq_sync_time = (select max(_cq_sync_time) from "cloudquery-athena-example"."aws_s3_buckets"))
```

In this way you will maintain a historical record of your data, but still be able to only query the latest data.

Going forward, we will continue making it easier and faster to load data into Athena, but also easier to maintain views over the latest data. Stay tuned for more updates!
