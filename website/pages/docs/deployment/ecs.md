---
title: Running CloudQuery on Amazon ECS
tag: tutorial
date: 2023/03/03
---

# Running CloudQuery on Amazon ECS

In this tutorial we will be deploying CloudQuery on AWS ECS using Fargate. You will be using the AWS CLI to create the required resources. You can also use the AWS Management Console to create the resources. At the end of the tutorial you will have a CloudQuery instance running on AWS ECS that will collect data from your AWS account and store it in an S3 bucket. You can then query the data using Athena.

## Prerequisites
Before starting the deployment process, you need to have the following prerequisites:
  * An AWS account
  * AWS CLI installed and configured
  * Basic understanding of AWS ECS and Fargate

## Step 1: Generate a CloudQuery Configuration File

Create a new file named `cloudquery.yml` with the following content:
```yaml
kind: source
spec:
  # Source spec section
  name: aws
  path: "cloudquery/aws"
  version: "VERSION_SOURCE_AWS"
  tables: ["*"]
  destinations: ["s3"] 
---
kind: destination
spec:
  name: "s3"
  path: "cloudquery/s3"
  version: "v2.2.3"
  write_mode: "append"
  spec:
    bucket: <REPLACE_WITH_S3_DESTINATION_BUCKET>
    path: "{{TABLE}}/{{UUID}}.parquet"
    format: "parquet"
    athena: true
```

This will create a configuration file that will instruct CloudQuery to collect data from AWS and store it in an S3 bucket. You will need to replace the `REPLACE_WITH_S3_DESTINATION_BUCKET` placeholder with the name of the S3 bucket you want to use to store the data. You can also modify the configuration file to collect only the data you need. For more information on how to create a configuration file, [visit our docs](/docs/reference/source-spec)


In order to inject the config file into the prebuilt container you will have to base64 encode the content of the `cloudquery.yml` file . Assuming you are running on a Linux or MacOS machine you can do this conversion by running the following command:
```bash
cat cloudquery.yml | base64
```

## Step 2: Create an ECS Cluster
The first step in deploying CloudQuery on AWS ECS is to create an ECS cluster. To create an ECS cluster, use the following command:

Prior to running replace `<REPLACE_WITH_ECS_CLUSTER_NAME>` with the name you want to give to your ECS cluster.
```bash
aws ecs create-cluster --cluster-name <REPLACE_WITH_ECS_CLUSTER_NAME>
```


## Step 3: Create a Log Group
The next step is to create a log group for your ECS task. To create a log group, use the following command:

Prior to running replace `<REPLACE_WITH_LOG_GROUP_NAME>` with the name you want to give to your log group.
```bash
aws logs create-log-group --log-group-name <REPLACE_WITH_LOG_GROUP_NAME>
```

## Step 4: Set Log Group Retention
After creating a log group, you need to set the retention policy for your log group. To set the retention policy, use the following command:

Replace `<REPLACE_WITH_LOG_GROUP_NAME>` with the name of the log group that you created in Step 3.

```bash
aws logs put-retention-policy --log-group-name <REPLACE_WITH_LOG_GROUP_NAME> --retention-in-days 14
```
This command will set the retention period for your log group to 14 days. You can modify the retention period based on your requirements.

## Step 5: Create an IAM Role
To allow the ECS task to access the required AWS services, you need to create an IAM role.

Create a new file named `task-role-trust-policy.json` with the following contents:
```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "",
      "Effect": "Allow",
      "Principal": {
        "Service": "ecs-tasks.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}
```

Use the file you just created to create an IAM role for your ECS task. To create an IAM role, use the following command:

```bash
# Prior to running the following command, make sure you have replaced the <REPLACE_WITH_TASK_ROLE_NAME> placeholder with the name of the IAM role you want to create.
aws iam create-role --role-name <REPLACE_WITH_TASK_ROLE_NAME> --assume-role-policy-document file://task-role-trust-policy.json;
```

Store the ARN of the IAM role you just created. You will need it in the next step.

