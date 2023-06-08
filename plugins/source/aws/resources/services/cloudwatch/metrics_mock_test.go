package cloudwatch

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/tableoptions"
	"github.com/cloudquery/plugin-sdk/v3/faker"
	"github.com/golang/mock/gomock"
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

	m.EXPECT().ListMetrics(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&cloudwatch.ListMetricsOutput{
			Metrics: []types.Metric{a},
		}, nil)

	return services
}

func TestCloudwatchMetrics(t *testing.T) {
	client.AwsMockTestHelper(t, Metrics(), buildCloudWatchMetricsMock, client.TestOptions{
		TableOptions: tableoptions.TableOptions{
			CloudwatchMetrics: &tableoptions.CloudwatchMetrics{
				ListMetricsOpts: []tableoptions.CustomCloudwatchListMetricsInput{{
					ListMetricsInput: cloudwatch.ListMetricsInput{
						MetricName: aws.String("foo"),
						Namespace:  aws.String("bar"),
					},
				}},
			},
		},
	})
}
