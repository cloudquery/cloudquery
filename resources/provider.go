package resources

import (
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func Provider() *provider.Provider {
	return &provider.Provider{
		Name:      "aws",
		Configure: client.Configure,
		ResourceMap: map[string]*schema.Table{
			"autoscaling.launch_configurations":     AutoscalingLaunchConfigurations(),
			"cloudfront.distributions":              CloudfrontDistributions(),
			"cloudfront.cache_policies":             CloudfrontCachePolicies(),
			"cloudtrail.trails":                     CloudtrailTrails(),
			"cloudwatch.alarms":                     CloudwatchAlarms(),
			"cloudwatchlogs.filters":                CloudwatchlogsFilters(),
			"s3.buckets":                            S3Buckets(),
			"directconnect.gateways":                DirectconnectGateways(),
			"directconnect.virtual_gateways":        DirectconnectVirtualGateways(),
			"directconnect.virtual_interfaces":      DirectconnectVirtualInterfaces(),
			"ec2.byoip_cidrs":                       Ec2ByoipCidrs(),
			"ec2.customer_gateways":                 Ec2CustomerGateways(),
			"ec2.flow_logs":                         Ec2FlowLogs(),
			"ec2.images":                            Ec2Images(),
			"ec2.internet_gateways":                 Ec2InternetGateways(),
			"ec2.nat_gateways":                      Ec2NatGateways(),
			"ec2.network_acls":                      Ec2NetworkAcls(),
			"ec2.route_tables":                      Ec2RouteTables(),
			"ec2.subnets":                           Ec2Subnets(),
			"ec2.transit_gateways":                  Ec2TransitGateways(),
			"ec2.vpc_peering_connections":           Ec2VpcPeeringConnections(),
			"ec2.vpc_endpoints":                     Ec2VpcEndpoints(),
			"ec2.vpcs":                              Ec2Vpcs(),
			"ec2.instances":                         Ec2Instances(),
			"ec2.security_groups":                   Ec2SecurityGroups(),
			"ec2.ebs_volumes":                       Ec2EbsVolumes(),
			"ecr.repositories":                      EcrRepositories(),
			"efs.filesystems":                       EfsFilesystems(),
			"eks.clusters":                          EksClusters(),
			"ecs.clusters":                          EcsClusters(),
			"elasticbeanstalk.environments":         ElasticbeanstalkEnvironments(),
			"elbv2.target_groups":                   Elbv2TargetGroups(),
			"elbv2.load_balancers":                  Elbv2LoadBalancers(),
			"emr.clusters":                          EmrClusters(),
			"fsx.backups":                           FsxBackups(),
			"iam.groups":                            IamGroups(),
			"iam.policies":                          IamPolicies(),
			"iam.password_policies":                 IamPasswordPolicies(),
			"iam.roles":                             IamRoles(),
			"iam.users":                             IamUsers(),
			"iam.virtual_mfa_devices":               IamVirtualMfaDevices(),
			"iam.openid_connect_identity_providers": IamOpenidConnectIdentityProviders(),
			"iam.saml_identity_providers":           IamSamlIdentityProviders(),
			"kms.keys":                              KmsKeys(),
			"organizations.accounts":                OrganizationsAccounts(),
			"sns.topics":                            SnsTopics(),
			"sns.subscriptions":                     SnsSubscriptions(),
			"rds.certificates":                      RdsCertificates(),
			"rds.clusters":                          RdsClusters(),
			"rds.db_subnet_groups":                  RdsSubnetGroups(),
			"rds.instances":                         RdsInstances(),
			"redshift.clusters":                     RedshiftClusters(),
			"redshift.subnet_groups":                RedshiftSubnetGroups(),
		},
		Config: func() provider.Config {
			return &client.Config{}
		},
	}
}