Create a new file named `data-access.json` with the following content:
``` json
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Action": [
                "s3:PutObject"
            ],
            "Resource": [
                "arn:aws:s3:::<REPLACE_WITH_S3_DESTINATION_BUCKET>/*"
            ],
            "Effect": "Allow"
        },
        {
            "Action": [
                "s3:GetObject"
            ],
            "Effect": "Deny",
            "NotResource": [
                "arn:aws:s3:::<REPLACE_WITH_S3_DESTINATION_BUCKET>/*"
            ]
        },
        {
            "Action": [
                "cloudformation:GetTemplate",
                "dynamodb:GetItem",
                "dynamodb:BatchGetItem",
                "dynamodb:Query",
                "dynamodb:Scan",
                "ec2:GetConsoleOutput",
                "ec2:GetConsoleScreenshot",
                "ecr:BatchGetImage",
                "ecr:GetAuthorizationToken",
                "ecr:GetDownloadUrlForLayer",
                "kinesis:Get*",
                "lambda:GetFunction",
                "logs:GetLogEvents",
                "sdb:Select*",
                "sqs:ReceiveMessage"
            ],
            "Resource": "*",
            "Effect": "Deny"
        }
    ]
}
```

Replace the `REPLACE_WITH_S3_DESTINATION_BUCKET` placeholder with the name of the S3 bucket you want to use to store the data.


This policy will allow the ECS task to write data to the S3 bucket you specified while also ensuring that CloudQuery never has access to any of your data.


Using the IAM policy that you just defined in `data-access.json`, you are going to attach it directly to the IAM role that the Fargate task will use. On top of the custom in-line policy you will also attach the `ReadOnlyAccess` policy and the `AmazonECSTaskExecutionRolePolicy` policy.
```bash
# Prior to running the following commands, make sure you have replaced the <REPLACE_TASK_ROLE_NAME> placeholder with the name of the IAM role you created earlier in this step.
aws iam put-role-policy --role-name < REPLACE_WITH_TASK_ROLE_NAME> --policy-name DenyData --policy-document file://data-access.json;
aws iam attach-role-policy --role-name < REPLACE_WITH_TASK_ROLE_NAME> --policy-arn arn:aws:iam::aws:policy/ReadOnlyAccess
aws iam attach-role-policy --role-name < REPLACE_WITH_TASK_ROLE_NAME>  --policy-arn arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy 

```

At this point you have a single IAM role that will be used by the Fargate task to access the required AWS services.

## Step 6: Register a Task Definition
A task definition is a blueprint that defines one or more containers that run together on the same host machine. In this step, we will create a task definition for the CloudQuery container.

In the task that you are going to create you will override the default entrypoint and create a custom command that will run the `cloudquery sync` command. This will enable you to store the configuration file as an environment variable, rather than having to create a custom image with the configuration file baked in.

You will also need to pass the CloudQuery configuration file to the container. To do that, you will need to base64 encode the configuration file and pass it as an environment variable to the container.

Create a new file named `task-definition.json` with the following content:
```json
{
  "containerDefinitions": [
    {
      "name": "ScheduledWorker",
      "image": "ghcr.io/cloudquery/cloudquery:<REPLACE_CQ_CLI_VERSION>",
      "command": [
        "/bin/sh",
        "-c",
        "echo $CQ_CONFIG| base64 -d  > ./file.yml;/app/cloudquery sync ./file.yml --log-console --log-format json"
      ],
      "environment": [
        { "name": "CQ_CONFIG", "value": "<REPLACE_CQ_BASE64_ENCODED_CONFIG>" }
      ],
      "logConfiguration": {
        "logDriver": "awslogs",
        "options": {
          "awslogs-group": "<REPLACE_LOG_GROUP_NAME>",
          "awslogs-region": "<REPLACE_AWS_REGION>",
          "awslogs-stream-prefix": "<REPLACE_PREFIX_FOR_STREAM>"
        }
      },
      "entryPoint": [""]
    }
  ],
  "family": "<REPLACE_TASK_FAMILY_NAME>",
  "requiresCompatibilities": ["FARGATE"],
  "cpu": "1024",
  "memory": "2048",
  "networkMode": "awsvpc",
  "taskRoleArn": "<REPLACE_TASK_ROLE_ARN>",
  "executionRoleArn": "<REPLACE_TASK_ROLE_ARN>"
}
```
Replace the following placeholders:
  - `<REPLACE_TASK_ROLE_ARN>` : The full arn of the role you created in Step 5.
  - `<REPLACE_CQ_CLI_VERSION>` : The version of the CloudQuery CLI you want to use. You can find the latest version [here](LINK TO GHCR)
  - `<REPLACE_CQ_BASE64_ENCODED_CONFIG>` : The base64 encoded version of the CloudQuery configuration file you created in Step 1.
  - `<REPLACE_LOG_GROUP_NAME>` : The name of the CloudWatch log group you created in Step 4.
  - `<REPLACE_AWS_REGION>` : The AWS region where you created the CloudWatch log group in Step 4.
  - `<REPLACE_PREFIX_FOR_STREAM>` : The prefix you want to use for the CloudWatch log stream.
  - `<REPLACE_TASK_FAMILY_NAME>` : The name of the task family you want to use.


