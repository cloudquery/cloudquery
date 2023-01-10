package scheduler

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/scheduler"
	"github.com/aws/aws-sdk-go-v2/service/scheduler/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildSchedulerScheduleGroupsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockSchedulerClient(ctrl)
	object := types.ScheduleGroupSummary{}
	err := faker.FakeObject(&object)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListScheduleGroups(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&scheduler.ListScheduleGroupsOutput{
			ScheduleGroups: []types.ScheduleGroupSummary{object},
		}, nil)

	tagsOutput := scheduler.ListTagsForResourceOutput{}
	err = faker.FakeObject(&tagsOutput)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any()).Return(&tagsOutput, nil).AnyTimes()
	return client.Services{
		Scheduler: m,
	}
}
func TestSchedulerSchedulerGroups(t *testing.T) {
	client.AwsMockTestHelper(t, ScheduleGroups(), buildSchedulerScheduleGroupsMock, client.TestOptions{})
}
