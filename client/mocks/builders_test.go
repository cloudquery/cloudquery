package mocks_test

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	autoscalingTypes "github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
	cloudtrailTypes "github.com/aws/aws-sdk-go-v2/service/cloudtrail/types"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	cloudwatchTypes "github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	cloudwatchlogsTypes "github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	"github.com/aws/aws-sdk-go-v2/service/directconnect"
	directconnectTypes "github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	ec2Types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	ecrTypes "github.com/aws/aws-sdk-go-v2/service/ecr/types"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	ecsTypes "github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/aws/aws-sdk-go-v2/service/efs"
	efsTypes "github.com/aws/aws-sdk-go-v2/service/efs/types"
	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
	elasticbeanstalkTypes "github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/types"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	elbv2Types "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	"github.com/aws/aws-sdk-go-v2/service/emr"
	emrTypes "github.com/aws/aws-sdk-go-v2/service/emr/types"
	"github.com/aws/aws-sdk-go-v2/service/fsx"
	fsxTypes "github.com/aws/aws-sdk-go-v2/service/fsx/types"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	iamTypes "github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	kmsTypes "github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/aws/aws-sdk-go-v2/service/organizations"
	organizationsTypes "github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	rdsTypes "github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
	redshiftTypes "github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	s3Types "github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	snsTypes "github.com/aws/aws-sdk-go-v2/service/sns/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildAutoscalingLaunchConfigurationsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockAutoscalingClient(ctrl)
	services := client.Services{
		Autoscaling: m,
	}
	l := autoscalingTypes.LaunchConfiguration{}
	err := faker.FakeData(&l)
	if err != nil {
		t.Fatal(err)
	}
	autoscalingLaunchConfigurations := &autoscaling.DescribeLaunchConfigurationsOutput{
		LaunchConfigurations: []autoscalingTypes.LaunchConfiguration{l},
	}
	m.EXPECT().DescribeLaunchConfigurations(gomock.Any(), gomock.Any(), gomock.Any()).Return(autoscalingLaunchConfigurations, nil)
	return services
}

func buildEcsClusterMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEcsClient(ctrl)
	services := client.Services{
		ECS: m,
	}
	c := ecsTypes.Cluster{}
	err := faker.FakeData(&c)
	if err != nil {
		t.Fatal(err)
	}
	ecsOutput := &ecs.DescribeClustersOutput{
		Clusters: []ecsTypes.Cluster{c},
	}
	m.EXPECT().DescribeClusters(gomock.Any(), gomock.Any(), gomock.Any()).Return(ecsOutput, nil)
	ecsListOutput := &ecs.ListClustersOutput{
		ClusterArns: []string{"randomClusteArn"},
	}
	m.EXPECT().ListClusters(gomock.Any(), gomock.Any(), gomock.Any()).Return(ecsListOutput, nil)
	return services
}

func buildCloudtrailTrailsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockCloudtrailClient(ctrl)
	services := client.Services{
		Cloudtrail: m,
	}
	trail := cloudtrailTypes.Trail{}
	err := faker.FakeData(&trail)
	if err != nil {
		t.Fatal(err)
	}
	trailStatus := cloudtrail.GetTrailStatusOutput{}
	err = faker.FakeData(&trailStatus)
	if err != nil {
		t.Fatal(err)
	}
	eventSelector := cloudtrailTypes.EventSelector{}
	err = faker.FakeData(&eventSelector)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeTrails(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&cloudtrail.DescribeTrailsOutput{
			TrailList: []cloudtrailTypes.Trail{trail},
		},
		nil,
	)
	m.EXPECT().GetTrailStatus(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&trailStatus,
		nil,
	)
	m.EXPECT().GetEventSelectors(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&cloudtrail.GetEventSelectorsOutput{
			EventSelectors: []cloudtrailTypes.EventSelector{eventSelector},
		},
		nil,
	)
	return services
}

func buildEc2ImagesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)
	services := client.Services{
		EC2: m,
	}
	g := ec2Types.Image{}
	err := faker.FakeData(&g)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeImages(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeImagesOutput{
			Images: []ec2Types.Image{g},
		}, nil)
	return services
}

func buildCloudWatchAlarmsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockCloudwatchClient(ctrl)
	services := client.Services{
		Cloudwatch: m,
	}
	a := cloudwatchTypes.MetricAlarm{}
	err := faker.FakeData(&a)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeAlarms(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&cloudwatch.DescribeAlarmsOutput{
			MetricAlarms: []cloudwatchTypes.MetricAlarm{a},
		}, nil)
	return services
}

func buildEc2FlowLogsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)

	g := ec2Types.FlowLog{}
	err := faker.FakeData(&g)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeFlowLogs(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeFlowLogsOutput{
			FlowLogs: []ec2Types.FlowLog{g},
		}, nil)
	return client.Services{
		EC2: m,
	}
}

func buildRedshiftClustersMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockRedshiftClient(ctrl)
	g := redshiftTypes.Cluster{}
	err := faker.FakeData(&g)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeClusters(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&redshift.DescribeClustersOutput{
			Clusters: []redshiftTypes.Cluster{g},
		}, nil)
	return client.Services{
		Redshift: m,
	}
}

func buildRedshiftSubnetGroupsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockRedshiftClient(ctrl)

	g := redshiftTypes.ClusterSubnetGroup{}
	err := faker.FakeData(&g)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeClusterSubnetGroups(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&redshift.DescribeClusterSubnetGroupsOutput{
			ClusterSubnetGroups: []redshiftTypes.ClusterSubnetGroup{g},
		}, nil)
	return client.Services{
		Redshift: m,
	}
}

func buildCloudwatchLogsFiltersMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockCloudwatchLogsClient(ctrl)
	l := cloudwatchlogsTypes.MetricFilter{}
	err := faker.FakeData(&l)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeMetricFilters(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&cloudwatchlogs.DescribeMetricFiltersOutput{
			MetricFilters: []cloudwatchlogsTypes.MetricFilter{l},
		}, nil)
	return client.Services{
		CloudwatchLogs: m,
	}
}

func buildDirectconnectGatewaysMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockDirectconnectClient(ctrl)
	l := directconnectTypes.DirectConnectGateway{}
	err := faker.FakeData(&l)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeDirectConnectGateways(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&directconnect.DescribeDirectConnectGatewaysOutput{
			DirectConnectGateways: []directconnectTypes.DirectConnectGateway{l},
		}, nil)
	return client.Services{
		Directconnect: m,
	}
}

func buildEc2ByoipCidrsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)
	l := ec2Types.ByoipCidr{}
	err := faker.FakeData(&l)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeByoipCidrs(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeByoipCidrsOutput{
			ByoipCidrs: []ec2Types.ByoipCidr{l},
		}, nil)
	return client.Services{
		EC2: m,
	}
}

func buildEc2CustomerGateways(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)
	l := ec2Types.CustomerGateway{}
	err := faker.FakeData(&l)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeCustomerGateways(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeCustomerGatewaysOutput{
			CustomerGateways: []ec2Types.CustomerGateway{l},
		}, nil)
	return client.Services{
		EC2: m,
	}
}

func buildEc2InternetGateways(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)
	l := ec2Types.InternetGateway{}
	err := faker.FakeData(&l)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeInternetGateways(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeInternetGatewaysOutput{
			InternetGateways: []ec2Types.InternetGateway{l},
		}, nil)
	return client.Services{
		EC2: m,
	}
}

func buildEc2NatGateways(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)
	l := ec2Types.NatGateway{}
	err := faker.FakeData(&l)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeNatGateways(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeNatGatewaysOutput{
			NatGateways: []ec2Types.NatGateway{l},
		}, nil)
	return client.Services{
		EC2: m,
	}
}

func buildEc2SecurityGroups(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)
	l := ec2Types.SecurityGroup{}
	err := faker.FakeData(&l)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeSecurityGroups(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeSecurityGroupsOutput{
			SecurityGroups: []ec2Types.SecurityGroup{l},
		}, nil)
	return client.Services{
		EC2: m,
	}
}

