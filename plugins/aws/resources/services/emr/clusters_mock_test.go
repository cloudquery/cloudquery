package emr

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/emr"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildEMRClusters(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockEmrClient(ctrl)
	var summary types.ClusterSummary
	if err := faker.FakeData(&summary); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListClusters(gomock.Any(), &emr.ListClustersInput{}, gomock.Any()).Return(
		&emr.ListClustersOutput{Clusters: []types.ClusterSummary{summary}},
		nil,
	)

	var cluster types.Cluster
	skipFields := []string{"Configurations", "InstanceCollectionType", "RepoUpgradeOnBoot", "ScaleDownBehavior"}
	if err := faker.FakeDataSkipFields(&cluster, skipFields); err != nil {
		t.Fatal(err)
	}
	cluster.InstanceCollectionType = types.InstanceCollectionTypeInstanceFleet
	cluster.RepoUpgradeOnBoot = types.RepoUpgradeOnBootNone
	cluster.ScaleDownBehavior = types.ScaleDownBehaviorTerminateAtInstanceHour
	var config types.Configuration
	if err := faker.FakeDataSkipFields(&config, []string{"Configurations"}); err != nil {
		t.Fatal(err)
	}
	config.Configurations = []types.Configuration{}
	cluster.Configurations = []types.Configuration{config}
	mock.EXPECT().DescribeCluster(gomock.Any(), &emr.DescribeClusterInput{ClusterId: summary.Id}).Return(
		&emr.DescribeClusterOutput{Cluster: &cluster},
		nil,
	)
	return client.Services{EMR: mock}
}

func TestEMRClusters(t *testing.T) {
	client.AwsMockTestHelper(t, EmrClusters(), buildEMRClusters, client.TestOptions{})
}
