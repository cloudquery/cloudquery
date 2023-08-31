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

func buildApigatewayVpcLinks(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockApigatewayClient(ctrl)

	l := types.VpcLink{}
	require.NoError(t, faker.FakeObject(&l))

	m.EXPECT().GetVpcLinks(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apigateway.GetVpcLinksOutput{
			Items: []types.VpcLink{l},
		}, nil)

	return client.Services{
		Apigateway: m,
	}
}

func TestVpcLinks(t *testing.T) {
	client.AwsMockTestHelper(t, VpcLinks(), buildApigatewayVpcLinks, client.TestOptions{})
}
