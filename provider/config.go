package provider

const configYaml = `
  - name: aws
#    accounts: # Optional. if you want to assume role to multiple account and fetch data from them
#      - role_arn: <CHANGE_THIS>
#    regions: # Optional. if commented out assumes all regions
#      - us-east-1
#      - us-west-2
#    log_level: debug # Optional. if commented out will enable AWS SDK debug logging. possible values: debug, debug_with_signing, debug_with_http_body, debug_with_request_retries, debug_with_request_error, debug_with_event_stream_body
#    max_retries: # Optional. The maximum number of times that a request will be retried for failures. Defaults to -1, which defers the max retry setting to the service specific configuration. 	
    resources: # You can comment resources your are not interested in for faster fetching.
      - name: autoscaling.launch_configurations
      - name: cloudtrail.trails
      - name: cloudwatch.alarms
      - name: cloudwatchlogs.metric_filters
      - name: directconnect.gateways
      - name: ec2.customer_gateways
      - name: ec2.flow_logs
      - name: ec2.images
      - name: ec2.instances
      - name: ec2.internet_gateways
      - name: ec2.nat_gateways
      - name: ec2.network_acls
      - name: ec2.route_tables
      - name: ec2.security_groups
      - name: ec2.subnets
      - name: ec2.vpc_peering_connections
      - name: ec2.vpcs
      - name: ecs.clusters
      - name: ecr.images
      - name: efs.filesystems
      - name: elasticbeanstalk.environments
      - name: elbv2.load_balancers
      - name: elbv2.target_groups
      - name: emr.clusters
      - name: fsx.backups
      - name: iam.groups
      - name: iam.password_policies
      - name: iam.policies
      - name: iam.roles
      - name: iam.users
      - name: iam.virtual_mfa_devices
      - name: kms.keys
      - name: organizations.accounts
      - name: rds.certificates
      - name: rds.clusters
      - name: rds.db_subnet_groups
      - name: redshift.clusters
      - name: redshift.cluster_subnet_groups
      - name: s3.buckets
      - name: sns.subscriptions
      - name: sns.topics`
