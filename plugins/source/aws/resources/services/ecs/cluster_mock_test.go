package ecs

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ecs"
	ecsTypes "github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildEcsClusterMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEcsClient(ctrl)
	services := client.Services{
		ECS: m,
	}
	faker.SetIgnoreInterface(true)
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
	var tags ecs.ListTagsForResourceOutput
	err = faker.FakeData(&tags)
	if err != nil {
		t.Fatal(err)
	}

	servicesList := ecs.ListServicesOutput{
		ServiceArns: []string{"test"},
	}
	m.EXPECT().ListServices(gomock.Any(), gomock.Any(), gomock.Any()).Return(&servicesList, nil)

	svcs := ecs.DescribeServicesOutput{}
	err = faker.FakeData(&svcs)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeServices(gomock.Any(), gomock.Any(), gomock.Any()).Return(&svcs, nil)

	instancesList := ecs.ListContainerInstancesOutput{
		ContainerInstanceArns: []string{"test"},
	}
	m.EXPECT().ListContainerInstances(gomock.Any(), gomock.Any(), gomock.Any()).Return(&instancesList, nil)

	instances := ecs.DescribeContainerInstancesOutput{}
	err = faker.FakeData(&instances)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeContainerInstances(gomock.Any(), gomock.Any(), gomock.Any()).Return(&instances, nil)

	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(&tags, nil)

	listTasks := ecs.ListTasksOutput{}
	err = faker.FakeData(&listTasks)
	if err != nil {
		t.Fatal(err)
	}
	listTasks.NextToken = nil
	m.EXPECT().ListTasks(gomock.Any(), gomock.Any(), gomock.Any()).Return(&listTasks, nil)

	tasks := ecs.DescribeTasksOutput{}
	err = faker.FakeData(&tasks)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeTasks(gomock.Any(), gomock.Any(), gomock.Any()).Return(&tasks, nil)

	return services
}

func TestEcsClusters(t *testing.T) {
	client.AwsMockTestHelper(t, Clusters(), buildEcsClusterMock, client.TestOptions{})
}
