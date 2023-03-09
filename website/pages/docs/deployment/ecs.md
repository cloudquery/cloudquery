---
title: Running CloudQuery on Amazon ECS
tag: tutorial
date: 2023/03/03
---


## Prerequisites
Before starting the deployment process, you need to have the following prerequisites:
  * An AWS account
  * AWS CLI installed on your local machine
  * Basic understanding of AWS ECS and Fargate

## Step 1: Create an ECS Cluster
The first step in deploying CloudQuery on AWS ECS is to create an ECS cluster. To create an ECS cluster, use the following command:
```bash
aws ecs create-cluster --cluster-name &lt;REPLACE_ECS_CLUSTER_NAME&gt;
```

Replace `&lt;REPLACE_ECS_CLUSTER_NAME&gt;` with the name you want to give to your ECS cluster.

## Step 2: Create a Log Group
The next step is to create a log group for your ECS task. To create a log group, use the following command:
```bash

aws logs create-log-group --log-group-name &lt;REPLACE_LOG_GROUP_NAME&gt;

```
Replace `&lt;REPLACE_LOG_GROUP_NAME&gt;` with the name you want to give to your log group.

## Step 3: Set Log Group Retention
After creating a log group, you need to set the retention policy for your log group. To set the retention policy, use the following command:
```bash
aws logs put-retention-policy --log-group-name &lt;REPLACE_LOG_GROUP_NAME&gt; --retention-in-days 14
```
This command will set the retention period for your log group to 14 days. You can modify the retention period based on your requirements.

## Step 4: Create an IAM Role
To allow the ECS task to access the required AWS services, you need to create an IAM role. To create an IAM role, use the following commands:
```bash

aws iam create-role --role-name &lt;REPLACE_TASK_ROLE_NAME&gt; --assume-role-policy-document file://task-role-trust-policy.json;

aws iam put-role-policy --role-name &lt;REPLACE_TASK_ROLE_NAME&gt; --policy-name DenyData --policy-document file://deny-data.json;

aws iam put-role-policy --role-name &lt;REPLACE_TASK_ROLE_NAME&gt; --policy-name WriteDataToS3Destination --policy-document file://write-data-s3-destination.json;

aws iam attach-role-policy --role-name &lt;REPLACE_TASK_ROLE_NAME&gt; --policy-arn arn:aws:iam::aws:policy/ReadOnlyAccess

aws iam attach-role-policy --role-name &lt;REPLACE_TASK_ROLE_NAME&gt;  --policy-arn arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy 

```
Replace `&lt;REPLACE_TASK_ROLE_NAME&gt;` with the name you want to give to your IAM role. The `task-role-trust-policy.json`, `deny-data.json`, and `write-data-s3-destination.json` files contain the IAM policy documents that define the permissions for the IAM role.


## Step 5: Register a Task Definition
A task definition is a blueprint that defines one or more containers that run together on the same host machine. In this step, we will create a task definition for the CloudQuery container.
Create a new file named `task-definition.json` with the following contents:
```json

{
  "containerDefinitions": [
    {
      "name": "ScheduledWorker",
      "image": "ghcr.io/cloudquery/cloudquery:&lt;REPLACE_CQ_CLI_VERSION&gt;",
      "command": [
        "/bin/sh",
        "-c",
        "echo $CQ_CONFIG| base64 -d  &gt; ./file.yml;/app/cloudquery sync ./file.yml --log-console --log-format json"
      ],
      "environment": [
        { "name": "CQ_CONFIG", "value": "&lt;REPLACE_CQ_BASE64_ENCODED_CONFIG&gt;" }
      ],
      "logConfiguration": {
        "logDriver": "awslogs",
        "options": {
          "awslogs-group": "&lt;REPLACE_LOG_GROUP_NAME&gt;",
          "awslogs-region": "&lt;REPLACE_AWS_REGION&gt;",
          "awslogs-stream-prefix": "&lt;REPLACE_PREFIX_FOR_STREAM&gt;"
        }
      },
      "entryPoint": [""]
    }
  ],
  "family": "&lt;REPLACE_TASK_FAMILY_NAME&gt;",
  "requiresCompatibilities": ["FARGATE"],
  "cpu": "1024",
  "memory": "2048",
  "networkMode": "awsvpc",
  "taskRoleArn": "&lt;REPLACE_TASK_ROLE_ARN&gt;",
  "executionRoleArn": "&lt;REPLACE_TASK_ROLE_ARN&gt;"
}

```
The `containerDefinitions` section defines the container that runs the CloudQuery CLI. In this case, the container is named "ScheduledWorker", and it uses the official CloudQuery Docker image from GitHub Container Registry.
The `command` section specifies the command that will be run in the container. The command downloads the CloudQuery configuration file, syncs the data, and logs the output in JSON format to the specified AWS CloudWatch Logs group.
The `environment` section specifies environment variables that are passed to the container. In this case, the environment variable `CQ_CONFIG` contains the Base64-encoded configuration file.
The `logConfiguration` section specifies the logging configuration for the container. In this case, the logs are sent to the specified AWS CloudWatch Logs group in the specified AWS region, with a specified prefix for the log stream name.
The remaining fields specify the task definition's name, compatibility, CPU and memory requirements, network mode, and the IAM roles associated with the task.
Once you have modified the `task-definition.json` file to include the correct values for your environment, you can register the task definition with AWS ECS using the following command:
```bash

aws ecs register-task-definition --cli-input-json file://task-definition.json

```
This command registers the task definition with AWS ECS and returns the task definition's ARN, which you will use in the next step when you run the task.

## Step 6: Run the ECS Task
After registering the task definition, you can run the ECS task. To run the ECS task, use the following command:
```bash

aws ecs run-task \
  --cluster &lt;CLUSTER NAME&gt; \
  --task-definition &lt;TASK_ARN&gt; \
  --launch-type FARGATE \
  --network-configuration 'awsvpcConfiguration={subnets=[&lt;SUBNET_1&gt;,&lt;SUBNET_2&gt;],securityGroups=[&lt;SG_1&gt;,&lt;SG_
```