func buildEc2NetworkAcls(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)

	l := ec2Types.NetworkAcl{}
	err := faker.FakeData(&l)
	if err != nil {
		t.Fatal(err)
	}
	l.IsDefault = false
	m.EXPECT().DescribeNetworkAcls(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeNetworkAclsOutput{
			NetworkAcls: []ec2Types.NetworkAcl{l},
		}, nil)
	return client.Services{
		EC2: m,
	}
}

func buildEc2RouteTables(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)
	l := ec2Types.RouteTable{}
	err := faker.FakeData(&l)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeRouteTables(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeRouteTablesOutput{
			RouteTables: []ec2Types.RouteTable{l},
		}, nil)
	return client.Services{
		EC2: m,
	}
}

func buildEc2Subnets(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)
	l := ec2Types.Subnet{}
	err := faker.FakeData(&l)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeSubnets(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeSubnetsOutput{
			Subnets: []ec2Types.Subnet{l},
		}, nil)
	return client.Services{
		EC2: m,
	}
}

func buildEc2Vpcs(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)
	l := ec2Types.Vpc{}
	err := faker.FakeData(&l)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeVpcs(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeVpcsOutput{
			Vpcs: []ec2Types.Vpc{l},
		}, nil)
	return client.Services{
		EC2: m,
	}
}

func buildEc2VpcsPeeringConnections(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)
	l := ec2Types.VpcPeeringConnection{}
	err := faker.FakeData(&l)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeVpcPeeringConnections(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeVpcPeeringConnectionsOutput{
			VpcPeeringConnections: []ec2Types.VpcPeeringConnection{l},
		}, nil)
	return client.Services{
		EC2: m,
	}
}

func buildEc2Instances(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)
	l := ec2Types.Reservation{}
	err := faker.FakeData(&l)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeInstances(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeInstancesOutput{
			Reservations: []ec2Types.Reservation{l},
		}, nil)
	return client.Services{
		EC2: m,
	}
}
func buildEfsFilesystemsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEfsClient(ctrl)
	l := efsTypes.FileSystemDescription{}
	err := faker.FakeData(&l)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeFileSystems(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&efs.DescribeFileSystemsOutput{
			FileSystems: []efsTypes.FileSystemDescription{l},
		}, nil)
	return client.Services{
		EFS: m,
	}
}
func buildEcrRepositoriesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEcrClient(ctrl)
	l := ecrTypes.Repository{}
	err := faker.FakeData(&l)
	if err != nil {
		t.Fatal(err)
	}
	i := ecrTypes.ImageDetail{}
	err = faker.FakeData(&i)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeRepositories(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ecr.DescribeRepositoriesOutput{
			Repositories: []ecrTypes.Repository{l},
		}, nil)

	m.EXPECT().DescribeImages(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ecr.DescribeImagesOutput{
			ImageDetails: []ecrTypes.ImageDetail{i},
		}, nil)
	return client.Services{
		ECR: m,
	}
}

func buildElasticbeanstalkEnvironments(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockElasticbeanstalkClient(ctrl)
	l := elasticbeanstalkTypes.EnvironmentDescription{}
	err := faker.FakeData(&l)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeEnvironments(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&elasticbeanstalk.DescribeEnvironmentsOutput{
			Environments: []elasticbeanstalkTypes.EnvironmentDescription{l},
		}, nil)
	return client.Services{
		ElasticBeanstalk: m,
	}
}

func buildElbv2LoadBalancers(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockElbV2Client(ctrl)
	l := elbv2Types.LoadBalancer{}
	err := faker.FakeData(&l)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeLoadBalancers(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&elasticloadbalancingv2.DescribeLoadBalancersOutput{
			LoadBalancers: []elbv2Types.LoadBalancer{l},
		}, nil)
	return client.Services{
		ELBv2: m,
	}
}

func buildElbv2TargetGroups(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockElbV2Client(ctrl)
	l := elbv2Types.TargetGroup{}
	err := faker.FakeData(&l)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeTargetGroups(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&elasticloadbalancingv2.DescribeTargetGroupsOutput{
			TargetGroups: []elbv2Types.TargetGroup{l},
		}, nil)
	return client.Services{
		ELBv2: m,
	}
}