Once you have modified the `task-definition.json` file to include the correct values for your environment, you can register the task definition with AWS ECS using the following command:
```bash

aws ecs register-task-definition --cli-input-json file://task-definition.json

```
This command registers the task definition with AWS ECS and returns the task definition's ARN, which you will use in the next step when you run the task.

## Step 7: Run the CloudQuery Task on ECS
Now that the task definition is registered, it's time to run the CloudQuery task on ECS using the `aws ecs run-task` command.

```bash
aws ecs run-task \
  --cluster <REPLACE_ECS_CLUSTER_NAME> \
  --task-definition <REPLACE_TASK_ARN> \
  --launch-type FARGATE \
  --network-configuration 'awsvpcConfiguration={subnets=[<REPLACE_SUBNET_1>,<REPLACE_SUBNET_2>],securityGroups=[<REPLACE_SG_1>,<REPLACE_SG_2>]}'
```
Replace the following placeholders: 
  - `<REPLACE_ECS_CLUSTER_NAME>` with the name of the ECS cluster you created in Step 2
  - `<REPLACE_TASK_ARN>` with the ARN of the task definition you registered in Step 6
  - `<REPLACE_SUBNET_1>` and `<REPLACE_SUBNET_2>` with the IDs of the subnets in which you want to run the task. You can specify any number of subnets that you want
  - `<REPLACE_SG_1>` and `<REPLACE_SG_2>` with the IDs of the security groups for the task. You can specify any number of security groups that you want


## Step 8: Schedule the Task to Run on a Regular Basis

Now that you have a task that runs CloudQuery, you can schedule it to run on a regular basis using AWS EventBridge scheduler. An EventBridge schedule is able to start a task on a regular basis, but to do so it needs a role that it can assume which has the `ecs:RunTask` permission. In this step, you will create a role that has the required permissions and then you will create a schedule that will run the task on a regular basis.


Create a file named `trust-policy.json` with the following content:
```json
{
  "Version": "2012-10-17",
  "Statement": [{
    "Version": "2012-10-17",
    "Statement": [{
      "Effect": "Allow",
      "Principal": {
        "Service": "scheduler.amazonaws.com"
      },
      "Action": "sts:AssumeRole",
      "Condition": {
        "StringEquals": {
          "aws:SourceArn": "arn:aws:scheduler:<REPLACE_AWS_REGION>:<REPLACE_AWS_ACCOUNT_ID>:schedule/default/<REPLACE_SCHEDULE_NAME>"
        }
      }
    }]
  }]
}
```

Replace the following placeholders:
  - `<REPLACE_AWS_REGION>` with the AWS region where you deploying this solution.
  - `<REPLACE_AWS_ACCOUNT_ID>` with the AWS account ID where you are deploying this solution.
  - `<REPLACE_SCHEDULE_NAME>` with the name of the schedule you will create later in this step.

