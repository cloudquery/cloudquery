package ecs

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildEcsClusterMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEcsClient(ctrl)
	services := client.Services{
		Ecs: m,
	}
	c := types.Cluster{}
	err := faker.FakeObject(&c)
	if err != nil {
		t.Fatal(err)
	}
	ecsOutput := &ecs.DescribeClustersOutput{
		Clusters: []types.Cluster{c},
	}
	m.EXPECT().DescribeClusters(gomock.Any(), gomock.Any(), gomock.Any()).Return(ecsOutput, nil)
	ecsListOutput := &ecs.ListClustersOutput{
		ClusterArns: []string{"randomClusteArn"},
	}
	m.EXPECT().ListClusters(gomock.Any(), gomock.Any(), gomock.Any()).Return(ecsListOutput, nil)

	servicesList := ecs.ListServicesOutput{
		ServiceArns: []string{"test"},
	}
	m.EXPECT().ListServices(gomock.Any(), gomock.Any(), gomock.Any()).Return(&servicesList, nil)

	svcs := ecs.DescribeServicesOutput{}
	err = faker.FakeObject(&svcs)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeServices(gomock.Any(), gomock.Any(), gomock.Any()).Return(&svcs, nil)

	instancesList := ecs.ListContainerInstancesOutput{
		ContainerInstanceArns: []string{"test"},
	}
	m.EXPECT().ListContainerInstances(gomock.Any(), gomock.Any(), gomock.Any()).Return(&instancesList, nil)

	instances := ecs.DescribeContainerInstancesOutput{}
	err = faker.FakeObject(&instances)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeContainerInstances(gomock.Any(), gomock.Any(), gomock.Any()).Return(&instances, nil)

	listTasks := ecs.ListTasksOutput{}
	err = faker.FakeObject(&listTasks)
	if err != nil {
		t.Fatal(err)
	}
	listTasks.NextToken = nil
	m.EXPECT().ListTasks(gomock.Any(), gomock.Any(), gomock.Any()).Return(&listTasks, nil)

	tasks := ecs.DescribeTasksOutput{}
	err = faker.FakeObject(&tasks)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeTasks(gomock.Any(), gomock.Any(), gomock.Any()).Return(&tasks, nil)

	protection := ecs.GetTaskProtectionOutput{}
	err = faker.FakeObject(&protection)
	if err != nil {
		t.Fatal(err)
	}
	protection.Failures = nil
	m.EXPECT().GetTaskProtection(gomock.Any(), gomock.Any(), gomock.Any()).Return(&protection, nil)

	taskSet := types.TaskSet{}
	err = faker.FakeObject(&taskSet)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeTaskSets(gomock.Any(), gomock.Any(), gomock.Any()).Return(&ecs.DescribeTaskSetsOutput{
		TaskSets: []types.TaskSet{taskSet},
		Failures: nil,
	}, nil)

	return services
}

func TestEcsClusters(t *testing.T) {
	client.AwsMockTestHelper(t, Clusters(), buildEcsClusterMock, client.TestOptions{})
}
