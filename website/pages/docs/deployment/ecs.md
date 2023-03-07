---
title: Running CloudQuery on Amazon ECS
tag: tutorial
date: 2023/03/03
---

# Deploy with Amazon ECS

In this tutorial, we will show you how to load AWS resources into Amazon S3 by running CloudQuery as a service in [Amazon ECS](https://github.com/features/actions), using the AWS source- and S3 destination plugins.

## Prerequisites

1. S3 Bucket that will hold the data. If the bucket is in a different account than where you will be deploying the ECS service then you will need to ensure that there is a bucket policy that allows for cross account access.


## Creating the CloudQuery configuration file

Under the root of your repository, create a new `cloudquery.yml` file with the following content:

```yaml copy
kind: source
spec:
  name: 'aws'
  path: cloudquery/aws
  version: "VERSION_SOURCE_AWS"
  destinations: ['s3']
---
kind: destination
spec:
  name: 's3'
  path: cloudquery/s3
  version: "VERSION_DESTINATION_S3"
  spec:
    bucket: <DESTINATION_BUCKET_NAME>
```

You will need to replace `<DESTINATION_BUCKET_NAME>` with the name of the S3 Bucket where you want to store the data. 


> For more configuration options, [visit our docs](/docs/reference/source-spec)

## Solution Overview

![alt text](/website/data/ecs-deployment/cloudquery_on_ecs.png "Title")


The CloudFormation will stand up a number of resources including:



- ***ECS Service and Task definition***: This uses the base image that the [CloudQuery Team maintains](https://github.com/cloudquery/cloudquery/pkgs/container/cloudquery). On top of that image this solution provides an alternative entrypoint that enables users to specify the configuration file as a base64 encoded string as an environment variable. This allows users to not have to build and maintain a custom image.
- ***EventBridge Schedule***: A Cloud native scheduling service that enables users to directly trigger an ECS task to run at a specific time or rate with support for timezones and daylight savings time. 
- ***CloudWatch Dashboard***: A dashboard of useful information and metrics. This dashboard is powered by CloudWatch Insight querying the structured logs that are produced by each sync.
- ***VPC and supporting resources***: By default this solution will deploy all of the VPC and networking resources required for you to have a running sync at the end of the deployment.  


### Deploying Solution

Start off by cloning:
```bash
git clone https://github.com/cloudquery/<NOT SURE HOW WE WILL RELEASE THIS>
cd <DIRECTORY>
```

Deploy CloudFormation Template:

```bash
aws cloudformation deploy --template-file template.yml \
  --stack-name $STACKNAME \
  --capabilities CAPABILITY_NAMED_IAM CAPABILITY_AUTO_EXPAND \
  --parameter-overrides \
    DestinationS3Bucket=$BUCKET_NAME \
    CQConfiguration=$(cat cloudquery.yml| base64) 
```

This call takes the `cloudquery.yml` configuration and injects it as a base64 encoded environment variable, so make sure not to hard code any sensitive information in the file. If you need to specify sensitive information you should use `Environment Variable Substitution` [INCLUDE LINK TO DOCS] 


Once this has finished deploying CloudQuery will be executed every `24 hours`. To trigger a manual sync run the following command:
```bash
aws ecs run-task \
  --cluster <CLUSTER NAME> \
  --task-definition <TASK_ARN> \
  --launch-type FARGATE \
  --network-configuration 'awsvpcConfiguration={subnets=[<SUBNET_1>,<SUBNET_2>],securityGroups=[<SG_1>,<SG_2>]}'
```



## Monitoring:

As part of this solution there is a CloudWatch dashboard that helps users understand what is going on with the sync by showing metrics on how many resources have been synced and shows all of the errors that occurred during a sync:


[INCLUDE SCREENSHOT HERE]
