package mocks_test

import (
	"context"
	"os"
	"testing"

	"github.com/cloudquery/cq-provider-sdk/provider/providertest"

	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/resources"
	"github.com/cloudquery/cq-provider-sdk/logging"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/hashicorp/go-hclog"
	"github.com/jackc/pgx/v4"
)

type TestResource struct {
	resource    string
	mockBuilder func(*testing.T, *gomock.Controller) client.Services
	mainTable   *schema.Table
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func TestResources(t *testing.T) {
	dbCfg, err := pgx.ParseConfig(getEnv("DATABASE_URL",
		"host=localhost user=postgres password=pass DB.name=postgres port=5432"))
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	conn, err := pgx.ConnectConfig(ctx, dbCfg)
	if err != nil {
		t.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(ctx)
	_ = faker.SetRandomMapAndSliceMinSize(1)
	err = faker.SetRandomMapAndSliceMaxSize(1)
	if err != nil {
		t.Fatal(err)
	}
	ctrl := gomock.NewController(t)
	testResourcesTable := []TestResource{
		{
			resource:    "autoscaling.launch_configurations",
			mockBuilder: buildAutoscalingLaunchConfigurationsMock,
			mainTable:   resources.AutoscalingLaunchConfigurations(),
		},
		{
			resource:    "cloudwatch.alarms",
			mockBuilder: buildCloudWatchAlarmsMock,
			mainTable:   resources.CloudwatchAlarms(),
		},
		{
			resource:    "cloudtrail.trails",
			mockBuilder: buildCloudtrailTrailsMock,
			mainTable:   resources.CloudtrailTrails(),
		},
		{
			resource:    "cloudwatchlogs.filters",
			mockBuilder: buildCloudwatchLogsFiltersMock,
			mainTable:   resources.CloudwatchlogsFilters(),
		},
		{
			resource:    "directconnect.gateways",
			mockBuilder: buildDirectconnectGatewaysMock,
			mainTable:   resources.DirectconnectGateways(),
		},
		{
			resource:    "ec2.byoip_cidrs",
			mockBuilder: buildEc2ByoipCidrsMock,
			mainTable:   resources.Ec2ByoipCidrs(),
		},
		{
			resource:    "directconnect.virtual_gateways",
			mockBuilder: buildDirectconnectVirtualGatewaysMock,
			mainTable:   resources.DirectconnectVirtualGateways(),
		},
		{
			resource:    "directconnect.virtual_interfaces",
			mockBuilder: buildDirectconnectVirtualInterfacesMock,
			mainTable:   resources.DirectconnectVirtualInterfaces(),
		},
		{
			resource:    "ec2.customer_gateways",
			mockBuilder: buildEc2CustomerGateways,
			mainTable:   resources.Ec2CustomerGateways(),
		},
		{
			resource:    "ec2.ebs_volumes",
			mockBuilder: buildEc2EbsVolumes,
			mainTable:   resources.Ec2EbsVolumes(),
		},
		{
			resource:    "ec2.flow_logs",
			mockBuilder: buildEc2FlowLogsMock,
			mainTable:   resources.Ec2FlowLogs(),
		},
		{
			resource:    "ec2.internet_gateways",
			mockBuilder: buildEc2InternetGateways,
			mainTable:   resources.Ec2InternetGateways(),
		},
		{
			resource:    "ec2.images",
			mockBuilder: buildEc2ImagesMock,
			mainTable:   resources.Ec2Images(),
		},
		{
			resource:    "ec2.nat_gateways",
			mockBuilder: buildEc2NatGateways,
			mainTable:   resources.Ec2NatGateways(),
		},
		{
			resource:    "ec2.network_acls",
			mockBuilder: buildEc2NetworkAcls,
			mainTable:   resources.Ec2NetworkAcls(),
		},
		{
			resource:    "ec2.route_tables",
			mockBuilder: buildEc2RouteTables,
			mainTable:   resources.Ec2RouteTables(),
		},
		{
			resource:    "ec2.instances",
			mockBuilder: buildEc2Instances,
			mainTable:   resources.Ec2Instances(),
		},
		{
			resource:    "ec2.security_groups",
			mockBuilder: buildEc2SecurityGroups,
			mainTable:   resources.Ec2SecurityGroups(),
		},
		{
			resource:    "ec2.subnets",
			mockBuilder: buildEc2Subnets,
			mainTable:   resources.Ec2Subnets(),
		},
		{
			resource:    "ec2.transit_gateways",
			mockBuilder: buildEc2TransitGateways,
			mainTable:   resources.Ec2TransitGateways(),
		},
		{
			resource:    "ec2.vpcs",
			mockBuilder: buildEc2Vpcs,
			mainTable:   resources.Ec2Vpcs(),
		},
		{
			resource:    "ec2.vpc_peering_connections",
			mockBuilder: buildEc2VpcsPeeringConnections,
			mainTable:   resources.Ec2VpcPeeringConnections(),
		},
		{
			resource:    "ecr.repositories",
			mockBuilder: buildEcrRepositoriesMock,
			mainTable:   resources.EcrRepositories(),
		},
		{
			resource:    "efs.filesystems",
			mockBuilder: buildEfsFilesystemsMock,
			mainTable:   resources.EfsFilesystems(),
		},
		{
			resource:    "rds.certificates",
			mockBuilder: buildRdsCertificates,
			mainTable:   resources.RdsCertificates(),
		},
		{
			resource:    "rds.clusters",
			mockBuilder: buildRdsDBClusters,
			mainTable:   resources.RdsClusters(),
		},
		{
			resource:    "rds.instances",
			mockBuilder: buildRdsDBInstances,
			mainTable:   resources.RdsInstances(),
		},
		{
			resource:    "rds.db_subnet_groups",
			mockBuilder: buildRdsDBSubnetGroups,
			mainTable:   resources.RdsSubnetGroups(),
		},
		{
			resource:    "elasticbeanstalk.environments",
			mockBuilder: buildElasticbeanstalkEnvironments,
			mainTable:   resources.ElasticbeanstalkEnvironments(),
		},
		{
			resource:    "elbv1.load_balancers",
			mockBuilder: buildElbv1LoadBalancers,
			mainTable:   resources.Elbv1LoadBalancers(),
		},
		{
			resource:    "elbv2.load_balancers",
			mockBuilder: buildElbv2LoadBalancers,
			mainTable:   resources.Elbv2LoadBalancers(),
		},
		{
			resource:    "elbv2.target_groups",
			mockBuilder: buildElbv2TargetGroups,
			mainTable:   resources.Elbv2TargetGroups(),
		},
		{
			resource:    "emr.clusters",
			mockBuilder: buildEmrClusters,
			mainTable:   resources.EmrClusters(),
		},
		{
			resource:    "eks.clusters",
			mockBuilder: buildEksClusters,
			mainTable:   resources.EksClusters(),
		},
		// Infinite loop in faker
		//{
		//	resource: "fsx.backups",
		//	mockFunc: testFsxBackups,
		//	tables: []string{"aws_fsx_backups"},
		//},
		{
			resource:    "sns.topics",
			mockBuilder: buildSnsTopics,
			mainTable:   resources.SnsTopics(),
		},
		{
			resource:    "sns.subscriptions",
			mockBuilder: buildSnsSubscriptions,
			mainTable:   resources.SnsSubscriptions(),
		},
		{
			resource:    "iam.groups",
			mainTable:   resources.IamGroups(),
			mockBuilder: buildIamGroups,
		},
		{
			resource:    "iam.policies",
			mainTable:   resources.IamPolicies(),
			mockBuilder: buildIamPolicies,
		},
		{
			resource:    "iam.password_policies",
			mainTable:   resources.IamPasswordPolicies(),
			mockBuilder: buildIamPasswordPolicies,
		},
		{
			resource:    "iam.roles",
			mainTable:   resources.IamRoles(),
			mockBuilder: buildIamRoles,
		},
		{
			resource:    "iam.virtual_mfa_devices",
			mainTable:   resources.IamVirtualMfaDevices(),
			mockBuilder: buildIamVirtualMfaDevices,
		},
		{
			resource:    "iam.openid_connect_identity_providers",
			mainTable:   resources.IamOpenidConnectIdentityProviders(),
			mockBuilder: buildIamOpenIDConnectProviders,
		},
		{
			resource:    "iam.saml_identity_providers",
			mainTable:   resources.IamSamlIdentityProviders(),
			mockBuilder: buildIamSAMLProviders,
		},
		{
			resource:    "kms.keys",
			mainTable:   resources.KmsKeys(),
			mockBuilder: buildKmsKeys,
		},
		{
			resource:    "organizations.accounts",
			mainTable:   resources.OrganizationsAccounts(),
			mockBuilder: buildOrganizationsAccounts,
		},
		{
			resource:    "redshift.clusters",
			mainTable:   resources.RedshiftClusters(),
			mockBuilder: buildRedshiftClustersMock,
		},
		{
			resource:    "redshift.subnet_groups",
			mainTable:   resources.RedshiftSubnetGroups(),
			mockBuilder: buildRedshiftSubnetGroupsMock,
		},
		{
			resource:    "ecs.clusters",
			mainTable:   resources.EcsClusters(),
			mockBuilder: buildEcsClusterMock,
		},
		{
			resource:    "cloudfront.distributions",
			mainTable:   resources.CloudfrontDistributions(),
			mockBuilder: buildCloudfrontDistributionsMock,
		},
		{
			resource:    "cloudfront.cache_policies",
			mainTable:   resources.CloudfrontCachePolicies(),
			mockBuilder: buildCloudfrontCachePoliciesMock,
		},
		{
			resource:    "ec2.vpc_endpoints",
			mainTable:   resources.Ec2VpcEndpoints(),
			mockBuilder: buildEc2VpcEndpoints,
		},
		{
			resource:    "route53.traffic_policies",
			mainTable:   resources.Route53TrafficPolicies(),
			mockBuilder: buildRoute53TrafficPoliciesMock,
		},
		{
			resource:    "route53.health_checks",
			mainTable:   resources.Route53HealthChecks(),
			mockBuilder: buildRoute53HealthChecksMock,
		},
		{
			resource:    "route53.reusable_delegation_sets",
			mainTable:   resources.Route53ReusableDelegationSets(),
			mockBuilder: buildRoute53DelegationSetsMock,
		},
	}
	for _, tc := range testResourcesTable {
		t.Run(tc.resource, func(t *testing.T) {
			cfg := client.Config{
				Regions:    []string{"us-east-1"},
				Accounts:   []client.Account{{ID: "testAccount", RoleARN: ""}},
				AWSDebug:   false,
				MaxRetries: 3,
				MaxBackoff: 60,
			}
			providertest.TestResource(t, resources.Provider, providertest.ResourceTestData{
				Table:  tc.mainTable,
				Config: cfg,
				Configure: func(logger hclog.Logger, i interface{}) (schema.ClientMeta, error) {
					c := client.NewAwsClient(logging.New(&hclog.LoggerOptions{
						Level: hclog.Warn,
					}), []string{"us-east-1"})
					c.ServicesManager.InitServicesForAccountAndRegion("testAccount", "us-east-1", tc.mockBuilder(t, ctrl))
					return &c, nil
				},
			})

		})
	}
}
