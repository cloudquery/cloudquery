package apigatewayv2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildApigatewayv2VpcLinks(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockApigatewayv2Client(ctrl)

	v := types.VpcLink{}
	err := faker.FakeObject(&v)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetVpcLinks(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apigatewayv2.GetVpcLinksOutput{
			Items: []types.VpcLink{v},
		}, nil)

	return client.Services{
		Apigatewayv2: m,
	}
}

func TestApigatewayv2VpcLinks(t *testing.T) {
	client.AwsMockTestHelper(t, VpcLinks(), buildApigatewayv2VpcLinks, client.TestOptions{})
}
