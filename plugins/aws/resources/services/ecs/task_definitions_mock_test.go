package ecs

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildEcsTaskDefinitions(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEcsClient(ctrl)

	faker.SetIgnoreInterface(true)
	listTaskDefinitionsOutput := ecs.ListTaskDefinitionsOutput{}
	err := faker.FakeData(&listTaskDefinitionsOutput)
	if err != nil {
		t.Fatal(err)
	}
	listTaskDefinitionsOutput.NextToken = nil
	m.EXPECT().ListTaskDefinitions(gomock.Any(), gomock.Any(), gomock.Any()).Return(&listTaskDefinitionsOutput, nil)

	taskDefinition := &ecs.DescribeTaskDefinitionOutput{}
	err = faker.FakeData(&taskDefinition)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeTaskDefinition(gomock.Any(), gomock.Any(), gomock.Any()).Return(taskDefinition, nil)

	return client.Services{
		ECS: m,
	}
}

func TestEcsTaskDefinitions(t *testing.T) {
	client.AwsMockTestHelper(t, EcsTaskDefinitions(), buildEcsTaskDefinitions, client.TestOptions{})
}
