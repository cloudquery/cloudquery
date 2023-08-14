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

func build30DayCost(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockCostexplorerClient(ctrl)

	gcuo := costexplorer.GetCostAndUsageOutput{}
	require.NoError(t, faker.FakeObject(&gcuo))
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
