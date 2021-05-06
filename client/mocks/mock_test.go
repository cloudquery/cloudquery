package mocks_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/resources"
	"github.com/cloudquery/cq-provider-sdk/logging"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/hashicorp/go-hclog"
	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

type TestResource struct {
	resource           string
	mockBuilder        func(*testing.T, *gomock.Controller) client.Services
	mainTable          *schema.Table
	verifyEmptyColumns bool
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
			resource:           "autoscaling.launch_configurations",
			mockBuilder:        buildAutoscalingLaunchConfigurationsMock,
			mainTable:          resources.AutoscalingLaunchConfigurations(),
			verifyEmptyColumns: true,
		},
		{
			resource:           "cloudwatch.alarms",
			mockBuilder:        buildCloudWatchAlarmsMock,
			mainTable:          resources.CloudwatchAlarms(),
			verifyEmptyColumns: true,
		},
		{
			resource:           "cloudtrail.trails",
			mockBuilder:        buildCloudtrailTrailsMock,
			mainTable:          resources.CloudtrailTrails(),
			verifyEmptyColumns: true,
		},
		{
			resource:           "cloudwatchlogs.filters",
			mockBuilder:        buildCloudwatchLogsFiltersMock,
			mainTable:          resources.CloudwatchlogsFilters(),
			verifyEmptyColumns: true,
		},
		{
			resource:           "directconnect.gateways",
			mockBuilder:        buildDirectconnectGatewaysMock,
			mainTable:          resources.DirectconnectGateways(),
			verifyEmptyColumns: true,
		},
		{
			resource:           "ec2.byoip_cidrs",
			mockBuilder:        buildEc2ByoipCidrsMock,
			mainTable:          resources.Ec2ByoipCidrs(),
			verifyEmptyColumns: true,
		},
		{
			resource:           "directconnect.virtual_gateways",
			mockBuilder:        buildDirectconnectVirtualGatewaysMock,
			mainTable:          resources.DirectconnectVirtualGateways(),
			verifyEmptyColumns: true,
		},
		{
			resource:           "directconnect.virtual_interfaces",
			mockBuilder:        buildDirectconnectVirtualInterfacesMock,
			mainTable:          resources.DirectconnectVirtualInterfaces(),
			verifyEmptyColumns: true,
		},
		{
			resource:           "ec2.customer_gateways",
			mockBuilder:        buildEc2CustomerGateways,
			mainTable:          resources.Ec2CustomerGateways(),
			verifyEmptyColumns: true,
		},
		{
			resource:           "ec2.ebs_volumes",
			mockBuilder:        buildEc2EbsVolumes,
			mainTable:          resources.Ec2EbsVolumes(),
			verifyEmptyColumns: true,
		},
		{
			resource:           "ec2.flow_logs",
			mockBuilder:        buildEc2FlowLogsMock,
			mainTable:          resources.Ec2FlowLogs(),
			verifyEmptyColumns: true,
		},
		{
			resource:           "ec2.internet_gateways",
			mockBuilder:        buildEc2InternetGateways,
			mainTable:          resources.Ec2InternetGateways(),
			verifyEmptyColumns: true,
		},
		{
			resource:           "ec2.images",
			mockBuilder:        buildEc2ImagesMock,
			mainTable:          resources.Ec2Images(),
			verifyEmptyColumns: true,
		},
		{
			resource:           "ec2.nat_gateways",
			mockBuilder:        buildEc2NatGateways,
			mainTable:          resources.Ec2NatGateways(),
			verifyEmptyColumns: true,
		},
		{
			resource:           "ec2.network_acls",
			mockBuilder:        buildEc2NetworkAcls,
			mainTable:          resources.Ec2NetworkAcls(),
			verifyEmptyColumns: true,
		},
		{
			resource:           "ec2.route_tables",
			mockBuilder:        buildEc2RouteTables,
			mainTable:          resources.Ec2RouteTables(),
			verifyEmptyColumns: true,
		},
		{
			resource:           "ec2.instances",
			mockBuilder:        buildEc2Instances,
			mainTable:          resources.Ec2Instances(),
			verifyEmptyColumns: true,
		},
		{
			resource:           "ec2.security_groups",
			mockBuilder:        buildEc2SecurityGroups,
			mainTable:          resources.Ec2SecurityGroups(),
			verifyEmptyColumns: true,
		},
		{
			resource:           "ec2.subnets",
			mockBuilder:        buildEc2Subnets,
			mainTable:          resources.Ec2Subnets(),
			verifyEmptyColumns: true,
		},
		{
			resource:           "ec2.transit_gateways",
			mockBuilder:        buildEc2TransitGateways,
			mainTable:          resources.Ec2TransitGateways(),
			verifyEmptyColumns: true,
		},
		{
			resource:           "ec2.vpcs",
			mockBuilder:        buildEc2Vpcs,
			mainTable:          resources.Ec2Vpcs(),
			verifyEmptyColumns: true,
		},
		{
			resource:           "ec2.vpc_peering_connections",
			mockBuilder:        buildEc2VpcsPeeringConnections,
			mainTable:          resources.Ec2VpcPeeringConnections(),
			verifyEmptyColumns: true,
		},
		{
			resource:           "ecr.repositories",
			mockBuilder:        buildEcrRepositoriesMock,
			mainTable:          resources.EcrRepositories(),
			verifyEmptyColumns: true,
		},
		{
			resource:           "efs.filesystems",
			mockBuilder:        buildEfsFilesystemsMock,
			mainTable:          resources.EfsFilesystems(),
			verifyEmptyColumns: true,
		},
		{
			resource:           "rds.certificates",
			mockBuilder:        buildRdsCertificates,
			mainTable:          resources.RdsCertificates(),
			verifyEmptyColumns: true,
		},
		{
			resource:           "rds.clusters",
			mockBuilder:        buildRdsDBClusters,
			mainTable:          resources.RdsClusters(),
			verifyEmptyColumns: true,
		},
		{
			resource:           "rds.instances",
			mockBuilder:        buildRdsDBInstances,
			mainTable:          resources.RdsInstances(),
			verifyEmptyColumns: true,
		},
		{
			resource:           "rds.db_subnet_groups",
			mockBuilder:        buildRdsDBSubnetGroups,
			mainTable:          resources.RdsSubnetGroups(),
			verifyEmptyColumns: true,
		},
		{
			resource:           "elasticbeanstalk.environments",
			mockBuilder:        buildElasticbeanstalkEnvironments,
			mainTable:          resources.ElasticbeanstalkEnvironments(),
			verifyEmptyColumns: true,
		},
		{
			resource:           "elbv2.load_balancers",
			mockBuilder:        buildElbv2LoadBalancers,
			mainTable:          resources.Elbv2LoadBalancers(),
			verifyEmptyColumns: true,
		},
		{
			resource:           "elbv2.target_groups",
			mockBuilder:        buildElbv2TargetGroups,
			mainTable:          resources.Elbv2TargetGroups(),
			verifyEmptyColumns: true,
		},
		{
			resource:           "emr.clusters",
			mockBuilder:        buildEmrClusters,
			mainTable:          resources.EmrClusters(),
			verifyEmptyColumns: true,
		},
		{
			resource:           "eks.clusters",
			mockBuilder:        buildEksClusters,
			mainTable:          resources.EksClusters(),
			verifyEmptyColumns: true,
		},
		// Infinite loop in faker
		//{
		//	resource: "fsx.backups",
		//	mockFunc: testFsxBackups,
		//	tables: []string{"aws_fsx_backups"},
		//},
		{
			resource:           "sns.topics",
			mockBuilder:        buildSnsTopics,
			mainTable:          resources.SnsTopics(),
			verifyEmptyColumns: true,
		},
		{
			resource:           "sns.subscriptions",
			mockBuilder:        buildSnsSubscriptions,
			mainTable:          resources.SnsSubscriptions(),
			verifyEmptyColumns: true,
		},
		{
			resource:           "iam.groups",
			mainTable:          resources.IamGroups(),
			mockBuilder:        buildIamGroups,
			verifyEmptyColumns: true,
		},
		{
			resource:           "iam.policies",
			mainTable:          resources.IamPolicies(),
			mockBuilder:        buildIamPolicies,
			verifyEmptyColumns: true,
		},
		{
			resource:           "iam.password_policies",
			mainTable:          resources.IamPasswordPolicies(),
			mockBuilder:        buildIamPasswordPolicies,
			verifyEmptyColumns: true,
		},
		{
			resource:           "iam.roles",
			mainTable:          resources.IamRoles(),
			mockBuilder:        buildIamRoles,
			verifyEmptyColumns: true,
		},
		{
			resource:           "iam.users",
			mainTable:          resources.IamUsers(),
			mockBuilder:        buildIamUsers,
			verifyEmptyColumns: true,
		},
		{
			resource:           "iam.virtual_mfa_devices",
			mainTable:          resources.IamVirtualMfaDevices(),
			mockBuilder:        buildIamVirtualMfaDevices,
			verifyEmptyColumns: true,
		},
		{
			resource:           "kms.keys",
			mainTable:          resources.KmsKeys(),
			mockBuilder:        buildKmsKeys,
			verifyEmptyColumns: true,
		},
		{
			resource:           "organizations.accounts",
			mainTable:          resources.OrganizationsAccounts(),
			mockBuilder:        buildOrganizationsAccounts,
			verifyEmptyColumns: true,
		},
		{
			resource:           "s3.buckets",
			mainTable:          resources.S3Buckets(),
			mockBuilder:        buildS3Buckets,
			verifyEmptyColumns: true,
		},
		{
			resource:           "redshift.clusters",
			mainTable:          resources.RedshiftClusters(),
			mockBuilder:        buildRedshiftClustersMock,
			verifyEmptyColumns: true,
		},
		{
			resource:           "redshift.subnet_groups",
			mainTable:          resources.RedshiftSubnetGroups(),
			mockBuilder:        buildRedshiftSubnetGroupsMock,
			verifyEmptyColumns: true,
		},
		{
			resource:           "ecs.clusters",
			mainTable:          resources.EcsClusters(),
			mockBuilder:        buildEcsClusterMock,
			verifyEmptyColumns: true,
		},
		{
			resource:           "cloudfront.distributions",
			mainTable:          resources.CloudfrontDistributions(),
			mockBuilder:        buildCloudfrontDistributionsMock,
			verifyEmptyColumns: true,
		},
		{
			resource:           "cloudfront.cache_policies",
			mainTable:          resources.CloudfrontCachePolicies(),
			mockBuilder:        buildCloudfrontCachePoliciesMock,
			verifyEmptyColumns: true,
		},
		{
			resource:           "ec2.vpc_endpoints",
			mainTable:          resources.Ec2VpcEndpoints(),
			mockBuilder:        buildEc2VpcEndpoints,
			verifyEmptyColumns: true,
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
				Resources: []client.Resource{{
					Name: tc.resource,
				},
				},
			}
			testProvider := resources.Provider()
			testProvider.Logger = logging.New(hclog.DefaultOptions)
			testProvider.Configure = func(logger hclog.Logger, i interface{}) (schema.ClientMeta, error) {
				client := client.NewAwsClient(logging.New(&hclog.LoggerOptions{
					Level: hclog.Warn,
				}), []string{"test-1"})
				client.SetAccountServices("testAccount", tc.mockBuilder(t, ctrl))
				return &client, nil
			}
			err := testProvider.Init("", "host=localhost user=postgres password=pass DB.name=postgres port=5432", false)
			assert.Nil(t, err)
			data, err := yaml.Marshal(cfg)
			assert.Nil(t, err)
			err = testProvider.Fetch(data)
			assert.Nil(t, err)
			if tc.verifyEmptyColumns {
				verifyNoEmptyColumns(t, tc, conn)
			}
		})
	}
}

func verifyNoEmptyColumns(t *testing.T, tc TestResource, conn *pgx.Conn) {
	// Test that we don't have missing columns and have exactly one entry for each table
	for _, table := range getTablesFromMainTable(tc.mainTable) {

		query := fmt.Sprintf("select * FROM %s ", table)
		rows, err := conn.Query(context.Background(), query)
		if err != nil {
			t.Fatal(err)
		}
		count := 0
		for rows.Next() {
			count += 1
		}
		if count < 1 {
			t.Fatalf("expected to have at least 1 entry at table %s got %d", table, count)
		}

		query = fmt.Sprintf("select t.* FROM %s as t WHERE to_jsonb(t) = jsonb_strip_nulls(to_jsonb(t))", table)
		rows, err = conn.Query(context.Background(), query)
		if err != nil {
			t.Fatal(err)
		}
		count = 0
		for rows.Next() {
			count += 1
		}
		if count < 1 {
			t.Fatalf("row at table %s has an empty column", table)
		}
	}
}

func getTablesFromMainTable(table *schema.Table) []string {
	var res []string
	res = append(res, table.Name)
	for _, t := range table.Relations {
		res = append(res, getTablesFromMainTable(t)...)
	}
	return res
}
