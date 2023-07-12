package cloudwatchlogs

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildMetricFiltersMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockCloudwatchlogsClient(ctrl)
	l := types.MetricFilter{}
	require.NoError(t, faker.FakeObject(&l))
	m.EXPECT().DescribeMetricFilters(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&cloudwatchlogs.DescribeMetricFiltersOutput{
			MetricFilters: []types.MetricFilter{l},
		}, nil)
	return client.Services{
		Cloudwatchlogs: m,
	}
}

func TestCloudwatchlogsFilter(t *testing.T) {
	client.AwsMockTestHelper(t, MetricFilters(), buildMetricFiltersMock, client.TestOptions{})
}
