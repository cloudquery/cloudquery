package costexplorer

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func build30DayForecast(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockCostexplorerClient(ctrl)

	gcfo := costexplorer.GetCostForecastOutput{}
	require.NoError(t, faker.FakeObject(&gcfo))
	m.EXPECT().GetCostForecast(gomock.Any(), gomock.Any(), gomock.Any()).MinTimes(1).Return(
		&gcfo, nil)

	return client.Services{
		Costexplorer: m,
	}
}

func TestCostExplorerThirtyDayCostForecast(t *testing.T) {
	client.AwsMockTestHelper(t, ThirtyDayCostForecast(), build30DayForecast, client.TestOptions{})
}
