package cloudwatchlogs

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildMetricFiltersMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockCloudwatchlogsClient(ctrl)
	l := types.MetricFilter{}
	err := faker.FakeObject(&l)
	if err != nil {
		t.Fatal(err)
	}
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
