:::callout{type="info"}
This feature is currently in a closed preview. Sign up for access using this [form](https://cloudquery.typeform.com/to/Dvdn6P2u)
:::

# Event-Based Sync

AWS CloudTrail enables users to get an audit log of events occurring within their account. By subscribing to a stream of AWS CloudTrail events in a Kinesis Data stream CloudQuery can trigger selective syncs to update just the singular resource that had a configuration change. 

Each table in the supported list is a top level table. When an event is received for a table, all child tables are re-synced too by default. To skip some child tables you can use `skip_tables`

## Supported Services and Events
Service | Event | CloudQuery Tables
-|-|-|
ec2.amazonaws.com | AssociateRouteTable | `aws_ec2_route_tables`
ec2.amazonaws.com | AttachInternetGateway | `aws_ec2_internet_gateways`
ec2.amazonaws.com | AuthorizeSecurityGroupEgress | `aws_ec2_security_groups`
ec2.amazonaws.com | AuthorizeSecurityGroupIngress | `aws_ec2_security_groups`
ec2.amazonaws.com | CreateImage | `aws_ec2_images`
ec2.amazonaws.com | CreateInternetGateway | `aws_ec2_internet_gateways`
ec2.amazonaws.com | CreateNetworkInterface | `aws_ec2_network_interfaces`
ec2.amazonaws.com | CreateSecurityGroup | `aws_ec2_security_groups`
ec2.amazonaws.com | CreateSubnet | `aws_ec2_subnets`
ec2.amazonaws.com | CreateTags | `aws_ec2_instances`
ec2.amazonaws.com | CreateVpc | `aws_ec2_vpcs`
ec2.amazonaws.com | DeleteTags | `aws_ec2_instances`
ec2.amazonaws.com | DetachInternetGateway | `aws_ec2_internet_gateways`
ec2.amazonaws.com | ModifySubnetAttribute | `aws_ec2_subnets`
ec2.amazonaws.com | RevokeSecurityGroupEgress | `aws_ec2_security_groups`
ec2.amazonaws.com | RevokeSecurityGroupIngress | `aws_ec2_security_groups`
ec2.amazonaws.com | RunInstances | `aws_ec2_instances`
iam.amazonaws.com | CreateGroup | `aws_iam_groups`
iam.amazonaws.com | CreateGroup | `aws_iam_groups`
iam.amazonaws.com | DeleteGroup | `aws_iam_groups`
iam.amazonaws.com | UpdateGroup | `aws_iam_groups`
iam.amazonaws.com | CreateRole | `aws_iam_roles`
iam.amazonaws.com | DeleteRole | `aws_iam_roles`
iam.amazonaws.com | TagRole | `aws_iam_roles`
iam.amazonaws.com | UntagRole | `aws_iam_roles`
iam.amazonaws.com | UpdateRole | `aws_iam_roles`
iam.amazonaws.com | UpdateRoleDescription | `aws_iam_roles`
iam.amazonaws.com | CreateUser | `aws_iam_users`
iam.amazonaws.com | DeleteUser | `aws_iam_users`
iam.amazonaws.com | TagUser | `aws_iam_users`
iam.amazonaws.com | UntagUser | `aws_iam_users`
iam.amazonaws.com | UpdateUser | `aws_iam_users`
rds.amazonaws.com | CreateDBCluster | `aws_rds_clusters`
rds.amazonaws.com | CreateDBInstance | `aws_rds_instances`
rds.amazonaws.com | ModifyDBCluster | `aws_rds_clusters`
rds.amazonaws.com | ModifyDBInstance | `aws_rds_instances`


## Configuration

1. Configure an AWS CloudTrail Trail to send management events to a Kinesis Data Stream via CloudWatch Logs. The most straight-forward way to do this is to use the CloudFormation template provided by CloudQuery.

The CloudFormation template will deploy the following architecture:

![Event based syncing cloud infrastructure](/images/docs/aws/event-based-sync-architecture.png)


```bash
aws cloudformation deploy --template-file ./streaming-deployment.yml --stack-name <STACK-NAME> --capabilities CAPABILITY_IAM --disable-rollback --region <DESIRED-REGION>
```



2. Copy the ARN of the Kinesis stream. If you used the CloudFormation template you can run the following command:
```bash
aws cloudformation describe-stacks --stack-name <STACK-NAME> --query "Stacks[].Outputs" --region <DESIRED-REGION>
```

3. Define a `config.yml` file like the one below

``` yaml
kind: source
spec:
  name: "aws-event-based"
  registry: "local"
  path: <PATH/TO/BINARY>
  tables:
    - aws_ec2_instances
    - aws_ec2_internet_gateways
    - aws_ec2_security_groups
    - aws_ec2_subnets
    - aws_ec2_vpcs
    - aws_ecs_cluster_tasks
    - aws_iam_groups
    - aws_iam_roles
    - aws_iam_users
    - aws_rds_instances
  destinations: ["DESTINATION_NAME"]
  spec:
    event_based_sync:
      # account:
      #  local_profile: "<ROLE-NAME>"
      kinesis_stream_arn: <OUTPUT-FROM-CLOUDFORMATION-STACK>
```

4. Sync the data! 
```bash
cloudquery sync config.yml
```

This will start a long lived process that will only stop when there is an error or you stop the process


### Limitations
- Kinesis Stream can only have a single shard. (This is a limitation that we expect to remove in the future)
- Stale records will only be deleted if the plugin stops consuming the Kinesis Stream, which only can occur if there is an error