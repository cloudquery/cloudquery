package frauddetector

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/frauddetector"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildModelVersions(t *testing.T, client *mocks.MockFrauddetectorClient) {
	data := types.ModelVersionDetail{}
	err := faker.FakeObject(&data)
	if err != nil {
		t.Fatal(err)
	}

	client.EXPECT().DescribeModelVersions(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&frauddetector.DescribeModelVersionsOutput{ModelVersionDetails: []types.ModelVersionDetail{data}}, nil,
	)
}
