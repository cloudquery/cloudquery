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

func buildApigatewayUsagePlans(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockApigatewayClient(ctrl)

	u := types.UsagePlan{}
	err := faker.FakeData(&u)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetUsagePlans(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apigateway.GetUsagePlansOutput{
			Items: []types.UsagePlan{u},
		}, nil)

	uk := types.UsagePlanKey{}
	err = faker.FakeData(&uk)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetUsagePlanKeys(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apigateway.GetUsagePlanKeysOutput{
			Items: []types.UsagePlanKey{uk},
		}, nil)

	return client.Services{
		Apigateway: m,
	}
}

func TestApigatewayUsagePlans(t *testing.T) {
	client.AwsMockTestHelper(t, ApigatewayUsagePlans(), buildApigatewayUsagePlans, client.TestOptions{})
}
