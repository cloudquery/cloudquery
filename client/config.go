package client

// Provider Configuration
type Account struct {
	ID      string
	RoleARN string `yaml:"role_arn"`
}

type Config struct {
	Regions    []string  `yaml:"regions"`
	Accounts   []Account `yaml:"accounts"`
	AWSDebug   bool      `yaml:"aws_debug"`
	LogLevel   *string   `yaml:"log_level"`
	MaxRetries int       `yaml:"max_retries" default:"5"`
	MaxBackoff int       `yaml:"max_backoff" default:"30"`
	Resources  []Resource
}

type Resource struct {
	Name  string
	Other map[string]interface{} `yaml:",inline"`
}

const DefaultConfigYaml = `
  - name: aws
#    accounts: # Optional. if you want to assume role to multiple account and fetch data from them
#      - role_arn: <CHANGE_THIS>
#    regions: # Optional. if commented out assumes all regions
#      - us-east-1
#      - us-west-2
#    aws_debug: false # Optional. if commented out will enable AWS SDK debug logging. 
#    max_retries: 5  # Optional. The maximum number of times that a request will be retried for failures. Defaults to 5 retry attempts.
#    max_backoff: 30 # Optional. The maximum back off delay between attempts. The backoff delays exponentially with a jitter based on the number of attempts. Defaults to 60 seconds.
    resources: # You can comment resources your are not interested in for faster fetching.
      - name: autoscaling.launch_configurations
      - name: cloudtrail.trails
      - name: cloudwatch.alarms
      - name: cloudwatchlogs.filters
      - name: directconnect.gateways
      - name: directconnect.virtual_interfaces
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
      - name: ec2.transit_gateways
      - name: ec2.vpc_peering_connections
      - name: ec2.vpcs
      - name: ecs.clusters
      - name: ecr.repositories
      - name: efs.filesystems
      - name: eks.clusters
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
      - name: rds.instances
      - name: redshift.clusters
      - name: redshift.subnet_groups
      - name: s3.buckets
      - name: sns.subscriptions
      - name: sns.topics`
