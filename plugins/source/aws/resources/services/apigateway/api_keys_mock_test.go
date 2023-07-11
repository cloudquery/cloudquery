package apigateway

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildApiKeysMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockApigatewayClient(ctrl)

	a := types.ApiKey{}
	require.NoError(t, faker.FakeObject(&a))

	m.EXPECT().GetApiKeys(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apigateway.GetApiKeysOutput{
			Items: []types.ApiKey{a},
		}, nil)

	return client.Services{
		Apigateway: m,
	}
}

func TestAPIKeys(t *testing.T) {
	client.AwsMockTestHelper(t, ApiKeys(), buildApiKeysMock, client.TestOptions{})
}
