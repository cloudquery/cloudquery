package apigatewayv2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildApigatewayv2DomainNames(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockApigatewayv2Client(ctrl)

	dn := types.DomainName{}
	require.NoError(t, faker.FakeObject(&dn))
	m.EXPECT().GetDomainNames(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apigatewayv2.GetDomainNamesOutput{
			Items: []types.DomainName{dn},
		}, nil)

	am := types.ApiMapping{}
	require.NoError(t, faker.FakeObject(&am))
	m.EXPECT().GetApiMappings(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apigatewayv2.GetApiMappingsOutput{
			Items: []types.ApiMapping{am},
		}, nil)

	return client.Services{
		Apigatewayv2: m,
	}
}

func TestApigatewayv2DomainNames(t *testing.T) {
	client.AwsMockTestHelper(t, DomainNames(), buildApigatewayv2DomainNames, client.TestOptions{})
}
