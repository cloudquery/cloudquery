package costexplorer

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v3/faker"
	"github.com/golang/mock/gomock"
)

func build30DayForecast(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockCostexplorerClient(ctrl)

	gcfo := costexplorer.GetCostForecastOutput{}
	err := faker.FakeObject(&gcfo)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetCostForecast(gomock.Any(), gomock.Any(), gomock.Any()).MinTimes(1).Return(
		&gcfo, nil)

	return client.Services{
		Costexplorer: m,
	}
}

func TestCostExplorerThirtyDayCostForecast(t *testing.T) {
	client.AwsMockTestHelper(t, ThirtyDayCostForecast(), build30DayForecast, client.TestOptions{})
}
