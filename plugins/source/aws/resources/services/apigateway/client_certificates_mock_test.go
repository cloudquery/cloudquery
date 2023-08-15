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

func buildApigatewayClientCertificates(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockApigatewayClient(ctrl)

	c := types.ClientCertificate{}
	require.NoError(t, faker.FakeObject(&c))

	m.EXPECT().GetClientCertificates(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apigateway.GetClientCertificatesOutput{
			Items: []types.ClientCertificate{c},
		}, nil)

	return client.Services{
		Apigateway: m,
	}
}

func TestClientCertificates(t *testing.T) {
	client.AwsMockTestHelper(t, ClientCertificates(), buildApigatewayClientCertificates, client.TestOptions{})
}
