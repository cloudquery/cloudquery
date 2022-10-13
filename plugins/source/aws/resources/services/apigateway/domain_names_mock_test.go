package apigateway

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildApigatewayDomainNames(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockApigatewayClient(ctrl)

	dm := types.DomainName{}
	err := faker.FakeObject(&dm)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetDomainNames(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apigateway.GetDomainNamesOutput{
			Items: []types.DomainName{dm},
		}, nil)

	bpm := types.BasePathMapping{}
	err = faker.FakeObject(&bpm)
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

func TestDomainNames(t *testing.T) {
	client.AwsMockTestHelper(t, DomainNames(), buildApigatewayDomainNames, client.TestOptions{})
}
