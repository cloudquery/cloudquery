package mocks_test

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	autoscalingTypes "github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	cloudfrontTypes "github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	cloudwatchTypes "github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	cloudwatchlogsTypes "github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice"
	databasemigrationserviceTypes "github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
	"github.com/aws/aws-sdk-go-v2/service/directconnect"
	directconnectTypes "github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	ec2Types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	ecrTypes "github.com/aws/aws-sdk-go-v2/service/ecr/types"
	"github.com/aws/aws-sdk-go-v2/service/efs"
	efsTypes "github.com/aws/aws-sdk-go-v2/service/efs/types"
	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	iamTypes "github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	kmsTypes "github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/aws/aws-sdk-go-v2/service/organizations"
	organizationsTypes "github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	route53Types "github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	snsTypes "github.com/aws/aws-sdk-go-v2/service/sns/types"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"

	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
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

func buildCloudfrontCachePoliciesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockCloudfrontClient(ctrl)
	services := client.Services{
		Cloudfront: m,
	}
	cp := cloudfrontTypes.CachePolicySummary{}
	if err := faker.FakeData(&cp); err != nil {
		t.Fatal(err)
	}

	cloudfrontOutput := &cloudfront.ListCachePoliciesOutput{
		CachePolicyList: &cloudfrontTypes.CachePolicyList{
			Items: []cloudfrontTypes.CachePolicySummary{cp},
		},
	}
	m.EXPECT().ListCachePolicies(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		cloudfrontOutput,
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

func buildRoute53TrafficPoliciesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockRoute53Client(ctrl)
	tps := route53Types.TrafficPolicySummary{}
	if err := faker.FakeData(&tps); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListTrafficPolicies(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&route53.ListTrafficPoliciesOutput{
			TrafficPolicySummaries: []route53Types.TrafficPolicySummary{tps},
		}, nil)
	tp := route53Types.TrafficPolicy{}
	if err := faker.FakeData(&tp); err != nil {
		t.Fatal(err)
	}
	tp.Id = tps.Id
	jsonStr := "{\"test\": \"test\"}"
	tp.Document = &jsonStr
	m.EXPECT().ListTrafficPolicyVersions(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&route53.ListTrafficPolicyVersionsOutput{
			TrafficPolicies: []route53Types.TrafficPolicy{tp},
		}, nil)
	return client.Services{
		Route53: m,
	}
}

func buildRoute53DelegationSetsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockRoute53Client(ctrl)
	ds := route53Types.DelegationSet{}
	if err := faker.FakeData(&ds); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListReusableDelegationSets(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&route53.ListReusableDelegationSetsOutput{
			DelegationSets: []route53Types.DelegationSet{ds},
		}, nil)
	return client.Services{
		Route53: m,
	}
}

func buildRoute53HealthChecksMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockRoute53Client(ctrl)
	hc := route53Types.HealthCheck{}
	if err := faker.FakeData(&hc); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListHealthChecks(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&route53.ListHealthChecksOutput{
			HealthChecks: []route53Types.HealthCheck{hc},
		}, nil)
	tag := route53Types.Tag{}
	if err := faker.FakeData(&tag); err != nil {
		t.Fatal(err)
	}
	//m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(
	//	&route53.ListTagsForResourceOutput{
	//		ResourceTagSet: &route53Types.ResourceTagSet{
	//			Tags: []route53Types.Tag{tag},
	//		},
	//	}, nil)
	m.EXPECT().ListTagsForResources(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&route53.ListTagsForResourcesOutput{
			ResourceTagSets: []route53Types.ResourceTagSet{
				{
					ResourceId: hc.Id,
					Tags:       []route53Types.Tag{tag},
				},
			},
		}, nil)
	return client.Services{
		Route53: m,
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
	association := directconnectTypes.DirectConnectGatewayAssociation{}
	attachment := directconnectTypes.DirectConnectGatewayAttachment{}
	err := faker.FakeData(&l)
	if err != nil {
		t.Fatal(err)
	}
	err = faker.FakeData(&association)
	if err != nil {
		t.Fatal(err)
	}
	err = faker.FakeData(&attachment)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeDirectConnectGateways(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&directconnect.DescribeDirectConnectGatewaysOutput{
			DirectConnectGateways: []directconnectTypes.DirectConnectGateway{l},
		}, nil)
	m.EXPECT().DescribeDirectConnectGatewayAssociations(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&directconnect.DescribeDirectConnectGatewayAssociationsOutput{
			DirectConnectGatewayAssociations: []directconnectTypes.DirectConnectGatewayAssociation{association},
		}, nil)
	m.EXPECT().DescribeDirectConnectGatewayAttachments(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&directconnect.DescribeDirectConnectGatewayAttachmentsOutput{
			DirectConnectGatewayAttachments: []directconnectTypes.DirectConnectGatewayAttachment{attachment},
		}, nil)
	return client.Services{
		Directconnect: m,
	}
}

func buildDirectconnectVirtualGatewaysMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockDirectconnectClient(ctrl)
	l := directconnectTypes.VirtualGateway{}
	err := faker.FakeData(&l)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeVirtualGateways(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&directconnect.DescribeVirtualGatewaysOutput{
			VirtualGateways: []directconnectTypes.VirtualGateway{l},
		}, nil)
	return client.Services{
		Directconnect: m,
	}
}

func buildDirectconnectVirtualInterfacesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockDirectconnectClient(ctrl)
	l := directconnectTypes.VirtualInterface{}
	err := faker.FakeData(&l)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeVirtualInterfaces(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&directconnect.DescribeVirtualInterfacesOutput{
			VirtualInterfaces: []directconnectTypes.VirtualInterface{l},
		}, nil)
	return client.Services{
		Directconnect: m,
	}
}

func buildDmsReplicationInstances(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockDatabasemigrationserviceClient(ctrl)
	l := databasemigrationserviceTypes.ReplicationInstance{}
	if err := faker.FakeData(&l); err != nil {
		t.Fatal(err)
	}
	l.ReplicationInstancePrivateIpAddress = aws.String("1.2.3.4")
	l.ReplicationInstancePrivateIpAddresses = []string{"1.2.3.4"}
	l.ReplicationInstancePublicIpAddress = aws.String("1.2.3.4")
	l.ReplicationInstancePublicIpAddresses = []string{"1.2.3.4"}
	m.EXPECT().DescribeReplicationInstances(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&databasemigrationservice.DescribeReplicationInstancesOutput{
			ReplicationInstances: []databasemigrationserviceTypes.ReplicationInstance{l},
		}, nil)
	lt := databasemigrationserviceTypes.Tag{}
	if err := faker.FakeData(&lt); err != nil {
		t.Fatal(err)
	}
	lt.ResourceArn = l.ReplicationInstanceArn
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&databasemigrationservice.ListTagsForResourceOutput{
			TagList: []databasemigrationserviceTypes.Tag{lt},
		}, nil)
	return client.Services{
		DMS: m,
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

func buildEc2EbsVolumes(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)
	l := ec2Types.Volume{}
	err := faker.FakeData(&l)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeVolumes(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeVolumesOutput{
			Volumes: []ec2Types.Volume{l},
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
	l.IsDefault = aws.Bool(false)
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

func buildEc2VpcEndpoints(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)
	e := ec2Types.VpcEndpoint{}
	if err := faker.FakeData(&e); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeVpcEndpoints(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeVpcEndpointsOutput{
			VpcEndpoints: []ec2Types.VpcEndpoint{e},
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
				"KmsMasterKeyId":            "test/key",
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

	//list policies
	var l []string
	err = faker.FakeData(&l)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListGroupPolicies(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&iam.ListGroupPoliciesOutput{
			PolicyNames: l,
		}, nil)

	//get policy
	gp := iam.GetGroupPolicyOutput{}
	err = faker.FakeData(&gp)
	if err != nil {
		t.Fatal(err)
	}
	document := "{\"test\": {\"t1\":1}}"
	gp.PolicyDocument = &document
	m.EXPECT().GetGroupPolicy(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&gp, nil)
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

func buildIamOpenIDConnectProviders(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIamClient(ctrl)
	l := iamTypes.OpenIDConnectProviderListEntry{}
	err := faker.FakeData(&l)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListOpenIDConnectProviders(gomock.Any(), gomock.Any()).Return(
		&iam.ListOpenIDConnectProvidersOutput{
			OpenIDConnectProviderList: []iamTypes.OpenIDConnectProviderListEntry{l},
		}, nil)

	p := iam.GetOpenIDConnectProviderOutput{}
	err = faker.FakeData(&p)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetOpenIDConnectProvider(gomock.Any(), gomock.Any()).Return(&p, nil)

	return client.Services{
		IAM: m,
	}
}

func buildIamSAMLProviders(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIamClient(ctrl)
	l := iamTypes.SAMLProviderListEntry{}
	err := faker.FakeData(&l)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListSAMLProviders(gomock.Any(), gomock.Any()).Return(
		&iam.ListSAMLProvidersOutput{
			SAMLProviderList: []iamTypes.SAMLProviderListEntry{l},
		}, nil)

	p := iam.GetSAMLProviderOutput{}
	err = faker.FakeData(&p)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetSAMLProvider(gomock.Any(), gomock.Any()).Return(&p, nil)

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

	tags := kms.ListResourceTagsOutput{}
	err = faker.FakeData(&tags)
	if err != nil {
		t.Fatal(err)
	}
	tags.NextMarker = nil

	m.EXPECT().ListKeys(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&kms.ListKeysOutput{
			Keys: []kmsTypes.KeyListEntry{k},
		}, nil)
	m.EXPECT().DescribeKey(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&km, nil)
	m.EXPECT().GetKeyRotationStatus(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&krs, nil)
	m.EXPECT().ListResourceTags(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&tags, nil)
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
