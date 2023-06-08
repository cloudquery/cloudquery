package cloudwatch

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/tableoptions"
	"github.com/cloudquery/plugin-sdk/v3/faker"
	"github.com/golang/mock/gomock"
)

func buildCloudWatchMetricStatsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockCloudwatchClient(ctrl)
	services := client.Services{
		Cloudwatch: m,
	}

	a := cloudwatch.GetMetricStatisticsOutput{}
	if err := faker.FakeObject(&a); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().GetMetricStatistics(gomock.Any(), gomock.Any(), gomock.Any()).Return(&a, nil)

	return services
}

func TestCloudwatchMetricStats(t *testing.T) {
	client.AwsMockTestHelper(t, MetricStatistics(), buildCloudWatchMetricStatsMock, client.TestOptions{
		TableOptions: tableoptions.TableOptions{
			CloudwatchMetricStats: &tableoptions.CloudwatchMetricStatistics{
				GetMetricStatisticsOpts: []tableoptions.CustomCloudwatchGetMetricStatisticsInput{{
					GetMetricStatisticsInput: cloudwatch.GetMetricStatisticsInput{
						MetricName: aws.String("foo"),
						Namespace:  aws.String("bar"),
					},
				}},
			},
		},
	})
}
