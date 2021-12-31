// +build mock

package cloudwatchlogs

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	"github.com/bxcodec/faker"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/golang/mock/gomock"
)

func buildCloudwatchLogsFiltersMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockCloudwatchLogsClient(ctrl)
	l := types.MetricFilter{}
	err := faker.FakeData(&l)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeMetricFilters(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&cloudwatchlogs.DescribeMetricFiltersOutput{
			MetricFilters: []types.MetricFilter{l},
		}, nil)
	return client.Services{
		CloudwatchLogs: m,
	}
}

func TestAutoscalingLaunchConfiguration(t *testing.T) {
	client.AwsMockTestHelper(t, CloudwatchlogsFilters(), buildCloudwatchLogsFiltersMock, client.TestOptions{})
}
