package batch

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/batch"
	"github.com/aws/aws-sdk-go-v2/service/batch/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v2/faker"
	"github.com/golang/mock/gomock"
)

func buildBatchJobDefinitionsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockBatchClient(ctrl)
	services := client.Services{
		Batch: m,
	}
	a := types.JobDefinition{}
	err := faker.FakeObject(&a)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeJobDefinitions(gomock.Any(), gomock.Any()).Return(
		&batch.DescribeJobDefinitionsOutput{
			JobDefinitions: []types.JobDefinition{a},
		}, nil)

	tagResponse := batch.ListTagsForResourceOutput{}
	err = faker.FakeObject(&tagResponse)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(&tagResponse, nil)

	return services
}

func TestBatchJobDefinitions(t *testing.T) {
	client.AwsMockTestHelper(t, JobDefinitions(), buildBatchJobDefinitionsMock, client.TestOptions{})
}
