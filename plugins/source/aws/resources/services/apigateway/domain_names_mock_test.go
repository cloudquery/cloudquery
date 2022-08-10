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

func buildApigatewayDomainNames(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockApigatewayClient(ctrl)

	dm := types.DomainName{}
	err := faker.FakeData(&dm)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetDomainNames(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apigateway.GetDomainNamesOutput{
			Items: []types.DomainName{dm},
		}, nil)

	bpm := types.BasePathMapping{}
	err = faker.FakeData(&bpm)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetBasePathMappings(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apigateway.GetBasePathMappingsOutput{
			Items: []types.BasePathMapping{bpm},
		}, nil)

	return client.Services{
		Apigateway: m,
	}
}

func TestApigatewayDomainNames(t *testing.T) {
	client.AwsMockTestHelper(t, ApigatewayDomainNames(), buildApigatewayDomainNames, client.TestOptions{})
}