Create an IAM role for the EventBridge scheduler using the following command:
```bash
aws iam create-role --role-name <REPLACE_EVENTBRIDGE_SCHEDULER_ROLE_NAME> --assume-role-policy-document file://trust-policy.json
```
Replace the `<REPLACE_EVENTBRIDGE_SCHEDULER_ROLE_NAME>` placeholder with the name of the role you want to create.

After creating a role with a trust policy that enables the scheduler service to assume it you will attach an inline policy that allows the scheduler to run the task you created. To do so
create a file named `ECSExecPolicy.json` with the following content:
```json
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "ecs:RunTask"
            ],
            "Resource": [
                "<REPLACE_ECS_TASK_ARN>"
            ],
            "Condition": {
                "ArnLike": {
                    "ecs:cluster": "<REPLACE_ECS_CLUSTER_ARN>"
                }
            }
        }
    ]
}
```
Replace the following placeholders:
  - `<REPLACE_ECS_TASK_ARN>` with the ARN of the task you created in Step 6.
  - `<REPLACE_ECS_CLUSTER_ARN>` with the ARN of the ECS cluster you created in Step 2.


```bash
aws iam put-role-policy --role-name <REPLACE_EVENTBRIDGE_SCHEDULER_ROLE_NAME> --policy-name ECSExecPolicy --policy-document file://ECSExecPolicy.json
```
Replace the `<REPLACE_EVENTBRIDGE_SCHEDULER_ROLE_NAME>` placeholder with the name of the role you just created.

Now that you have an IAM role that the scheduler service create the following JSON file that will contain all of the arguments required to create the schedule:
```json
{
    "Name": "<REPLACE_SCHEDULE_NAME>",
    "Description": "",
    "State": "ENABLED",
    "ScheduleExpression": "rate(24 hours)",
    "ScheduleExpressionTimezone": "UTC",
    "FlexibleTimeWindow": {
      "Mode": "OFF"
    },
    "Target": {
      "RoleArn": "<REPLACE_EVENTBRIDGE_SCHEDULER_ROLE_ARN>",
      "RetryPolicy": {
        "MaximumRetryAttempts": 1
      },
      "Arn": "<REPLACE_ECS_CLUSTER_ARN>",
      "Input": "{}",
      "EcsParameters": {
        "LaunchType": "FARGATE",
        "PlatformVersion": "1.4",
        "TaskCount": 1,
        "NetworkConfiguration": {
          "AwsvpcConfiguration": {
            "Subnets": [
              "<REPLACE_SUBNET_1>",
              "<REPLACE_SUBNET_2>"
            ],
            "SecurityGroups": [
              "<REPLACE_SG_1>",
              "<REPLACE_SG_2>"
            ],
          }
        },
        "TaskDefinitionArn": "<REPLACE_ECS_TASK_ARN>"
      }
    }
  }
```
Replace the following placeholders:
  - `<REPLACE_SCHEDULE_NAME>` with the name of the schedule you want to create.
  - `<REPLACE_EVENTBRIDGE_SCHEDULER_ROLE_ARN>` with the ARN of the role you created in the previous step.
  - `<REPLACE_ECS_CLUSTER_ARN>` with the ARN of the ECS cluster you created in Step 2.
  - `<REPLACE_SUBNET_1>` and `<REPLACE_SUBNET_2>` with the IDs of the subnets in which you want to run the task. You can specify any number of subnets that you want
  - `<REPLACE_SG_1>` and `<REPLACE_SG_2>` with the IDs of the security groups for the task. You can specify any number of security groups that you want
  - `<REPLACE_ECS_TASK_ARN>` with the ARN of the task you created in Step 6.

Finally, create the schedule using the following command: 

```bash
aws scheduler create-schedule  --cli-input-json file://scheduler-params.json
```



## Conclusion

You now have a working CloudQuery deployment that runs on a regular basis and stores the results in an S3 bucket. This is a great base for iterating on to learn more about performance you can check out the [performance tuning guide](/docs/advanced-topics/performance-tuning).

If you have any questions or comments, please feel free to reach out to us on [GitHub](https://github.com/cloudquery/cloudquery) or [Discord](https://cloudquery.io/discord)!
