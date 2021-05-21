package resources

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildApigatewayApiKeys(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockApigatewayClient(ctrl)

	a := types.ApiKey{}
	err := faker.FakeData(&a)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetApiKeys(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apigateway.GetApiKeysOutput{
			Items: []types.ApiKey{a},
		}, nil)

	return client.Services{
		Apigateway: m,
	}
}

func TestApigatewayAPIKeys(t *testing.T) {
	awsTestHelper(t, ApigatewayAPIKeys(), buildApigatewayApiKeys)
}
