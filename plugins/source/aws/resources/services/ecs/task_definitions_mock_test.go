package ecs

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildEcsTaskDefinitions(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEcsClient(ctrl)

	listTaskDefinitionsOutput := ecs.ListTaskDefinitionsOutput{}
	require.NoError(t, faker.FakeObject(&listTaskDefinitionsOutput))
	listTaskDefinitionsOutput.NextToken = nil
	m.EXPECT().ListTaskDefinitions(gomock.Any(), gomock.Any(), gomock.Any()).Return(&listTaskDefinitionsOutput, nil)

	taskDefinition := &ecs.DescribeTaskDefinitionOutput{}
	require.NoError(t, faker.FakeObject(&taskDefinition))
	m.EXPECT().DescribeTaskDefinition(gomock.Any(), gomock.Any(), gomock.Any()).Return(taskDefinition, nil)

	return client.Services{
		Ecs: m,
	}
}

func TestEcsTaskDefinitions(t *testing.T) {
	client.AwsMockTestHelper(t, TaskDefinitions(), buildEcsTaskDefinitions, client.TestOptions{})
}
