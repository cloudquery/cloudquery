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

func buildApigatewayClientCertificates(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockApigatewayClient(ctrl)

	с := types.ClientCertificate{}
	err := faker.FakeData(&с)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetClientCertificates(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apigateway.GetClientCertificatesOutput{
			Items: []types.ClientCertificate{с},
		}, nil)

	return client.Services{
		Apigateway: m,
	}
}

func TestApigatewayClientCertificates(t *testing.T) {
	awsTestHelper(t, ApigatewayClientCertificates(), buildApigatewayClientCertificates, TestOptions{})
}
