package costexplorer

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v2/faker"
	"github.com/golang/mock/gomock"
)

func build30DayCost(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockCostexplorerClient(ctrl)

	gcuo := costexplorer.GetCostAndUsageOutput{}
	err := faker.FakeObject(&gcuo)
	if err != nil {
		t.Fatal(err)
	}
	gcuo.NextPageToken = nil
	m.EXPECT().GetCostAndUsage(gomock.Any(), gomock.Any(), gomock.Any()).MinTimes(1).Return(
		&gcuo, nil)

	return client.Services{
		Costexplorer: m,
	}
}

func TestCostExplorerCurrentMonth(t *testing.T) {
	client.AwsMockTestHelper(t, ThirtyDayCost(), build30DayCost, client.TestOptions{})
}
