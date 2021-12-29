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

func buildApigatewayVpcLinks(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockApigatewayClient(ctrl)

	l := types.VpcLink{}
	err := faker.FakeData(&l)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetVpcLinks(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apigateway.GetVpcLinksOutput{
			Items: []types.VpcLink{l},
		}, nil)

	return client.Services{
		Apigateway: m,
	}
}

func TestApigatewayVpcLinks(t *testing.T) {
	client.AwsMockTestHelper(t, ApigatewayVpcLinks(), buildApigatewayVpcLinks, client.TestOptions{})
}
