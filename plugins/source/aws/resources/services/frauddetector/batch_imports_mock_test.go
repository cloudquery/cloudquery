package frauddetector

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/frauddetector"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v2/faker"
	"github.com/golang/mock/gomock"
)

func buildBatchImports(t *testing.T, ctrl *gomock.Controller) client.Services {
	fdClient := mocks.NewMockFrauddetectorClient(ctrl)

	data := types.BatchImport{}
	err := faker.FakeObject(&data)
	if err != nil {
		t.Fatal(err)
	}

	fdClient.EXPECT().GetBatchImportJobs(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&frauddetector.GetBatchImportJobsOutput{BatchImports: []types.BatchImport{data}}, nil,
	)

	return client.Services{
		Frauddetector: fdClient,
	}
}

func TestBatchImports(t *testing.T) {
	client.AwsMockTestHelper(t, BatchImports(), buildBatchImports, client.TestOptions{})
}