func buildEmrClusters(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEmrClient(ctrl)
	l := emrTypes.ClusterSummary{}
	err := faker.FakeData(&l)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListClusters(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&emr.ListClustersOutput{
			Clusters: []emrTypes.ClusterSummary{l},
		}, nil)
	return client.Services{
		EMR: m,
	}
}

func buildEksClusters(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEksClient(ctrl)
	l := eks.DescribeClusterOutput{}
	err := faker.FakeData(&l)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListClusters(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&eks.ListClustersOutput{
			Clusters: []string{"test-cluster"},
		}, nil)
	m.EXPECT().DescribeCluster(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&l, nil)
	return client.Services{
		Eks: m,
	}
}

func buildFsxBackups(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockFsxClient(ctrl)
	l := fsxTypes.Backup{}
	err := faker.FakeData(&l)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeBackups(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&fsx.DescribeBackupsOutput{
			Backups: []fsxTypes.Backup{l},
		}, nil)
	return client.Services{
		FSX: m,
	}
}

func buildSnsTopics(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockSnsClient(ctrl)
	topic := snsTypes.Topic{}
	err := faker.FakeData(&topic)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListTopics(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&sns.ListTopicsOutput{
			Topics: []snsTypes.Topic{topic},
		}, nil)
	m.EXPECT().GetTopicAttributes(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&sns.GetTopicAttributesOutput{
			Attributes: map[string]string{
				"SubscriptionsConfirmed":    "5",
				"SubscriptionsDeleted":      "3",
				"SubscriptionsPending":      "0",
				"FifoTopic":                 "false",
				"ContentBasedDeduplication": "true",
				"DisplayName":               "cloudquery",
				"Owner":                     "owner",
				"Policy":                    `{"stuff": 3}`,
				"DeliveryPolicy":            `{"stuff": 3}`,
				"EffectiveDeliveryPolicy":   `{"stuff": 3}`,
			},
		}, nil)
	return client.Services{
		SNS: m,
	}
}

func buildSnsSubscriptions(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockSnsClient(ctrl)
	sub := snsTypes.Subscription{}
	err := faker.FakeData(&sub)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListSubscriptions(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&sns.ListSubscriptionsOutput{
			Subscriptions: []snsTypes.Subscription{sub},
		}, nil)
	return client.Services{
		SNS: m,
	}
}

func buildRdsCertificates(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockRdsClient(ctrl)
	l := rdsTypes.Certificate{}
	err := faker.FakeData(&l)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeCertificates(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&rds.DescribeCertificatesOutput{
			Certificates: []rdsTypes.Certificate{l},
		}, nil)
	return client.Services{
		RDS: m,
	}
}

func buildRdsDBClusters(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockRdsClient(ctrl)
	l := rdsTypes.DBCluster{}
	err := faker.FakeData(&l)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeDBClusters(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&rds.DescribeDBClustersOutput{
			DBClusters: []rdsTypes.DBCluster{l},
		}, nil)
	return client.Services{
		RDS: m,
	}
}

func buildRdsDBInstances(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockRdsClient(ctrl)
	l := rdsTypes.DBInstance{}
	err := faker.FakeData(&l)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeDBInstances(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&rds.DescribeDBInstancesOutput{
			DBInstances: []rdsTypes.DBInstance{l},
		}, nil)
	return client.Services{
		RDS: m,
	}
}

func buildRdsDBSubnetGroups(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockRdsClient(ctrl)
	l := rdsTypes.DBSubnetGroup{}
	err := faker.FakeData(&l)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeDBSubnetGroups(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&rds.DescribeDBSubnetGroupsOutput{
			DBSubnetGroups: []rdsTypes.DBSubnetGroup{l},
		}, nil)
	return client.Services{
		RDS: m,
	}
}

