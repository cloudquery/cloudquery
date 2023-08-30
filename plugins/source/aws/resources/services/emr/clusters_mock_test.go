package emr

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/emr"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildEMRClusters(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockEmrClient(ctrl)
	var summary1 types.ClusterSummary
	require.NoError(t, faker.FakeObject(&summary1))

	summary2 := summary1
	summary1.Id = aws.String("cluster1")
	summary2.Id = aws.String("cluster2")
	mock.EXPECT().ListClusters(gomock.Any(), &emr.ListClustersInput{ClusterStates: []types.ClusterState{
		types.ClusterStateRunning,
		types.ClusterStateStarting,
		types.ClusterStateBootstrapping,
		types.ClusterStateWaiting,
	}}, gomock.Any()).Return(
		&emr.ListClustersOutput{Clusters: []types.ClusterSummary{summary1, summary2}},
		nil,
	)

	var cluster1 types.Cluster
	require.NoError(t, faker.FakeObject(&cluster1))

	cluster1.Id = summary1.Id
	cluster1.InstanceCollectionType = types.InstanceCollectionTypeInstanceFleet
	cluster1.RepoUpgradeOnBoot = types.RepoUpgradeOnBootNone
	cluster1.ScaleDownBehavior = types.ScaleDownBehaviorTerminateAtInstanceHour
	var config types.Configuration
	require.NoError(t, faker.FakeObject(&config))

	config.Configurations = []types.Configuration{}
	cluster1.Configurations = []types.Configuration{config}

	cluster2 := cluster1
	cluster2.Id = summary2.Id
	cluster1.InstanceCollectionType = types.InstanceCollectionTypeInstanceFleet
	cluster2.InstanceCollectionType = types.InstanceCollectionTypeInstanceGroup

	mock.EXPECT().DescribeCluster(gomock.Any(), &emr.DescribeClusterInput{ClusterId: summary1.Id}, gomock.Any()).Return(
		&emr.DescribeClusterOutput{Cluster: &cluster1},
		nil,
	)

	mock.EXPECT().DescribeCluster(gomock.Any(), &emr.DescribeClusterInput{ClusterId: summary2.Id}, gomock.Any()).Return(
		&emr.DescribeClusterOutput{Cluster: &cluster2},
		nil,
	)

	var instanceFleet types.InstanceFleet
	require.NoError(t, faker.FakeObject(&instanceFleet))

	mock.EXPECT().ListInstanceFleets(gomock.Any(), &emr.ListInstanceFleetsInput{ClusterId: summary1.Id}, gomock.Any()).Return(
		&emr.ListInstanceFleetsOutput{InstanceFleets: []types.InstanceFleet{instanceFleet}},
		nil,
	)

	var instanceGroup types.InstanceGroup
	require.NoError(t, faker.FakeObject(&instanceGroup))

	mock.EXPECT().ListInstanceGroups(gomock.Any(), &emr.ListInstanceGroupsInput{ClusterId: summary2.Id}, gomock.Any()).Return(
		&emr.ListInstanceGroupsOutput{InstanceGroups: []types.InstanceGroup{instanceGroup}},
		nil,
	)

	var instance types.Instance
	require.NoError(t, faker.FakeObject(&instance))

	mock.EXPECT().ListInstances(gomock.Any(), &emr.ListInstancesInput{ClusterId: summary1.Id}, gomock.Any()).Return(
		&emr.ListInstancesOutput{Instances: []types.Instance{instance}},
		nil,
	)

	mock.EXPECT().ListInstances(gomock.Any(), &emr.ListInstancesInput{ClusterId: summary2.Id}, gomock.Any()).Return(
		&emr.ListInstancesOutput{Instances: []types.Instance{instance}},
		nil,
	)

	var notebookExecutionSummary types.NotebookExecutionSummary
	require.NoError(t, faker.FakeObject(&notebookExecutionSummary))

	mock.EXPECT().ListNotebookExecutions(gomock.Any(), &emr.ListNotebookExecutionsInput{ExecutionEngineId: summary1.Id}, gomock.Any()).Return(
		&emr.ListNotebookExecutionsOutput{NotebookExecutions: []types.NotebookExecutionSummary{notebookExecutionSummary}},
		nil,
	)

	mock.EXPECT().ListNotebookExecutions(gomock.Any(), &emr.ListNotebookExecutionsInput{ExecutionEngineId: summary2.Id}, gomock.Any()).Return(
		&emr.ListNotebookExecutionsOutput{NotebookExecutions: []types.NotebookExecutionSummary{}},
		nil,
	)

	var notebookExecution types.NotebookExecution
	require.NoError(t, faker.FakeObject(&notebookExecution))

	mock.EXPECT().DescribeNotebookExecution(gomock.Any(), &emr.DescribeNotebookExecutionInput{NotebookExecutionId: notebookExecutionSummary.NotebookExecutionId}, gomock.Any()).Return(
		&emr.DescribeNotebookExecutionOutput{NotebookExecution: &notebookExecution},
		nil,
	)

	var stepSummary types.StepSummary
	require.NoError(t, faker.FakeObject(&stepSummary))

	mock.EXPECT().ListSteps(gomock.Any(), &emr.ListStepsInput{ClusterId: summary1.Id}, gomock.Any()).Return(
		&emr.ListStepsOutput{Steps: []types.StepSummary{stepSummary}},
		nil,
	)

	mock.EXPECT().ListSteps(gomock.Any(), &emr.ListStepsInput{ClusterId: summary2.Id}, gomock.Any()).Return(
		&emr.ListStepsOutput{Steps: []types.StepSummary{}},
		nil,
	)

	var step types.Step
	require.NoError(t, faker.FakeObject(&step))

	mock.EXPECT().DescribeStep(gomock.Any(), &emr.DescribeStepInput{ClusterId: summary1.Id, StepId: stepSummary.Id}, gomock.Any()).Return(
		&emr.DescribeStepOutput{Step: &step},
		nil,
	)

	return client.Services{Emr: mock}
}

func TestEMRClusters(t *testing.T) {
	client.AwsMockTestHelper(t, Clusters(), buildEMRClusters, client.TestOptions{})
}
