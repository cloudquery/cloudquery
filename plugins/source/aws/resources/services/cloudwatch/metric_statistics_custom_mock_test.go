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

func buildCloudWatchCustomMetricStatisticsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
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

func TestCloudwatchCustomMetricStatistics(t *testing.T) {
	client.AwsMockTestHelper(t, MetricStatisticsCustom(), buildCloudWatchCustomMetricStatisticsMock, client.TestOptions{
		TableOptions: tableoptions.TableOptions{
			CloudwatchCustomMetricStats: &tableoptions.CloudwatchCustomMetricStatistics{
				GetMetricStatisticsOpts: []tableoptions.CloudwatchGetMetricStatisticsInput{{
					GetMetricStatisticsInput: cloudwatch.GetMetricStatisticsInput{
						MetricName: aws.String("foo"),
						Namespace:  aws.String("bar"),
					},
				}},
			},
		},
	})
}