func buildIamGroups(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIamClient(ctrl)
	g := iamTypes.Group{}
	err := faker.FakeData(&g)
	if err != nil {
		t.Fatal(err)
	}

	p := iamTypes.AttachedPolicy{}
	err = faker.FakeData(&p)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListGroups(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&iam.ListGroupsOutput{
			Groups: []iamTypes.Group{g},
		}, nil)
	m.EXPECT().ListAttachedGroupPolicies(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&iam.ListAttachedGroupPoliciesOutput{
			AttachedPolicies: []iamTypes.AttachedPolicy{p},
		}, nil)
	return client.Services{
		IAM: m,
	}
}

func buildIamPolicies(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIamClient(ctrl)
	g := iamTypes.ManagedPolicyDetail{}
	err := faker.FakeData(&g)
	if err != nil {
		t.Fatal(err)
	}
	document := `{"stuff": 3}`
	// generate valid json
	for i := range g.PolicyVersionList {
		g.PolicyVersionList[i].Document = &document
	}

	m.EXPECT().GetAccountAuthorizationDetails(gomock.Any(), gomock.Any()).Return(
		&iam.GetAccountAuthorizationDetailsOutput{
			Policies: []iamTypes.ManagedPolicyDetail{g},
		}, nil)
	return client.Services{
		IAM: m,
	}
}

func buildIamPasswordPolicies(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIamClient(ctrl)
	g := iamTypes.PasswordPolicy{}
	err := faker.FakeData(&g)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().GetAccountPasswordPolicy(gomock.Any(), gomock.Any()).Return(
		&iam.GetAccountPasswordPolicyOutput{
			PasswordPolicy: &g,
		}, nil)
	return client.Services{
		IAM: m,
	}
}

func buildIamVirtualMfaDevices(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIamClient(ctrl)
	g := iamTypes.VirtualMFADevice{}
	err := faker.FakeData(&g)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListVirtualMFADevices(gomock.Any(), gomock.Any()).Return(
		&iam.ListVirtualMFADevicesOutput{
			VirtualMFADevices: []iamTypes.VirtualMFADevice{g},
		}, nil)
	return client.Services{
		IAM: m,
	}
}

func buildIamRoles(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIamClient(ctrl)
	r := iamTypes.Role{}
	err := faker.FakeData(&r)
	if err != nil {
		t.Fatal(err)
	}

	p := iamTypes.AttachedPolicy{}
	err = faker.FakeData(&p)
	if err != nil {
		t.Fatal(err)
	}

	// generate valid json
	document := `{"stuff": 3}`
	r.AssumeRolePolicyDocument = &document

	m.EXPECT().ListRoles(gomock.Any(), gomock.Any()).Return(
		&iam.ListRolesOutput{
			Roles: []iamTypes.Role{r},
		}, nil)
	m.EXPECT().ListAttachedRolePolicies(gomock.Any(), gomock.Any()).Return(
		&iam.ListAttachedRolePoliciesOutput{
			AttachedPolicies: []iamTypes.AttachedPolicy{p},
		}, nil)
	return client.Services{
		IAM: m,
	}
}

func buildIamUsers(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIamClient(ctrl)
	u := iamTypes.User{}
	err := faker.FakeData(&u)
	if err != nil {
		t.Fatal(err)
	}
	g := iamTypes.Group{}
	err = faker.FakeData(&g)
	if err != nil {
		t.Fatal(err)
	}
	km := iamTypes.AccessKeyMetadata{}
	err = faker.FakeData(&km)
	if err != nil {
		t.Fatal(err)
	}
	aup := iamTypes.AttachedPolicy{}
	err = faker.FakeData(&aup)
	if err != nil {
		t.Fatal(err)
	}
	akl := iam.GetAccessKeyLastUsedOutput{}
	err = faker.FakeData(&akl)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListUsers(gomock.Any(), gomock.Any()).Return(
		&iam.ListUsersOutput{
			Users: []iamTypes.User{u},
		}, nil)
	m.EXPECT().ListGroupsForUser(gomock.Any(), gomock.Any()).Return(
		&iam.ListGroupsForUserOutput{
			Groups: []iamTypes.Group{g},
		}, nil)
	m.EXPECT().GetCredentialReport(gomock.Any(), gomock.Any()).Return(
		nil, nil)
	m.EXPECT().ListAccessKeys(gomock.Any(), gomock.Any()).Return(
		&iam.ListAccessKeysOutput{
			AccessKeyMetadata: []iamTypes.AccessKeyMetadata{km},
		}, nil)
	m.EXPECT().ListAttachedUserPolicies(gomock.Any(), gomock.Any()).Return(
		&iam.ListAttachedUserPoliciesOutput{
			AttachedPolicies: []iamTypes.AttachedPolicy{aup},
		}, nil)
	m.EXPECT().GetAccessKeyLastUsed(gomock.Any(), gomock.Any()).Return(
		&akl, nil)
	return client.Services{
		IAM: m,
	}
}

