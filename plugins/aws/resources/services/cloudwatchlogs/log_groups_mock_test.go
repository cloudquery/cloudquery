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

func buildCloudwatchLogsLogGroupsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockCloudwatchLogsClient(ctrl)
	l := types.LogGroup{}
	err := faker.FakeData(&l)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeLogGroups(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&cloudwatchlogs.DescribeLogGroupsOutput{
			LogGroups: []types.LogGroup{l},
		}, nil)

	tags := &cloudwatchlogs.ListTagsLogGroupOutput{}
	err = faker.FakeData(tags)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListTagsLogGroup(gomock.Any(), gomock.Any(), gomock.Any()).Return(tags, nil)

	return client.Services{
		CloudwatchLogs: m,
	}
}

func TestCloudwatchlogsLogGroups(t *testing.T) {
	client.AwsMockTestHelper(t, LogGroups(), buildCloudwatchLogsLogGroupsMock, client.TestOptions{})
}
