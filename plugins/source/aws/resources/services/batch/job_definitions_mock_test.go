package batch

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/batch"
	"github.com/aws/aws-sdk-go-v2/service/batch/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildBatchJobDefinitionsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockBatchClient(ctrl)
	services := client.Services{
		Batch: m,
	}
	a := types.JobDefinition{}
	require.NoError(t, faker.FakeObject(&a))

	m.EXPECT().DescribeJobDefinitions(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&batch.DescribeJobDefinitionsOutput{
			JobDefinitions: []types.JobDefinition{a},
		}, nil)

	tagResponse := batch.ListTagsForResourceOutput{}
	require.NoError(t, faker.FakeObject(&tagResponse))
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(&tagResponse, nil)

	return services
}

func TestBatchJobDefinitions(t *testing.T) {
	client.AwsMockTestHelper(t, JobDefinitions(), buildBatchJobDefinitionsMock, client.TestOptions{})
}