func buildKmsKeys(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockKmsClient(ctrl)
	k := kmsTypes.KeyListEntry{}
	err := faker.FakeData(&k)
	if err != nil {
		t.Fatal(err)
	}

	km := kms.DescribeKeyOutput{}
	err = faker.FakeData(&km)
	if err != nil {
		t.Fatal(err)
	}

	krs := kms.GetKeyRotationStatusOutput{}
	err = faker.FakeData(&krs)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListKeys(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&kms.ListKeysOutput{
			Keys: []kmsTypes.KeyListEntry{k},
		}, nil)
	m.EXPECT().DescribeKey(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&km, nil)
	m.EXPECT().GetKeyRotationStatus(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&krs, nil)
	return client.Services{
		KMS: m,
	}
}

func buildOrganizationsAccounts(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockOrganizationsClient(ctrl)
	g := organizationsTypes.Account{}
	err := faker.FakeData(&g)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListAccounts(gomock.Any(), gomock.Any()).Return(
		&organizations.ListAccountsOutput{
			Accounts: []organizationsTypes.Account{g},
		}, nil)
	return client.Services{
		Organizations: m,
	}
}

func buildS3Buckets(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockS3Client(ctrl)
	b := s3Types.Bucket{}
	err := faker.FakeData(&b)
	if err != nil {
		t.Fatal(err)
	}
	bloc := s3.GetBucketLocationOutput{}
	err = faker.FakeData(&bloc)
	if err != nil {
		t.Fatal(err)
	}
	blog := s3.GetBucketLoggingOutput{}
	err = faker.FakeData(&blog)
	if err != nil {
		t.Fatal(err)
	}
	bpol := s3.GetBucketPolicyOutput{}
	err = faker.FakeData(&bpol)
	if err != nil {
		t.Fatal(err)
	}
	jsonDoc := `{"stuff": 3}`
	bpol.Policy = &jsonDoc
	bver := s3.GetBucketVersioningOutput{}
	err = faker.FakeData(&bver)
	if err != nil {
		t.Fatal(err)
	}
	bgrant := s3Types.Grant{}
	err = faker.FakeData(&bgrant)
	if err != nil {
		t.Fatal(err)
	}
	bcors := s3Types.CORSRule{}
	err = faker.FakeData(&bcors)
	if err != nil {
		t.Fatal(err)
	}
	bencryption := s3.GetBucketEncryptionOutput{}
	err = faker.FakeData(&bencryption)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListBuckets(gomock.Any(), gomock.Any()).Return(
		&s3.ListBucketsOutput{
			Buckets: []s3Types.Bucket{b},
		}, nil)
	m.EXPECT().GetBucketLocation(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&bloc, nil)
	m.EXPECT().GetBucketLogging(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&blog, nil)
	m.EXPECT().GetBucketPolicy(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&bpol, nil)
	m.EXPECT().GetBucketVersioning(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&bver, nil)
	m.EXPECT().GetBucketAcl(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&s3.GetBucketAclOutput{
			Grants: []s3Types.Grant{bgrant},
		}, nil)
	m.EXPECT().GetBucketCors(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&s3.GetBucketCorsOutput{
			CORSRules: []s3Types.CORSRule{bcors},
		}, nil)
	m.EXPECT().GetBucketEncryption(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&bencryption, nil)
	return client.Services{
		S3: m,
	}
}
