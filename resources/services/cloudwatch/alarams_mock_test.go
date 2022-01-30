package cloudwatch

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildCloudWatchAlarmsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockCloudwatchClient(ctrl)
	services := client.Services{
		Cloudwatch: m,
	}
	a := types.MetricAlarm{}
	err := faker.FakeData(&a)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeAlarms(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&cloudwatch.DescribeAlarmsOutput{
			MetricAlarms: []types.MetricAlarm{a},
		}, nil)
	return services
}

func TestCloudwatchAlarms(t *testing.T) {
	client.AwsMockTestHelper(t, CloudwatchAlarms(), buildCloudWatchAlarmsMock, client.TestOptions{})
}
