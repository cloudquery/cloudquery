package ecs

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildEcsClusterMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEcsClient(ctrl)
	services := client.Services{
		Ecs: m,
	}
	c := types.Cluster{}
	require.NoError(t, faker.FakeObject(&c))
	ecsOutput := &ecs.DescribeClustersOutput{
		Clusters: []types.Cluster{c},
	}
	m.EXPECT().DescribeClusters(gomock.Any(), gomock.Any(), gomock.Any()).Return(ecsOutput, nil)
	ecsListOutput := &ecs.ListClustersOutput{
		ClusterArns: []string{"randomClusterArn"},
	}
	m.EXPECT().ListClusters(gomock.Any(), gomock.Any(), gomock.Any()).Return(ecsListOutput, nil)

	servicesList := ecs.ListServicesOutput{
		ServiceArns: []string{"test"},
	}
	m.EXPECT().ListServices(gomock.Any(), gomock.Any(), gomock.Any()).Return(&servicesList, nil)

	svcs := ecs.DescribeServicesOutput{}
	require.NoError(t, faker.FakeObject(&svcs))
	m.EXPECT().DescribeServices(gomock.Any(), gomock.Any(), gomock.Any()).Return(&svcs, nil)

	instancesList := ecs.ListContainerInstancesOutput{
		ContainerInstanceArns: []string{"test"},
	}
	m.EXPECT().ListContainerInstances(gomock.Any(), gomock.Any(), gomock.Any()).Return(&instancesList, nil)

	instances := ecs.DescribeContainerInstancesOutput{}
	require.NoError(t, faker.FakeObject(&instances))
	m.EXPECT().DescribeContainerInstances(gomock.Any(), gomock.Any(), gomock.Any()).Return(&instances, nil)

	listTasks := ecs.ListTasksOutput{}
	require.NoError(t, faker.FakeObject(&listTasks))
	listTasks.NextToken = nil
	m.EXPECT().ListTasks(gomock.Any(), gomock.Any(), gomock.Any()).Return(&listTasks, nil)

	tasks := ecs.DescribeTasksOutput{}
	require.NoError(t, faker.FakeObject(&tasks))
	m.EXPECT().DescribeTasks(gomock.Any(), gomock.Any(), gomock.Any()).Return(&tasks, nil)

	protection := ecs.GetTaskProtectionOutput{}
	require.NoError(t, faker.FakeObject(&protection))
	protection.Failures = nil
	m.EXPECT().GetTaskProtection(gomock.Any(), gomock.Any(), gomock.Any()).Return(&protection, nil)

	taskSet := types.TaskSet{}
	require.NoError(t, faker.FakeObject(&taskSet))
	m.EXPECT().DescribeTaskSets(gomock.Any(), gomock.Any(), gomock.Any()).Return(&ecs.DescribeTaskSetsOutput{
		TaskSets: []types.TaskSet{taskSet},
		Failures: nil,
	}, nil)

	return services
}

func TestEcsClusters(t *testing.T) {
	client.AwsMockTestHelper(t, Clusters(), buildEcsClusterMock, client.TestOptions{})
}
