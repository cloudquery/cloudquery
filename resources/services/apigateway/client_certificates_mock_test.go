//go:build mock
// +build mock

package apigateway

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildApigatewayClientCertificates(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockApigatewayClient(ctrl)

	c := types.ClientCertificate{}
	err := faker.FakeData(&c)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetClientCertificates(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apigateway.GetClientCertificatesOutput{
			Items: []types.ClientCertificate{c},
		}, nil)

	return client.Services{
		Apigateway: m,
	}
}

func TestApigatewayClientCertificates(t *testing.T) {
	client.AwsMockTestHelper(t, ApigatewayClientCertificates(), buildApigatewayClientCertificates, client.TestOptions{})
}
