package emr

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/emr"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildEMRClusters(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockEmrClient(ctrl)
	var summary types.ClusterSummary
	if err := faker.FakeObject(&summary); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListClusters(gomock.Any(), &emr.ListClustersInput{ClusterStates: []types.ClusterState{
		types.ClusterStateRunning,
		types.ClusterStateStarting,
		types.ClusterStateBootstrapping,
		types.ClusterStateWaiting,
	}}, gomock.Any()).Return(
		&emr.ListClustersOutput{Clusters: []types.ClusterSummary{summary}},
		nil,
	)

	var cluster types.Cluster
	if err := faker.FakeObject(&cluster); err != nil {
		t.Fatal(err)
	}
	cluster.InstanceCollectionType = types.InstanceCollectionTypeInstanceFleet
	cluster.RepoUpgradeOnBoot = types.RepoUpgradeOnBootNone
	cluster.ScaleDownBehavior = types.ScaleDownBehaviorTerminateAtInstanceHour
	var config types.Configuration
	if err := faker.FakeObject(&config); err != nil {
		t.Fatal(err)
	}
	config.Configurations = []types.Configuration{}
	cluster.Configurations = []types.Configuration{config}
	mock.EXPECT().DescribeCluster(gomock.Any(), &emr.DescribeClusterInput{ClusterId: summary.Id}).Return(
		&emr.DescribeClusterOutput{Cluster: &cluster},
		nil,
	)
	return client.Services{Emr: mock}
}

func TestEMRClusters(t *testing.T) {
	client.AwsMockTestHelper(t, Clusters(), buildEMRClusters, client.TestOptions{})
}
