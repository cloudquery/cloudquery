package ssm

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildParameters(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockSSMClient(ctrl)
	var pm types.ParameterMetadata
	if err := faker.FakeData(&pm); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().DescribeParameters(
		gomock.Any(),
		&ssm.DescribeParametersInput{},
	).Return(
		&ssm.DescribeParametersOutput{Parameters: []types.ParameterMetadata{pm}},
		nil,
	)
	return client.Services{SSM: mock}
}

func TestParameters(t *testing.T) {
	client.AwsMockTestHelper(t, Parameters(), buildParameters, client.TestOptions{})
}
