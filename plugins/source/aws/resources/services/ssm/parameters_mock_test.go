package ssm

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildParameters(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockSsmClient(ctrl)
	var pm types.ParameterMetadata
	require.NoError(t, faker.FakeObject(&pm))

	mock.EXPECT().DescribeParameters(
		gomock.Any(),
		&ssm.DescribeParametersInput{},
		gomock.Any(),
	).Return(
		&ssm.DescribeParametersOutput{Parameters: []types.ParameterMetadata{pm}},
		nil,
	)
	return client.Services{Ssm: mock}
}

func TestParameters(t *testing.T) {
	client.AwsMockTestHelper(t, Parameters(), buildParameters, client.TestOptions{})
}
