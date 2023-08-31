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

func buildBatchJobQueuesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockBatchClient(ctrl)

	buildBatchJobsMock(t, m)

	services := client.Services{
		Batch: m,
	}
	a := types.JobQueueDetail{}
	require.NoError(t, faker.FakeObject(&a))

	m.EXPECT().DescribeJobQueues(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&batch.DescribeJobQueuesOutput{
			JobQueues: []types.JobQueueDetail{a},
		}, nil)

	tagResponse := batch.ListTagsForResourceOutput{}
	require.NoError(t, faker.FakeObject(&tagResponse))
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(&tagResponse, nil)

	return services
}

func TestBatchJobQueues(t *testing.T) {
	client.AwsMockTestHelper(t, JobQueues(), buildBatchJobQueuesMock, client.TestOptions{})
}
