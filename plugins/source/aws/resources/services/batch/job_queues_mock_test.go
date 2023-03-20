package batch

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/batch"
	"github.com/aws/aws-sdk-go-v2/service/batch/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildBatchJobQueuesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockBatchClient(ctrl)

	buildBatchJobsMock(t, m)

	services := client.Services{
		Batch: m,
	}
	a := types.JobQueueDetail{}
	err := faker.FakeObject(&a)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeJobQueues(gomock.Any(), gomock.Any()).Return(
		&batch.DescribeJobQueuesOutput{
			JobQueues: []types.JobQueueDetail{a},
		}, nil)

	tagResponse := batch.ListTagsForResourceOutput{}
	err = faker.FakeObject(&tagResponse)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(&tagResponse, nil)

	return services
}

func TestBatchJobQueues(t *testing.T) {
	client.AwsMockTestHelper(t, JobQueues(), buildBatchJobQueuesMock, client.TestOptions{})
}
