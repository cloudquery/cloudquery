package cloudwatch

import (
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/tableoptions"
	"github.com/cloudquery/plugin-sdk/v3/faker"
	"github.com/golang/mock/gomock"
)

var (
	st = time.Now().Add(-time.Hour * 24 * 7)
	et = time.Now()
)

func buildCloudWatchMetricsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockCloudwatchClient(ctrl)
	services := client.Services{
		Cloudwatch: m,
	}

	a := types.Metric{}
	if err := faker.FakeObject(&a); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListMetrics(gomock.Any(), gomock.Any(), gomock.Any()).Return(&cloudwatch.ListMetricsOutput{
		Metrics: []types.Metric{a},
	}, nil)

	o := cloudwatch.GetMetricStatisticsOutput{}
	if err := faker.FakeObject(&o); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().GetMetricStatistics(gomock.Any(), &cloudwatch.GetMetricStatisticsInput{
		MetricName: a.MetricName,
		Namespace:  a.Namespace,
		Dimensions: a.Dimensions,
		StartTime:  &st,
		EndTime:    &et,
	}, gomock.Any()).Return(&o, nil)

	return services
}

func TestCloudwatchMetrics(t *testing.T) {
	client.AwsMockTestHelper(t, Metrics(), buildCloudWatchMetricsMock, client.TestOptions{
		TableOptions: tableoptions.TableOptions{
			CloudwatchMetrics: []tableoptions.CloudwatchMetrics{
				{
					ListMetricsOpts: tableoptions.CloudwatchListMetricsInput{},
					GetMetricStatisticsOpts: []tableoptions.CloudwatchGetMetricStatisticsInput{{
						GetMetricStatisticsInput: cloudwatch.GetMetricStatisticsInput{
							StartTime: &st,
							EndTime:   &et,
						},
					}},
				},
			},
		},
	})
}
