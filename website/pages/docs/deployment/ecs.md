---
title: Running CloudQuery on Amazon ECS
description: In this tutorial we will be deploying CloudQuery on AWS ECS using Fargate. You will be using the AWS CLI and AWS CloudFormation to create the required resources. At the end of the tutorial you will have a CloudQuery instance running on AWS ECS that will periodically collect data from your AWS account and store it in an S3 bucket. You can then query the data using Athena.
tag: tutorial
date: 2023/03/03
---

# Running CloudQuery on Amazon ECS

In this tutorial we will be deploying CloudQuery on AWS ECS using Fargate. You will be using the AWS CLI and AWS CloudFormation to create the required resources. At the end of the tutorial you will have a CloudQuery instance running on AWS ECS that will periodically collect data from your AWS account and store it in an S3 bucket. You can then query the data using Athena.

## Prerequisites
Before starting the deployment process, you need to have the following prerequisites:
  * AWS CLI installed and configured more information can be found [here](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-getting-started.html) 
  * Basic understanding of [AWS ECS and Fargate](https://docs.aws.amazon.com/AmazonECS/latest/userguide/what-is-fargate.html)

## Step 1: Generate a CloudQuery Configuration File

Create a new file named `cloudquery.yml` with the following content:
```yaml
kind: source
spec:
  # Source spec section
  name: aws
  path: "cloudquery/aws"
  registry: "cloudquery"
  version: "VERSION_SOURCE_AWS"
  tables: ["aws_s3_buckets"]
  destinations: ["s3"] 
---
kind: destination
spec:
  name: "s3"
  path: "cloudquery/s3"
  registry: "cloudquery"
  version: "VERSION_DESTINATION_S3"
  write_mode: "append"
  spec:
    bucket: <REPLACE_WITH_S3_DESTINATION_BUCKET>
    region: <REPLACE_WITH_AWS_REGION>
    path: "{{TABLE}}/{{UUID}}.parquet"
    format: "parquet"
    athena: true
```

This will create a configuration file that will instruct CloudQuery to collect data from AWS and store it in an S3 bucket. You will need to replace the `REPLACE_WITH_S3_DESTINATION_BUCKET` placeholder with the name of the S3 bucket you want to use to store the data. You can also modify the configuration file to collect only the data you need. For more information on how to create a configuration file, [visit our docs](/docs/reference/source-spec)


In order to inject the configuration file into the prebuilt container you will have to base64 encode the content of the `cloudquery.yml` file . Assuming you are running on a Linux or MacOS machine you can do this conversion by running the following command:
```bash
cat cloudquery.yml | base64
```


## Step 2: Create a Secret to store a CloudQuery API Key

Downloading integrations requires users to be authenticated, normally this means running `cloudquery login` but that is not doable in ECS. The recommended way to handle this is to use an API key. More information on generating an API Key can be found [here](/docs/deployment/generate-api-key).

Once you have a CloudQuery API Key you are going to create a Secret in AWS Secrets Manager to store the API Key. To create a secret, use the following command:

```bash
aws secretsmanager create-secret \
    --name CQ-APIKey \
    --description "API Key to authenticate with CloudQuery hub" \
    --secret-string "<CQ_API_KEY>"
```



## Step 3: Create a CloudFormation template
This template will create the required resources for the deployment of CloudQuery on AWS ECS. Create a new file named `cloudquery-ecs-resources.yaml` with the following content:

```yaml
AWSTemplateFormatVersion: "2010-09-09"
Description: "Deploy CloudQuery on AWS ECS with Fargate"

Parameters:
  CQVersion:
    Description: Please enter the version of CloudQuery you want to deploy. This should be in the format  X, X.Y, X.Y.Z, or `latest`
    Type: String
  ScheduleExpression:
    Description: Please enter the Eventbridge Schedule Expression. This can either be a Rate or a cron expression. This will define how often CloudQuery will run the sync.
    Type: String
    Default: "rate(24 hours)"
  PrivateSubnetIds:
    Description: Please enter a comma separated list of Subnet Ids where you want to run CloudQuery and the Database.
    Type: String
  SecurityGroupIds:
    Description: Please enter a comma separated list of Security Group IDs that you want to attach to the Fargate node.
    Type: String
  DestinationS3Bucket:
    Description: Please enter the name of the S3 bucket where you want to store the CloudQuery results
    Type: String
  CQConfiguration:
    Description: Please enter the CloudQuery configuration file encoded in base64
    Type: String
  CQApiKey:
    Description: ARN of the secret containing the CloudQuery API Key
    Type: String
  AWSMarketplace:
    Description: If you are using the AWS Marketplace version of CloudQuery, set this to true
    Type: String
    AllowedValues: [true, false]
    Default: false
Resources:
  #### ECS Cluster:
  ECSCluster:
    Type: "AWS::ECS::Cluster"

  ScheduledWorkerTask:
    Type: "AWS::ECS::TaskDefinition"
    Properties:
      RequiresCompatibilities:
        - FARGATE
      NetworkMode: awsvpc
      Cpu: 1024
      Memory: 2GB
      ExecutionRoleArn: !GetAtt ExecutionRole.Arn
      TaskRoleArn: !GetAtt ExecutionRole.Arn
      ContainerDefinitions:
        - Essential: "true"
          Command:
            - "echo $CQ_CONFIG| base64 -d  > ./file.yml;/app/cloudquery sync ./file.yml --log-console --log-format json"
          Image: !Sub ghcr.io/cloudquery/cloudquery:${CQVersion}
          Name: ScheduledWorker
          EntryPoint:
            - "/bin/sh"
            - "-c"
          Environment:
            - Name: CQ_INSTALL_SRC
              Value: CLOUDFORMATION
            - Name: CQ_CONFIG
              Value: !Ref CQConfiguration
            - Name: CQ_AWS_MARKETPLACE_CONTAINER
              Value: !Ref AWSMarketplace
          Secrets:
            - ValueFrom: !Ref CQApiKey
              Name: CLOUDQUERY_API_KEY
          LogConfiguration:
            LogDriver: awslogs
            Options:
              awslogs-group: !Ref LogGroup
              awslogs-region: !Ref AWS::Region
              awslogs-stream-prefix: !Ref AWS::StackName
  LogGroup:
    DeletionPolicy: Retain
    UpdateReplacePolicy: Retain
    Type: AWS::Logs::LogGroup
    Properties:
      RetentionInDays: 14
  # Scheduler Configurations
  Schedule:
    Type: AWS::Scheduler::Schedule
    Properties:
      FlexibleTimeWindow:
        Mode: "OFF"
      ScheduleExpression: !Ref ScheduleExpression
      State: ENABLED
      Target:
        Arn: !GetAtt ECSCluster.Arn
        EcsParameters:
          NetworkConfiguration:
            AwsvpcConfiguration:
              AssignPublicIp: DISABLED
              SecurityGroups: !Split [",", !Ref SecurityGroupIds]
              Subnets: !Split [",", !Ref PrivateSubnetIds]
          LaunchType: FARGATE
          PlatformVersion: 1.4.0
          TaskCount: 1
          TaskDefinitionArn: !Ref ScheduledWorkerTask
        RoleArn: !GetAtt SchedulerExecutionRole.Arn

  SchedulerExecutionRole:
    Type: AWS::IAM::Role
    Properties:
      AssumeRolePolicyDocument:
        Version: "2012-10-17"
        Statement:
          - Action:
              - sts:AssumeRole
            Effect: Allow
            Principal:
              Service:
                - scheduler.amazonaws.com
      Path: "/"
      Policies:
        - PolicyName: SchedulerExecutionRole
          PolicyDocument:
            Version: "2012-10-17"
            Statement:
              - Action:
                  - ecs:RunTask
                Effect: Allow
                Resource: !Ref ScheduledWorkerTask
                Condition:
                  ArnEquals:
                    ecs:cluster: !GetAtt ECSCluster.Arn
              - Effect: Allow
                Action:
                  - iam:PassRole
                Resource: !GetAtt ExecutionRole.Arn

  ####### IAM role for Fargate execution#####
  ExecutionRole:
    Type: AWS::IAM::Role
    Properties:
      AssumeRolePolicyDocument:
        Statement:
          - Effect: Allow
            Principal:
              Service: ecs-tasks.amazonaws.com
            Action: "sts:AssumeRole"
      Policies:
        - PolicyName: AccessAPIKeySecret
          PolicyDocument:
            Version: 2012-10-17
            Statement:
              - Effect: Allow
                Action:
                  - secretsmanager:GetSecretValue
                Resource: !Sub ${CQApiKey}
        - PolicyName: WriteDataToS3Destination
          PolicyDocument:
            Version: 2012-10-17
            Statement:
              - Effect: Allow
                Action:
                  - s3:PutObject
                Resource: !Sub arn:${AWS::Partition}:s3:::${DestinationS3Bucket}/*
        - PolicyName: DenyData
          PolicyDocument:
            Version: 2012-10-17
            Statement:
              - Effect: Deny
                NotResource:
                  - !Sub arn:${AWS::Partition}:s3:::${DestinationS3Bucket}/*
                Action:
                  - s3:GetObject
              - Effect: Deny
                Resource: "*"
                Action:
                  - cloudformation:GetTemplate
                  - dynamodb:GetItem
                  - dynamodb:BatchGetItem
                  - dynamodb:Query
                  - dynamodb:Scan
                  - ec2:GetConsoleOutput
                  - ec2:GetConsoleScreenshot
                  - kinesis:Get*
                  - lambda:GetFunction
                  - logs:GetLogEvents
                  - sdb:Select*
                  - sqs:ReceiveMessage
      ManagedPolicyArns:
        - "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"
        - "arn:aws:iam::aws:policy/ReadOnlyAccess"
        - "arn:aws:iam::aws:policy/AWSMarketplaceMeteringFullAccess"
Outputs:
  ClusterId:
    Value: !Ref ECSCluster
  TaskArn:
    Value: !Ref ScheduledWorkerTask
```


## Step 4: Deploy the CloudFormation Template

You can deploy the CloudFormation template using the `aws cloudformation deploy` command. This command will create the required resources for the deployment of CloudQuery on AWS ECS with Fargate. If you are using the AWS Marketplace version of CloudQuery, you can set the `AWSMarketplace` parameter to `true`.

``` bash
aws cloudformation deploy --template-file cloudquery-ecs-resources.yaml --stack-name <STACK_NAME> --parameter-overrides CQApiKey=<SECRET_ARN> CQVersion=latest PrivateSubnetIds=<SUBNET_ID_1>,<SUBNET_ID_2> SecurityGroupIds=<SecurityGroup_ID_1>  DestinationS3Bucket=<DESTINATION_BUCKET> CQConfiguration=<BASE64_ENCODED_CONFIG>
```



## Step 5: Run a CloudQuery sync

To get the values for Cluster Name and Task ARN you can use the following command:

```bash
aws cloudformation describe-stacks --stack-name <STACK_NAME> --capabilities CAPABILITY_IAM --query "Stacks[].Outputs"
```

Now that the task definition is registered, it's time to run the CloudQuery task on ECS using the `aws ecs run-task` command:

```bash
aws ecs run-task \
  --cluster <REPLACE_WITH_ECS_CLUSTER_ID> \
  --task-definition <REPLACE_WITH_TASK_ARN> \
  --launch-type FARGATE \
  --network-configuration 'awsvpcConfiguration={subnets=[<REPLACE_WITH_SUBNET_1>,<REPLACE_WITH_SUBNET_2>],securityGroups=[<REPLACE_WITH_SG_1>,<REPLACE_WITH_SG_2>],assignPublicIp=ENABLED}'
```
Replace the following placeholders: 
  - `<REPLACE_WITH_ECS_CLUSTER_NAME>` with the name of the ECS cluster you created in Step 2
  - `<REPLACE_WITH_TASK_ARN>` with the ARN of the task definition you registered in Step 6
  - `<REPLACE_WITH_SUBNET_1>` and `<REPLACE_WITH_SUBNET_2>` with the IDs of the subnets in which you want to run the task. You can specify any number of subnets that you want
  - `<REPLACE_WITH_SG_1>` and `<REPLACE_WITH_SG_2>` with the IDs of the security groups for the task. You can specify any number of security groups that you want

Note: if you are deploying this in a private subnet you will need to set the `assignPublicIp` to `DISABLED`



## Conclusion

You now have a working CloudQuery deployment that runs on a regular basis and stores the results in an S3 bucket. This is a great base for iterating on. To learn more about performance you can check out the [performance tuning guide](/docs/advanced-topics/performance-tuning).

If you have any questions or comments, please feel free to reach out to us on [GitHub](https://github.com/cloudquery/cloudquery) or our [Community](https://community.cloudquery.io)!
