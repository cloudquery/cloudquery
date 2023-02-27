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

func buildCloudwatchLogsLogGroupsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockCloudwatchlogsClient(ctrl)
	l := types.LogGroup{}
	err := faker.FakeObject(&l)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeLogGroups(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&cloudwatchlogs.DescribeLogGroupsOutput{
			LogGroups: []types.LogGroup{l},
		}, nil)

	sf := types.SubscriptionFilter{}
	if err = faker.FakeObject(&sf); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeSubscriptionFilters(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&cloudwatchlogs.DescribeSubscriptionFiltersOutput{
			SubscriptionFilters: []types.SubscriptionFilter{sf},
		}, nil)

	tags := &cloudwatchlogs.ListTagsLogGroupOutput{}
	err = faker.FakeObject(tags)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListTagsLogGroup(gomock.Any(), gomock.Any(), gomock.Any()).Return(tags, nil)

	return client.Services{
		Cloudwatchlogs: m,
	}
}

func TestCloudwatchlogsLogGroups(t *testing.T) {
	client.AwsMockTestHelper(t, LogGroups(), buildCloudwatchLogsLogGroupsMock, client.TestOptions{})
}
