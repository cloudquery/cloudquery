package cloudwatch

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildCloudWatchAlarmsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockCloudwatchClient(ctrl)
	services := client.Services{
		Cloudwatch: m,
	}
	a := types.MetricAlarm{}
	err := faker.FakeObject(&a)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeAlarms(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&cloudwatch.DescribeAlarmsOutput{
			MetricAlarms: []types.MetricAlarm{a},
		}, nil)

	tagResponse := cloudwatch.ListTagsForResourceOutput{}
	err = faker.FakeObject(&tagResponse)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(&tagResponse, nil)

	return services
}

func TestCloudwatchAlarms(t *testing.T) {
	client.AwsMockTestHelper(t, Alarms(), buildCloudWatchAlarmsMock, client.TestOptions{})
}
